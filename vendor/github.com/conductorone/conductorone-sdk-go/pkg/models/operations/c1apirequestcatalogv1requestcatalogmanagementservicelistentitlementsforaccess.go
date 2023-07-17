// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsForAccessRequest struct {
	CatalogID string `pathParam:"style=simple,explode=false,name=catalog_id"`
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsForAccessRequest) GetCatalogID() string {
	if o == nil {
		return ""
	}
	return o.CatalogID
}

type C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsForAccessResponse struct {
	ContentType string
	// Successful response
	RequestCatalogManagementServiceListEntitlementsForAccessResponse *shared.RequestCatalogManagementServiceListEntitlementsForAccessResponse
	StatusCode                                                       int
	RawResponse                                                      *http.Response
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsForAccessResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsForAccessResponse) GetRequestCatalogManagementServiceListEntitlementsForAccessResponse() *shared.RequestCatalogManagementServiceListEntitlementsForAccessResponse {
	if o == nil {
		return nil
	}
	return o.RequestCatalogManagementServiceListEntitlementsForAccessResponse
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsForAccessResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsForAccessResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}