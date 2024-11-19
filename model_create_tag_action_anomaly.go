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

// checks if the CreateTagActionAnomaly type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTagActionAnomaly{}

// CreateTagActionAnomaly struct for CreateTagActionAnomaly
type CreateTagActionAnomaly struct {
	// \\ When set to true, the notification is enabled. When set to false, the notification is disabled.
	Enabled bool `json:"enabled"`
	// The maximum number of notifications that will be triggered within a window of time, controlled by the `min_report_period` parameter.
	MinReportCount  int32      `json:"min_report_count"`
	MinReportPeriod PeriodEnum `json:"min_report_period"`
	// The list of notification targets for this notification.
	Targets []CreateTagTarget `json:"targets"`
	// Always set to \"Alert\".
	Type *string `json:"type,omitempty"`
}

type _CreateTagActionAnomaly CreateTagActionAnomaly

// NewCreateTagActionAnomaly instantiates a new CreateTagActionAnomaly object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTagActionAnomaly(enabled bool, minReportCount int32, minReportPeriod PeriodEnum, targets []CreateTagTarget) *CreateTagActionAnomaly {
	this := CreateTagActionAnomaly{}
	this.Enabled = enabled
	this.MinReportCount = minReportCount
	this.MinReportPeriod = minReportPeriod
	this.Targets = targets
	return &this
}

// NewCreateTagActionAnomalyWithDefaults instantiates a new CreateTagActionAnomaly object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTagActionAnomalyWithDefaults() *CreateTagActionAnomaly {
	this := CreateTagActionAnomaly{}
	var minReportCount int32 = 1
	this.MinReportCount = minReportCount
	var minReportPeriod PeriodEnum = PeriodEnum(HOUR)
	this.MinReportPeriod = minReportPeriod
	return &this
}

// GetEnabled returns the Enabled field value
func (o *CreateTagActionAnomaly) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *CreateTagActionAnomaly) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *CreateTagActionAnomaly) SetEnabled(v bool) {
	o.Enabled = v
}

// GetMinReportCount returns the MinReportCount field value
func (o *CreateTagActionAnomaly) GetMinReportCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MinReportCount
}

// GetMinReportCountOk returns a tuple with the MinReportCount field value
// and a boolean to check if the value has been set.
func (o *CreateTagActionAnomaly) GetMinReportCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinReportCount, true
}

// SetMinReportCount sets field value
func (o *CreateTagActionAnomaly) SetMinReportCount(v int32) {
	o.MinReportCount = v
}

// GetMinReportPeriod returns the MinReportPeriod field value
func (o *CreateTagActionAnomaly) GetMinReportPeriod() PeriodEnum {
	if o == nil {
		var ret PeriodEnum
		return ret
	}

	return o.MinReportPeriod
}

// GetMinReportPeriodOk returns a tuple with the MinReportPeriod field value
// and a boolean to check if the value has been set.
func (o *CreateTagActionAnomaly) GetMinReportPeriodOk() (*PeriodEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinReportPeriod, true
}

// SetMinReportPeriod sets field value
func (o *CreateTagActionAnomaly) SetMinReportPeriod(v PeriodEnum) {
	o.MinReportPeriod = v
}

// GetTargets returns the Targets field value
func (o *CreateTagActionAnomaly) GetTargets() []CreateTagTarget {
	if o == nil {
		var ret []CreateTagTarget
		return ret
	}

	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value
// and a boolean to check if the value has been set.
func (o *CreateTagActionAnomaly) GetTargetsOk() ([]CreateTagTarget, bool) {
	if o == nil {
		return nil, false
	}
	return o.Targets, true
}

// SetTargets sets field value
func (o *CreateTagActionAnomaly) SetTargets(v []CreateTagTarget) {
	o.Targets = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateTagActionAnomaly) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTagActionAnomaly) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateTagActionAnomaly) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CreateTagActionAnomaly) SetType(v string) {
	o.Type = &v
}

func (o CreateTagActionAnomaly) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTagActionAnomaly) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	toSerialize["min_report_count"] = o.MinReportCount
	toSerialize["min_report_period"] = o.MinReportPeriod
	toSerialize["targets"] = o.Targets
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

func (o *CreateTagActionAnomaly) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enabled",
		"min_report_count",
		"min_report_period",
		"targets",
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

	varCreateTagActionAnomaly := _CreateTagActionAnomaly{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateTagActionAnomaly)

	if err != nil {
		return err
	}

	*o = CreateTagActionAnomaly(varCreateTagActionAnomaly)

	return err
}

type NullableCreateTagActionAnomaly struct {
	value *CreateTagActionAnomaly
	isSet bool
}

func (v NullableCreateTagActionAnomaly) Get() *CreateTagActionAnomaly {
	return v.value
}

func (v *NullableCreateTagActionAnomaly) Set(val *CreateTagActionAnomaly) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTagActionAnomaly) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTagActionAnomaly) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTagActionAnomaly(val *CreateTagActionAnomaly) *NullableCreateTagActionAnomaly {
	return &NullableCreateTagActionAnomaly{value: val, isSet: true}
}

func (v NullableCreateTagActionAnomaly) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTagActionAnomaly) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}