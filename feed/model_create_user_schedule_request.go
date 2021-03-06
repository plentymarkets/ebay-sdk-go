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

// CreateUserScheduleRequest The type that defines the fields for the createSchedule method.
type CreateUserScheduleRequest struct {
	// The name of the feed type for the created schedule. Match the feed_type from the schedule template associated with this schedule.
	FeedType *string `json:"feedType,omitempty"`
	// The preferred day of the month to trigger the schedule. This field can be used with preferredTriggerHour for monthly schedules. The last day of the month is used for numbers larger than the actual number of days in the month. This field is available as specified by the template (scheduleTemplateId). The template can specify this field as optional or required, and optionally provides a default value. Minimum: 1 Maximum: 31
	PreferredTriggerDayOfMonth *int32 `json:"preferredTriggerDayOfMonth,omitempty"`
	// The preferred day of the week to trigger the schedule. This field can be used with preferredTriggerHour for weekly schedules. This field is available as specified by the template (scheduleTemplateId). The template can specify this field as optional or required, and optionally provides a default value. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/sell/feed/types/api:DayOfWeekEnum'>eBay API documentation</a>
	PreferredTriggerDayOfWeek *string `json:"preferredTriggerDayOfWeek,omitempty"`
	// The preferred two-digit hour of the day to trigger the schedule. This field is available as specified by the template (scheduleTemplateId). The template can specify this field as optional or required, and optionally provides a default value. Format: UTC hhZ For example, the following represents 11:00 am UTC: 11Z
	PreferredTriggerHour *string `json:"preferredTriggerHour,omitempty"`
	// The timestamp on which the report generation (subscription) ends. After this date, the schedule status becomes INACTIVE. Use this field, if available, to end the schedule in the future. This value must be later than scheduleStartDate (if supplied). This field is available as specified by the template (scheduleTemplateId). The template can specify this field as optional or required, and optionally provides a default value. Format: UTC yyyy-MM-ddTHHZ For example, the following represents UTC October 10, 2021 at 10:00 AM: 2021-10-10T10Z
	ScheduleEndDate *string `json:"scheduleEndDate,omitempty"`
	// The schedule name assigned by the user for the created schedule.
	ScheduleName *string `json:"scheduleName,omitempty"`
	// The timestamp to start generating the report. After this timestamp, the schedule status becomes active until either the scheduleEndDate occurs or the scheduleTemplateId becomes inactive. Use this field, if available, to start the schedule in the future but before the scheduleEndDate (if supplied). This field is available as specified by the template (scheduleTemplateId). The template can specify this field as optional or required, and optionally provides a default value. Format: UTC yyyy-MM-ddTHHZ For example, the following represents a schedule start date of UTC October 01, 2020 at 12:00 PM: 2020-01-01T12Z
	ScheduleStartDate *string `json:"scheduleStartDate,omitempty"`
	// The ID of the template associated with the schedule ID. You can get this ID from the documentation or by calling the getScheduleTemplates method. This method requires a schedule template ID that is ACTIVE.
	ScheduleTemplateId *string `json:"scheduleTemplateId,omitempty"`
	// The schema version of the schedule feedType. This field is required if the feedType has a schema version. This field is available as specified by the template (scheduleTemplateId). The template can specify this field as optional or required, and optionally provides a default value.
	SchemaVersion *string `json:"schemaVersion,omitempty"`
}

// NewCreateUserScheduleRequest instantiates a new CreateUserScheduleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUserScheduleRequest() *CreateUserScheduleRequest {
	this := CreateUserScheduleRequest{}
	return &this
}

// NewCreateUserScheduleRequestWithDefaults instantiates a new CreateUserScheduleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUserScheduleRequestWithDefaults() *CreateUserScheduleRequest {
	this := CreateUserScheduleRequest{}
	return &this
}

// GetFeedType returns the FeedType field value if set, zero value otherwise.
func (o *CreateUserScheduleRequest) GetFeedType() string {
	if o == nil || o.FeedType == nil {
		var ret string
		return ret
	}
	return *o.FeedType
}

// GetFeedTypeOk returns a tuple with the FeedType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserScheduleRequest) GetFeedTypeOk() (*string, bool) {
	if o == nil || o.FeedType == nil {
		return nil, false
	}
	return o.FeedType, true
}

// HasFeedType returns a boolean if a field has been set.
func (o *CreateUserScheduleRequest) HasFeedType() bool {
	if o != nil && o.FeedType != nil {
		return true
	}

	return false
}

// SetFeedType gets a reference to the given string and assigns it to the FeedType field.
func (o *CreateUserScheduleRequest) SetFeedType(v string) {
	o.FeedType = &v
}

// GetPreferredTriggerDayOfMonth returns the PreferredTriggerDayOfMonth field value if set, zero value otherwise.
func (o *CreateUserScheduleRequest) GetPreferredTriggerDayOfMonth() int32 {
	if o == nil || o.PreferredTriggerDayOfMonth == nil {
		var ret int32
		return ret
	}
	return *o.PreferredTriggerDayOfMonth
}

// GetPreferredTriggerDayOfMonthOk returns a tuple with the PreferredTriggerDayOfMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserScheduleRequest) GetPreferredTriggerDayOfMonthOk() (*int32, bool) {
	if o == nil || o.PreferredTriggerDayOfMonth == nil {
		return nil, false
	}
	return o.PreferredTriggerDayOfMonth, true
}

// HasPreferredTriggerDayOfMonth returns a boolean if a field has been set.
func (o *CreateUserScheduleRequest) HasPreferredTriggerDayOfMonth() bool {
	if o != nil && o.PreferredTriggerDayOfMonth != nil {
		return true
	}

	return false
}

// SetPreferredTriggerDayOfMonth gets a reference to the given int32 and assigns it to the PreferredTriggerDayOfMonth field.
func (o *CreateUserScheduleRequest) SetPreferredTriggerDayOfMonth(v int32) {
	o.PreferredTriggerDayOfMonth = &v
}

// GetPreferredTriggerDayOfWeek returns the PreferredTriggerDayOfWeek field value if set, zero value otherwise.
func (o *CreateUserScheduleRequest) GetPreferredTriggerDayOfWeek() string {
	if o == nil || o.PreferredTriggerDayOfWeek == nil {
		var ret string
		return ret
	}
	return *o.PreferredTriggerDayOfWeek
}

// GetPreferredTriggerDayOfWeekOk returns a tuple with the PreferredTriggerDayOfWeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserScheduleRequest) GetPreferredTriggerDayOfWeekOk() (*string, bool) {
	if o == nil || o.PreferredTriggerDayOfWeek == nil {
		return nil, false
	}
	return o.PreferredTriggerDayOfWeek, true
}

// HasPreferredTriggerDayOfWeek returns a boolean if a field has been set.
func (o *CreateUserScheduleRequest) HasPreferredTriggerDayOfWeek() bool {
	if o != nil && o.PreferredTriggerDayOfWeek != nil {
		return true
	}

	return false
}

// SetPreferredTriggerDayOfWeek gets a reference to the given string and assigns it to the PreferredTriggerDayOfWeek field.
func (o *CreateUserScheduleRequest) SetPreferredTriggerDayOfWeek(v string) {
	o.PreferredTriggerDayOfWeek = &v
}

// GetPreferredTriggerHour returns the PreferredTriggerHour field value if set, zero value otherwise.
func (o *CreateUserScheduleRequest) GetPreferredTriggerHour() string {
	if o == nil || o.PreferredTriggerHour == nil {
		var ret string
		return ret
	}
	return *o.PreferredTriggerHour
}

// GetPreferredTriggerHourOk returns a tuple with the PreferredTriggerHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserScheduleRequest) GetPreferredTriggerHourOk() (*string, bool) {
	if o == nil || o.PreferredTriggerHour == nil {
		return nil, false
	}
	return o.PreferredTriggerHour, true
}

// HasPreferredTriggerHour returns a boolean if a field has been set.
func (o *CreateUserScheduleRequest) HasPreferredTriggerHour() bool {
	if o != nil && o.PreferredTriggerHour != nil {
		return true
	}

	return false
}

// SetPreferredTriggerHour gets a reference to the given string and assigns it to the PreferredTriggerHour field.
func (o *CreateUserScheduleRequest) SetPreferredTriggerHour(v string) {
	o.PreferredTriggerHour = &v
}

// GetScheduleEndDate returns the ScheduleEndDate field value if set, zero value otherwise.
func (o *CreateUserScheduleRequest) GetScheduleEndDate() string {
	if o == nil || o.ScheduleEndDate == nil {
		var ret string
		return ret
	}
	return *o.ScheduleEndDate
}

// GetScheduleEndDateOk returns a tuple with the ScheduleEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserScheduleRequest) GetScheduleEndDateOk() (*string, bool) {
	if o == nil || o.ScheduleEndDate == nil {
		return nil, false
	}
	return o.ScheduleEndDate, true
}

// HasScheduleEndDate returns a boolean if a field has been set.
func (o *CreateUserScheduleRequest) HasScheduleEndDate() bool {
	if o != nil && o.ScheduleEndDate != nil {
		return true
	}

	return false
}

// SetScheduleEndDate gets a reference to the given string and assigns it to the ScheduleEndDate field.
func (o *CreateUserScheduleRequest) SetScheduleEndDate(v string) {
	o.ScheduleEndDate = &v
}

// GetScheduleName returns the ScheduleName field value if set, zero value otherwise.
func (o *CreateUserScheduleRequest) GetScheduleName() string {
	if o == nil || o.ScheduleName == nil {
		var ret string
		return ret
	}
	return *o.ScheduleName
}

// GetScheduleNameOk returns a tuple with the ScheduleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserScheduleRequest) GetScheduleNameOk() (*string, bool) {
	if o == nil || o.ScheduleName == nil {
		return nil, false
	}
	return o.ScheduleName, true
}

// HasScheduleName returns a boolean if a field has been set.
func (o *CreateUserScheduleRequest) HasScheduleName() bool {
	if o != nil && o.ScheduleName != nil {
		return true
	}

	return false
}

// SetScheduleName gets a reference to the given string and assigns it to the ScheduleName field.
func (o *CreateUserScheduleRequest) SetScheduleName(v string) {
	o.ScheduleName = &v
}

// GetScheduleStartDate returns the ScheduleStartDate field value if set, zero value otherwise.
func (o *CreateUserScheduleRequest) GetScheduleStartDate() string {
	if o == nil || o.ScheduleStartDate == nil {
		var ret string
		return ret
	}
	return *o.ScheduleStartDate
}

// GetScheduleStartDateOk returns a tuple with the ScheduleStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserScheduleRequest) GetScheduleStartDateOk() (*string, bool) {
	if o == nil || o.ScheduleStartDate == nil {
		return nil, false
	}
	return o.ScheduleStartDate, true
}

// HasScheduleStartDate returns a boolean if a field has been set.
func (o *CreateUserScheduleRequest) HasScheduleStartDate() bool {
	if o != nil && o.ScheduleStartDate != nil {
		return true
	}

	return false
}

// SetScheduleStartDate gets a reference to the given string and assigns it to the ScheduleStartDate field.
func (o *CreateUserScheduleRequest) SetScheduleStartDate(v string) {
	o.ScheduleStartDate = &v
}

// GetScheduleTemplateId returns the ScheduleTemplateId field value if set, zero value otherwise.
func (o *CreateUserScheduleRequest) GetScheduleTemplateId() string {
	if o == nil || o.ScheduleTemplateId == nil {
		var ret string
		return ret
	}
	return *o.ScheduleTemplateId
}

// GetScheduleTemplateIdOk returns a tuple with the ScheduleTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserScheduleRequest) GetScheduleTemplateIdOk() (*string, bool) {
	if o == nil || o.ScheduleTemplateId == nil {
		return nil, false
	}
	return o.ScheduleTemplateId, true
}

// HasScheduleTemplateId returns a boolean if a field has been set.
func (o *CreateUserScheduleRequest) HasScheduleTemplateId() bool {
	if o != nil && o.ScheduleTemplateId != nil {
		return true
	}

	return false
}

// SetScheduleTemplateId gets a reference to the given string and assigns it to the ScheduleTemplateId field.
func (o *CreateUserScheduleRequest) SetScheduleTemplateId(v string) {
	o.ScheduleTemplateId = &v
}

// GetSchemaVersion returns the SchemaVersion field value if set, zero value otherwise.
func (o *CreateUserScheduleRequest) GetSchemaVersion() string {
	if o == nil || o.SchemaVersion == nil {
		var ret string
		return ret
	}
	return *o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserScheduleRequest) GetSchemaVersionOk() (*string, bool) {
	if o == nil || o.SchemaVersion == nil {
		return nil, false
	}
	return o.SchemaVersion, true
}

// HasSchemaVersion returns a boolean if a field has been set.
func (o *CreateUserScheduleRequest) HasSchemaVersion() bool {
	if o != nil && o.SchemaVersion != nil {
		return true
	}

	return false
}

// SetSchemaVersion gets a reference to the given string and assigns it to the SchemaVersion field.
func (o *CreateUserScheduleRequest) SetSchemaVersion(v string) {
	o.SchemaVersion = &v
}

func (o CreateUserScheduleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FeedType != nil {
		toSerialize["feedType"] = o.FeedType
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
	return json.Marshal(toSerialize)
}

type NullableCreateUserScheduleRequest struct {
	value *CreateUserScheduleRequest
	isSet bool
}

func (v NullableCreateUserScheduleRequest) Get() *CreateUserScheduleRequest {
	return v.value
}

func (v *NullableCreateUserScheduleRequest) Set(val *CreateUserScheduleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUserScheduleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUserScheduleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUserScheduleRequest(val *CreateUserScheduleRequest) *NullableCreateUserScheduleRequest {
	return &NullableCreateUserScheduleRequest{value: val, isSet: true}
}

func (v NullableCreateUserScheduleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUserScheduleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


