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

// checks if the GetTargetAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTargetAction{}

// GetTargetAction struct for GetTargetAction
type GetTargetAction struct {
	// The name of the notification target.
	Name *string `json:"name,omitempty"`
	// The description of the notifiaction target.
	Description     *string                  `json:"description,omitempty"`
	Type            *NotificationTypeEnum    `json:"type,omitempty"`
	ParamsSet       *TargetResponseParamsSet `json:"params_set,omitempty"`
	AlertContentSet *TargetAlertContentSet   `json:"alert_content_set,omitempty"`
	// A list of key-value pairs that may indicate some auxiliary information about the notification target.
	UserData map[string]interface{} `json:"user_data,omitempty"`
}

// NewGetTargetAction instantiates a new GetTargetAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTargetAction() *GetTargetAction {
	this := GetTargetAction{}
	return &this
}

// NewGetTargetActionWithDefaults instantiates a new GetTargetAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTargetActionWithDefaults() *GetTargetAction {
	this := GetTargetAction{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetTargetAction) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTargetAction) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetTargetAction) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetTargetAction) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GetTargetAction) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTargetAction) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GetTargetAction) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GetTargetAction) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetTargetAction) GetType() NotificationTypeEnum {
	if o == nil || IsNil(o.Type) {
		var ret NotificationTypeEnum
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTargetAction) GetTypeOk() (*NotificationTypeEnum, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetTargetAction) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given NotificationTypeEnum and assigns it to the Type field.
func (o *GetTargetAction) SetType(v NotificationTypeEnum) {
	o.Type = &v
}

// GetParamsSet returns the ParamsSet field value if set, zero value otherwise.
func (o *GetTargetAction) GetParamsSet() TargetResponseParamsSet {
	if o == nil || IsNil(o.ParamsSet) {
		var ret TargetResponseParamsSet
		return ret
	}
	return *o.ParamsSet
}

// GetParamsSetOk returns a tuple with the ParamsSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTargetAction) GetParamsSetOk() (*TargetResponseParamsSet, bool) {
	if o == nil || IsNil(o.ParamsSet) {
		return nil, false
	}
	return o.ParamsSet, true
}

// HasParamsSet returns a boolean if a field has been set.
func (o *GetTargetAction) HasParamsSet() bool {
	if o != nil && !IsNil(o.ParamsSet) {
		return true
	}

	return false
}

// SetParamsSet gets a reference to the given TargetResponseParamsSet and assigns it to the ParamsSet field.
func (o *GetTargetAction) SetParamsSet(v TargetResponseParamsSet) {
	o.ParamsSet = &v
}

// GetAlertContentSet returns the AlertContentSet field value if set, zero value otherwise.
func (o *GetTargetAction) GetAlertContentSet() TargetAlertContentSet {
	if o == nil || IsNil(o.AlertContentSet) {
		var ret TargetAlertContentSet
		return ret
	}
	return *o.AlertContentSet
}

// GetAlertContentSetOk returns a tuple with the AlertContentSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTargetAction) GetAlertContentSetOk() (*TargetAlertContentSet, bool) {
	if o == nil || IsNil(o.AlertContentSet) {
		return nil, false
	}
	return o.AlertContentSet, true
}

// HasAlertContentSet returns a boolean if a field has been set.
func (o *GetTargetAction) HasAlertContentSet() bool {
	if o != nil && !IsNil(o.AlertContentSet) {
		return true
	}

	return false
}

// SetAlertContentSet gets a reference to the given TargetAlertContentSet and assigns it to the AlertContentSet field.
func (o *GetTargetAction) SetAlertContentSet(v TargetAlertContentSet) {
	o.AlertContentSet = &v
}

// GetUserData returns the UserData field value if set, zero value otherwise.
func (o *GetTargetAction) GetUserData() map[string]interface{} {
	if o == nil || IsNil(o.UserData) {
		var ret map[string]interface{}
		return ret
	}
	return o.UserData
}

// GetUserDataOk returns a tuple with the UserData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTargetAction) GetUserDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.UserData) {
		return map[string]interface{}{}, false
	}
	return o.UserData, true
}

// HasUserData returns a boolean if a field has been set.
func (o *GetTargetAction) HasUserData() bool {
	if o != nil && !IsNil(o.UserData) {
		return true
	}

	return false
}

// SetUserData gets a reference to the given map[string]interface{} and assigns it to the UserData field.
func (o *GetTargetAction) SetUserData(v map[string]interface{}) {
	o.UserData = v
}

func (o GetTargetAction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTargetAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.ParamsSet) {
		toSerialize["params_set"] = o.ParamsSet
	}
	if !IsNil(o.AlertContentSet) {
		toSerialize["alert_content_set"] = o.AlertContentSet
	}
	if !IsNil(o.UserData) {
		toSerialize["user_data"] = o.UserData
	}
	return toSerialize, nil
}

type NullableGetTargetAction struct {
	value *GetTargetAction
	isSet bool
}

func (v NullableGetTargetAction) Get() *GetTargetAction {
	return v.value
}

func (v *NullableGetTargetAction) Set(val *GetTargetAction) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTargetAction) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTargetAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTargetAction(val *GetTargetAction) *NullableGetTargetAction {
	return &NullableGetTargetAction{value: val, isSet: true}
}

func (v NullableGetTargetAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTargetAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}