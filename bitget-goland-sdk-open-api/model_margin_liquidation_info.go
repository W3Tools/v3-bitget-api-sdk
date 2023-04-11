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

// MarginLiquidationInfo struct for MarginLiquidationInfo
type MarginLiquidationInfo struct {
	Ctime                *string `json:"ctime,omitempty"`
	LiqEndTime           *string `json:"liqEndTime,omitempty"`
	LiqFee               *string `json:"liqFee,omitempty"`
	LiqId                *string `json:"liqId,omitempty"`
	LiqRisk              *string `json:"liqRisk,omitempty"`
	LiqStartTime         *string `json:"liqStartTime,omitempty"`
	TotalAssets          *string `json:"totalAssets,omitempty"`
	TotalDebt            *string `json:"totalDebt,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarginLiquidationInfo MarginLiquidationInfo

// NewMarginLiquidationInfo instantiates a new MarginLiquidationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginLiquidationInfo() *MarginLiquidationInfo {
	this := MarginLiquidationInfo{}
	return &this
}

// NewMarginLiquidationInfoWithDefaults instantiates a new MarginLiquidationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginLiquidationInfoWithDefaults() *MarginLiquidationInfo {
	this := MarginLiquidationInfo{}
	return &this
}

// GetCtime returns the Ctime field value if set, zero value otherwise.
func (o *MarginLiquidationInfo) GetCtime() string {
	if o == nil || isNil(o.Ctime) {
		var ret string
		return ret
	}
	return *o.Ctime
}

// GetCtimeOk returns a tuple with the Ctime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginLiquidationInfo) GetCtimeOk() (*string, bool) {
	if o == nil || isNil(o.Ctime) {
		return nil, false
	}
	return o.Ctime, true
}

// HasCtime returns a boolean if a field has been set.
func (o *MarginLiquidationInfo) HasCtime() bool {
	if o != nil && !isNil(o.Ctime) {
		return true
	}

	return false
}

// SetCtime gets a reference to the given string and assigns it to the Ctime field.
func (o *MarginLiquidationInfo) SetCtime(v string) {
	o.Ctime = &v
}

// GetLiqEndTime returns the LiqEndTime field value if set, zero value otherwise.
func (o *MarginLiquidationInfo) GetLiqEndTime() string {
	if o == nil || isNil(o.LiqEndTime) {
		var ret string
		return ret
	}
	return *o.LiqEndTime
}

// GetLiqEndTimeOk returns a tuple with the LiqEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginLiquidationInfo) GetLiqEndTimeOk() (*string, bool) {
	if o == nil || isNil(o.LiqEndTime) {
		return nil, false
	}
	return o.LiqEndTime, true
}

// HasLiqEndTime returns a boolean if a field has been set.
func (o *MarginLiquidationInfo) HasLiqEndTime() bool {
	if o != nil && !isNil(o.LiqEndTime) {
		return true
	}

	return false
}

// SetLiqEndTime gets a reference to the given string and assigns it to the LiqEndTime field.
func (o *MarginLiquidationInfo) SetLiqEndTime(v string) {
	o.LiqEndTime = &v
}

// GetLiqFee returns the LiqFee field value if set, zero value otherwise.
func (o *MarginLiquidationInfo) GetLiqFee() string {
	if o == nil || isNil(o.LiqFee) {
		var ret string
		return ret
	}
	return *o.LiqFee
}

// GetLiqFeeOk returns a tuple with the LiqFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginLiquidationInfo) GetLiqFeeOk() (*string, bool) {
	if o == nil || isNil(o.LiqFee) {
		return nil, false
	}
	return o.LiqFee, true
}

// HasLiqFee returns a boolean if a field has been set.
func (o *MarginLiquidationInfo) HasLiqFee() bool {
	if o != nil && !isNil(o.LiqFee) {
		return true
	}

	return false
}

// SetLiqFee gets a reference to the given string and assigns it to the LiqFee field.
func (o *MarginLiquidationInfo) SetLiqFee(v string) {
	o.LiqFee = &v
}

// GetLiqId returns the LiqId field value if set, zero value otherwise.
func (o *MarginLiquidationInfo) GetLiqId() string {
	if o == nil || isNil(o.LiqId) {
		var ret string
		return ret
	}
	return *o.LiqId
}

// GetLiqIdOk returns a tuple with the LiqId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginLiquidationInfo) GetLiqIdOk() (*string, bool) {
	if o == nil || isNil(o.LiqId) {
		return nil, false
	}
	return o.LiqId, true
}

// HasLiqId returns a boolean if a field has been set.
func (o *MarginLiquidationInfo) HasLiqId() bool {
	if o != nil && !isNil(o.LiqId) {
		return true
	}

	return false
}

// SetLiqId gets a reference to the given string and assigns it to the LiqId field.
func (o *MarginLiquidationInfo) SetLiqId(v string) {
	o.LiqId = &v
}

// GetLiqRisk returns the LiqRisk field value if set, zero value otherwise.
func (o *MarginLiquidationInfo) GetLiqRisk() string {
	if o == nil || isNil(o.LiqRisk) {
		var ret string
		return ret
	}
	return *o.LiqRisk
}

// GetLiqRiskOk returns a tuple with the LiqRisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginLiquidationInfo) GetLiqRiskOk() (*string, bool) {
	if o == nil || isNil(o.LiqRisk) {
		return nil, false
	}
	return o.LiqRisk, true
}

// HasLiqRisk returns a boolean if a field has been set.
func (o *MarginLiquidationInfo) HasLiqRisk() bool {
	if o != nil && !isNil(o.LiqRisk) {
		return true
	}

	return false
}

// SetLiqRisk gets a reference to the given string and assigns it to the LiqRisk field.
func (o *MarginLiquidationInfo) SetLiqRisk(v string) {
	o.LiqRisk = &v
}

// GetLiqStartTime returns the LiqStartTime field value if set, zero value otherwise.
func (o *MarginLiquidationInfo) GetLiqStartTime() string {
	if o == nil || isNil(o.LiqStartTime) {
		var ret string
		return ret
	}
	return *o.LiqStartTime
}

// GetLiqStartTimeOk returns a tuple with the LiqStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginLiquidationInfo) GetLiqStartTimeOk() (*string, bool) {
	if o == nil || isNil(o.LiqStartTime) {
		return nil, false
	}
	return o.LiqStartTime, true
}

// HasLiqStartTime returns a boolean if a field has been set.
func (o *MarginLiquidationInfo) HasLiqStartTime() bool {
	if o != nil && !isNil(o.LiqStartTime) {
		return true
	}

	return false
}

// SetLiqStartTime gets a reference to the given string and assigns it to the LiqStartTime field.
func (o *MarginLiquidationInfo) SetLiqStartTime(v string) {
	o.LiqStartTime = &v
}

// GetTotalAssets returns the TotalAssets field value if set, zero value otherwise.
func (o *MarginLiquidationInfo) GetTotalAssets() string {
	if o == nil || isNil(o.TotalAssets) {
		var ret string
		return ret
	}
	return *o.TotalAssets
}

// GetTotalAssetsOk returns a tuple with the TotalAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginLiquidationInfo) GetTotalAssetsOk() (*string, bool) {
	if o == nil || isNil(o.TotalAssets) {
		return nil, false
	}
	return o.TotalAssets, true
}

// HasTotalAssets returns a boolean if a field has been set.
func (o *MarginLiquidationInfo) HasTotalAssets() bool {
	if o != nil && !isNil(o.TotalAssets) {
		return true
	}

	return false
}

// SetTotalAssets gets a reference to the given string and assigns it to the TotalAssets field.
func (o *MarginLiquidationInfo) SetTotalAssets(v string) {
	o.TotalAssets = &v
}

// GetTotalDebt returns the TotalDebt field value if set, zero value otherwise.
func (o *MarginLiquidationInfo) GetTotalDebt() string {
	if o == nil || isNil(o.TotalDebt) {
		var ret string
		return ret
	}
	return *o.TotalDebt
}

// GetTotalDebtOk returns a tuple with the TotalDebt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginLiquidationInfo) GetTotalDebtOk() (*string, bool) {
	if o == nil || isNil(o.TotalDebt) {
		return nil, false
	}
	return o.TotalDebt, true
}

// HasTotalDebt returns a boolean if a field has been set.
func (o *MarginLiquidationInfo) HasTotalDebt() bool {
	if o != nil && !isNil(o.TotalDebt) {
		return true
	}

	return false
}

// SetTotalDebt gets a reference to the given string and assigns it to the TotalDebt field.
func (o *MarginLiquidationInfo) SetTotalDebt(v string) {
	o.TotalDebt = &v
}

func (o MarginLiquidationInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ctime) {
		toSerialize["ctime"] = o.Ctime
	}
	if !isNil(o.LiqEndTime) {
		toSerialize["liqEndTime"] = o.LiqEndTime
	}
	if !isNil(o.LiqFee) {
		toSerialize["liqFee"] = o.LiqFee
	}
	if !isNil(o.LiqId) {
		toSerialize["liqId"] = o.LiqId
	}
	if !isNil(o.LiqRisk) {
		toSerialize["liqRisk"] = o.LiqRisk
	}
	if !isNil(o.LiqStartTime) {
		toSerialize["liqStartTime"] = o.LiqStartTime
	}
	if !isNil(o.TotalAssets) {
		toSerialize["totalAssets"] = o.TotalAssets
	}
	if !isNil(o.TotalDebt) {
		toSerialize["totalDebt"] = o.TotalDebt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MarginLiquidationInfo) UnmarshalJSON(bytes []byte) (err error) {
	varMarginLiquidationInfo := _MarginLiquidationInfo{}

	if err = json.Unmarshal(bytes, &varMarginLiquidationInfo); err == nil {
		*o = MarginLiquidationInfo(varMarginLiquidationInfo)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ctime")
		delete(additionalProperties, "liqEndTime")
		delete(additionalProperties, "liqFee")
		delete(additionalProperties, "liqId")
		delete(additionalProperties, "liqRisk")
		delete(additionalProperties, "liqStartTime")
		delete(additionalProperties, "totalAssets")
		delete(additionalProperties, "totalDebt")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMarginLiquidationInfo struct {
	value *MarginLiquidationInfo
	isSet bool
}

func (v NullableMarginLiquidationInfo) Get() *MarginLiquidationInfo {
	return v.value
}

func (v *NullableMarginLiquidationInfo) Set(val *MarginLiquidationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginLiquidationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginLiquidationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginLiquidationInfo(val *MarginLiquidationInfo) *NullableMarginLiquidationInfo {
	return &NullableMarginLiquidationInfo{value: val, isSet: true}
}

func (v NullableMarginLiquidationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginLiquidationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}