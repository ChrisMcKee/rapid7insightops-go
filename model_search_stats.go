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

// checks if the SearchStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchStats{}

// SearchStats It includes the metrics of the executed query.
type SearchStats struct {
	// The time taken to complete the query, in milliseconds.
	DurationMs *int32 `json:"duration_ms,omitempty"`
	// The total number of bytes.
	BytesAll *int32 `json:"bytes_all,omitempty"`
	// The number of bytes checked.
	BytesChecked *int32 `json:"bytes_checked,omitempty"`
	// The total number of events.
	EventsAll *int32 `json:"events_all,omitempty"`
	// The number of events checked.
	EventsChecked *int32 `json:"events_checked,omitempty"`
	// The number of events matched.
	EventsMatched *int32 `json:"events_matched,omitempty"`
	// This shows if the query was able to use our indexing technology to help reduce the query duration. * Value of 0: Indexing could not be used. * Value of 1: Full indexing was achieved to reduce the query duration.
	IndexFactor *float32 `json:"index_factor,omitempty"`
}

// NewSearchStats instantiates a new SearchStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchStats() *SearchStats {
	this := SearchStats{}
	return &this
}

// NewSearchStatsWithDefaults instantiates a new SearchStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchStatsWithDefaults() *SearchStats {
	this := SearchStats{}
	return &this
}

// GetDurationMs returns the DurationMs field value if set, zero value otherwise.
func (o *SearchStats) GetDurationMs() int32 {
	if o == nil || IsNil(o.DurationMs) {
		var ret int32
		return ret
	}
	return *o.DurationMs
}

// GetDurationMsOk returns a tuple with the DurationMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchStats) GetDurationMsOk() (*int32, bool) {
	if o == nil || IsNil(o.DurationMs) {
		return nil, false
	}
	return o.DurationMs, true
}

// HasDurationMs returns a boolean if a field has been set.
func (o *SearchStats) HasDurationMs() bool {
	if o != nil && !IsNil(o.DurationMs) {
		return true
	}

	return false
}

// SetDurationMs gets a reference to the given int32 and assigns it to the DurationMs field.
func (o *SearchStats) SetDurationMs(v int32) {
	o.DurationMs = &v
}

// GetBytesAll returns the BytesAll field value if set, zero value otherwise.
func (o *SearchStats) GetBytesAll() int32 {
	if o == nil || IsNil(o.BytesAll) {
		var ret int32
		return ret
	}
	return *o.BytesAll
}

// GetBytesAllOk returns a tuple with the BytesAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchStats) GetBytesAllOk() (*int32, bool) {
	if o == nil || IsNil(o.BytesAll) {
		return nil, false
	}
	return o.BytesAll, true
}

// HasBytesAll returns a boolean if a field has been set.
func (o *SearchStats) HasBytesAll() bool {
	if o != nil && !IsNil(o.BytesAll) {
		return true
	}

	return false
}

// SetBytesAll gets a reference to the given int32 and assigns it to the BytesAll field.
func (o *SearchStats) SetBytesAll(v int32) {
	o.BytesAll = &v
}

// GetBytesChecked returns the BytesChecked field value if set, zero value otherwise.
func (o *SearchStats) GetBytesChecked() int32 {
	if o == nil || IsNil(o.BytesChecked) {
		var ret int32
		return ret
	}
	return *o.BytesChecked
}

// GetBytesCheckedOk returns a tuple with the BytesChecked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchStats) GetBytesCheckedOk() (*int32, bool) {
	if o == nil || IsNil(o.BytesChecked) {
		return nil, false
	}
	return o.BytesChecked, true
}

// HasBytesChecked returns a boolean if a field has been set.
func (o *SearchStats) HasBytesChecked() bool {
	if o != nil && !IsNil(o.BytesChecked) {
		return true
	}

	return false
}

// SetBytesChecked gets a reference to the given int32 and assigns it to the BytesChecked field.
func (o *SearchStats) SetBytesChecked(v int32) {
	o.BytesChecked = &v
}

// GetEventsAll returns the EventsAll field value if set, zero value otherwise.
func (o *SearchStats) GetEventsAll() int32 {
	if o == nil || IsNil(o.EventsAll) {
		var ret int32
		return ret
	}
	return *o.EventsAll
}

// GetEventsAllOk returns a tuple with the EventsAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchStats) GetEventsAllOk() (*int32, bool) {
	if o == nil || IsNil(o.EventsAll) {
		return nil, false
	}
	return o.EventsAll, true
}

// HasEventsAll returns a boolean if a field has been set.
func (o *SearchStats) HasEventsAll() bool {
	if o != nil && !IsNil(o.EventsAll) {
		return true
	}

	return false
}

// SetEventsAll gets a reference to the given int32 and assigns it to the EventsAll field.
func (o *SearchStats) SetEventsAll(v int32) {
	o.EventsAll = &v
}

// GetEventsChecked returns the EventsChecked field value if set, zero value otherwise.
func (o *SearchStats) GetEventsChecked() int32 {
	if o == nil || IsNil(o.EventsChecked) {
		var ret int32
		return ret
	}
	return *o.EventsChecked
}

// GetEventsCheckedOk returns a tuple with the EventsChecked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchStats) GetEventsCheckedOk() (*int32, bool) {
	if o == nil || IsNil(o.EventsChecked) {
		return nil, false
	}
	return o.EventsChecked, true
}

// HasEventsChecked returns a boolean if a field has been set.
func (o *SearchStats) HasEventsChecked() bool {
	if o != nil && !IsNil(o.EventsChecked) {
		return true
	}

	return false
}

// SetEventsChecked gets a reference to the given int32 and assigns it to the EventsChecked field.
func (o *SearchStats) SetEventsChecked(v int32) {
	o.EventsChecked = &v
}

// GetEventsMatched returns the EventsMatched field value if set, zero value otherwise.
func (o *SearchStats) GetEventsMatched() int32 {
	if o == nil || IsNil(o.EventsMatched) {
		var ret int32
		return ret
	}
	return *o.EventsMatched
}

// GetEventsMatchedOk returns a tuple with the EventsMatched field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchStats) GetEventsMatchedOk() (*int32, bool) {
	if o == nil || IsNil(o.EventsMatched) {
		return nil, false
	}
	return o.EventsMatched, true
}

// HasEventsMatched returns a boolean if a field has been set.
func (o *SearchStats) HasEventsMatched() bool {
	if o != nil && !IsNil(o.EventsMatched) {
		return true
	}

	return false
}

// SetEventsMatched gets a reference to the given int32 and assigns it to the EventsMatched field.
func (o *SearchStats) SetEventsMatched(v int32) {
	o.EventsMatched = &v
}

// GetIndexFactor returns the IndexFactor field value if set, zero value otherwise.
func (o *SearchStats) GetIndexFactor() float32 {
	if o == nil || IsNil(o.IndexFactor) {
		var ret float32
		return ret
	}
	return *o.IndexFactor
}

// GetIndexFactorOk returns a tuple with the IndexFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchStats) GetIndexFactorOk() (*float32, bool) {
	if o == nil || IsNil(o.IndexFactor) {
		return nil, false
	}
	return o.IndexFactor, true
}

// HasIndexFactor returns a boolean if a field has been set.
func (o *SearchStats) HasIndexFactor() bool {
	if o != nil && !IsNil(o.IndexFactor) {
		return true
	}

	return false
}

// SetIndexFactor gets a reference to the given float32 and assigns it to the IndexFactor field.
func (o *SearchStats) SetIndexFactor(v float32) {
	o.IndexFactor = &v
}

func (o SearchStats) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchStats) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.IndexFactor) {
		toSerialize["index_factor"] = o.IndexFactor
	}
	return toSerialize, nil
}

type NullableSearchStats struct {
	value *SearchStats
	isSet bool
}

func (v NullableSearchStats) Get() *SearchStats {
	return v.value
}

func (v *NullableSearchStats) Set(val *SearchStats) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchStats) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchStats(val *SearchStats) *NullableSearchStats {
	return &NullableSearchStats{value: val, isSet: true}
}

func (v NullableSearchStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}