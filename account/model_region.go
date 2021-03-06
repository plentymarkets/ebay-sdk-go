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

// Region This type defines information for a region.
type Region struct {
	// A string that indicates the name of a region, as defined by eBay. A \"region\" can be either a 'world region' (e.g., the \"Middle East\" or \"Southeast Asia\") or a country, as represented with a two-letter country code. Use <b>GeteBayDetails</b> to get the values accepted by this field. <p>The values that you're allowed to use for a specific <b>regionName</b> field depend on the context in which you are setting the value. For details on how to set the values for this field, see <a href=\"/api-docs/sell/static/seller-accounts/ht_shipping-worldwide.html#shipToLocations\">The shipToLocations container</a>.
	RegionName *string `json:"regionName,omitempty"`
	// Reserved for future use. <!--The region's type, which can be one of the following: 'COUNTRY', 'COUNTRY_REGION', 'STATE_OR_PROVINCE', 'WORLD_REGION', or 'WORLDWIDE'.--> For implementation help, refer to <a href='https://developer.ebay.com/api-docs/sell/account/types/ba:RegionTypeEnum'>eBay API documentation</a>
	RegionType *string `json:"regionType,omitempty"`
}

// NewRegion instantiates a new Region object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegion() *Region {
	this := Region{}
	return &this
}

// NewRegionWithDefaults instantiates a new Region object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegionWithDefaults() *Region {
	this := Region{}
	return &this
}

// GetRegionName returns the RegionName field value if set, zero value otherwise.
func (o *Region) GetRegionName() string {
	if o == nil || o.RegionName == nil {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Region) GetRegionNameOk() (*string, bool) {
	if o == nil || o.RegionName == nil {
		return nil, false
	}
	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *Region) HasRegionName() bool {
	if o != nil && o.RegionName != nil {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *Region) SetRegionName(v string) {
	o.RegionName = &v
}

// GetRegionType returns the RegionType field value if set, zero value otherwise.
func (o *Region) GetRegionType() string {
	if o == nil || o.RegionType == nil {
		var ret string
		return ret
	}
	return *o.RegionType
}

// GetRegionTypeOk returns a tuple with the RegionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Region) GetRegionTypeOk() (*string, bool) {
	if o == nil || o.RegionType == nil {
		return nil, false
	}
	return o.RegionType, true
}

// HasRegionType returns a boolean if a field has been set.
func (o *Region) HasRegionType() bool {
	if o != nil && o.RegionType != nil {
		return true
	}

	return false
}

// SetRegionType gets a reference to the given string and assigns it to the RegionType field.
func (o *Region) SetRegionType(v string) {
	o.RegionType = &v
}

func (o Region) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RegionName != nil {
		toSerialize["regionName"] = o.RegionName
	}
	if o.RegionType != nil {
		toSerialize["regionType"] = o.RegionType
	}
	return json.Marshal(toSerialize)
}

type NullableRegion struct {
	value *Region
	isSet bool
}

func (v NullableRegion) Get() *Region {
	return v.value
}

func (v *NullableRegion) Set(val *Region) {
	v.value = val
	v.isSet = true
}

func (v NullableRegion) IsSet() bool {
	return v.isSet
}

func (v *NullableRegion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegion(val *Region) *NullableRegion {
	return &NullableRegion{value: val, isSet: true}
}

func (v NullableRegion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


