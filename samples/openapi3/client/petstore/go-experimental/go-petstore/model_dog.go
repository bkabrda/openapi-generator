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

// Dog struct for Dog
type Dog struct {
	Animal
	Breed *string `json:"breed,omitempty"`
}

// NewDog instantiates a new Dog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDog() *Dog {
    this := Dog{}
    return &this
}

// NewDogWithDefaults instantiates a new Dog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDogWithDefaults() *Dog {
    this := Dog{}
    return &this
}

// GetBreed returns the Breed field value if set, zero value otherwise.
func (o *Dog) GetBreed() string {
	if o == nil || o.Breed == nil {
		var ret string
		return ret
	}
	return *o.Breed
}

// GetBreedOk returns a tuple with the Breed field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Dog) GetBreedOk() (string, bool) {
	if o == nil || o.Breed == nil {
		var ret string
		return ret, false
	}
	return *o.Breed, true
}

// HasBreed returns a boolean if a field has been set.
func (o *Dog) HasBreed() bool {
	if o != nil && o.Breed != nil {
		return true
	}

	return false
}

// SetBreed gets a reference to the given string and assigns it to the Breed field.
func (o *Dog) SetBreed(v string) {
	o.Breed = &v
}

func (o Dog) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAnimal, errAnimal := json.Marshal(o.Animal)
	if errAnimal != nil {
		return []byte{}, errAnimal
	}
	errAnimal = json.Unmarshal([]byte(serializedAnimal), &toSerialize)
	if errAnimal != nil {
		return []byte{}, errAnimal
	}
	if o.Breed != nil {
		toSerialize["breed"] = o.Breed
	}
	return json.Marshal(toSerialize)
}

type NullableDog struct {
	value *Dog
	isSet bool
}

func (v NullableDog) Get() *Dog {
	return v.value
}

func (v *NullableDog) Set(val *Dog) {
	v.value = val
	v.isSet = true
}

func (v NullableDog) IsSet() bool {
	return v.isSet
}

func (v *NullableDog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDog(val *Dog) *NullableDog {
	return &NullableDog{value: val, isSet: true}
}

func (v NullableDog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
