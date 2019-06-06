/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package petstore
import (
	"encoding/json"
)

type CatAllOf struct {
	Declawed *bool `json:"declawed,omitempty" xml:"declawed"`

}

// GetDeclawed returns the Declawed field if non-nil, zero value otherwise.
func (o *CatAllOf) GetDeclawed() bool {
	if o == nil || o.Declawed == nil {
		var ret bool
		return ret
	}
	return *o.Declawed
}

// GetDeclawedOk returns a tuple with the Declawed field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CatAllOf) GetDeclawedOk() (bool, bool) {
	if o == nil || o.Declawed == nil {
		var ret bool
		return ret, false
	}
	return *o.Declawed, true
}

// HasDeclawed returns a boolean if a field has been set.
func (o *CatAllOf) HasDeclawed() bool {
	if o != nil && o.Declawed != nil {
		return true
	}

	return false
}

// SetDeclawed gets a reference to the given bool and assigns it to the Declawed field.
func (o *CatAllOf) SetDeclawed(v bool) {
	o.Declawed = &v
}


func (o CatAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Declawed != nil {
		toSerialize["declawed"] = o.Declawed
	}
	return json.Marshal(toSerialize)
}


