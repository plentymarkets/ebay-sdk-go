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

// SalesTaxApiService SalesTaxApi service
type SalesTaxApiService service

type ApiCreateOrReplaceSalesTaxRequest struct {
	ctx _context.Context
	ApiService *SalesTaxApiService
	countryCode string
	jurisdictionId string
	salesTaxBase *SalesTaxBase
}

// A container that describes the how the sales tax is calculated.
func (r ApiCreateOrReplaceSalesTaxRequest) SalesTaxBase(salesTaxBase SalesTaxBase) ApiCreateOrReplaceSalesTaxRequest {
	r.salesTaxBase = &salesTaxBase
	return r
}

func (r ApiCreateOrReplaceSalesTaxRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.CreateOrReplaceSalesTaxExecute(r)
}

/*
CreateOrReplaceSalesTax Method for CreateOrReplaceSalesTax

This method creates or updates a sales tax table entry for a jurisdiction. Specify the tax table entry you want to configure using the two path parameters: <b>countryCode</b> and <b>jurisdictionId</b>.  <br><br>A tax table entry for a jurisdiction is comprised of two fields: one for the jurisdiction's sales-tax rate and another that's a boolean value indicating whether or not shipping and handling are taxed in the jurisdiction.  <br><br>You can set up <i>tax tables</i> for countries that support different <i>tax jurisdictions</i>. Currently, only Canada, India, and the US support separate tax jurisdictions. If you sell into any of these countries, you can set up tax tables for any of the country's jurisdictions. Retrieve valid jurisdiction IDs using <b>getSalesTaxJurisdictions</b> in the Metadata API.  <br><br>For details on using this call, see <a href="/api-docs/sell/static/seller-accounts/tax-tables.html">Establishing sales-tax tables</a>. <br><br><span class="tablenote"><b>Important!</b> Starting in January 2019, eBay will begin to calculate, collect, and remit sales tax on behalf of sellers for items shipped to customers. This feature is rolling out on specific dates to specific US states as defined on the following page: <a href="https://www.ebay.com/help/selling/fees-credits-invoices/taxes-import-charges?id=4121#section4" target="_balnk">eBay sales tax collection</a>. <br><br>Once eBay starts to collect sales tax for a state, no action is required on the seller's part and there will be no charges or fees for eBay automatically calculating, collecting and remitting sales tax. The sales-tax collection process will apply to all the sales in the states that support this feature, whether the seller is located in or outside of the United States.  <br><br>When a buyer purchases an item on eBay, and the ship-to address is one of the states where eBay collects the sales tax, eBay will calculate and add the applicable sales tax at checkout. The buyer will pay both the cost of the item along with the sales tax. eBay will collect and remit the tax.</span>

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param countryCode This path parameter specifies the two-letter <a href=\"https://www.iso.org/iso-3166-country-codes.html\" title=\"https://www.iso.org\" target=\"_blank\">ISO 3166</a> code for the country for which you want to create tax table entry.
 @param jurisdictionId This path parameter specifies the ID of the sales-tax jurisdiction for the table entry you want to create.
 @return ApiCreateOrReplaceSalesTaxRequest
*/
func (a *SalesTaxApiService) CreateOrReplaceSalesTax(ctx _context.Context, countryCode string, jurisdictionId string) ApiCreateOrReplaceSalesTaxRequest {
	return ApiCreateOrReplaceSalesTaxRequest{
		ApiService: a,
		ctx: ctx,
		countryCode: countryCode,
		jurisdictionId: jurisdictionId,
	}
}

// Execute executes the request
func (a *SalesTaxApiService) CreateOrReplaceSalesTaxExecute(r ApiCreateOrReplaceSalesTaxRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SalesTaxApiService.CreateOrReplaceSalesTax")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sales_tax/{countryCode}/{jurisdictionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"countryCode"+"}", _neturl.PathEscape(parameterToString(r.countryCode, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"jurisdictionId"+"}", _neturl.PathEscape(parameterToString(r.jurisdictionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.salesTaxBase == nil {
		return nil, reportError("salesTaxBase is required and must be specified")
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
	// body params
	localVarPostBody = r.salesTaxBase
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

type ApiDeleteSalesTaxRequest struct {
	ctx _context.Context
	ApiService *SalesTaxApiService
	countryCode string
	jurisdictionId string
}


func (r ApiDeleteSalesTaxRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteSalesTaxExecute(r)
}

/*
DeleteSalesTax Method for DeleteSalesTax

This call deletes a tax table entry for a jurisdiction. Specify the jurisdiction to delete using the <b>countryCode</b> and <b>jurisdictionId</b> path parameters.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param countryCode This path parameter specifies the two-letter <a href=\"https://www.iso.org/iso-3166-country-codes.html\" title=\"https://www.iso.org\" target=\"_blank\">ISO 3166</a> code for the country whose tax table entry you want to delete.
 @param jurisdictionId This path parameter specifies the ID of the sales tax jurisdiction whose table entry you want to delete.
 @return ApiDeleteSalesTaxRequest
*/
func (a *SalesTaxApiService) DeleteSalesTax(ctx _context.Context, countryCode string, jurisdictionId string) ApiDeleteSalesTaxRequest {
	return ApiDeleteSalesTaxRequest{
		ApiService: a,
		ctx: ctx,
		countryCode: countryCode,
		jurisdictionId: jurisdictionId,
	}
}

// Execute executes the request
func (a *SalesTaxApiService) DeleteSalesTaxExecute(r ApiDeleteSalesTaxRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SalesTaxApiService.DeleteSalesTax")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sales_tax/{countryCode}/{jurisdictionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"countryCode"+"}", _neturl.PathEscape(parameterToString(r.countryCode, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"jurisdictionId"+"}", _neturl.PathEscape(parameterToString(r.jurisdictionId, "")), -1)

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

type ApiGetSalesTaxRequest struct {
	ctx _context.Context
	ApiService *SalesTaxApiService
	countryCode string
	jurisdictionId string
}


func (r ApiGetSalesTaxRequest) Execute() (SalesTax, *_nethttp.Response, error) {
	return r.ApiService.GetSalesTaxExecute(r)
}

/*
GetSalesTax Method for GetSalesTax

This call gets the current tax table entry for a specific tax jurisdiction. Specify the jurisdiction to retrieve using the <b>countryCode</b> and <b>jurisdictionId</b> path parameters.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param countryCode This path parameter specifies the two-letter <a href=\"https://www.iso.org/iso-3166-country-codes.html\" title=\"https://www.iso.org\" target=\"_blank\">ISO 3166</a> code for the country whose tax table you want to retrieve.
 @param jurisdictionId This path parameter specifies the ID of the sales tax jurisdiction for the tax table entry you want to retrieve.
 @return ApiGetSalesTaxRequest
*/
func (a *SalesTaxApiService) GetSalesTax(ctx _context.Context, countryCode string, jurisdictionId string) ApiGetSalesTaxRequest {
	return ApiGetSalesTaxRequest{
		ApiService: a,
		ctx: ctx,
		countryCode: countryCode,
		jurisdictionId: jurisdictionId,
	}
}

// Execute executes the request
//  @return SalesTax
func (a *SalesTaxApiService) GetSalesTaxExecute(r ApiGetSalesTaxRequest) (SalesTax, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  SalesTax
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SalesTaxApiService.GetSalesTax")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sales_tax/{countryCode}/{jurisdictionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"countryCode"+"}", _neturl.PathEscape(parameterToString(r.countryCode, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"jurisdictionId"+"}", _neturl.PathEscape(parameterToString(r.jurisdictionId, "")), -1)

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

type ApiGetSalesTaxesRequest struct {
	ctx _context.Context
	ApiService *SalesTaxApiService
	countryCode *string
}

// This path parameter specifies the two-letter &lt;a href&#x3D;\&quot;https://www.iso.org/iso-3166-country-codes.html\&quot; title&#x3D;\&quot;https://www.iso.org\&quot; target&#x3D;\&quot;_blank\&quot;&gt;ISO 3166&lt;/a&gt; code for the country whose tax table you want to retrieve. For implementation help, refer to eBay API documentation at https://developer.ebay.com/api-docs/sell/account/types/ba:CountryCodeEnum
func (r ApiGetSalesTaxesRequest) CountryCode(countryCode string) ApiGetSalesTaxesRequest {
	r.countryCode = &countryCode
	return r
}

func (r ApiGetSalesTaxesRequest) Execute() (SalesTaxes, *_nethttp.Response, error) {
	return r.ApiService.GetSalesTaxesExecute(r)
}

/*
GetSalesTaxes Method for GetSalesTaxes

Use this call to retrieve a sales tax table that the seller established for a specific country. Specify the tax table to retrieve using the <code>country_code</code> query parameter.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSalesTaxesRequest
*/
func (a *SalesTaxApiService) GetSalesTaxes(ctx _context.Context) ApiGetSalesTaxesRequest {
	return ApiGetSalesTaxesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SalesTaxes
func (a *SalesTaxApiService) GetSalesTaxesExecute(r ApiGetSalesTaxesRequest) (SalesTaxes, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  SalesTaxes
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SalesTaxApiService.GetSalesTaxes")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sales_tax"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.countryCode == nil {
		return localVarReturnValue, nil, reportError("countryCode is required and must be specified")
	}

	localVarQueryParams.Add("country_code", parameterToString(*r.countryCode, ""))
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
