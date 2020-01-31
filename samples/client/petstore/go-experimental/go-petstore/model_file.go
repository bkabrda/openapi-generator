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

// File Must be named `File` for test.
type File struct {
	// Test capitalization
	SourceURI *string `json:"sourceURI,omitempty"`
}

func NewFile() *File {
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

type NullableFile struct {
	Value File
	ExplicitNull bool
}

func (v NullableFile) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableFile) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
