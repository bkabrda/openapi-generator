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

// ApiResponse struct for ApiResponse
type ApiResponse struct {
	Code *int32 `json:"code,omitempty"`
	Type *string `json:"type,omitempty"`
	Message *string `json:"message,omitempty"`
}

func NewApiResponse() *ApiResponse {
    this := ApiResponse{}
    return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ApiResponse) GetCode() int32 {
	if o == nil || o.Code == nil {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ApiResponse) GetCodeOk() (int32, bool) {
	if o == nil || o.Code == nil {
		var ret int32
		return ret, false
	}
	return *o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ApiResponse) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *ApiResponse) SetCode(v int32) {
	o.Code = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ApiResponse) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ApiResponse) GetTypeOk() (string, bool) {
	if o == nil || o.Type == nil {
		var ret string
		return ret, false
	}
	return *o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApiResponse) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ApiResponse) SetType(v string) {
	o.Type = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ApiResponse) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ApiResponse) GetMessageOk() (string, bool) {
	if o == nil || o.Message == nil {
		var ret string
		return ret, false
	}
	return *o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ApiResponse) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ApiResponse) SetMessage(v string) {
	o.Message = &v
}

type NullableApiResponse struct {
	Value ApiResponse
	ExplicitNull bool
}

func (v NullableApiResponse) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableApiResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
