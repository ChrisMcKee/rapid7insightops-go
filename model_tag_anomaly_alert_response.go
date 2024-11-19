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

// checks if the TagAnomalyAlertResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TagAnomalyAlertResponse{}

// TagAnomalyAlertResponse struct for TagAnomalyAlertResponse
type TagAnomalyAlertResponse struct {
	// Always set to \"AlertNotify\".
	Type *string `json:"type,omitempty"`
	// The unique ID of the detection rule.
	Id *string `json:"id,omitempty"`
	// The name of the detection rule.
	Name *string `json:"name,omitempty"`
	// The description of the detection rule.
	Description *string `json:"description,omitempty"`
	// The collection of logs the detection rule operates on.
	Sources []TagSourceResponse `json:"sources,omitempty"`
	// The notifications attached to the detection rule.
	Actions []ActionResponseAnomaly `json:"actions,omitempty"`
	// Use the `leql` parameter instead of this parameter.
	// Deprecated
	Patterns []string `json:"patterns,omitempty"`
	// A list of key-value pairs that may indicate some auxiliary information about the change detection rule.
	UserData map[string]interface{} `json:"user_data,omitempty"`
	// ID of the scheduled query associated with this change detection rule.
	ScheduledQueryId *string `json:"scheduled_query_id,omitempty"`
	// Always set to \"AnomalyAlert\".
	SubType        *string                 `json:"sub_type,omitempty"`
	ScheduledQuery *ScheduledQueryResponse `json:"scheduled_query,omitempty"`
	// The percentage that defines when to trigger notifications. The percentage can be positive or negative. For example, if the percentage is set to -50, then notifications will be triggered when the query result for the time range defined by the `time_value` and `time_period` parameters is 50% smaller when compared to the same query over the time range defined by the `timeframe_value` and `timeframe_period` parameters.
	ThresholdValue *int32 `json:"threshold_value,omitempty"`
	// This ensures investigations are ordered by priority in Investigation Management. Must be >=0
	Priority *int32 `json:"priority,omitempty"`
}

// NewTagAnomalyAlertResponse instantiates a new TagAnomalyAlertResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagAnomalyAlertResponse() *TagAnomalyAlertResponse {
	this := TagAnomalyAlertResponse{}
	return &this
}

// NewTagAnomalyAlertResponseWithDefaults instantiates a new TagAnomalyAlertResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagAnomalyAlertResponseWithDefaults() *TagAnomalyAlertResponse {
	this := TagAnomalyAlertResponse{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TagAnomalyAlertResponse) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagAnomalyAlertResponse) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TagAnomalyAlertResponse) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TagAnomalyAlertResponse) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TagAnomalyAlertResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagAnomalyAlertResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TagAnomalyAlertResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TagAnomalyAlertResponse) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TagAnomalyAlertResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagAnomalyAlertResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TagAnomalyAlertResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TagAnomalyAlertResponse) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TagAnomalyAlertResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagAnomalyAlertResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TagAnomalyAlertResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TagAnomalyAlertResponse) SetDescription(v string) {
	o.Description = &v
}

// GetSources returns the Sources field value if set, zero value otherwise.
func (o *TagAnomalyAlertResponse) GetSources() []TagSourceResponse {
	if o == nil || IsNil(o.Sources) {
		var ret []TagSourceResponse
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagAnomalyAlertResponse) GetSourcesOk() ([]TagSourceResponse, bool) {
	if o == nil || IsNil(o.Sources) {
		return nil, false
	}
	return o.Sources, true
}

// HasSources returns a boolean if a field has been set.
func (o *TagAnomalyAlertResponse) HasSources() bool {
	if o != nil && !IsNil(o.Sources) {
		return true
	}

	return false
}

// SetSources gets a reference to the given []TagSourceResponse and assigns it to the Sources field.
func (o *TagAnomalyAlertResponse) SetSources(v []TagSourceResponse) {
	o.Sources = v
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *TagAnomalyAlertResponse) GetActions() []ActionResponseAnomaly {
	if o == nil || IsNil(o.Actions) {
		var ret []ActionResponseAnomaly
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagAnomalyAlertResponse) GetActionsOk() ([]ActionResponseAnomaly, bool) {
	if o == nil || IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *TagAnomalyAlertResponse) HasActions() bool {
	if o != nil && !IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given []ActionResponseAnomaly and assigns it to the Actions field.
func (o *TagAnomalyAlertResponse) SetActions(v []ActionResponseAnomaly) {
	o.Actions = v
}

// GetPatterns returns the Patterns field value if set, zero value otherwise.
// Deprecated
func (o *TagAnomalyAlertResponse) GetPatterns() []string {
	if o == nil || IsNil(o.Patterns) {
		var ret []string
		return ret
	}
	return o.Patterns
}

// GetPatternsOk returns a tuple with the Patterns field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *TagAnomalyAlertResponse) GetPatternsOk() ([]string, bool) {
	if o == nil || IsNil(o.Patterns) {
		return nil, false
	}
	return o.Patterns, true
}

// HasPatterns returns a boolean if a field has been set.
func (o *TagAnomalyAlertResponse) HasPatterns() bool {
	if o != nil && !IsNil(o.Patterns) {
		return true
	}

	return false
}

// SetPatterns gets a reference to the given []string and assigns it to the Patterns field.
// Deprecated
func (o *TagAnomalyAlertResponse) SetPatterns(v []string) {
	o.Patterns = v
}

// GetUserData returns the UserData field value if set, zero value otherwise.
func (o *TagAnomalyAlertResponse) GetUserData() map[string]interface{} {
	if o == nil || IsNil(o.UserData) {
		var ret map[string]interface{}
		return ret
	}
	return o.UserData
}

// GetUserDataOk returns a tuple with the UserData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagAnomalyAlertResponse) GetUserDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.UserData) {
		return map[string]interface{}{}, false
	}
	return o.UserData, true
}

// HasUserData returns a boolean if a field has been set.
func (o *TagAnomalyAlertResponse) HasUserData() bool {
	if o != nil && !IsNil(o.UserData) {
		return true
	}

	return false
}

// SetUserData gets a reference to the given map[string]interface{} and assigns it to the UserData field.
func (o *TagAnomalyAlertResponse) SetUserData(v map[string]interface{}) {
	o.UserData = v
}

// GetScheduledQueryId returns the ScheduledQueryId field value if set, zero value otherwise.
func (o *TagAnomalyAlertResponse) GetScheduledQueryId() string {
	if o == nil || IsNil(o.ScheduledQueryId) {
		var ret string
		return ret
	}
	return *o.ScheduledQueryId
}

// GetScheduledQueryIdOk returns a tuple with the ScheduledQueryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagAnomalyAlertResponse) GetScheduledQueryIdOk() (*string, bool) {
	if o == nil || IsNil(o.ScheduledQueryId) {
		return nil, false
	}
	return o.ScheduledQueryId, true
}

// HasScheduledQueryId returns a boolean if a field has been set.
func (o *TagAnomalyAlertResponse) HasScheduledQueryId() bool {
	if o != nil && !IsNil(o.ScheduledQueryId) {
		return true
	}

	return false
}

// SetScheduledQueryId gets a reference to the given string and assigns it to the ScheduledQueryId field.
func (o *TagAnomalyAlertResponse) SetScheduledQueryId(v string) {
	o.ScheduledQueryId = &v
}

// GetSubType returns the SubType field value if set, zero value otherwise.
func (o *TagAnomalyAlertResponse) GetSubType() string {
	if o == nil || IsNil(o.SubType) {
		var ret string
		return ret
	}
	return *o.SubType
}

// GetSubTypeOk returns a tuple with the SubType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagAnomalyAlertResponse) GetSubTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SubType) {
		return nil, false
	}
	return o.SubType, true
}

// HasSubType returns a boolean if a field has been set.
func (o *TagAnomalyAlertResponse) HasSubType() bool {
	if o != nil && !IsNil(o.SubType) {
		return true
	}

	return false
}

// SetSubType gets a reference to the given string and assigns it to the SubType field.
func (o *TagAnomalyAlertResponse) SetSubType(v string) {
	o.SubType = &v
}

// GetScheduledQuery returns the ScheduledQuery field value if set, zero value otherwise.
func (o *TagAnomalyAlertResponse) GetScheduledQuery() ScheduledQueryResponse {
	if o == nil || IsNil(o.ScheduledQuery) {
		var ret ScheduledQueryResponse
		return ret
	}
	return *o.ScheduledQuery
}

// GetScheduledQueryOk returns a tuple with the ScheduledQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagAnomalyAlertResponse) GetScheduledQueryOk() (*ScheduledQueryResponse, bool) {
	if o == nil || IsNil(o.ScheduledQuery) {
		return nil, false
	}
	return o.ScheduledQuery, true
}

// HasScheduledQuery returns a boolean if a field has been set.
func (o *TagAnomalyAlertResponse) HasScheduledQuery() bool {
	if o != nil && !IsNil(o.ScheduledQuery) {
		return true
	}

	return false
}

// SetScheduledQuery gets a reference to the given ScheduledQueryResponse and assigns it to the ScheduledQuery field.
func (o *TagAnomalyAlertResponse) SetScheduledQuery(v ScheduledQueryResponse) {
	o.ScheduledQuery = &v
}

// GetThresholdValue returns the ThresholdValue field value if set, zero value otherwise.
func (o *TagAnomalyAlertResponse) GetThresholdValue() int32 {
	if o == nil || IsNil(o.ThresholdValue) {
		var ret int32
		return ret
	}
	return *o.ThresholdValue
}

// GetThresholdValueOk returns a tuple with the ThresholdValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagAnomalyAlertResponse) GetThresholdValueOk() (*int32, bool) {
	if o == nil || IsNil(o.ThresholdValue) {
		return nil, false
	}
	return o.ThresholdValue, true
}

// HasThresholdValue returns a boolean if a field has been set.
func (o *TagAnomalyAlertResponse) HasThresholdValue() bool {
	if o != nil && !IsNil(o.ThresholdValue) {
		return true
	}

	return false
}

// SetThresholdValue gets a reference to the given int32 and assigns it to the ThresholdValue field.
func (o *TagAnomalyAlertResponse) SetThresholdValue(v int32) {
	o.ThresholdValue = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *TagAnomalyAlertResponse) GetPriority() int32 {
	if o == nil || IsNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagAnomalyAlertResponse) GetPriorityOk() (*int32, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *TagAnomalyAlertResponse) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *TagAnomalyAlertResponse) SetPriority(v int32) {
	o.Priority = &v
}

func (o TagAnomalyAlertResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagAnomalyAlertResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Sources) {
		toSerialize["sources"] = o.Sources
	}
	if !IsNil(o.Actions) {
		toSerialize["actions"] = o.Actions
	}
	if !IsNil(o.Patterns) {
		toSerialize["patterns"] = o.Patterns
	}
	if !IsNil(o.UserData) {
		toSerialize["user_data"] = o.UserData
	}
	if !IsNil(o.ScheduledQueryId) {
		toSerialize["scheduled_query_id"] = o.ScheduledQueryId
	}
	if !IsNil(o.SubType) {
		toSerialize["sub_type"] = o.SubType
	}
	if !IsNil(o.ScheduledQuery) {
		toSerialize["scheduled_query"] = o.ScheduledQuery
	}
	if !IsNil(o.ThresholdValue) {
		toSerialize["threshold_value"] = o.ThresholdValue
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	return toSerialize, nil
}

type NullableTagAnomalyAlertResponse struct {
	value *TagAnomalyAlertResponse
	isSet bool
}

func (v NullableTagAnomalyAlertResponse) Get() *TagAnomalyAlertResponse {
	return v.value
}

func (v *NullableTagAnomalyAlertResponse) Set(val *TagAnomalyAlertResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTagAnomalyAlertResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTagAnomalyAlertResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagAnomalyAlertResponse(val *TagAnomalyAlertResponse) *NullableTagAnomalyAlertResponse {
	return &NullableTagAnomalyAlertResponse{value: val, isSet: true}
}

func (v NullableTagAnomalyAlertResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagAnomalyAlertResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}