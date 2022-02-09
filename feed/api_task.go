/*
Feed API

<p>The <strong>Feed API</strong> lets sellers upload input files, download reports and files including their status, filter reports using URI parameters, and retrieve customer service metrics task details.</p>

API version: v1.3.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package feed

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

// TaskApiService TaskApi service
type TaskApiService service

type ApiCreateTaskRequest struct {
	ctx _context.Context
	ApiService *TaskApiService
	createTaskRequest *CreateTaskRequest
	xEBAYCMARKETPLACEID *string
}

// description not needed
func (r ApiCreateTaskRequest) CreateTaskRequest(createTaskRequest CreateTaskRequest) ApiCreateTaskRequest {
	r.createTaskRequest = &createTaskRequest
	return r
}
// The ID of the eBay marketplace where the item is hosted. Note: This value is case sensitive. For example: X-EBAY-C-MARKETPLACE-ID:EBAY_US This identifies the eBay marketplace that applies to this task. See MarketplaceIdEnum.
func (r ApiCreateTaskRequest) XEBAYCMARKETPLACEID(xEBAYCMARKETPLACEID string) ApiCreateTaskRequest {
	r.xEBAYCMARKETPLACEID = &xEBAYCMARKETPLACEID
	return r
}

func (r ApiCreateTaskRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.CreateTaskExecute(r)
}

/*
CreateTask Method for CreateTask

This method creates an upload task or a download task without filter criteria. When using this method, specify the feedType and the feed file schemaVersion. The feed type specified sets the task as a download or an upload task. For details about the upload and download flows, see Working with Order Feeds in the Selling Integration Guide. Note: The scope depends on the feed type. An error message results when an unsupported scope or feed type is specified. The following list contains this method's authorization scopes and their corresponding feed types: https://api.ebay.com/oauth/api_scope/sell.inventory: See LMS FeedTypes https://api.ebay.com/oauth/api_scope/sell.fulfillment: LMS_ORDER_ACK (specify for upload tasks). Also see LMS FeedTypes https://api.ebay.com/oauth/api_scope/sell.marketing: None* https://api.ebay.com/oauth/api_scope/commerce.catalog.readonly: None* * Reserved for future release

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateTaskRequest
*/
func (a *TaskApiService) CreateTask(ctx _context.Context) ApiCreateTaskRequest {
	return ApiCreateTaskRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *TaskApiService) CreateTaskExecute(r ApiCreateTaskRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskApiService.CreateTask")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/task"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.createTaskRequest == nil {
		return nil, reportError("createTaskRequest is required and must be specified")
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
	if r.xEBAYCMARKETPLACEID != nil {
		localVarHeaderParams["X-EBAY-C-MARKETPLACE-ID"] = parameterToString(*r.xEBAYCMARKETPLACEID, "")
	}
	// body params
	localVarPostBody = r.createTaskRequest
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

type ApiGetInputFileRequest struct {
	ctx _context.Context
	ApiService *TaskApiService
	taskId string
}


func (r ApiGetInputFileRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetInputFileExecute(r)
}

/*
GetInputFile Method for GetInputFile

This method downloads the file previously uploaded using uploadFile. Specify the task_id from the uploadFile call. Note: With respect to LMS, this method applies to all feed types except LMS_ORDER_REPORT and LMS_ACTIVE_INVENTORY_REPORT. See LMS API Feeds in the Selling Integration Guide.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param taskId The task ID associated with the file to be downloaded.
 @return ApiGetInputFileRequest
*/
func (a *TaskApiService) GetInputFile(ctx _context.Context, taskId string) ApiGetInputFileRequest {
	return ApiGetInputFileRequest{
		ApiService: a,
		ctx: ctx,
		taskId: taskId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *TaskApiService) GetInputFileExecute(r ApiGetInputFileRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskApiService.GetInputFile")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/task/{task_id}/download_input_file"
	localVarPath = strings.Replace(localVarPath, "{"+"task_id"+"}", _neturl.PathEscape(parameterToString(r.taskId, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/octet-stream"}

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

type ApiGetResultFileRequest struct {
	ctx _context.Context
	ApiService *TaskApiService
	taskId string
}


func (r ApiGetResultFileRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetResultFileExecute(r)
}

/*
GetResultFile Method for GetResultFile

This method retrieves the generated file that is associated with the specified task ID. The response of this call is a compressed or uncompressed CSV, XML, or JSON file, with the applicable file extension (for example: csv.gz). For details about how this method is used, see Working with Order Feeds in the Selling Integration Guide. Note: The status of the task to retrieve must be in the COMPLETED or COMPLETED_WITH_ERROR state before this method can retrieve the file. You can use the getTask or getTasks method to retrieve the status of the task.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param taskId The ID of the task associated with the file you want to download. This ID was generated when the task was created.
 @return ApiGetResultFileRequest
*/
func (a *TaskApiService) GetResultFile(ctx _context.Context, taskId string) ApiGetResultFileRequest {
	return ApiGetResultFileRequest{
		ApiService: a,
		ctx: ctx,
		taskId: taskId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *TaskApiService) GetResultFileExecute(r ApiGetResultFileRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskApiService.GetResultFile")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/task/{task_id}/download_result_file"
	localVarPath = strings.Replace(localVarPath, "{"+"task_id"+"}", _neturl.PathEscape(parameterToString(r.taskId, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/octet-stream"}

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

type ApiGetTaskRequest struct {
	ctx _context.Context
	ApiService *TaskApiService
	taskId string
}


func (r ApiGetTaskRequest) Execute() (Task, *_nethttp.Response, error) {
	return r.ApiService.GetTaskExecute(r)
}

/*
GetTask Method for GetTask

This method retrieves the details and status of the specified task. The input is task_id. For details of how this method is used, see Working with Order Feeds in the Selling Integration Guide.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param taskId The ID of the task. This ID was generated when the task was created.
 @return ApiGetTaskRequest
*/
func (a *TaskApiService) GetTask(ctx _context.Context, taskId string) ApiGetTaskRequest {
	return ApiGetTaskRequest{
		ApiService: a,
		ctx: ctx,
		taskId: taskId,
	}
}

// Execute executes the request
//  @return Task
func (a *TaskApiService) GetTaskExecute(r ApiGetTaskRequest) (Task, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Task
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskApiService.GetTask")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/task/{task_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"task_id"+"}", _neturl.PathEscape(parameterToString(r.taskId, "")), -1)

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

type ApiGetTasksRequest struct {
	ctx _context.Context
	ApiService *TaskApiService
	dateRange *string
	feedType *string
	limit *string
	lookBackDays *string
	offset *string
	scheduleId *string
}

// Specifies the range of task creation dates used to filter the results. The results are filtered to include only tasks with a creation date that is equal to this date or is within specified range. Only tasks that are less than 90 days can be retrieved. Note: Maximum date range window size is 90 days. Valid Format (UTC):yyyy-MM-ddThh:mm:ss.SSSZ..yyyy-MM-ddThh:mm:ss.SSSZ For example: Tasks created on September 8, 2019 2019-09-08T00:00:00.000Z..2019-09-09T00:00:00.000Z
func (r ApiGetTasksRequest) DateRange(dateRange string) ApiGetTasksRequest {
	r.dateRange = &dateRange
	return r
}
// The feed type associated with the tasks to be returned. Only use a feedType that is available for your API: Order Feeds: LMS_ORDER_ACK, LMS_ORDER_REPORT Large Merchant Services (LMS) Feeds: See Available FeedTypes Do not use with the schedule_id parameter. Since schedules are based on feed types, you can specify a schedule (schedule_id) that returns the needed feed_type.
func (r ApiGetTasksRequest) FeedType(feedType string) ApiGetTasksRequest {
	r.feedType = &feedType
	return r
}
// The maximum number of tasks that can be returned on each page of the paginated response. Use this parameter in conjunction with the offset parameter to control the pagination of the output. Note: This feature employs a zero-based list, where the first item in the list has an offset of 0. For example, if offset is set to 10 and limit is set to 10, the call retrieves tasks 11 thru 20 from the result set. If this parameter is omitted, the default value is used. Default: 10 Maximum: 500
func (r ApiGetTasksRequest) Limit(limit string) ApiGetTasksRequest {
	r.limit = &limit
	return r
}
// The number of previous days in which to search for tasks. Do not use with the date_range parameter. If both date_range and look_back_days are omitted, this parameter&#39;s default value is used. Default: 7 Range: 1-90 (inclusive)
func (r ApiGetTasksRequest) LookBackDays(lookBackDays string) ApiGetTasksRequest {
	r.lookBackDays = &lookBackDays
	return r
}
// The number of tasks to skip in the result set before returning the first task in the paginated response. Combine offset with the limit query parameter to control the items returned in the response. For example, if you supply an offset of 0 and a limit of 10, the first page of the response contains the first 10 items from the complete list of items retrieved by the call. If offset is 10 and limit is 20, the first page of the response contains items 11-30 from the complete result set. If this query parameter is not set, the default value is used and the first page of records is returned. Default: 0
func (r ApiGetTasksRequest) Offset(offset string) ApiGetTasksRequest {
	r.offset = &offset
	return r
}
// The schedule ID associated with the task. A schedule periodically generates a report for the feed type specified by the schedule template (see scheduleTemplateId in createSchedule). Do not use with the feed_type parameter. Since schedules are based on feed types, you can specify a schedule (schedule_id) that returns the needed feed_type.
func (r ApiGetTasksRequest) ScheduleId(scheduleId string) ApiGetTasksRequest {
	r.scheduleId = &scheduleId
	return r
}

func (r ApiGetTasksRequest) Execute() (TaskCollection, *_nethttp.Response, error) {
	return r.ApiService.GetTasksExecute(r)
}

/*
GetTasks Method for GetTasks

This method returns the details and status for an array of tasks based on a specified feed_type or scheduledId. Specifying both feed_type and scheduledId results in an error. Since schedules are based on feed types, you can specify a schedule (schedule_id) that returns the needed feed_type. If specifying the feed_type, limit which tasks are returned by specifying filters, such as the creation date range or period of time using look_back_days. Also, by specifying the feed_type, both on-demand and scheduled reports are returned. If specifying a scheduledId, the schedule template (that the schedule ID is based on) determines which tasks are returned (see schedule_id for additional information). Each scheduledId applies to one feed_type.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetTasksRequest
*/
func (a *TaskApiService) GetTasks(ctx _context.Context) ApiGetTasksRequest {
	return ApiGetTasksRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TaskCollection
func (a *TaskApiService) GetTasksExecute(r ApiGetTasksRequest) (TaskCollection, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TaskCollection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskApiService.GetTasks")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/task"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.dateRange != nil {
		localVarQueryParams.Add("date_range", parameterToString(*r.dateRange, ""))
	}
	if r.feedType != nil {
		localVarQueryParams.Add("feed_type", parameterToString(*r.feedType, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.lookBackDays != nil {
		localVarQueryParams.Add("look_back_days", parameterToString(*r.lookBackDays, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.scheduleId != nil {
		localVarQueryParams.Add("schedule_id", parameterToString(*r.scheduleId, ""))
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

type ApiUploadFileRequest struct {
	ctx _context.Context
	ApiService *TaskApiService
	taskId string
	creationDate *string
	fileName *string
	modificationDate *string
	name *string
	readDate *string
	size *int32
	type_ *string
}

// The file creation date. Format: UTC yyyy-MM-ddThh:mm:ss.SSSZ For example: Created on September 8, 2019 2019-09-08T00:00:00.000Z
func (r ApiUploadFileRequest) CreationDate(creationDate string) ApiUploadFileRequest {
	r.creationDate = &creationDate
	return r
}
// The name of the file including its extension (for example, xml or csv) to be uploaded.
func (r ApiUploadFileRequest) FileName(fileName string) ApiUploadFileRequest {
	r.fileName = &fileName
	return r
}
// The file modified date. Format: UTC yyyy-MM-ddThh:mm:ss.SSSZ For example: Created on September 9, 2019 2019-09-09T00:00:00.000Z
func (r ApiUploadFileRequest) ModificationDate(modificationDate string) ApiUploadFileRequest {
	r.modificationDate = &modificationDate
	return r
}
// A content identifier. The only presently supported name is file.
func (r ApiUploadFileRequest) Name(name string) ApiUploadFileRequest {
	r.name = &name
	return r
}
// The date you read the file. Format: UTC yyyy-MM-ddThh:mm:ss.SSSZ For example: Created on September 10, 2019 2019-09-10T00:00:00.000Z
func (r ApiUploadFileRequest) ReadDate(readDate string) ApiUploadFileRequest {
	r.readDate = &readDate
	return r
}
// The size of the file.
func (r ApiUploadFileRequest) Size(size int32) ApiUploadFileRequest {
	r.size = &size
	return r
}
// The file type. The only presently supported type is form-data.
func (r ApiUploadFileRequest) Type_(type_ string) ApiUploadFileRequest {
	r.type_ = &type_
	return r
}

func (r ApiUploadFileRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.UploadFileExecute(r)
}

/*
UploadFile Method for UploadFile

This method associates the specified file with the specified task ID and uploads the input file. After the file has been uploaded, the processing of the file begins. Reports often take time to generate and it's common for this method to return an HTTP status of 202, which indicates the report is being generated. Use the getTask with the task ID or getTasks to determine the status of a report. The status flow is QUEUED &gt; IN_PROCESS &gt; COMPLETED or COMPLETED_WITH_ERROR. When the status is COMPLETED or COMPLETED_WITH_ERROR, this indicates the file has been processed and the order report can be downloaded. If there are errors, they will be indicated in the report file. For details of how this method is used in the upload flow, see Working with Order Feeds in the Selling Integration Guide. Note: This method applies to all File Exchange feed types and LMS feed types except LMS_ORDER_REPORT and LMS_ACTIVE_INVENTORY_REPORT. See LMS API Feeds in the Selling Integration Guide and File Exchange FeedTypes in the File Exchange Migration Guide. Note: You must use a Content-Type header with its value set to &quot;multipart/form-data&quot;. See Samples for information.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param taskId The task_id associated with the file that will be uploaded. This ID was generated when the specified task was created.
 @return ApiUploadFileRequest
*/
func (a *TaskApiService) UploadFile(ctx _context.Context, taskId string) ApiUploadFileRequest {
	return ApiUploadFileRequest{
		ApiService: a,
		ctx: ctx,
		taskId: taskId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *TaskApiService) UploadFileExecute(r ApiUploadFileRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskApiService.UploadFile")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/task/{task_id}/upload_file"
	localVarPath = strings.Replace(localVarPath, "{"+"task_id"+"}", _neturl.PathEscape(parameterToString(r.taskId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	if r.creationDate != nil {
		localVarFormParams.Add("creationDate", parameterToString(*r.creationDate, ""))
	}
	if r.fileName != nil {
		localVarFormParams.Add("fileName", parameterToString(*r.fileName, ""))
	}
	if r.modificationDate != nil {
		localVarFormParams.Add("modificationDate", parameterToString(*r.modificationDate, ""))
	}
	if r.name != nil {
		localVarFormParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.readDate != nil {
		localVarFormParams.Add("readDate", parameterToString(*r.readDate, ""))
	}
	if r.size != nil {
		localVarFormParams.Add("size", parameterToString(*r.size, ""))
	}
	if r.type_ != nil {
		localVarFormParams.Add("type", parameterToString(*r.type_, ""))
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