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

// TraceSettingResult struct for TraceSettingResult
type TraceSettingResult struct {
	IsMyTrader           *bool                              `json:"isMyTrader,omitempty"`
	ProfitRate           *string                            `json:"profitRate,omitempty"`
	SettingType          *string                            `json:"settingType,omitempty"`
	SettledInDays        *string                            `json:"settledInDays,omitempty"`
	TraceBatchDetails    []TraceSettingBatchDetailsResult   `json:"traceBatchDetails,omitempty"`
	TraceProductConfigs  []TraceSettingProductConfigsResult `json:"traceProductConfigs,omitempty"`
	TraderHeadPic        *string                            `json:"traderHeadPic,omitempty"`
	TraderNickName       *string                            `json:"traderNickName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TraceSettingResult TraceSettingResult

// NewTraceSettingResult instantiates a new TraceSettingResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTraceSettingResult() *TraceSettingResult {
	this := TraceSettingResult{}
	return &this
}

// NewTraceSettingResultWithDefaults instantiates a new TraceSettingResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTraceSettingResultWithDefaults() *TraceSettingResult {
	this := TraceSettingResult{}
	return &this
}

// GetIsMyTrader returns the IsMyTrader field value if set, zero value otherwise.
func (o *TraceSettingResult) GetIsMyTrader() bool {
	if o == nil || isNil(o.IsMyTrader) {
		var ret bool
		return ret
	}
	return *o.IsMyTrader
}

// GetIsMyTraderOk returns a tuple with the IsMyTrader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceSettingResult) GetIsMyTraderOk() (*bool, bool) {
	if o == nil || isNil(o.IsMyTrader) {
		return nil, false
	}
	return o.IsMyTrader, true
}

// HasIsMyTrader returns a boolean if a field has been set.
func (o *TraceSettingResult) HasIsMyTrader() bool {
	if o != nil && !isNil(o.IsMyTrader) {
		return true
	}

	return false
}

// SetIsMyTrader gets a reference to the given bool and assigns it to the IsMyTrader field.
func (o *TraceSettingResult) SetIsMyTrader(v bool) {
	o.IsMyTrader = &v
}

// GetProfitRate returns the ProfitRate field value if set, zero value otherwise.
func (o *TraceSettingResult) GetProfitRate() string {
	if o == nil || isNil(o.ProfitRate) {
		var ret string
		return ret
	}
	return *o.ProfitRate
}

// GetProfitRateOk returns a tuple with the ProfitRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceSettingResult) GetProfitRateOk() (*string, bool) {
	if o == nil || isNil(o.ProfitRate) {
		return nil, false
	}
	return o.ProfitRate, true
}

// HasProfitRate returns a boolean if a field has been set.
func (o *TraceSettingResult) HasProfitRate() bool {
	if o != nil && !isNil(o.ProfitRate) {
		return true
	}

	return false
}

// SetProfitRate gets a reference to the given string and assigns it to the ProfitRate field.
func (o *TraceSettingResult) SetProfitRate(v string) {
	o.ProfitRate = &v
}

// GetSettingType returns the SettingType field value if set, zero value otherwise.
func (o *TraceSettingResult) GetSettingType() string {
	if o == nil || isNil(o.SettingType) {
		var ret string
		return ret
	}
	return *o.SettingType
}

// GetSettingTypeOk returns a tuple with the SettingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceSettingResult) GetSettingTypeOk() (*string, bool) {
	if o == nil || isNil(o.SettingType) {
		return nil, false
	}
	return o.SettingType, true
}

// HasSettingType returns a boolean if a field has been set.
func (o *TraceSettingResult) HasSettingType() bool {
	if o != nil && !isNil(o.SettingType) {
		return true
	}

	return false
}

// SetSettingType gets a reference to the given string and assigns it to the SettingType field.
func (o *TraceSettingResult) SetSettingType(v string) {
	o.SettingType = &v
}

// GetSettledInDays returns the SettledInDays field value if set, zero value otherwise.
func (o *TraceSettingResult) GetSettledInDays() string {
	if o == nil || isNil(o.SettledInDays) {
		var ret string
		return ret
	}
	return *o.SettledInDays
}

// GetSettledInDaysOk returns a tuple with the SettledInDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceSettingResult) GetSettledInDaysOk() (*string, bool) {
	if o == nil || isNil(o.SettledInDays) {
		return nil, false
	}
	return o.SettledInDays, true
}

// HasSettledInDays returns a boolean if a field has been set.
func (o *TraceSettingResult) HasSettledInDays() bool {
	if o != nil && !isNil(o.SettledInDays) {
		return true
	}

	return false
}

// SetSettledInDays gets a reference to the given string and assigns it to the SettledInDays field.
func (o *TraceSettingResult) SetSettledInDays(v string) {
	o.SettledInDays = &v
}

// GetTraceBatchDetails returns the TraceBatchDetails field value if set, zero value otherwise.
func (o *TraceSettingResult) GetTraceBatchDetails() []TraceSettingBatchDetailsResult {
	if o == nil || isNil(o.TraceBatchDetails) {
		var ret []TraceSettingBatchDetailsResult
		return ret
	}
	return o.TraceBatchDetails
}

// GetTraceBatchDetailsOk returns a tuple with the TraceBatchDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceSettingResult) GetTraceBatchDetailsOk() ([]TraceSettingBatchDetailsResult, bool) {
	if o == nil || isNil(o.TraceBatchDetails) {
		return nil, false
	}
	return o.TraceBatchDetails, true
}

// HasTraceBatchDetails returns a boolean if a field has been set.
func (o *TraceSettingResult) HasTraceBatchDetails() bool {
	if o != nil && !isNil(o.TraceBatchDetails) {
		return true
	}

	return false
}

// SetTraceBatchDetails gets a reference to the given []TraceSettingBatchDetailsResult and assigns it to the TraceBatchDetails field.
func (o *TraceSettingResult) SetTraceBatchDetails(v []TraceSettingBatchDetailsResult) {
	o.TraceBatchDetails = v
}

// GetTraceProductConfigs returns the TraceProductConfigs field value if set, zero value otherwise.
func (o *TraceSettingResult) GetTraceProductConfigs() []TraceSettingProductConfigsResult {
	if o == nil || isNil(o.TraceProductConfigs) {
		var ret []TraceSettingProductConfigsResult
		return ret
	}
	return o.TraceProductConfigs
}

// GetTraceProductConfigsOk returns a tuple with the TraceProductConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceSettingResult) GetTraceProductConfigsOk() ([]TraceSettingProductConfigsResult, bool) {
	if o == nil || isNil(o.TraceProductConfigs) {
		return nil, false
	}
	return o.TraceProductConfigs, true
}

// HasTraceProductConfigs returns a boolean if a field has been set.
func (o *TraceSettingResult) HasTraceProductConfigs() bool {
	if o != nil && !isNil(o.TraceProductConfigs) {
		return true
	}

	return false
}

// SetTraceProductConfigs gets a reference to the given []TraceSettingProductConfigsResult and assigns it to the TraceProductConfigs field.
func (o *TraceSettingResult) SetTraceProductConfigs(v []TraceSettingProductConfigsResult) {
	o.TraceProductConfigs = v
}

// GetTraderHeadPic returns the TraderHeadPic field value if set, zero value otherwise.
func (o *TraceSettingResult) GetTraderHeadPic() string {
	if o == nil || isNil(o.TraderHeadPic) {
		var ret string
		return ret
	}
	return *o.TraderHeadPic
}

// GetTraderHeadPicOk returns a tuple with the TraderHeadPic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceSettingResult) GetTraderHeadPicOk() (*string, bool) {
	if o == nil || isNil(o.TraderHeadPic) {
		return nil, false
	}
	return o.TraderHeadPic, true
}

// HasTraderHeadPic returns a boolean if a field has been set.
func (o *TraceSettingResult) HasTraderHeadPic() bool {
	if o != nil && !isNil(o.TraderHeadPic) {
		return true
	}

	return false
}

// SetTraderHeadPic gets a reference to the given string and assigns it to the TraderHeadPic field.
func (o *TraceSettingResult) SetTraderHeadPic(v string) {
	o.TraderHeadPic = &v
}

// GetTraderNickName returns the TraderNickName field value if set, zero value otherwise.
func (o *TraceSettingResult) GetTraderNickName() string {
	if o == nil || isNil(o.TraderNickName) {
		var ret string
		return ret
	}
	return *o.TraderNickName
}

// GetTraderNickNameOk returns a tuple with the TraderNickName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceSettingResult) GetTraderNickNameOk() (*string, bool) {
	if o == nil || isNil(o.TraderNickName) {
		return nil, false
	}
	return o.TraderNickName, true
}

// HasTraderNickName returns a boolean if a field has been set.
func (o *TraceSettingResult) HasTraderNickName() bool {
	if o != nil && !isNil(o.TraderNickName) {
		return true
	}

	return false
}

// SetTraderNickName gets a reference to the given string and assigns it to the TraderNickName field.
func (o *TraceSettingResult) SetTraderNickName(v string) {
	o.TraderNickName = &v
}

func (o TraceSettingResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IsMyTrader) {
		toSerialize["isMyTrader"] = o.IsMyTrader
	}
	if !isNil(o.ProfitRate) {
		toSerialize["profitRate"] = o.ProfitRate
	}
	if !isNil(o.SettingType) {
		toSerialize["settingType"] = o.SettingType
	}
	if !isNil(o.SettledInDays) {
		toSerialize["settledInDays"] = o.SettledInDays
	}
	if !isNil(o.TraceBatchDetails) {
		toSerialize["traceBatchDetails"] = o.TraceBatchDetails
	}
	if !isNil(o.TraceProductConfigs) {
		toSerialize["traceProductConfigs"] = o.TraceProductConfigs
	}
	if !isNil(o.TraderHeadPic) {
		toSerialize["traderHeadPic"] = o.TraderHeadPic
	}
	if !isNil(o.TraderNickName) {
		toSerialize["traderNickName"] = o.TraderNickName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TraceSettingResult) UnmarshalJSON(bytes []byte) (err error) {
	varTraceSettingResult := _TraceSettingResult{}

	if err = json.Unmarshal(bytes, &varTraceSettingResult); err == nil {
		*o = TraceSettingResult(varTraceSettingResult)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "isMyTrader")
		delete(additionalProperties, "profitRate")
		delete(additionalProperties, "settingType")
		delete(additionalProperties, "settledInDays")
		delete(additionalProperties, "traceBatchDetails")
		delete(additionalProperties, "traceProductConfigs")
		delete(additionalProperties, "traderHeadPic")
		delete(additionalProperties, "traderNickName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTraceSettingResult struct {
	value *TraceSettingResult
	isSet bool
}

func (v NullableTraceSettingResult) Get() *TraceSettingResult {
	return v.value
}

func (v *NullableTraceSettingResult) Set(val *TraceSettingResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTraceSettingResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTraceSettingResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTraceSettingResult(val *TraceSettingResult) *NullableTraceSettingResult {
	return &NullableTraceSettingResult{value: val, isSet: true}
}

func (v NullableTraceSettingResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTraceSettingResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
