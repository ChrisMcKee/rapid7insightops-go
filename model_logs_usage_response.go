/*
InsightOps REST API

### Overview  Our REST API lets you use InsightOps through HTTP requests. Currently, the REST API allows you to perform the majority of the actions available through the UI, and has some additional functionality that is not available through the UI. You may use this API to automate common tasks (for example, via shell scripts), and to generally interact with InsightOps programmatically.  This page precisely describes the REST API and serves as a reference for the API. Each HTTP method and each URL endpoint is documented in a self-contained unit so that users only need to read about the HTTP methods relevant to them.  ### Terminology  * A **log entry** is an individual log event. * A **log** is a collection of log entries, or a single log stream. * A **log set** is a logical-only collection of logs, i.e. logs can be in multiple logsets and deleting a logset only deletes the relation between logs, not the logs themselves. * [Log Entry Query Language](https://docs.rapid7.com/insightops/log-search) (**LEQL**) is the query language used in Insight Ops to search log data.

API version: latest
Contact: support@rapid7.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package insightops

import (
	"encoding/json"
)

// checks if the LogsUsageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogsUsageResponse{}

// LogsUsageResponse struct for LogsUsageResponse
type LogsUsageResponse struct {
	PerDayUsage *LogsUsageResponsePerDayUsage `json:"per_day_usage,omitempty"`
}

// NewLogsUsageResponse instantiates a new LogsUsageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogsUsageResponse() *LogsUsageResponse {
	this := LogsUsageResponse{}
	return &this
}

// NewLogsUsageResponseWithDefaults instantiates a new LogsUsageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogsUsageResponseWithDefaults() *LogsUsageResponse {
	this := LogsUsageResponse{}
	return &this
}

// GetPerDayUsage returns the PerDayUsage field value if set, zero value otherwise.
func (o *LogsUsageResponse) GetPerDayUsage() LogsUsageResponsePerDayUsage {
	if o == nil || IsNil(o.PerDayUsage) {
		var ret LogsUsageResponsePerDayUsage
		return ret
	}
	return *o.PerDayUsage
}

// GetPerDayUsageOk returns a tuple with the PerDayUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsUsageResponse) GetPerDayUsageOk() (*LogsUsageResponsePerDayUsage, bool) {
	if o == nil || IsNil(o.PerDayUsage) {
		return nil, false
	}
	return o.PerDayUsage, true
}

// HasPerDayUsage returns a boolean if a field has been set.
func (o *LogsUsageResponse) HasPerDayUsage() bool {
	if o != nil && !IsNil(o.PerDayUsage) {
		return true
	}

	return false
}

// SetPerDayUsage gets a reference to the given LogsUsageResponsePerDayUsage and assigns it to the PerDayUsage field.
func (o *LogsUsageResponse) SetPerDayUsage(v LogsUsageResponsePerDayUsage) {
	o.PerDayUsage = &v
}

func (o LogsUsageResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogsUsageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PerDayUsage) {
		toSerialize["per_day_usage"] = o.PerDayUsage
	}
	return toSerialize, nil
}

type NullableLogsUsageResponse struct {
	value *LogsUsageResponse
	isSet bool
}

func (v NullableLogsUsageResponse) Get() *LogsUsageResponse {
	return v.value
}

func (v *NullableLogsUsageResponse) Set(val *LogsUsageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsUsageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsUsageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsUsageResponse(val *LogsUsageResponse) *NullableLogsUsageResponse {
	return &NullableLogsUsageResponse{value: val, isSet: true}
}

func (v NullableLogsUsageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsUsageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}