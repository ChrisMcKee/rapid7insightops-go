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

// checks if the CreateScheduledQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateScheduledQuery{}

// CreateScheduledQuery The configurations for the two queries that are executed at regular 10-minute intervals. Notifications will be triggered when there is a large enough difference between the two (defined by the `threshold_value` parameter).
type CreateScheduledQuery struct {
	// The name of the scheduled query.
	Name string      `json:"name"`
	Leql AnomalyLeql `json:"leql"`
	// Defines the time range for the two queries, along with the `time_period` parameter. Two queries will run: the time range of the first, more recent query, will end at the current system time; the time range of the older query will end at the current system time minus the time specified by the `timeframe_value` and `timeframe_period` parameters.
	TimeValue int32 `json:"time_value"`
	// Defines the time range for the two queries, along with the `time_value` parameter.
	TimePeriod string `json:"time_period"`
	// Defines the \"offset\" for the older of the two queries along with the `timeframe_period` parameter. For example, if the `timeframe_value` is set to 1 and `timeframe_period` is set to `Week`, then the time range for the older of the two queries will end at the current system time minus 1 week.
	TimeframeValue int32 `json:"timeframe_value"`
	// Defines the \"offset\" for the older of the two queries, along with the `timeframe_value` parameter.
	TimeframePeriod string `json:"timeframe_period"`
}

type _CreateScheduledQuery CreateScheduledQuery

// NewCreateScheduledQuery instantiates a new CreateScheduledQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateScheduledQuery(name string, leql AnomalyLeql, timeValue int32, timePeriod string, timeframeValue int32, timeframePeriod string) *CreateScheduledQuery {
	this := CreateScheduledQuery{}
	this.Name = name
	this.Leql = leql
	this.TimeValue = timeValue
	this.TimePeriod = timePeriod
	this.TimeframeValue = timeframeValue
	this.TimeframePeriod = timeframePeriod
	return &this
}

// NewCreateScheduledQueryWithDefaults instantiates a new CreateScheduledQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateScheduledQueryWithDefaults() *CreateScheduledQuery {
	this := CreateScheduledQuery{}
	return &this
}

// GetName returns the Name field value
func (o *CreateScheduledQuery) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateScheduledQuery) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateScheduledQuery) SetName(v string) {
	o.Name = v
}

// GetLeql returns the Leql field value
func (o *CreateScheduledQuery) GetLeql() AnomalyLeql {
	if o == nil {
		var ret AnomalyLeql
		return ret
	}

	return o.Leql
}

// GetLeqlOk returns a tuple with the Leql field value
// and a boolean to check if the value has been set.
func (o *CreateScheduledQuery) GetLeqlOk() (*AnomalyLeql, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Leql, true
}

// SetLeql sets field value
func (o *CreateScheduledQuery) SetLeql(v AnomalyLeql) {
	o.Leql = v
}

// GetTimeValue returns the TimeValue field value
func (o *CreateScheduledQuery) GetTimeValue() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TimeValue
}

// GetTimeValueOk returns a tuple with the TimeValue field value
// and a boolean to check if the value has been set.
func (o *CreateScheduledQuery) GetTimeValueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeValue, true
}

// SetTimeValue sets field value
func (o *CreateScheduledQuery) SetTimeValue(v int32) {
	o.TimeValue = v
}

// GetTimePeriod returns the TimePeriod field value
func (o *CreateScheduledQuery) GetTimePeriod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimePeriod
}

// GetTimePeriodOk returns a tuple with the TimePeriod field value
// and a boolean to check if the value has been set.
func (o *CreateScheduledQuery) GetTimePeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimePeriod, true
}

// SetTimePeriod sets field value
func (o *CreateScheduledQuery) SetTimePeriod(v string) {
	o.TimePeriod = v
}

// GetTimeframeValue returns the TimeframeValue field value
func (o *CreateScheduledQuery) GetTimeframeValue() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TimeframeValue
}

// GetTimeframeValueOk returns a tuple with the TimeframeValue field value
// and a boolean to check if the value has been set.
func (o *CreateScheduledQuery) GetTimeframeValueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeframeValue, true
}

// SetTimeframeValue sets field value
func (o *CreateScheduledQuery) SetTimeframeValue(v int32) {
	o.TimeframeValue = v
}

// GetTimeframePeriod returns the TimeframePeriod field value
func (o *CreateScheduledQuery) GetTimeframePeriod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeframePeriod
}

// GetTimeframePeriodOk returns a tuple with the TimeframePeriod field value
// and a boolean to check if the value has been set.
func (o *CreateScheduledQuery) GetTimeframePeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeframePeriod, true
}

// SetTimeframePeriod sets field value
func (o *CreateScheduledQuery) SetTimeframePeriod(v string) {
	o.TimeframePeriod = v
}

func (o CreateScheduledQuery) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateScheduledQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["leql"] = o.Leql
	toSerialize["time_value"] = o.TimeValue
	toSerialize["time_period"] = o.TimePeriod
	toSerialize["timeframe_value"] = o.TimeframeValue
	toSerialize["timeframe_period"] = o.TimeframePeriod
	return toSerialize, nil
}

func (o *CreateScheduledQuery) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"leql",
		"time_value",
		"time_period",
		"timeframe_value",
		"timeframe_period",
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

	varCreateScheduledQuery := _CreateScheduledQuery{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateScheduledQuery)

	if err != nil {
		return err
	}

	*o = CreateScheduledQuery(varCreateScheduledQuery)

	return err
}

type NullableCreateScheduledQuery struct {
	value *CreateScheduledQuery
	isSet bool
}

func (v NullableCreateScheduledQuery) Get() *CreateScheduledQuery {
	return v.value
}

func (v *NullableCreateScheduledQuery) Set(val *CreateScheduledQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateScheduledQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateScheduledQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateScheduledQuery(val *CreateScheduledQuery) *NullableCreateScheduledQuery {
	return &NullableCreateScheduledQuery{value: val, isSet: true}
}

func (v NullableCreateScheduledQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateScheduledQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
