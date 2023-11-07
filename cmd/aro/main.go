package main

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	_ "net/http/pprof"

	utillog "github.com/Azure/ARO-RP/pkg/util/log"
	_ "github.com/Azure/ARO-RP/pkg/util/scheme"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/davecgh/go-spew/spew"

	utilgraph "github.com/Azure/ARO-RP/pkg/util/graph"
	msgraph_models "github.com/Azure/ARO-RP/pkg/util/graph/graphsdk/models"
	msgraph_errors "github.com/Azure/ARO-RP/pkg/util/graph/graphsdk/models/odataerrors"
)

/*
func usage() {
	fmt.Fprint(flag.CommandLine.Output(), "usage:\n")
	fmt.Fprintf(flag.CommandLine.Output(), "  %s dbtoken\n", os.Args[0])
	fmt.Fprintf(flag.CommandLine.Output(), "  %s deploy config.yaml location\n", os.Args[0])
	fmt.Fprintf(flag.CommandLine.Output(), "  %s gateway\n", os.Args[0])
	fmt.Fprintf(flag.CommandLine.Output(), "  %s mirror [release_image...]\n", os.Args[0])
	fmt.Fprintf(flag.CommandLine.Output(), "  %s monitor\n", os.Args[0])
	fmt.Fprintf(flag.CommandLine.Output(), "  %s portal\n", os.Args[0])
	fmt.Fprintf(flag.CommandLine.Output(), "  %s rp\n", os.Args[0])
	fmt.Fprintf(flag.CommandLine.Output(), "  %s operator {master,worker}\n", os.Args[0])
	fmt.Fprintf(flag.CommandLine.Output(), "  %s update-versions\n", os.Args[0])
	flag.PrintDefaults()
}
*/

func main() {

	log := utillog.GetLogger()
	log.Print("************* niontive test")

	// Assume env set
	spTokenCredential, err := azidentity.NewEnvironmentCredential(nil)
	if err != nil {
		log.Fatalf("failed to get environment credential: %s", err)
	}

	// Inspired by pkg/util/azureclient/environments.go
	client, err := utilgraph.NewGraphServiceClientWithCredentials(spTokenCredential, nil)
	if err != nil {
		log.Fatalf("failed to create graph client: %s", err)
	}
	client.GetAdapter().SetBaseUrl("https://graph.microsoft.com/" + "v1.0")

	// Inspired by pkg/util/cluster/aad.go
	displayName := "niontive-e2e-test"
	appBody := msgraph_models.NewApplication()
	appBody.SetDisplayName(&displayName)

	log.Printf("appBody display name: %s", *appBody.GetDisplayName())
	appResult, err := client.Applications().Post(context.Background(), appBody, nil)
	if err != nil {
		if oDataError, ok := err.(msgraph_errors.ODataErrorable); ok {
			spew.Dump(oDataError.GetErrorEscaped())
		}
		log.Fatalf("failed to create application: %s", err)
	}

	log.Printf("appResult: %v", appResult)

	log.Print("************* niontive test done!")

	/*
		rand.Seed(time.Now().UnixNano())

		flag.Usage = usage
		flag.Parse()

		ctx := context.Background()
		audit := utillog.GetAuditEntry()
		log := utillog.GetLogger()

		go func() {
			log.Warn(http.ListenAndServe("localhost:6060", nil))
		}()

		log.Printf("starting, git commit %s", version.GitCommit)

		var err error
		switch strings.ToLower(flag.Arg(0)) {
		case "dbtoken":
			checkArgs(1)
			err = dbtoken(ctx, log)
		case "deploy":
			checkArgs(3)
			err = deploy(ctx, log)
		case "gateway":
			checkArgs(1)
			err = gateway(ctx, log)
		case "mirror":
			checkMinArgs(1)
			err = mirror(ctx, log)
		case "monitor":
			checkArgs(1)
			err = monitor(ctx, log)
		case "rp":
			checkArgs(1)
			err = rp(ctx, log, audit)
		case "portal":
			checkArgs(1)
			err = portal(ctx, log, audit)
		case "operator":
			checkArgs(2)
			err = operator(ctx, log)
		case "update-versions":
			checkArgs(1)
			err = updateOCPVersions(ctx, log)
		default:
			usage()
			os.Exit(2)
		}

		if err != nil {
			log.Fatal(err)
		}
	*/
}

/*
func checkArgs(required int) {
	if len(flag.Args()) != required {
		usage()
		os.Exit(2)
	}
}

func checkMinArgs(required int) {
	if len(flag.Args()) < required {
		usage()
		os.Exit(2)
	}
}

func DBName(isLocalDevelopmentMode bool) (string, error) {
	if !isLocalDevelopmentMode {
		return "ARO", nil
	}

	if err := env.ValidateVars(DatabaseName); err != nil {
		return "", fmt.Errorf("%v (development mode)", err.Error())
	}

	return os.Getenv(DatabaseName), nil
}
*/
