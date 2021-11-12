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

// ErrorParameter A complex type that indicates a parameter that caused an error and the value of the parameter which caused the error.
type ErrorParameter struct {
	// Name of the parameter that caused the error.
	Name *string `json:"name,omitempty"`
	// The value of the parameter that caused the error.
	Value *string `json:"value,omitempty"`
}

// NewErrorParameter instantiates a new ErrorParameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorParameter() *ErrorParameter {
	this := ErrorParameter{}
	return &this
}

// NewErrorParameterWithDefaults instantiates a new ErrorParameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorParameterWithDefaults() *ErrorParameter {
	this := ErrorParameter{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ErrorParameter) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorParameter) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ErrorParameter) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ErrorParameter) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ErrorParameter) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorParameter) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ErrorParameter) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ErrorParameter) SetValue(v string) {
	o.Value = &v
}

func (o ErrorParameter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableErrorParameter struct {
	value *ErrorParameter
	isSet bool
}

func (v NullableErrorParameter) Get() *ErrorParameter {
	return v.value
}

func (v *NullableErrorParameter) Set(val *ErrorParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorParameter(val *ErrorParameter) *NullableErrorParameter {
	return &NullableErrorParameter{value: val, isSet: true}
}

func (v NullableErrorParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


