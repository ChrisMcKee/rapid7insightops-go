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

// checks if the CreateVariable201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateVariable201Response{}

// CreateVariable201Response struct for CreateVariable201Response
type CreateVariable201Response struct {
	Variable *LeqlVariableResponse `json:"variable,omitempty"`
}

// NewCreateVariable201Response instantiates a new CreateVariable201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateVariable201Response() *CreateVariable201Response {
	this := CreateVariable201Response{}
	return &this
}

// NewCreateVariable201ResponseWithDefaults instantiates a new CreateVariable201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateVariable201ResponseWithDefaults() *CreateVariable201Response {
	this := CreateVariable201Response{}
	return &this
}

// GetVariable returns the Variable field value if set, zero value otherwise.
func (o *CreateVariable201Response) GetVariable() LeqlVariableResponse {
	if o == nil || IsNil(o.Variable) {
		var ret LeqlVariableResponse
		return ret
	}
	return *o.Variable
}

// GetVariableOk returns a tuple with the Variable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVariable201Response) GetVariableOk() (*LeqlVariableResponse, bool) {
	if o == nil || IsNil(o.Variable) {
		return nil, false
	}
	return o.Variable, true
}

// HasVariable returns a boolean if a field has been set.
func (o *CreateVariable201Response) HasVariable() bool {
	if o != nil && !IsNil(o.Variable) {
		return true
	}

	return false
}

// SetVariable gets a reference to the given LeqlVariableResponse and assigns it to the Variable field.
func (o *CreateVariable201Response) SetVariable(v LeqlVariableResponse) {
	o.Variable = &v
}

func (o CreateVariable201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateVariable201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Variable) {
		toSerialize["variable"] = o.Variable
	}
	return toSerialize, nil
}

type NullableCreateVariable201Response struct {
	value *CreateVariable201Response
	isSet bool
}

func (v NullableCreateVariable201Response) Get() *CreateVariable201Response {
	return v.value
}

func (v *NullableCreateVariable201Response) Set(val *CreateVariable201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateVariable201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateVariable201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateVariable201Response(val *CreateVariable201Response) *NullableCreateVariable201Response {
	return &NullableCreateVariable201Response{value: val, isSet: true}
}

func (v NullableCreateVariable201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateVariable201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
