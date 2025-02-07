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

// checks if the MetricsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricsResponse{}

// MetricsResponse struct for MetricsResponse
type MetricsResponse struct {
	// The UUID of the pre-computed query.
	Id *string `json:"id,omitempty"`
	// The information of the logs used in the pre-computed query.
	Logs []LogInfoResponse `json:"logs,omitempty"`
	// The information on the log sets used in the pre-computed query.
	Logsets    []LogsetInfo       `json:"logsets,omitempty"`
	Leql       *LeqlMetrics       `json:"leql,omitempty"`
	Statistics *StatisticsMetrics `json:"statistics,omitempty"`
}

// NewMetricsResponse instantiates a new MetricsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsResponse() *MetricsResponse {
	this := MetricsResponse{}
	return &this
}

// NewMetricsResponseWithDefaults instantiates a new MetricsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsResponseWithDefaults() *MetricsResponse {
	this := MetricsResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MetricsResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MetricsResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MetricsResponse) SetId(v string) {
	o.Id = &v
}

// GetLogs returns the Logs field value if set, zero value otherwise.
func (o *MetricsResponse) GetLogs() []LogInfoResponse {
	if o == nil || IsNil(o.Logs) {
		var ret []LogInfoResponse
		return ret
	}
	return o.Logs
}

// GetLogsOk returns a tuple with the Logs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsResponse) GetLogsOk() ([]LogInfoResponse, bool) {
	if o == nil || IsNil(o.Logs) {
		return nil, false
	}
	return o.Logs, true
}

// HasLogs returns a boolean if a field has been set.
func (o *MetricsResponse) HasLogs() bool {
	if o != nil && !IsNil(o.Logs) {
		return true
	}

	return false
}

// SetLogs gets a reference to the given []LogInfoResponse and assigns it to the Logs field.
func (o *MetricsResponse) SetLogs(v []LogInfoResponse) {
	o.Logs = v
}

// GetLogsets returns the Logsets field value if set, zero value otherwise.
func (o *MetricsResponse) GetLogsets() []LogsetInfo {
	if o == nil || IsNil(o.Logsets) {
		var ret []LogsetInfo
		return ret
	}
	return o.Logsets
}

// GetLogsetsOk returns a tuple with the Logsets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsResponse) GetLogsetsOk() ([]LogsetInfo, bool) {
	if o == nil || IsNil(o.Logsets) {
		return nil, false
	}
	return o.Logsets, true
}

// HasLogsets returns a boolean if a field has been set.
func (o *MetricsResponse) HasLogsets() bool {
	if o != nil && !IsNil(o.Logsets) {
		return true
	}

	return false
}

// SetLogsets gets a reference to the given []LogsetInfo and assigns it to the Logsets field.
func (o *MetricsResponse) SetLogsets(v []LogsetInfo) {
	o.Logsets = v
}

// GetLeql returns the Leql field value if set, zero value otherwise.
func (o *MetricsResponse) GetLeql() LeqlMetrics {
	if o == nil || IsNil(o.Leql) {
		var ret LeqlMetrics
		return ret
	}
	return *o.Leql
}

// GetLeqlOk returns a tuple with the Leql field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsResponse) GetLeqlOk() (*LeqlMetrics, bool) {
	if o == nil || IsNil(o.Leql) {
		return nil, false
	}
	return o.Leql, true
}

// HasLeql returns a boolean if a field has been set.
func (o *MetricsResponse) HasLeql() bool {
	if o != nil && !IsNil(o.Leql) {
		return true
	}

	return false
}

// SetLeql gets a reference to the given LeqlMetrics and assigns it to the Leql field.
func (o *MetricsResponse) SetLeql(v LeqlMetrics) {
	o.Leql = &v
}

// GetStatistics returns the Statistics field value if set, zero value otherwise.
func (o *MetricsResponse) GetStatistics() StatisticsMetrics {
	if o == nil || IsNil(o.Statistics) {
		var ret StatisticsMetrics
		return ret
	}
	return *o.Statistics
}

// GetStatisticsOk returns a tuple with the Statistics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsResponse) GetStatisticsOk() (*StatisticsMetrics, bool) {
	if o == nil || IsNil(o.Statistics) {
		return nil, false
	}
	return o.Statistics, true
}

// HasStatistics returns a boolean if a field has been set.
func (o *MetricsResponse) HasStatistics() bool {
	if o != nil && !IsNil(o.Statistics) {
		return true
	}

	return false
}

// SetStatistics gets a reference to the given StatisticsMetrics and assigns it to the Statistics field.
func (o *MetricsResponse) SetStatistics(v StatisticsMetrics) {
	o.Statistics = &v
}

func (o MetricsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Logs) {
		toSerialize["logs"] = o.Logs
	}
	if !IsNil(o.Logsets) {
		toSerialize["logsets"] = o.Logsets
	}
	if !IsNil(o.Leql) {
		toSerialize["leql"] = o.Leql
	}
	if !IsNil(o.Statistics) {
		toSerialize["statistics"] = o.Statistics
	}
	return toSerialize, nil
}

type NullableMetricsResponse struct {
	value *MetricsResponse
	isSet bool
}

func (v NullableMetricsResponse) Get() *MetricsResponse {
	return v.value
}

func (v *NullableMetricsResponse) Set(val *MetricsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsResponse(val *MetricsResponse) *NullableMetricsResponse {
	return &NullableMetricsResponse{value: val, isSet: true}
}

func (v NullableMetricsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
