/*
Account API

The <b>Account API</b> gives sellers the ability to configure their eBay seller accounts, including the seller's policies (the Fulfillment Policy, Payment Policy, and Return Policy), opt in and out of eBay seller programs, configure sales tax tables, and get account information.  <br><br>For details on the availability of the methods in this API, see <a href=\"/api-docs/sell/account/overview.html#requirements\">Account API requirements and restrictions</a>.

API version: v1.6.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ebayaccount

import (
	"encoding/json"
)

// RateTableResponse The response container for with information on a seller's shipping rate tables.
type RateTableResponse struct {
	// A list of elements that provide information on the seller-defined shipping rate tables.
	RateTables *[]RateTable `json:"rateTables,omitempty"`
}

// NewRateTableResponse instantiates a new RateTableResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRateTableResponse() *RateTableResponse {
	this := RateTableResponse{}
	return &this
}

// NewRateTableResponseWithDefaults instantiates a new RateTableResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRateTableResponseWithDefaults() *RateTableResponse {
	this := RateTableResponse{}
	return &this
}

// GetRateTables returns the RateTables field value if set, zero value otherwise.
func (o *RateTableResponse) GetRateTables() []RateTable {
	if o == nil || o.RateTables == nil {
		var ret []RateTable
		return ret
	}
	return *o.RateTables
}

// GetRateTablesOk returns a tuple with the RateTables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateTableResponse) GetRateTablesOk() (*[]RateTable, bool) {
	if o == nil || o.RateTables == nil {
		return nil, false
	}
	return o.RateTables, true
}

// HasRateTables returns a boolean if a field has been set.
func (o *RateTableResponse) HasRateTables() bool {
	if o != nil && o.RateTables != nil {
		return true
	}

	return false
}

// SetRateTables gets a reference to the given []RateTable and assigns it to the RateTables field.
func (o *RateTableResponse) SetRateTables(v []RateTable) {
	o.RateTables = &v
}

func (o RateTableResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RateTables != nil {
		toSerialize["rateTables"] = o.RateTables
	}
	return json.Marshal(toSerialize)
}

type NullableRateTableResponse struct {
	value *RateTableResponse
	isSet bool
}

func (v NullableRateTableResponse) Get() *RateTableResponse {
	return v.value
}

func (v *NullableRateTableResponse) Set(val *RateTableResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRateTableResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRateTableResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRateTableResponse(val *RateTableResponse) *NullableRateTableResponse {
	return &NullableRateTableResponse{value: val, isSet: true}
}

func (v NullableRateTableResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRateTableResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


