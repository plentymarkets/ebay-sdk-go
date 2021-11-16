/*
Account API

The <b>Account API</b> gives sellers the ability to configure their eBay seller accounts, including the seller's policies (the Fulfillment Policy, Payment Policy, and Return Policy), opt in and out of eBay seller programs, configure sales tax tables, and get account information.  <br><br>For details on the availability of the methods in this API, see <a href=\"/api-docs/sell/account/overview.html#requirements\">Account API requirements and restrictions</a>.

API version: v1.6.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package account

import (
	"encoding/json"
)

// PaymentPolicy Root container that defines the fields for a seller's payment policy. The <b>paymentPolicy</b> encapsulates a seller's payment terms and consists of payment details for the seller, the name and description of the policy, and the marketplace and category group(s) covered by the payment policy. While each seller must define at least one payment policy for every marketplace into which they sell, sellers can define multiple payment policies for a single marketplace by specifying different configurations for the unique policies.
type PaymentPolicy struct {
	// The <b>CategoryTypeEnum</b> value to which this policy applies. Used to discern accounts that sell motor vehicles from those that don't. (Currently, each policy can be set to only one <b>categoryTypes</b> value at a time.)
	CategoryTypes *[]CategoryType `json:"categoryTypes,omitempty"`
	Deposit *Deposit `json:"deposit,omitempty"`
	// An optional seller-defined description of the payment policy.  <br><br><b>Max length</b>: 250
	Description *string `json:"description,omitempty"`
	FullPaymentDueIn *TimeDuration `json:"fullPaymentDueIn,omitempty"`
	// If this field is returned as <code>true</code>, immediate payment is required from the buyer for a fixed-price item or for an auction item where the buyer uses the 'Buy it Now' option.<br /><br />It is possible for the seller to set this field as <code>true</code> in the payment policy, but it will not apply in some scenarios. For example, immediate payment is not applicable for auction listings that have a winning bidder, for buyer purchases that involve the Best Offer feature, or for transactions that happen offline between the buyer and seller.
	ImmediatePay *bool `json:"immediatePay,omitempty"`
	// The ID of the eBay marketplace to which the payment policy applies. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/sell/account/types/ba:MarketplaceIdEnum'>eBay API documentation</a>
	MarketplaceId *string `json:"marketplaceId,omitempty"`
	// A user-defined name for this payment policy. Names must be unique for policies assigned to the same marketplace. <br><br><b>Max length</b>: 64
	Name *string `json:"name,omitempty"`
	// This free-form string field gives sellers the ability add detailed payment instructions to their listings. The payment instructions appear on eBay's View Item and Checkout pages.  <br><br>eBay recommends sellers use this field to clarify payment policies for motor vehicle listings on eBay Motors. For example, sellers can include the specifics on the deposit (if required), pickup/delivery arrangements, and full payment details on the vehicle. <br><br>The field allows only 500 characters as input, but due to the way the eBay web site UI treats characters, this field can return more than 500 characters in the response. For example, characters like & and ' (ampersand and single quote) count as 5 characters each.  <br><br><b>Max length:</b> 1000.
	PaymentInstructions *string `json:"paymentInstructions,omitempty"`
	// This container is returned to show the payment methods that are accepted for the payment business policy.  <br><br>Sellers enabled for managed payments do not have to specify any electronic payment methods for listings, so this array will often be returned empty unless the payment business policy is intended for motor vehicle listings or other items in categories where offline payments are required or supported. 
	PaymentMethods *[]PaymentMethod `json:"paymentMethods,omitempty"`
	// A unique eBay-assigned ID for a payment policy. This ID is generated when the policy is created.
	PaymentPolicyId *string `json:"paymentPolicyId,omitempty"`
}

// NewPaymentPolicy instantiates a new PaymentPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentPolicy() *PaymentPolicy {
	this := PaymentPolicy{}
	return &this
}

// NewPaymentPolicyWithDefaults instantiates a new PaymentPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentPolicyWithDefaults() *PaymentPolicy {
	this := PaymentPolicy{}
	return &this
}

// GetCategoryTypes returns the CategoryTypes field value if set, zero value otherwise.
func (o *PaymentPolicy) GetCategoryTypes() []CategoryType {
	if o == nil || o.CategoryTypes == nil {
		var ret []CategoryType
		return ret
	}
	return *o.CategoryTypes
}

// GetCategoryTypesOk returns a tuple with the CategoryTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentPolicy) GetCategoryTypesOk() (*[]CategoryType, bool) {
	if o == nil || o.CategoryTypes == nil {
		return nil, false
	}
	return o.CategoryTypes, true
}

// HasCategoryTypes returns a boolean if a field has been set.
func (o *PaymentPolicy) HasCategoryTypes() bool {
	if o != nil && o.CategoryTypes != nil {
		return true
	}

	return false
}

// SetCategoryTypes gets a reference to the given []CategoryType and assigns it to the CategoryTypes field.
func (o *PaymentPolicy) SetCategoryTypes(v []CategoryType) {
	o.CategoryTypes = &v
}

// GetDeposit returns the Deposit field value if set, zero value otherwise.
func (o *PaymentPolicy) GetDeposit() Deposit {
	if o == nil || o.Deposit == nil {
		var ret Deposit
		return ret
	}
	return *o.Deposit
}

// GetDepositOk returns a tuple with the Deposit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentPolicy) GetDepositOk() (*Deposit, bool) {
	if o == nil || o.Deposit == nil {
		return nil, false
	}
	return o.Deposit, true
}

// HasDeposit returns a boolean if a field has been set.
func (o *PaymentPolicy) HasDeposit() bool {
	if o != nil && o.Deposit != nil {
		return true
	}

	return false
}

// SetDeposit gets a reference to the given Deposit and assigns it to the Deposit field.
func (o *PaymentPolicy) SetDeposit(v Deposit) {
	o.Deposit = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PaymentPolicy) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentPolicy) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PaymentPolicy) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PaymentPolicy) SetDescription(v string) {
	o.Description = &v
}

// GetFullPaymentDueIn returns the FullPaymentDueIn field value if set, zero value otherwise.
func (o *PaymentPolicy) GetFullPaymentDueIn() TimeDuration {
	if o == nil || o.FullPaymentDueIn == nil {
		var ret TimeDuration
		return ret
	}
	return *o.FullPaymentDueIn
}

// GetFullPaymentDueInOk returns a tuple with the FullPaymentDueIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentPolicy) GetFullPaymentDueInOk() (*TimeDuration, bool) {
	if o == nil || o.FullPaymentDueIn == nil {
		return nil, false
	}
	return o.FullPaymentDueIn, true
}

// HasFullPaymentDueIn returns a boolean if a field has been set.
func (o *PaymentPolicy) HasFullPaymentDueIn() bool {
	if o != nil && o.FullPaymentDueIn != nil {
		return true
	}

	return false
}

// SetFullPaymentDueIn gets a reference to the given TimeDuration and assigns it to the FullPaymentDueIn field.
func (o *PaymentPolicy) SetFullPaymentDueIn(v TimeDuration) {
	o.FullPaymentDueIn = &v
}

// GetImmediatePay returns the ImmediatePay field value if set, zero value otherwise.
func (o *PaymentPolicy) GetImmediatePay() bool {
	if o == nil || o.ImmediatePay == nil {
		var ret bool
		return ret
	}
	return *o.ImmediatePay
}

// GetImmediatePayOk returns a tuple with the ImmediatePay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentPolicy) GetImmediatePayOk() (*bool, bool) {
	if o == nil || o.ImmediatePay == nil {
		return nil, false
	}
	return o.ImmediatePay, true
}

// HasImmediatePay returns a boolean if a field has been set.
func (o *PaymentPolicy) HasImmediatePay() bool {
	if o != nil && o.ImmediatePay != nil {
		return true
	}

	return false
}

// SetImmediatePay gets a reference to the given bool and assigns it to the ImmediatePay field.
func (o *PaymentPolicy) SetImmediatePay(v bool) {
	o.ImmediatePay = &v
}

// GetMarketplaceId returns the MarketplaceId field value if set, zero value otherwise.
func (o *PaymentPolicy) GetMarketplaceId() string {
	if o == nil || o.MarketplaceId == nil {
		var ret string
		return ret
	}
	return *o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentPolicy) GetMarketplaceIdOk() (*string, bool) {
	if o == nil || o.MarketplaceId == nil {
		return nil, false
	}
	return o.MarketplaceId, true
}

// HasMarketplaceId returns a boolean if a field has been set.
func (o *PaymentPolicy) HasMarketplaceId() bool {
	if o != nil && o.MarketplaceId != nil {
		return true
	}

	return false
}

// SetMarketplaceId gets a reference to the given string and assigns it to the MarketplaceId field.
func (o *PaymentPolicy) SetMarketplaceId(v string) {
	o.MarketplaceId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PaymentPolicy) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentPolicy) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PaymentPolicy) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PaymentPolicy) SetName(v string) {
	o.Name = &v
}

// GetPaymentInstructions returns the PaymentInstructions field value if set, zero value otherwise.
func (o *PaymentPolicy) GetPaymentInstructions() string {
	if o == nil || o.PaymentInstructions == nil {
		var ret string
		return ret
	}
	return *o.PaymentInstructions
}

// GetPaymentInstructionsOk returns a tuple with the PaymentInstructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentPolicy) GetPaymentInstructionsOk() (*string, bool) {
	if o == nil || o.PaymentInstructions == nil {
		return nil, false
	}
	return o.PaymentInstructions, true
}

// HasPaymentInstructions returns a boolean if a field has been set.
func (o *PaymentPolicy) HasPaymentInstructions() bool {
	if o != nil && o.PaymentInstructions != nil {
		return true
	}

	return false
}

// SetPaymentInstructions gets a reference to the given string and assigns it to the PaymentInstructions field.
func (o *PaymentPolicy) SetPaymentInstructions(v string) {
	o.PaymentInstructions = &v
}

// GetPaymentMethods returns the PaymentMethods field value if set, zero value otherwise.
func (o *PaymentPolicy) GetPaymentMethods() []PaymentMethod {
	if o == nil || o.PaymentMethods == nil {
		var ret []PaymentMethod
		return ret
	}
	return *o.PaymentMethods
}

// GetPaymentMethodsOk returns a tuple with the PaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentPolicy) GetPaymentMethodsOk() (*[]PaymentMethod, bool) {
	if o == nil || o.PaymentMethods == nil {
		return nil, false
	}
	return o.PaymentMethods, true
}

// HasPaymentMethods returns a boolean if a field has been set.
func (o *PaymentPolicy) HasPaymentMethods() bool {
	if o != nil && o.PaymentMethods != nil {
		return true
	}

	return false
}

// SetPaymentMethods gets a reference to the given []PaymentMethod and assigns it to the PaymentMethods field.
func (o *PaymentPolicy) SetPaymentMethods(v []PaymentMethod) {
	o.PaymentMethods = &v
}

// GetPaymentPolicyId returns the PaymentPolicyId field value if set, zero value otherwise.
func (o *PaymentPolicy) GetPaymentPolicyId() string {
	if o == nil || o.PaymentPolicyId == nil {
		var ret string
		return ret
	}
	return *o.PaymentPolicyId
}

// GetPaymentPolicyIdOk returns a tuple with the PaymentPolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentPolicy) GetPaymentPolicyIdOk() (*string, bool) {
	if o == nil || o.PaymentPolicyId == nil {
		return nil, false
	}
	return o.PaymentPolicyId, true
}

// HasPaymentPolicyId returns a boolean if a field has been set.
func (o *PaymentPolicy) HasPaymentPolicyId() bool {
	if o != nil && o.PaymentPolicyId != nil {
		return true
	}

	return false
}

// SetPaymentPolicyId gets a reference to the given string and assigns it to the PaymentPolicyId field.
func (o *PaymentPolicy) SetPaymentPolicyId(v string) {
	o.PaymentPolicyId = &v
}

func (o PaymentPolicy) MarshalJSON() ([]byte, error) {
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
	if o.PaymentPolicyId != nil {
		toSerialize["paymentPolicyId"] = o.PaymentPolicyId
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentPolicy struct {
	value *PaymentPolicy
	isSet bool
}

func (v NullablePaymentPolicy) Get() *PaymentPolicy {
	return v.value
}

func (v *NullablePaymentPolicy) Set(val *PaymentPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentPolicy(val *PaymentPolicy) *NullablePaymentPolicy {
	return &NullablePaymentPolicy{value: val, isSet: true}
}

func (v NullablePaymentPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

