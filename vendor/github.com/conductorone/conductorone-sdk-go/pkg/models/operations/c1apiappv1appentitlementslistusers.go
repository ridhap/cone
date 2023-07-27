// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppEntitlementsListUsersRequest struct {
	AppEntitlementID string   `pathParam:"style=simple,explode=false,name=app_entitlement_id"`
	AppID            string   `pathParam:"style=simple,explode=false,name=app_id"`
	PageSize         *float64 `queryParam:"style=form,explode=true,name=page_size"`
	PageToken        *string  `queryParam:"style=form,explode=true,name=page_token"`
}

func (o *C1APIAppV1AppEntitlementsListUsersRequest) GetAppEntitlementID() string {
	if o == nil {
		return ""
	}
	return o.AppEntitlementID
}

func (o *C1APIAppV1AppEntitlementsListUsersRequest) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *C1APIAppV1AppEntitlementsListUsersRequest) GetPageSize() *float64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *C1APIAppV1AppEntitlementsListUsersRequest) GetPageToken() *string {
	if o == nil {
		return nil
	}
	return o.PageToken
}

type C1APIAppV1AppEntitlementsListUsersResponse struct {
	ContentType string
	// The ListAppEntitlementUsersResponse message contains a list of results and a nextPageToken if applicable.
	ListAppEntitlementUsersResponse *shared.ListAppEntitlementUsersResponse
	StatusCode                      int
	RawResponse                     *http.Response
}

func (o *C1APIAppV1AppEntitlementsListUsersResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1AppEntitlementsListUsersResponse) GetListAppEntitlementUsersResponse() *shared.ListAppEntitlementUsersResponse {
	if o == nil {
		return nil
	}
	return o.ListAppEntitlementUsersResponse
}

func (o *C1APIAppV1AppEntitlementsListUsersResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1AppEntitlementsListUsersResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}