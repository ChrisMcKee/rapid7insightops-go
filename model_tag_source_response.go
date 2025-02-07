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

// checks if the TagSourceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TagSourceResponse{}

// TagSourceResponse struct for TagSourceResponse
type TagSourceResponse struct {
	// The UUID of the log.
	Id              *string              `json:"id,omitempty"`
	RetentionPeriod *RetentionPeriodEnum `json:"retention_period,omitempty"`
	// The log token(s) used for writing to the log. This only applies to token type logs (view the `source_type` parameter).
	Token *string `json:"token,omitempty"`
	// The name of the log.
	Name *string `json:"name,omitempty"`
	// The number of days stored.
	StoredDays interface{} `json:"stored_days,omitempty"`
}

// NewTagSourceResponse instantiates a new TagSourceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagSourceResponse() *TagSourceResponse {
	this := TagSourceResponse{}
	return &this
}

// NewTagSourceResponseWithDefaults instantiates a new TagSourceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagSourceResponseWithDefaults() *TagSourceResponse {
	this := TagSourceResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TagSourceResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagSourceResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TagSourceResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TagSourceResponse) SetId(v string) {
	o.Id = &v
}

// GetRetentionPeriod returns the RetentionPeriod field value if set, zero value otherwise.
func (o *TagSourceResponse) GetRetentionPeriod() RetentionPeriodEnum {
	if o == nil || IsNil(o.RetentionPeriod) {
		var ret RetentionPeriodEnum
		return ret
	}
	return *o.RetentionPeriod
}

// GetRetentionPeriodOk returns a tuple with the RetentionPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagSourceResponse) GetRetentionPeriodOk() (*RetentionPeriodEnum, bool) {
	if o == nil || IsNil(o.RetentionPeriod) {
		return nil, false
	}
	return o.RetentionPeriod, true
}

// HasRetentionPeriod returns a boolean if a field has been set.
func (o *TagSourceResponse) HasRetentionPeriod() bool {
	if o != nil && !IsNil(o.RetentionPeriod) {
		return true
	}

	return false
}

// SetRetentionPeriod gets a reference to the given RetentionPeriodEnum and assigns it to the RetentionPeriod field.
func (o *TagSourceResponse) SetRetentionPeriod(v RetentionPeriodEnum) {
	o.RetentionPeriod = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *TagSourceResponse) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagSourceResponse) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *TagSourceResponse) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *TagSourceResponse) SetToken(v string) {
	o.Token = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TagSourceResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagSourceResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TagSourceResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TagSourceResponse) SetName(v string) {
	o.Name = &v
}

// GetStoredDays returns the StoredDays field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TagSourceResponse) GetStoredDays() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.StoredDays
}

// GetStoredDaysOk returns a tuple with the StoredDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TagSourceResponse) GetStoredDaysOk() (*interface{}, bool) {
	if o == nil || IsNil(o.StoredDays) {
		return nil, false
	}
	return &o.StoredDays, true
}

// HasStoredDays returns a boolean if a field has been set.
func (o *TagSourceResponse) HasStoredDays() bool {
	if o != nil && !IsNil(o.StoredDays) {
		return true
	}

	return false
}

// SetStoredDays gets a reference to the given interface{} and assigns it to the StoredDays field.
func (o *TagSourceResponse) SetStoredDays(v interface{}) {
	o.StoredDays = v
}

func (o TagSourceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagSourceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.RetentionPeriod) {
		toSerialize["retention_period"] = o.RetentionPeriod
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.StoredDays != nil {
		toSerialize["stored_days"] = o.StoredDays
	}
	return toSerialize, nil
}

type NullableTagSourceResponse struct {
	value *TagSourceResponse
	isSet bool
}

func (v NullableTagSourceResponse) Get() *TagSourceResponse {
	return v.value
}

func (v *NullableTagSourceResponse) Set(val *TagSourceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTagSourceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTagSourceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagSourceResponse(val *TagSourceResponse) *NullableTagSourceResponse {
	return &NullableTagSourceResponse{value: val, isSet: true}
}

func (v NullableTagSourceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagSourceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
