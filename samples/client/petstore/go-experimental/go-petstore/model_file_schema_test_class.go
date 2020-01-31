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

// FileSchemaTestClass struct for FileSchemaTestClass
type FileSchemaTestClass struct {
	File *File `json:"file,omitempty"`
	Files *[]File `json:"files,omitempty"`
}

func NewFileSchemaTestClass() *FileSchemaTestClass {
    this := FileSchemaTestClass{}
    return &this
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *FileSchemaTestClass) GetFile() File {
	if o == nil || o.File == nil {
		var ret File
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FileSchemaTestClass) GetFileOk() (File, bool) {
	if o == nil || o.File == nil {
		var ret File
		return ret, false
	}
	return *o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *FileSchemaTestClass) HasFile() bool {
	if o != nil && o.File != nil {
		return true
	}

	return false
}

// SetFile gets a reference to the given File and assigns it to the File field.
func (o *FileSchemaTestClass) SetFile(v File) {
	o.File = &v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *FileSchemaTestClass) GetFiles() []File {
	if o == nil || o.Files == nil {
		var ret []File
		return ret
	}
	return *o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FileSchemaTestClass) GetFilesOk() ([]File, bool) {
	if o == nil || o.Files == nil {
		var ret []File
		return ret, false
	}
	return *o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *FileSchemaTestClass) HasFiles() bool {
	if o != nil && o.Files != nil {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []File and assigns it to the Files field.
func (o *FileSchemaTestClass) SetFiles(v []File) {
	o.Files = &v
}

type NullableFileSchemaTestClass struct {
	Value FileSchemaTestClass
	ExplicitNull bool
}

func (v NullableFileSchemaTestClass) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableFileSchemaTestClass) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
