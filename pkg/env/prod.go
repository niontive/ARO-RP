package env

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"net"
	"os"
	"strings"

	mgmtcompute "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2020-06-01/compute"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/sirupsen/logrus"

	"github.com/Azure/ARO-RP/pkg/proxy"
	"github.com/Azure/ARO-RP/pkg/util/azureclient"
	"github.com/Azure/ARO-RP/pkg/util/azureclient/mgmt/compute"
	"github.com/Azure/ARO-RP/pkg/util/clientauthorizer"
	"github.com/Azure/ARO-RP/pkg/util/keyvault"
	"github.com/Azure/ARO-RP/pkg/util/refreshable"
	"github.com/Azure/ARO-RP/pkg/util/version"
)

type prod struct {
	Core
	proxy.Dialer
	ARMHelper

	armClientAuthorizer   clientauthorizer.ClientAuthorizer
	adminClientAuthorizer clientauthorizer.ClientAuthorizer

	acrDomain string
	vmskus    map[string]*mgmtcompute.ResourceSku

	fpCertificate *x509.Certificate
	fpPrivateKey  *rsa.PrivateKey
	fpClientID    string

	clusterKeyvault keyvault.Manager
	serviceKeyvault keyvault.Manager

	clusterGenevaLoggingCertificate   *x509.Certificate
	clusterGenevaLoggingPrivateKey    *rsa.PrivateKey
	clusterGenevaLoggingAccount       string
	clusterGenevaLoggingConfigVersion string
	clusterGenevaLoggingEnvironment   string
	clusterGenevaLoggingNamespace     string

	log *logrus.Entry

	features map[Feature]bool
}

func newProd(ctx context.Context, log *logrus.Entry) (*prod, error) {
	for _, key := range []string{
		"AZURE_FP_CLIENT_ID",
		"DOMAIN_NAME",
	} {
		if _, found := os.LookupEnv(key); !found {
			return nil, fmt.Errorf("environment variable %q unset", key)
		}
	}

	if !IsLocalDevelopmentMode() {
		for _, key := range []string{
			"CLUSTER_MDSD_CONFIG_VERSION",
			"CLUSTER_MDSD_ACCOUNT",
			"MDSD_ENVIRONMENT",
			"CLUSTER_MDSD_NAMESPACE",
		} {
			if _, found := os.LookupEnv(key); !found {
				return nil, fmt.Errorf("environment variable %q unset", key)
			}
		}
	}

	core, err := NewCore(ctx, log)
	if err != nil {
		return nil, err
	}

	dialer, err := proxy.NewDialer(core.IsLocalDevelopmentMode())
	if err != nil {
		return nil, err
	}

	p := &prod{
		Core:   core,
		Dialer: dialer,

		fpClientID: os.Getenv("AZURE_FP_CLIENT_ID"),

		clusterGenevaLoggingAccount:       os.Getenv("CLUSTER_MDSD_ACCOUNT"),
		clusterGenevaLoggingConfigVersion: os.Getenv("CLUSTER_MDSD_CONFIG_VERSION"),
		clusterGenevaLoggingEnvironment:   os.Getenv("MDSD_ENVIRONMENT"),
		clusterGenevaLoggingNamespace:     os.Getenv("CLUSTER_MDSD_NAMESPACE"),

		log: log,

		features: map[Feature]bool{},
	}

	features := os.Getenv("RP_FEATURES")
	if features != "" {
		for _, feature := range strings.Split(features, ",") {
			f, err := FeatureString("Feature" + feature)
			if err != nil {
				return nil, err
			}

			p.features[f] = true
		}
	}

	msiAuthorizer, err := p.NewMSIAuthorizer(MSIContextRP, p.Environment().ResourceManagerEndpoint)
	if err != nil {
		return nil, err
	}

	msiKVAuthorizer, err := p.NewMSIAuthorizer(MSIContextRP, p.Environment().ResourceIdentifiers.KeyVault)
	if err != nil {
		return nil, err
	}

	serviceKeyvaultURI, err := keyvault.URI(p, ServiceKeyvaultSuffix)
	if err != nil {
		return nil, err
	}

	p.serviceKeyvault = keyvault.NewManager(msiKVAuthorizer, serviceKeyvaultURI)

	resourceSkusClient := compute.NewResourceSkusClient(p.Environment(), p.SubscriptionID(), msiAuthorizer)
	err = p.populateVMSkus(ctx, resourceSkusClient)
	if err != nil {
		return nil, err
	}

	fpPrivateKey, fpCertificates, err := p.serviceKeyvault.GetCertificateSecret(ctx, RPFirstPartySecretName)
	if err != nil {
		return nil, err
	}

	p.fpPrivateKey = fpPrivateKey
	p.fpCertificate = fpCertificates[0]

	localFPKVAuthorizer, err := p.FPAuthorizer(p.TenantID(), p.Environment().ResourceIdentifiers.KeyVault)
	if err != nil {
		return nil, err
	}

	clusterKeyvaultURI, err := keyvault.URI(p, ClusterKeyvaultSuffix)
	if err != nil {
		return nil, err
	}

	p.clusterKeyvault = keyvault.NewManager(localFPKVAuthorizer, clusterKeyvaultURI)

	clusterGenevaLoggingPrivateKey, clusterGenevaLoggingCertificates, err := p.serviceKeyvault.GetCertificateSecret(ctx, ClusterLoggingSecretName)
	if err != nil {
		return nil, err
	}

	p.clusterGenevaLoggingPrivateKey = clusterGenevaLoggingPrivateKey
	p.clusterGenevaLoggingCertificate = clusterGenevaLoggingCertificates[0]

	if p.ACRResourceID() != "" { // TODO: ugh!
		acrResource, err := azure.ParseResourceID(p.ACRResourceID())
		if err != nil {
			return nil, err
		}
		p.acrDomain = acrResource.ResourceName + "." + p.Environment().ContainerRegistryDNSSuffix
	} else {
		p.acrDomain = "arointsvc" + "." + azureclient.PublicCloud.ContainerRegistryDNSSuffix // TODO: make cloud aware once this is set up for US Gov Cloud
	}

	p.ARMHelper, err = newARMHelper(ctx, log, p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (p *prod) InitializeAuthorizers() error {
	if !p.FeatureIsSet(FeatureEnableDevelopmentAuthorizer) {
		p.armClientAuthorizer = clientauthorizer.NewARM(p.log, p.Core)

	} else {
		armClientAuthorizer, err := clientauthorizer.NewSubjectNameAndIssuer(
			p.log,
			"/etc/aro-rp/arm-ca-bundle.pem",
			os.Getenv("ARM_API_CLIENT_CERT_COMMON_NAME"),
		)
		if err != nil {
			return err
		}

		p.armClientAuthorizer = armClientAuthorizer
	}

	adminClientAuthorizer, err := clientauthorizer.NewSubjectNameAndIssuer(
		p.log,
		"/etc/aro-rp/admin-ca-bundle.pem",
		os.Getenv("ADMIN_API_CLIENT_CERT_COMMON_NAME"),
	)
	if err != nil {
		return err
	}

	p.adminClientAuthorizer = adminClientAuthorizer
	return nil
}

func (p *prod) ArmClientAuthorizer() clientauthorizer.ClientAuthorizer {
	return p.armClientAuthorizer
}

func (p *prod) AdminClientAuthorizer() clientauthorizer.ClientAuthorizer {
	return p.adminClientAuthorizer
}

func (p *prod) ACRResourceID() string {
	return os.Getenv("ACR_RESOURCE_ID")
}

func (p *prod) ACRDomain() string {
	return p.acrDomain
}

func (p *prod) AROOperatorImage() string {
	return fmt.Sprintf("%s/aro:%s", p.acrDomain, version.GitCommit)
}

func (p *prod) populateVMSkus(ctx context.Context, resourceSkusClient compute.ResourceSkusClient) error {
	// Filtering is poorly documented, but currently (API version 2019-04-01)
	// it seems that the API returns all SKUs without a filter and with invalid
	// value in the filter.
	// Filtering gives significant optimisation: at the moment of writing,
	// we get ~1.2M response in eastus vs ~37M unfiltered (467 items vs 16618).
	filter := fmt.Sprintf("location eq '%s'", p.Location())
	skus, err := resourceSkusClient.List(ctx, filter)
	if err != nil {
		return err
	}

	p.vmskus = map[string]*mgmtcompute.ResourceSku{}
	for _, sku := range skus {
		// TODO(mjudeikis): At some point some SKU's stopped returning zones and
		// locations. IcM is open with MSFT but this might take a while.
		// Revert once we find out right behaviour.
		// https://github.com/Azure/ARO-RP/issues/1515
		if len(*sku.Locations) == 0 || !strings.EqualFold((*sku.Locations)[0], p.Location()) ||
			*sku.ResourceType != "virtualMachines" {
			continue
		}

		if len(*sku.LocationInfo) == 0 { // happened in eastus2euap
			continue
		}

		// We copy only part of the object so we don't have to keep
		// a lot of data in memory.
		p.vmskus[*sku.Name] = &mgmtcompute.ResourceSku{
			Name:         sku.Name,
			LocationInfo: sku.LocationInfo,
			Capabilities: sku.Capabilities,
		}
	}

	return nil
}

func (p *prod) ClusterGenevaLoggingAccount() string {
	return p.clusterGenevaLoggingAccount
}

func (p *prod) ClusterGenevaLoggingConfigVersion() string {
	return p.clusterGenevaLoggingConfigVersion
}

func (p *prod) ClusterGenevaLoggingEnvironment() string {
	return p.clusterGenevaLoggingEnvironment
}

func (p *prod) ClusterGenevaLoggingNamespace() string {
	return p.clusterGenevaLoggingNamespace
}

func (p *prod) ClusterGenevaLoggingSecret() (*rsa.PrivateKey, *x509.Certificate) {
	return p.clusterGenevaLoggingPrivateKey, p.clusterGenevaLoggingCertificate
}

func (p *prod) ClusterKeyvault() keyvault.Manager {
	return p.clusterKeyvault
}

func (p *prod) Domain() string {
	return os.Getenv("DOMAIN_NAME")
}

func (p *prod) FeatureIsSet(f Feature) bool {
	return p.features[f]
}

func (p *prod) FPAuthorizer(tenantID, resource string) (refreshable.Authorizer, error) {
	oauthConfig, err := adal.NewOAuthConfig(p.Environment().ActiveDirectoryEndpoint, tenantID)
	if err != nil {
		return nil, err
	}

	sp, err := adal.NewServicePrincipalTokenFromCertificate(*oauthConfig, p.fpClientID, p.fpCertificate, p.fpPrivateKey, resource)
	if err != nil {
		return nil, err
	}

	return refreshable.NewAuthorizer(sp), nil
}

func (p *prod) FPClientID() string {
	return p.fpClientID
}

func (p *prod) Listen() (net.Listener, error) {
	return net.Listen("tcp", ":8443")
}

func (p *prod) ServiceKeyvault() keyvault.Manager {
	return p.serviceKeyvault
}

func (p *prod) VMSku(vmSize string) (*mgmtcompute.ResourceSku, error) {
	vmsku, found := p.vmskus[vmSize]
	if !found {
		return nil, fmt.Errorf("sku information not found for vm size %q", vmSize)
	}
	return vmsku, nil
}
