/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package petstore

import (
	"bytes"
	"encoding/json"
)

// Model200Response Model for testing model name starting with number
type Model200Response struct {
	Name *int32 `json:"name,omitempty"`
	Class *string `json:"class,omitempty"`
}

// NewModel200Response instantiates a new Model200Response object
func NewModel200Response() *Model200Response {
    this := Model200Response{}
    return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Model200Response) GetName() int32 {
	if o == nil || o.Name == nil {
		var ret int32
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Model200Response) GetNameOk() (int32, bool) {
	if o == nil || o.Name == nil {
		var ret int32
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Model200Response) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given int32 and assigns it to the Name field.
func (o *Model200Response) SetName(v int32) {
	o.Name = &v
}

// GetClass returns the Class field value if set, zero value otherwise.
func (o *Model200Response) GetClass() string {
	if o == nil || o.Class == nil {
		var ret string
		return ret
	}
	return *o.Class
}

// GetClassOk returns a tuple with the Class field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Model200Response) GetClassOk() (string, bool) {
	if o == nil || o.Class == nil {
		var ret string
		return ret, false
	}
	return *o.Class, true
}

// HasClass returns a boolean if a field has been set.
func (o *Model200Response) HasClass() bool {
	if o != nil && o.Class != nil {
		return true
	}

	return false
}

// SetClass gets a reference to the given string and assigns it to the Class field.
func (o *Model200Response) SetClass(v string) {
	o.Class = &v
}

type NullableModel200Response struct {
	Value Model200Response
	ExplicitNull bool
}

func (v NullableModel200Response) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableModel200Response) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
