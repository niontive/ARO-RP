package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EntitlementManagementAccessPackagesFilterByCurrentUserWithOnResponse 
// Deprecated: This class is obsolete. Use filterByCurrentUserWithOnGetResponse instead.
type EntitlementManagementAccessPackagesFilterByCurrentUserWithOnResponse struct {
    EntitlementManagementAccessPackagesFilterByCurrentUserWithOnGetResponse
}
// NewEntitlementManagementAccessPackagesFilterByCurrentUserWithOnResponse instantiates a new EntitlementManagementAccessPackagesFilterByCurrentUserWithOnResponse and sets the default values.
func NewEntitlementManagementAccessPackagesFilterByCurrentUserWithOnResponse()(*EntitlementManagementAccessPackagesFilterByCurrentUserWithOnResponse) {
    m := &EntitlementManagementAccessPackagesFilterByCurrentUserWithOnResponse{
        EntitlementManagementAccessPackagesFilterByCurrentUserWithOnGetResponse: *NewEntitlementManagementAccessPackagesFilterByCurrentUserWithOnGetResponse(),
    }
    return m
}
// CreateEntitlementManagementAccessPackagesFilterByCurrentUserWithOnResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEntitlementManagementAccessPackagesFilterByCurrentUserWithOnResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEntitlementManagementAccessPackagesFilterByCurrentUserWithOnResponse(), nil
}
// EntitlementManagementAccessPackagesFilterByCurrentUserWithOnResponseable 
// Deprecated: This class is obsolete. Use filterByCurrentUserWithOnGetResponse instead.
type EntitlementManagementAccessPackagesFilterByCurrentUserWithOnResponseable interface {
    EntitlementManagementAccessPackagesFilterByCurrentUserWithOnGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
