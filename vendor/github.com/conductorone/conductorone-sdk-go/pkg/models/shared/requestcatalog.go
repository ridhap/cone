// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// RequestCatalog - The RequestCatalog message.
type RequestCatalog struct {
	// The accessEntitlements field.
	AccessEntitlements []AppEntitlement `json:"accessEntitlements,omitempty"`
	// The appIds field.
	AppIds    []string   `json:"appIds,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The createdByUserId field.
	CreatedByUserID *string `json:"createdByUserId,omitempty"`
	// The description field.
	Description *string `json:"description,omitempty"`
	// The displayName field.
	DisplayName *string `json:"displayName,omitempty"`
	// The id field.
	ID *string `json:"id,omitempty"`
	// The published field.
	Published *bool      `json:"published,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// The visibleToEveryone field.
	VisibleToEveryone *bool `json:"visibleToEveryone,omitempty"`
}

func (o *RequestCatalog) GetAccessEntitlements() []AppEntitlement {
	if o == nil {
		return nil
	}
	return o.AccessEntitlements
}

func (o *RequestCatalog) GetAppIds() []string {
	if o == nil {
		return nil
	}
	return o.AppIds
}

func (o *RequestCatalog) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *RequestCatalog) GetCreatedByUserID() *string {
	if o == nil {
		return nil
	}
	return o.CreatedByUserID
}

func (o *RequestCatalog) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *RequestCatalog) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *RequestCatalog) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *RequestCatalog) GetPublished() *bool {
	if o == nil {
		return nil
	}
	return o.Published
}

func (o *RequestCatalog) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *RequestCatalog) GetVisibleToEveryone() *bool {
	if o == nil {
		return nil
	}
	return o.VisibleToEveryone
}