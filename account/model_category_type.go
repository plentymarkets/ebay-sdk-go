/*
Account API

The <b>Account API</b> gives sellers the ability to configure their eBay seller accounts, including the seller's policies (seller-defined custom policies and eBay business policies), opt in and out of eBay seller programs, configure sales tax tables, and get account information.  <br><br>For details on the availability of the methods in this API, see <a href=\"/api-docs/sell/account/overview.html#requirements\">Account API requirements and restrictions</a>.

API version: v1.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package account

import (
	"encoding/json"
)

// CategoryType The category type discerns whether the policy covers the sale of motor vehicles (via eBay Motors), or the sale of everything except motor vehicles. <br><br>Each business policy can be associated with either or both categories ('MOTORS_VEHICLES' and 'ALL_EXCLUDING_MOTORS_VEHICLES'); however,the 'MOTORS_VEHICLES' category type is not valid for return policies&ndash;return policies cannot be used with motor vehicle listings.
type CategoryType struct {
	// Specifies the default policy for a <b>marketplaceId</b> and <b>categoryTypes.name</b> pair. Sellers can create multiple policies for any <b>marketplaceId</b> and <b>categoryTypes.name</b> combination. For example, you can create multiple fulfillment policies for one marketplace, where they all target the same category type <b>name</b>. However, only one policy can be the default for any <b>marketplaceId</b> and <b>name</b> combination, and eBay designates the first policy created for a combination as the default.  <br><br>If set to <code>true</code>, this policy is the default policy for the associated <b>categoryTypes.name</b> and <b>marketplaceId</b> pair.<br><br><span class=\"tablenote\"><b>Note</b>: eBay considers the status of this field only when you create listings through the Web flow. If you create listings using the APIs, you must specifically set the policies you want applied to a listing in the payload of the call you use to create the listing. If you use the Web flow to create item listings, eBay uses the default policy for the marketplace and category type specified, unless you override the default.</span> <br><br>For more on default policies, see <a href=\"/api-docs/sell/static/seller-accounts/business-policies.html#default_policy\">Changing the default policy for a category type</a>.
	Default *bool `json:"default,omitempty"`
	// The category type to which the policy applies (motor vehicles or non-motor vehicles). <br /><br /><b>Restrictions</b>: <ul><li>The <code>MOTORS_VEHICLES</code> category type is not valid for return policies. eBay flows do not support the return of motor vehicles.</li><li>Only the <code>ALL_EXCLUDING_MOTORS_VEHICLES</code> category type is supported for sellers who opt-in to the <a href=\"/managed-payments\" title=\"eBay Developers Program page\" target=\"_blank\">managed payments</a> program. Managed payments does not currently support the sale of motor vehicles.</li></ul> For implementation help, refer to <a href='https://developer.ebay.com/api-docs/sell/account/types/api:CategoryTypeEnum'>eBay API documentation</a>
	Name *string `json:"name,omitempty"`
}

// NewCategoryType instantiates a new CategoryType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryType() *CategoryType {
	this := CategoryType{}
	return &this
}

// NewCategoryTypeWithDefaults instantiates a new CategoryType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryTypeWithDefaults() *CategoryType {
	this := CategoryType{}
	return &this
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *CategoryType) GetDefault() bool {
	if o == nil || o.Default == nil {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryType) GetDefaultOk() (*bool, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *CategoryType) HasDefault() bool {
	if o != nil && o.Default != nil {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *CategoryType) SetDefault(v bool) {
	o.Default = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CategoryType) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryType) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CategoryType) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CategoryType) SetName(v string) {
	o.Name = &v
}

func (o CategoryType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableCategoryType struct {
	value *CategoryType
	isSet bool
}

func (v NullableCategoryType) Get() *CategoryType {
	return v.value
}

func (v *NullableCategoryType) Set(val *CategoryType) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryType) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryType(val *CategoryType) *NullableCategoryType {
	return &NullableCategoryType{value: val, isSet: true}
}

func (v NullableCategoryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoryType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


