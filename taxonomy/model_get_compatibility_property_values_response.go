/*
Taxonomy API

Use the Taxonomy API to discover the most appropriate eBay categories under which sellers can offer inventory items for sale, and the most likely categories under which buyers can browse or search for items to purchase. In addition, the Taxonomy API provides metadata about the required and recommended category aspects to include in listings, and also has two operations to retrieve parts compatibility information.

API version: v1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package taxonomy

import (
	"encoding/json"
)

// GetCompatibilityPropertyValuesResponse The base response type of the getCompatibilityPropertyValues method.
type GetCompatibilityPropertyValuesResponse struct {
	// This array contains all compatible vehicle property values that match the specified eBay marketplace, specified eBay category, and filters in the request. If the compatibility_property parameter value in the request is 'Trim', each value returned in each value field will be a different vehicle trim, applicable to any filters that are set in the filter query parameter of the request, and also based on the eBay marketplace and category specified in the call request.
	CompatibilityPropertyValues *[]CompatibilityPropertyValue `json:"compatibilityPropertyValues,omitempty"`
}

// NewGetCompatibilityPropertyValuesResponse instantiates a new GetCompatibilityPropertyValuesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCompatibilityPropertyValuesResponse() *GetCompatibilityPropertyValuesResponse {
	this := GetCompatibilityPropertyValuesResponse{}
	return &this
}

// NewGetCompatibilityPropertyValuesResponseWithDefaults instantiates a new GetCompatibilityPropertyValuesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCompatibilityPropertyValuesResponseWithDefaults() *GetCompatibilityPropertyValuesResponse {
	this := GetCompatibilityPropertyValuesResponse{}
	return &this
}

// GetCompatibilityPropertyValues returns the CompatibilityPropertyValues field value if set, zero value otherwise.
func (o *GetCompatibilityPropertyValuesResponse) GetCompatibilityPropertyValues() []CompatibilityPropertyValue {
	if o == nil || o.CompatibilityPropertyValues == nil {
		var ret []CompatibilityPropertyValue
		return ret
	}
	return *o.CompatibilityPropertyValues
}

// GetCompatibilityPropertyValuesOk returns a tuple with the CompatibilityPropertyValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCompatibilityPropertyValuesResponse) GetCompatibilityPropertyValuesOk() (*[]CompatibilityPropertyValue, bool) {
	if o == nil || o.CompatibilityPropertyValues == nil {
		return nil, false
	}
	return o.CompatibilityPropertyValues, true
}

// HasCompatibilityPropertyValues returns a boolean if a field has been set.
func (o *GetCompatibilityPropertyValuesResponse) HasCompatibilityPropertyValues() bool {
	if o != nil && o.CompatibilityPropertyValues != nil {
		return true
	}

	return false
}

// SetCompatibilityPropertyValues gets a reference to the given []CompatibilityPropertyValue and assigns it to the CompatibilityPropertyValues field.
func (o *GetCompatibilityPropertyValuesResponse) SetCompatibilityPropertyValues(v []CompatibilityPropertyValue) {
	o.CompatibilityPropertyValues = &v
}

func (o GetCompatibilityPropertyValuesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CompatibilityPropertyValues != nil {
		toSerialize["compatibilityPropertyValues"] = o.CompatibilityPropertyValues
	}
	return json.Marshal(toSerialize)
}

type NullableGetCompatibilityPropertyValuesResponse struct {
	value *GetCompatibilityPropertyValuesResponse
	isSet bool
}

func (v NullableGetCompatibilityPropertyValuesResponse) Get() *GetCompatibilityPropertyValuesResponse {
	return v.value
}

func (v *NullableGetCompatibilityPropertyValuesResponse) Set(val *GetCompatibilityPropertyValuesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCompatibilityPropertyValuesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCompatibilityPropertyValuesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCompatibilityPropertyValuesResponse(val *GetCompatibilityPropertyValuesResponse) *NullableGetCompatibilityPropertyValuesResponse {
	return &NullableGetCompatibilityPropertyValuesResponse{value: val, isSet: true}
}

func (v NullableGetCompatibilityPropertyValuesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCompatibilityPropertyValuesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


