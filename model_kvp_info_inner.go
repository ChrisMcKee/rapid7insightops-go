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

// checks if the KvpInfoInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KvpInfoInner{}

// KvpInfoInner struct for KvpInfoInner
type KvpInfoInner struct {
	Key   KvpInfoInnerKey   `json:"key"`
	Value KvpInfoInnerValue `json:"value"`
}

type _KvpInfoInner KvpInfoInner

// NewKvpInfoInner instantiates a new KvpInfoInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKvpInfoInner(key KvpInfoInnerKey, value KvpInfoInnerValue) *KvpInfoInner {
	this := KvpInfoInner{}
	this.Key = key
	this.Value = value
	return &this
}

// NewKvpInfoInnerWithDefaults instantiates a new KvpInfoInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKvpInfoInnerWithDefaults() *KvpInfoInner {
	this := KvpInfoInner{}
	return &this
}

// GetKey returns the Key field value
func (o *KvpInfoInner) GetKey() KvpInfoInnerKey {
	if o == nil {
		var ret KvpInfoInnerKey
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *KvpInfoInner) GetKeyOk() (*KvpInfoInnerKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *KvpInfoInner) SetKey(v KvpInfoInnerKey) {
	o.Key = v
}

// GetValue returns the Value field value
func (o *KvpInfoInner) GetValue() KvpInfoInnerValue {
	if o == nil {
		var ret KvpInfoInnerValue
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *KvpInfoInner) GetValueOk() (*KvpInfoInnerValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *KvpInfoInner) SetValue(v KvpInfoInnerValue) {
	o.Value = v
}

func (o KvpInfoInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KvpInfoInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *KvpInfoInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
		"value",
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

	varKvpInfoInner := _KvpInfoInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKvpInfoInner)

	if err != nil {
		return err
	}

	*o = KvpInfoInner(varKvpInfoInner)

	return err
}

type NullableKvpInfoInner struct {
	value *KvpInfoInner
	isSet bool
}

func (v NullableKvpInfoInner) Get() *KvpInfoInner {
	return v.value
}

func (v *NullableKvpInfoInner) Set(val *KvpInfoInner) {
	v.value = val
	v.isSet = true
}

func (v NullableKvpInfoInner) IsSet() bool {
	return v.isSet
}

func (v *NullableKvpInfoInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKvpInfoInner(val *KvpInfoInner) *NullableKvpInfoInner {
	return &NullableKvpInfoInner{value: val, isSet: true}
}

func (v NullableKvpInfoInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKvpInfoInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}