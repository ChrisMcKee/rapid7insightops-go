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

// checks if the MetricsInputMetricLogs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricsInputMetricLogs{}

// MetricsInputMetricLogs The logs to use in the pre-computed query.
type MetricsInputMetricLogs struct {
	// The UUID of the logs.
	Id *string `json:"id,omitempty"`
}

// NewMetricsInputMetricLogs instantiates a new MetricsInputMetricLogs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsInputMetricLogs() *MetricsInputMetricLogs {
	this := MetricsInputMetricLogs{}
	return &this
}

// NewMetricsInputMetricLogsWithDefaults instantiates a new MetricsInputMetricLogs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsInputMetricLogsWithDefaults() *MetricsInputMetricLogs {
	this := MetricsInputMetricLogs{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MetricsInputMetricLogs) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsInputMetricLogs) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MetricsInputMetricLogs) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MetricsInputMetricLogs) SetId(v string) {
	o.Id = &v
}

func (o MetricsInputMetricLogs) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricsInputMetricLogs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableMetricsInputMetricLogs struct {
	value *MetricsInputMetricLogs
	isSet bool
}

func (v NullableMetricsInputMetricLogs) Get() *MetricsInputMetricLogs {
	return v.value
}

func (v *NullableMetricsInputMetricLogs) Set(val *MetricsInputMetricLogs) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsInputMetricLogs) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsInputMetricLogs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsInputMetricLogs(val *MetricsInputMetricLogs) *NullableMetricsInputMetricLogs {
	return &NullableMetricsInputMetricLogs{value: val, isSet: true}
}

func (v NullableMetricsInputMetricLogs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsInputMetricLogs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
