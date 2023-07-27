// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// OAuth2AuthorizedAs - OAuth2AuthorizedAs tracks the user that OAuthed with the connector.
type OAuth2AuthorizedAs struct {
	// authEmail is the email of the user that authorized the connector using OAuth.
	AuthEmail    *string    `json:"authEmail,omitempty"`
	AuthorizedAt *time.Time `json:"authorizedAt,omitempty"`
}

func (o *OAuth2AuthorizedAs) GetAuthEmail() *string {
	if o == nil {
		return nil
	}
	return o.AuthEmail
}

func (o *OAuth2AuthorizedAs) GetAuthorizedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.AuthorizedAt
}