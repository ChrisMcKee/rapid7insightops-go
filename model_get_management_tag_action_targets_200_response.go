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

// checks if the GetManagementTagActionTargets200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetManagementTagActionTargets200Response{}

// GetManagementTagActionTargets200Response struct for GetManagementTagActionTargets200Response
type GetManagementTagActionTargets200Response struct {
	// All the notification targets for this account.
	Targets []GetTargetAction `json:"targets,omitempty"`
}

// NewGetManagementTagActionTargets200Response instantiates a new GetManagementTagActionTargets200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetManagementTagActionTargets200Response() *GetManagementTagActionTargets200Response {
	this := GetManagementTagActionTargets200Response{}
	return &this
}

// NewGetManagementTagActionTargets200ResponseWithDefaults instantiates a new GetManagementTagActionTargets200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetManagementTagActionTargets200ResponseWithDefaults() *GetManagementTagActionTargets200Response {
	this := GetManagementTagActionTargets200Response{}
	return &this
}

// GetTargets returns the Targets field value if set, zero value otherwise.
func (o *GetManagementTagActionTargets200Response) GetTargets() []GetTargetAction {
	if o == nil || IsNil(o.Targets) {
		var ret []GetTargetAction
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetManagementTagActionTargets200Response) GetTargetsOk() ([]GetTargetAction, bool) {
	if o == nil || IsNil(o.Targets) {
		return nil, false
	}
	return o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *GetManagementTagActionTargets200Response) HasTargets() bool {
	if o != nil && !IsNil(o.Targets) {
		return true
	}

	return false
}

// SetTargets gets a reference to the given []GetTargetAction and assigns it to the Targets field.
func (o *GetManagementTagActionTargets200Response) SetTargets(v []GetTargetAction) {
	o.Targets = v
}

func (o GetManagementTagActionTargets200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetManagementTagActionTargets200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Targets) {
		toSerialize["targets"] = o.Targets
	}
	return toSerialize, nil
}

type NullableGetManagementTagActionTargets200Response struct {
	value *GetManagementTagActionTargets200Response
	isSet bool
}

func (v NullableGetManagementTagActionTargets200Response) Get() *GetManagementTagActionTargets200Response {
	return v.value
}

func (v *NullableGetManagementTagActionTargets200Response) Set(val *GetManagementTagActionTargets200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetManagementTagActionTargets200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetManagementTagActionTargets200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetManagementTagActionTargets200Response(val *GetManagementTagActionTargets200Response) *NullableGetManagementTagActionTargets200Response {
	return &NullableGetManagementTagActionTargets200Response{value: val, isSet: true}
}

func (v NullableGetManagementTagActionTargets200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetManagementTagActionTargets200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
