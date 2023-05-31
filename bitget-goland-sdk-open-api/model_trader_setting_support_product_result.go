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

// TraderSettingSupportProductResult struct for TraderSettingSupportProductResult
type TraderSettingSupportProductResult struct {
	OpenCopyTrace        *bool   `json:"openCopyTrace,omitempty"`
	ProductCode          *string `json:"productCode,omitempty"`
	ProductName          *string `json:"productName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TraderSettingSupportProductResult TraderSettingSupportProductResult

// NewTraderSettingSupportProductResult instantiates a new TraderSettingSupportProductResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTraderSettingSupportProductResult() *TraderSettingSupportProductResult {
	this := TraderSettingSupportProductResult{}
	return &this
}

// NewTraderSettingSupportProductResultWithDefaults instantiates a new TraderSettingSupportProductResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTraderSettingSupportProductResultWithDefaults() *TraderSettingSupportProductResult {
	this := TraderSettingSupportProductResult{}
	return &this
}

// GetOpenCopyTrace returns the OpenCopyTrace field value if set, zero value otherwise.
func (o *TraderSettingSupportProductResult) GetOpenCopyTrace() bool {
	if o == nil || isNil(o.OpenCopyTrace) {
		var ret bool
		return ret
	}
	return *o.OpenCopyTrace
}

// GetOpenCopyTraceOk returns a tuple with the OpenCopyTrace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraderSettingSupportProductResult) GetOpenCopyTraceOk() (*bool, bool) {
	if o == nil || isNil(o.OpenCopyTrace) {
		return nil, false
	}
	return o.OpenCopyTrace, true
}

// HasOpenCopyTrace returns a boolean if a field has been set.
func (o *TraderSettingSupportProductResult) HasOpenCopyTrace() bool {
	if o != nil && !isNil(o.OpenCopyTrace) {
		return true
	}

	return false
}

// SetOpenCopyTrace gets a reference to the given bool and assigns it to the OpenCopyTrace field.
func (o *TraderSettingSupportProductResult) SetOpenCopyTrace(v bool) {
	o.OpenCopyTrace = &v
}

// GetProductCode returns the ProductCode field value if set, zero value otherwise.
func (o *TraderSettingSupportProductResult) GetProductCode() string {
	if o == nil || isNil(o.ProductCode) {
		var ret string
		return ret
	}
	return *o.ProductCode
}

// GetProductCodeOk returns a tuple with the ProductCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraderSettingSupportProductResult) GetProductCodeOk() (*string, bool) {
	if o == nil || isNil(o.ProductCode) {
		return nil, false
	}
	return o.ProductCode, true
}

// HasProductCode returns a boolean if a field has been set.
func (o *TraderSettingSupportProductResult) HasProductCode() bool {
	if o != nil && !isNil(o.ProductCode) {
		return true
	}

	return false
}

// SetProductCode gets a reference to the given string and assigns it to the ProductCode field.
func (o *TraderSettingSupportProductResult) SetProductCode(v string) {
	o.ProductCode = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *TraderSettingSupportProductResult) GetProductName() string {
	if o == nil || isNil(o.ProductName) {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraderSettingSupportProductResult) GetProductNameOk() (*string, bool) {
	if o == nil || isNil(o.ProductName) {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *TraderSettingSupportProductResult) HasProductName() bool {
	if o != nil && !isNil(o.ProductName) {
		return true
	}

	return false
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *TraderSettingSupportProductResult) SetProductName(v string) {
	o.ProductName = &v
}

func (o TraderSettingSupportProductResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.OpenCopyTrace) {
		toSerialize["openCopyTrace"] = o.OpenCopyTrace
	}
	if !isNil(o.ProductCode) {
		toSerialize["productCode"] = o.ProductCode
	}
	if !isNil(o.ProductName) {
		toSerialize["productName"] = o.ProductName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TraderSettingSupportProductResult) UnmarshalJSON(bytes []byte) (err error) {
	varTraderSettingSupportProductResult := _TraderSettingSupportProductResult{}

	if err = json.Unmarshal(bytes, &varTraderSettingSupportProductResult); err == nil {
		*o = TraderSettingSupportProductResult(varTraderSettingSupportProductResult)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "openCopyTrace")
		delete(additionalProperties, "productCode")
		delete(additionalProperties, "productName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTraderSettingSupportProductResult struct {
	value *TraderSettingSupportProductResult
	isSet bool
}

func (v NullableTraderSettingSupportProductResult) Get() *TraderSettingSupportProductResult {
	return v.value
}

func (v *NullableTraderSettingSupportProductResult) Set(val *TraderSettingSupportProductResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTraderSettingSupportProductResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTraderSettingSupportProductResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTraderSettingSupportProductResult(val *TraderSettingSupportProductResult) *NullableTraderSettingSupportProductResult {
	return &NullableTraderSettingSupportProductResult{value: val, isSet: true}
}

func (v NullableTraderSettingSupportProductResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTraderSettingSupportProductResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
