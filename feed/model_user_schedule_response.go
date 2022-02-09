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

// UserScheduleResponse The type that defines the fields for a paginated result set of available schedules. The response consists of 0 or more sequenced pages where each page has 0 or more items.
type UserScheduleResponse struct {
	// The ID of the schedule. This ID is generated when the schedule was created by the createSchedule method.
	ScheduleId *string `json:"scheduleId,omitempty"`
	// The creation date of the schedule in hours based on the 24-hour Coordinated Universal Time (UTC) clock.
	CreationDate *string `json:"creationDate,omitempty"`
	// The feedType associated with the schedule.
	FeedType *string `json:"feedType,omitempty"`
	// The date the schedule was last modified.
	LastModifiedDate *string `json:"lastModifiedDate,omitempty"`
	// The preferred day of the month to trigger the schedule. This field can be used with preferredTriggerHour for monthly schedules. The last day of the month is used for numbers larger than the number of days in the month.
	PreferredTriggerDayOfMonth *int32 `json:"preferredTriggerDayOfMonth,omitempty"`
	// The preferred day of the week to trigger the schedule. This field can be used with preferredTriggerHour for weekly schedules. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/sell/feed/types/api:DayOfWeekEnum'>eBay API documentation</a>
	PreferredTriggerDayOfWeek *string `json:"preferredTriggerDayOfWeek,omitempty"`
	// The preferred two-digit hour of the day to trigger the schedule. Format: UTC hhZ For example, the following represents 11:00 am UTC: 11Z
	PreferredTriggerHour *string `json:"preferredTriggerHour,omitempty"`
	// The timestamp on which the report generation (subscription) ends. After this date, the schedule status becomes INACTIVE.
	ScheduleEndDate *string `json:"scheduleEndDate,omitempty"`
	// The schedule name assigned by the user for the created schedule. Users assign this name for their reference.
	ScheduleName *string `json:"scheduleName,omitempty"`
	// The timestamp that indicates the start of the report generation.
	ScheduleStartDate *string `json:"scheduleStartDate,omitempty"`
	// The ID of the template used to create this schedule.
	ScheduleTemplateId *string `json:"scheduleTemplateId,omitempty"`
	// The schema version of the feedType for the schedule.
	SchemaVersion *string `json:"schemaVersion,omitempty"`
	// The enumeration value that indicates the state of the schedule. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/sell/feed/types/api:StatusEnum'>eBay API documentation</a>
	Status *string `json:"status,omitempty"`
	// The reason the schedule is inactive. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/sell/feed/types/api:StatusReasonEnum'>eBay API documentation</a>
	StatusReason *string `json:"statusReason,omitempty"`
}

// NewUserScheduleResponse instantiates a new UserScheduleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserScheduleResponse() *UserScheduleResponse {
	this := UserScheduleResponse{}
	return &this
}

// NewUserScheduleResponseWithDefaults instantiates a new UserScheduleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserScheduleResponseWithDefaults() *UserScheduleResponse {
	this := UserScheduleResponse{}
	return &this
}

// GetScheduleId returns the ScheduleId field value if set, zero value otherwise.
func (o *UserScheduleResponse) GetScheduleId() string {
	if o == nil || o.ScheduleId == nil {
		var ret string
		return ret
	}
	return *o.ScheduleId
}

// GetScheduleIdOk returns a tuple with the ScheduleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserScheduleResponse) GetScheduleIdOk() (*string, bool) {
	if o == nil || o.ScheduleId == nil {
		return nil, false
	}
	return o.ScheduleId, true
}

// HasScheduleId returns a boolean if a field has been set.
func (o *UserScheduleResponse) HasScheduleId() bool {
	if o != nil && o.ScheduleId != nil {
		return true
	}

	return false
}

// SetScheduleId gets a reference to the given string and assigns it to the ScheduleId field.
func (o *UserScheduleResponse) SetScheduleId(v string) {
	o.ScheduleId = &v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *UserScheduleResponse) GetCreationDate() string {
	if o == nil || o.CreationDate == nil {
		var ret string
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserScheduleResponse) GetCreationDateOk() (*string, bool) {
	if o == nil || o.CreationDate == nil {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *UserScheduleResponse) HasCreationDate() bool {
	if o != nil && o.CreationDate != nil {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given string and assigns it to the CreationDate field.
func (o *UserScheduleResponse) SetCreationDate(v string) {
	o.CreationDate = &v
}

// GetFeedType returns the FeedType field value if set, zero value otherwise.
func (o *UserScheduleResponse) GetFeedType() string {
	if o == nil || o.FeedType == nil {
		var ret string
		return ret
	}
	return *o.FeedType
}

// GetFeedTypeOk returns a tuple with the FeedType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserScheduleResponse) GetFeedTypeOk() (*string, bool) {
	if o == nil || o.FeedType == nil {
		return nil, false
	}
	return o.FeedType, true
}

// HasFeedType returns a boolean if a field has been set.
func (o *UserScheduleResponse) HasFeedType() bool {
	if o != nil && o.FeedType != nil {
		return true
	}

	return false
}

// SetFeedType gets a reference to the given string and assigns it to the FeedType field.
func (o *UserScheduleResponse) SetFeedType(v string) {
	o.FeedType = &v
}

// GetLastModifiedDate returns the LastModifiedDate field value if set, zero value otherwise.
func (o *UserScheduleResponse) GetLastModifiedDate() string {
	if o == nil || o.LastModifiedDate == nil {
		var ret string
		return ret
	}
	return *o.LastModifiedDate
}

// GetLastModifiedDateOk returns a tuple with the LastModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserScheduleResponse) GetLastModifiedDateOk() (*string, bool) {
	if o == nil || o.LastModifiedDate == nil {
		return nil, false
	}
	return o.LastModifiedDate, true
}

// HasLastModifiedDate returns a boolean if a field has been set.
func (o *UserScheduleResponse) HasLastModifiedDate() bool {
	if o != nil && o.LastModifiedDate != nil {
		return true
	}

	return false
}

// SetLastModifiedDate gets a reference to the given string and assigns it to the LastModifiedDate field.
func (o *UserScheduleResponse) SetLastModifiedDate(v string) {
	o.LastModifiedDate = &v
}

// GetPreferredTriggerDayOfMonth returns the PreferredTriggerDayOfMonth field value if set, zero value otherwise.
func (o *UserScheduleResponse) GetPreferredTriggerDayOfMonth() int32 {
	if o == nil || o.PreferredTriggerDayOfMonth == nil {
		var ret int32
		return ret
	}
	return *o.PreferredTriggerDayOfMonth
}

// GetPreferredTriggerDayOfMonthOk returns a tuple with the PreferredTriggerDayOfMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserScheduleResponse) GetPreferredTriggerDayOfMonthOk() (*int32, bool) {
	if o == nil || o.PreferredTriggerDayOfMonth == nil {
		return nil, false
	}
	return o.PreferredTriggerDayOfMonth, true
}

// HasPreferredTriggerDayOfMonth returns a boolean if a field has been set.
func (o *UserScheduleResponse) HasPreferredTriggerDayOfMonth() bool {
	if o != nil && o.PreferredTriggerDayOfMonth != nil {
		return true
	}

	return false
}

// SetPreferredTriggerDayOfMonth gets a reference to the given int32 and assigns it to the PreferredTriggerDayOfMonth field.
func (o *UserScheduleResponse) SetPreferredTriggerDayOfMonth(v int32) {
	o.PreferredTriggerDayOfMonth = &v
}

// GetPreferredTriggerDayOfWeek returns the PreferredTriggerDayOfWeek field value if set, zero value otherwise.
func (o *UserScheduleResponse) GetPreferredTriggerDayOfWeek() string {
	if o == nil || o.PreferredTriggerDayOfWeek == nil {
		var ret string
		return ret
	}
	return *o.PreferredTriggerDayOfWeek
}

// GetPreferredTriggerDayOfWeekOk returns a tuple with the PreferredTriggerDayOfWeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserScheduleResponse) GetPreferredTriggerDayOfWeekOk() (*string, bool) {
	if o == nil || o.PreferredTriggerDayOfWeek == nil {
		return nil, false
	}
	return o.PreferredTriggerDayOfWeek, true
}

// HasPreferredTriggerDayOfWeek returns a boolean if a field has been set.
func (o *UserScheduleResponse) HasPreferredTriggerDayOfWeek() bool {
	if o != nil && o.PreferredTriggerDayOfWeek != nil {
		return true
	}

	return false
}

// SetPreferredTriggerDayOfWeek gets a reference to the given string and assigns it to the PreferredTriggerDayOfWeek field.
func (o *UserScheduleResponse) SetPreferredTriggerDayOfWeek(v string) {
	o.PreferredTriggerDayOfWeek = &v
}

// GetPreferredTriggerHour returns the PreferredTriggerHour field value if set, zero value otherwise.
func (o *UserScheduleResponse) GetPreferredTriggerHour() string {
	if o == nil || o.PreferredTriggerHour == nil {
		var ret string
		return ret
	}
	return *o.PreferredTriggerHour
}

// GetPreferredTriggerHourOk returns a tuple with the PreferredTriggerHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserScheduleResponse) GetPreferredTriggerHourOk() (*string, bool) {
	if o == nil || o.PreferredTriggerHour == nil {
		return nil, false
	}
	return o.PreferredTriggerHour, true
}

// HasPreferredTriggerHour returns a boolean if a field has been set.
func (o *UserScheduleResponse) HasPreferredTriggerHour() bool {
	if o != nil && o.PreferredTriggerHour != nil {
		return true
	}

	return false
}

// SetPreferredTriggerHour gets a reference to the given string and assigns it to the PreferredTriggerHour field.
func (o *UserScheduleResponse) SetPreferredTriggerHour(v string) {
	o.PreferredTriggerHour = &v
}

// GetScheduleEndDate returns the ScheduleEndDate field value if set, zero value otherwise.
func (o *UserScheduleResponse) GetScheduleEndDate() string {
	if o == nil || o.ScheduleEndDate == nil {
		var ret string
		return ret
	}
	return *o.ScheduleEndDate
}

// GetScheduleEndDateOk returns a tuple with the ScheduleEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserScheduleResponse) GetScheduleEndDateOk() (*string, bool) {
	if o == nil || o.ScheduleEndDate == nil {
		return nil, false
	}
	return o.ScheduleEndDate, true
}

// HasScheduleEndDate returns a boolean if a field has been set.
func (o *UserScheduleResponse) HasScheduleEndDate() bool {
	if o != nil && o.ScheduleEndDate != nil {
		return true
	}

	return false
}

// SetScheduleEndDate gets a reference to the given string and assigns it to the ScheduleEndDate field.
func (o *UserScheduleResponse) SetScheduleEndDate(v string) {
	o.ScheduleEndDate = &v
}

// GetScheduleName returns the ScheduleName field value if set, zero value otherwise.
func (o *UserScheduleResponse) GetScheduleName() string {
	if o == nil || o.ScheduleName == nil {
		var ret string
		return ret
	}
	return *o.ScheduleName
}

// GetScheduleNameOk returns a tuple with the ScheduleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserScheduleResponse) GetScheduleNameOk() (*string, bool) {
	if o == nil || o.ScheduleName == nil {
		return nil, false
	}
	return o.ScheduleName, true
}

// HasScheduleName returns a boolean if a field has been set.
func (o *UserScheduleResponse) HasScheduleName() bool {
	if o != nil && o.ScheduleName != nil {
		return true
	}

	return false
}

// SetScheduleName gets a reference to the given string and assigns it to the ScheduleName field.
func (o *UserScheduleResponse) SetScheduleName(v string) {
	o.ScheduleName = &v
}

// GetScheduleStartDate returns the ScheduleStartDate field value if set, zero value otherwise.
func (o *UserScheduleResponse) GetScheduleStartDate() string {
	if o == nil || o.ScheduleStartDate == nil {
		var ret string
		return ret
	}
	return *o.ScheduleStartDate
}

// GetScheduleStartDateOk returns a tuple with the ScheduleStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserScheduleResponse) GetScheduleStartDateOk() (*string, bool) {
	if o == nil || o.ScheduleStartDate == nil {
		return nil, false
	}
	return o.ScheduleStartDate, true
}

// HasScheduleStartDate returns a boolean if a field has been set.
func (o *UserScheduleResponse) HasScheduleStartDate() bool {
	if o != nil && o.ScheduleStartDate != nil {
		return true
	}

	return false
}

// SetScheduleStartDate gets a reference to the given string and assigns it to the ScheduleStartDate field.
func (o *UserScheduleResponse) SetScheduleStartDate(v string) {
	o.ScheduleStartDate = &v
}

// GetScheduleTemplateId returns the ScheduleTemplateId field value if set, zero value otherwise.
func (o *UserScheduleResponse) GetScheduleTemplateId() string {
	if o == nil || o.ScheduleTemplateId == nil {
		var ret string
		return ret
	}
	return *o.ScheduleTemplateId
}

// GetScheduleTemplateIdOk returns a tuple with the ScheduleTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserScheduleResponse) GetScheduleTemplateIdOk() (*string, bool) {
	if o == nil || o.ScheduleTemplateId == nil {
		return nil, false
	}
	return o.ScheduleTemplateId, true
}

// HasScheduleTemplateId returns a boolean if a field has been set.
func (o *UserScheduleResponse) HasScheduleTemplateId() bool {
	if o != nil && o.ScheduleTemplateId != nil {
		return true
	}

	return false
}

// SetScheduleTemplateId gets a reference to the given string and assigns it to the ScheduleTemplateId field.
func (o *UserScheduleResponse) SetScheduleTemplateId(v string) {
	o.ScheduleTemplateId = &v
}

// GetSchemaVersion returns the SchemaVersion field value if set, zero value otherwise.
func (o *UserScheduleResponse) GetSchemaVersion() string {
	if o == nil || o.SchemaVersion == nil {
		var ret string
		return ret
	}
	return *o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserScheduleResponse) GetSchemaVersionOk() (*string, bool) {
	if o == nil || o.SchemaVersion == nil {
		return nil, false
	}
	return o.SchemaVersion, true
}

// HasSchemaVersion returns a boolean if a field has been set.
func (o *UserScheduleResponse) HasSchemaVersion() bool {
	if o != nil && o.SchemaVersion != nil {
		return true
	}

	return false
}

// SetSchemaVersion gets a reference to the given string and assigns it to the SchemaVersion field.
func (o *UserScheduleResponse) SetSchemaVersion(v string) {
	o.SchemaVersion = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UserScheduleResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserScheduleResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UserScheduleResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *UserScheduleResponse) SetStatus(v string) {
	o.Status = &v
}

// GetStatusReason returns the StatusReason field value if set, zero value otherwise.
func (o *UserScheduleResponse) GetStatusReason() string {
	if o == nil || o.StatusReason == nil {
		var ret string
		return ret
	}
	return *o.StatusReason
}

// GetStatusReasonOk returns a tuple with the StatusReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserScheduleResponse) GetStatusReasonOk() (*string, bool) {
	if o == nil || o.StatusReason == nil {
		return nil, false
	}
	return o.StatusReason, true
}

// HasStatusReason returns a boolean if a field has been set.
func (o *UserScheduleResponse) HasStatusReason() bool {
	if o != nil && o.StatusReason != nil {
		return true
	}

	return false
}

// SetStatusReason gets a reference to the given string and assigns it to the StatusReason field.
func (o *UserScheduleResponse) SetStatusReason(v string) {
	o.StatusReason = &v
}

func (o UserScheduleResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ScheduleId != nil {
		toSerialize["scheduleId"] = o.ScheduleId
	}
	if o.CreationDate != nil {
		toSerialize["creationDate"] = o.CreationDate
	}
	if o.FeedType != nil {
		toSerialize["feedType"] = o.FeedType
	}
	if o.LastModifiedDate != nil {
		toSerialize["lastModifiedDate"] = o.LastModifiedDate
	}
	if o.PreferredTriggerDayOfMonth != nil {
		toSerialize["preferredTriggerDayOfMonth"] = o.PreferredTriggerDayOfMonth
	}
	if o.PreferredTriggerDayOfWeek != nil {
		toSerialize["preferredTriggerDayOfWeek"] = o.PreferredTriggerDayOfWeek
	}
	if o.PreferredTriggerHour != nil {
		toSerialize["preferredTriggerHour"] = o.PreferredTriggerHour
	}
	if o.ScheduleEndDate != nil {
		toSerialize["scheduleEndDate"] = o.ScheduleEndDate
	}
	if o.ScheduleName != nil {
		toSerialize["scheduleName"] = o.ScheduleName
	}
	if o.ScheduleStartDate != nil {
		toSerialize["scheduleStartDate"] = o.ScheduleStartDate
	}
	if o.ScheduleTemplateId != nil {
		toSerialize["scheduleTemplateId"] = o.ScheduleTemplateId
	}
	if o.SchemaVersion != nil {
		toSerialize["schemaVersion"] = o.SchemaVersion
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.StatusReason != nil {
		toSerialize["statusReason"] = o.StatusReason
	}
	return json.Marshal(toSerialize)
}

type NullableUserScheduleResponse struct {
	value *UserScheduleResponse
	isSet bool
}

func (v NullableUserScheduleResponse) Get() *UserScheduleResponse {
	return v.value
}

func (v *NullableUserScheduleResponse) Set(val *UserScheduleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserScheduleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserScheduleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserScheduleResponse(val *UserScheduleResponse) *NullableUserScheduleResponse {
	return &NullableUserScheduleResponse{value: val, isSet: true}
}

func (v NullableUserScheduleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserScheduleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


