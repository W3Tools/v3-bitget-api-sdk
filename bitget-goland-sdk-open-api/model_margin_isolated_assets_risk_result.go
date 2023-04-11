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

// MarginIsolatedAssetsRiskResult struct for MarginIsolatedAssetsRiskResult
type MarginIsolatedAssetsRiskResult struct {
	RiskRate             *string `json:"riskRate,omitempty"`
	Symbol               *string `json:"symbol,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarginIsolatedAssetsRiskResult MarginIsolatedAssetsRiskResult

// NewMarginIsolatedAssetsRiskResult instantiates a new MarginIsolatedAssetsRiskResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginIsolatedAssetsRiskResult() *MarginIsolatedAssetsRiskResult {
	this := MarginIsolatedAssetsRiskResult{}
	return &this
}

// NewMarginIsolatedAssetsRiskResultWithDefaults instantiates a new MarginIsolatedAssetsRiskResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginIsolatedAssetsRiskResultWithDefaults() *MarginIsolatedAssetsRiskResult {
	this := MarginIsolatedAssetsRiskResult{}
	return &this
}

// GetRiskRate returns the RiskRate field value if set, zero value otherwise.
func (o *MarginIsolatedAssetsRiskResult) GetRiskRate() string {
	if o == nil || isNil(o.RiskRate) {
		var ret string
		return ret
	}
	return *o.RiskRate
}

// GetRiskRateOk returns a tuple with the RiskRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginIsolatedAssetsRiskResult) GetRiskRateOk() (*string, bool) {
	if o == nil || isNil(o.RiskRate) {
		return nil, false
	}
	return o.RiskRate, true
}

// HasRiskRate returns a boolean if a field has been set.
func (o *MarginIsolatedAssetsRiskResult) HasRiskRate() bool {
	if o != nil && !isNil(o.RiskRate) {
		return true
	}

	return false
}

// SetRiskRate gets a reference to the given string and assigns it to the RiskRate field.
func (o *MarginIsolatedAssetsRiskResult) SetRiskRate(v string) {
	o.RiskRate = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *MarginIsolatedAssetsRiskResult) GetSymbol() string {
	if o == nil || isNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginIsolatedAssetsRiskResult) GetSymbolOk() (*string, bool) {
	if o == nil || isNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *MarginIsolatedAssetsRiskResult) HasSymbol() bool {
	if o != nil && !isNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *MarginIsolatedAssetsRiskResult) SetSymbol(v string) {
	o.Symbol = &v
}

func (o MarginIsolatedAssetsRiskResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RiskRate) {
		toSerialize["riskRate"] = o.RiskRate
	}
	if !isNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MarginIsolatedAssetsRiskResult) UnmarshalJSON(bytes []byte) (err error) {
	varMarginIsolatedAssetsRiskResult := _MarginIsolatedAssetsRiskResult{}

	if err = json.Unmarshal(bytes, &varMarginIsolatedAssetsRiskResult); err == nil {
		*o = MarginIsolatedAssetsRiskResult(varMarginIsolatedAssetsRiskResult)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "riskRate")
		delete(additionalProperties, "symbol")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMarginIsolatedAssetsRiskResult struct {
	value *MarginIsolatedAssetsRiskResult
	isSet bool
}

func (v NullableMarginIsolatedAssetsRiskResult) Get() *MarginIsolatedAssetsRiskResult {
	return v.value
}

func (v *NullableMarginIsolatedAssetsRiskResult) Set(val *MarginIsolatedAssetsRiskResult) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginIsolatedAssetsRiskResult) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginIsolatedAssetsRiskResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginIsolatedAssetsRiskResult(val *MarginIsolatedAssetsRiskResult) *NullableMarginIsolatedAssetsRiskResult {
	return &NullableMarginIsolatedAssetsRiskResult{value: val, isSet: true}
}

func (v NullableMarginIsolatedAssetsRiskResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginIsolatedAssetsRiskResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}