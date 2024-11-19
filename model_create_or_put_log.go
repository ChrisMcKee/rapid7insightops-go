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

// checks if the CreateOrPutLog type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrPutLog{}

// CreateOrPutLog struct for CreateOrPutLog
type CreateOrPutLog struct {
	Log CreateOrPutLogLog `json:"log"`
}

type _CreateOrPutLog CreateOrPutLog

// NewCreateOrPutLog instantiates a new CreateOrPutLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrPutLog(log CreateOrPutLogLog) *CreateOrPutLog {
	this := CreateOrPutLog{}
	this.Log = log
	return &this
}

// NewCreateOrPutLogWithDefaults instantiates a new CreateOrPutLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrPutLogWithDefaults() *CreateOrPutLog {
	this := CreateOrPutLog{}
	return &this
}

// GetLog returns the Log field value
func (o *CreateOrPutLog) GetLog() CreateOrPutLogLog {
	if o == nil {
		var ret CreateOrPutLogLog
		return ret
	}

	return o.Log
}

// GetLogOk returns a tuple with the Log field value
// and a boolean to check if the value has been set.
func (o *CreateOrPutLog) GetLogOk() (*CreateOrPutLogLog, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Log, true
}

// SetLog sets field value
func (o *CreateOrPutLog) SetLog(v CreateOrPutLogLog) {
	o.Log = v
}

func (o CreateOrPutLog) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrPutLog) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["log"] = o.Log
	return toSerialize, nil
}

func (o *CreateOrPutLog) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"log",
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

	varCreateOrPutLog := _CreateOrPutLog{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateOrPutLog)

	if err != nil {
		return err
	}

	*o = CreateOrPutLog(varCreateOrPutLog)

	return err
}

type NullableCreateOrPutLog struct {
	value *CreateOrPutLog
	isSet bool
}

func (v NullableCreateOrPutLog) Get() *CreateOrPutLog {
	return v.value
}

func (v *NullableCreateOrPutLog) Set(val *CreateOrPutLog) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrPutLog) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrPutLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrPutLog(val *CreateOrPutLog) *NullableCreateOrPutLog {
	return &NullableCreateOrPutLog{value: val, isSet: true}
}

func (v NullableCreateOrPutLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrPutLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}