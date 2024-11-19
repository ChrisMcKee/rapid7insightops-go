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

// checks if the LabelsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LabelsResponse{}

// LabelsResponse struct for LabelsResponse
type LabelsResponse struct {
	// The labels attached to the detection rule.
	Labels []LabelResponse `json:"labels,omitempty"`
}

// NewLabelsResponse instantiates a new LabelsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabelsResponse() *LabelsResponse {
	this := LabelsResponse{}
	return &this
}

// NewLabelsResponseWithDefaults instantiates a new LabelsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabelsResponseWithDefaults() *LabelsResponse {
	this := LabelsResponse{}
	return &this
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *LabelsResponse) GetLabels() []LabelResponse {
	if o == nil || IsNil(o.Labels) {
		var ret []LabelResponse
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelsResponse) GetLabelsOk() ([]LabelResponse, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *LabelsResponse) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []LabelResponse and assigns it to the Labels field.
func (o *LabelsResponse) SetLabels(v []LabelResponse) {
	o.Labels = v
}

func (o LabelsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LabelsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	return toSerialize, nil
}

type NullableLabelsResponse struct {
	value *LabelsResponse
	isSet bool
}

func (v NullableLabelsResponse) Get() *LabelsResponse {
	return v.value
}

func (v *NullableLabelsResponse) Set(val *LabelsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLabelsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLabelsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabelsResponse(val *LabelsResponse) *NullableLabelsResponse {
	return &NullableLabelsResponse{value: val, isSet: true}
}

func (v NullableLabelsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabelsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
