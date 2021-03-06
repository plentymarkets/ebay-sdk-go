# PaymentPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryTypes** | Pointer to [**[]CategoryType**](CategoryType.md) | The &lt;b&gt;CategoryTypeEnum&lt;/b&gt; value to which this policy applies. Used to discern accounts that sell motor vehicles from those that don&#39;t. (Currently, each policy can be set to only one &lt;b&gt;categoryTypes&lt;/b&gt; value at a time.) | [optional] 
**Deposit** | Pointer to [**Deposit**](Deposit.md) |  | [optional] 
**Description** | Pointer to **string** | An optional seller-defined description of the payment policy for internal use (this value is not displayed to end users).  &lt;br&gt;&lt;br&gt;&lt;b&gt;Max length&lt;/b&gt;: 250 | [optional] 
**FullPaymentDueIn** | Pointer to [**TimeDuration**](TimeDuration.md) |  | [optional] 
**ImmediatePay** | Pointer to **bool** | If set to &lt;code&gt;true&lt;/code&gt;, payment is due upon receipt (eBay generates a receipt when the buyer agrees to purchase an item). &lt;br&gt;&lt;br&gt;This boolean must be set in the payment policy if the seller wants to create a listing that has an \&quot;immediate payment\&quot; requirement. The seller can change the immediate payment requirement at any time during the life cycle of a listing. &lt;br&gt;&lt;br&gt;The following must be true before a seller can apply an immediate payment requirement to an item:&lt;ul&gt;&lt;li&gt;The seller must have a PayPal Business account.&lt;/li&gt;&lt;li&gt;The Buy It Now price cannot be higher than $60,000 USD.&lt;/li&gt;&lt;li&gt;The eBay marketplace on which the item is listed must support PayPal payments.&lt;/li&gt;&lt;li&gt;The listing type must be fixed-price, or an auction with a Buy It Now option.&lt;/li&gt;&lt;/ul&gt;To enable the immediate payment requirement, the seller must also perform the following actions via API calls:&lt;ul&gt;&lt;li&gt;Provide a valid &lt;b&gt;paymentMethods.recipientAccountReference.referenceId&lt;/b&gt; value.&lt;/li&gt;&lt;li&gt;Offer PayPal as the only payment method for the item(s).&lt;/li&gt;&lt;li&gt;Specify all related costs to the buyer (because the buyer is not be able to use the Buyer Request Total feature in an immediate payment listing); these costs include flat-rate shipping costs for each domestic and international shipping service offered, package handling costs, and any shipping surcharges.&lt;/li&gt;&lt;li&gt;Include and set the &lt;b&gt;shippingProfileDiscountInfo&lt;/b&gt; container values if you are going to use promotional shipping discounts.&lt;/li&gt;&lt;/ul&gt;For more information, see the &lt;a href&#x3D;\&quot;http://pages.ebay.com/help/pay/require-immediate-payment.html\&quot;&gt;Understanding immediate payment&lt;/a&gt; Help page. &lt;br&gt;&lt;br&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt;Note:&lt;/b&gt; Listings created with the Inventory API must reference a payment policy that has &lt;b&gt;immediatePay&lt;/b&gt; is set to &lt;code&gt;true&lt;/code&gt;. Items listed with the Inventory API must also be fixed-price good-till-canceled (GTC) listings where PayPal is the only supported payment method (&lt;b&gt;paymentMethod&lt;/b&gt; must be set to &lt;code&gt;PAYPAL&lt;/code&gt;).&lt;/span&gt; &lt;br&gt;&lt;br&gt;&lt;b&gt;Default&lt;/b&gt;: false | [optional] 
**MarketplaceId** | Pointer to **string** | The ID of the eBay marketplace to which the payment policy applies. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/account/types/ba:MarketplaceIdEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**Name** | Pointer to **string** | A user-defined name for this payment policy. Names must be unique for policies assigned to the same marketplace. &lt;br&gt;&lt;br&gt;&lt;b&gt;Max length&lt;/b&gt;: 64 | [optional] 
**PaymentInstructions** | Pointer to **string** | This free-form string field gives sellers the ability add detailed payment instructions to their listings. The payment instructions appear on eBay&#39;s View Item and Checkout pages.  &lt;br&gt;&lt;br&gt;eBay recommends sellers use this field to clarify payment policies for motor vehicle listings on eBay Motors. For example, sellers can include the specifics on the deposit (if required), pickup/delivery arrangements, and full payment details on the vehicle. &lt;br&gt;&lt;br&gt;The field allows only 500 characters as input, but due to the way the eBay web site UI treats characters, this field can return more than 500 characters in the response. For example, characters like &amp; and &#39; (ampersand and single quote) count as 5 characters each.  &lt;br&gt;&lt;br&gt;&lt;b&gt;Max length:&lt;/b&gt; 1000. | [optional] 
**PaymentMethods** | Pointer to [**[]PaymentMethod**](PaymentMethod.md) | This container is returned to show the payment methods that are accepted for the payment business policy.  &lt;br&gt;&lt;br&gt;Sellers enabled for managed payments do not have to specify any electronic payment methods for listings, so this array will often be returned empty unless the payment business policy is intended for motor vehicle listings or other items in categories where offline payments are required or supported.  | [optional] 
**PaymentPolicyId** | Pointer to **string** | A unique eBay-assigned ID for a payment policy. This ID is generated when the policy is created. | [optional] 

## Methods

### NewPaymentPolicy

`func NewPaymentPolicy() *PaymentPolicy`

NewPaymentPolicy instantiates a new PaymentPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentPolicyWithDefaults

`func NewPaymentPolicyWithDefaults() *PaymentPolicy`

NewPaymentPolicyWithDefaults instantiates a new PaymentPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryTypes

`func (o *PaymentPolicy) GetCategoryTypes() []CategoryType`

GetCategoryTypes returns the CategoryTypes field if non-nil, zero value otherwise.

### GetCategoryTypesOk

`func (o *PaymentPolicy) GetCategoryTypesOk() (*[]CategoryType, bool)`

GetCategoryTypesOk returns a tuple with the CategoryTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryTypes

`func (o *PaymentPolicy) SetCategoryTypes(v []CategoryType)`

SetCategoryTypes sets CategoryTypes field to given value.

### HasCategoryTypes

`func (o *PaymentPolicy) HasCategoryTypes() bool`

HasCategoryTypes returns a boolean if a field has been set.

### GetDeposit

`func (o *PaymentPolicy) GetDeposit() Deposit`

GetDeposit returns the Deposit field if non-nil, zero value otherwise.

### GetDepositOk

`func (o *PaymentPolicy) GetDepositOk() (*Deposit, bool)`

GetDepositOk returns a tuple with the Deposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeposit

`func (o *PaymentPolicy) SetDeposit(v Deposit)`

SetDeposit sets Deposit field to given value.

### HasDeposit

`func (o *PaymentPolicy) HasDeposit() bool`

HasDeposit returns a boolean if a field has been set.

### GetDescription

`func (o *PaymentPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFullPaymentDueIn

`func (o *PaymentPolicy) GetFullPaymentDueIn() TimeDuration`

GetFullPaymentDueIn returns the FullPaymentDueIn field if non-nil, zero value otherwise.

### GetFullPaymentDueInOk

`func (o *PaymentPolicy) GetFullPaymentDueInOk() (*TimeDuration, bool)`

GetFullPaymentDueInOk returns a tuple with the FullPaymentDueIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullPaymentDueIn

`func (o *PaymentPolicy) SetFullPaymentDueIn(v TimeDuration)`

SetFullPaymentDueIn sets FullPaymentDueIn field to given value.

### HasFullPaymentDueIn

`func (o *PaymentPolicy) HasFullPaymentDueIn() bool`

HasFullPaymentDueIn returns a boolean if a field has been set.

### GetImmediatePay

`func (o *PaymentPolicy) GetImmediatePay() bool`

GetImmediatePay returns the ImmediatePay field if non-nil, zero value otherwise.

### GetImmediatePayOk

`func (o *PaymentPolicy) GetImmediatePayOk() (*bool, bool)`

GetImmediatePayOk returns a tuple with the ImmediatePay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmediatePay

`func (o *PaymentPolicy) SetImmediatePay(v bool)`

SetImmediatePay sets ImmediatePay field to given value.

### HasImmediatePay

`func (o *PaymentPolicy) HasImmediatePay() bool`

HasImmediatePay returns a boolean if a field has been set.

### GetMarketplaceId

`func (o *PaymentPolicy) GetMarketplaceId() string`

GetMarketplaceId returns the MarketplaceId field if non-nil, zero value otherwise.

### GetMarketplaceIdOk

`func (o *PaymentPolicy) GetMarketplaceIdOk() (*string, bool)`

GetMarketplaceIdOk returns a tuple with the MarketplaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceId

`func (o *PaymentPolicy) SetMarketplaceId(v string)`

SetMarketplaceId sets MarketplaceId field to given value.

### HasMarketplaceId

`func (o *PaymentPolicy) HasMarketplaceId() bool`

HasMarketplaceId returns a boolean if a field has been set.

### GetName

`func (o *PaymentPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaymentPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaymentPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PaymentPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPaymentInstructions

`func (o *PaymentPolicy) GetPaymentInstructions() string`

GetPaymentInstructions returns the PaymentInstructions field if non-nil, zero value otherwise.

### GetPaymentInstructionsOk

`func (o *PaymentPolicy) GetPaymentInstructionsOk() (*string, bool)`

GetPaymentInstructionsOk returns a tuple with the PaymentInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstructions

`func (o *PaymentPolicy) SetPaymentInstructions(v string)`

SetPaymentInstructions sets PaymentInstructions field to given value.

### HasPaymentInstructions

`func (o *PaymentPolicy) HasPaymentInstructions() bool`

HasPaymentInstructions returns a boolean if a field has been set.

### GetPaymentMethods

`func (o *PaymentPolicy) GetPaymentMethods() []PaymentMethod`

GetPaymentMethods returns the PaymentMethods field if non-nil, zero value otherwise.

### GetPaymentMethodsOk

`func (o *PaymentPolicy) GetPaymentMethodsOk() (*[]PaymentMethod, bool)`

GetPaymentMethodsOk returns a tuple with the PaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethods

`func (o *PaymentPolicy) SetPaymentMethods(v []PaymentMethod)`

SetPaymentMethods sets PaymentMethods field to given value.

### HasPaymentMethods

`func (o *PaymentPolicy) HasPaymentMethods() bool`

HasPaymentMethods returns a boolean if a field has been set.

### GetPaymentPolicyId

`func (o *PaymentPolicy) GetPaymentPolicyId() string`

GetPaymentPolicyId returns the PaymentPolicyId field if non-nil, zero value otherwise.

### GetPaymentPolicyIdOk

`func (o *PaymentPolicy) GetPaymentPolicyIdOk() (*string, bool)`

GetPaymentPolicyIdOk returns a tuple with the PaymentPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentPolicyId

`func (o *PaymentPolicy) SetPaymentPolicyId(v string)`

SetPaymentPolicyId sets PaymentPolicyId field to given value.

### HasPaymentPolicyId

`func (o *PaymentPolicy) HasPaymentPolicyId() bool`

HasPaymentPolicyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


