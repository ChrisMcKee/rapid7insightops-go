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

// checks if the AuditLogList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditLogList{}

// AuditLogList The list of log sets.
type AuditLogList struct {
	Logs []AuditLog `json:"logs,omitempty"`
}

// NewAuditLogList instantiates a new AuditLogList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditLogList() *AuditLogList {
	this := AuditLogList{}
	return &this
}

// NewAuditLogListWithDefaults instantiates a new AuditLogList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditLogListWithDefaults() *AuditLogList {
	this := AuditLogList{}
	return &this
}

// GetLogs returns the Logs field value if set, zero value otherwise.
func (o *AuditLogList) GetLogs() []AuditLog {
	if o == nil || IsNil(o.Logs) {
		var ret []AuditLog
		return ret
	}
	return o.Logs
}

// GetLogsOk returns a tuple with the Logs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogList) GetLogsOk() ([]AuditLog, bool) {
	if o == nil || IsNil(o.Logs) {
		return nil, false
	}
	return o.Logs, true
}

// HasLogs returns a boolean if a field has been set.
func (o *AuditLogList) HasLogs() bool {
	if o != nil && !IsNil(o.Logs) {
		return true
	}

	return false
}

// SetLogs gets a reference to the given []AuditLog and assigns it to the Logs field.
func (o *AuditLogList) SetLogs(v []AuditLog) {
	o.Logs = v
}

func (o AuditLogList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuditLogList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Logs) {
		toSerialize["logs"] = o.Logs
	}
	return toSerialize, nil
}

type NullableAuditLogList struct {
	value *AuditLogList
	isSet bool
}

func (v NullableAuditLogList) Get() *AuditLogList {
	return v.value
}

func (v *NullableAuditLogList) Set(val *AuditLogList) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditLogList) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditLogList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditLogList(val *AuditLogList) *NullableAuditLogList {
	return &NullableAuditLogList{value: val, isSet: true}
}

func (v NullableAuditLogList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditLogList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
