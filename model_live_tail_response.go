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

// checks if the LiveTailResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LiveTailResponse{}

// LiveTailResponse struct for LiveTailResponse
type LiveTailResponse struct {
	// The log keys of the logs for the live tail feed.
	Logs []string     `json:"logs"`
	Leql LeqlLiveTail `json:"leql"`
	// Always empty.
	Events []EventLiveTail `json:"events"`
	Links  []LinksInner    `json:"links"`
}

type _LiveTailResponse LiveTailResponse

// NewLiveTailResponse instantiates a new LiveTailResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiveTailResponse(logs []string, leql LeqlLiveTail, events []EventLiveTail, links []LinksInner) *LiveTailResponse {
	this := LiveTailResponse{}
	this.Logs = logs
	this.Leql = leql
	this.Events = events
	this.Links = links
	return &this
}

// NewLiveTailResponseWithDefaults instantiates a new LiveTailResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiveTailResponseWithDefaults() *LiveTailResponse {
	this := LiveTailResponse{}
	return &this
}

// GetLogs returns the Logs field value
func (o *LiveTailResponse) GetLogs() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Logs
}

// GetLogsOk returns a tuple with the Logs field value
// and a boolean to check if the value has been set.
func (o *LiveTailResponse) GetLogsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Logs, true
}

// SetLogs sets field value
func (o *LiveTailResponse) SetLogs(v []string) {
	o.Logs = v
}

// GetLeql returns the Leql field value
func (o *LiveTailResponse) GetLeql() LeqlLiveTail {
	if o == nil {
		var ret LeqlLiveTail
		return ret
	}

	return o.Leql
}

// GetLeqlOk returns a tuple with the Leql field value
// and a boolean to check if the value has been set.
func (o *LiveTailResponse) GetLeqlOk() (*LeqlLiveTail, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Leql, true
}

// SetLeql sets field value
func (o *LiveTailResponse) SetLeql(v LeqlLiveTail) {
	o.Leql = v
}

// GetEvents returns the Events field value
func (o *LiveTailResponse) GetEvents() []EventLiveTail {
	if o == nil {
		var ret []EventLiveTail
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *LiveTailResponse) GetEventsOk() ([]EventLiveTail, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *LiveTailResponse) SetEvents(v []EventLiveTail) {
	o.Events = v
}

// GetLinks returns the Links field value
func (o *LiveTailResponse) GetLinks() []LinksInner {
	if o == nil {
		var ret []LinksInner
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *LiveTailResponse) GetLinksOk() ([]LinksInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *LiveTailResponse) SetLinks(v []LinksInner) {
	o.Links = v
}

func (o LiveTailResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LiveTailResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["logs"] = o.Logs
	toSerialize["leql"] = o.Leql
	toSerialize["events"] = o.Events
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *LiveTailResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"logs",
		"leql",
		"events",
		"links",
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

	varLiveTailResponse := _LiveTailResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLiveTailResponse)

	if err != nil {
		return err
	}

	*o = LiveTailResponse(varLiveTailResponse)

	return err
}

type NullableLiveTailResponse struct {
	value *LiveTailResponse
	isSet bool
}

func (v NullableLiveTailResponse) Get() *LiveTailResponse {
	return v.value
}

func (v *NullableLiveTailResponse) Set(val *LiveTailResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLiveTailResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLiveTailResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiveTailResponse(val *LiveTailResponse) *NullableLiveTailResponse {
	return &NullableLiveTailResponse{value: val, isSet: true}
}

func (v NullableLiveTailResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLiveTailResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
