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

// checks if the SourcesIdArrayInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourcesIdArrayInner{}

// SourcesIdArrayInner struct for SourcesIdArrayInner
type SourcesIdArrayInner struct {
	// The UUID of the log.
	Id *string `json:"id,omitempty"`
}

// NewSourcesIdArrayInner instantiates a new SourcesIdArrayInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourcesIdArrayInner() *SourcesIdArrayInner {
	this := SourcesIdArrayInner{}
	return &this
}

// NewSourcesIdArrayInnerWithDefaults instantiates a new SourcesIdArrayInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourcesIdArrayInnerWithDefaults() *SourcesIdArrayInner {
	this := SourcesIdArrayInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SourcesIdArrayInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesIdArrayInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SourcesIdArrayInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SourcesIdArrayInner) SetId(v string) {
	o.Id = &v
}

func (o SourcesIdArrayInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourcesIdArrayInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableSourcesIdArrayInner struct {
	value *SourcesIdArrayInner
	isSet bool
}

func (v NullableSourcesIdArrayInner) Get() *SourcesIdArrayInner {
	return v.value
}

func (v *NullableSourcesIdArrayInner) Set(val *SourcesIdArrayInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSourcesIdArrayInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSourcesIdArrayInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourcesIdArrayInner(val *SourcesIdArrayInner) *NullableSourcesIdArrayInner {
	return &NullableSourcesIdArrayInner{value: val, isSet: true}
}

func (v NullableSourcesIdArrayInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourcesIdArrayInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
