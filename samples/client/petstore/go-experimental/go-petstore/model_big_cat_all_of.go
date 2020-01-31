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

// BigCatAllOf struct for BigCatAllOf
type BigCatAllOf struct {
	Kind *string `json:"kind,omitempty"`
}

// NewBigCatAllOf instantiates a new BigCatAllOf object
func NewBigCatAllOf() *BigCatAllOf {
    this := BigCatAllOf{}
    return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *BigCatAllOf) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *BigCatAllOf) GetKindOk() (string, bool) {
	if o == nil || o.Kind == nil {
		var ret string
		return ret, false
	}
	return *o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *BigCatAllOf) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *BigCatAllOf) SetKind(v string) {
	o.Kind = &v
}

type NullableBigCatAllOf struct {
	Value BigCatAllOf
	ExplicitNull bool
}

func (v NullableBigCatAllOf) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableBigCatAllOf) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
