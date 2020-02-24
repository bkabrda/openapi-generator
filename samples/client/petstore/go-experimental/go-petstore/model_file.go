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

// File Must be named `File` for test.
type File struct {
	// Test capitalization
	SourceURI *string `json:"sourceURI,omitempty"`
}

// NewFile instantiates a new File object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFile() *File {
    this := File{}
    return &this
}

// NewFileWithDefaults instantiates a new File object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileWithDefaults() *File {
    this := File{}
    return &this
}

// GetSourceURI returns the SourceURI field value if set, zero value otherwise.
func (o *File) GetSourceURI() string {
	if o == nil || o.SourceURI == nil {
		var ret string
		return ret
	}
	return *o.SourceURI
}

// GetSourceURIOk returns a tuple with the SourceURI field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *File) GetSourceURIOk() (string, bool) {
	if o == nil || o.SourceURI == nil {
		var ret string
		return ret, false
	}
    return *o.SourceURI, true
}

// HasSourceURI returns a boolean if a field has been set.
func (o *File) HasSourceURI() bool {
    if o != nil && o.SourceURI != nil {
		return true
	}

	return false
}

// SetSourceURI gets a reference to the given string and assigns it to the SourceURI field.
func (o *File) SetSourceURI(v string) {
	o.SourceURI = &v
}

func (o File) MarshalJSON() ([]byte, error) {
    toSerialize := map[string]interface{}{}
    if o.SourceURI != nil {
        toSerialize["sourceURI"] = o.SourceURI
    }
    return json.Marshal(toSerialize)
}

type NullableFile struct {
	value *File
	isSet bool
}

func (v NullableFile) Get() *File {
    return v.value
}

func (v NullableFile) Set(val *File) {
    v.value = val
    v.isSet = true
}

func (v NullableFile) IsSet() bool {
    return v.isSet
}

func (v NullableFile) Unset() {
    v.value = nil
    v.isSet = false
}

func NewNullableFile(val *File) *NullableFile {
    return &NullableFile{value: val, isSet: true}
}

func (v NullableFile) MarshalJSON() ([]byte, error) {
    return json.Marshal(v.value)
}

func (v *NullableFile) UnmarshalJSON(src []byte) error {
    v.isSet = true
	return json.Unmarshal(src, &v.value)
}
