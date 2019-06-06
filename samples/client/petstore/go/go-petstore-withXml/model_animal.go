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
	"errors"
)

type Animal struct {
	ClassName *string `json:"className,omitempty" xml:"className"`

	Color *string `json:"color,omitempty" xml:"color"`

}

// GetClassName returns the ClassName field if non-nil, zero value otherwise.
func (o *Animal) GetClassName() string {
	if o == nil || o.ClassName == nil {
		var ret string
		return ret
	}
	return *o.ClassName
}

// GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Animal) GetClassNameOk() (string, bool) {
	if o == nil || o.ClassName == nil {
		var ret string
		return ret, false
	}
	return *o.ClassName, true
}

// HasClassName returns a boolean if a field has been set.
func (o *Animal) HasClassName() bool {
	if o != nil && o.ClassName != nil {
		return true
	}

	return false
}

// SetClassName gets a reference to the given string and assigns it to the ClassName field.
func (o *Animal) SetClassName(v string) {
	o.ClassName = &v
}

// GetColor returns the Color field if non-nil, zero value otherwise.
func (o *Animal) GetColor() string {
	if o == nil || o.Color == nil {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Animal) GetColorOk() (string, bool) {
	if o == nil || o.Color == nil {
		var ret string
		return ret, false
	}
	return *o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *Animal) HasColor() bool {
	if o != nil && o.Color != nil {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *Animal) SetColor(v string) {
	o.Color = &v
}


func (o Animal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClassName == nil {
		return nil, errors.New("ClassName is required and not nullable, but was not set on Animal")
	}
	if o.ClassName != nil {
		toSerialize["className"] = o.ClassName
	}
	if o.Color != nil {
		toSerialize["color"] = o.Color
	}
	return json.Marshal(toSerialize)
}


