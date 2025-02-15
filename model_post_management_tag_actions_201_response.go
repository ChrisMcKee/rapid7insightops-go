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

// checks if the PostManagementTagActions201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostManagementTagActions201Response{}

// PostManagementTagActions201Response struct for PostManagementTagActions201Response
type PostManagementTagActions201Response struct {
	Action *AlertNotificationSetting `json:"action,omitempty"`
}

// NewPostManagementTagActions201Response instantiates a new PostManagementTagActions201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostManagementTagActions201Response() *PostManagementTagActions201Response {
	this := PostManagementTagActions201Response{}
	return &this
}

// NewPostManagementTagActions201ResponseWithDefaults instantiates a new PostManagementTagActions201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostManagementTagActions201ResponseWithDefaults() *PostManagementTagActions201Response {
	this := PostManagementTagActions201Response{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *PostManagementTagActions201Response) GetAction() AlertNotificationSetting {
	if o == nil || IsNil(o.Action) {
		var ret AlertNotificationSetting
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostManagementTagActions201Response) GetActionOk() (*AlertNotificationSetting, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *PostManagementTagActions201Response) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given AlertNotificationSetting and assigns it to the Action field.
func (o *PostManagementTagActions201Response) SetAction(v AlertNotificationSetting) {
	o.Action = &v
}

func (o PostManagementTagActions201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostManagementTagActions201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	return toSerialize, nil
}

type NullablePostManagementTagActions201Response struct {
	value *PostManagementTagActions201Response
	isSet bool
}

func (v NullablePostManagementTagActions201Response) Get() *PostManagementTagActions201Response {
	return v.value
}

func (v *NullablePostManagementTagActions201Response) Set(val *PostManagementTagActions201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePostManagementTagActions201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePostManagementTagActions201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostManagementTagActions201Response(val *PostManagementTagActions201Response) *NullablePostManagementTagActions201Response {
	return &NullablePostManagementTagActions201Response{value: val, isSet: true}
}

func (v NullablePostManagementTagActions201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostManagementTagActions201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
