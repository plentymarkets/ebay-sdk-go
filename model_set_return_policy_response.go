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

// SetReturnPolicyResponse Complex type that that gets populated with a response containing a return policy.
type SetReturnPolicyResponse struct {
	// For return policies, this field always returns <code>ALL_EXCLUDING_MOTORS_VEHICLES</code> (returns on motor vehicles are not processed through eBay flows.) <br><br><b>Default</b>: <code>ALL_EXCLUDING_MOTORS_VEHICLES</code> (for return policies only)
	CategoryTypes *[]CategoryType `json:"categoryTypes,omitempty"`
	// An optional seller-defined description of the return policy for internal use (this value is not displayed to end users).
	Description *string `json:"description,omitempty"`
	// <p class=\"tablenote\"><span  style=\"color: #dd1e31;\"><b>Important!</b></span> This field has been deprecated as of version 1.2.0, released on May 31, 2018. Any value supplied in this field is ignored, it is neither read nor returned.</p>  <p>If set to <code>true</code>, the seller offers an <em>Extended Holiday Returns</em> policy for their listings. <p><span class=\"tablenote\"><strong>IMPORTANT:</strong> Extended Holiday Returns is a seasonally available feature that is offered on some eBay marketplaces. To see if the feature is enabled in any given year, check the eBay Seller Center <a href=\"http://pages.ebay.com/sellerinformation/returns-on-eBay/\">Returns on eBay</a> page of before the holiday season begins.</p>
	ExtendedHolidayReturnsOffered *bool `json:"extendedHolidayReturnsOffered,omitempty"`
	InternationalOverride *InternationalReturnOverrideType `json:"internationalOverride,omitempty"`
	// The ID of the eBay marketplace to which this return policy applies. If this value is not specified, value defaults to the seller's eBay registration site. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/sell/account/types/ba:MarketplaceIdEnum'>eBay API documentation</a>
	MarketplaceId *string `json:"marketplaceId,omitempty"`
	// A user-defined name for this return policy. Names must be unique for policies assigned to the same marketplace. <br><br><b>Max length</b>: 64
	Name *string `json:"name,omitempty"`
	// <p class=\"tablenote\"><span  style=\"color: #dd1e31;\"><b>Important!</b></span> This field has been deprecated as of version 1.2.0, released on May 31, 2018. Any value other than <code>MONEY_BACK</code> will be treated as <code>MONEY_BACK</code> (although for a period of time, eBay will store and return the legacy values to preserve backwards compatibility).</p>  Indicates the method the seller uses to compensate the buyer for returned items. The return method specified applies only to <a href=\"http://developer.ebay.com/DevZone/guides/features-guide/default.html#Development/Post-Order-Returns.html#return-reasons\" target=\"_blank\">remorse returns</a>. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/sell/account/types/api:RefundMethodEnum'>eBay API documentation</a>
	RefundMethod *string `json:"refundMethod,omitempty"`
	// <p class=\"tablenote\"><span  style=\"color: #dd1e31;\"><b>Important!</b></span> This field has been deprecated as of version 1.2.0, released on May 31, 2018. Any value supplied in this field is ignored, it is neither read nor returned.</p>  <p>Optionally set by the seller, the percentage charged if the seller charges buyers a a restocking fee when items are returned due to buyer remorse and/or a purchasing mistake. The total amount charged to the buyer is the cost of the item multiplied by the percentage indicated in this field.
	RestockingFeePercentage *string `json:"restockingFeePercentage,omitempty"`
	// This field contains the seller's detailed explanation for their return policy and is displayed in the Return Policy section of the View Item page. This field is valid in only the following marketplaces (the field is otherwise ignored): <ul> <li>Germany (DE)</li> <li>Spain (ES)</li> <li>France (FR)</li> <li>Italy (IT)</li> </ul>
	ReturnInstructions *string `json:"returnInstructions,omitempty"`
	// This field indicates the method in which the seller handles non-money back return requests for <a href=\"http://developer.ebay.com/DevZone/guides/features-guide/default.html#Development/Post-Order-Returns.html#return-reasons\" target=\"_blank\">remorse returns</a>. This field is valid in only the US marketplace and the only valid value is <code>REPLACEMENT</code>. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/sell/account/types/api:ReturnMethodEnum'>eBay API documentation</a>
	ReturnMethod *string `json:"returnMethod,omitempty"`
	ReturnPeriod *TimeDuration `json:"returnPeriod,omitempty"`
	// A unique eBay-assigned ID for a return policy. This ID is generated when the policy is created.
	ReturnPolicyId *string `json:"returnPolicyId,omitempty"`
	// If set to <code>true</code>, the seller accepts returns. If set to <code>false</code>, this field indicates that the seller does not accept returns.
	ReturnsAccepted *bool `json:"returnsAccepted,omitempty"`
	// This field indicates who is responsible for paying for the shipping charges for returned items. The field can be set to either <code>BUYER</code> or <code>SELLER</code>.  <br><br>Depending on the return policy and specifics of the return, either the buyer or the seller can be responsible for the return shipping costs. Note that the seller is always responsible for return shipping costs for SNAD-related issues. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/sell/account/types/api:ReturnShippingCostPayerEnum'>eBay API documentation</a>
	ReturnShippingCostPayer *string `json:"returnShippingCostPayer,omitempty"`
	// A list of warnings related to request. This field normally returns empty, which indicates the request did not generate any warnings.
	Warnings *[]Error `json:"warnings,omitempty"`
}

// NewSetReturnPolicyResponse instantiates a new SetReturnPolicyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetReturnPolicyResponse() *SetReturnPolicyResponse {
	this := SetReturnPolicyResponse{}
	return &this
}

// NewSetReturnPolicyResponseWithDefaults instantiates a new SetReturnPolicyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetReturnPolicyResponseWithDefaults() *SetReturnPolicyResponse {
	this := SetReturnPolicyResponse{}
	return &this
}

// GetCategoryTypes returns the CategoryTypes field value if set, zero value otherwise.
func (o *SetReturnPolicyResponse) GetCategoryTypes() []CategoryType {
	if o == nil || o.CategoryTypes == nil {
		var ret []CategoryType
		return ret
	}
	return *o.CategoryTypes
}

// GetCategoryTypesOk returns a tuple with the CategoryTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetReturnPolicyResponse) GetCategoryTypesOk() (*[]CategoryType, bool) {
	if o == nil || o.CategoryTypes == nil {
		return nil, false
	}
	return o.CategoryTypes, true
}

// HasCategoryTypes returns a boolean if a field has been set.
func (o *SetReturnPolicyResponse) HasCategoryTypes() bool {
	if o != nil && o.CategoryTypes != nil {
		return true
	}

	return false
}

// SetCategoryTypes gets a reference to the given []CategoryType and assigns it to the CategoryTypes field.
func (o *SetReturnPolicyResponse) SetCategoryTypes(v []CategoryType) {
	o.CategoryTypes = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SetReturnPolicyResponse) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetReturnPolicyResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SetReturnPolicyResponse) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SetReturnPolicyResponse) SetDescription(v string) {
	o.Description = &v
}

// GetExtendedHolidayReturnsOffered returns the ExtendedHolidayReturnsOffered field value if set, zero value otherwise.
func (o *SetReturnPolicyResponse) GetExtendedHolidayReturnsOffered() bool {
	if o == nil || o.ExtendedHolidayReturnsOffered == nil {
		var ret bool
		return ret
	}
	return *o.ExtendedHolidayReturnsOffered
}

// GetExtendedHolidayReturnsOfferedOk returns a tuple with the ExtendedHolidayReturnsOffered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetReturnPolicyResponse) GetExtendedHolidayReturnsOfferedOk() (*bool, bool) {
	if o == nil || o.ExtendedHolidayReturnsOffered == nil {
		return nil, false
	}
	return o.ExtendedHolidayReturnsOffered, true
}

// HasExtendedHolidayReturnsOffered returns a boolean if a field has been set.
func (o *SetReturnPolicyResponse) HasExtendedHolidayReturnsOffered() bool {
	if o != nil && o.ExtendedHolidayReturnsOffered != nil {
		return true
	}

	return false
}

// SetExtendedHolidayReturnsOffered gets a reference to the given bool and assigns it to the ExtendedHolidayReturnsOffered field.
func (o *SetReturnPolicyResponse) SetExtendedHolidayReturnsOffered(v bool) {
	o.ExtendedHolidayReturnsOffered = &v
}

// GetInternationalOverride returns the InternationalOverride field value if set, zero value otherwise.
func (o *SetReturnPolicyResponse) GetInternationalOverride() InternationalReturnOverrideType {
	if o == nil || o.InternationalOverride == nil {
		var ret InternationalReturnOverrideType
		return ret
	}
	return *o.InternationalOverride
}

// GetInternationalOverrideOk returns a tuple with the InternationalOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetReturnPolicyResponse) GetInternationalOverrideOk() (*InternationalReturnOverrideType, bool) {
	if o == nil || o.InternationalOverride == nil {
		return nil, false
	}
	return o.InternationalOverride, true
}

// HasInternationalOverride returns a boolean if a field has been set.
func (o *SetReturnPolicyResponse) HasInternationalOverride() bool {
	if o != nil && o.InternationalOverride != nil {
		return true
	}

	return false
}

// SetInternationalOverride gets a reference to the given InternationalReturnOverrideType and assigns it to the InternationalOverride field.
func (o *SetReturnPolicyResponse) SetInternationalOverride(v InternationalReturnOverrideType) {
	o.InternationalOverride = &v
}

// GetMarketplaceId returns the MarketplaceId field value if set, zero value otherwise.
func (o *SetReturnPolicyResponse) GetMarketplaceId() string {
	if o == nil || o.MarketplaceId == nil {
		var ret string
		return ret
	}
	return *o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetReturnPolicyResponse) GetMarketplaceIdOk() (*string, bool) {
	if o == nil || o.MarketplaceId == nil {
		return nil, false
	}
	return o.MarketplaceId, true
}

// HasMarketplaceId returns a boolean if a field has been set.
func (o *SetReturnPolicyResponse) HasMarketplaceId() bool {
	if o != nil && o.MarketplaceId != nil {
		return true
	}

	return false
}

// SetMarketplaceId gets a reference to the given string and assigns it to the MarketplaceId field.
func (o *SetReturnPolicyResponse) SetMarketplaceId(v string) {
	o.MarketplaceId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SetReturnPolicyResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetReturnPolicyResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SetReturnPolicyResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SetReturnPolicyResponse) SetName(v string) {
	o.Name = &v
}

// GetRefundMethod returns the RefundMethod field value if set, zero value otherwise.
func (o *SetReturnPolicyResponse) GetRefundMethod() string {
	if o == nil || o.RefundMethod == nil {
		var ret string
		return ret
	}
	return *o.RefundMethod
}

// GetRefundMethodOk returns a tuple with the RefundMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetReturnPolicyResponse) GetRefundMethodOk() (*string, bool) {
	if o == nil || o.RefundMethod == nil {
		return nil, false
	}
	return o.RefundMethod, true
}

// HasRefundMethod returns a boolean if a field has been set.
func (o *SetReturnPolicyResponse) HasRefundMethod() bool {
	if o != nil && o.RefundMethod != nil {
		return true
	}

	return false
}

// SetRefundMethod gets a reference to the given string and assigns it to the RefundMethod field.
func (o *SetReturnPolicyResponse) SetRefundMethod(v string) {
	o.RefundMethod = &v
}

// GetRestockingFeePercentage returns the RestockingFeePercentage field value if set, zero value otherwise.
func (o *SetReturnPolicyResponse) GetRestockingFeePercentage() string {
	if o == nil || o.RestockingFeePercentage == nil {
		var ret string
		return ret
	}
	return *o.RestockingFeePercentage
}

// GetRestockingFeePercentageOk returns a tuple with the RestockingFeePercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetReturnPolicyResponse) GetRestockingFeePercentageOk() (*string, bool) {
	if o == nil || o.RestockingFeePercentage == nil {
		return nil, false
	}
	return o.RestockingFeePercentage, true
}

// HasRestockingFeePercentage returns a boolean if a field has been set.
func (o *SetReturnPolicyResponse) HasRestockingFeePercentage() bool {
	if o != nil && o.RestockingFeePercentage != nil {
		return true
	}

	return false
}

// SetRestockingFeePercentage gets a reference to the given string and assigns it to the RestockingFeePercentage field.
func (o *SetReturnPolicyResponse) SetRestockingFeePercentage(v string) {
	o.RestockingFeePercentage = &v
}

// GetReturnInstructions returns the ReturnInstructions field value if set, zero value otherwise.
func (o *SetReturnPolicyResponse) GetReturnInstructions() string {
	if o == nil || o.ReturnInstructions == nil {
		var ret string
		return ret
	}
	return *o.ReturnInstructions
}

// GetReturnInstructionsOk returns a tuple with the ReturnInstructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetReturnPolicyResponse) GetReturnInstructionsOk() (*string, bool) {
	if o == nil || o.ReturnInstructions == nil {
		return nil, false
	}
	return o.ReturnInstructions, true
}

// HasReturnInstructions returns a boolean if a field has been set.
func (o *SetReturnPolicyResponse) HasReturnInstructions() bool {
	if o != nil && o.ReturnInstructions != nil {
		return true
	}

	return false
}

// SetReturnInstructions gets a reference to the given string and assigns it to the ReturnInstructions field.
func (o *SetReturnPolicyResponse) SetReturnInstructions(v string) {
	o.ReturnInstructions = &v
}

// GetReturnMethod returns the ReturnMethod field value if set, zero value otherwise.
func (o *SetReturnPolicyResponse) GetReturnMethod() string {
	if o == nil || o.ReturnMethod == nil {
		var ret string
		return ret
	}
	return *o.ReturnMethod
}

// GetReturnMethodOk returns a tuple with the ReturnMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetReturnPolicyResponse) GetReturnMethodOk() (*string, bool) {
	if o == nil || o.ReturnMethod == nil {
		return nil, false
	}
	return o.ReturnMethod, true
}

// HasReturnMethod returns a boolean if a field has been set.
func (o *SetReturnPolicyResponse) HasReturnMethod() bool {
	if o != nil && o.ReturnMethod != nil {
		return true
	}

	return false
}

// SetReturnMethod gets a reference to the given string and assigns it to the ReturnMethod field.
func (o *SetReturnPolicyResponse) SetReturnMethod(v string) {
	o.ReturnMethod = &v
}

// GetReturnPeriod returns the ReturnPeriod field value if set, zero value otherwise.
func (o *SetReturnPolicyResponse) GetReturnPeriod() TimeDuration {
	if o == nil || o.ReturnPeriod == nil {
		var ret TimeDuration
		return ret
	}
	return *o.ReturnPeriod
}

// GetReturnPeriodOk returns a tuple with the ReturnPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetReturnPolicyResponse) GetReturnPeriodOk() (*TimeDuration, bool) {
	if o == nil || o.ReturnPeriod == nil {
		return nil, false
	}
	return o.ReturnPeriod, true
}

// HasReturnPeriod returns a boolean if a field has been set.
func (o *SetReturnPolicyResponse) HasReturnPeriod() bool {
	if o != nil && o.ReturnPeriod != nil {
		return true
	}

	return false
}

// SetReturnPeriod gets a reference to the given TimeDuration and assigns it to the ReturnPeriod field.
func (o *SetReturnPolicyResponse) SetReturnPeriod(v TimeDuration) {
	o.ReturnPeriod = &v
}

// GetReturnPolicyId returns the ReturnPolicyId field value if set, zero value otherwise.
func (o *SetReturnPolicyResponse) GetReturnPolicyId() string {
	if o == nil || o.ReturnPolicyId == nil {
		var ret string
		return ret
	}
	return *o.ReturnPolicyId
}

// GetReturnPolicyIdOk returns a tuple with the ReturnPolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetReturnPolicyResponse) GetReturnPolicyIdOk() (*string, bool) {
	if o == nil || o.ReturnPolicyId == nil {
		return nil, false
	}
	return o.ReturnPolicyId, true
}

// HasReturnPolicyId returns a boolean if a field has been set.
func (o *SetReturnPolicyResponse) HasReturnPolicyId() bool {
	if o != nil && o.ReturnPolicyId != nil {
		return true
	}

	return false
}

// SetReturnPolicyId gets a reference to the given string and assigns it to the ReturnPolicyId field.
func (o *SetReturnPolicyResponse) SetReturnPolicyId(v string) {
	o.ReturnPolicyId = &v
}

// GetReturnsAccepted returns the ReturnsAccepted field value if set, zero value otherwise.
func (o *SetReturnPolicyResponse) GetReturnsAccepted() bool {
	if o == nil || o.ReturnsAccepted == nil {
		var ret bool
		return ret
	}
	return *o.ReturnsAccepted
}

// GetReturnsAcceptedOk returns a tuple with the ReturnsAccepted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetReturnPolicyResponse) GetReturnsAcceptedOk() (*bool, bool) {
	if o == nil || o.ReturnsAccepted == nil {
		return nil, false
	}
	return o.ReturnsAccepted, true
}

// HasReturnsAccepted returns a boolean if a field has been set.
func (o *SetReturnPolicyResponse) HasReturnsAccepted() bool {
	if o != nil && o.ReturnsAccepted != nil {
		return true
	}

	return false
}

// SetReturnsAccepted gets a reference to the given bool and assigns it to the ReturnsAccepted field.
func (o *SetReturnPolicyResponse) SetReturnsAccepted(v bool) {
	o.ReturnsAccepted = &v
}

// GetReturnShippingCostPayer returns the ReturnShippingCostPayer field value if set, zero value otherwise.
func (o *SetReturnPolicyResponse) GetReturnShippingCostPayer() string {
	if o == nil || o.ReturnShippingCostPayer == nil {
		var ret string
		return ret
	}
	return *o.ReturnShippingCostPayer
}

// GetReturnShippingCostPayerOk returns a tuple with the ReturnShippingCostPayer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetReturnPolicyResponse) GetReturnShippingCostPayerOk() (*string, bool) {
	if o == nil || o.ReturnShippingCostPayer == nil {
		return nil, false
	}
	return o.ReturnShippingCostPayer, true
}

// HasReturnShippingCostPayer returns a boolean if a field has been set.
func (o *SetReturnPolicyResponse) HasReturnShippingCostPayer() bool {
	if o != nil && o.ReturnShippingCostPayer != nil {
		return true
	}

	return false
}

// SetReturnShippingCostPayer gets a reference to the given string and assigns it to the ReturnShippingCostPayer field.
func (o *SetReturnPolicyResponse) SetReturnShippingCostPayer(v string) {
	o.ReturnShippingCostPayer = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *SetReturnPolicyResponse) GetWarnings() []Error {
	if o == nil || o.Warnings == nil {
		var ret []Error
		return ret
	}
	return *o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetReturnPolicyResponse) GetWarningsOk() (*[]Error, bool) {
	if o == nil || o.Warnings == nil {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *SetReturnPolicyResponse) HasWarnings() bool {
	if o != nil && o.Warnings != nil {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []Error and assigns it to the Warnings field.
func (o *SetReturnPolicyResponse) SetWarnings(v []Error) {
	o.Warnings = &v
}

func (o SetReturnPolicyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CategoryTypes != nil {
		toSerialize["categoryTypes"] = o.CategoryTypes
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ExtendedHolidayReturnsOffered != nil {
		toSerialize["extendedHolidayReturnsOffered"] = o.ExtendedHolidayReturnsOffered
	}
	if o.InternationalOverride != nil {
		toSerialize["internationalOverride"] = o.InternationalOverride
	}
	if o.MarketplaceId != nil {
		toSerialize["marketplaceId"] = o.MarketplaceId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.RefundMethod != nil {
		toSerialize["refundMethod"] = o.RefundMethod
	}
	if o.RestockingFeePercentage != nil {
		toSerialize["restockingFeePercentage"] = o.RestockingFeePercentage
	}
	if o.ReturnInstructions != nil {
		toSerialize["returnInstructions"] = o.ReturnInstructions
	}
	if o.ReturnMethod != nil {
		toSerialize["returnMethod"] = o.ReturnMethod
	}
	if o.ReturnPeriod != nil {
		toSerialize["returnPeriod"] = o.ReturnPeriod
	}
	if o.ReturnPolicyId != nil {
		toSerialize["returnPolicyId"] = o.ReturnPolicyId
	}
	if o.ReturnsAccepted != nil {
		toSerialize["returnsAccepted"] = o.ReturnsAccepted
	}
	if o.ReturnShippingCostPayer != nil {
		toSerialize["returnShippingCostPayer"] = o.ReturnShippingCostPayer
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}
	return json.Marshal(toSerialize)
}

type NullableSetReturnPolicyResponse struct {
	value *SetReturnPolicyResponse
	isSet bool
}

func (v NullableSetReturnPolicyResponse) Get() *SetReturnPolicyResponse {
	return v.value
}

func (v *NullableSetReturnPolicyResponse) Set(val *SetReturnPolicyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSetReturnPolicyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSetReturnPolicyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetReturnPolicyResponse(val *SetReturnPolicyResponse) *NullableSetReturnPolicyResponse {
	return &NullableSetReturnPolicyResponse{value: val, isSet: true}
}

func (v NullableSetReturnPolicyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetReturnPolicyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

