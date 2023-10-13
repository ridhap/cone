// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAttributeV1AttributesCreateAttributeValueResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// CreateAttributeValueResponse is the response for creating an attribute value.
	CreateAttributeValueResponse *shared.CreateAttributeValueResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *C1APIAttributeV1AttributesCreateAttributeValueResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAttributeV1AttributesCreateAttributeValueResponse) GetCreateAttributeValueResponse() *shared.CreateAttributeValueResponse {
	if o == nil {
		return nil
	}
	return o.CreateAttributeValueResponse
}

func (o *C1APIAttributeV1AttributesCreateAttributeValueResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAttributeV1AttributesCreateAttributeValueResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}