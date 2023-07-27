// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ConnectorServiceRotateCredentialResponse - ConnectorServiceRotateCredentialResponse is the response returned by the rotate method.
type ConnectorServiceRotateCredentialResponse struct {
	// ConnectorCredential is used by a connector to authenticate with conductor one.
	ConnectorCredential *ConnectorCredential `json:"credential,omitempty"`
	// The new clientSecret returned after rotating the connector credential.
	ClientSecret *string `json:"clientSecret,omitempty"`
}

func (o *ConnectorServiceRotateCredentialResponse) GetConnectorCredential() *ConnectorCredential {
	if o == nil {
		return nil
	}
	return o.ConnectorCredential
}

func (o *ConnectorServiceRotateCredentialResponse) GetClientSecret() *string {
	if o == nil {
		return nil
	}
	return o.ClientSecret
}