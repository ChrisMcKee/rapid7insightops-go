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

// checks if the Stats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Stats{}

// Stats struct for Stats
type Stats struct {
	Average           *float32 `json:"average,omitempty"`
	Min               *float32 `json:"min,omitempty"`
	Max               *float32 `json:"max,omitempty"`
	Median            *float32 `json:"median,omitempty"`
	Percentile1To99   *float32 `json:"percentile_<1 to 99>*,omitempty"`
	Count             *float32 `json:"count,omitempty"`
	Sum               *float32 `json:"sum,omitempty"`
	Standarddeviation *float32 `json:"standarddeviation,omitempty"`
	Unique            *float32 `json:"unique,omitempty"`
	Bytes             *float32 `json:"bytes,omitempty"`
}

// NewStats instantiates a new Stats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStats() *Stats {
	this := Stats{}
	return &this
}

// NewStatsWithDefaults instantiates a new Stats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatsWithDefaults() *Stats {
	this := Stats{}
	return &this
}

// GetAverage returns the Average field value if set, zero value otherwise.
func (o *Stats) GetAverage() float32 {
	if o == nil || IsNil(o.Average) {
		var ret float32
		return ret
	}
	return *o.Average
}

// GetAverageOk returns a tuple with the Average field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stats) GetAverageOk() (*float32, bool) {
	if o == nil || IsNil(o.Average) {
		return nil, false
	}
	return o.Average, true
}

// HasAverage returns a boolean if a field has been set.
func (o *Stats) HasAverage() bool {
	if o != nil && !IsNil(o.Average) {
		return true
	}

	return false
}

// SetAverage gets a reference to the given float32 and assigns it to the Average field.
func (o *Stats) SetAverage(v float32) {
	o.Average = &v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *Stats) GetMin() float32 {
	if o == nil || IsNil(o.Min) {
		var ret float32
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stats) GetMinOk() (*float32, bool) {
	if o == nil || IsNil(o.Min) {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *Stats) HasMin() bool {
	if o != nil && !IsNil(o.Min) {
		return true
	}

	return false
}

// SetMin gets a reference to the given float32 and assigns it to the Min field.
func (o *Stats) SetMin(v float32) {
	o.Min = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *Stats) GetMax() float32 {
	if o == nil || IsNil(o.Max) {
		var ret float32
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stats) GetMaxOk() (*float32, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *Stats) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given float32 and assigns it to the Max field.
func (o *Stats) SetMax(v float32) {
	o.Max = &v
}

// GetMedian returns the Median field value if set, zero value otherwise.
func (o *Stats) GetMedian() float32 {
	if o == nil || IsNil(o.Median) {
		var ret float32
		return ret
	}
	return *o.Median
}

// GetMedianOk returns a tuple with the Median field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stats) GetMedianOk() (*float32, bool) {
	if o == nil || IsNil(o.Median) {
		return nil, false
	}
	return o.Median, true
}

// HasMedian returns a boolean if a field has been set.
func (o *Stats) HasMedian() bool {
	if o != nil && !IsNil(o.Median) {
		return true
	}

	return false
}

// SetMedian gets a reference to the given float32 and assigns it to the Median field.
func (o *Stats) SetMedian(v float32) {
	o.Median = &v
}

// GetPercentile1To99 returns the Percentile1To99 field value if set, zero value otherwise.
func (o *Stats) GetPercentile1To99() float32 {
	if o == nil || IsNil(o.Percentile1To99) {
		var ret float32
		return ret
	}
	return *o.Percentile1To99
}

// GetPercentile1To99Ok returns a tuple with the Percentile1To99 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stats) GetPercentile1To99Ok() (*float32, bool) {
	if o == nil || IsNil(o.Percentile1To99) {
		return nil, false
	}
	return o.Percentile1To99, true
}

// HasPercentile1To99 returns a boolean if a field has been set.
func (o *Stats) HasPercentile1To99() bool {
	if o != nil && !IsNil(o.Percentile1To99) {
		return true
	}

	return false
}

// SetPercentile1To99 gets a reference to the given float32 and assigns it to the Percentile1To99 field.
func (o *Stats) SetPercentile1To99(v float32) {
	o.Percentile1To99 = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *Stats) GetCount() float32 {
	if o == nil || IsNil(o.Count) {
		var ret float32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stats) GetCountOk() (*float32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *Stats) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given float32 and assigns it to the Count field.
func (o *Stats) SetCount(v float32) {
	o.Count = &v
}

// GetSum returns the Sum field value if set, zero value otherwise.
func (o *Stats) GetSum() float32 {
	if o == nil || IsNil(o.Sum) {
		var ret float32
		return ret
	}
	return *o.Sum
}

// GetSumOk returns a tuple with the Sum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stats) GetSumOk() (*float32, bool) {
	if o == nil || IsNil(o.Sum) {
		return nil, false
	}
	return o.Sum, true
}

// HasSum returns a boolean if a field has been set.
func (o *Stats) HasSum() bool {
	if o != nil && !IsNil(o.Sum) {
		return true
	}

	return false
}

// SetSum gets a reference to the given float32 and assigns it to the Sum field.
func (o *Stats) SetSum(v float32) {
	o.Sum = &v
}

// GetStandarddeviation returns the Standarddeviation field value if set, zero value otherwise.
func (o *Stats) GetStandarddeviation() float32 {
	if o == nil || IsNil(o.Standarddeviation) {
		var ret float32
		return ret
	}
	return *o.Standarddeviation
}

// GetStandarddeviationOk returns a tuple with the Standarddeviation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stats) GetStandarddeviationOk() (*float32, bool) {
	if o == nil || IsNil(o.Standarddeviation) {
		return nil, false
	}
	return o.Standarddeviation, true
}

// HasStandarddeviation returns a boolean if a field has been set.
func (o *Stats) HasStandarddeviation() bool {
	if o != nil && !IsNil(o.Standarddeviation) {
		return true
	}

	return false
}

// SetStandarddeviation gets a reference to the given float32 and assigns it to the Standarddeviation field.
func (o *Stats) SetStandarddeviation(v float32) {
	o.Standarddeviation = &v
}

// GetUnique returns the Unique field value if set, zero value otherwise.
func (o *Stats) GetUnique() float32 {
	if o == nil || IsNil(o.Unique) {
		var ret float32
		return ret
	}
	return *o.Unique
}

// GetUniqueOk returns a tuple with the Unique field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stats) GetUniqueOk() (*float32, bool) {
	if o == nil || IsNil(o.Unique) {
		return nil, false
	}
	return o.Unique, true
}

// HasUnique returns a boolean if a field has been set.
func (o *Stats) HasUnique() bool {
	if o != nil && !IsNil(o.Unique) {
		return true
	}

	return false
}

// SetUnique gets a reference to the given float32 and assigns it to the Unique field.
func (o *Stats) SetUnique(v float32) {
	o.Unique = &v
}

// GetBytes returns the Bytes field value if set, zero value otherwise.
func (o *Stats) GetBytes() float32 {
	if o == nil || IsNil(o.Bytes) {
		var ret float32
		return ret
	}
	return *o.Bytes
}

// GetBytesOk returns a tuple with the Bytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stats) GetBytesOk() (*float32, bool) {
	if o == nil || IsNil(o.Bytes) {
		return nil, false
	}
	return o.Bytes, true
}

// HasBytes returns a boolean if a field has been set.
func (o *Stats) HasBytes() bool {
	if o != nil && !IsNil(o.Bytes) {
		return true
	}

	return false
}

// SetBytes gets a reference to the given float32 and assigns it to the Bytes field.
func (o *Stats) SetBytes(v float32) {
	o.Bytes = &v
}

func (o Stats) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Stats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Average) {
		toSerialize["average"] = o.Average
	}
	if !IsNil(o.Min) {
		toSerialize["min"] = o.Min
	}
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	if !IsNil(o.Median) {
		toSerialize["median"] = o.Median
	}
	if !IsNil(o.Percentile1To99) {
		toSerialize["percentile_<1 to 99>*"] = o.Percentile1To99
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Sum) {
		toSerialize["sum"] = o.Sum
	}
	if !IsNil(o.Standarddeviation) {
		toSerialize["standarddeviation"] = o.Standarddeviation
	}
	if !IsNil(o.Unique) {
		toSerialize["unique"] = o.Unique
	}
	if !IsNil(o.Bytes) {
		toSerialize["bytes"] = o.Bytes
	}
	return toSerialize, nil
}

type NullableStats struct {
	value *Stats
	isSet bool
}

func (v NullableStats) Get() *Stats {
	return v.value
}

func (v *NullableStats) Set(val *Stats) {
	v.value = val
	v.isSet = true
}

func (v NullableStats) IsSet() bool {
	return v.isSet
}

func (v *NullableStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStats(val *Stats) *NullableStats {
	return &NullableStats{value: val, isSet: true}
}

func (v NullableStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
