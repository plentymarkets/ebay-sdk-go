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

// CategorySuggestion This type contains information about a suggested category tree leaf node that corresponds to keywords provided in the request. It includes details about each of the category's ancestor nodes extending up to the root of the category tree.
type CategorySuggestion struct {
	Category *Category `json:"category,omitempty"`
	// An ordered list of category references that describes the location of the suggested category in the specified category tree. The list identifies the category's ancestry as a sequence of parent nodes, from the current node's immediate parent to the root node of the category tree. Note: The root node of a full default category tree includes categoryId and categoryName fields, but their values should not be relied upon. They provide no useful information for application development.
	CategoryTreeNodeAncestors *[]AncestorReference `json:"categoryTreeNodeAncestors,omitempty"`
	// The absolute level of the category tree node in the hierarchy of its category tree. Note: The root node of any full category tree is always at level 0.
	CategoryTreeNodeLevel *int32 `json:"categoryTreeNodeLevel,omitempty"`
	// This field is reserved for internal or future use.
	Relevancy *string `json:"relevancy,omitempty"`
}

// NewCategorySuggestion instantiates a new CategorySuggestion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategorySuggestion() *CategorySuggestion {
	this := CategorySuggestion{}
	return &this
}

// NewCategorySuggestionWithDefaults instantiates a new CategorySuggestion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategorySuggestionWithDefaults() *CategorySuggestion {
	this := CategorySuggestion{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *CategorySuggestion) GetCategory() Category {
	if o == nil || o.Category == nil {
		var ret Category
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategorySuggestion) GetCategoryOk() (*Category, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *CategorySuggestion) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given Category and assigns it to the Category field.
func (o *CategorySuggestion) SetCategory(v Category) {
	o.Category = &v
}

// GetCategoryTreeNodeAncestors returns the CategoryTreeNodeAncestors field value if set, zero value otherwise.
func (o *CategorySuggestion) GetCategoryTreeNodeAncestors() []AncestorReference {
	if o == nil || o.CategoryTreeNodeAncestors == nil {
		var ret []AncestorReference
		return ret
	}
	return *o.CategoryTreeNodeAncestors
}

// GetCategoryTreeNodeAncestorsOk returns a tuple with the CategoryTreeNodeAncestors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategorySuggestion) GetCategoryTreeNodeAncestorsOk() (*[]AncestorReference, bool) {
	if o == nil || o.CategoryTreeNodeAncestors == nil {
		return nil, false
	}
	return o.CategoryTreeNodeAncestors, true
}

// HasCategoryTreeNodeAncestors returns a boolean if a field has been set.
func (o *CategorySuggestion) HasCategoryTreeNodeAncestors() bool {
	if o != nil && o.CategoryTreeNodeAncestors != nil {
		return true
	}

	return false
}

// SetCategoryTreeNodeAncestors gets a reference to the given []AncestorReference and assigns it to the CategoryTreeNodeAncestors field.
func (o *CategorySuggestion) SetCategoryTreeNodeAncestors(v []AncestorReference) {
	o.CategoryTreeNodeAncestors = &v
}

// GetCategoryTreeNodeLevel returns the CategoryTreeNodeLevel field value if set, zero value otherwise.
func (o *CategorySuggestion) GetCategoryTreeNodeLevel() int32 {
	if o == nil || o.CategoryTreeNodeLevel == nil {
		var ret int32
		return ret
	}
	return *o.CategoryTreeNodeLevel
}

// GetCategoryTreeNodeLevelOk returns a tuple with the CategoryTreeNodeLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategorySuggestion) GetCategoryTreeNodeLevelOk() (*int32, bool) {
	if o == nil || o.CategoryTreeNodeLevel == nil {
		return nil, false
	}
	return o.CategoryTreeNodeLevel, true
}

// HasCategoryTreeNodeLevel returns a boolean if a field has been set.
func (o *CategorySuggestion) HasCategoryTreeNodeLevel() bool {
	if o != nil && o.CategoryTreeNodeLevel != nil {
		return true
	}

	return false
}

// SetCategoryTreeNodeLevel gets a reference to the given int32 and assigns it to the CategoryTreeNodeLevel field.
func (o *CategorySuggestion) SetCategoryTreeNodeLevel(v int32) {
	o.CategoryTreeNodeLevel = &v
}

// GetRelevancy returns the Relevancy field value if set, zero value otherwise.
func (o *CategorySuggestion) GetRelevancy() string {
	if o == nil || o.Relevancy == nil {
		var ret string
		return ret
	}
	return *o.Relevancy
}

// GetRelevancyOk returns a tuple with the Relevancy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategorySuggestion) GetRelevancyOk() (*string, bool) {
	if o == nil || o.Relevancy == nil {
		return nil, false
	}
	return o.Relevancy, true
}

// HasRelevancy returns a boolean if a field has been set.
func (o *CategorySuggestion) HasRelevancy() bool {
	if o != nil && o.Relevancy != nil {
		return true
	}

	return false
}

// SetRelevancy gets a reference to the given string and assigns it to the Relevancy field.
func (o *CategorySuggestion) SetRelevancy(v string) {
	o.Relevancy = &v
}

func (o CategorySuggestion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.CategoryTreeNodeAncestors != nil {
		toSerialize["categoryTreeNodeAncestors"] = o.CategoryTreeNodeAncestors
	}
	if o.CategoryTreeNodeLevel != nil {
		toSerialize["categoryTreeNodeLevel"] = o.CategoryTreeNodeLevel
	}
	if o.Relevancy != nil {
		toSerialize["relevancy"] = o.Relevancy
	}
	return json.Marshal(toSerialize)
}

type NullableCategorySuggestion struct {
	value *CategorySuggestion
	isSet bool
}

func (v NullableCategorySuggestion) Get() *CategorySuggestion {
	return v.value
}

func (v *NullableCategorySuggestion) Set(val *CategorySuggestion) {
	v.value = val
	v.isSet = true
}

func (v NullableCategorySuggestion) IsSet() bool {
	return v.isSet
}

func (v *NullableCategorySuggestion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategorySuggestion(val *CategorySuggestion) *NullableCategorySuggestion {
	return &NullableCategorySuggestion{value: val, isSet: true}
}

func (v NullableCategorySuggestion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategorySuggestion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


