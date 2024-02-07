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

// checks if the ApplicationListRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationListRes{}

// ApplicationListRes struct for ApplicationListRes
type ApplicationListRes struct {
	Count int64 `json:"count"`
	Data []Application `json:"data"`
}

type _ApplicationListRes ApplicationListRes

// NewApplicationListRes instantiates a new ApplicationListRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationListRes(count int64, data []Application) *ApplicationListRes {
	this := ApplicationListRes{}
	this.Count = count
	this.Data = data
	return &this
}

// NewApplicationListResWithDefaults instantiates a new ApplicationListRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationListResWithDefaults() *ApplicationListRes {
	this := ApplicationListRes{}
	return &this
}

// GetCount returns the Count field value
func (o *ApplicationListRes) GetCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *ApplicationListRes) GetCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *ApplicationListRes) SetCount(v int64) {
	o.Count = v
}

// GetData returns the Data field value
func (o *ApplicationListRes) GetData() []Application {
	if o == nil {
		var ret []Application
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ApplicationListRes) GetDataOk() ([]Application, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ApplicationListRes) SetData(v []Application) {
	o.Data = v
}

func (o ApplicationListRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationListRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *ApplicationListRes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"count",
		"data",
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

	varApplicationListRes := _ApplicationListRes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApplicationListRes)

	if err != nil {
		return err
	}

	*o = ApplicationListRes(varApplicationListRes)

	return err
}

type NullableApplicationListRes struct {
	value *ApplicationListRes
	isSet bool
}

func (v NullableApplicationListRes) Get() *ApplicationListRes {
	return v.value
}

func (v *NullableApplicationListRes) Set(val *ApplicationListRes) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationListRes) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationListRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationListRes(val *ApplicationListRes) *NullableApplicationListRes {
	return &NullableApplicationListRes{value: val, isSet: true}
}

func (v NullableApplicationListRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationListRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


