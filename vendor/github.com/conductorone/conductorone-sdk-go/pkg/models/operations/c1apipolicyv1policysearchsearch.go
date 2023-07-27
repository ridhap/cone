// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIPolicyV1PolicySearchSearchResponse struct {
	ContentType string
	// Successful response
	ListPolicyResponse *shared.ListPolicyResponse
	StatusCode         int
	RawResponse        *http.Response
}

func (o *C1APIPolicyV1PolicySearchSearchResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIPolicyV1PolicySearchSearchResponse) GetListPolicyResponse() *shared.ListPolicyResponse {
	if o == nil {
		return nil
	}
	return o.ListPolicyResponse
}

func (o *C1APIPolicyV1PolicySearchSearchResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIPolicyV1PolicySearchSearchResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}