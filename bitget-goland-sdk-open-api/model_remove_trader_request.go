/*
Bitget Open API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// RemoveTraderRequest struct for RemoveTraderRequest
type RemoveTraderRequest struct {
	// traderUserId
	TraderUserId         string `json:"traderUserId"`
	AdditionalProperties map[string]interface{}
}

type _RemoveTraderRequest RemoveTraderRequest

// NewRemoveTraderRequest instantiates a new RemoveTraderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveTraderRequest(traderUserId string) *RemoveTraderRequest {
	this := RemoveTraderRequest{}
	this.TraderUserId = traderUserId
	return &this
}

// NewRemoveTraderRequestWithDefaults instantiates a new RemoveTraderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveTraderRequestWithDefaults() *RemoveTraderRequest {
	this := RemoveTraderRequest{}
	return &this
}

// GetTraderUserId returns the TraderUserId field value
func (o *RemoveTraderRequest) GetTraderUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TraderUserId
}

// GetTraderUserIdOk returns a tuple with the TraderUserId field value
// and a boolean to check if the value has been set.
func (o *RemoveTraderRequest) GetTraderUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TraderUserId, true
}

// SetTraderUserId sets field value
func (o *RemoveTraderRequest) SetTraderUserId(v string) {
	o.TraderUserId = v
}

func (o RemoveTraderRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["traderUserId"] = o.TraderUserId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RemoveTraderRequest) UnmarshalJSON(bytes []byte) (err error) {
	varRemoveTraderRequest := _RemoveTraderRequest{}

	if err = json.Unmarshal(bytes, &varRemoveTraderRequest); err == nil {
		*o = RemoveTraderRequest(varRemoveTraderRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "traderUserId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRemoveTraderRequest struct {
	value *RemoveTraderRequest
	isSet bool
}

func (v NullableRemoveTraderRequest) Get() *RemoveTraderRequest {
	return v.value
}

func (v *NullableRemoveTraderRequest) Set(val *RemoveTraderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveTraderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveTraderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveTraderRequest(val *RemoveTraderRequest) *NullableRemoveTraderRequest {
	return &NullableRemoveTraderRequest{value: val, isSet: true}
}

func (v NullableRemoveTraderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveTraderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
