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
	"fmt"
)

// checks if the Statistics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Statistics{}

// Statistics The result of the 'calculate' function and/or any 'groupby' queries. 202 responses may provide a partial result.
type Statistics struct {
	// Holds the overall result when query does not contain a 'groupby' clause.  Examples: * `\"stats\": {\"key1\": {\"sum\": <RESULT>}}` for a 'calculate(SUM:key1)' query  * `\"stats\": {\"number\": {\"average\": <RESULT>}}` for a 'calculate(AVERAGE:number)' query  * `\"stats\": {\"global_timeseries\": {\"count\": <RESULT>}}` for a 'count' query
	Stats map[string]Stats `json:"stats"`
	// Holds the overall result for each group in a 'groupby' query.
	Groups []map[string]Stats `json:"groups"`
	// The time window in milliseconds for each time slice in the time series.
	Granularity int32 `json:"granularity"`
	// Holds the query results for each timeslice (each partition of the time_range) (10 by default). For non-'groupby' queries.
	Timeseries map[string][]Stats `json:"timeseries"`
	// For 'groupby' queries, holds the timeseries object for each group.
	GroupsTimeseries []map[string]GroupsTimeseriesInnerValue `json:"groups_timeseries"`
	// The start of the time range for the query, as a UNIX timestamp in milliseconds.
	From int32 `json:"from"`
	// The end of the time range for the query, as a UNIX timestamp in milliseconds.
	To int32 `json:"to"`
	// The type of function performed, for example, \"count\", \"max\", \"average\", \"standarddeviation\".
	Type string `json:"type"`
	// The key which the function of the 'calculate' clause is applied to (for example, `number` in `calculate(SUM:number)`).
	Key *string `json:"key,omitempty"`
	// Always 0. (Not implemented)
	Cardinality int32 `json:"cardinality"`
	// Not yet implemented.
	Others map[string]interface{} `json:"others"`
	// holds a status code for the query(for example, 200, 202, 204, 400), potentially different from the status code of the response.
	Status int32 `json:"status"`
	// Boolean indicating whether groups are calculated approximately (approximated if a groupby query involves over 10,000 groups). Equal to \"null\" if query is not a groupby query.
	AllExactResults      NullableBool `json:"all_exact_results"`
	AdditionalProperties map[string]interface{}
}

type _Statistics Statistics

// NewStatistics instantiates a new Statistics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatistics(stats map[string]Stats, groups []map[string]Stats, granularity int32, timeseries map[string][]Stats, groupsTimeseries []map[string]GroupsTimeseriesInnerValue, from int32, to int32, type_ string, cardinality int32, others map[string]interface{}, status int32, allExactResults NullableBool) *Statistics {
	this := Statistics{}
	this.Stats = stats
	this.Groups = groups
	this.Granularity = granularity
	this.Timeseries = timeseries
	this.GroupsTimeseries = groupsTimeseries
	this.From = from
	this.To = to
	this.Type = type_
	this.Cardinality = cardinality
	this.Others = others
	this.Status = status
	this.AllExactResults = allExactResults
	return &this
}

// NewStatisticsWithDefaults instantiates a new Statistics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatisticsWithDefaults() *Statistics {
	this := Statistics{}
	return &this
}

// GetStats returns the Stats field value
func (o *Statistics) GetStats() map[string]Stats {
	if o == nil {
		var ret map[string]Stats
		return ret
	}

	return o.Stats
}

// GetStatsOk returns a tuple with the Stats field value
// and a boolean to check if the value has been set.
func (o *Statistics) GetStatsOk() (*map[string]Stats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stats, true
}

// SetStats sets field value
func (o *Statistics) SetStats(v map[string]Stats) {
	o.Stats = v
}

// GetGroups returns the Groups field value
func (o *Statistics) GetGroups() []map[string]Stats {
	if o == nil {
		var ret []map[string]Stats
		return ret
	}

	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value
// and a boolean to check if the value has been set.
func (o *Statistics) GetGroupsOk() ([]map[string]Stats, bool) {
	if o == nil {
		return nil, false
	}
	return o.Groups, true
}

// SetGroups sets field value
func (o *Statistics) SetGroups(v []map[string]Stats) {
	o.Groups = v
}

// GetGranularity returns the Granularity field value
func (o *Statistics) GetGranularity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Granularity
}

// GetGranularityOk returns a tuple with the Granularity field value
// and a boolean to check if the value has been set.
func (o *Statistics) GetGranularityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Granularity, true
}

// SetGranularity sets field value
func (o *Statistics) SetGranularity(v int32) {
	o.Granularity = v
}

// GetTimeseries returns the Timeseries field value
func (o *Statistics) GetTimeseries() map[string][]Stats {
	if o == nil {
		var ret map[string][]Stats
		return ret
	}

	return o.Timeseries
}

// GetTimeseriesOk returns a tuple with the Timeseries field value
// and a boolean to check if the value has been set.
func (o *Statistics) GetTimeseriesOk() (*map[string][]Stats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timeseries, true
}

// SetTimeseries sets field value
func (o *Statistics) SetTimeseries(v map[string][]Stats) {
	o.Timeseries = v
}

// GetGroupsTimeseries returns the GroupsTimeseries field value
func (o *Statistics) GetGroupsTimeseries() []map[string]GroupsTimeseriesInnerValue {
	if o == nil {
		var ret []map[string]GroupsTimeseriesInnerValue
		return ret
	}

	return o.GroupsTimeseries
}

// GetGroupsTimeseriesOk returns a tuple with the GroupsTimeseries field value
// and a boolean to check if the value has been set.
func (o *Statistics) GetGroupsTimeseriesOk() ([]map[string]GroupsTimeseriesInnerValue, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupsTimeseries, true
}

// SetGroupsTimeseries sets field value
func (o *Statistics) SetGroupsTimeseries(v []map[string]GroupsTimeseriesInnerValue) {
	o.GroupsTimeseries = v
}

// GetFrom returns the From field value
func (o *Statistics) GetFrom() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *Statistics) GetFromOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *Statistics) SetFrom(v int32) {
	o.From = v
}

// GetTo returns the To field value
func (o *Statistics) GetTo() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *Statistics) GetToOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *Statistics) SetTo(v int32) {
	o.To = v
}

// GetType returns the Type field value
func (o *Statistics) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Statistics) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Statistics) SetType(v string) {
	o.Type = v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *Statistics) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Statistics) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *Statistics) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *Statistics) SetKey(v string) {
	o.Key = &v
}

// GetCardinality returns the Cardinality field value
func (o *Statistics) GetCardinality() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Cardinality
}

// GetCardinalityOk returns a tuple with the Cardinality field value
// and a boolean to check if the value has been set.
func (o *Statistics) GetCardinalityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cardinality, true
}

// SetCardinality sets field value
func (o *Statistics) SetCardinality(v int32) {
	o.Cardinality = v
}

// GetOthers returns the Others field value
func (o *Statistics) GetOthers() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Others
}

// GetOthersOk returns a tuple with the Others field value
// and a boolean to check if the value has been set.
func (o *Statistics) GetOthersOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Others, true
}

// SetOthers sets field value
func (o *Statistics) SetOthers(v map[string]interface{}) {
	o.Others = v
}

// GetStatus returns the Status field value
func (o *Statistics) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Statistics) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Statistics) SetStatus(v int32) {
	o.Status = v
}

// GetAllExactResults returns the AllExactResults field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *Statistics) GetAllExactResults() bool {
	if o == nil || o.AllExactResults.Get() == nil {
		var ret bool
		return ret
	}

	return *o.AllExactResults.Get()
}

// GetAllExactResultsOk returns a tuple with the AllExactResults field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Statistics) GetAllExactResultsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllExactResults.Get(), o.AllExactResults.IsSet()
}

// SetAllExactResults sets field value
func (o *Statistics) SetAllExactResults(v bool) {
	o.AllExactResults.Set(&v)
}

func (o Statistics) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Statistics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["stats"] = o.Stats
	toSerialize["groups"] = o.Groups
	toSerialize["granularity"] = o.Granularity
	toSerialize["timeseries"] = o.Timeseries
	toSerialize["groups_timeseries"] = o.GroupsTimeseries
	toSerialize["from"] = o.From
	toSerialize["to"] = o.To
	toSerialize["type"] = o.Type
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	toSerialize["cardinality"] = o.Cardinality
	toSerialize["others"] = o.Others
	toSerialize["status"] = o.Status
	toSerialize["all_exact_results"] = o.AllExactResults.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Statistics) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"stats",
		"groups",
		"granularity",
		"timeseries",
		"groups_timeseries",
		"from",
		"to",
		"type",
		"cardinality",
		"others",
		"status",
		"all_exact_results",
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

	varStatistics := _Statistics{}

	err = json.Unmarshal(data, &varStatistics)

	if err != nil {
		return err
	}

	*o = Statistics(varStatistics)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "stats")
		delete(additionalProperties, "groups")
		delete(additionalProperties, "granularity")
		delete(additionalProperties, "timeseries")
		delete(additionalProperties, "groups_timeseries")
		delete(additionalProperties, "from")
		delete(additionalProperties, "to")
		delete(additionalProperties, "type")
		delete(additionalProperties, "key")
		delete(additionalProperties, "cardinality")
		delete(additionalProperties, "others")
		delete(additionalProperties, "status")
		delete(additionalProperties, "all_exact_results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStatistics struct {
	value *Statistics
	isSet bool
}

func (v NullableStatistics) Get() *Statistics {
	return v.value
}

func (v *NullableStatistics) Set(val *Statistics) {
	v.value = val
	v.isSet = true
}

func (v NullableStatistics) IsSet() bool {
	return v.isSet
}

func (v *NullableStatistics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatistics(val *Statistics) *NullableStatistics {
	return &NullableStatistics{value: val, isSet: true}
}

func (v NullableStatistics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatistics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
