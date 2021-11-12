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

// PaymentPolicyRequest This root container defines a seller's payment policy for a specific marketplace and category type. Used when creating or updating a payment policy, <b>paymentPolicyRequest</b> encapsulates a seller's terms for how buyers can pay for the items they buy. While each seller must define at least one payment policy for every marketplace into which they sell, sellers can define multiple payment policies for a single marketplace by specifying different configurations for the unique policies.  <br><br>A successful call returns a <b>paymentPolicyId</b>, plus the <b>Location</b> response header contains the URI to the resource.  <br><br>Policy instructions can be localized by providing a locale in the <code>Content-Language</code> HTTP request header. For example: <code>Content-Language: de-DE</code>.  <p class=\"tablenote\"><b>Tip: </b>For more on using business policies, see <a href=\"/api-docs/sell/static/seller-accounts/business-policies.html\">eBay business policies</a>.</p>
type PaymentPolicyRequest struct {
	// The <b>CategoryTypeEnum</b> value to which this policy applies. This container is used to discern accounts that sell motor vehicles from those that do not.<br /><br /><b>Restriction:</b> Currently, each policy can be set to only one <b>categoryTypes</b> value at a time.
	CategoryTypes *[]CategoryType `json:"categoryTypes,omitempty"`
	Deposit *Deposit `json:"deposit,omitempty"`
	// An optional seller-defined description of the payment policy for internal use (this value is not displayed to end users).  <br><br><b>Max length</b>: 250
	Description *string `json:"description,omitempty"`
	FullPaymentDueIn *TimeDuration `json:"fullPaymentDueIn,omitempty"`
	// This field should be included and set to 'true' if the seller wants to require immediate payment from the buyer for a fixed-price item or for an auction item where the buyer is using the 'Buy it Now' option.<br /><br />Sellers who are not onboarded with eBay managed payments have the following additional requirements:<ul><li>Offer PayPal as the only payment method</li><li>Provide a valid PayPal payment email address in the <b>paymentMethods.recipientAccountReference.referenceId</b> field</li></ul><br /><br /><b>Default:</b> False
	ImmediatePay *bool `json:"immediatePay,omitempty"`
	// The ID of the eBay marketplace to which this payment policy applies. If this value is not specified, the value defaults to the seller's eBay registration site. <p class=\"tablenote\"><b>Note:</b> A limited number of sellers, on a limited number of eBay marketplaces, are currently opted-in to the eBay managed payments program. To view the eBay marketplaces where managed payments are currently supported, see the <a href=\"/managed-payments\" title=\"eBay Developers Program page\" target=\"_blank\">managed payments</a> landing page.</p> For implementation help, refer to <a href='https://developer.ebay.com/api-docs/sell/account/types/ba:MarketplaceIdEnum'>eBay API documentation</a>
	MarketplaceId *string `json:"marketplaceId,omitempty"`
	// A user-defined name for this payment policy. Names must be unique for policies assigned to the same marketplace. <p class=\"tablenote\"><b>Note:</b> eBay will create a new payment policy for sellers who opt-in to the <a href=\"/managed-payments\" title=\"eBay Developers Program page\" target=\"_blank\">managed payments</a> program.</p><b>Max length:</b> 64
	Name *string `json:"name,omitempty"`
	// <p class=\"tablenote\"><b>Note:</b> DO NOT USE THIS FIELD. Payment instructions are no longer supported by payment business policies.</p>A free-form string field that allows sellers to add detailed payment instructions to their listings. The payment instructions appear on eBay's View Item and Checkout pages. <br><br>eBay recommends sellers use this field to clarify payment policies for motor vehicle listings on eBay Motors. For example, sellers can include the specifics on the deposit (if required), pickup/delivery arrangements, and full payment details on the vehicle. <br><br>The field allows only 500 characters as input, but due to the way the eBay web site UI treats characters, this field can return more than 500 characters in the response. For example, characters like & and ' (ampersand and single quote) count as 5 characters each. <br /><br /><b>Restriction:</b> This container is not supported for sellers who opt-in to the <a href=\"/managed-payments\" title=\"eBay Developers Program page\" target=\"_blank\">managed payments</a> program.  <br><br><b>Max length:</b> 1000
	PaymentInstructions *string `json:"paymentInstructions,omitempty"`
	// This array is used to specify one or more payment methods that will be accepted for order payment.<br /><br />Because available electronic payments are strictly managed by eBay for sellers who are onboarded with eBay managed payments, this array is only needed by those sellers to specify offline payments (if applicable).<br /><br />For sellers who are not onboarded with managed payments, at least one electronic payment method is required. For most marketplaces and categories, this payment method is 'PAYPAL'.<br /><br />For a motor vehicle payment policy, the payment policy must specify at least one of the following as an offline payment method:<ul><li>CASH_ON_PICKUP (not available in the US)</li><li>CASHIER_CHECK</li><li>MONEY_ORDER</li><li>PERSONAL_CHECK</li></ul>
	PaymentMethods *[]PaymentMethod `json:"paymentMethods,omitempty"`
}

// NewPaymentPolicyRequest instantiates a new PaymentPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentPolicyRequest() *PaymentPolicyRequest {
	this := PaymentPolicyRequest{}
	return &this
}

// NewPaymentPolicyRequestWithDefaults instantiates a new PaymentPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentPolicyRequestWithDefaults() *PaymentPolicyRequest {
	this := PaymentPolicyRequest{}
	return &this
}

// GetCategoryTypes returns the CategoryTypes field value if set, zero value otherwise.
func (o *PaymentPolicyRequest) GetCategoryTypes() []CategoryType {
	if o == nil || o.CategoryTypes == nil {
		var ret []CategoryType
		return ret
	}
	return *o.CategoryTypes
}

// GetCategoryTypesOk returns a tuple with the CategoryTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentPolicyRequest) GetCategoryTypesOk() (*[]CategoryType, bool) {
	if o == nil || o.CategoryTypes == nil {
		return nil, false
	}
	return o.CategoryTypes, true
}

// HasCategoryTypes returns a boolean if a field has been set.
func (o *PaymentPolicyRequest) HasCategoryTypes() bool {
	if o != nil && o.CategoryTypes != nil {
		return true
	}

	return false
}

// SetCategoryTypes gets a reference to the given []CategoryType and assigns it to the CategoryTypes field.
func (o *PaymentPolicyRequest) SetCategoryTypes(v []CategoryType) {
	o.CategoryTypes = &v
}

// GetDeposit returns the Deposit field value if set, zero value otherwise.
func (o *PaymentPolicyRequest) GetDeposit() Deposit {
	if o == nil || o.Deposit == nil {
		var ret Deposit
		return ret
	}
	return *o.Deposit
}

// GetDepositOk returns a tuple with the Deposit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentPolicyRequest) GetDepositOk() (*Deposit, bool) {
	if o == nil || o.Deposit == nil {
		return nil, false
	}
	return o.Deposit, true
}

// HasDeposit returns a boolean if a field has been set.
func (o *PaymentPolicyRequest) HasDeposit() bool {
	if o != nil && o.Deposit != nil {
		return true
	}

	return false
}

// SetDeposit gets a reference to the given Deposit and assigns it to the Deposit field.
func (o *PaymentPolicyRequest) SetDeposit(v Deposit) {
	o.Deposit = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PaymentPolicyRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentPolicyRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PaymentPolicyRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PaymentPolicyRequest) SetDescription(v string) {
	o.Description = &v
}

// GetFullPaymentDueIn returns the FullPaymentDueIn field value if set, zero value otherwise.
func (o *PaymentPolicyRequest) GetFullPaymentDueIn() TimeDuration {
	if o == nil || o.FullPaymentDueIn == nil {
		var ret TimeDuration
		return ret
	}
	return *o.FullPaymentDueIn
}

// GetFullPaymentDueInOk returns a tuple with the FullPaymentDueIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentPolicyRequest) GetFullPaymentDueInOk() (*TimeDuration, bool) {
	if o == nil || o.FullPaymentDueIn == nil {
		return nil, false
	}
	return o.FullPaymentDueIn, true
}

// HasFullPaymentDueIn returns a boolean if a field has been set.
func (o *PaymentPolicyRequest) HasFullPaymentDueIn() bool {
	if o != nil && o.FullPaymentDueIn != nil {
		return true
	}

	return false
}

// SetFullPaymentDueIn gets a reference to the given TimeDuration and assigns it to the FullPaymentDueIn field.
func (o *PaymentPolicyRequest) SetFullPaymentDueIn(v TimeDuration) {
	o.FullPaymentDueIn = &v
}

// GetImmediatePay returns the ImmediatePay field value if set, zero value otherwise.
func (o *PaymentPolicyRequest) GetImmediatePay() bool {
	if o == nil || o.ImmediatePay == nil {
		var ret bool
		return ret
	}
	return *o.ImmediatePay
}

// GetImmediatePayOk returns a tuple with the ImmediatePay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentPolicyRequest) GetImmediatePayOk() (*bool, bool) {
	if o == nil || o.ImmediatePay == nil {
		return nil, false
	}
	return o.ImmediatePay, true
}

// HasImmediatePay returns a boolean if a field has been set.
func (o *PaymentPolicyRequest) HasImmediatePay() bool {
	if o != nil && o.ImmediatePay != nil {
		return true
	}

	return false
}

// SetImmediatePay gets a reference to the given bool and assigns it to the ImmediatePay field.
func (o *PaymentPolicyRequest) SetImmediatePay(v bool) {
	o.ImmediatePay = &v
}

// GetMarketplaceId returns the MarketplaceId field value if set, zero value otherwise.
func (o *PaymentPolicyRequest) GetMarketplaceId() string {
	if o == nil || o.MarketplaceId == nil {
		var ret string
		return ret
	}
	return *o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentPolicyRequest) GetMarketplaceIdOk() (*string, bool) {
	if o == nil || o.MarketplaceId == nil {
		return nil, false
	}
	return o.MarketplaceId, true
}

// HasMarketplaceId returns a boolean if a field has been set.
func (o *PaymentPolicyRequest) HasMarketplaceId() bool {
	if o != nil && o.MarketplaceId != nil {
		return true
	}

	return false
}

// SetMarketplaceId gets a reference to the given string and assigns it to the MarketplaceId field.
func (o *PaymentPolicyRequest) SetMarketplaceId(v string) {
	o.MarketplaceId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PaymentPolicyRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentPolicyRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PaymentPolicyRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PaymentPolicyRequest) SetName(v string) {
	o.Name = &v
}

// GetPaymentInstructions returns the PaymentInstructions field value if set, zero value otherwise.
func (o *PaymentPolicyRequest) GetPaymentInstructions() string {
	if o == nil || o.PaymentInstructions == nil {
		var ret string
		return ret
	}
	return *o.PaymentInstructions
}

// GetPaymentInstructionsOk returns a tuple with the PaymentInstructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentPolicyRequest) GetPaymentInstructionsOk() (*string, bool) {
	if o == nil || o.PaymentInstructions == nil {
		return nil, false
	}
	return o.PaymentInstructions, true
}

// HasPaymentInstructions returns a boolean if a field has been set.
func (o *PaymentPolicyRequest) HasPaymentInstructions() bool {
	if o != nil && o.PaymentInstructions != nil {
		return true
	}

	return false
}

// SetPaymentInstructions gets a reference to the given string and assigns it to the PaymentInstructions field.
func (o *PaymentPolicyRequest) SetPaymentInstructions(v string) {
	o.PaymentInstructions = &v
}

// GetPaymentMethods returns the PaymentMethods field value if set, zero value otherwise.
func (o *PaymentPolicyRequest) GetPaymentMethods() []PaymentMethod {
	if o == nil || o.PaymentMethods == nil {
		var ret []PaymentMethod
		return ret
	}
	return *o.PaymentMethods
}

// GetPaymentMethodsOk returns a tuple with the PaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentPolicyRequest) GetPaymentMethodsOk() (*[]PaymentMethod, bool) {
	if o == nil || o.PaymentMethods == nil {
		return nil, false
	}
	return o.PaymentMethods, true
}

// HasPaymentMethods returns a boolean if a field has been set.
func (o *PaymentPolicyRequest) HasPaymentMethods() bool {
	if o != nil && o.PaymentMethods != nil {
		return true
	}

	return false
}

// SetPaymentMethods gets a reference to the given []PaymentMethod and assigns it to the PaymentMethods field.
func (o *PaymentPolicyRequest) SetPaymentMethods(v []PaymentMethod) {
	o.PaymentMethods = &v
}

func (o PaymentPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CategoryTypes != nil {
		toSerialize["categoryTypes"] = o.CategoryTypes
	}
	if o.Deposit != nil {
		toSerialize["deposit"] = o.Deposit
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.FullPaymentDueIn != nil {
		toSerialize["fullPaymentDueIn"] = o.FullPaymentDueIn
	}
	if o.ImmediatePay != nil {
		toSerialize["immediatePay"] = o.ImmediatePay
	}
	if o.MarketplaceId != nil {
		toSerialize["marketplaceId"] = o.MarketplaceId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PaymentInstructions != nil {
		toSerialize["paymentInstructions"] = o.PaymentInstructions
	}
	if o.PaymentMethods != nil {
		toSerialize["paymentMethods"] = o.PaymentMethods
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentPolicyRequest struct {
	value *PaymentPolicyRequest
	isSet bool
}

func (v NullablePaymentPolicyRequest) Get() *PaymentPolicyRequest {
	return v.value
}

func (v *NullablePaymentPolicyRequest) Set(val *PaymentPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentPolicyRequest(val *PaymentPolicyRequest) *NullablePaymentPolicyRequest {
	return &NullablePaymentPolicyRequest{value: val, isSet: true}
}

func (v NullablePaymentPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


