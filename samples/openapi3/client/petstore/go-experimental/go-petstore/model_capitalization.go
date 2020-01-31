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

// Capitalization struct for Capitalization
type Capitalization struct {
	SmallCamel *string `json:"smallCamel,omitempty"`
	CapitalCamel *string `json:"CapitalCamel,omitempty"`
	SmallSnake *string `json:"small_Snake,omitempty"`
	CapitalSnake *string `json:"Capital_Snake,omitempty"`
	SCAETHFlowPoints *string `json:"SCA_ETH_Flow_Points,omitempty"`
	// Name of the pet 
	ATT_NAME *string `json:"ATT_NAME,omitempty"`
}

func NewCapitalization() *Capitalization {
    this := Capitalization{}
    return &this
}

// GetSmallCamel returns the SmallCamel field value if set, zero value otherwise.
func (o *Capitalization) GetSmallCamel() string {
	if o == nil || o.SmallCamel == nil {
		var ret string
		return ret
	}
	return *o.SmallCamel
}

// GetSmallCamelOk returns a tuple with the SmallCamel field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Capitalization) GetSmallCamelOk() (string, bool) {
	if o == nil || o.SmallCamel == nil {
		var ret string
		return ret, false
	}
	return *o.SmallCamel, true
}

// HasSmallCamel returns a boolean if a field has been set.
func (o *Capitalization) HasSmallCamel() bool {
	if o != nil && o.SmallCamel != nil {
		return true
	}

	return false
}

// SetSmallCamel gets a reference to the given string and assigns it to the SmallCamel field.
func (o *Capitalization) SetSmallCamel(v string) {
	o.SmallCamel = &v
}

// GetCapitalCamel returns the CapitalCamel field value if set, zero value otherwise.
func (o *Capitalization) GetCapitalCamel() string {
	if o == nil || o.CapitalCamel == nil {
		var ret string
		return ret
	}
	return *o.CapitalCamel
}

// GetCapitalCamelOk returns a tuple with the CapitalCamel field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Capitalization) GetCapitalCamelOk() (string, bool) {
	if o == nil || o.CapitalCamel == nil {
		var ret string
		return ret, false
	}
	return *o.CapitalCamel, true
}

// HasCapitalCamel returns a boolean if a field has been set.
func (o *Capitalization) HasCapitalCamel() bool {
	if o != nil && o.CapitalCamel != nil {
		return true
	}

	return false
}

// SetCapitalCamel gets a reference to the given string and assigns it to the CapitalCamel field.
func (o *Capitalization) SetCapitalCamel(v string) {
	o.CapitalCamel = &v
}

// GetSmallSnake returns the SmallSnake field value if set, zero value otherwise.
func (o *Capitalization) GetSmallSnake() string {
	if o == nil || o.SmallSnake == nil {
		var ret string
		return ret
	}
	return *o.SmallSnake
}

// GetSmallSnakeOk returns a tuple with the SmallSnake field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Capitalization) GetSmallSnakeOk() (string, bool) {
	if o == nil || o.SmallSnake == nil {
		var ret string
		return ret, false
	}
	return *o.SmallSnake, true
}

// HasSmallSnake returns a boolean if a field has been set.
func (o *Capitalization) HasSmallSnake() bool {
	if o != nil && o.SmallSnake != nil {
		return true
	}

	return false
}

// SetSmallSnake gets a reference to the given string and assigns it to the SmallSnake field.
func (o *Capitalization) SetSmallSnake(v string) {
	o.SmallSnake = &v
}

// GetCapitalSnake returns the CapitalSnake field value if set, zero value otherwise.
func (o *Capitalization) GetCapitalSnake() string {
	if o == nil || o.CapitalSnake == nil {
		var ret string
		return ret
	}
	return *o.CapitalSnake
}

// GetCapitalSnakeOk returns a tuple with the CapitalSnake field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Capitalization) GetCapitalSnakeOk() (string, bool) {
	if o == nil || o.CapitalSnake == nil {
		var ret string
		return ret, false
	}
	return *o.CapitalSnake, true
}

// HasCapitalSnake returns a boolean if a field has been set.
func (o *Capitalization) HasCapitalSnake() bool {
	if o != nil && o.CapitalSnake != nil {
		return true
	}

	return false
}

// SetCapitalSnake gets a reference to the given string and assigns it to the CapitalSnake field.
func (o *Capitalization) SetCapitalSnake(v string) {
	o.CapitalSnake = &v
}

// GetSCAETHFlowPoints returns the SCAETHFlowPoints field value if set, zero value otherwise.
func (o *Capitalization) GetSCAETHFlowPoints() string {
	if o == nil || o.SCAETHFlowPoints == nil {
		var ret string
		return ret
	}
	return *o.SCAETHFlowPoints
}

// GetSCAETHFlowPointsOk returns a tuple with the SCAETHFlowPoints field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Capitalization) GetSCAETHFlowPointsOk() (string, bool) {
	if o == nil || o.SCAETHFlowPoints == nil {
		var ret string
		return ret, false
	}
	return *o.SCAETHFlowPoints, true
}

// HasSCAETHFlowPoints returns a boolean if a field has been set.
func (o *Capitalization) HasSCAETHFlowPoints() bool {
	if o != nil && o.SCAETHFlowPoints != nil {
		return true
	}

	return false
}

// SetSCAETHFlowPoints gets a reference to the given string and assigns it to the SCAETHFlowPoints field.
func (o *Capitalization) SetSCAETHFlowPoints(v string) {
	o.SCAETHFlowPoints = &v
}

// GetATT_NAME returns the ATT_NAME field value if set, zero value otherwise.
func (o *Capitalization) GetATT_NAME() string {
	if o == nil || o.ATT_NAME == nil {
		var ret string
		return ret
	}
	return *o.ATT_NAME
}

// GetATT_NAMEOk returns a tuple with the ATT_NAME field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Capitalization) GetATT_NAMEOk() (string, bool) {
	if o == nil || o.ATT_NAME == nil {
		var ret string
		return ret, false
	}
	return *o.ATT_NAME, true
}

// HasATT_NAME returns a boolean if a field has been set.
func (o *Capitalization) HasATT_NAME() bool {
	if o != nil && o.ATT_NAME != nil {
		return true
	}

	return false
}

// SetATT_NAME gets a reference to the given string and assigns it to the ATT_NAME field.
func (o *Capitalization) SetATT_NAME(v string) {
	o.ATT_NAME = &v
}

type NullableCapitalization struct {
	Value Capitalization
	ExplicitNull bool
}

func (v NullableCapitalization) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCapitalization) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
