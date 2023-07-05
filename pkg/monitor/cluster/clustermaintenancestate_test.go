package cluster

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"testing"

	"github.com/Azure/ARO-RP/pkg/api"
	"github.com/golang/mock/gomock"
	
	mock_metrics "github.com/Azure/ARO-RP/pkg/util/mocks/metrics"
)

func TestEmitClusterMaintenanceState(t *testing.T) {
	ctx := context.Background()

	controller := gomock.NewController(t)
	defer controller.Finish()

	m := mock_metrics.NewMockEmitter(controller)

	mon := &Monitor{
		m: m,
		oc: &api.OpenShiftCluster{
			Properties: api.OpenShiftClusterProperties{
				MaintenanceState: api.MaintenanceStateNone,
			},
		},
	}

	m.EXPECT().EmitGauge("cluster.maintenanceState", int64(1), map[string]string{
		"maintenanceState": "None",
	})

	err := mon.emitClusterMaintenanceState(ctx)
	if err != nil {
		t.Fatal(err)
	}
}
