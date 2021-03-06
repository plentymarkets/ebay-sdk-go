/*
Account API

The <b>Account API</b> gives sellers the ability to configure their eBay seller accounts, including the seller's policies (seller-defined custom policies and eBay business policies), opt in and out of eBay seller programs, configure sales tax tables, and get account information.  <br><br>For details on the availability of the methods in this API, see <a href=\"/api-docs/sell/account/overview.html#requirements\">Account API requirements and restrictions</a>.

API version: v1.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package account

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

// KycApiService KycApi service
type KycApiService service

type ApiGetKYCRequest struct {
	ctx _context.Context
	ApiService *KycApiService
}


func (r ApiGetKYCRequest) Execute() (KycResponse, *_nethttp.Response, error) {
	return r.ApiService.GetKYCExecute(r)
}

/*
GetKYC Method for GetKYC

This method is used by sellers onboarded for eBay managed payments, or sellers who are currently going through, or who are eligible for onboarding for eBay managed payments. With this method, a seller can discover if there are any action items in regards to providing more documentation and/or information about themselves, their company, or the bank account they are or will be using for seller payouts. These 'action items' are also know as 'Know Your Customer' (or KYC) checks.<br><br>This method does not require any parameters other than the OAuth user token associated with the seller's account.<br><br>If the managed payments seller does not currently have any pending 'KYC' action items, only a <code>204 No Content</code> HTTP status code is returned, and no response payload. <br><br><span class="tablenote"><b>Note</b>: This method is not applicable for sellers who are not eligible for eBay managed payments. For sellers who sell on one or more eBay marketplaces that currently support managed payments, they can check on their managed payments onboarding status by using the <a href="../../payments_program/onboarding/methods/getPaymentsProgramOnboarding">getPaymentsProgramOnboarding</a> method. </span>

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetKYCRequest
*/
func (a *KycApiService) GetKYC(ctx _context.Context) ApiGetKYCRequest {
	return ApiGetKYCRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KycResponse
func (a *KycApiService) GetKYCExecute(r ApiGetKYCRequest) (KycResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  KycResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KycApiService.GetKYC")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/kyc"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
