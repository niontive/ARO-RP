package devices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GetAvailableExtensionPropertiesResponse 
// Deprecated: This class is obsolete. Use getAvailableExtensionPropertiesPostResponse instead.
type GetAvailableExtensionPropertiesResponse struct {
    GetAvailableExtensionPropertiesPostResponse
}
// NewGetAvailableExtensionPropertiesResponse instantiates a new GetAvailableExtensionPropertiesResponse and sets the default values.
func NewGetAvailableExtensionPropertiesResponse()(*GetAvailableExtensionPropertiesResponse) {
    m := &GetAvailableExtensionPropertiesResponse{
        GetAvailableExtensionPropertiesPostResponse: *NewGetAvailableExtensionPropertiesPostResponse(),
    }
    return m
}
// CreateGetAvailableExtensionPropertiesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGetAvailableExtensionPropertiesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetAvailableExtensionPropertiesResponse(), nil
}
// GetAvailableExtensionPropertiesResponseable 
// Deprecated: This class is obsolete. Use getAvailableExtensionPropertiesPostResponse instead.
type GetAvailableExtensionPropertiesResponseable interface {
    GetAvailableExtensionPropertiesPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
