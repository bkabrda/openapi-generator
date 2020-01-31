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
	"os"
	"time"
)

// FormatTest struct for FormatTest
type FormatTest struct {
	Integer *int32 `json:"integer,omitempty"`
	Int32 *int32 `json:"int32,omitempty"`
	Int64 *int64 `json:"int64,omitempty"`
	Number float32 `json:"number"`
	Float *float32 `json:"float,omitempty"`
	Double *float64 `json:"double,omitempty"`
	String *string `json:"string,omitempty"`
	Byte string `json:"byte"`
	Binary **os.File `json:"binary,omitempty"`
	Date string `json:"date"`
	DateTime *time.Time `json:"dateTime,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
	Password string `json:"password"`
	// A string that is a 10 digit number. Can have leading zeros.
	PatternWithDigits *string `json:"pattern_with_digits,omitempty"`
	// A string starting with 'image_' (case insensitive) and one to three digits following i.e. Image_01.
	PatternWithDigitsAndDelimiter *string `json:"pattern_with_digits_and_delimiter,omitempty"`
}

func NewFormatTest() *FormatTest {
    this := FormatTest{}
    return &this
}

// GetInteger returns the Integer field value if set, zero value otherwise.
func (o *FormatTest) GetInteger() int32 {
	if o == nil || o.Integer == nil {
		var ret int32
		return ret
	}
	return *o.Integer
}

// GetIntegerOk returns a tuple with the Integer field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FormatTest) GetIntegerOk() (int32, bool) {
	if o == nil || o.Integer == nil {
		var ret int32
		return ret, false
	}
	return *o.Integer, true
}

// HasInteger returns a boolean if a field has been set.
func (o *FormatTest) HasInteger() bool {
	if o != nil && o.Integer != nil {
		return true
	}

	return false
}

// SetInteger gets a reference to the given int32 and assigns it to the Integer field.
func (o *FormatTest) SetInteger(v int32) {
	o.Integer = &v
}

// GetInt32 returns the Int32 field value if set, zero value otherwise.
func (o *FormatTest) GetInt32() int32 {
	if o == nil || o.Int32 == nil {
		var ret int32
		return ret
	}
	return *o.Int32
}

// GetInt32Ok returns a tuple with the Int32 field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FormatTest) GetInt32Ok() (int32, bool) {
	if o == nil || o.Int32 == nil {
		var ret int32
		return ret, false
	}
	return *o.Int32, true
}

// HasInt32 returns a boolean if a field has been set.
func (o *FormatTest) HasInt32() bool {
	if o != nil && o.Int32 != nil {
		return true
	}

	return false
}

// SetInt32 gets a reference to the given int32 and assigns it to the Int32 field.
func (o *FormatTest) SetInt32(v int32) {
	o.Int32 = &v
}

// GetInt64 returns the Int64 field value if set, zero value otherwise.
func (o *FormatTest) GetInt64() int64 {
	if o == nil || o.Int64 == nil {
		var ret int64
		return ret
	}
	return *o.Int64
}

// GetInt64Ok returns a tuple with the Int64 field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FormatTest) GetInt64Ok() (int64, bool) {
	if o == nil || o.Int64 == nil {
		var ret int64
		return ret, false
	}
	return *o.Int64, true
}

// HasInt64 returns a boolean if a field has been set.
func (o *FormatTest) HasInt64() bool {
	if o != nil && o.Int64 != nil {
		return true
	}

	return false
}

// SetInt64 gets a reference to the given int64 and assigns it to the Int64 field.
func (o *FormatTest) SetInt64(v int64) {
	o.Int64 = &v
}

// GetNumber returns the Number field value
func (o *FormatTest) GetNumber() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Number
}

// SetNumber sets field value
func (o *FormatTest) SetNumber(v float32) {
	o.Number = v
}

// GetFloat returns the Float field value if set, zero value otherwise.
func (o *FormatTest) GetFloat() float32 {
	if o == nil || o.Float == nil {
		var ret float32
		return ret
	}
	return *o.Float
}

// GetFloatOk returns a tuple with the Float field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FormatTest) GetFloatOk() (float32, bool) {
	if o == nil || o.Float == nil {
		var ret float32
		return ret, false
	}
	return *o.Float, true
}

// HasFloat returns a boolean if a field has been set.
func (o *FormatTest) HasFloat() bool {
	if o != nil && o.Float != nil {
		return true
	}

	return false
}

// SetFloat gets a reference to the given float32 and assigns it to the Float field.
func (o *FormatTest) SetFloat(v float32) {
	o.Float = &v
}

// GetDouble returns the Double field value if set, zero value otherwise.
func (o *FormatTest) GetDouble() float64 {
	if o == nil || o.Double == nil {
		var ret float64
		return ret
	}
	return *o.Double
}

// GetDoubleOk returns a tuple with the Double field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FormatTest) GetDoubleOk() (float64, bool) {
	if o == nil || o.Double == nil {
		var ret float64
		return ret, false
	}
	return *o.Double, true
}

// HasDouble returns a boolean if a field has been set.
func (o *FormatTest) HasDouble() bool {
	if o != nil && o.Double != nil {
		return true
	}

	return false
}

// SetDouble gets a reference to the given float64 and assigns it to the Double field.
func (o *FormatTest) SetDouble(v float64) {
	o.Double = &v
}

// GetString returns the String field value if set, zero value otherwise.
func (o *FormatTest) GetString() string {
	if o == nil || o.String == nil {
		var ret string
		return ret
	}
	return *o.String
}

// GetStringOk returns a tuple with the String field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FormatTest) GetStringOk() (string, bool) {
	if o == nil || o.String == nil {
		var ret string
		return ret, false
	}
	return *o.String, true
}

// HasString returns a boolean if a field has been set.
func (o *FormatTest) HasString() bool {
	if o != nil && o.String != nil {
		return true
	}

	return false
}

// SetString gets a reference to the given string and assigns it to the String field.
func (o *FormatTest) SetString(v string) {
	o.String = &v
}

// GetByte returns the Byte field value
func (o *FormatTest) GetByte() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Byte
}

// SetByte sets field value
func (o *FormatTest) SetByte(v string) {
	o.Byte = v
}

// GetBinary returns the Binary field value if set, zero value otherwise.
func (o *FormatTest) GetBinary() *os.File {
	if o == nil || o.Binary == nil {
		var ret *os.File
		return ret
	}
	return *o.Binary
}

// GetBinaryOk returns a tuple with the Binary field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FormatTest) GetBinaryOk() (*os.File, bool) {
	if o == nil || o.Binary == nil {
		var ret *os.File
		return ret, false
	}
	return *o.Binary, true
}

// HasBinary returns a boolean if a field has been set.
func (o *FormatTest) HasBinary() bool {
	if o != nil && o.Binary != nil {
		return true
	}

	return false
}

// SetBinary gets a reference to the given *os.File and assigns it to the Binary field.
func (o *FormatTest) SetBinary(v *os.File) {
	o.Binary = &v
}

// GetDate returns the Date field value
func (o *FormatTest) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// SetDate sets field value
func (o *FormatTest) SetDate(v string) {
	o.Date = v
}

// GetDateTime returns the DateTime field value if set, zero value otherwise.
func (o *FormatTest) GetDateTime() time.Time {
	if o == nil || o.DateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.DateTime
}

// GetDateTimeOk returns a tuple with the DateTime field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FormatTest) GetDateTimeOk() (time.Time, bool) {
	if o == nil || o.DateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.DateTime, true
}

// HasDateTime returns a boolean if a field has been set.
func (o *FormatTest) HasDateTime() bool {
	if o != nil && o.DateTime != nil {
		return true
	}

	return false
}

// SetDateTime gets a reference to the given time.Time and assigns it to the DateTime field.
func (o *FormatTest) SetDateTime(v time.Time) {
	o.DateTime = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *FormatTest) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FormatTest) GetUuidOk() (string, bool) {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret, false
	}
	return *o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *FormatTest) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *FormatTest) SetUuid(v string) {
	o.Uuid = &v
}

// GetPassword returns the Password field value
func (o *FormatTest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// SetPassword sets field value
func (o *FormatTest) SetPassword(v string) {
	o.Password = v
}

// GetPatternWithDigits returns the PatternWithDigits field value if set, zero value otherwise.
func (o *FormatTest) GetPatternWithDigits() string {
	if o == nil || o.PatternWithDigits == nil {
		var ret string
		return ret
	}
	return *o.PatternWithDigits
}

// GetPatternWithDigitsOk returns a tuple with the PatternWithDigits field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FormatTest) GetPatternWithDigitsOk() (string, bool) {
	if o == nil || o.PatternWithDigits == nil {
		var ret string
		return ret, false
	}
	return *o.PatternWithDigits, true
}

// HasPatternWithDigits returns a boolean if a field has been set.
func (o *FormatTest) HasPatternWithDigits() bool {
	if o != nil && o.PatternWithDigits != nil {
		return true
	}

	return false
}

// SetPatternWithDigits gets a reference to the given string and assigns it to the PatternWithDigits field.
func (o *FormatTest) SetPatternWithDigits(v string) {
	o.PatternWithDigits = &v
}

// GetPatternWithDigitsAndDelimiter returns the PatternWithDigitsAndDelimiter field value if set, zero value otherwise.
func (o *FormatTest) GetPatternWithDigitsAndDelimiter() string {
	if o == nil || o.PatternWithDigitsAndDelimiter == nil {
		var ret string
		return ret
	}
	return *o.PatternWithDigitsAndDelimiter
}

// GetPatternWithDigitsAndDelimiterOk returns a tuple with the PatternWithDigitsAndDelimiter field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FormatTest) GetPatternWithDigitsAndDelimiterOk() (string, bool) {
	if o == nil || o.PatternWithDigitsAndDelimiter == nil {
		var ret string
		return ret, false
	}
	return *o.PatternWithDigitsAndDelimiter, true
}

// HasPatternWithDigitsAndDelimiter returns a boolean if a field has been set.
func (o *FormatTest) HasPatternWithDigitsAndDelimiter() bool {
	if o != nil && o.PatternWithDigitsAndDelimiter != nil {
		return true
	}

	return false
}

// SetPatternWithDigitsAndDelimiter gets a reference to the given string and assigns it to the PatternWithDigitsAndDelimiter field.
func (o *FormatTest) SetPatternWithDigitsAndDelimiter(v string) {
	o.PatternWithDigitsAndDelimiter = &v
}

type NullableFormatTest struct {
	Value FormatTest
	ExplicitNull bool
}

func (v NullableFormatTest) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableFormatTest) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
