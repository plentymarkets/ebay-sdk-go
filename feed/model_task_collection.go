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

// TaskCollection The type that defines the fields for a paginated result set of tasks. The response consists of 0 or more sequenced pages where each page has 0 or more items.
type TaskCollection struct {
	// The path to the call URI that produced the current page of results.
	Href *string `json:"href,omitempty"`
	// The value of the limit parameter submitted in the request, which is the maximum number of tasks to return per page, from the result set. A result set is the complete set of tasks returned by the method. Note: Though this parameter is not required to be submitted in the request, the parameter defaults to 10 if omitted. Note: If this is the last or only page of the result set, the page may contain fewer tasks than the limit value. To determine the number of pages in a result set, divide the total value (total number of tasks matching input criteria) by this limit value, and then round up to the next integer. For example, if the total value was 120 (120 total tasks) and the limit value was 50 (show 50 tasks per page), the total number of pages in the result set is three, so the seller would have to make three separate getTasks calls to view all tasks matching the input criteria.
	Limit *int32 `json:"limit,omitempty"`
	// The path to the call URI for the next page of results. This value is returned if there is an additional page of results to return from the result set.
	Next *string `json:"next,omitempty"`
	// The number of results skipped in the result set before listing the first returned result. This value can be set in the request with the offset query parameter. Note: The items in a paginated result set use a zero-based list where the first item in the list has an offset of 0.
	Offset *int32 `json:"offset,omitempty"`
	// The path to the call URI for the previous page of results. This is returned if there is a previous page of results from the result set.
	Prev *string `json:"prev,omitempty"`
	// An array of the tasks on this page. The tasks are sorted by creation date. An empty array is returned if the filter criteria excludes all tasks.
	Tasks *[]Task `json:"tasks,omitempty"`
	// The total number of tasks that match the input criteria.
	Total *int32 `json:"total,omitempty"`
}

// NewTaskCollection instantiates a new TaskCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskCollection() *TaskCollection {
	this := TaskCollection{}
	return &this
}

// NewTaskCollectionWithDefaults instantiates a new TaskCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskCollectionWithDefaults() *TaskCollection {
	this := TaskCollection{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *TaskCollection) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskCollection) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *TaskCollection) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *TaskCollection) SetHref(v string) {
	o.Href = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *TaskCollection) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskCollection) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *TaskCollection) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *TaskCollection) SetLimit(v int32) {
	o.Limit = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *TaskCollection) GetNext() string {
	if o == nil || o.Next == nil {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskCollection) GetNextOk() (*string, bool) {
	if o == nil || o.Next == nil {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *TaskCollection) HasNext() bool {
	if o != nil && o.Next != nil {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *TaskCollection) SetNext(v string) {
	o.Next = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *TaskCollection) GetOffset() int32 {
	if o == nil || o.Offset == nil {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskCollection) GetOffsetOk() (*int32, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *TaskCollection) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *TaskCollection) SetOffset(v int32) {
	o.Offset = &v
}

// GetPrev returns the Prev field value if set, zero value otherwise.
func (o *TaskCollection) GetPrev() string {
	if o == nil || o.Prev == nil {
		var ret string
		return ret
	}
	return *o.Prev
}

// GetPrevOk returns a tuple with the Prev field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskCollection) GetPrevOk() (*string, bool) {
	if o == nil || o.Prev == nil {
		return nil, false
	}
	return o.Prev, true
}

// HasPrev returns a boolean if a field has been set.
func (o *TaskCollection) HasPrev() bool {
	if o != nil && o.Prev != nil {
		return true
	}

	return false
}

// SetPrev gets a reference to the given string and assigns it to the Prev field.
func (o *TaskCollection) SetPrev(v string) {
	o.Prev = &v
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *TaskCollection) GetTasks() []Task {
	if o == nil || o.Tasks == nil {
		var ret []Task
		return ret
	}
	return *o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskCollection) GetTasksOk() (*[]Task, bool) {
	if o == nil || o.Tasks == nil {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *TaskCollection) HasTasks() bool {
	if o != nil && o.Tasks != nil {
		return true
	}

	return false
}

// SetTasks gets a reference to the given []Task and assigns it to the Tasks field.
func (o *TaskCollection) SetTasks(v []Task) {
	o.Tasks = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *TaskCollection) GetTotal() int32 {
	if o == nil || o.Total == nil {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskCollection) GetTotalOk() (*int32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *TaskCollection) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *TaskCollection) SetTotal(v int32) {
	o.Total = &v
}

func (o TaskCollection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Next != nil {
		toSerialize["next"] = o.Next
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	if o.Prev != nil {
		toSerialize["prev"] = o.Prev
	}
	if o.Tasks != nil {
		toSerialize["tasks"] = o.Tasks
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableTaskCollection struct {
	value *TaskCollection
	isSet bool
}

func (v NullableTaskCollection) Get() *TaskCollection {
	return v.value
}

func (v *NullableTaskCollection) Set(val *TaskCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskCollection(val *TaskCollection) *NullableTaskCollection {
	return &NullableTaskCollection{value: val, isSet: true}
}

func (v NullableTaskCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


