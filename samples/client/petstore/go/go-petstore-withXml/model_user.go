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

type User struct {
	Id *int64 `json:"id,omitempty" xml:"id"`

	Username *string `json:"username,omitempty" xml:"username"`

	FirstName *string `json:"firstName,omitempty" xml:"firstName"`

	LastName *string `json:"lastName,omitempty" xml:"lastName"`

	Email *string `json:"email,omitempty" xml:"email"`

	Password *string `json:"password,omitempty" xml:"password"`

	Phone *string `json:"phone,omitempty" xml:"phone"`

	// User Status
	UserStatus *int32 `json:"userStatus,omitempty" xml:"userStatus"`

}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *User) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *User) GetIdOk() (int64, bool) {
	if o == nil || o.Id == nil {
		var ret int64
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *User) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *User) SetId(v int64) {
	o.Id = &v
}

// GetUsername returns the Username field if non-nil, zero value otherwise.
func (o *User) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *User) GetUsernameOk() (string, bool) {
	if o == nil || o.Username == nil {
		var ret string
		return ret, false
	}
	return *o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *User) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *User) SetUsername(v string) {
	o.Username = &v
}

// GetFirstName returns the FirstName field if non-nil, zero value otherwise.
func (o *User) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *User) GetFirstNameOk() (string, bool) {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret, false
	}
	return *o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *User) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *User) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field if non-nil, zero value otherwise.
func (o *User) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLastNameOk() (string, bool) {
	if o == nil || o.LastName == nil {
		var ret string
		return ret, false
	}
	return *o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *User) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *User) SetLastName(v string) {
	o.LastName = &v
}

// GetEmail returns the Email field if non-nil, zero value otherwise.
func (o *User) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *User) GetEmailOk() (string, bool) {
	if o == nil || o.Email == nil {
		var ret string
		return ret, false
	}
	return *o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *User) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *User) SetEmail(v string) {
	o.Email = &v
}

// GetPassword returns the Password field if non-nil, zero value otherwise.
func (o *User) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *User) GetPasswordOk() (string, bool) {
	if o == nil || o.Password == nil {
		var ret string
		return ret, false
	}
	return *o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *User) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *User) SetPassword(v string) {
	o.Password = &v
}

// GetPhone returns the Phone field if non-nil, zero value otherwise.
func (o *User) GetPhone() string {
	if o == nil || o.Phone == nil {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *User) GetPhoneOk() (string, bool) {
	if o == nil || o.Phone == nil {
		var ret string
		return ret, false
	}
	return *o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *User) HasPhone() bool {
	if o != nil && o.Phone != nil {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *User) SetPhone(v string) {
	o.Phone = &v
}

// GetUserStatus returns the UserStatus field if non-nil, zero value otherwise.
func (o *User) GetUserStatus() int32 {
	if o == nil || o.UserStatus == nil {
		var ret int32
		return ret
	}
	return *o.UserStatus
}

// GetUserStatusOk returns a tuple with the UserStatus field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *User) GetUserStatusOk() (int32, bool) {
	if o == nil || o.UserStatus == nil {
		var ret int32
		return ret, false
	}
	return *o.UserStatus, true
}

// HasUserStatus returns a boolean if a field has been set.
func (o *User) HasUserStatus() bool {
	if o != nil && o.UserStatus != nil {
		return true
	}

	return false
}

// SetUserStatus gets a reference to the given int32 and assigns it to the UserStatus field.
func (o *User) SetUserStatus(v int32) {
	o.UserStatus = &v
}


func (o User) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.FirstName != nil {
		toSerialize["firstName"] = o.FirstName
	}
	if o.LastName != nil {
		toSerialize["lastName"] = o.LastName
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.Phone != nil {
		toSerialize["phone"] = o.Phone
	}
	if o.UserStatus != nil {
		toSerialize["userStatus"] = o.UserStatus
	}
	return json.Marshal(toSerialize)
}


