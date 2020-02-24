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

// EnumClass the model 'EnumClass'
type EnumClass string

// List of EnumClass
const (
	ENUMCLASS_ABC EnumClass = "_abc"
	ENUMCLASS_EFG EnumClass = "-efg"
	ENUMCLASS_XYZ EnumClass = "(xyz)"
)

type NullableEnumClass struct {
	value *EnumClass
	isSet bool
}

func (v NullableEnumClass) Get() *EnumClass {
    return v.value
}

func (v NullableEnumClass) Set(val *EnumClass) {
    v.value = val
    v.isSet = true
}

func (v NullableEnumClass) IsSet() bool {
    return v.isSet
}

func (v NullableEnumClass) Unset() {
    v.value = nil
    v.isSet = false
}

func NewNullableEnumClass(val *EnumClass) *NullableEnumClass {
    return &NullableEnumClass{value: val, isSet: true}
}

func (v NullableEnumClass) MarshalJSON() ([]byte, error) {
    return json.Marshal(v.value)
}

func (v *NullableEnumClass) UnmarshalJSON(src []byte) error {
    v.isSet = true
	return json.Unmarshal(src, &v.value)
}
