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

// ShippingService A complex type that defines the available shipping services offered in the parent <b>shippingOptions</b> container. The shipping services specified here must be able to accommodate the <b>optionType</b> defined in the parent <b>shippingOption</b> container (either DOMESTIC or INTERNATIONAL). <p class=\"tablenote\"><b>Tip:</b> For more on setting up shipping services, see <a href=\"/api-docs/sell/static/seller-accounts/ht_shipping-free.html#shippingServices\">Setting the shipping carrier and shipping service values</a>.</p>
type ShippingService struct {
	AdditionalShippingCost *Amount `json:"additionalShippingCost,omitempty"`
	// This field is only applicable to vehicle categories on eBay Motors (US and Canada). If set to <code>true</code>, the buyer is responsible for picking up the vehicle. Otherwise, the seller should specify the vehicle pickup arrangements in the item description. <br><br>The seller cannot modify this flag if the vehicle has bids or if the listing ends within 12 hours. <br><br><b>Default</b>: false
	BuyerResponsibleForPickup *bool `json:"buyerResponsibleForPickup,omitempty"`
	// This field is applicable for only items listed in vehicle categories on eBay Motors (US and Canada). If set to <code>true</code>, the buyer is responsible for the shipment of the vehicle. Otherwise, the seller should specify the vehicle shipping arrangements in the item description. <br><br>The seller cannot modify this flag if the vehicle has bids or if the listing ends within 12 hours. <br><br><b>Default</b>: false
	BuyerResponsibleForShipping *bool `json:"buyerResponsibleForShipping,omitempty"`
	CashOnDeliveryFee *Amount `json:"cashOnDeliveryFee,omitempty"`
	// If set to <code>true</code>, the seller offers free shipping to the buyer. This field can only be included and set to 'true' for the first domestic shipping service option specified in the <b>shippingServices</b> container (it is ignored if set for subsequent shipping services). The first specified shipping service option has a <b>sortOrder</b> value of <code>1</code> or (if the sortOrderId field is not used) it is the shipping service option that's specified first in the <b>shippingServices</b> container.
	FreeShipping *bool `json:"freeShipping,omitempty"`
	// The shipping carrier, such as 'USPS', 'FedEx', 'UPS', and so on.
	ShippingCarrierCode *string `json:"shippingCarrierCode,omitempty"`
	ShippingCost *Amount `json:"shippingCost,omitempty"`
	// The shipping service that the shipping carrier uses to ship an item. For example, an overnight, two-day delivery, or other type of service. For details on configuring shipping services, see <a href=\"/api-docs/sell/static/seller-accounts/ht_shipping-free.html#shippingServices\">Setting the shipping carrier and shipping service values</a>.
	ShippingServiceCode *string `json:"shippingServiceCode,omitempty"`
	ShipToLocations *RegionSet `json:"shipToLocations,omitempty"`
	// This integer value controls the order that this shipping service option appears in the View Item and Checkout pages, as related to the other specified shipping service options. <br><br>Sellers can specify up to four domestic shipping services (in four separate <b>shippingService</b> containers), so valid values are 1, 2, 3, and 4. A shipping service option with a <b>sortOrder</b> value of '1' appears at the top of View Item and Checkout pages. Conversely, a shipping service option with a <b>sortOrder</b> value of '4' appears at the bottom of the list. <br><br>Sellers can specify up to five international shipping services (in five separate <b>shippingService</b> containers, so valid values for international shipping services are 1, 2, 3, 4, and 5. Similarly to domestic shipping service options, the <b>sortOrder</b> value of a international shipping service option controls the placement of that shipping service option in the View Item and Checkout pages. Set up different domestic and international services by configuring two <b>shippingOptions</b> containers, where you set <b>shippingOptions.optionType</b> to either <code>DOMESTIC</code> or <code>INTERNATIONAL</code> to indicate the area supported by the listed shipping services. <br><br>If the <b>sortOrder</b> field is not supplied, the order of domestic and international shipping service options is determined by the order in which they are listed in the API call. <br><br><b>Min</b>: 1. <b>Max</b>: 4 (for domestic shipping service) or 5 (for international shipping service).
	SortOrder *int32 `json:"sortOrder,omitempty"`
	Surcharge *Amount `json:"surcharge,omitempty"`
}

// NewShippingService instantiates a new ShippingService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingService() *ShippingService {
	this := ShippingService{}
	return &this
}

// NewShippingServiceWithDefaults instantiates a new ShippingService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingServiceWithDefaults() *ShippingService {
	this := ShippingService{}
	return &this
}

// GetAdditionalShippingCost returns the AdditionalShippingCost field value if set, zero value otherwise.
func (o *ShippingService) GetAdditionalShippingCost() Amount {
	if o == nil || o.AdditionalShippingCost == nil {
		var ret Amount
		return ret
	}
	return *o.AdditionalShippingCost
}

// GetAdditionalShippingCostOk returns a tuple with the AdditionalShippingCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingService) GetAdditionalShippingCostOk() (*Amount, bool) {
	if o == nil || o.AdditionalShippingCost == nil {
		return nil, false
	}
	return o.AdditionalShippingCost, true
}

// HasAdditionalShippingCost returns a boolean if a field has been set.
func (o *ShippingService) HasAdditionalShippingCost() bool {
	if o != nil && o.AdditionalShippingCost != nil {
		return true
	}

	return false
}

// SetAdditionalShippingCost gets a reference to the given Amount and assigns it to the AdditionalShippingCost field.
func (o *ShippingService) SetAdditionalShippingCost(v Amount) {
	o.AdditionalShippingCost = &v
}

// GetBuyerResponsibleForPickup returns the BuyerResponsibleForPickup field value if set, zero value otherwise.
func (o *ShippingService) GetBuyerResponsibleForPickup() bool {
	if o == nil || o.BuyerResponsibleForPickup == nil {
		var ret bool
		return ret
	}
	return *o.BuyerResponsibleForPickup
}

// GetBuyerResponsibleForPickupOk returns a tuple with the BuyerResponsibleForPickup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingService) GetBuyerResponsibleForPickupOk() (*bool, bool) {
	if o == nil || o.BuyerResponsibleForPickup == nil {
		return nil, false
	}
	return o.BuyerResponsibleForPickup, true
}

// HasBuyerResponsibleForPickup returns a boolean if a field has been set.
func (o *ShippingService) HasBuyerResponsibleForPickup() bool {
	if o != nil && o.BuyerResponsibleForPickup != nil {
		return true
	}

	return false
}

// SetBuyerResponsibleForPickup gets a reference to the given bool and assigns it to the BuyerResponsibleForPickup field.
func (o *ShippingService) SetBuyerResponsibleForPickup(v bool) {
	o.BuyerResponsibleForPickup = &v
}

// GetBuyerResponsibleForShipping returns the BuyerResponsibleForShipping field value if set, zero value otherwise.
func (o *ShippingService) GetBuyerResponsibleForShipping() bool {
	if o == nil || o.BuyerResponsibleForShipping == nil {
		var ret bool
		return ret
	}
	return *o.BuyerResponsibleForShipping
}

// GetBuyerResponsibleForShippingOk returns a tuple with the BuyerResponsibleForShipping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingService) GetBuyerResponsibleForShippingOk() (*bool, bool) {
	if o == nil || o.BuyerResponsibleForShipping == nil {
		return nil, false
	}
	return o.BuyerResponsibleForShipping, true
}

// HasBuyerResponsibleForShipping returns a boolean if a field has been set.
func (o *ShippingService) HasBuyerResponsibleForShipping() bool {
	if o != nil && o.BuyerResponsibleForShipping != nil {
		return true
	}

	return false
}

// SetBuyerResponsibleForShipping gets a reference to the given bool and assigns it to the BuyerResponsibleForShipping field.
func (o *ShippingService) SetBuyerResponsibleForShipping(v bool) {
	o.BuyerResponsibleForShipping = &v
}

// GetCashOnDeliveryFee returns the CashOnDeliveryFee field value if set, zero value otherwise.
func (o *ShippingService) GetCashOnDeliveryFee() Amount {
	if o == nil || o.CashOnDeliveryFee == nil {
		var ret Amount
		return ret
	}
	return *o.CashOnDeliveryFee
}

// GetCashOnDeliveryFeeOk returns a tuple with the CashOnDeliveryFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingService) GetCashOnDeliveryFeeOk() (*Amount, bool) {
	if o == nil || o.CashOnDeliveryFee == nil {
		return nil, false
	}
	return o.CashOnDeliveryFee, true
}

// HasCashOnDeliveryFee returns a boolean if a field has been set.
func (o *ShippingService) HasCashOnDeliveryFee() bool {
	if o != nil && o.CashOnDeliveryFee != nil {
		return true
	}

	return false
}

// SetCashOnDeliveryFee gets a reference to the given Amount and assigns it to the CashOnDeliveryFee field.
func (o *ShippingService) SetCashOnDeliveryFee(v Amount) {
	o.CashOnDeliveryFee = &v
}

// GetFreeShipping returns the FreeShipping field value if set, zero value otherwise.
func (o *ShippingService) GetFreeShipping() bool {
	if o == nil || o.FreeShipping == nil {
		var ret bool
		return ret
	}
	return *o.FreeShipping
}

// GetFreeShippingOk returns a tuple with the FreeShipping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingService) GetFreeShippingOk() (*bool, bool) {
	if o == nil || o.FreeShipping == nil {
		return nil, false
	}
	return o.FreeShipping, true
}

// HasFreeShipping returns a boolean if a field has been set.
func (o *ShippingService) HasFreeShipping() bool {
	if o != nil && o.FreeShipping != nil {
		return true
	}

	return false
}

// SetFreeShipping gets a reference to the given bool and assigns it to the FreeShipping field.
func (o *ShippingService) SetFreeShipping(v bool) {
	o.FreeShipping = &v
}

// GetShippingCarrierCode returns the ShippingCarrierCode field value if set, zero value otherwise.
func (o *ShippingService) GetShippingCarrierCode() string {
	if o == nil || o.ShippingCarrierCode == nil {
		var ret string
		return ret
	}
	return *o.ShippingCarrierCode
}

// GetShippingCarrierCodeOk returns a tuple with the ShippingCarrierCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingService) GetShippingCarrierCodeOk() (*string, bool) {
	if o == nil || o.ShippingCarrierCode == nil {
		return nil, false
	}
	return o.ShippingCarrierCode, true
}

// HasShippingCarrierCode returns a boolean if a field has been set.
func (o *ShippingService) HasShippingCarrierCode() bool {
	if o != nil && o.ShippingCarrierCode != nil {
		return true
	}

	return false
}

// SetShippingCarrierCode gets a reference to the given string and assigns it to the ShippingCarrierCode field.
func (o *ShippingService) SetShippingCarrierCode(v string) {
	o.ShippingCarrierCode = &v
}

// GetShippingCost returns the ShippingCost field value if set, zero value otherwise.
func (o *ShippingService) GetShippingCost() Amount {
	if o == nil || o.ShippingCost == nil {
		var ret Amount
		return ret
	}
	return *o.ShippingCost
}

// GetShippingCostOk returns a tuple with the ShippingCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingService) GetShippingCostOk() (*Amount, bool) {
	if o == nil || o.ShippingCost == nil {
		return nil, false
	}
	return o.ShippingCost, true
}

// HasShippingCost returns a boolean if a field has been set.
func (o *ShippingService) HasShippingCost() bool {
	if o != nil && o.ShippingCost != nil {
		return true
	}

	return false
}

// SetShippingCost gets a reference to the given Amount and assigns it to the ShippingCost field.
func (o *ShippingService) SetShippingCost(v Amount) {
	o.ShippingCost = &v
}

// GetShippingServiceCode returns the ShippingServiceCode field value if set, zero value otherwise.
func (o *ShippingService) GetShippingServiceCode() string {
	if o == nil || o.ShippingServiceCode == nil {
		var ret string
		return ret
	}
	return *o.ShippingServiceCode
}

// GetShippingServiceCodeOk returns a tuple with the ShippingServiceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingService) GetShippingServiceCodeOk() (*string, bool) {
	if o == nil || o.ShippingServiceCode == nil {
		return nil, false
	}
	return o.ShippingServiceCode, true
}

// HasShippingServiceCode returns a boolean if a field has been set.
func (o *ShippingService) HasShippingServiceCode() bool {
	if o != nil && o.ShippingServiceCode != nil {
		return true
	}

	return false
}

// SetShippingServiceCode gets a reference to the given string and assigns it to the ShippingServiceCode field.
func (o *ShippingService) SetShippingServiceCode(v string) {
	o.ShippingServiceCode = &v
}

// GetShipToLocations returns the ShipToLocations field value if set, zero value otherwise.
func (o *ShippingService) GetShipToLocations() RegionSet {
	if o == nil || o.ShipToLocations == nil {
		var ret RegionSet
		return ret
	}
	return *o.ShipToLocations
}

// GetShipToLocationsOk returns a tuple with the ShipToLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingService) GetShipToLocationsOk() (*RegionSet, bool) {
	if o == nil || o.ShipToLocations == nil {
		return nil, false
	}
	return o.ShipToLocations, true
}

// HasShipToLocations returns a boolean if a field has been set.
func (o *ShippingService) HasShipToLocations() bool {
	if o != nil && o.ShipToLocations != nil {
		return true
	}

	return false
}

// SetShipToLocations gets a reference to the given RegionSet and assigns it to the ShipToLocations field.
func (o *ShippingService) SetShipToLocations(v RegionSet) {
	o.ShipToLocations = &v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *ShippingService) GetSortOrder() int32 {
	if o == nil || o.SortOrder == nil {
		var ret int32
		return ret
	}
	return *o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingService) GetSortOrderOk() (*int32, bool) {
	if o == nil || o.SortOrder == nil {
		return nil, false
	}
	return o.SortOrder, true
}

// HasSortOrder returns a boolean if a field has been set.
func (o *ShippingService) HasSortOrder() bool {
	if o != nil && o.SortOrder != nil {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given int32 and assigns it to the SortOrder field.
func (o *ShippingService) SetSortOrder(v int32) {
	o.SortOrder = &v
}

// GetSurcharge returns the Surcharge field value if set, zero value otherwise.
func (o *ShippingService) GetSurcharge() Amount {
	if o == nil || o.Surcharge == nil {
		var ret Amount
		return ret
	}
	return *o.Surcharge
}

// GetSurchargeOk returns a tuple with the Surcharge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingService) GetSurchargeOk() (*Amount, bool) {
	if o == nil || o.Surcharge == nil {
		return nil, false
	}
	return o.Surcharge, true
}

// HasSurcharge returns a boolean if a field has been set.
func (o *ShippingService) HasSurcharge() bool {
	if o != nil && o.Surcharge != nil {
		return true
	}

	return false
}

// SetSurcharge gets a reference to the given Amount and assigns it to the Surcharge field.
func (o *ShippingService) SetSurcharge(v Amount) {
	o.Surcharge = &v
}

func (o ShippingService) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdditionalShippingCost != nil {
		toSerialize["additionalShippingCost"] = o.AdditionalShippingCost
	}
	if o.BuyerResponsibleForPickup != nil {
		toSerialize["buyerResponsibleForPickup"] = o.BuyerResponsibleForPickup
	}
	if o.BuyerResponsibleForShipping != nil {
		toSerialize["buyerResponsibleForShipping"] = o.BuyerResponsibleForShipping
	}
	if o.CashOnDeliveryFee != nil {
		toSerialize["cashOnDeliveryFee"] = o.CashOnDeliveryFee
	}
	if o.FreeShipping != nil {
		toSerialize["freeShipping"] = o.FreeShipping
	}
	if o.ShippingCarrierCode != nil {
		toSerialize["shippingCarrierCode"] = o.ShippingCarrierCode
	}
	if o.ShippingCost != nil {
		toSerialize["shippingCost"] = o.ShippingCost
	}
	if o.ShippingServiceCode != nil {
		toSerialize["shippingServiceCode"] = o.ShippingServiceCode
	}
	if o.ShipToLocations != nil {
		toSerialize["shipToLocations"] = o.ShipToLocations
	}
	if o.SortOrder != nil {
		toSerialize["sortOrder"] = o.SortOrder
	}
	if o.Surcharge != nil {
		toSerialize["surcharge"] = o.Surcharge
	}
	return json.Marshal(toSerialize)
}

type NullableShippingService struct {
	value *ShippingService
	isSet bool
}

func (v NullableShippingService) Get() *ShippingService {
	return v.value
}

func (v *NullableShippingService) Set(val *ShippingService) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingService) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingService(val *ShippingService) *NullableShippingService {
	return &NullableShippingService{value: val, isSet: true}
}

func (v NullableShippingService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


