/*
InsightOps REST API

### Overview  Our REST API lets you use InsightOps through HTTP requests. Currently, the REST API allows you to perform the majority of the actions available through the UI, and has some additional functionality that is not available through the UI. You may use this API to automate common tasks (for example, via shell scripts), and to generally interact with InsightOps programmatically.  This page precisely describes the REST API and serves as a reference for the API. Each HTTP method and each URL endpoint is documented in a self-contained unit so that users only need to read about the HTTP methods relevant to them.  ### Terminology  * A **log entry** is an individual log event. * A **log** is a collection of log entries, or a single log stream. * A **log set** is a logical-only collection of logs, i.e. logs can be in multiple logsets and deleting a logset only deletes the relation between logs, not the logs themselves. * [Log Entry Query Language](https://docs.rapid7.com/insightops/log-search) (**LEQL**) is the query language used in Insight Ops to search log data.

API version: latest
Contact: support@rapid7.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package insightops

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the SavedQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SavedQuery{}

// SavedQuery struct for SavedQuery
type SavedQuery struct {
	// The unique id of the saved query.
	Id string `json:"id"`
	// The name of the saved query.
	Name string `json:"name"`
	Leql *Leql  `json:"leql,omitempty"`
	// The log keys of the logs which the query is run against.
	Logs []string `json:"logs,omitempty"`
}

type _SavedQuery SavedQuery

// NewSavedQuery instantiates a new SavedQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSavedQuery(id string, name string) *SavedQuery {
	this := SavedQuery{}
	this.Id = id
	this.Name = name
	return &this
}

// NewSavedQueryWithDefaults instantiates a new SavedQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSavedQueryWithDefaults() *SavedQuery {
	this := SavedQuery{}
	return &this
}

// GetId returns the Id field value
func (o *SavedQuery) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SavedQuery) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SavedQuery) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *SavedQuery) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SavedQuery) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SavedQuery) SetName(v string) {
	o.Name = v
}

// GetLeql returns the Leql field value if set, zero value otherwise.
func (o *SavedQuery) GetLeql() Leql {
	if o == nil || IsNil(o.Leql) {
		var ret Leql
		return ret
	}
	return *o.Leql
}

// GetLeqlOk returns a tuple with the Leql field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedQuery) GetLeqlOk() (*Leql, bool) {
	if o == nil || IsNil(o.Leql) {
		return nil, false
	}
	return o.Leql, true
}

// HasLeql returns a boolean if a field has been set.
func (o *SavedQuery) HasLeql() bool {
	if o != nil && !IsNil(o.Leql) {
		return true
	}

	return false
}

// SetLeql gets a reference to the given Leql and assigns it to the Leql field.
func (o *SavedQuery) SetLeql(v Leql) {
	o.Leql = &v
}

// GetLogs returns the Logs field value if set, zero value otherwise.
func (o *SavedQuery) GetLogs() []string {
	if o == nil || IsNil(o.Logs) {
		var ret []string
		return ret
	}
	return o.Logs
}

// GetLogsOk returns a tuple with the Logs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedQuery) GetLogsOk() ([]string, bool) {
	if o == nil || IsNil(o.Logs) {
		return nil, false
	}
	return o.Logs, true
}

// HasLogs returns a boolean if a field has been set.
func (o *SavedQuery) HasLogs() bool {
	if o != nil && !IsNil(o.Logs) {
		return true
	}

	return false
}

// SetLogs gets a reference to the given []string and assigns it to the Logs field.
func (o *SavedQuery) SetLogs(v []string) {
	o.Logs = v
}

func (o SavedQuery) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SavedQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.Leql) {
		toSerialize["leql"] = o.Leql
	}
	if !IsNil(o.Logs) {
		toSerialize["logs"] = o.Logs
	}
	return toSerialize, nil
}

func (o *SavedQuery) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSavedQuery := _SavedQuery{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSavedQuery)

	if err != nil {
		return err
	}

	*o = SavedQuery(varSavedQuery)

	return err
}

type NullableSavedQuery struct {
	value *SavedQuery
	isSet bool
}

func (v NullableSavedQuery) Get() *SavedQuery {
	return v.value
}

func (v *NullableSavedQuery) Set(val *SavedQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableSavedQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableSavedQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSavedQuery(val *SavedQuery) *NullableSavedQuery {
	return &NullableSavedQuery{value: val, isSet: true}
}

func (v NullableSavedQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSavedQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
