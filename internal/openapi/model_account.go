/*
Kanthor SDK API

SDK API

API version: 2022.1201.1701
Contact: support@kanthorlabs.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Account type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Account{}

// Account struct for Account
type Account struct {
	Metadata map[string]string `json:"metadata"`
	Name string `json:"name"`
	Sub string `json:"sub"`
}

type _Account Account

// NewAccount instantiates a new Account object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccount(metadata map[string]string, name string, sub string) *Account {
	this := Account{}
	this.Metadata = metadata
	this.Name = name
	this.Sub = sub
	return &this
}

// NewAccountWithDefaults instantiates a new Account object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountWithDefaults() *Account {
	this := Account{}
	return &this
}

// GetMetadata returns the Metadata field value
func (o *Account) GetMetadata() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *Account) GetMetadataOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *Account) SetMetadata(v map[string]string) {
	o.Metadata = v
}

// GetName returns the Name field value
func (o *Account) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Account) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Account) SetName(v string) {
	o.Name = v
}

// GetSub returns the Sub field value
func (o *Account) GetSub() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sub
}

// GetSubOk returns a tuple with the Sub field value
// and a boolean to check if the value has been set.
func (o *Account) GetSubOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sub, true
}

// SetSub sets field value
func (o *Account) SetSub(v string) {
	o.Sub = v
}

func (o Account) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Account) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["metadata"] = o.Metadata
	toSerialize["name"] = o.Name
	toSerialize["sub"] = o.Sub
	return toSerialize, nil
}

func (o *Account) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"metadata",
		"name",
		"sub",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAccount := _Account{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccount)

	if err != nil {
		return err
	}

	*o = Account(varAccount)

	return err
}

type NullableAccount struct {
	value *Account
	isSet bool
}

func (v NullableAccount) Get() *Account {
	return v.value
}

func (v *NullableAccount) Set(val *Account) {
	v.value = val
	v.isSet = true
}

func (v NullableAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccount(val *Account) *NullableAccount {
	return &NullableAccount{value: val, isSet: true}
}

func (v NullableAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


