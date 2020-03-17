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
	"encoding/json"
)

// BigCatAllOf struct for BigCatAllOf
type BigCatAllOf struct {
	Kind *string `json:"kind,omitempty"`
}

// NewBigCatAllOf instantiates a new BigCatAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBigCatAllOf() *BigCatAllOf {
    this := BigCatAllOf{}
    return &this
}

// NewBigCatAllOfWithDefaults instantiates a new BigCatAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBigCatAllOfWithDefaults() *BigCatAllOf {
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

func (o BigCatAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	return json.Marshal(toSerialize)
}

type NullableBigCatAllOf struct {
	value *BigCatAllOf
	isSet bool
}

func (v NullableBigCatAllOf) Get() *BigCatAllOf {
	return v.value
}

func (v NullableBigCatAllOf) Set(val *BigCatAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBigCatAllOf) IsSet() bool {
	return v.isSet
}

func (v NullableBigCatAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBigCatAllOf(val *BigCatAllOf) *NullableBigCatAllOf {
	return &NullableBigCatAllOf{value: val, isSet: true}
}

func (v NullableBigCatAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBigCatAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
