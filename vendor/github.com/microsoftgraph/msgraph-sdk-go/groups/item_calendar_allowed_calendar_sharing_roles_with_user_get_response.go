package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemCalendarAllowedCalendarSharingRolesWithUserGetResponse 
type ItemCalendarAllowedCalendarSharingRolesWithUserGetResponse struct {
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BaseCollectionPaginationCountResponse
}
// NewItemCalendarAllowedCalendarSharingRolesWithUserGetResponse instantiates a new ItemCalendarAllowedCalendarSharingRolesWithUserGetResponse and sets the default values.
func NewItemCalendarAllowedCalendarSharingRolesWithUserGetResponse()(*ItemCalendarAllowedCalendarSharingRolesWithUserGetResponse) {
    m := &ItemCalendarAllowedCalendarSharingRolesWithUserGetResponse{
        BaseCollectionPaginationCountResponse: *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateItemCalendarAllowedCalendarSharingRolesWithUserGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCalendarAllowedCalendarSharingRolesWithUserGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarAllowedCalendarSharingRolesWithUserGetResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemCalendarAllowedCalendarSharingRolesWithUserGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ParseCalendarRoleType)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CalendarRoleType, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CalendarRoleType))
                }
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *ItemCalendarAllowedCalendarSharingRolesWithUserGetResponse) GetValue()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CalendarRoleType) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CalendarRoleType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ItemCalendarAllowedCalendarSharingRolesWithUserGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        err = writer.WriteCollectionOfStringValues("value", iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SerializeCalendarRoleType(m.GetValue()))
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *ItemCalendarAllowedCalendarSharingRolesWithUserGetResponse) SetValue(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CalendarRoleType)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// ItemCalendarAllowedCalendarSharingRolesWithUserGetResponseable 
type ItemCalendarAllowedCalendarSharingRolesWithUserGetResponseable interface {
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CalendarRoleType)
    SetValue(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CalendarRoleType)()
}
