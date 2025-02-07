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

// checks if the CreatePutTargetTarget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreatePutTargetTarget{}

// CreatePutTargetTarget This creates a target.
type CreatePutTargetTarget struct {
	// The name of the target.
	Name string `json:"name"`
	// The description of the target.
	Description     *string                    `json:"description,omitempty"`
	Type            NotificationTypeEnum       `json:"type"`
	ParamsSet       CreatePatchTargetParamsSet `json:"params_set"`
	AlertContentSet TargetAlertContentSet      `json:"alert_content_set"`
	// A list of key-value pairs that may indicate some auxiliary information about the notification target.
	UserData map[string]interface{} `json:"user_data"`
}

type _CreatePutTargetTarget CreatePutTargetTarget

// NewCreatePutTargetTarget instantiates a new CreatePutTargetTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePutTargetTarget(name string, type_ NotificationTypeEnum, paramsSet CreatePatchTargetParamsSet, alertContentSet TargetAlertContentSet, userData map[string]interface{}) *CreatePutTargetTarget {
	this := CreatePutTargetTarget{}
	this.Name = name
	this.Type = type_
	this.ParamsSet = paramsSet
	this.AlertContentSet = alertContentSet
	this.UserData = userData
	return &this
}

// NewCreatePutTargetTargetWithDefaults instantiates a new CreatePutTargetTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePutTargetTargetWithDefaults() *CreatePutTargetTarget {
	this := CreatePutTargetTarget{}
	return &this
}

// GetName returns the Name field value
func (o *CreatePutTargetTarget) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreatePutTargetTarget) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreatePutTargetTarget) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreatePutTargetTarget) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePutTargetTarget) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreatePutTargetTarget) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreatePutTargetTarget) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value
func (o *CreatePutTargetTarget) GetType() NotificationTypeEnum {
	if o == nil {
		var ret NotificationTypeEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreatePutTargetTarget) GetTypeOk() (*NotificationTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreatePutTargetTarget) SetType(v NotificationTypeEnum) {
	o.Type = v
}

// GetParamsSet returns the ParamsSet field value
func (o *CreatePutTargetTarget) GetParamsSet() CreatePatchTargetParamsSet {
	if o == nil {
		var ret CreatePatchTargetParamsSet
		return ret
	}

	return o.ParamsSet
}

// GetParamsSetOk returns a tuple with the ParamsSet field value
// and a boolean to check if the value has been set.
func (o *CreatePutTargetTarget) GetParamsSetOk() (*CreatePatchTargetParamsSet, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParamsSet, true
}

// SetParamsSet sets field value
func (o *CreatePutTargetTarget) SetParamsSet(v CreatePatchTargetParamsSet) {
	o.ParamsSet = v
}

// GetAlertContentSet returns the AlertContentSet field value
func (o *CreatePutTargetTarget) GetAlertContentSet() TargetAlertContentSet {
	if o == nil {
		var ret TargetAlertContentSet
		return ret
	}

	return o.AlertContentSet
}

// GetAlertContentSetOk returns a tuple with the AlertContentSet field value
// and a boolean to check if the value has been set.
func (o *CreatePutTargetTarget) GetAlertContentSetOk() (*TargetAlertContentSet, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlertContentSet, true
}

// SetAlertContentSet sets field value
func (o *CreatePutTargetTarget) SetAlertContentSet(v TargetAlertContentSet) {
	o.AlertContentSet = v
}

// GetUserData returns the UserData field value
func (o *CreatePutTargetTarget) GetUserData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.UserData
}

// GetUserDataOk returns a tuple with the UserData field value
// and a boolean to check if the value has been set.
func (o *CreatePutTargetTarget) GetUserDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.UserData, true
}

// SetUserData sets field value
func (o *CreatePutTargetTarget) SetUserData(v map[string]interface{}) {
	o.UserData = v
}

func (o CreatePutTargetTarget) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatePutTargetTarget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["type"] = o.Type
	toSerialize["params_set"] = o.ParamsSet
	toSerialize["alert_content_set"] = o.AlertContentSet
	toSerialize["user_data"] = o.UserData
	return toSerialize, nil
}

func (o *CreatePutTargetTarget) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"type",
		"params_set",
		"alert_content_set",
		"user_data",
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

	varCreatePutTargetTarget := _CreatePutTargetTarget{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreatePutTargetTarget)

	if err != nil {
		return err
	}

	*o = CreatePutTargetTarget(varCreatePutTargetTarget)

	return err
}

type NullableCreatePutTargetTarget struct {
	value *CreatePutTargetTarget
	isSet bool
}

func (v NullableCreatePutTargetTarget) Get() *CreatePutTargetTarget {
	return v.value
}

func (v *NullableCreatePutTargetTarget) Set(val *CreatePutTargetTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePutTargetTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePutTargetTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePutTargetTarget(val *CreatePutTargetTarget) *NullableCreatePutTargetTarget {
	return &NullableCreatePutTargetTarget{value: val, isSet: true}
}

func (v NullableCreatePutTargetTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePutTargetTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
