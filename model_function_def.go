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

// checks if the FunctionDef type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FunctionDef{}

// FunctionDef struct for FunctionDef
type FunctionDef struct {
	Calculate *string `json:"calculate,omitempty"`
}

// NewFunctionDef instantiates a new FunctionDef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFunctionDef() *FunctionDef {
	this := FunctionDef{}
	return &this
}

// NewFunctionDefWithDefaults instantiates a new FunctionDef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFunctionDefWithDefaults() *FunctionDef {
	this := FunctionDef{}
	return &this
}

// GetCalculate returns the Calculate field value if set, zero value otherwise.
func (o *FunctionDef) GetCalculate() string {
	if o == nil || IsNil(o.Calculate) {
		var ret string
		return ret
	}
	return *o.Calculate
}

// GetCalculateOk returns a tuple with the Calculate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionDef) GetCalculateOk() (*string, bool) {
	if o == nil || IsNil(o.Calculate) {
		return nil, false
	}
	return o.Calculate, true
}

// HasCalculate returns a boolean if a field has been set.
func (o *FunctionDef) HasCalculate() bool {
	if o != nil && !IsNil(o.Calculate) {
		return true
	}

	return false
}

// SetCalculate gets a reference to the given string and assigns it to the Calculate field.
func (o *FunctionDef) SetCalculate(v string) {
	o.Calculate = &v
}

func (o FunctionDef) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FunctionDef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Calculate) {
		toSerialize["calculate"] = o.Calculate
	}
	return toSerialize, nil
}

type NullableFunctionDef struct {
	value *FunctionDef
	isSet bool
}

func (v NullableFunctionDef) Get() *FunctionDef {
	return v.value
}

func (v *NullableFunctionDef) Set(val *FunctionDef) {
	v.value = val
	v.isSet = true
}

func (v NullableFunctionDef) IsSet() bool {
	return v.isSet
}

func (v *NullableFunctionDef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFunctionDef(val *FunctionDef) *NullableFunctionDef {
	return &NullableFunctionDef{value: val, isSet: true}
}

func (v NullableFunctionDef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFunctionDef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
