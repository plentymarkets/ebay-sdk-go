/*
Feed API

<p>The <strong>Feed API</strong> lets sellers upload input files, download reports and files including their status, filter reports using URI parameters, and retrieve customer service metrics task details.</p>

API version: v1.3.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package feed

import (
	"encoding/json"
)

// ServiceMetricsTask struct for ServiceMetricsTask
type ServiceMetricsTask struct {
	// The timestamp when the customer service metrics task went into the COMPLETED or COMPLETED_WITH_ERROR state. This field is only returned if the status is one of the two completed values. This state means that eBay has compiled the report for the seller based on the seller&rsquo;s filter criteria, and the seller can run a getResultFile call to download the report.
	CompletionDate *string `json:"completionDate,omitempty"`
	// The date the customer service metrics task was created.
	CreationDate *string `json:"creationDate,omitempty"`
	// The relative getCustomerServiceMetricTask call URI path to retrieve the corresponding task.
	DetailHref *string `json:"detailHref,omitempty"`
	// The feed type associated with the task.
	FeedType *string `json:"feedType,omitempty"`
	FilterCriteria *CustomerServiceMetricsFilterCriteria `json:"filterCriteria,omitempty"`
	// The schema version number of the file format. If omitted, the default value is used. Default value: 1.0
	SchemaVersion *string `json:"schemaVersion,omitempty"`
	// An enumeration value that indicates the state of the task. See FeedStatusEnum for values. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/sell/feed/types/api:FeedStatusEnum'>eBay API documentation</a>
	Status *string `json:"status,omitempty"`
	// The unique eBay-assigned ID of the task.
	TaskId *string `json:"taskId,omitempty"`
}

// NewServiceMetricsTask instantiates a new ServiceMetricsTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceMetricsTask() *ServiceMetricsTask {
	this := ServiceMetricsTask{}
	return &this
}

// NewServiceMetricsTaskWithDefaults instantiates a new ServiceMetricsTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceMetricsTaskWithDefaults() *ServiceMetricsTask {
	this := ServiceMetricsTask{}
	return &this
}

// GetCompletionDate returns the CompletionDate field value if set, zero value otherwise.
func (o *ServiceMetricsTask) GetCompletionDate() string {
	if o == nil || o.CompletionDate == nil {
		var ret string
		return ret
	}
	return *o.CompletionDate
}

// GetCompletionDateOk returns a tuple with the CompletionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceMetricsTask) GetCompletionDateOk() (*string, bool) {
	if o == nil || o.CompletionDate == nil {
		return nil, false
	}
	return o.CompletionDate, true
}

// HasCompletionDate returns a boolean if a field has been set.
func (o *ServiceMetricsTask) HasCompletionDate() bool {
	if o != nil && o.CompletionDate != nil {
		return true
	}

	return false
}

// SetCompletionDate gets a reference to the given string and assigns it to the CompletionDate field.
func (o *ServiceMetricsTask) SetCompletionDate(v string) {
	o.CompletionDate = &v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *ServiceMetricsTask) GetCreationDate() string {
	if o == nil || o.CreationDate == nil {
		var ret string
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceMetricsTask) GetCreationDateOk() (*string, bool) {
	if o == nil || o.CreationDate == nil {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *ServiceMetricsTask) HasCreationDate() bool {
	if o != nil && o.CreationDate != nil {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given string and assigns it to the CreationDate field.
func (o *ServiceMetricsTask) SetCreationDate(v string) {
	o.CreationDate = &v
}

// GetDetailHref returns the DetailHref field value if set, zero value otherwise.
func (o *ServiceMetricsTask) GetDetailHref() string {
	if o == nil || o.DetailHref == nil {
		var ret string
		return ret
	}
	return *o.DetailHref
}

// GetDetailHrefOk returns a tuple with the DetailHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceMetricsTask) GetDetailHrefOk() (*string, bool) {
	if o == nil || o.DetailHref == nil {
		return nil, false
	}
	return o.DetailHref, true
}

// HasDetailHref returns a boolean if a field has been set.
func (o *ServiceMetricsTask) HasDetailHref() bool {
	if o != nil && o.DetailHref != nil {
		return true
	}

	return false
}

// SetDetailHref gets a reference to the given string and assigns it to the DetailHref field.
func (o *ServiceMetricsTask) SetDetailHref(v string) {
	o.DetailHref = &v
}

// GetFeedType returns the FeedType field value if set, zero value otherwise.
func (o *ServiceMetricsTask) GetFeedType() string {
	if o == nil || o.FeedType == nil {
		var ret string
		return ret
	}
	return *o.FeedType
}

// GetFeedTypeOk returns a tuple with the FeedType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceMetricsTask) GetFeedTypeOk() (*string, bool) {
	if o == nil || o.FeedType == nil {
		return nil, false
	}
	return o.FeedType, true
}

// HasFeedType returns a boolean if a field has been set.
func (o *ServiceMetricsTask) HasFeedType() bool {
	if o != nil && o.FeedType != nil {
		return true
	}

	return false
}

// SetFeedType gets a reference to the given string and assigns it to the FeedType field.
func (o *ServiceMetricsTask) SetFeedType(v string) {
	o.FeedType = &v
}

// GetFilterCriteria returns the FilterCriteria field value if set, zero value otherwise.
func (o *ServiceMetricsTask) GetFilterCriteria() CustomerServiceMetricsFilterCriteria {
	if o == nil || o.FilterCriteria == nil {
		var ret CustomerServiceMetricsFilterCriteria
		return ret
	}
	return *o.FilterCriteria
}

// GetFilterCriteriaOk returns a tuple with the FilterCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceMetricsTask) GetFilterCriteriaOk() (*CustomerServiceMetricsFilterCriteria, bool) {
	if o == nil || o.FilterCriteria == nil {
		return nil, false
	}
	return o.FilterCriteria, true
}

// HasFilterCriteria returns a boolean if a field has been set.
func (o *ServiceMetricsTask) HasFilterCriteria() bool {
	if o != nil && o.FilterCriteria != nil {
		return true
	}

	return false
}

// SetFilterCriteria gets a reference to the given CustomerServiceMetricsFilterCriteria and assigns it to the FilterCriteria field.
func (o *ServiceMetricsTask) SetFilterCriteria(v CustomerServiceMetricsFilterCriteria) {
	o.FilterCriteria = &v
}

// GetSchemaVersion returns the SchemaVersion field value if set, zero value otherwise.
func (o *ServiceMetricsTask) GetSchemaVersion() string {
	if o == nil || o.SchemaVersion == nil {
		var ret string
		return ret
	}
	return *o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceMetricsTask) GetSchemaVersionOk() (*string, bool) {
	if o == nil || o.SchemaVersion == nil {
		return nil, false
	}
	return o.SchemaVersion, true
}

// HasSchemaVersion returns a boolean if a field has been set.
func (o *ServiceMetricsTask) HasSchemaVersion() bool {
	if o != nil && o.SchemaVersion != nil {
		return true
	}

	return false
}

// SetSchemaVersion gets a reference to the given string and assigns it to the SchemaVersion field.
func (o *ServiceMetricsTask) SetSchemaVersion(v string) {
	o.SchemaVersion = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ServiceMetricsTask) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceMetricsTask) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ServiceMetricsTask) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ServiceMetricsTask) SetStatus(v string) {
	o.Status = &v
}

// GetTaskId returns the TaskId field value if set, zero value otherwise.
func (o *ServiceMetricsTask) GetTaskId() string {
	if o == nil || o.TaskId == nil {
		var ret string
		return ret
	}
	return *o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceMetricsTask) GetTaskIdOk() (*string, bool) {
	if o == nil || o.TaskId == nil {
		return nil, false
	}
	return o.TaskId, true
}

// HasTaskId returns a boolean if a field has been set.
func (o *ServiceMetricsTask) HasTaskId() bool {
	if o != nil && o.TaskId != nil {
		return true
	}

	return false
}

// SetTaskId gets a reference to the given string and assigns it to the TaskId field.
func (o *ServiceMetricsTask) SetTaskId(v string) {
	o.TaskId = &v
}

func (o ServiceMetricsTask) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CompletionDate != nil {
		toSerialize["completionDate"] = o.CompletionDate
	}
	if o.CreationDate != nil {
		toSerialize["creationDate"] = o.CreationDate
	}
	if o.DetailHref != nil {
		toSerialize["detailHref"] = o.DetailHref
	}
	if o.FeedType != nil {
		toSerialize["feedType"] = o.FeedType
	}
	if o.FilterCriteria != nil {
		toSerialize["filterCriteria"] = o.FilterCriteria
	}
	if o.SchemaVersion != nil {
		toSerialize["schemaVersion"] = o.SchemaVersion
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.TaskId != nil {
		toSerialize["taskId"] = o.TaskId
	}
	return json.Marshal(toSerialize)
}

type NullableServiceMetricsTask struct {
	value *ServiceMetricsTask
	isSet bool
}

func (v NullableServiceMetricsTask) Get() *ServiceMetricsTask {
	return v.value
}

func (v *NullableServiceMetricsTask) Set(val *ServiceMetricsTask) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceMetricsTask) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceMetricsTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceMetricsTask(val *ServiceMetricsTask) *NullableServiceMetricsTask {
	return &NullableServiceMetricsTask{value: val, isSet: true}
}

func (v NullableServiceMetricsTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceMetricsTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


