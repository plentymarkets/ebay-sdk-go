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

// PaymentsProgramOnboardingResponse Contains the payments program onboarding response
type PaymentsProgramOnboardingResponse struct {
	// This enumeration value indicates the eligibility of payment onboarding for the registered site. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/sell/account/types/api:PaymentsProgramOnboardingStatus'>eBay API documentation</a>
	OnboardingStatus *string `json:"onboardingStatus,omitempty"`
	// An array of the active process steps for payment onboarding and the status of each step. This array includes the step <strong>name</strong>, step <strong>status</strong>, and a <strong>webUrl</strong> to the <code>IN_PROGRESS</code> step. The step names are returned in sequential order. 
	Steps *[]PaymentsProgramOnboardingSteps `json:"steps,omitempty"`
}

// NewPaymentsProgramOnboardingResponse instantiates a new PaymentsProgramOnboardingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentsProgramOnboardingResponse() *PaymentsProgramOnboardingResponse {
	this := PaymentsProgramOnboardingResponse{}
	return &this
}

// NewPaymentsProgramOnboardingResponseWithDefaults instantiates a new PaymentsProgramOnboardingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentsProgramOnboardingResponseWithDefaults() *PaymentsProgramOnboardingResponse {
	this := PaymentsProgramOnboardingResponse{}
	return &this
}

// GetOnboardingStatus returns the OnboardingStatus field value if set, zero value otherwise.
func (o *PaymentsProgramOnboardingResponse) GetOnboardingStatus() string {
	if o == nil || o.OnboardingStatus == nil {
		var ret string
		return ret
	}
	return *o.OnboardingStatus
}

// GetOnboardingStatusOk returns a tuple with the OnboardingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentsProgramOnboardingResponse) GetOnboardingStatusOk() (*string, bool) {
	if o == nil || o.OnboardingStatus == nil {
		return nil, false
	}
	return o.OnboardingStatus, true
}

// HasOnboardingStatus returns a boolean if a field has been set.
func (o *PaymentsProgramOnboardingResponse) HasOnboardingStatus() bool {
	if o != nil && o.OnboardingStatus != nil {
		return true
	}

	return false
}

// SetOnboardingStatus gets a reference to the given string and assigns it to the OnboardingStatus field.
func (o *PaymentsProgramOnboardingResponse) SetOnboardingStatus(v string) {
	o.OnboardingStatus = &v
}

// GetSteps returns the Steps field value if set, zero value otherwise.
func (o *PaymentsProgramOnboardingResponse) GetSteps() []PaymentsProgramOnboardingSteps {
	if o == nil || o.Steps == nil {
		var ret []PaymentsProgramOnboardingSteps
		return ret
	}
	return *o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentsProgramOnboardingResponse) GetStepsOk() (*[]PaymentsProgramOnboardingSteps, bool) {
	if o == nil || o.Steps == nil {
		return nil, false
	}
	return o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *PaymentsProgramOnboardingResponse) HasSteps() bool {
	if o != nil && o.Steps != nil {
		return true
	}

	return false
}

// SetSteps gets a reference to the given []PaymentsProgramOnboardingSteps and assigns it to the Steps field.
func (o *PaymentsProgramOnboardingResponse) SetSteps(v []PaymentsProgramOnboardingSteps) {
	o.Steps = &v
}

func (o PaymentsProgramOnboardingResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OnboardingStatus != nil {
		toSerialize["onboardingStatus"] = o.OnboardingStatus
	}
	if o.Steps != nil {
		toSerialize["steps"] = o.Steps
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentsProgramOnboardingResponse struct {
	value *PaymentsProgramOnboardingResponse
	isSet bool
}

func (v NullablePaymentsProgramOnboardingResponse) Get() *PaymentsProgramOnboardingResponse {
	return v.value
}

func (v *NullablePaymentsProgramOnboardingResponse) Set(val *PaymentsProgramOnboardingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentsProgramOnboardingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentsProgramOnboardingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentsProgramOnboardingResponse(val *PaymentsProgramOnboardingResponse) *NullablePaymentsProgramOnboardingResponse {
	return &NullablePaymentsProgramOnboardingResponse{value: val, isSet: true}
}

func (v NullablePaymentsProgramOnboardingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentsProgramOnboardingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


