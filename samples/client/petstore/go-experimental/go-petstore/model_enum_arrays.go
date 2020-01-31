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

// EnumArrays struct for EnumArrays
type EnumArrays struct {
	JustSymbol *string `json:"just_symbol,omitempty"`
	ArrayEnum *[]string `json:"array_enum,omitempty"`
}

// NewEnumArrays instantiates a new EnumArrays object
func NewEnumArrays() *EnumArrays {
    this := EnumArrays{}
    return &this
}

// GetJustSymbol returns the JustSymbol field value if set, zero value otherwise.
func (o *EnumArrays) GetJustSymbol() string {
	if o == nil || o.JustSymbol == nil {
		var ret string
		return ret
	}
	return *o.JustSymbol
}

// GetJustSymbolOk returns a tuple with the JustSymbol field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *EnumArrays) GetJustSymbolOk() (string, bool) {
	if o == nil || o.JustSymbol == nil {
		var ret string
		return ret, false
	}
	return *o.JustSymbol, true
}

// HasJustSymbol returns a boolean if a field has been set.
func (o *EnumArrays) HasJustSymbol() bool {
	if o != nil && o.JustSymbol != nil {
		return true
	}

	return false
}

// SetJustSymbol gets a reference to the given string and assigns it to the JustSymbol field.
func (o *EnumArrays) SetJustSymbol(v string) {
	o.JustSymbol = &v
}

// GetArrayEnum returns the ArrayEnum field value if set, zero value otherwise.
func (o *EnumArrays) GetArrayEnum() []string {
	if o == nil || o.ArrayEnum == nil {
		var ret []string
		return ret
	}
	return *o.ArrayEnum
}

// GetArrayEnumOk returns a tuple with the ArrayEnum field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *EnumArrays) GetArrayEnumOk() ([]string, bool) {
	if o == nil || o.ArrayEnum == nil {
		var ret []string
		return ret, false
	}
	return *o.ArrayEnum, true
}

// HasArrayEnum returns a boolean if a field has been set.
func (o *EnumArrays) HasArrayEnum() bool {
	if o != nil && o.ArrayEnum != nil {
		return true
	}

	return false
}

// SetArrayEnum gets a reference to the given []string and assigns it to the ArrayEnum field.
func (o *EnumArrays) SetArrayEnum(v []string) {
	o.ArrayEnum = &v
}

type NullableEnumArrays struct {
	Value EnumArrays
	ExplicitNull bool
}

func (v NullableEnumArrays) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableEnumArrays) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
