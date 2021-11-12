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

// SalesTaxBase A container that describes the how the sales tax rate is calculated.
type SalesTaxBase struct {
	// The sales tax rate, as a percentage of the sale.
	SalesTaxPercentage *string `json:"salesTaxPercentage,omitempty"`
	// If set to <code>true</code>, shipping and handling charges are taxed.
	ShippingAndHandlingTaxed *bool `json:"shippingAndHandlingTaxed,omitempty"`
}

// NewSalesTaxBase instantiates a new SalesTaxBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSalesTaxBase() *SalesTaxBase {
	this := SalesTaxBase{}
	return &this
}

// NewSalesTaxBaseWithDefaults instantiates a new SalesTaxBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSalesTaxBaseWithDefaults() *SalesTaxBase {
	this := SalesTaxBase{}
	return &this
}

// GetSalesTaxPercentage returns the SalesTaxPercentage field value if set, zero value otherwise.
func (o *SalesTaxBase) GetSalesTaxPercentage() string {
	if o == nil || o.SalesTaxPercentage == nil {
		var ret string
		return ret
	}
	return *o.SalesTaxPercentage
}

// GetSalesTaxPercentageOk returns a tuple with the SalesTaxPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesTaxBase) GetSalesTaxPercentageOk() (*string, bool) {
	if o == nil || o.SalesTaxPercentage == nil {
		return nil, false
	}
	return o.SalesTaxPercentage, true
}

// HasSalesTaxPercentage returns a boolean if a field has been set.
func (o *SalesTaxBase) HasSalesTaxPercentage() bool {
	if o != nil && o.SalesTaxPercentage != nil {
		return true
	}

	return false
}

// SetSalesTaxPercentage gets a reference to the given string and assigns it to the SalesTaxPercentage field.
func (o *SalesTaxBase) SetSalesTaxPercentage(v string) {
	o.SalesTaxPercentage = &v
}

// GetShippingAndHandlingTaxed returns the ShippingAndHandlingTaxed field value if set, zero value otherwise.
func (o *SalesTaxBase) GetShippingAndHandlingTaxed() bool {
	if o == nil || o.ShippingAndHandlingTaxed == nil {
		var ret bool
		return ret
	}
	return *o.ShippingAndHandlingTaxed
}

// GetShippingAndHandlingTaxedOk returns a tuple with the ShippingAndHandlingTaxed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesTaxBase) GetShippingAndHandlingTaxedOk() (*bool, bool) {
	if o == nil || o.ShippingAndHandlingTaxed == nil {
		return nil, false
	}
	return o.ShippingAndHandlingTaxed, true
}

// HasShippingAndHandlingTaxed returns a boolean if a field has been set.
func (o *SalesTaxBase) HasShippingAndHandlingTaxed() bool {
	if o != nil && o.ShippingAndHandlingTaxed != nil {
		return true
	}

	return false
}

// SetShippingAndHandlingTaxed gets a reference to the given bool and assigns it to the ShippingAndHandlingTaxed field.
func (o *SalesTaxBase) SetShippingAndHandlingTaxed(v bool) {
	o.ShippingAndHandlingTaxed = &v
}

func (o SalesTaxBase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SalesTaxPercentage != nil {
		toSerialize["salesTaxPercentage"] = o.SalesTaxPercentage
	}
	if o.ShippingAndHandlingTaxed != nil {
		toSerialize["shippingAndHandlingTaxed"] = o.ShippingAndHandlingTaxed
	}
	return json.Marshal(toSerialize)
}

type NullableSalesTaxBase struct {
	value *SalesTaxBase
	isSet bool
}

func (v NullableSalesTaxBase) Get() *SalesTaxBase {
	return v.value
}

func (v *NullableSalesTaxBase) Set(val *SalesTaxBase) {
	v.value = val
	v.isSet = true
}

func (v NullableSalesTaxBase) IsSet() bool {
	return v.isSet
}

func (v *NullableSalesTaxBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSalesTaxBase(val *SalesTaxBase) *NullableSalesTaxBase {
	return &NullableSalesTaxBase{value: val, isSet: true}
}

func (v NullableSalesTaxBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSalesTaxBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


