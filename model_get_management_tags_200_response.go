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

// checks if the GetManagementTags200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetManagementTags200Response{}

// GetManagementTags200Response struct for GetManagementTags200Response
type GetManagementTags200Response struct {
	// All the basic detection rules for this account.
	Tags []GetManagementTags200ResponseTagsInner `json:"tags,omitempty"`
}

// NewGetManagementTags200Response instantiates a new GetManagementTags200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetManagementTags200Response() *GetManagementTags200Response {
	this := GetManagementTags200Response{}
	return &this
}

// NewGetManagementTags200ResponseWithDefaults instantiates a new GetManagementTags200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetManagementTags200ResponseWithDefaults() *GetManagementTags200Response {
	this := GetManagementTags200Response{}
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GetManagementTags200Response) GetTags() []GetManagementTags200ResponseTagsInner {
	if o == nil || IsNil(o.Tags) {
		var ret []GetManagementTags200ResponseTagsInner
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetManagementTags200Response) GetTagsOk() ([]GetManagementTags200ResponseTagsInner, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GetManagementTags200Response) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []GetManagementTags200ResponseTagsInner and assigns it to the Tags field.
func (o *GetManagementTags200Response) SetTags(v []GetManagementTags200ResponseTagsInner) {
	o.Tags = v
}

func (o GetManagementTags200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetManagementTags200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableGetManagementTags200Response struct {
	value *GetManagementTags200Response
	isSet bool
}

func (v NullableGetManagementTags200Response) Get() *GetManagementTags200Response {
	return v.value
}

func (v *NullableGetManagementTags200Response) Set(val *GetManagementTags200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetManagementTags200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetManagementTags200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetManagementTags200Response(val *GetManagementTags200Response) *NullableGetManagementTags200Response {
	return &NullableGetManagementTags200Response{value: val, isSet: true}
}

func (v NullableGetManagementTags200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetManagementTags200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
