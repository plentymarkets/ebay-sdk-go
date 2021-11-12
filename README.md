# Go API client for ebayaccount

The <b>Account API</b> gives sellers the ability to configure their eBay seller accounts, including the seller's policies (the Fulfillment Policy, Payment Policy, and Return Policy), opt in and out of eBay seller programs, configure sales tax tables, and get account information.  <br><br>For details on the availability of the methods in this API, see <a href=\"/api-docs/sell/account/overview.html#requirements\">Account API requirements and restrictions</a>.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1.6.3
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./ebayaccount"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.ebay.com/sell/account/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*FulfillmentPolicyApi* | [**CreateFulfillmentPolicy**](docs/FulfillmentPolicyApi.md#createfulfillmentpolicy) | **Post** /fulfillment_policy | 
*FulfillmentPolicyApi* | [**DeleteFulfillmentPolicy**](docs/FulfillmentPolicyApi.md#deletefulfillmentpolicy) | **Delete** /fulfillment_policy/{fulfillmentPolicyId} | 
*FulfillmentPolicyApi* | [**GetFulfillmentPolicies**](docs/FulfillmentPolicyApi.md#getfulfillmentpolicies) | **Get** /fulfillment_policy | 
*FulfillmentPolicyApi* | [**GetFulfillmentPolicy**](docs/FulfillmentPolicyApi.md#getfulfillmentpolicy) | **Get** /fulfillment_policy/{fulfillmentPolicyId} | 
*FulfillmentPolicyApi* | [**GetFulfillmentPolicyByName**](docs/FulfillmentPolicyApi.md#getfulfillmentpolicybyname) | **Get** /fulfillment_policy/get_by_policy_name | 
*FulfillmentPolicyApi* | [**UpdateFulfillmentPolicy**](docs/FulfillmentPolicyApi.md#updatefulfillmentpolicy) | **Put** /fulfillment_policy/{fulfillmentPolicyId} | 
*KycApi* | [**GetKYC**](docs/KycApi.md#getkyc) | **Get** /kyc | 
*OnboardingApi* | [**GetPaymentsProgramOnboarding**](docs/OnboardingApi.md#getpaymentsprogramonboarding) | **Get** /payments_program/{marketplace_id}/{payments_program_type}/onboarding | 
*PaymentPolicyApi* | [**CreatePaymentPolicy**](docs/PaymentPolicyApi.md#createpaymentpolicy) | **Post** /payment_policy | 
*PaymentPolicyApi* | [**DeletePaymentPolicy**](docs/PaymentPolicyApi.md#deletepaymentpolicy) | **Delete** /payment_policy/{payment_policy_id} | 
*PaymentPolicyApi* | [**GetPaymentPolicies**](docs/PaymentPolicyApi.md#getpaymentpolicies) | **Get** /payment_policy | 
*PaymentPolicyApi* | [**GetPaymentPolicy**](docs/PaymentPolicyApi.md#getpaymentpolicy) | **Get** /payment_policy/{payment_policy_id} | 
*PaymentPolicyApi* | [**GetPaymentPolicyByName**](docs/PaymentPolicyApi.md#getpaymentpolicybyname) | **Get** /payment_policy/get_by_policy_name | 
*PaymentPolicyApi* | [**UpdatePaymentPolicy**](docs/PaymentPolicyApi.md#updatepaymentpolicy) | **Put** /payment_policy/{payment_policy_id} | 
*PaymentsProgramApi* | [**GetPaymentsProgram**](docs/PaymentsProgramApi.md#getpaymentsprogram) | **Get** /payments_program/{marketplace_id}/{payments_program_type} | 
*PrivilegeApi* | [**GetPrivileges**](docs/PrivilegeApi.md#getprivileges) | **Get** /privilege | 
*ProgramApi* | [**GetOptedInPrograms**](docs/ProgramApi.md#getoptedinprograms) | **Get** /program/get_opted_in_programs | 
*ProgramApi* | [**OptInToProgram**](docs/ProgramApi.md#optintoprogram) | **Post** /program/opt_in | 
*ProgramApi* | [**OptOutOfProgram**](docs/ProgramApi.md#optoutofprogram) | **Post** /program/opt_out | 
*RateTableApi* | [**GetRateTables**](docs/RateTableApi.md#getratetables) | **Get** /rate_table | 
*ReturnPolicyApi* | [**CreateReturnPolicy**](docs/ReturnPolicyApi.md#createreturnpolicy) | **Post** /return_policy | 
*ReturnPolicyApi* | [**DeleteReturnPolicy**](docs/ReturnPolicyApi.md#deletereturnpolicy) | **Delete** /return_policy/{return_policy_id} | 
*ReturnPolicyApi* | [**GetReturnPolicies**](docs/ReturnPolicyApi.md#getreturnpolicies) | **Get** /return_policy | 
*ReturnPolicyApi* | [**GetReturnPolicy**](docs/ReturnPolicyApi.md#getreturnpolicy) | **Get** /return_policy/{return_policy_id} | 
*ReturnPolicyApi* | [**GetReturnPolicyByName**](docs/ReturnPolicyApi.md#getreturnpolicybyname) | **Get** /return_policy/get_by_policy_name | 
*ReturnPolicyApi* | [**UpdateReturnPolicy**](docs/ReturnPolicyApi.md#updatereturnpolicy) | **Put** /return_policy/{return_policy_id} | 
*SalesTaxApi* | [**CreateOrReplaceSalesTax**](docs/SalesTaxApi.md#createorreplacesalestax) | **Put** /sales_tax/{countryCode}/{jurisdictionId} | 
*SalesTaxApi* | [**DeleteSalesTax**](docs/SalesTaxApi.md#deletesalestax) | **Delete** /sales_tax/{countryCode}/{jurisdictionId} | 
*SalesTaxApi* | [**GetSalesTax**](docs/SalesTaxApi.md#getsalestax) | **Get** /sales_tax/{countryCode}/{jurisdictionId} | 
*SalesTaxApi* | [**GetSalesTaxes**](docs/SalesTaxApi.md#getsalestaxes) | **Get** /sales_tax | 


## Documentation For Models

 - [Amount](docs/Amount.md)
 - [CategoryType](docs/CategoryType.md)
 - [Deposit](docs/Deposit.md)
 - [Error](docs/Error.md)
 - [ErrorParameter](docs/ErrorParameter.md)
 - [FulfillmentPolicy](docs/FulfillmentPolicy.md)
 - [FulfillmentPolicyRequest](docs/FulfillmentPolicyRequest.md)
 - [FulfillmentPolicyResponse](docs/FulfillmentPolicyResponse.md)
 - [InternationalReturnOverrideType](docs/InternationalReturnOverrideType.md)
 - [KycCheck](docs/KycCheck.md)
 - [KycResponse](docs/KycResponse.md)
 - [PaymentMethod](docs/PaymentMethod.md)
 - [PaymentPolicy](docs/PaymentPolicy.md)
 - [PaymentPolicyRequest](docs/PaymentPolicyRequest.md)
 - [PaymentPolicyResponse](docs/PaymentPolicyResponse.md)
 - [PaymentsProgramOnboardingResponse](docs/PaymentsProgramOnboardingResponse.md)
 - [PaymentsProgramOnboardingSteps](docs/PaymentsProgramOnboardingSteps.md)
 - [PaymentsProgramResponse](docs/PaymentsProgramResponse.md)
 - [Program](docs/Program.md)
 - [Programs](docs/Programs.md)
 - [RateTable](docs/RateTable.md)
 - [RateTableResponse](docs/RateTableResponse.md)
 - [RecipientAccountReference](docs/RecipientAccountReference.md)
 - [Region](docs/Region.md)
 - [RegionSet](docs/RegionSet.md)
 - [ReturnPolicy](docs/ReturnPolicy.md)
 - [ReturnPolicyRequest](docs/ReturnPolicyRequest.md)
 - [ReturnPolicyResponse](docs/ReturnPolicyResponse.md)
 - [SalesTax](docs/SalesTax.md)
 - [SalesTaxBase](docs/SalesTaxBase.md)
 - [SalesTaxes](docs/SalesTaxes.md)
 - [SellingLimit](docs/SellingLimit.md)
 - [SellingPrivileges](docs/SellingPrivileges.md)
 - [SetFulfillmentPolicyResponse](docs/SetFulfillmentPolicyResponse.md)
 - [SetPaymentPolicyResponse](docs/SetPaymentPolicyResponse.md)
 - [SetReturnPolicyResponse](docs/SetReturnPolicyResponse.md)
 - [ShippingOption](docs/ShippingOption.md)
 - [ShippingService](docs/ShippingService.md)
 - [TimeDuration](docs/TimeDuration.md)


## Documentation For Authorization



### api_auth


- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://auth.ebay.com/oauth2/authorize
- **Scopes**: 
 - **https://api.ebay.com/oauth/api_scope/sell.account.readonly**: View your account settings
 - **https://api.ebay.com/oauth/api_scope/sell.account**: View and manage your account settings

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author


