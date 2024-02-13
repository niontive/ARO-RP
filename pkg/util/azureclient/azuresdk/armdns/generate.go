package armdns

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

//go:generate rm -rf ../../../../../pkg/util/mocks/azureclient/azuresdk/$GOPACKAGE
//go:generate go run ../../../../../vendor/github.com/golang/mock/mockgen -destination=../../../../../pkg/util/mocks/azureclient/azuresdk/$GOPACKAGE/armdns.go -source=armdns.go
//go:generate go run ../../../../../vendor/golang.org/x/tools/cmd/goimports -local=github.com/Azure/ARO-RP -e -w ../../../../../../pkg/util/mocks/azureclient/azuresdk/$GOPACKAGE/armdns.go
