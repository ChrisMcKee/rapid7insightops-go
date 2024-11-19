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

// checks if the SearchStatsInnerStatistics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchStatsInnerStatistics{}

// SearchStatsInnerStatistics struct for SearchStatsInnerStatistics
type SearchStatsInnerStatistics struct {
	// The time taken to complete the query, in milliseconds.
	DurationMs *int32 `json:"duration_ms,omitempty"`
	// The total amount of log data queried in bytes.
	BytesAll *int32 `json:"bytes_all,omitempty"`
	// The amount of log data that was scanned by this query in bytes. This amount may be much less than the total amount of log data that is queried. This is a result of data indexing performed by Log Search when the LEQL statement matches a small subset of the data.
	BytesChecked *int32 `json:"bytes_checked,omitempty"`
	// The total number of log entries queried.
	EventsAll *int32 `json:"events_all,omitempty"`
	// The number of log entries that were scanned by this query. This amount may be much less than the total amount of log data that is queried. This is a result of data indexing performed by Log Search when the LEQL statement matches a small subset of the data.
	EventsChecked *int32 `json:"events_checked,omitempty"`
	// The number of log entries that matched the LEQL statement of the query.
	EventsMatched *int32 `json:"events_matched,omitempty"`
}

// NewSearchStatsInnerStatistics instantiates a new SearchStatsInnerStatistics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchStatsInnerStatistics() *SearchStatsInnerStatistics {
	this := SearchStatsInnerStatistics{}
	return &this
}

// NewSearchStatsInnerStatisticsWithDefaults instantiates a new SearchStatsInnerStatistics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchStatsInnerStatisticsWithDefaults() *SearchStatsInnerStatistics {
	this := SearchStatsInnerStatistics{}
	return &this
}

// GetDurationMs returns the DurationMs field value if set, zero value otherwise.
func (o *SearchStatsInnerStatistics) GetDurationMs() int32 {
	if o == nil || IsNil(o.DurationMs) {
		var ret int32
		return ret
	}
	return *o.DurationMs
}

// GetDurationMsOk returns a tuple with the DurationMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchStatsInnerStatistics) GetDurationMsOk() (*int32, bool) {
	if o == nil || IsNil(o.DurationMs) {
		return nil, false
	}
	return o.DurationMs, true
}

// HasDurationMs returns a boolean if a field has been set.
func (o *SearchStatsInnerStatistics) HasDurationMs() bool {
	if o != nil && !IsNil(o.DurationMs) {
		return true
	}

	return false
}

// SetDurationMs gets a reference to the given int32 and assigns it to the DurationMs field.
func (o *SearchStatsInnerStatistics) SetDurationMs(v int32) {
	o.DurationMs = &v
}

// GetBytesAll returns the BytesAll field value if set, zero value otherwise.
func (o *SearchStatsInnerStatistics) GetBytesAll() int32 {
	if o == nil || IsNil(o.BytesAll) {
		var ret int32
		return ret
	}
	return *o.BytesAll
}

// GetBytesAllOk returns a tuple with the BytesAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchStatsInnerStatistics) GetBytesAllOk() (*int32, bool) {
	if o == nil || IsNil(o.BytesAll) {
		return nil, false
	}
	return o.BytesAll, true
}

// HasBytesAll returns a boolean if a field has been set.
func (o *SearchStatsInnerStatistics) HasBytesAll() bool {
	if o != nil && !IsNil(o.BytesAll) {
		return true
	}

	return false
}

// SetBytesAll gets a reference to the given int32 and assigns it to the BytesAll field.
func (o *SearchStatsInnerStatistics) SetBytesAll(v int32) {
	o.BytesAll = &v
}

// GetBytesChecked returns the BytesChecked field value if set, zero value otherwise.
func (o *SearchStatsInnerStatistics) GetBytesChecked() int32 {
	if o == nil || IsNil(o.BytesChecked) {
		var ret int32
		return ret
	}
	return *o.BytesChecked
}

// GetBytesCheckedOk returns a tuple with the BytesChecked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchStatsInnerStatistics) GetBytesCheckedOk() (*int32, bool) {
	if o == nil || IsNil(o.BytesChecked) {
		return nil, false
	}
	return o.BytesChecked, true
}

// HasBytesChecked returns a boolean if a field has been set.
func (o *SearchStatsInnerStatistics) HasBytesChecked() bool {
	if o != nil && !IsNil(o.BytesChecked) {
		return true
	}

	return false
}

// SetBytesChecked gets a reference to the given int32 and assigns it to the BytesChecked field.
func (o *SearchStatsInnerStatistics) SetBytesChecked(v int32) {
	o.BytesChecked = &v
}

// GetEventsAll returns the EventsAll field value if set, zero value otherwise.
func (o *SearchStatsInnerStatistics) GetEventsAll() int32 {
	if o == nil || IsNil(o.EventsAll) {
		var ret int32
		return ret
	}
	return *o.EventsAll
}

// GetEventsAllOk returns a tuple with the EventsAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchStatsInnerStatistics) GetEventsAllOk() (*int32, bool) {
	if o == nil || IsNil(o.EventsAll) {
		return nil, false
	}
	return o.EventsAll, true
}

// HasEventsAll returns a boolean if a field has been set.
func (o *SearchStatsInnerStatistics) HasEventsAll() bool {
	if o != nil && !IsNil(o.EventsAll) {
		return true
	}

	return false
}

// SetEventsAll gets a reference to the given int32 and assigns it to the EventsAll field.
func (o *SearchStatsInnerStatistics) SetEventsAll(v int32) {
	o.EventsAll = &v
}

// GetEventsChecked returns the EventsChecked field value if set, zero value otherwise.
func (o *SearchStatsInnerStatistics) GetEventsChecked() int32 {
	if o == nil || IsNil(o.EventsChecked) {
		var ret int32
		return ret
	}
	return *o.EventsChecked
}

// GetEventsCheckedOk returns a tuple with the EventsChecked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchStatsInnerStatistics) GetEventsCheckedOk() (*int32, bool) {
	if o == nil || IsNil(o.EventsChecked) {
		return nil, false
	}
	return o.EventsChecked, true
}

// HasEventsChecked returns a boolean if a field has been set.
func (o *SearchStatsInnerStatistics) HasEventsChecked() bool {
	if o != nil && !IsNil(o.EventsChecked) {
		return true
	}

	return false
}

// SetEventsChecked gets a reference to the given int32 and assigns it to the EventsChecked field.
func (o *SearchStatsInnerStatistics) SetEventsChecked(v int32) {
	o.EventsChecked = &v
}

// GetEventsMatched returns the EventsMatched field value if set, zero value otherwise.
func (o *SearchStatsInnerStatistics) GetEventsMatched() int32 {
	if o == nil || IsNil(o.EventsMatched) {
		var ret int32
		return ret
	}
	return *o.EventsMatched
}

// GetEventsMatchedOk returns a tuple with the EventsMatched field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchStatsInnerStatistics) GetEventsMatchedOk() (*int32, bool) {
	if o == nil || IsNil(o.EventsMatched) {
		return nil, false
	}
	return o.EventsMatched, true
}

// HasEventsMatched returns a boolean if a field has been set.
func (o *SearchStatsInnerStatistics) HasEventsMatched() bool {
	if o != nil && !IsNil(o.EventsMatched) {
		return true
	}

	return false
}

// SetEventsMatched gets a reference to the given int32 and assigns it to the EventsMatched field.
func (o *SearchStatsInnerStatistics) SetEventsMatched(v int32) {
	o.EventsMatched = &v
}

func (o SearchStatsInnerStatistics) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchStatsInnerStatistics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DurationMs) {
		toSerialize["duration_ms"] = o.DurationMs
	}
	if !IsNil(o.BytesAll) {
		toSerialize["bytes_all"] = o.BytesAll
	}
	if !IsNil(o.BytesChecked) {
		toSerialize["bytes_checked"] = o.BytesChecked
	}
	if !IsNil(o.EventsAll) {
		toSerialize["events_all"] = o.EventsAll
	}
	if !IsNil(o.EventsChecked) {
		toSerialize["events_checked"] = o.EventsChecked
	}
	if !IsNil(o.EventsMatched) {
		toSerialize["events_matched"] = o.EventsMatched
	}
	return toSerialize, nil
}

type NullableSearchStatsInnerStatistics struct {
	value *SearchStatsInnerStatistics
	isSet bool
}

func (v NullableSearchStatsInnerStatistics) Get() *SearchStatsInnerStatistics {
	return v.value
}

func (v *NullableSearchStatsInnerStatistics) Set(val *SearchStatsInnerStatistics) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchStatsInnerStatistics) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchStatsInnerStatistics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchStatsInnerStatistics(val *SearchStatsInnerStatistics) *NullableSearchStatsInnerStatistics {
	return &NullableSearchStatsInnerStatistics{value: val, isSet: true}
}

func (v NullableSearchStatsInnerStatistics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchStatsInnerStatistics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}