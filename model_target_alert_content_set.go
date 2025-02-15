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

// checks if the TargetAlertContentSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TargetAlertContentSet{}

// TargetAlertContentSet The list of key-value pairs to modify the content of the notification.
type TargetAlertContentSet struct {
	// When set to \"true\", the log context is set to appear in the notification. When set to \"false\", the log context is set to not appear in the notification.
	LeContext *string `json:"le_context,omitempty"`
	// When set to \"true\", the log entry which triggered the alert will appear in the notification. When set to \"false\", the log entry which triggered the alert will not appear in the notification.
	LeTriggerEvent []string `json:"le_trigger_event,omitempty"`
}

// NewTargetAlertContentSet instantiates a new TargetAlertContentSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetAlertContentSet() *TargetAlertContentSet {
	this := TargetAlertContentSet{}
	return &this
}

// NewTargetAlertContentSetWithDefaults instantiates a new TargetAlertContentSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetAlertContentSetWithDefaults() *TargetAlertContentSet {
	this := TargetAlertContentSet{}
	return &this
}

// GetLeContext returns the LeContext field value if set, zero value otherwise.
func (o *TargetAlertContentSet) GetLeContext() string {
	if o == nil || IsNil(o.LeContext) {
		var ret string
		return ret
	}
	return *o.LeContext
}

// GetLeContextOk returns a tuple with the LeContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetAlertContentSet) GetLeContextOk() (*string, bool) {
	if o == nil || IsNil(o.LeContext) {
		return nil, false
	}
	return o.LeContext, true
}

// HasLeContext returns a boolean if a field has been set.
func (o *TargetAlertContentSet) HasLeContext() bool {
	if o != nil && !IsNil(o.LeContext) {
		return true
	}

	return false
}

// SetLeContext gets a reference to the given string and assigns it to the LeContext field.
func (o *TargetAlertContentSet) SetLeContext(v string) {
	o.LeContext = &v
}

// GetLeTriggerEvent returns the LeTriggerEvent field value if set, zero value otherwise.
func (o *TargetAlertContentSet) GetLeTriggerEvent() []string {
	if o == nil || IsNil(o.LeTriggerEvent) {
		var ret []string
		return ret
	}
	return o.LeTriggerEvent
}

// GetLeTriggerEventOk returns a tuple with the LeTriggerEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetAlertContentSet) GetLeTriggerEventOk() ([]string, bool) {
	if o == nil || IsNil(o.LeTriggerEvent) {
		return nil, false
	}
	return o.LeTriggerEvent, true
}

// HasLeTriggerEvent returns a boolean if a field has been set.
func (o *TargetAlertContentSet) HasLeTriggerEvent() bool {
	if o != nil && !IsNil(o.LeTriggerEvent) {
		return true
	}

	return false
}

// SetLeTriggerEvent gets a reference to the given []string and assigns it to the LeTriggerEvent field.
func (o *TargetAlertContentSet) SetLeTriggerEvent(v []string) {
	o.LeTriggerEvent = v
}

func (o TargetAlertContentSet) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TargetAlertContentSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LeContext) {
		toSerialize["le_context"] = o.LeContext
	}
	if !IsNil(o.LeTriggerEvent) {
		toSerialize["le_trigger_event"] = o.LeTriggerEvent
	}
	return toSerialize, nil
}

type NullableTargetAlertContentSet struct {
	value *TargetAlertContentSet
	isSet bool
}

func (v NullableTargetAlertContentSet) Get() *TargetAlertContentSet {
	return v.value
}

func (v *NullableTargetAlertContentSet) Set(val *TargetAlertContentSet) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetAlertContentSet) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetAlertContentSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetAlertContentSet(val *TargetAlertContentSet) *NullableTargetAlertContentSet {
	return &NullableTargetAlertContentSet{value: val, isSet: true}
}

func (v NullableTargetAlertContentSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTargetAlertContentSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
