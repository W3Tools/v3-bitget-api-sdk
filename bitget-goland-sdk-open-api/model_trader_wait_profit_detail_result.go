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

// TraderWaitProfitDetailResult struct for TraderWaitProfitDetailResult
type TraderWaitProfitDetailResult struct {
	CoinName             *string `json:"coinName,omitempty"`
	DistributeRatio      *string `json:"distributeRatio,omitempty"`
	Profit               *string `json:"profit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TraderWaitProfitDetailResult TraderWaitProfitDetailResult

// NewTraderWaitProfitDetailResult instantiates a new TraderWaitProfitDetailResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTraderWaitProfitDetailResult() *TraderWaitProfitDetailResult {
	this := TraderWaitProfitDetailResult{}
	return &this
}

// NewTraderWaitProfitDetailResultWithDefaults instantiates a new TraderWaitProfitDetailResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTraderWaitProfitDetailResultWithDefaults() *TraderWaitProfitDetailResult {
	this := TraderWaitProfitDetailResult{}
	return &this
}

// GetCoinName returns the CoinName field value if set, zero value otherwise.
func (o *TraderWaitProfitDetailResult) GetCoinName() string {
	if o == nil || isNil(o.CoinName) {
		var ret string
		return ret
	}
	return *o.CoinName
}

// GetCoinNameOk returns a tuple with the CoinName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraderWaitProfitDetailResult) GetCoinNameOk() (*string, bool) {
	if o == nil || isNil(o.CoinName) {
		return nil, false
	}
	return o.CoinName, true
}

// HasCoinName returns a boolean if a field has been set.
func (o *TraderWaitProfitDetailResult) HasCoinName() bool {
	if o != nil && !isNil(o.CoinName) {
		return true
	}

	return false
}

// SetCoinName gets a reference to the given string and assigns it to the CoinName field.
func (o *TraderWaitProfitDetailResult) SetCoinName(v string) {
	o.CoinName = &v
}

// GetDistributeRatio returns the DistributeRatio field value if set, zero value otherwise.
func (o *TraderWaitProfitDetailResult) GetDistributeRatio() string {
	if o == nil || isNil(o.DistributeRatio) {
		var ret string
		return ret
	}
	return *o.DistributeRatio
}

// GetDistributeRatioOk returns a tuple with the DistributeRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraderWaitProfitDetailResult) GetDistributeRatioOk() (*string, bool) {
	if o == nil || isNil(o.DistributeRatio) {
		return nil, false
	}
	return o.DistributeRatio, true
}

// HasDistributeRatio returns a boolean if a field has been set.
func (o *TraderWaitProfitDetailResult) HasDistributeRatio() bool {
	if o != nil && !isNil(o.DistributeRatio) {
		return true
	}

	return false
}

// SetDistributeRatio gets a reference to the given string and assigns it to the DistributeRatio field.
func (o *TraderWaitProfitDetailResult) SetDistributeRatio(v string) {
	o.DistributeRatio = &v
}

// GetProfit returns the Profit field value if set, zero value otherwise.
func (o *TraderWaitProfitDetailResult) GetProfit() string {
	if o == nil || isNil(o.Profit) {
		var ret string
		return ret
	}
	return *o.Profit
}

// GetProfitOk returns a tuple with the Profit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraderWaitProfitDetailResult) GetProfitOk() (*string, bool) {
	if o == nil || isNil(o.Profit) {
		return nil, false
	}
	return o.Profit, true
}

// HasProfit returns a boolean if a field has been set.
func (o *TraderWaitProfitDetailResult) HasProfit() bool {
	if o != nil && !isNil(o.Profit) {
		return true
	}

	return false
}

// SetProfit gets a reference to the given string and assigns it to the Profit field.
func (o *TraderWaitProfitDetailResult) SetProfit(v string) {
	o.Profit = &v
}

func (o TraderWaitProfitDetailResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CoinName) {
		toSerialize["coinName"] = o.CoinName
	}
	if !isNil(o.DistributeRatio) {
		toSerialize["distributeRatio"] = o.DistributeRatio
	}
	if !isNil(o.Profit) {
		toSerialize["profit"] = o.Profit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TraderWaitProfitDetailResult) UnmarshalJSON(bytes []byte) (err error) {
	varTraderWaitProfitDetailResult := _TraderWaitProfitDetailResult{}

	if err = json.Unmarshal(bytes, &varTraderWaitProfitDetailResult); err == nil {
		*o = TraderWaitProfitDetailResult(varTraderWaitProfitDetailResult)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "coinName")
		delete(additionalProperties, "distributeRatio")
		delete(additionalProperties, "profit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTraderWaitProfitDetailResult struct {
	value *TraderWaitProfitDetailResult
	isSet bool
}

func (v NullableTraderWaitProfitDetailResult) Get() *TraderWaitProfitDetailResult {
	return v.value
}

func (v *NullableTraderWaitProfitDetailResult) Set(val *TraderWaitProfitDetailResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTraderWaitProfitDetailResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTraderWaitProfitDetailResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTraderWaitProfitDetailResult(val *TraderWaitProfitDetailResult) *NullableTraderWaitProfitDetailResult {
	return &NullableTraderWaitProfitDetailResult{value: val, isSet: true}
}

func (v NullableTraderWaitProfitDetailResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTraderWaitProfitDetailResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}