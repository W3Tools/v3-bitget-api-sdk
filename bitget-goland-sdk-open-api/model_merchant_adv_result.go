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

// MerchantAdvResult struct for MerchantAdvResult
type MerchantAdvResult struct {
	AdvList              []MerchantAdvInfo `json:"advList,omitempty"`
	MinId                *string           `json:"minId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MerchantAdvResult MerchantAdvResult

// NewMerchantAdvResult instantiates a new MerchantAdvResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantAdvResult() *MerchantAdvResult {
	this := MerchantAdvResult{}
	return &this
}

// NewMerchantAdvResultWithDefaults instantiates a new MerchantAdvResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantAdvResultWithDefaults() *MerchantAdvResult {
	this := MerchantAdvResult{}
	return &this
}

// GetAdvList returns the AdvList field value if set, zero value otherwise.
func (o *MerchantAdvResult) GetAdvList() []MerchantAdvInfo {
	if o == nil || isNil(o.AdvList) {
		var ret []MerchantAdvInfo
		return ret
	}
	return o.AdvList
}

// GetAdvListOk returns a tuple with the AdvList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantAdvResult) GetAdvListOk() ([]MerchantAdvInfo, bool) {
	if o == nil || isNil(o.AdvList) {
		return nil, false
	}
	return o.AdvList, true
}

// HasAdvList returns a boolean if a field has been set.
func (o *MerchantAdvResult) HasAdvList() bool {
	if o != nil && !isNil(o.AdvList) {
		return true
	}

	return false
}

// SetAdvList gets a reference to the given []MerchantAdvInfo and assigns it to the AdvList field.
func (o *MerchantAdvResult) SetAdvList(v []MerchantAdvInfo) {
	o.AdvList = v
}

// GetMinId returns the MinId field value if set, zero value otherwise.
func (o *MerchantAdvResult) GetMinId() string {
	if o == nil || isNil(o.MinId) {
		var ret string
		return ret
	}
	return *o.MinId
}

// GetMinIdOk returns a tuple with the MinId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantAdvResult) GetMinIdOk() (*string, bool) {
	if o == nil || isNil(o.MinId) {
		return nil, false
	}
	return o.MinId, true
}

// HasMinId returns a boolean if a field has been set.
func (o *MerchantAdvResult) HasMinId() bool {
	if o != nil && !isNil(o.MinId) {
		return true
	}

	return false
}

// SetMinId gets a reference to the given string and assigns it to the MinId field.
func (o *MerchantAdvResult) SetMinId(v string) {
	o.MinId = &v
}

func (o MerchantAdvResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AdvList) {
		toSerialize["advList"] = o.AdvList
	}
	if !isNil(o.MinId) {
		toSerialize["minId"] = o.MinId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MerchantAdvResult) UnmarshalJSON(bytes []byte) (err error) {
	varMerchantAdvResult := _MerchantAdvResult{}

	if err = json.Unmarshal(bytes, &varMerchantAdvResult); err == nil {
		*o = MerchantAdvResult(varMerchantAdvResult)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "advList")
		delete(additionalProperties, "minId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMerchantAdvResult struct {
	value *MerchantAdvResult
	isSet bool
}

func (v NullableMerchantAdvResult) Get() *MerchantAdvResult {
	return v.value
}

func (v *NullableMerchantAdvResult) Set(val *MerchantAdvResult) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantAdvResult) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantAdvResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantAdvResult(val *MerchantAdvResult) *NullableMerchantAdvResult {
	return &NullableMerchantAdvResult{value: val, isSet: true}
}

func (v NullableMerchantAdvResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantAdvResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}