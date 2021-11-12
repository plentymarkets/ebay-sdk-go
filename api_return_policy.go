/*
Account API

The <b>Account API</b> gives sellers the ability to configure their eBay seller accounts, including the seller's policies (the Fulfillment Policy, Payment Policy, and Return Policy), opt in and out of eBay seller programs, configure sales tax tables, and get account information.  <br><br>For details on the availability of the methods in this API, see <a href=\"/api-docs/sell/account/overview.html#requirements\">Account API requirements and restrictions</a>.

API version: v1.6.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ebayaccount

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// ReturnPolicyApiService ReturnPolicyApi service
type ReturnPolicyApiService service

type ApiCreateReturnPolicyRequest struct {
	ctx _context.Context
	ApiService *ReturnPolicyApiService
	returnPolicyRequest *ReturnPolicyRequest
}

// Return policy request
func (r ApiCreateReturnPolicyRequest) ReturnPolicyRequest(returnPolicyRequest ReturnPolicyRequest) ApiCreateReturnPolicyRequest {
	r.returnPolicyRequest = &returnPolicyRequest
	return r
}

func (r ApiCreateReturnPolicyRequest) Execute() (SetReturnPolicyResponse, *_nethttp.Response, error) {
	return r.ApiService.CreateReturnPolicyExecute(r)
}

/*
CreateReturnPolicy Method for CreateReturnPolicy

This method creates a new return policy where the policy encapsulates seller's terms for returning items. Use the Metadata API method <b>getReturnPolicies</b> to determine which categories require you to supply a return policy for the marketplace(s) into which you list.  <br><br>Each policy targets a <b>marketplaceId</b> and <code>categoryTypes.</code><b>name</b> combination and you can create multiple policies for each combination.  <br><br>A successful request returns the URI to the new policy in the <b>Location</b> response header and the ID for the new policy is returned in the response payload.  <p class="tablenote"><b>Tip:</b> For details on creating and using the business policies supported by the Account API, see <a href="/api-docs/sell/static/seller-accounts/business-policies.html">eBay business policies</a>.</p>  <p><b>Marketplaces and locales</b></p>  <p>Policy instructions can be localized by providing a locale in the <code>Accept-Language</code> HTTP request header. For example, the following setting displays field values from the request body in German: <code>Accept-Language: de-DE</code>.</p>  <p>Target the specific locale of a marketplace that supports multiple locales using the <code>Content-Language</code> request header. For example, target the French locale of the Canadian marketplace by specifying the <code>fr-CA</code> locale for <code>Content-Language</code>. Likewise, target the Dutch locale of the Belgium marketplace by setting <code>Content-Language: nl-BE</code>.</p> <p class="tablenote"><b>Tip:</b> For details on headers, see <a href="/api-docs/static/rest-request-components.html#HTTP">HTTP request headers</a>.</p>

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateReturnPolicyRequest
*/
func (a *ReturnPolicyApiService) CreateReturnPolicy(ctx _context.Context) ApiCreateReturnPolicyRequest {
	return ApiCreateReturnPolicyRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SetReturnPolicyResponse
func (a *ReturnPolicyApiService) CreateReturnPolicyExecute(r ApiCreateReturnPolicyRequest) (SetReturnPolicyResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SetReturnPolicyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReturnPolicyApiService.CreateReturnPolicy")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/return_policy"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.returnPolicyRequest == nil {
		return localVarReturnValue, nil, reportError("returnPolicyRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.returnPolicyRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
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

type ApiDeleteReturnPolicyRequest struct {
	ctx _context.Context
	ApiService *ReturnPolicyApiService
	returnPolicyId string
}


func (r ApiDeleteReturnPolicyRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteReturnPolicyExecute(r)
}

/*
DeleteReturnPolicy Method for DeleteReturnPolicy

This method deletes a return policy. Supply the ID of the policy you want to delete in the <b>returnPolicyId</b> path parameter. Note that you cannot delete the default return policy.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param returnPolicyId This path parameter specifies the ID of the return policy you want to delete.
 @return ApiDeleteReturnPolicyRequest
*/
func (a *ReturnPolicyApiService) DeleteReturnPolicy(ctx _context.Context, returnPolicyId string) ApiDeleteReturnPolicyRequest {
	return ApiDeleteReturnPolicyRequest{
		ApiService: a,
		ctx: ctx,
		returnPolicyId: returnPolicyId,
	}
}

// Execute executes the request
func (a *ReturnPolicyApiService) DeleteReturnPolicyExecute(r ApiDeleteReturnPolicyRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReturnPolicyApiService.DeleteReturnPolicy")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/return_policy/{return_policy_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"return_policy_id"+"}", _neturl.PathEscape(parameterToString(r.returnPolicyId, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetReturnPoliciesRequest struct {
	ctx _context.Context
	ApiService *ReturnPolicyApiService
	marketplaceId *string
}

// This query parameter specifies the ID of the eBay marketplace of the policy you want to retrieve. For implementation help, refer to eBay API documentation at https://developer.ebay.com/api-docs/sell/account/types/ba:MarketplaceIdEnum
func (r ApiGetReturnPoliciesRequest) MarketplaceId(marketplaceId string) ApiGetReturnPoliciesRequest {
	r.marketplaceId = &marketplaceId
	return r
}

func (r ApiGetReturnPoliciesRequest) Execute() (ReturnPolicyResponse, *_nethttp.Response, error) {
	return r.ApiService.GetReturnPoliciesExecute(r)
}

/*
GetReturnPolicies Method for GetReturnPolicies

This method retrieves all the return policies configured for the marketplace you specify using the <code>marketplace_id</code> query parameter.  <br><br><b>Marketplaces and locales</b>  <br><br>Get the correct policies for a marketplace that supports multiple locales using the <code>Content-Language</code> request header. For example, get the policies for the French locale of the Canadian marketplace by specifying <code>fr-CA</code> for the <code>Content-Language</code> header. Likewise, target the Dutch locale of the Belgium marketplace by setting <code>Content-Language: nl-BE</code>. For details on header values, see <a href="/api-docs/static/rest-request-components.html#HTTP" target="_blank">HTTP request headers</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetReturnPoliciesRequest
*/
func (a *ReturnPolicyApiService) GetReturnPolicies(ctx _context.Context) ApiGetReturnPoliciesRequest {
	return ApiGetReturnPoliciesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ReturnPolicyResponse
func (a *ReturnPolicyApiService) GetReturnPoliciesExecute(r ApiGetReturnPoliciesRequest) (ReturnPolicyResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ReturnPolicyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReturnPolicyApiService.GetReturnPolicies")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/return_policy"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.marketplaceId == nil {
		return localVarReturnValue, nil, reportError("marketplaceId is required and must be specified")
	}

	localVarQueryParams.Add("marketplace_id", parameterToString(*r.marketplaceId, ""))
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
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

type ApiGetReturnPolicyRequest struct {
	ctx _context.Context
	ApiService *ReturnPolicyApiService
	returnPolicyId string
}


func (r ApiGetReturnPolicyRequest) Execute() (ReturnPolicy, *_nethttp.Response, error) {
	return r.ApiService.GetReturnPolicyExecute(r)
}

/*
GetReturnPolicy Method for GetReturnPolicy

This method retrieves the complete details of the return policy specified by the <b>returnPolicyId</b> path parameter.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param returnPolicyId This path parameter specifies the of the return policy you want to retrieve.
 @return ApiGetReturnPolicyRequest
*/
func (a *ReturnPolicyApiService) GetReturnPolicy(ctx _context.Context, returnPolicyId string) ApiGetReturnPolicyRequest {
	return ApiGetReturnPolicyRequest{
		ApiService: a,
		ctx: ctx,
		returnPolicyId: returnPolicyId,
	}
}

// Execute executes the request
//  @return ReturnPolicy
func (a *ReturnPolicyApiService) GetReturnPolicyExecute(r ApiGetReturnPolicyRequest) (ReturnPolicy, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ReturnPolicy
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReturnPolicyApiService.GetReturnPolicy")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/return_policy/{return_policy_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"return_policy_id"+"}", _neturl.PathEscape(parameterToString(r.returnPolicyId, "")), -1)

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
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

type ApiGetReturnPolicyByNameRequest struct {
	ctx _context.Context
	ApiService *ReturnPolicyApiService
	marketplaceId *string
	name *string
}

// This query parameter specifies the ID of the eBay marketplace of the policy you want to retrieve. For implementation help, refer to eBay API documentation at https://developer.ebay.com/api-docs/sell/account/types/ba:MarketplaceIdEnum
func (r ApiGetReturnPolicyByNameRequest) MarketplaceId(marketplaceId string) ApiGetReturnPolicyByNameRequest {
	r.marketplaceId = &marketplaceId
	return r
}
// This query parameter specifies the user-defined name of the return policy you want to retrieve.
func (r ApiGetReturnPolicyByNameRequest) Name(name string) ApiGetReturnPolicyByNameRequest {
	r.name = &name
	return r
}

func (r ApiGetReturnPolicyByNameRequest) Execute() (ReturnPolicy, *_nethttp.Response, error) {
	return r.ApiService.GetReturnPolicyByNameExecute(r)
}

/*
GetReturnPolicyByName Method for GetReturnPolicyByName

This method retrieves the complete details of a single return policy. Supply both the policy <code>name</code> and its associated <code>marketplace_id</code> in the request query parameters.   <br><br><b>Marketplaces and locales</b>  <br><br>Get the correct policy for a marketplace that supports multiple locales using the <code>Content-Language</code> request header. For example, get a policy for the French locale of the Canadian marketplace by specifying <code>fr-CA</code> for the <code>Content-Language</code> header. Likewise, target the Dutch locale of the Belgium marketplace by setting <code>Content-Language: nl-BE</code>. For details on header values, see <a href="/api-docs/static/rest-request-components.html#HTTP">HTTP request headers</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetReturnPolicyByNameRequest
*/
func (a *ReturnPolicyApiService) GetReturnPolicyByName(ctx _context.Context) ApiGetReturnPolicyByNameRequest {
	return ApiGetReturnPolicyByNameRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ReturnPolicy
func (a *ReturnPolicyApiService) GetReturnPolicyByNameExecute(r ApiGetReturnPolicyByNameRequest) (ReturnPolicy, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ReturnPolicy
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReturnPolicyApiService.GetReturnPolicyByName")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/return_policy/get_by_policy_name"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.marketplaceId == nil {
		return localVarReturnValue, nil, reportError("marketplaceId is required and must be specified")
	}
	if r.name == nil {
		return localVarReturnValue, nil, reportError("name is required and must be specified")
	}

	localVarQueryParams.Add("marketplace_id", parameterToString(*r.marketplaceId, ""))
	localVarQueryParams.Add("name", parameterToString(*r.name, ""))
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
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

type ApiUpdateReturnPolicyRequest struct {
	ctx _context.Context
	ApiService *ReturnPolicyApiService
	returnPolicyId string
	returnPolicyRequest *ReturnPolicyRequest
}

// Container for a return policy request.
func (r ApiUpdateReturnPolicyRequest) ReturnPolicyRequest(returnPolicyRequest ReturnPolicyRequest) ApiUpdateReturnPolicyRequest {
	r.returnPolicyRequest = &returnPolicyRequest
	return r
}

func (r ApiUpdateReturnPolicyRequest) Execute() (SetReturnPolicyResponse, *_nethttp.Response, error) {
	return r.ApiService.UpdateReturnPolicyExecute(r)
}

/*
UpdateReturnPolicy Method for UpdateReturnPolicy

This method updates an existing return policy. Specify the policy you want to update using the <b>return_policy_id</b> path parameter. Supply a complete policy payload with the updates you want to make; this call overwrites the existing policy with the new details specified in the payload.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param returnPolicyId This path parameter specifies the ID of the return policy you want to update.
 @return ApiUpdateReturnPolicyRequest
*/
func (a *ReturnPolicyApiService) UpdateReturnPolicy(ctx _context.Context, returnPolicyId string) ApiUpdateReturnPolicyRequest {
	return ApiUpdateReturnPolicyRequest{
		ApiService: a,
		ctx: ctx,
		returnPolicyId: returnPolicyId,
	}
}

// Execute executes the request
//  @return SetReturnPolicyResponse
func (a *ReturnPolicyApiService) UpdateReturnPolicyExecute(r ApiUpdateReturnPolicyRequest) (SetReturnPolicyResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SetReturnPolicyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReturnPolicyApiService.UpdateReturnPolicy")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/return_policy/{return_policy_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"return_policy_id"+"}", _neturl.PathEscape(parameterToString(r.returnPolicyId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.returnPolicyRequest == nil {
		return localVarReturnValue, nil, reportError("returnPolicyRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.returnPolicyRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
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