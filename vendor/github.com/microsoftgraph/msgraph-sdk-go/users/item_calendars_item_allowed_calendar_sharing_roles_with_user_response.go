package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse 
// Deprecated: This class is obsolete. Use allowedCalendarSharingRolesWithUserGetResponse instead.
type ItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse struct {
    ItemCalendarsItemAllowedCalendarSharingRolesWithUserGetResponse
}
// NewItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse instantiates a new ItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse and sets the default values.
func NewItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse()(*ItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse) {
    m := &ItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse{
        ItemCalendarsItemAllowedCalendarSharingRolesWithUserGetResponse: *NewItemCalendarsItemAllowedCalendarSharingRolesWithUserGetResponse(),
    }
    return m
}
// CreateItemCalendarsItemAllowedCalendarSharingRolesWithUserResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCalendarsItemAllowedCalendarSharingRolesWithUserResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse(), nil
}
// ItemCalendarsItemAllowedCalendarSharingRolesWithUserResponseable 
// Deprecated: This class is obsolete. Use allowedCalendarSharingRolesWithUserGetResponse instead.
type ItemCalendarsItemAllowedCalendarSharingRolesWithUserResponseable interface {
    ItemCalendarsItemAllowedCalendarSharingRolesWithUserGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
