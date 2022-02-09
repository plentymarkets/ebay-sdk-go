/*
Feed API

<p>The <strong>Feed API</strong> lets sellers upload input files, download reports and files including their status, filter reports using URI parameters, and retrieve customer service metrics task details.</p>

API version: v1.3.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package feed

import (
	"encoding/json"
)

// CustomerServiceMetricsFilterCriteria A complex data type that filters data for report creation. See CustomerServiceMetricsFilterCriteria for fields and descriptions.
type CustomerServiceMetricsFilterCriteria struct {
	// An enumeration value that specifies the customer service metric that eBay tracks to measure seller performance. See CustomerServiceMetricTypeEnum for values. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/sell/feed/types/api:CustomerServiceMetricTypeEnum'>eBay API documentation</a>
	CustomerServiceMetricType *string `json:"customerServiceMetricType,omitempty"`
	// An enumeration value that specifies the eBay marketplace where the evaluation occurs. See MarketplaceIdEnum for values. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/sell/feed/types/bas:MarketplaceIdEnum'>eBay API documentation</a>
	EvaluationMarketplaceId *string `json:"evaluationMarketplaceId,omitempty"`
	// A list of listing category IDs on which the service metric is measured. A seller can use one or more L1 (top-level) eBay categories to get metrics specific to those L1 categories. The Category IDs for each L1 category are required. Category ID values for L1 categories can be retrieved using the Taxonomy API. Note: Pass this attribute to narrow down your filter results for the ITEM_NOT_AS_DESCRIBED customerServiceMetricType. Supported categories include: primary(L1) category Id
	ListingCategories *[]string `json:"listingCategories,omitempty"`
	// A list of shipping region enumeration values on which the service metric is measured. This comma delimited array allows the seller to customize the report to focus on domestic or international shipping. Note: Pass this attribute to narrow down your filter results for the ITEM_NOT_RECEIVED customerServiceMetricType. Supported categories include: primary(L1) category IdSee ShippingRegionTypeEnum for values
	ShippingRegions *[]string `json:"shippingRegions,omitempty"`
}

// NewCustomerServiceMetricsFilterCriteria instantiates a new CustomerServiceMetricsFilterCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerServiceMetricsFilterCriteria() *CustomerServiceMetricsFilterCriteria {
	this := CustomerServiceMetricsFilterCriteria{}
	return &this
}

// NewCustomerServiceMetricsFilterCriteriaWithDefaults instantiates a new CustomerServiceMetricsFilterCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerServiceMetricsFilterCriteriaWithDefaults() *CustomerServiceMetricsFilterCriteria {
	this := CustomerServiceMetricsFilterCriteria{}
	return &this
}

// GetCustomerServiceMetricType returns the CustomerServiceMetricType field value if set, zero value otherwise.
func (o *CustomerServiceMetricsFilterCriteria) GetCustomerServiceMetricType() string {
	if o == nil || o.CustomerServiceMetricType == nil {
		var ret string
		return ret
	}
	return *o.CustomerServiceMetricType
}

// GetCustomerServiceMetricTypeOk returns a tuple with the CustomerServiceMetricType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerServiceMetricsFilterCriteria) GetCustomerServiceMetricTypeOk() (*string, bool) {
	if o == nil || o.CustomerServiceMetricType == nil {
		return nil, false
	}
	return o.CustomerServiceMetricType, true
}

// HasCustomerServiceMetricType returns a boolean if a field has been set.
func (o *CustomerServiceMetricsFilterCriteria) HasCustomerServiceMetricType() bool {
	if o != nil && o.CustomerServiceMetricType != nil {
		return true
	}

	return false
}

// SetCustomerServiceMetricType gets a reference to the given string and assigns it to the CustomerServiceMetricType field.
func (o *CustomerServiceMetricsFilterCriteria) SetCustomerServiceMetricType(v string) {
	o.CustomerServiceMetricType = &v
}

// GetEvaluationMarketplaceId returns the EvaluationMarketplaceId field value if set, zero value otherwise.
func (o *CustomerServiceMetricsFilterCriteria) GetEvaluationMarketplaceId() string {
	if o == nil || o.EvaluationMarketplaceId == nil {
		var ret string
		return ret
	}
	return *o.EvaluationMarketplaceId
}

// GetEvaluationMarketplaceIdOk returns a tuple with the EvaluationMarketplaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerServiceMetricsFilterCriteria) GetEvaluationMarketplaceIdOk() (*string, bool) {
	if o == nil || o.EvaluationMarketplaceId == nil {
		return nil, false
	}
	return o.EvaluationMarketplaceId, true
}

// HasEvaluationMarketplaceId returns a boolean if a field has been set.
func (o *CustomerServiceMetricsFilterCriteria) HasEvaluationMarketplaceId() bool {
	if o != nil && o.EvaluationMarketplaceId != nil {
		return true
	}

	return false
}

// SetEvaluationMarketplaceId gets a reference to the given string and assigns it to the EvaluationMarketplaceId field.
func (o *CustomerServiceMetricsFilterCriteria) SetEvaluationMarketplaceId(v string) {
	o.EvaluationMarketplaceId = &v
}

// GetListingCategories returns the ListingCategories field value if set, zero value otherwise.
func (o *CustomerServiceMetricsFilterCriteria) GetListingCategories() []string {
	if o == nil || o.ListingCategories == nil {
		var ret []string
		return ret
	}
	return *o.ListingCategories
}

// GetListingCategoriesOk returns a tuple with the ListingCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerServiceMetricsFilterCriteria) GetListingCategoriesOk() (*[]string, bool) {
	if o == nil || o.ListingCategories == nil {
		return nil, false
	}
	return o.ListingCategories, true
}

// HasListingCategories returns a boolean if a field has been set.
func (o *CustomerServiceMetricsFilterCriteria) HasListingCategories() bool {
	if o != nil && o.ListingCategories != nil {
		return true
	}

	return false
}

// SetListingCategories gets a reference to the given []string and assigns it to the ListingCategories field.
func (o *CustomerServiceMetricsFilterCriteria) SetListingCategories(v []string) {
	o.ListingCategories = &v
}

// GetShippingRegions returns the ShippingRegions field value if set, zero value otherwise.
func (o *CustomerServiceMetricsFilterCriteria) GetShippingRegions() []string {
	if o == nil || o.ShippingRegions == nil {
		var ret []string
		return ret
	}
	return *o.ShippingRegions
}

// GetShippingRegionsOk returns a tuple with the ShippingRegions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerServiceMetricsFilterCriteria) GetShippingRegionsOk() (*[]string, bool) {
	if o == nil || o.ShippingRegions == nil {
		return nil, false
	}
	return o.ShippingRegions, true
}

// HasShippingRegions returns a boolean if a field has been set.
func (o *CustomerServiceMetricsFilterCriteria) HasShippingRegions() bool {
	if o != nil && o.ShippingRegions != nil {
		return true
	}

	return false
}

// SetShippingRegions gets a reference to the given []string and assigns it to the ShippingRegions field.
func (o *CustomerServiceMetricsFilterCriteria) SetShippingRegions(v []string) {
	o.ShippingRegions = &v
}

func (o CustomerServiceMetricsFilterCriteria) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CustomerServiceMetricType != nil {
		toSerialize["customerServiceMetricType"] = o.CustomerServiceMetricType
	}
	if o.EvaluationMarketplaceId != nil {
		toSerialize["evaluationMarketplaceId"] = o.EvaluationMarketplaceId
	}
	if o.ListingCategories != nil {
		toSerialize["listingCategories"] = o.ListingCategories
	}
	if o.ShippingRegions != nil {
		toSerialize["shippingRegions"] = o.ShippingRegions
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerServiceMetricsFilterCriteria struct {
	value *CustomerServiceMetricsFilterCriteria
	isSet bool
}

func (v NullableCustomerServiceMetricsFilterCriteria) Get() *CustomerServiceMetricsFilterCriteria {
	return v.value
}

func (v *NullableCustomerServiceMetricsFilterCriteria) Set(val *CustomerServiceMetricsFilterCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerServiceMetricsFilterCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerServiceMetricsFilterCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerServiceMetricsFilterCriteria(val *CustomerServiceMetricsFilterCriteria) *NullableCustomerServiceMetricsFilterCriteria {
	return &NullableCustomerServiceMetricsFilterCriteria{value: val, isSet: true}
}

func (v NullableCustomerServiceMetricsFilterCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerServiceMetricsFilterCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

