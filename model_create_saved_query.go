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

// checks if the CreateSavedQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSavedQuery{}

// CreateSavedQuery struct for CreateSavedQuery
type CreateSavedQuery struct {
	SavedQuery *CreateSavedQuerySavedQuery `json:"saved_query,omitempty"`
}

// NewCreateSavedQuery instantiates a new CreateSavedQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSavedQuery() *CreateSavedQuery {
	this := CreateSavedQuery{}
	return &this
}

// NewCreateSavedQueryWithDefaults instantiates a new CreateSavedQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSavedQueryWithDefaults() *CreateSavedQuery {
	this := CreateSavedQuery{}
	return &this
}

// GetSavedQuery returns the SavedQuery field value if set, zero value otherwise.
func (o *CreateSavedQuery) GetSavedQuery() CreateSavedQuerySavedQuery {
	if o == nil || IsNil(o.SavedQuery) {
		var ret CreateSavedQuerySavedQuery
		return ret
	}
	return *o.SavedQuery
}

// GetSavedQueryOk returns a tuple with the SavedQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSavedQuery) GetSavedQueryOk() (*CreateSavedQuerySavedQuery, bool) {
	if o == nil || IsNil(o.SavedQuery) {
		return nil, false
	}
	return o.SavedQuery, true
}

// HasSavedQuery returns a boolean if a field has been set.
func (o *CreateSavedQuery) HasSavedQuery() bool {
	if o != nil && !IsNil(o.SavedQuery) {
		return true
	}

	return false
}

// SetSavedQuery gets a reference to the given CreateSavedQuerySavedQuery and assigns it to the SavedQuery field.
func (o *CreateSavedQuery) SetSavedQuery(v CreateSavedQuerySavedQuery) {
	o.SavedQuery = &v
}

func (o CreateSavedQuery) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSavedQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SavedQuery) {
		toSerialize["saved_query"] = o.SavedQuery
	}
	return toSerialize, nil
}

type NullableCreateSavedQuery struct {
	value *CreateSavedQuery
	isSet bool
}

func (v NullableCreateSavedQuery) Get() *CreateSavedQuery {
	return v.value
}

func (v *NullableCreateSavedQuery) Set(val *CreateSavedQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSavedQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSavedQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSavedQuery(val *CreateSavedQuery) *NullableCreateSavedQuery {
	return &NullableCreateSavedQuery{value: val, isSet: true}
}

func (v NullableCreateSavedQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSavedQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}