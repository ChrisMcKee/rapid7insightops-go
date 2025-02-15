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

// checks if the ListSavedQueries200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListSavedQueries200Response{}

// ListSavedQueries200Response struct for ListSavedQueries200Response
type ListSavedQueries200Response struct {
	// All the saved queries for this account.
	SavedQueries []SavedQuery `json:"saved_queries,omitempty"`
}

// NewListSavedQueries200Response instantiates a new ListSavedQueries200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSavedQueries200Response() *ListSavedQueries200Response {
	this := ListSavedQueries200Response{}
	return &this
}

// NewListSavedQueries200ResponseWithDefaults instantiates a new ListSavedQueries200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSavedQueries200ResponseWithDefaults() *ListSavedQueries200Response {
	this := ListSavedQueries200Response{}
	return &this
}

// GetSavedQueries returns the SavedQueries field value if set, zero value otherwise.
func (o *ListSavedQueries200Response) GetSavedQueries() []SavedQuery {
	if o == nil || IsNil(o.SavedQueries) {
		var ret []SavedQuery
		return ret
	}
	return o.SavedQueries
}

// GetSavedQueriesOk returns a tuple with the SavedQueries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSavedQueries200Response) GetSavedQueriesOk() ([]SavedQuery, bool) {
	if o == nil || IsNil(o.SavedQueries) {
		return nil, false
	}
	return o.SavedQueries, true
}

// HasSavedQueries returns a boolean if a field has been set.
func (o *ListSavedQueries200Response) HasSavedQueries() bool {
	if o != nil && !IsNil(o.SavedQueries) {
		return true
	}

	return false
}

// SetSavedQueries gets a reference to the given []SavedQuery and assigns it to the SavedQueries field.
func (o *ListSavedQueries200Response) SetSavedQueries(v []SavedQuery) {
	o.SavedQueries = v
}

func (o ListSavedQueries200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListSavedQueries200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SavedQueries) {
		toSerialize["saved_queries"] = o.SavedQueries
	}
	return toSerialize, nil
}

type NullableListSavedQueries200Response struct {
	value *ListSavedQueries200Response
	isSet bool
}

func (v NullableListSavedQueries200Response) Get() *ListSavedQueries200Response {
	return v.value
}

func (v *NullableListSavedQueries200Response) Set(val *ListSavedQueries200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListSavedQueries200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListSavedQueries200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSavedQueries200Response(val *ListSavedQueries200Response) *NullableListSavedQueries200Response {
	return &NullableListSavedQueries200Response{value: val, isSet: true}
}

func (v NullableListSavedQueries200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSavedQueries200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
