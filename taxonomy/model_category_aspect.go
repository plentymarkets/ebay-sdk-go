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

// CategoryAspect struct for CategoryAspect
type CategoryAspect struct {
	Category *Category `json:"category,omitempty"`
	// A list of aspect metadata that is used to describe the items in a particular leaf category.
	Aspects *[]Aspect `json:"aspects,omitempty"`
}

// NewCategoryAspect instantiates a new CategoryAspect object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryAspect() *CategoryAspect {
	this := CategoryAspect{}
	return &this
}

// NewCategoryAspectWithDefaults instantiates a new CategoryAspect object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryAspectWithDefaults() *CategoryAspect {
	this := CategoryAspect{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *CategoryAspect) GetCategory() Category {
	if o == nil || o.Category == nil {
		var ret Category
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryAspect) GetCategoryOk() (*Category, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *CategoryAspect) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given Category and assigns it to the Category field.
func (o *CategoryAspect) SetCategory(v Category) {
	o.Category = &v
}

// GetAspects returns the Aspects field value if set, zero value otherwise.
func (o *CategoryAspect) GetAspects() []Aspect {
	if o == nil || o.Aspects == nil {
		var ret []Aspect
		return ret
	}
	return *o.Aspects
}

// GetAspectsOk returns a tuple with the Aspects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryAspect) GetAspectsOk() (*[]Aspect, bool) {
	if o == nil || o.Aspects == nil {
		return nil, false
	}
	return o.Aspects, true
}

// HasAspects returns a boolean if a field has been set.
func (o *CategoryAspect) HasAspects() bool {
	if o != nil && o.Aspects != nil {
		return true
	}

	return false
}

// SetAspects gets a reference to the given []Aspect and assigns it to the Aspects field.
func (o *CategoryAspect) SetAspects(v []Aspect) {
	o.Aspects = &v
}

func (o CategoryAspect) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.Aspects != nil {
		toSerialize["aspects"] = o.Aspects
	}
	return json.Marshal(toSerialize)
}

type NullableCategoryAspect struct {
	value *CategoryAspect
	isSet bool
}

func (v NullableCategoryAspect) Get() *CategoryAspect {
	return v.value
}

func (v *NullableCategoryAspect) Set(val *CategoryAspect) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryAspect) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryAspect) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryAspect(val *CategoryAspect) *NullableCategoryAspect {
	return &NullableCategoryAspect{value: val, isSet: true}
}

func (v NullableCategoryAspect) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoryAspect) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

