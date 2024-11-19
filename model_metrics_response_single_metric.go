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

// checks if the MetricsResponseSingleMetric type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricsResponseSingleMetric{}

// MetricsResponseSingleMetric struct for MetricsResponseSingleMetric
type MetricsResponseSingleMetric struct {
	// The UUID of the pre-computed query.
	Id *string `json:"id,omitempty"`
	// The name of the pre-computed query.
	Name *string `json:"name,omitempty"`
	// The description of the pre-computed query.
	Description *string `json:"description,omitempty"`
	// This indicates whether the pre-computed query is enabled or not.
	Enabled interface{} `json:"enabled,omitempty"`
	// The information of the logs used in the pre-computed query.
	Logs []LogInfoResponse `json:"logs,omitempty"`
	// The information on the log sets used in the pre-computed query.
	Logsets []LogsetInfo          `json:"logsets,omitempty"`
	Leql    *LeqlMetricsEndpoints `json:"leql,omitempty"`
	// The time window in seconds, that each PCQ datapoint corresponds to.
	ResolutionSecs *int32 `json:"resolution_secs,omitempty"`
	// The length of time in seconds, that the PCQ will be stored for.
	RetentionSecs *int32 `json:"retention_secs,omitempty"`
	// The time the PCQ was enabled.
	EnabledSince *string `json:"enabled_since,omitempty"`
	// The time the PCQ was created.
	Created *string `json:"created,omitempty"`
}

// NewMetricsResponseSingleMetric instantiates a new MetricsResponseSingleMetric object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsResponseSingleMetric() *MetricsResponseSingleMetric {
	this := MetricsResponseSingleMetric{}
	return &this
}

// NewMetricsResponseSingleMetricWithDefaults instantiates a new MetricsResponseSingleMetric object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsResponseSingleMetricWithDefaults() *MetricsResponseSingleMetric {
	this := MetricsResponseSingleMetric{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MetricsResponseSingleMetric) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsResponseSingleMetric) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MetricsResponseSingleMetric) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MetricsResponseSingleMetric) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MetricsResponseSingleMetric) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsResponseSingleMetric) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MetricsResponseSingleMetric) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MetricsResponseSingleMetric) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MetricsResponseSingleMetric) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsResponseSingleMetric) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MetricsResponseSingleMetric) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MetricsResponseSingleMetric) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetricsResponseSingleMetric) GetEnabled() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetricsResponseSingleMetric) GetEnabledOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return &o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *MetricsResponseSingleMetric) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given interface{} and assigns it to the Enabled field.
func (o *MetricsResponseSingleMetric) SetEnabled(v interface{}) {
	o.Enabled = v
}

// GetLogs returns the Logs field value if set, zero value otherwise.
func (o *MetricsResponseSingleMetric) GetLogs() []LogInfoResponse {
	if o == nil || IsNil(o.Logs) {
		var ret []LogInfoResponse
		return ret
	}
	return o.Logs
}

// GetLogsOk returns a tuple with the Logs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsResponseSingleMetric) GetLogsOk() ([]LogInfoResponse, bool) {
	if o == nil || IsNil(o.Logs) {
		return nil, false
	}
	return o.Logs, true
}

// HasLogs returns a boolean if a field has been set.
func (o *MetricsResponseSingleMetric) HasLogs() bool {
	if o != nil && !IsNil(o.Logs) {
		return true
	}

	return false
}

// SetLogs gets a reference to the given []LogInfoResponse and assigns it to the Logs field.
func (o *MetricsResponseSingleMetric) SetLogs(v []LogInfoResponse) {
	o.Logs = v
}

// GetLogsets returns the Logsets field value if set, zero value otherwise.
func (o *MetricsResponseSingleMetric) GetLogsets() []LogsetInfo {
	if o == nil || IsNil(o.Logsets) {
		var ret []LogsetInfo
		return ret
	}
	return o.Logsets
}

// GetLogsetsOk returns a tuple with the Logsets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsResponseSingleMetric) GetLogsetsOk() ([]LogsetInfo, bool) {
	if o == nil || IsNil(o.Logsets) {
		return nil, false
	}
	return o.Logsets, true
}

// HasLogsets returns a boolean if a field has been set.
func (o *MetricsResponseSingleMetric) HasLogsets() bool {
	if o != nil && !IsNil(o.Logsets) {
		return true
	}

	return false
}

// SetLogsets gets a reference to the given []LogsetInfo and assigns it to the Logsets field.
func (o *MetricsResponseSingleMetric) SetLogsets(v []LogsetInfo) {
	o.Logsets = v
}

// GetLeql returns the Leql field value if set, zero value otherwise.
func (o *MetricsResponseSingleMetric) GetLeql() LeqlMetricsEndpoints {
	if o == nil || IsNil(o.Leql) {
		var ret LeqlMetricsEndpoints
		return ret
	}
	return *o.Leql
}

// GetLeqlOk returns a tuple with the Leql field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsResponseSingleMetric) GetLeqlOk() (*LeqlMetricsEndpoints, bool) {
	if o == nil || IsNil(o.Leql) {
		return nil, false
	}
	return o.Leql, true
}

// HasLeql returns a boolean if a field has been set.
func (o *MetricsResponseSingleMetric) HasLeql() bool {
	if o != nil && !IsNil(o.Leql) {
		return true
	}

	return false
}

// SetLeql gets a reference to the given LeqlMetricsEndpoints and assigns it to the Leql field.
func (o *MetricsResponseSingleMetric) SetLeql(v LeqlMetricsEndpoints) {
	o.Leql = &v
}

// GetResolutionSecs returns the ResolutionSecs field value if set, zero value otherwise.
func (o *MetricsResponseSingleMetric) GetResolutionSecs() int32 {
	if o == nil || IsNil(o.ResolutionSecs) {
		var ret int32
		return ret
	}
	return *o.ResolutionSecs
}

// GetResolutionSecsOk returns a tuple with the ResolutionSecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsResponseSingleMetric) GetResolutionSecsOk() (*int32, bool) {
	if o == nil || IsNil(o.ResolutionSecs) {
		return nil, false
	}
	return o.ResolutionSecs, true
}

// HasResolutionSecs returns a boolean if a field has been set.
func (o *MetricsResponseSingleMetric) HasResolutionSecs() bool {
	if o != nil && !IsNil(o.ResolutionSecs) {
		return true
	}

	return false
}

// SetResolutionSecs gets a reference to the given int32 and assigns it to the ResolutionSecs field.
func (o *MetricsResponseSingleMetric) SetResolutionSecs(v int32) {
	o.ResolutionSecs = &v
}

// GetRetentionSecs returns the RetentionSecs field value if set, zero value otherwise.
func (o *MetricsResponseSingleMetric) GetRetentionSecs() int32 {
	if o == nil || IsNil(o.RetentionSecs) {
		var ret int32
		return ret
	}
	return *o.RetentionSecs
}

// GetRetentionSecsOk returns a tuple with the RetentionSecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsResponseSingleMetric) GetRetentionSecsOk() (*int32, bool) {
	if o == nil || IsNil(o.RetentionSecs) {
		return nil, false
	}
	return o.RetentionSecs, true
}

// HasRetentionSecs returns a boolean if a field has been set.
func (o *MetricsResponseSingleMetric) HasRetentionSecs() bool {
	if o != nil && !IsNil(o.RetentionSecs) {
		return true
	}

	return false
}

// SetRetentionSecs gets a reference to the given int32 and assigns it to the RetentionSecs field.
func (o *MetricsResponseSingleMetric) SetRetentionSecs(v int32) {
	o.RetentionSecs = &v
}

// GetEnabledSince returns the EnabledSince field value if set, zero value otherwise.
func (o *MetricsResponseSingleMetric) GetEnabledSince() string {
	if o == nil || IsNil(o.EnabledSince) {
		var ret string
		return ret
	}
	return *o.EnabledSince
}

// GetEnabledSinceOk returns a tuple with the EnabledSince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsResponseSingleMetric) GetEnabledSinceOk() (*string, bool) {
	if o == nil || IsNil(o.EnabledSince) {
		return nil, false
	}
	return o.EnabledSince, true
}

// HasEnabledSince returns a boolean if a field has been set.
func (o *MetricsResponseSingleMetric) HasEnabledSince() bool {
	if o != nil && !IsNil(o.EnabledSince) {
		return true
	}

	return false
}

// SetEnabledSince gets a reference to the given string and assigns it to the EnabledSince field.
func (o *MetricsResponseSingleMetric) SetEnabledSince(v string) {
	o.EnabledSince = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *MetricsResponseSingleMetric) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsResponseSingleMetric) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *MetricsResponseSingleMetric) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *MetricsResponseSingleMetric) SetCreated(v string) {
	o.Created = &v
}

func (o MetricsResponseSingleMetric) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricsResponseSingleMetric) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Logs) {
		toSerialize["logs"] = o.Logs
	}
	if !IsNil(o.Logsets) {
		toSerialize["logsets"] = o.Logsets
	}
	if !IsNil(o.Leql) {
		toSerialize["leql"] = o.Leql
	}
	if !IsNil(o.ResolutionSecs) {
		toSerialize["resolution_secs"] = o.ResolutionSecs
	}
	if !IsNil(o.RetentionSecs) {
		toSerialize["retention_secs"] = o.RetentionSecs
	}
	if !IsNil(o.EnabledSince) {
		toSerialize["enabled_since"] = o.EnabledSince
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	return toSerialize, nil
}

type NullableMetricsResponseSingleMetric struct {
	value *MetricsResponseSingleMetric
	isSet bool
}

func (v NullableMetricsResponseSingleMetric) Get() *MetricsResponseSingleMetric {
	return v.value
}

func (v *NullableMetricsResponseSingleMetric) Set(val *MetricsResponseSingleMetric) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsResponseSingleMetric) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsResponseSingleMetric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsResponseSingleMetric(val *MetricsResponseSingleMetric) *NullableMetricsResponseSingleMetric {
	return &NullableMetricsResponseSingleMetric{value: val, isSet: true}
}

func (v NullableMetricsResponseSingleMetric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsResponseSingleMetric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
