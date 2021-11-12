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

// SalesTax The applicable sales tax rate, as a percentage of the sale amount, for a given country and sales tax jurisdiction within that country.
type SalesTax struct {
	// The country code identifying the country to which this tax rate applies. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/sell/account/types/ba:CountryCodeEnum'>eBay API documentation</a>
	CountryCode *string `json:"countryCode,omitempty"`
	// A unique ID that identifies the sales tax jurisdiction to which the tax rate applies (for example a state within the United States).
	SalesTaxJurisdictionId *string `json:"salesTaxJurisdictionId,omitempty"`
	// The sales tax rate (as a percentage of the sale) applied to sales transactions made in this country and sales tax jurisdiction.
	SalesTaxPercentage *string `json:"salesTaxPercentage,omitempty"`
	// If set to <code>true</code>, shipping and handling charges are taxed.
	ShippingAndHandlingTaxed *bool `json:"shippingAndHandlingTaxed,omitempty"`
}

// NewSalesTax instantiates a new SalesTax object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSalesTax() *SalesTax {
	this := SalesTax{}
	return &this
}

// NewSalesTaxWithDefaults instantiates a new SalesTax object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSalesTaxWithDefaults() *SalesTax {
	this := SalesTax{}
	return &this
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *SalesTax) GetCountryCode() string {
	if o == nil || o.CountryCode == nil {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesTax) GetCountryCodeOk() (*string, bool) {
	if o == nil || o.CountryCode == nil {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *SalesTax) HasCountryCode() bool {
	if o != nil && o.CountryCode != nil {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *SalesTax) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetSalesTaxJurisdictionId returns the SalesTaxJurisdictionId field value if set, zero value otherwise.
func (o *SalesTax) GetSalesTaxJurisdictionId() string {
	if o == nil || o.SalesTaxJurisdictionId == nil {
		var ret string
		return ret
	}
	return *o.SalesTaxJurisdictionId
}

// GetSalesTaxJurisdictionIdOk returns a tuple with the SalesTaxJurisdictionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesTax) GetSalesTaxJurisdictionIdOk() (*string, bool) {
	if o == nil || o.SalesTaxJurisdictionId == nil {
		return nil, false
	}
	return o.SalesTaxJurisdictionId, true
}

// HasSalesTaxJurisdictionId returns a boolean if a field has been set.
func (o *SalesTax) HasSalesTaxJurisdictionId() bool {
	if o != nil && o.SalesTaxJurisdictionId != nil {
		return true
	}

	return false
}

// SetSalesTaxJurisdictionId gets a reference to the given string and assigns it to the SalesTaxJurisdictionId field.
func (o *SalesTax) SetSalesTaxJurisdictionId(v string) {
	o.SalesTaxJurisdictionId = &v
}

// GetSalesTaxPercentage returns the SalesTaxPercentage field value if set, zero value otherwise.
func (o *SalesTax) GetSalesTaxPercentage() string {
	if o == nil || o.SalesTaxPercentage == nil {
		var ret string
		return ret
	}
	return *o.SalesTaxPercentage
}

// GetSalesTaxPercentageOk returns a tuple with the SalesTaxPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesTax) GetSalesTaxPercentageOk() (*string, bool) {
	if o == nil || o.SalesTaxPercentage == nil {
		return nil, false
	}
	return o.SalesTaxPercentage, true
}

// HasSalesTaxPercentage returns a boolean if a field has been set.
func (o *SalesTax) HasSalesTaxPercentage() bool {
	if o != nil && o.SalesTaxPercentage != nil {
		return true
	}

	return false
}

// SetSalesTaxPercentage gets a reference to the given string and assigns it to the SalesTaxPercentage field.
func (o *SalesTax) SetSalesTaxPercentage(v string) {
	o.SalesTaxPercentage = &v
}

// GetShippingAndHandlingTaxed returns the ShippingAndHandlingTaxed field value if set, zero value otherwise.
func (o *SalesTax) GetShippingAndHandlingTaxed() bool {
	if o == nil || o.ShippingAndHandlingTaxed == nil {
		var ret bool
		return ret
	}
	return *o.ShippingAndHandlingTaxed
}

// GetShippingAndHandlingTaxedOk returns a tuple with the ShippingAndHandlingTaxed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesTax) GetShippingAndHandlingTaxedOk() (*bool, bool) {
	if o == nil || o.ShippingAndHandlingTaxed == nil {
		return nil, false
	}
	return o.ShippingAndHandlingTaxed, true
}

// HasShippingAndHandlingTaxed returns a boolean if a field has been set.
func (o *SalesTax) HasShippingAndHandlingTaxed() bool {
	if o != nil && o.ShippingAndHandlingTaxed != nil {
		return true
	}

	return false
}

// SetShippingAndHandlingTaxed gets a reference to the given bool and assigns it to the ShippingAndHandlingTaxed field.
func (o *SalesTax) SetShippingAndHandlingTaxed(v bool) {
	o.ShippingAndHandlingTaxed = &v
}

func (o SalesTax) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CountryCode != nil {
		toSerialize["countryCode"] = o.CountryCode
	}
	if o.SalesTaxJurisdictionId != nil {
		toSerialize["salesTaxJurisdictionId"] = o.SalesTaxJurisdictionId
	}
	if o.SalesTaxPercentage != nil {
		toSerialize["salesTaxPercentage"] = o.SalesTaxPercentage
	}
	if o.ShippingAndHandlingTaxed != nil {
		toSerialize["shippingAndHandlingTaxed"] = o.ShippingAndHandlingTaxed
	}
	return json.Marshal(toSerialize)
}

type NullableSalesTax struct {
	value *SalesTax
	isSet bool
}

func (v NullableSalesTax) Get() *SalesTax {
	return v.value
}

func (v *NullableSalesTax) Set(val *SalesTax) {
	v.value = val
	v.isSet = true
}

func (v NullableSalesTax) IsSet() bool {
	return v.isSet
}

func (v *NullableSalesTax) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSalesTax(val *SalesTax) *NullableSalesTax {
	return &NullableSalesTax{value: val, isSet: true}
}

func (v NullableSalesTax) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSalesTax) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


