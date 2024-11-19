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

// checks if the CreateOrPutTagInactivityAlertTag type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrPutTagInactivityAlertTag{}

// CreateOrPutTagInactivityAlertTag struct for CreateOrPutTagInactivityAlertTag
type CreateOrPutTagInactivityAlertTag struct {
	// The name of the detection rule.
	Name string `json:"name"`
	// Use the `leql` parameter instead of this parameter.
	// Deprecated
	Patterns []string `json:"patterns,omitempty"`
	// The IDs of the logs that the detection rule operates on.
	Sources []SourcesIdArrayInner `json:"sources"`
	// The Always set to \"AlertNotify\".
	Type string `json:"type"`
	// Always set to \"InactivityAlert\".
	SubType string `json:"sub_type"`
	// The description of the detection rule.
	Description *string `json:"description,omitempty"`
	// The notifications attached to the detection rule.
	Actions []CreateTagActionInactivity `json:"actions,omitempty"`
	// A list of key-value pairs that may indicate some auxiliary information about the change detection rule.
	UserData map[string]interface{} `json:"user_data,omitempty"`
	// Defines the duration of inactivity before an alert triggers along with the `timeframe_period` parameter.
	TimeframeValue  *int32           `json:"timeframe_value,omitempty"`
	TimeframePeriod *TimeframePeriod `json:"timeframe_period,omitempty"`
	Leql            InactivityLeql   `json:"leql"`
	// This ensures investigations are ordered by priority in Investigation Management. Must be >=0
	Priority *int32 `json:"priority,omitempty"`
}

type _CreateOrPutTagInactivityAlertTag CreateOrPutTagInactivityAlertTag

// NewCreateOrPutTagInactivityAlertTag instantiates a new CreateOrPutTagInactivityAlertTag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrPutTagInactivityAlertTag(name string, sources []SourcesIdArrayInner, type_ string, subType string, leql InactivityLeql) *CreateOrPutTagInactivityAlertTag {
	this := CreateOrPutTagInactivityAlertTag{}
	this.Name = name
	this.Sources = sources
	this.Type = type_
	this.SubType = subType
	this.Leql = leql
	return &this
}

// NewCreateOrPutTagInactivityAlertTagWithDefaults instantiates a new CreateOrPutTagInactivityAlertTag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrPutTagInactivityAlertTagWithDefaults() *CreateOrPutTagInactivityAlertTag {
	this := CreateOrPutTagInactivityAlertTag{}
	return &this
}

// GetName returns the Name field value
func (o *CreateOrPutTagInactivityAlertTag) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateOrPutTagInactivityAlertTag) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateOrPutTagInactivityAlertTag) SetName(v string) {
	o.Name = v
}

// GetPatterns returns the Patterns field value if set, zero value otherwise.
// Deprecated
func (o *CreateOrPutTagInactivityAlertTag) GetPatterns() []string {
	if o == nil || IsNil(o.Patterns) {
		var ret []string
		return ret
	}
	return o.Patterns
}

// GetPatternsOk returns a tuple with the Patterns field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CreateOrPutTagInactivityAlertTag) GetPatternsOk() ([]string, bool) {
	if o == nil || IsNil(o.Patterns) {
		return nil, false
	}
	return o.Patterns, true
}

// HasPatterns returns a boolean if a field has been set.
func (o *CreateOrPutTagInactivityAlertTag) HasPatterns() bool {
	if o != nil && !IsNil(o.Patterns) {
		return true
	}

	return false
}

// SetPatterns gets a reference to the given []string and assigns it to the Patterns field.
// Deprecated
func (o *CreateOrPutTagInactivityAlertTag) SetPatterns(v []string) {
	o.Patterns = v
}

// GetSources returns the Sources field value
func (o *CreateOrPutTagInactivityAlertTag) GetSources() []SourcesIdArrayInner {
	if o == nil {
		var ret []SourcesIdArrayInner
		return ret
	}

	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value
// and a boolean to check if the value has been set.
func (o *CreateOrPutTagInactivityAlertTag) GetSourcesOk() ([]SourcesIdArrayInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sources, true
}

// SetSources sets field value
func (o *CreateOrPutTagInactivityAlertTag) SetSources(v []SourcesIdArrayInner) {
	o.Sources = v
}

// GetType returns the Type field value
func (o *CreateOrPutTagInactivityAlertTag) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateOrPutTagInactivityAlertTag) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateOrPutTagInactivityAlertTag) SetType(v string) {
	o.Type = v
}

// GetSubType returns the SubType field value
func (o *CreateOrPutTagInactivityAlertTag) GetSubType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubType
}

// GetSubTypeOk returns a tuple with the SubType field value
// and a boolean to check if the value has been set.
func (o *CreateOrPutTagInactivityAlertTag) GetSubTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubType, true
}

// SetSubType sets field value
func (o *CreateOrPutTagInactivityAlertTag) SetSubType(v string) {
	o.SubType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateOrPutTagInactivityAlertTag) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrPutTagInactivityAlertTag) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateOrPutTagInactivityAlertTag) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateOrPutTagInactivityAlertTag) SetDescription(v string) {
	o.Description = &v
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *CreateOrPutTagInactivityAlertTag) GetActions() []CreateTagActionInactivity {
	if o == nil || IsNil(o.Actions) {
		var ret []CreateTagActionInactivity
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrPutTagInactivityAlertTag) GetActionsOk() ([]CreateTagActionInactivity, bool) {
	if o == nil || IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *CreateOrPutTagInactivityAlertTag) HasActions() bool {
	if o != nil && !IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given []CreateTagActionInactivity and assigns it to the Actions field.
func (o *CreateOrPutTagInactivityAlertTag) SetActions(v []CreateTagActionInactivity) {
	o.Actions = v
}

// GetUserData returns the UserData field value if set, zero value otherwise.
func (o *CreateOrPutTagInactivityAlertTag) GetUserData() map[string]interface{} {
	if o == nil || IsNil(o.UserData) {
		var ret map[string]interface{}
		return ret
	}
	return o.UserData
}

// GetUserDataOk returns a tuple with the UserData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrPutTagInactivityAlertTag) GetUserDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.UserData) {
		return map[string]interface{}{}, false
	}
	return o.UserData, true
}

// HasUserData returns a boolean if a field has been set.
func (o *CreateOrPutTagInactivityAlertTag) HasUserData() bool {
	if o != nil && !IsNil(o.UserData) {
		return true
	}

	return false
}

// SetUserData gets a reference to the given map[string]interface{} and assigns it to the UserData field.
func (o *CreateOrPutTagInactivityAlertTag) SetUserData(v map[string]interface{}) {
	o.UserData = v
}

// GetTimeframeValue returns the TimeframeValue field value if set, zero value otherwise.
func (o *CreateOrPutTagInactivityAlertTag) GetTimeframeValue() int32 {
	if o == nil || IsNil(o.TimeframeValue) {
		var ret int32
		return ret
	}
	return *o.TimeframeValue
}

// GetTimeframeValueOk returns a tuple with the TimeframeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrPutTagInactivityAlertTag) GetTimeframeValueOk() (*int32, bool) {
	if o == nil || IsNil(o.TimeframeValue) {
		return nil, false
	}
	return o.TimeframeValue, true
}

// HasTimeframeValue returns a boolean if a field has been set.
func (o *CreateOrPutTagInactivityAlertTag) HasTimeframeValue() bool {
	if o != nil && !IsNil(o.TimeframeValue) {
		return true
	}

	return false
}

// SetTimeframeValue gets a reference to the given int32 and assigns it to the TimeframeValue field.
func (o *CreateOrPutTagInactivityAlertTag) SetTimeframeValue(v int32) {
	o.TimeframeValue = &v
}

// GetTimeframePeriod returns the TimeframePeriod field value if set, zero value otherwise.
func (o *CreateOrPutTagInactivityAlertTag) GetTimeframePeriod() TimeframePeriod {
	if o == nil || IsNil(o.TimeframePeriod) {
		var ret TimeframePeriod
		return ret
	}
	return *o.TimeframePeriod
}

// GetTimeframePeriodOk returns a tuple with the TimeframePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrPutTagInactivityAlertTag) GetTimeframePeriodOk() (*TimeframePeriod, bool) {
	if o == nil || IsNil(o.TimeframePeriod) {
		return nil, false
	}
	return o.TimeframePeriod, true
}

// HasTimeframePeriod returns a boolean if a field has been set.
func (o *CreateOrPutTagInactivityAlertTag) HasTimeframePeriod() bool {
	if o != nil && !IsNil(o.TimeframePeriod) {
		return true
	}

	return false
}

// SetTimeframePeriod gets a reference to the given TimeframePeriod and assigns it to the TimeframePeriod field.
func (o *CreateOrPutTagInactivityAlertTag) SetTimeframePeriod(v TimeframePeriod) {
	o.TimeframePeriod = &v
}

// GetLeql returns the Leql field value
func (o *CreateOrPutTagInactivityAlertTag) GetLeql() InactivityLeql {
	if o == nil {
		var ret InactivityLeql
		return ret
	}

	return o.Leql
}

// GetLeqlOk returns a tuple with the Leql field value
// and a boolean to check if the value has been set.
func (o *CreateOrPutTagInactivityAlertTag) GetLeqlOk() (*InactivityLeql, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Leql, true
}

// SetLeql sets field value
func (o *CreateOrPutTagInactivityAlertTag) SetLeql(v InactivityLeql) {
	o.Leql = v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *CreateOrPutTagInactivityAlertTag) GetPriority() int32 {
	if o == nil || IsNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrPutTagInactivityAlertTag) GetPriorityOk() (*int32, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *CreateOrPutTagInactivityAlertTag) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *CreateOrPutTagInactivityAlertTag) SetPriority(v int32) {
	o.Priority = &v
}

func (o CreateOrPutTagInactivityAlertTag) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrPutTagInactivityAlertTag) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Patterns) {
		toSerialize["patterns"] = o.Patterns
	}
	toSerialize["sources"] = o.Sources
	toSerialize["type"] = o.Type
	toSerialize["sub_type"] = o.SubType
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Actions) {
		toSerialize["actions"] = o.Actions
	}
	if !IsNil(o.UserData) {
		toSerialize["user_data"] = o.UserData
	}
	if !IsNil(o.TimeframeValue) {
		toSerialize["timeframe_value"] = o.TimeframeValue
	}
	if !IsNil(o.TimeframePeriod) {
		toSerialize["timeframe_period"] = o.TimeframePeriod
	}
	toSerialize["leql"] = o.Leql
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	return toSerialize, nil
}

func (o *CreateOrPutTagInactivityAlertTag) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"sources",
		"type",
		"sub_type",
		"leql",
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

	varCreateOrPutTagInactivityAlertTag := _CreateOrPutTagInactivityAlertTag{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateOrPutTagInactivityAlertTag)

	if err != nil {
		return err
	}

	*o = CreateOrPutTagInactivityAlertTag(varCreateOrPutTagInactivityAlertTag)

	return err
}

type NullableCreateOrPutTagInactivityAlertTag struct {
	value *CreateOrPutTagInactivityAlertTag
	isSet bool
}

func (v NullableCreateOrPutTagInactivityAlertTag) Get() *CreateOrPutTagInactivityAlertTag {
	return v.value
}

func (v *NullableCreateOrPutTagInactivityAlertTag) Set(val *CreateOrPutTagInactivityAlertTag) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrPutTagInactivityAlertTag) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrPutTagInactivityAlertTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrPutTagInactivityAlertTag(val *CreateOrPutTagInactivityAlertTag) *NullableCreateOrPutTagInactivityAlertTag {
	return &NullableCreateOrPutTagInactivityAlertTag{value: val, isSet: true}
}

func (v NullableCreateOrPutTagInactivityAlertTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrPutTagInactivityAlertTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
