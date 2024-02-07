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

// checks if the EndpointListRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndpointListRes{}

// EndpointListRes struct for EndpointListRes
type EndpointListRes struct {
	Count int64 `json:"count"`
	Data []Endpoint `json:"data"`
}

type _EndpointListRes EndpointListRes

// NewEndpointListRes instantiates a new EndpointListRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointListRes(count int64, data []Endpoint) *EndpointListRes {
	this := EndpointListRes{}
	this.Count = count
	this.Data = data
	return &this
}

// NewEndpointListResWithDefaults instantiates a new EndpointListRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointListResWithDefaults() *EndpointListRes {
	this := EndpointListRes{}
	return &this
}

// GetCount returns the Count field value
func (o *EndpointListRes) GetCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *EndpointListRes) GetCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *EndpointListRes) SetCount(v int64) {
	o.Count = v
}

// GetData returns the Data field value
func (o *EndpointListRes) GetData() []Endpoint {
	if o == nil {
		var ret []Endpoint
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *EndpointListRes) GetDataOk() ([]Endpoint, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *EndpointListRes) SetData(v []Endpoint) {
	o.Data = v
}

func (o EndpointListRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointListRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *EndpointListRes) UnmarshalJSON(data []byte) (err error) {
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

	varEndpointListRes := _EndpointListRes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEndpointListRes)

	if err != nil {
		return err
	}

	*o = EndpointListRes(varEndpointListRes)

	return err
}

type NullableEndpointListRes struct {
	value *EndpointListRes
	isSet bool
}

func (v NullableEndpointListRes) Get() *EndpointListRes {
	return v.value
}

func (v *NullableEndpointListRes) Set(val *EndpointListRes) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointListRes) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointListRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointListRes(val *EndpointListRes) *NullableEndpointListRes {
	return &NullableEndpointListRes{value: val, isSet: true}
}

func (v NullableEndpointListRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointListRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

