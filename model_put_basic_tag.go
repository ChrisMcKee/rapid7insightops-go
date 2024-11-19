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

// checks if the PutBasicTag type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutBasicTag{}

// PutBasicTag struct for PutBasicTag
type PutBasicTag struct {
	Tag *PutBasicTagTag `json:"tag,omitempty"`
}

// NewPutBasicTag instantiates a new PutBasicTag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutBasicTag() *PutBasicTag {
	this := PutBasicTag{}
	return &this
}

// NewPutBasicTagWithDefaults instantiates a new PutBasicTag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutBasicTagWithDefaults() *PutBasicTag {
	this := PutBasicTag{}
	return &this
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *PutBasicTag) GetTag() PutBasicTagTag {
	if o == nil || IsNil(o.Tag) {
		var ret PutBasicTagTag
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutBasicTag) GetTagOk() (*PutBasicTagTag, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *PutBasicTag) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given PutBasicTagTag and assigns it to the Tag field.
func (o *PutBasicTag) SetTag(v PutBasicTagTag) {
	o.Tag = &v
}

func (o PutBasicTag) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutBasicTag) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	return toSerialize, nil
}

type NullablePutBasicTag struct {
	value *PutBasicTag
	isSet bool
}

func (v NullablePutBasicTag) Get() *PutBasicTag {
	return v.value
}

func (v *NullablePutBasicTag) Set(val *PutBasicTag) {
	v.value = val
	v.isSet = true
}

func (v NullablePutBasicTag) IsSet() bool {
	return v.isSet
}

func (v *NullablePutBasicTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutBasicTag(val *PutBasicTag) *NullablePutBasicTag {
	return &NullablePutBasicTag{value: val, isSet: true}
}

func (v NullablePutBasicTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutBasicTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
