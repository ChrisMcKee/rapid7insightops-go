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

// checks if the ContextResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContextResponse{}

// ContextResponse struct for ContextResponse
type ContextResponse struct {
	// The log key of the log which contains the contextual log entries.
	Logs []string    `json:"logs"`
	Leql LeqlContext `json:"leql"`
	// The contextual log entries for the given log entry. It is empty for a 202 response.
	Events []EventQuery           `json:"events"`
	Links  []ContextApiLinksInner `json:"links,omitempty"`
}

type _ContextResponse ContextResponse

// NewContextResponse instantiates a new ContextResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContextResponse(logs []string, leql LeqlContext, events []EventQuery) *ContextResponse {
	this := ContextResponse{}
	this.Logs = logs
	this.Leql = leql
	this.Events = events
	return &this
}

// NewContextResponseWithDefaults instantiates a new ContextResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContextResponseWithDefaults() *ContextResponse {
	this := ContextResponse{}
	return &this
}

// GetLogs returns the Logs field value
func (o *ContextResponse) GetLogs() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Logs
}

// GetLogsOk returns a tuple with the Logs field value
// and a boolean to check if the value has been set.
func (o *ContextResponse) GetLogsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Logs, true
}

// SetLogs sets field value
func (o *ContextResponse) SetLogs(v []string) {
	o.Logs = v
}

// GetLeql returns the Leql field value
func (o *ContextResponse) GetLeql() LeqlContext {
	if o == nil {
		var ret LeqlContext
		return ret
	}

	return o.Leql
}

// GetLeqlOk returns a tuple with the Leql field value
// and a boolean to check if the value has been set.
func (o *ContextResponse) GetLeqlOk() (*LeqlContext, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Leql, true
}

// SetLeql sets field value
func (o *ContextResponse) SetLeql(v LeqlContext) {
	o.Leql = v
}

// GetEvents returns the Events field value
func (o *ContextResponse) GetEvents() []EventQuery {
	if o == nil {
		var ret []EventQuery
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *ContextResponse) GetEventsOk() ([]EventQuery, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *ContextResponse) SetEvents(v []EventQuery) {
	o.Events = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ContextResponse) GetLinks() []ContextApiLinksInner {
	if o == nil || IsNil(o.Links) {
		var ret []ContextApiLinksInner
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextResponse) GetLinksOk() ([]ContextApiLinksInner, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ContextResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ContextApiLinksInner and assigns it to the Links field.
func (o *ContextResponse) SetLinks(v []ContextApiLinksInner) {
	o.Links = v
}

func (o ContextResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContextResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["logs"] = o.Logs
	toSerialize["leql"] = o.Leql
	toSerialize["events"] = o.Events
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

func (o *ContextResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"logs",
		"leql",
		"events",
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

	varContextResponse := _ContextResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContextResponse)

	if err != nil {
		return err
	}

	*o = ContextResponse(varContextResponse)

	return err
}

type NullableContextResponse struct {
	value *ContextResponse
	isSet bool
}

func (v NullableContextResponse) Get() *ContextResponse {
	return v.value
}

func (v *NullableContextResponse) Set(val *ContextResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContextResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContextResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContextResponse(val *ContextResponse) *NullableContextResponse {
	return &NullableContextResponse{value: val, isSet: true}
}

func (v NullableContextResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContextResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
