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

// EnumTest struct for EnumTest
type EnumTest struct {
	EnumString *string `json:"enum_string,omitempty"`
	EnumStringRequired string `json:"enum_string_required"`
	EnumInteger *int32 `json:"enum_integer,omitempty"`
	EnumNumber *float64 `json:"enum_number,omitempty"`
	OuterEnum *OuterEnum `json:"outerEnum,omitempty"`
}

// NewEnumTest instantiates a new EnumTest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnumTest(enumStringRequired string, ) *EnumTest {
	this := EnumTest{}
	this.EnumStringRequired = enumStringRequired
	return &this
}

// NewEnumTestWithDefaults instantiates a new EnumTest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnumTestWithDefaults() *EnumTest {
	this := EnumTest{}
	return &this
}

// GetEnumString returns the EnumString field value if set, zero value otherwise.
func (o *EnumTest) GetEnumString() string {
	if o == nil || o.EnumString == nil {
		var ret string
		return ret
	}
	return *o.EnumString
}

// GetEnumStringOk returns a tuple with the EnumString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnumTest) GetEnumStringOk() (*string, bool) {
	if o == nil || o.EnumString == nil {
		return nil, false
	}
	return o.EnumString, true
}

// HasEnumString returns a boolean if a field has been set.
func (o *EnumTest) HasEnumString() bool {
	if o != nil && o.EnumString != nil {
		return true
	}

	return false
}

// SetEnumString gets a reference to the given string and assigns it to the EnumString field.
func (o *EnumTest) SetEnumString(v string) {
	o.EnumString = &v
}

// GetEnumStringRequired returns the EnumStringRequired field value
func (o *EnumTest) GetEnumStringRequired() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.EnumStringRequired
}

// GetEnumStringRequiredOk returns a tuple with the EnumStringRequired field value
// and a boolean to check if the value has been set.
func (o *EnumTest) GetEnumStringRequiredOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EnumStringRequired, true
}

// SetEnumStringRequired sets field value
func (o *EnumTest) SetEnumStringRequired(v string) {
	o.EnumStringRequired = v
}

// GetEnumInteger returns the EnumInteger field value if set, zero value otherwise.
func (o *EnumTest) GetEnumInteger() int32 {
	if o == nil || o.EnumInteger == nil {
		var ret int32
		return ret
	}
	return *o.EnumInteger
}

// GetEnumIntegerOk returns a tuple with the EnumInteger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnumTest) GetEnumIntegerOk() (*int32, bool) {
	if o == nil || o.EnumInteger == nil {
		return nil, false
	}
	return o.EnumInteger, true
}

// HasEnumInteger returns a boolean if a field has been set.
func (o *EnumTest) HasEnumInteger() bool {
	if o != nil && o.EnumInteger != nil {
		return true
	}

	return false
}

// SetEnumInteger gets a reference to the given int32 and assigns it to the EnumInteger field.
func (o *EnumTest) SetEnumInteger(v int32) {
	o.EnumInteger = &v
}

// GetEnumNumber returns the EnumNumber field value if set, zero value otherwise.
func (o *EnumTest) GetEnumNumber() float64 {
	if o == nil || o.EnumNumber == nil {
		var ret float64
		return ret
	}
	return *o.EnumNumber
}

// GetEnumNumberOk returns a tuple with the EnumNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnumTest) GetEnumNumberOk() (*float64, bool) {
	if o == nil || o.EnumNumber == nil {
		return nil, false
	}
	return o.EnumNumber, true
}

// HasEnumNumber returns a boolean if a field has been set.
func (o *EnumTest) HasEnumNumber() bool {
	if o != nil && o.EnumNumber != nil {
		return true
	}

	return false
}

// SetEnumNumber gets a reference to the given float64 and assigns it to the EnumNumber field.
func (o *EnumTest) SetEnumNumber(v float64) {
	o.EnumNumber = &v
}

// GetOuterEnum returns the OuterEnum field value if set, zero value otherwise.
func (o *EnumTest) GetOuterEnum() OuterEnum {
	if o == nil || o.OuterEnum == nil {
		var ret OuterEnum
		return ret
	}
	return *o.OuterEnum
}

// GetOuterEnumOk returns a tuple with the OuterEnum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnumTest) GetOuterEnumOk() (*OuterEnum, bool) {
	if o == nil || o.OuterEnum == nil {
		return nil, false
	}
	return o.OuterEnum, true
}

// HasOuterEnum returns a boolean if a field has been set.
func (o *EnumTest) HasOuterEnum() bool {
	if o != nil && o.OuterEnum != nil {
		return true
	}

	return false
}

// SetOuterEnum gets a reference to the given OuterEnum and assigns it to the OuterEnum field.
func (o *EnumTest) SetOuterEnum(v OuterEnum) {
	o.OuterEnum = &v
}

func (o EnumTest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EnumString != nil {
		toSerialize["enum_string"] = o.EnumString
	}
	if true {
		toSerialize["enum_string_required"] = o.EnumStringRequired
	}
	if o.EnumInteger != nil {
		toSerialize["enum_integer"] = o.EnumInteger
	}
	if o.EnumNumber != nil {
		toSerialize["enum_number"] = o.EnumNumber
	}
	if o.OuterEnum != nil {
		toSerialize["outerEnum"] = o.OuterEnum
	}
	return json.Marshal(toSerialize)
}

type NullableEnumTest struct {
	value *EnumTest
	isSet bool
}

func (v NullableEnumTest) Get() *EnumTest {
	return v.value
}

func (v *NullableEnumTest) Set(val *EnumTest) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumTest) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumTest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumTest(val *EnumTest) *NullableEnumTest {
	return &NullableEnumTest{value: val, isSet: true}
}

func (v NullableEnumTest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumTest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


