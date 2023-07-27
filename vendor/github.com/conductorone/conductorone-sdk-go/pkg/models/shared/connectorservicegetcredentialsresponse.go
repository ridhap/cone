// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ConnectorServiceGetCredentialsResponse - ConnectorServiceGetCredentialsResponse is the response returned by the get method.
type ConnectorServiceGetCredentialsResponse struct {
	// ConnectorCredential is used by a connector to authenticate with conductor one.
	ConnectorCredential *ConnectorCredential `json:"credential,omitempty"`
}

func (o *ConnectorServiceGetCredentialsResponse) GetConnectorCredential() *ConnectorCredential {
	if o == nil {
		return nil
	}
	return o.ConnectorCredential
}