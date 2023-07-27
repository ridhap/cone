// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// RequestCatalogManagementServiceGetResponse - The request catalog management service get response returns a request catalog view with the expanded items in the expanded array indicated by the expand mask in the request.
type RequestCatalogManagementServiceGetResponse struct {
	// The request catalog view contains the serialized request catalog and paths to objects referenced by the request catalog.
	RequestCatalogView *RequestCatalogView `json:"requestCatalogView,omitempty"`
	// List of serialized related objects.
	Expanded []map[string]interface{} `json:"expanded,omitempty"`
}

func (o *RequestCatalogManagementServiceGetResponse) GetRequestCatalogView() *RequestCatalogView {
	if o == nil {
		return nil
	}
	return o.RequestCatalogView
}

func (o *RequestCatalogManagementServiceGetResponse) GetExpanded() []map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.Expanded
}
