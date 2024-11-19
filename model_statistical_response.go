/*
InsightOps REST API

### Overview  Our REST API lets you use InsightOps through HTTP requests. Currently, the REST API allows you to perform the majority of the actions available through the UI, and has some additional functionality that is not available through the UI. You may use this API to automate common tasks (for example, via shell scripts), and to generally interact with InsightOps programmatically.  This page precisely describes the REST API and serves as a reference for the API. Each HTTP method and each URL endpoint is documented in a self-contained unit so that users only need to read about the HTTP methods relevant to them.  ### Terminology  * A **log entry** is an individual log event. * A **log** is a collection of log entries, or a single log stream. * A **log set** is a logical-only collection of logs, i.e. logs can be in multiple logsets and deleting a logset only deletes the relation between logs, not the logs themselves. * [Log Entry Query Language](https://docs.rapid7.com/insightops/log-search) (**LEQL**) is the query language used in Insight Ops to search log data.

API version: latest
Contact: support@rapid7.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package insightops

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the StatisticalResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatisticalResponse{}

// StatisticalResponse struct for StatisticalResponse
type StatisticalResponse struct {
	// The log keys of the logs which the query is run against.
	Logs        []string    `json:"logs"`
	Leql        Leql        `json:"leql"`
	Statistics  Statistics  `json:"statistics"`
	SearchStats SearchStats `json:"search_stats"`
}

type _StatisticalResponse StatisticalResponse

// NewStatisticalResponse instantiates a new StatisticalResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatisticalResponse(logs []string, leql Leql, statistics Statistics, searchStats SearchStats) *StatisticalResponse {
	this := StatisticalResponse{}
	this.Logs = logs
	this.Leql = leql
	this.Statistics = statistics
	this.SearchStats = searchStats
	return &this
}

// NewStatisticalResponseWithDefaults instantiates a new StatisticalResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatisticalResponseWithDefaults() *StatisticalResponse {
	this := StatisticalResponse{}
	return &this
}

// GetLogs returns the Logs field value
func (o *StatisticalResponse) GetLogs() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Logs
}

// GetLogsOk returns a tuple with the Logs field value
// and a boolean to check if the value has been set.
func (o *StatisticalResponse) GetLogsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Logs, true
}

// SetLogs sets field value
func (o *StatisticalResponse) SetLogs(v []string) {
	o.Logs = v
}

// GetLeql returns the Leql field value
func (o *StatisticalResponse) GetLeql() Leql {
	if o == nil {
		var ret Leql
		return ret
	}

	return o.Leql
}

// GetLeqlOk returns a tuple with the Leql field value
// and a boolean to check if the value has been set.
func (o *StatisticalResponse) GetLeqlOk() (*Leql, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Leql, true
}

// SetLeql sets field value
func (o *StatisticalResponse) SetLeql(v Leql) {
	o.Leql = v
}

// GetStatistics returns the Statistics field value
func (o *StatisticalResponse) GetStatistics() Statistics {
	if o == nil {
		var ret Statistics
		return ret
	}

	return o.Statistics
}

// GetStatisticsOk returns a tuple with the Statistics field value
// and a boolean to check if the value has been set.
func (o *StatisticalResponse) GetStatisticsOk() (*Statistics, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Statistics, true
}

// SetStatistics sets field value
func (o *StatisticalResponse) SetStatistics(v Statistics) {
	o.Statistics = v
}

// GetSearchStats returns the SearchStats field value
func (o *StatisticalResponse) GetSearchStats() SearchStats {
	if o == nil {
		var ret SearchStats
		return ret
	}

	return o.SearchStats
}

// GetSearchStatsOk returns a tuple with the SearchStats field value
// and a boolean to check if the value has been set.
func (o *StatisticalResponse) GetSearchStatsOk() (*SearchStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SearchStats, true
}

// SetSearchStats sets field value
func (o *StatisticalResponse) SetSearchStats(v SearchStats) {
	o.SearchStats = v
}

func (o StatisticalResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatisticalResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["logs"] = o.Logs
	toSerialize["leql"] = o.Leql
	toSerialize["statistics"] = o.Statistics
	toSerialize["search_stats"] = o.SearchStats
	return toSerialize, nil
}

func (o *StatisticalResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"logs",
		"leql",
		"statistics",
		"search_stats",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varStatisticalResponse := _StatisticalResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStatisticalResponse)

	if err != nil {
		return err
	}

	*o = StatisticalResponse(varStatisticalResponse)

	return err
}

type NullableStatisticalResponse struct {
	value *StatisticalResponse
	isSet bool
}

func (v NullableStatisticalResponse) Get() *StatisticalResponse {
	return v.value
}

func (v *NullableStatisticalResponse) Set(val *StatisticalResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStatisticalResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStatisticalResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatisticalResponse(val *StatisticalResponse) *NullableStatisticalResponse {
	return &NullableStatisticalResponse{value: val, isSet: true}
}

func (v NullableStatisticalResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatisticalResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}