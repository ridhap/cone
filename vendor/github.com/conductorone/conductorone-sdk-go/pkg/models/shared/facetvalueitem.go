// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// FacetValueItem - The FacetValueItem message.
type FacetValueItem struct {
	// An array of facet values.
	Values []FacetValue `json:"values,omitempty"`
}

func (o *FacetValueItem) GetValues() []FacetValue {
	if o == nil {
		return nil
	}
	return o.Values
}