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
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// CustomPolicyApiService CustomPolicyApi service
type CustomPolicyApiService service

type ApiCreateCustomPolicyRequest struct {
	ctx _context.Context
	ApiService *CustomPolicyApiService
	xEBAYCMARKETPLACEID *string
	customPolicyCreateRequest *CustomPolicyCreateRequest
}

// This header parameter specifies the eBay markeplace for the custom policy that is being created. Supported values for this header can be found in the &lt;a href&#x3D;\&quot;/api-docs/sell/compliance/types/bas:MarketplaceIdEnum\&quot; target&#x3D;\&quot;_blank\&quot;&gt;MarketplaceIdEnum&lt;/a&gt; type definition.&lt;br/&gt; &lt;br/&gt; &lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;strong&gt;Note:&lt;/strong&gt; The following eBay marketplaces support Custom Policies: &lt;ul&gt;&lt;li&gt;Germany (EBAY_DE)&lt;/li&gt; &lt;li&gt;Canada (EBAY_CA)&lt;/li&gt; &lt;li&gt;Australia (EBAY_AU)&lt;/li&gt; &lt;li&gt;United States (EBAY_US)&lt;/li&gt; &lt;li&gt;France (EBAY_FR)&lt;/li&gt;&lt;/ul&gt;&lt;/span&gt;
func (r ApiCreateCustomPolicyRequest) XEBAYCMARKETPLACEID(xEBAYCMARKETPLACEID string) ApiCreateCustomPolicyRequest {
	r.xEBAYCMARKETPLACEID = &xEBAYCMARKETPLACEID
	return r
}
// Request to create a new Custom Policy.
func (r ApiCreateCustomPolicyRequest) CustomPolicyCreateRequest(customPolicyCreateRequest CustomPolicyCreateRequest) ApiCreateCustomPolicyRequest {
	r.customPolicyCreateRequest = &customPolicyCreateRequest
	return r
}

func (r ApiCreateCustomPolicyRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.CreateCustomPolicyExecute(r)
}

/*
CreateCustomPolicy Method for CreateCustomPolicy

This method creates a new custom policy in which a seller specifies their terms for complying with local governmental regulations. <br/><br/>Two Custom Policy types are supported: <ul><li>Product Compliance (PRODUCT_COMPLIANCE)</li> <li>Takeback (TAKE_BACK)</li></ul>Each Custom Policy targets a <b>policyType</b> and <b>eBay marketplace</b> combination. Multiple policies may be created as follows: <ul><li><b>Product Compliance</b>: a maximum of 10 policies per eBay marketplace may be created</li> <li><b>Takeback</b>: a maximum of 3 policies per eBay marketplace may be created</li></ul>A successful create policy call returns an HTTP status code of <b>201 Created</b> with the system-generated policy ID included in the <b>Location</b> response header.<br/><br/><b>Product Compliance Policy</b><br/><br/>Product Compliance policies disclose product information as required for regulatory compliance.<br/><br/><span class="tablenote"><strong>Note:</strong> A maximum of 10 Product Compliance policies per eBay marketplace may be created.</span> <br/><br/> <b>Takeback Policy</b><br/><br/>Takeback policies describe the seller's legal obligation to take back a previously purchased item when the buyer purchases a new one.<br/><br/><span class="tablenote"><strong>Note:</strong> A maximum of 3 Takeback policies per eBay marketplace may be created.</span>

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateCustomPolicyRequest
*/
func (a *CustomPolicyApiService) CreateCustomPolicy(ctx _context.Context) ApiCreateCustomPolicyRequest {
	return ApiCreateCustomPolicyRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *CustomPolicyApiService) CreateCustomPolicyExecute(r ApiCreateCustomPolicyRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomPolicyApiService.CreateCustomPolicy")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/custom_policy/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xEBAYCMARKETPLACEID == nil {
		return localVarReturnValue, nil, reportError("xEBAYCMARKETPLACEID is required and must be specified")
	}
	if r.customPolicyCreateRequest == nil {
		return localVarReturnValue, nil, reportError("customPolicyCreateRequest is required and must be specified")
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
	localVarHeaderParams["X-EBAY-C-MARKETPLACE-ID"] = parameterToString(*r.xEBAYCMARKETPLACEID, "")
	// body params
	localVarPostBody = r.customPolicyCreateRequest
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

type ApiGetCustomPoliciesRequest struct {
	ctx _context.Context
	ApiService *CustomPolicyApiService
	xEBAYCMARKETPLACEID *string
	policyTypes *string
}

// This header parameter specifies the eBay markeplace for the custom policy that is being created. Supported values for this header can be found in the &lt;a href&#x3D;\&quot;/api-docs/sell/compliance/types/bas:MarketplaceIdEnum\&quot; target&#x3D;\&quot;_blank\&quot;&gt;MarketplaceIdEnum&lt;/a&gt; type definition.&lt;br/&gt; &lt;br/&gt; &lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;strong&gt;Note:&lt;/strong&gt; The following eBay marketplaces support Custom Policies: &lt;ul&gt;&lt;li&gt;Germany (EBAY_DE)&lt;/li&gt; &lt;li&gt;Canada (EBAY_CA)&lt;/li&gt; &lt;li&gt;Australia (EBAY_AU)&lt;/li&gt; &lt;li&gt;United States (EBAY_US)&lt;/li&gt; &lt;li&gt;France (EBAY_FR)&lt;/li&gt;&lt;/ul&gt;&lt;/span&gt;
func (r ApiGetCustomPoliciesRequest) XEBAYCMARKETPLACEID(xEBAYCMARKETPLACEID string) ApiGetCustomPoliciesRequest {
	r.xEBAYCMARKETPLACEID = &xEBAYCMARKETPLACEID
	return r
}
// This query parameter specifies the type of custom policies to be returned.&lt;br /&gt;&lt;br /&gt;Multiple policy types may be requested in a single call by providing a comma-delimited set of all policy types to be returned.&lt;br/&gt;&lt;br/&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;strong&gt;Note:&lt;/strong&gt; Omitting this query parameter from a request will also return policies of all policy types.&lt;/span&gt;&lt;br/&gt;&lt;br/&gt;Two Custom Policy types are supported: &lt;ul&gt;&lt;li&gt;Product Compliance (PRODUCT_COMPLIANCE)&lt;/li&gt; &lt;li&gt;Takeback (TAKE_BACK)&lt;/li&gt;&lt;/ul&gt;
func (r ApiGetCustomPoliciesRequest) PolicyTypes(policyTypes string) ApiGetCustomPoliciesRequest {
	r.policyTypes = &policyTypes
	return r
}

func (r ApiGetCustomPoliciesRequest) Execute() (CustomPolicyResponse, *_nethttp.Response, error) {
	return r.ApiService.GetCustomPoliciesExecute(r)
}

/*
GetCustomPolicies Method for GetCustomPolicies

This method retrieves the list of custom policies specified by the <b>policy_types</b> query parameter for the selected eBay marketplace.<br/> <br/> <span class="tablenote"><strong>Note:</strong> The following eBay marketplaces support Custom Policies: <ul><li>Germany (EBAY_DE)</li> <li>Canada (EBAY_CA)</li> <li>Australia (EBAY_AU)</li> <li>United States (EBAY_US)</li> <li>France (EBAY_FR)</li></ul></span><br/><br/>For details on header values, see <a href="/api-docs/static/rest-request-components.html#HTTP" target="_blank">HTTP request headers</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetCustomPoliciesRequest
*/
func (a *CustomPolicyApiService) GetCustomPolicies(ctx _context.Context) ApiGetCustomPoliciesRequest {
	return ApiGetCustomPoliciesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CustomPolicyResponse
func (a *CustomPolicyApiService) GetCustomPoliciesExecute(r ApiGetCustomPoliciesRequest) (CustomPolicyResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  CustomPolicyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomPolicyApiService.GetCustomPolicies")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/custom_policy/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xEBAYCMARKETPLACEID == nil {
		return localVarReturnValue, nil, reportError("xEBAYCMARKETPLACEID is required and must be specified")
	}

	if r.policyTypes != nil {
		localVarQueryParams.Add("policy_types", parameterToString(*r.policyTypes, ""))
	}
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
	localVarHeaderParams["X-EBAY-C-MARKETPLACE-ID"] = parameterToString(*r.xEBAYCMARKETPLACEID, "")
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

type ApiGetCustomPolicyRequest struct {
	ctx _context.Context
	ApiService *CustomPolicyApiService
	customPolicyId string
	xEBAYCMARKETPLACEID *string
}

// This header parameter specifies the eBay markeplace for the custom policy that is being created. Supported values for this header can be found in the &lt;a href&#x3D;\&quot;/api-docs/sell/compliance/types/bas:MarketplaceIdEnum\&quot; target&#x3D;\&quot;_blank\&quot;&gt;MarketplaceIdEnum&lt;/a&gt; type definition.&lt;br/&gt; &lt;br/&gt; &lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;strong&gt;Note:&lt;/strong&gt; The following eBay marketplaces support Custom Policies: &lt;ul&gt;&lt;li&gt;Germany (EBAY_DE)&lt;/li&gt; &lt;li&gt;Canada (EBAY_CA)&lt;/li&gt; &lt;li&gt;Australia (EBAY_AU)&lt;/li&gt; &lt;li&gt;United States (EBAY_US)&lt;/li&gt; &lt;li&gt;France (EBAY_FR)&lt;/li&gt;&lt;/ul&gt;&lt;/span&gt;
func (r ApiGetCustomPolicyRequest) XEBAYCMARKETPLACEID(xEBAYCMARKETPLACEID string) ApiGetCustomPolicyRequest {
	r.xEBAYCMARKETPLACEID = &xEBAYCMARKETPLACEID
	return r
}

func (r ApiGetCustomPolicyRequest) Execute() (CustomPolicy, *_nethttp.Response, error) {
	return r.ApiService.GetCustomPolicyExecute(r)
}

/*
GetCustomPolicy Method for GetCustomPolicy

This method retrieves the custom policy specified by the <b>custom_policy_id</b> path parameter for the selected eBay marketplace.<br/> <br/> <span class="tablenote"><strong>Note:</strong> The following eBay marketplaces support Custom Policies: <ul><li>Germany (EBAY_DE)</li> <li>Canada (EBAY_CA)</li> <li>Australia (EBAY_AU)</li> <li>United States (EBAY_US)</li> <li>France (EBAY_FR)</li></ul></span><br/><br/>For details on header values, see <a href="/api-docs/static/rest-request-components.html#HTTP" target="_blank">HTTP request headers</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customPolicyId This path parameter is the unique custom policy identifier for the policy to be returned.<br/><br/><span class=\"tablenote\"><strong>Note:</strong> This value is automatically assigned by the system when the policy is created.</span>
 @return ApiGetCustomPolicyRequest
*/
func (a *CustomPolicyApiService) GetCustomPolicy(ctx _context.Context, customPolicyId string) ApiGetCustomPolicyRequest {
	return ApiGetCustomPolicyRequest{
		ApiService: a,
		ctx: ctx,
		customPolicyId: customPolicyId,
	}
}

// Execute executes the request
//  @return CustomPolicy
func (a *CustomPolicyApiService) GetCustomPolicyExecute(r ApiGetCustomPolicyRequest) (CustomPolicy, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  CustomPolicy
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomPolicyApiService.GetCustomPolicy")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/custom_policy/{custom_policy_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"custom_policy_id"+"}", _neturl.PathEscape(parameterToString(r.customPolicyId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xEBAYCMARKETPLACEID == nil {
		return localVarReturnValue, nil, reportError("xEBAYCMARKETPLACEID is required and must be specified")
	}

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
	localVarHeaderParams["X-EBAY-C-MARKETPLACE-ID"] = parameterToString(*r.xEBAYCMARKETPLACEID, "")
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

type ApiUpdateCustomPolicyRequest struct {
	ctx _context.Context
	ApiService *CustomPolicyApiService
	customPolicyId string
	xEBAYCMARKETPLACEID *string
	customPolicyRequest *CustomPolicyRequest
}

// This header parameter specifies the eBay markeplace for the custom policy that is being created. Supported values for this header can be found in the &lt;a href&#x3D;\&quot;/api-docs/sell/compliance/types/bas:MarketplaceIdEnum\&quot; target&#x3D;\&quot;_blank\&quot;&gt;MarketplaceIdEnum&lt;/a&gt; type definition.&lt;br/&gt; &lt;br/&gt; &lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;strong&gt;Note:&lt;/strong&gt; The following eBay marketplaces support Custom Policies: &lt;ul&gt;&lt;li&gt;Germany (EBAY_DE)&lt;/li&gt; &lt;li&gt;Canada (EBAY_CA)&lt;/li&gt; &lt;li&gt;Australia (EBAY_AU)&lt;/li&gt; &lt;li&gt;United States (EBAY_US)&lt;/li&gt; &lt;li&gt;France (EBAY_FR)&lt;/li&gt;&lt;/ul&gt;&lt;/span&gt;
func (r ApiUpdateCustomPolicyRequest) XEBAYCMARKETPLACEID(xEBAYCMARKETPLACEID string) ApiUpdateCustomPolicyRequest {
	r.xEBAYCMARKETPLACEID = &xEBAYCMARKETPLACEID
	return r
}
// Request to update a current custom policy.
func (r ApiUpdateCustomPolicyRequest) CustomPolicyRequest(customPolicyRequest CustomPolicyRequest) ApiUpdateCustomPolicyRequest {
	r.customPolicyRequest = &customPolicyRequest
	return r
}

func (r ApiUpdateCustomPolicyRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.UpdateCustomPolicyExecute(r)
}

/*
UpdateCustomPolicy Method for UpdateCustomPolicy

This method updates an existing custom policy specified by the <b>custom_policy_id</b> path parameter for the selected marketplace. This method overwrites the policy's <b>Name</b>, <b>Label</b>, and <b>Description</b> fields. Therefore, the complete, current text of all three policy fields must be included in the request payload even when one or two of these fields will not actually be updated.<br/> <br/>For example, the value for the <b>Label</b> field is to be updated, but the <b>Name</b> and <b>Description</b> values will remain unchanged. The existing <b>Name</b> and <b>Description</b> values, as they are defined in the current policy, must also be passed in. <br/><br/>A successful policy update call returns an HTTP status code of <b>204 No Content</b>.<br/><br/><span class="tablenote"><strong>Note:</strong> The following eBay marketplaces support Custom Policies: <ul><li>Germany (EBAY_DE)</li> <li>Canada (EBAY_CA)</li> <li>Australia (EBAY_AU)</li> <li>United States (EBAY_US)</li> <li>France (EBAY_FR)</li></ul></span><br/><br/>For details on header values, see <a href="/api-docs/static/rest-request-components.html#HTTP">HTTP request headers</a>.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customPolicyId This path parameter is the unique custom policy identifier for the policy to be returned.<br/><br/><span class=\"tablenote\"><strong>Note:</strong> This value is automatically assigned by the system when the policy is created.</span>
 @return ApiUpdateCustomPolicyRequest
*/
func (a *CustomPolicyApiService) UpdateCustomPolicy(ctx _context.Context, customPolicyId string) ApiUpdateCustomPolicyRequest {
	return ApiUpdateCustomPolicyRequest{
		ApiService: a,
		ctx: ctx,
		customPolicyId: customPolicyId,
	}
}

// Execute executes the request
func (a *CustomPolicyApiService) UpdateCustomPolicyExecute(r ApiUpdateCustomPolicyRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomPolicyApiService.UpdateCustomPolicy")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/custom_policy/{custom_policy_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"custom_policy_id"+"}", _neturl.PathEscape(parameterToString(r.customPolicyId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xEBAYCMARKETPLACEID == nil {
		return nil, reportError("xEBAYCMARKETPLACEID is required and must be specified")
	}
	if r.customPolicyRequest == nil {
		return nil, reportError("customPolicyRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	localVarHeaderParams["X-EBAY-C-MARKETPLACE-ID"] = parameterToString(*r.xEBAYCMARKETPLACEID, "")
	// body params
	localVarPostBody = r.customPolicyRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
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
