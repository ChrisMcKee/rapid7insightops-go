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

// checks if the PatchLabel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchLabel{}

// PatchLabel struct for PatchLabel
type PatchLabel struct {
	Label *PatchLabelLabel `json:"label,omitempty"`
}

// NewPatchLabel instantiates a new PatchLabel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchLabel() *PatchLabel {
	this := PatchLabel{}
	return &this
}

// NewPatchLabelWithDefaults instantiates a new PatchLabel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchLabelWithDefaults() *PatchLabel {
	this := PatchLabel{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PatchLabel) GetLabel() PatchLabelLabel {
	if o == nil || IsNil(o.Label) {
		var ret PatchLabelLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchLabel) GetLabelOk() (*PatchLabelLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PatchLabel) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given PatchLabelLabel and assigns it to the Label field.
func (o *PatchLabel) SetLabel(v PatchLabelLabel) {
	o.Label = &v
}

func (o PatchLabel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchLabel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	return toSerialize, nil
}

type NullablePatchLabel struct {
	value *PatchLabel
	isSet bool
}

func (v NullablePatchLabel) Get() *PatchLabel {
	return v.value
}

func (v *NullablePatchLabel) Set(val *PatchLabel) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchLabel) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchLabel(val *PatchLabel) *NullablePatchLabel {
	return &NullablePatchLabel{value: val, isSet: true}
}

func (v NullablePatchLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
