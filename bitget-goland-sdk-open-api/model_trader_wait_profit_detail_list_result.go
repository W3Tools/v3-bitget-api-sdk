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

// TraderWaitProfitDetailListResult struct for TraderWaitProfitDetailListResult
type TraderWaitProfitDetailListResult struct {
	NextFlag             *bool                          `json:"nextFlag,omitempty"`
	ResultList           []TraderWaitProfitDetailResult `json:"resultList,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TraderWaitProfitDetailListResult TraderWaitProfitDetailListResult

// NewTraderWaitProfitDetailListResult instantiates a new TraderWaitProfitDetailListResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTraderWaitProfitDetailListResult() *TraderWaitProfitDetailListResult {
	this := TraderWaitProfitDetailListResult{}
	return &this
}

// NewTraderWaitProfitDetailListResultWithDefaults instantiates a new TraderWaitProfitDetailListResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTraderWaitProfitDetailListResultWithDefaults() *TraderWaitProfitDetailListResult {
	this := TraderWaitProfitDetailListResult{}
	return &this
}

// GetNextFlag returns the NextFlag field value if set, zero value otherwise.
func (o *TraderWaitProfitDetailListResult) GetNextFlag() bool {
	if o == nil || isNil(o.NextFlag) {
		var ret bool
		return ret
	}
	return *o.NextFlag
}

// GetNextFlagOk returns a tuple with the NextFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraderWaitProfitDetailListResult) GetNextFlagOk() (*bool, bool) {
	if o == nil || isNil(o.NextFlag) {
		return nil, false
	}
	return o.NextFlag, true
}

// HasNextFlag returns a boolean if a field has been set.
func (o *TraderWaitProfitDetailListResult) HasNextFlag() bool {
	if o != nil && !isNil(o.NextFlag) {
		return true
	}

	return false
}

// SetNextFlag gets a reference to the given bool and assigns it to the NextFlag field.
func (o *TraderWaitProfitDetailListResult) SetNextFlag(v bool) {
	o.NextFlag = &v
}

// GetResultList returns the ResultList field value if set, zero value otherwise.
func (o *TraderWaitProfitDetailListResult) GetResultList() []TraderWaitProfitDetailResult {
	if o == nil || isNil(o.ResultList) {
		var ret []TraderWaitProfitDetailResult
		return ret
	}
	return o.ResultList
}

// GetResultListOk returns a tuple with the ResultList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraderWaitProfitDetailListResult) GetResultListOk() ([]TraderWaitProfitDetailResult, bool) {
	if o == nil || isNil(o.ResultList) {
		return nil, false
	}
	return o.ResultList, true
}

// HasResultList returns a boolean if a field has been set.
func (o *TraderWaitProfitDetailListResult) HasResultList() bool {
	if o != nil && !isNil(o.ResultList) {
		return true
	}

	return false
}

// SetResultList gets a reference to the given []TraderWaitProfitDetailResult and assigns it to the ResultList field.
func (o *TraderWaitProfitDetailListResult) SetResultList(v []TraderWaitProfitDetailResult) {
	o.ResultList = v
}

func (o TraderWaitProfitDetailListResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NextFlag) {
		toSerialize["nextFlag"] = o.NextFlag
	}
	if !isNil(o.ResultList) {
		toSerialize["resultList"] = o.ResultList
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TraderWaitProfitDetailListResult) UnmarshalJSON(bytes []byte) (err error) {
	varTraderWaitProfitDetailListResult := _TraderWaitProfitDetailListResult{}

	if err = json.Unmarshal(bytes, &varTraderWaitProfitDetailListResult); err == nil {
		*o = TraderWaitProfitDetailListResult(varTraderWaitProfitDetailListResult)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "nextFlag")
		delete(additionalProperties, "resultList")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTraderWaitProfitDetailListResult struct {
	value *TraderWaitProfitDetailListResult
	isSet bool
}

func (v NullableTraderWaitProfitDetailListResult) Get() *TraderWaitProfitDetailListResult {
	return v.value
}

func (v *NullableTraderWaitProfitDetailListResult) Set(val *TraderWaitProfitDetailListResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTraderWaitProfitDetailListResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTraderWaitProfitDetailListResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTraderWaitProfitDetailListResult(val *TraderWaitProfitDetailListResult) *NullableTraderWaitProfitDetailListResult {
	return &NullableTraderWaitProfitDetailListResult{value: val, isSet: true}
}

func (v NullableTraderWaitProfitDetailListResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTraderWaitProfitDetailListResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
