package sites

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksResponse 
// Deprecated: This class is obsolete. Use getRecentNotebooksWithIncludePersonalNotebooksGetResponse instead.
type ItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksResponse struct {
    ItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksGetResponse
}
// NewItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksResponse instantiates a new ItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksResponse and sets the default values.
func NewItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksResponse()(*ItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksResponse) {
    m := &ItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksResponse{
        ItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksGetResponse: *NewItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksGetResponse(),
    }
    return m
}
// CreateItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksResponse(), nil
}
// ItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksResponseable 
// Deprecated: This class is obsolete. Use getRecentNotebooksWithIncludePersonalNotebooksGetResponse instead.
type ItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksResponseable interface {
    ItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
