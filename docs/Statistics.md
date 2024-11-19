# Statistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stats** | [**map[string]Stats**](Stats.md) | Holds the overall result when query does not contain a &#39;groupby&#39; clause.  Examples: * &#x60;\&quot;stats\&quot;: {\&quot;key1\&quot;: {\&quot;sum\&quot;: &lt;RESULT&gt;}}&#x60; for a &#39;calculate(SUM:key1)&#39; query  * &#x60;\&quot;stats\&quot;: {\&quot;number\&quot;: {\&quot;average\&quot;: &lt;RESULT&gt;}}&#x60; for a &#39;calculate(AVERAGE:number)&#39; query  * &#x60;\&quot;stats\&quot;: {\&quot;global_timeseries\&quot;: {\&quot;count\&quot;: &lt;RESULT&gt;}}&#x60; for a &#39;count&#39; query  | 
**Groups** | [**[]map[string]Stats**](map[string]Stats.md) | Holds the overall result for each group in a &#39;groupby&#39; query. | 
**Granularity** | **int32** | The time window in milliseconds for each time slice in the time series. | 
**Timeseries** | [**map[string][]Stats**](array.md) | Holds the query results for each timeslice (each partition of the time_range) (10 by default). For non-&#39;groupby&#39; queries. | 
**GroupsTimeseries** | [**[]map[string]GroupsTimeseriesInnerValue**](map[string]GroupsTimeseriesInnerValue.md) | For &#39;groupby&#39; queries, holds the timeseries object for each group. | 
**From** | **int32** | The start of the time range for the query, as a UNIX timestamp in milliseconds. | 
**To** | **int32** | The end of the time range for the query, as a UNIX timestamp in milliseconds. | 
**Type** | **string** | The type of function performed, for example, \&quot;count\&quot;, \&quot;max\&quot;, \&quot;average\&quot;, \&quot;standarddeviation\&quot;. | 
**Key** | Pointer to **string** | The key which the function of the &#39;calculate&#39; clause is applied to (for example, &#x60;number&#x60; in &#x60;calculate(SUM:number)&#x60;). | [optional] 
**Cardinality** | **int32** | Always 0. (Not implemented) | 
**Others** | **map[string]interface{}** | Not yet implemented. | 
**Status** | **int32** | holds a status code for the query(for example, 200, 202, 204, 400), potentially different from the status code of the response. | 
**AllExactResults** | **NullableBool** | Boolean indicating whether groups are calculated approximately (approximated if a groupby query involves over 10,000 groups). Equal to \&quot;null\&quot; if query is not a groupby query.  | 

## Methods

### NewStatistics

`func NewStatistics(stats map[string]Stats, groups []map[string]Stats, granularity int32, timeseries map[string][]Stats, groupsTimeseries []map[string]GroupsTimeseriesInnerValue, from int32, to int32, type_ string, cardinality int32, others map[string]interface{}, status int32, allExactResults NullableBool, ) *Statistics`

NewStatistics instantiates a new Statistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatisticsWithDefaults

`func NewStatisticsWithDefaults() *Statistics`

NewStatisticsWithDefaults instantiates a new Statistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStats

`func (o *Statistics) GetStats() map[string]Stats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *Statistics) GetStatsOk() (*map[string]Stats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *Statistics) SetStats(v map[string]Stats)`

SetStats sets Stats field to given value.


### GetGroups

`func (o *Statistics) GetGroups() []map[string]Stats`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *Statistics) GetGroupsOk() (*[]map[string]Stats, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *Statistics) SetGroups(v []map[string]Stats)`

SetGroups sets Groups field to given value.


### GetGranularity

`func (o *Statistics) GetGranularity() int32`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *Statistics) GetGranularityOk() (*int32, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *Statistics) SetGranularity(v int32)`

SetGranularity sets Granularity field to given value.


### GetTimeseries

`func (o *Statistics) GetTimeseries() map[string][]Stats`

GetTimeseries returns the Timeseries field if non-nil, zero value otherwise.

### GetTimeseriesOk

`func (o *Statistics) GetTimeseriesOk() (*map[string][]Stats, bool)`

GetTimeseriesOk returns a tuple with the Timeseries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeseries

`func (o *Statistics) SetTimeseries(v map[string][]Stats)`

SetTimeseries sets Timeseries field to given value.


### GetGroupsTimeseries

`func (o *Statistics) GetGroupsTimeseries() []map[string]GroupsTimeseriesInnerValue`

GetGroupsTimeseries returns the GroupsTimeseries field if non-nil, zero value otherwise.

### GetGroupsTimeseriesOk

`func (o *Statistics) GetGroupsTimeseriesOk() (*[]map[string]GroupsTimeseriesInnerValue, bool)`

GetGroupsTimeseriesOk returns a tuple with the GroupsTimeseries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsTimeseries

`func (o *Statistics) SetGroupsTimeseries(v []map[string]GroupsTimeseriesInnerValue)`

SetGroupsTimeseries sets GroupsTimeseries field to given value.


### GetFrom

`func (o *Statistics) GetFrom() int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *Statistics) GetFromOk() (*int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *Statistics) SetFrom(v int32)`

SetFrom sets From field to given value.


### GetTo

`func (o *Statistics) GetTo() int32`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *Statistics) GetToOk() (*int32, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *Statistics) SetTo(v int32)`

SetTo sets To field to given value.


### GetType

`func (o *Statistics) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Statistics) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Statistics) SetType(v string)`

SetType sets Type field to given value.


### GetKey

`func (o *Statistics) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Statistics) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Statistics) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *Statistics) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetCardinality

`func (o *Statistics) GetCardinality() int32`

GetCardinality returns the Cardinality field if non-nil, zero value otherwise.

### GetCardinalityOk

`func (o *Statistics) GetCardinalityOk() (*int32, bool)`

GetCardinalityOk returns a tuple with the Cardinality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardinality

`func (o *Statistics) SetCardinality(v int32)`

SetCardinality sets Cardinality field to given value.


### GetOthers

`func (o *Statistics) GetOthers() map[string]interface{}`

GetOthers returns the Others field if non-nil, zero value otherwise.

### GetOthersOk

`func (o *Statistics) GetOthersOk() (*map[string]interface{}, bool)`

GetOthersOk returns a tuple with the Others field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOthers

`func (o *Statistics) SetOthers(v map[string]interface{})`

SetOthers sets Others field to given value.


### GetStatus

`func (o *Statistics) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Statistics) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Statistics) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetAllExactResults

`func (o *Statistics) GetAllExactResults() bool`

GetAllExactResults returns the AllExactResults field if non-nil, zero value otherwise.

### GetAllExactResultsOk

`func (o *Statistics) GetAllExactResultsOk() (*bool, bool)`

GetAllExactResultsOk returns a tuple with the AllExactResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllExactResults

`func (o *Statistics) SetAllExactResults(v bool)`

SetAllExactResults sets AllExactResults field to given value.


### SetAllExactResultsNil

`func (o *Statistics) SetAllExactResultsNil(b bool)`

 SetAllExactResultsNil sets the value for AllExactResults to be an explicit nil

### UnsetAllExactResults
`func (o *Statistics) UnsetAllExactResults()`

UnsetAllExactResults ensures that no value is present for AllExactResults, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


