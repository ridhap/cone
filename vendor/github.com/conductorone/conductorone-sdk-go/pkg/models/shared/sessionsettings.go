// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// The SessionSettings message.
type SessionSettings struct {
	MaxSessionLength *string `json:"maxSessionLength,omitempty"`
}

func (o *SessionSettings) GetMaxSessionLength() *string {
	if o == nil {
		return nil
	}
	return o.MaxSessionLength
}