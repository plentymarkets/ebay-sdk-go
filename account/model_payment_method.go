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

// PaymentMethod Container specifying a payment method that is accepted by the seller. Specify multiple payment methods by repeating this container. For more on payment methods, see <a href=\"http://pages.ebay.com/help/policies/accepted-payments-policy.html\">Accepted payments policy</a>.  <p>Note that payment methods are not applicable to classified ad listings &ndash; all classified ad payments are handled off of the eBay platform.</p>
type PaymentMethod struct {
	// It's important to note that the credit card brands Visa and MasterCard must both be listed if either one is listed, as is shown in the following code fragment:  <br><br><code>\"paymentMethods\": [{ \"brands\": [VISA, MASTERCARD] }] ...</code> <p class=\"tablenote\"><b>Note:</b> Different eBay marketplaces may or may not support this field. <br><br>Use the Trading API <b>GetCategoryFeatures</b> call with <b>FeatureID</b> set to <code>PaymentMethods</code> and <b>DetailLevel</b> set to <code>ReturnAll</code> to see what credit card brands different marketplaces support. If the <b>GetCategoryFeatures</b> call returns details on credit card brands for the categories in which you sell, then you can use this field to list the credit card brands the seller accepts. If, on the other hand, <b>GetCategoryFeatures</b> does not enumerate credit card brands for your target site (for example, if it returns <b>PaymentMethod</b> set to <code>CCAccepted</code>), then you cannot enumerate specific credit card brands with this field for that marketplace.</p> <p><i>Required if </i> <b>paymentMethodType</b> is set to <code>CREDIT_CARD</code>.  <br><br>A list of credit card brands accepted by the seller.</p>
	Brands *[]string `json:"brands,omitempty"`
	// The payment method, selected from the supported payment method types.  <p>Use <b>GetCategoryFeatures</b> in the Trading API to retrieve the payment methods allowed for a category on a specific marketplace, as well as the default payment method for that marketplace (review the <b>SiteDefaults.PaymentMethod</b> field). For example, the response from <b>GetCategoryFeatures</b> shows that on the US marketplace, most categories allow only electronic payments via credit cards, PayPal, and the like. Also, note that <b>GeteBayDetails</b> does not return payment method information.  <p class=\"tablenote\"><b>Note:</b> If you create item listings using the Inventory API, you must set this field to <code>PAYPAL</code> (currently, the Inventory API supports only fixed-prince GTC items where the only supported <b>paymentMethod</b> is PayPal).</p> For implementation help, refer to <a href='https://developer.ebay.com/api-docs/sell/account/types/api:PaymentMethodTypeEnum'>eBay API documentation</a>
	PaymentMethodType *string `json:"paymentMethodType,omitempty"`
	RecipientAccountReference *RecipientAccountReference `json:"recipientAccountReference,omitempty"`
}

// NewPaymentMethod instantiates a new PaymentMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethod() *PaymentMethod {
	this := PaymentMethod{}
	return &this
}

// NewPaymentMethodWithDefaults instantiates a new PaymentMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodWithDefaults() *PaymentMethod {
	this := PaymentMethod{}
	return &this
}

// GetBrands returns the Brands field value if set, zero value otherwise.
func (o *PaymentMethod) GetBrands() []string {
	if o == nil || o.Brands == nil {
		var ret []string
		return ret
	}
	return *o.Brands
}

// GetBrandsOk returns a tuple with the Brands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetBrandsOk() (*[]string, bool) {
	if o == nil || o.Brands == nil {
		return nil, false
	}
	return o.Brands, true
}

// HasBrands returns a boolean if a field has been set.
func (o *PaymentMethod) HasBrands() bool {
	if o != nil && o.Brands != nil {
		return true
	}

	return false
}

// SetBrands gets a reference to the given []string and assigns it to the Brands field.
func (o *PaymentMethod) SetBrands(v []string) {
	o.Brands = &v
}

// GetPaymentMethodType returns the PaymentMethodType field value if set, zero value otherwise.
func (o *PaymentMethod) GetPaymentMethodType() string {
	if o == nil || o.PaymentMethodType == nil {
		var ret string
		return ret
	}
	return *o.PaymentMethodType
}

// GetPaymentMethodTypeOk returns a tuple with the PaymentMethodType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetPaymentMethodTypeOk() (*string, bool) {
	if o == nil || o.PaymentMethodType == nil {
		return nil, false
	}
	return o.PaymentMethodType, true
}

// HasPaymentMethodType returns a boolean if a field has been set.
func (o *PaymentMethod) HasPaymentMethodType() bool {
	if o != nil && o.PaymentMethodType != nil {
		return true
	}

	return false
}

// SetPaymentMethodType gets a reference to the given string and assigns it to the PaymentMethodType field.
func (o *PaymentMethod) SetPaymentMethodType(v string) {
	o.PaymentMethodType = &v
}

// GetRecipientAccountReference returns the RecipientAccountReference field value if set, zero value otherwise.
func (o *PaymentMethod) GetRecipientAccountReference() RecipientAccountReference {
	if o == nil || o.RecipientAccountReference == nil {
		var ret RecipientAccountReference
		return ret
	}
	return *o.RecipientAccountReference
}

// GetRecipientAccountReferenceOk returns a tuple with the RecipientAccountReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetRecipientAccountReferenceOk() (*RecipientAccountReference, bool) {
	if o == nil || o.RecipientAccountReference == nil {
		return nil, false
	}
	return o.RecipientAccountReference, true
}

// HasRecipientAccountReference returns a boolean if a field has been set.
func (o *PaymentMethod) HasRecipientAccountReference() bool {
	if o != nil && o.RecipientAccountReference != nil {
		return true
	}

	return false
}

// SetRecipientAccountReference gets a reference to the given RecipientAccountReference and assigns it to the RecipientAccountReference field.
func (o *PaymentMethod) SetRecipientAccountReference(v RecipientAccountReference) {
	o.RecipientAccountReference = &v
}

func (o PaymentMethod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Brands != nil {
		toSerialize["brands"] = o.Brands
	}
	if o.PaymentMethodType != nil {
		toSerialize["paymentMethodType"] = o.PaymentMethodType
	}
	if o.RecipientAccountReference != nil {
		toSerialize["recipientAccountReference"] = o.RecipientAccountReference
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentMethod struct {
	value *PaymentMethod
	isSet bool
}

func (v NullablePaymentMethod) Get() *PaymentMethod {
	return v.value
}

func (v *NullablePaymentMethod) Set(val *PaymentMethod) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethod) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethod(val *PaymentMethod) *NullablePaymentMethod {
	return &NullablePaymentMethod{value: val, isSet: true}
}

func (v NullablePaymentMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


