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

// KycResponse This is the base response type of the <b>getKYC</b> method.
type KycResponse struct {
	// This array contains one or more KYC checks required from a managed payments seller. The seller may need to provide more documentation and/or information about themselves, their company, or the bank account they are using for seller payouts.<br/><br/>If no KYC checks are currently required from the seller, this array is not returned, and the seller only receives a <code>204 No Content</code> HTTP status code.
	KycChecks *[]KycCheck `json:"kycChecks,omitempty"`
}

// NewKycResponse instantiates a new KycResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKycResponse() *KycResponse {
	this := KycResponse{}
	return &this
}

// NewKycResponseWithDefaults instantiates a new KycResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKycResponseWithDefaults() *KycResponse {
	this := KycResponse{}
	return &this
}

// GetKycChecks returns the KycChecks field value if set, zero value otherwise.
func (o *KycResponse) GetKycChecks() []KycCheck {
	if o == nil || o.KycChecks == nil {
		var ret []KycCheck
		return ret
	}
	return *o.KycChecks
}

// GetKycChecksOk returns a tuple with the KycChecks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KycResponse) GetKycChecksOk() (*[]KycCheck, bool) {
	if o == nil || o.KycChecks == nil {
		return nil, false
	}
	return o.KycChecks, true
}

// HasKycChecks returns a boolean if a field has been set.
func (o *KycResponse) HasKycChecks() bool {
	if o != nil && o.KycChecks != nil {
		return true
	}

	return false
}

// SetKycChecks gets a reference to the given []KycCheck and assigns it to the KycChecks field.
func (o *KycResponse) SetKycChecks(v []KycCheck) {
	o.KycChecks = &v
}

func (o KycResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.KycChecks != nil {
		toSerialize["kycChecks"] = o.KycChecks
	}
	return json.Marshal(toSerialize)
}

type NullableKycResponse struct {
	value *KycResponse
	isSet bool
}

func (v NullableKycResponse) Get() *KycResponse {
	return v.value
}

func (v *NullableKycResponse) Set(val *KycResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKycResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKycResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKycResponse(val *KycResponse) *NullableKycResponse {
	return &NullableKycResponse{value: val, isSet: true}
}

func (v NullableKycResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKycResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


