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

// HistoryOrderListRequest struct for HistoryOrderListRequest
type HistoryOrderListRequest struct {
	// mixId
	MixId *string `json:"mixId,omitempty"`
	// pageSize
	PageSize *string `json:"pageSize,omitempty"`
	// trackingNo
	TrackingNo           *string `json:"trackingNo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HistoryOrderListRequest HistoryOrderListRequest

// NewHistoryOrderListRequest instantiates a new HistoryOrderListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoryOrderListRequest() *HistoryOrderListRequest {
	this := HistoryOrderListRequest{}
	return &this
}

// NewHistoryOrderListRequestWithDefaults instantiates a new HistoryOrderListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoryOrderListRequestWithDefaults() *HistoryOrderListRequest {
	this := HistoryOrderListRequest{}
	return &this
}

// GetMixId returns the MixId field value if set, zero value otherwise.
func (o *HistoryOrderListRequest) GetMixId() string {
	if o == nil || isNil(o.MixId) {
		var ret string
		return ret
	}
	return *o.MixId
}

// GetMixIdOk returns a tuple with the MixId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryOrderListRequest) GetMixIdOk() (*string, bool) {
	if o == nil || isNil(o.MixId) {
		return nil, false
	}
	return o.MixId, true
}

// HasMixId returns a boolean if a field has been set.
func (o *HistoryOrderListRequest) HasMixId() bool {
	if o != nil && !isNil(o.MixId) {
		return true
	}

	return false
}

// SetMixId gets a reference to the given string and assigns it to the MixId field.
func (o *HistoryOrderListRequest) SetMixId(v string) {
	o.MixId = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *HistoryOrderListRequest) GetPageSize() string {
	if o == nil || isNil(o.PageSize) {
		var ret string
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryOrderListRequest) GetPageSizeOk() (*string, bool) {
	if o == nil || isNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *HistoryOrderListRequest) HasPageSize() bool {
	if o != nil && !isNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given string and assigns it to the PageSize field.
func (o *HistoryOrderListRequest) SetPageSize(v string) {
	o.PageSize = &v
}

// GetTrackingNo returns the TrackingNo field value if set, zero value otherwise.
func (o *HistoryOrderListRequest) GetTrackingNo() string {
	if o == nil || isNil(o.TrackingNo) {
		var ret string
		return ret
	}
	return *o.TrackingNo
}

// GetTrackingNoOk returns a tuple with the TrackingNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryOrderListRequest) GetTrackingNoOk() (*string, bool) {
	if o == nil || isNil(o.TrackingNo) {
		return nil, false
	}
	return o.TrackingNo, true
}

// HasTrackingNo returns a boolean if a field has been set.
func (o *HistoryOrderListRequest) HasTrackingNo() bool {
	if o != nil && !isNil(o.TrackingNo) {
		return true
	}

	return false
}

// SetTrackingNo gets a reference to the given string and assigns it to the TrackingNo field.
func (o *HistoryOrderListRequest) SetTrackingNo(v string) {
	o.TrackingNo = &v
}

func (o HistoryOrderListRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MixId) {
		toSerialize["mixId"] = o.MixId
	}
	if !isNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	if !isNil(o.TrackingNo) {
		toSerialize["trackingNo"] = o.TrackingNo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HistoryOrderListRequest) UnmarshalJSON(bytes []byte) (err error) {
	varHistoryOrderListRequest := _HistoryOrderListRequest{}

	if err = json.Unmarshal(bytes, &varHistoryOrderListRequest); err == nil {
		*o = HistoryOrderListRequest(varHistoryOrderListRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "mixId")
		delete(additionalProperties, "pageSize")
		delete(additionalProperties, "trackingNo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHistoryOrderListRequest struct {
	value *HistoryOrderListRequest
	isSet bool
}

func (v NullableHistoryOrderListRequest) Get() *HistoryOrderListRequest {
	return v.value
}

func (v *NullableHistoryOrderListRequest) Set(val *HistoryOrderListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoryOrderListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoryOrderListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoryOrderListRequest(val *HistoryOrderListRequest) *NullableHistoryOrderListRequest {
	return &NullableHistoryOrderListRequest{value: val, isSet: true}
}

func (v NullableHistoryOrderListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoryOrderListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
