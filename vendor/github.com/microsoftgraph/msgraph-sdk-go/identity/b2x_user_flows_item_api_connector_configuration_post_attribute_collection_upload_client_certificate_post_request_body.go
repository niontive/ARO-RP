package identity

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// B2xUserFlowsItemApiConnectorConfigurationPostAttributeCollectionUploadClientCertificatePostRequestBody 
type B2xUserFlowsItemApiConnectorConfigurationPostAttributeCollectionUploadClientCertificatePostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewB2xUserFlowsItemApiConnectorConfigurationPostAttributeCollectionUploadClientCertificatePostRequestBody instantiates a new B2xUserFlowsItemApiConnectorConfigurationPostAttributeCollectionUploadClientCertificatePostRequestBody and sets the default values.
func NewB2xUserFlowsItemApiConnectorConfigurationPostAttributeCollectionUploadClientCertificatePostRequestBody()(*B2xUserFlowsItemApiConnectorConfigurationPostAttributeCollectionUploadClientCertificatePostRequestBody) {
    m := &B2xUserFlowsItemApiConnectorConfigurationPostAttributeCollectionUploadClientCertificatePostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateB2xUserFlowsItemApiConnectorConfigurationPostAttributeCollectionUploadClientCertificatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateB2xUserFlowsItemApiConnectorConfigurationPostAttributeCollectionUploadClientCertificatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewB2xUserFlowsItemApiConnectorConfigurationPostAttributeCollectionUploadClientCertificatePostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *B2xUserFlowsItemApiConnectorConfigurationPostAttributeCollectionUploadClientCertificatePostRequestBody) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *B2xUserFlowsItemApiConnectorConfigurationPostAttributeCollectionUploadClientCertificatePostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *B2xUserFlowsItemApiConnectorConfigurationPostAttributeCollectionUploadClientCertificatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["password"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPassword(val)
        }
        return nil
    }
    res["pkcs12Value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPkcs12Value(val)
        }
        return nil
    }
    return res
}
// GetPassword gets the password property value. The password property
func (m *B2xUserFlowsItemApiConnectorConfigurationPostAttributeCollectionUploadClientCertificatePostRequestBody) GetPassword()(*string) {
    val, err := m.GetBackingStore().Get("password")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPkcs12Value gets the pkcs12Value property value. The pkcs12Value property
func (m *B2xUserFlowsItemApiConnectorConfigurationPostAttributeCollectionUploadClientCertificatePostRequestBody) GetPkcs12Value()(*string) {
    val, err := m.GetBackingStore().Get("pkcs12Value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *B2xUserFlowsItemApiConnectorConfigurationPostAttributeCollectionUploadClientCertificatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("password", m.GetPassword())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("pkcs12Value", m.GetPkcs12Value())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *B2xUserFlowsItemApiConnectorConfigurationPostAttributeCollectionUploadClientCertificatePostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *B2xUserFlowsItemApiConnectorConfigurationPostAttributeCollectionUploadClientCertificatePostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetPassword sets the password property value. The password property
func (m *B2xUserFlowsItemApiConnectorConfigurationPostAttributeCollectionUploadClientCertificatePostRequestBody) SetPassword(value *string)() {
    err := m.GetBackingStore().Set("password", value)
    if err != nil {
        panic(err)
    }
}
// SetPkcs12Value sets the pkcs12Value property value. The pkcs12Value property
func (m *B2xUserFlowsItemApiConnectorConfigurationPostAttributeCollectionUploadClientCertificatePostRequestBody) SetPkcs12Value(value *string)() {
    err := m.GetBackingStore().Set("pkcs12Value", value)
    if err != nil {
        panic(err)
    }
}
// B2xUserFlowsItemApiConnectorConfigurationPostAttributeCollectionUploadClientCertificatePostRequestBodyable 
type B2xUserFlowsItemApiConnectorConfigurationPostAttributeCollectionUploadClientCertificatePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetPassword()(*string)
    GetPkcs12Value()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetPassword(value *string)()
    SetPkcs12Value(value *string)()
}
