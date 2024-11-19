# SearchStatsInnerStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DurationMs** | Pointer to **int32** | The time taken to complete the query, in milliseconds. | [optional] 
**BytesAll** | Pointer to **int32** | The total amount of log data queried in bytes. | [optional] 
**BytesChecked** | Pointer to **int32** | The amount of log data that was scanned by this query in bytes. This amount may be much less than the total amount of log data that is queried. This is a result of data indexing performed by Log Search when the LEQL statement matches a small subset of the data. | [optional] 
**EventsAll** | Pointer to **int32** | The total number of log entries queried. | [optional] 
**EventsChecked** | Pointer to **int32** | The number of log entries that were scanned by this query. This amount may be much less than the total amount of log data that is queried. This is a result of data indexing performed by Log Search when the LEQL statement matches a small subset of the data. | [optional] 
**EventsMatched** | Pointer to **int32** | The number of log entries that matched the LEQL statement of the query. | [optional] 

## Methods

### NewSearchStatsInnerStatistics

`func NewSearchStatsInnerStatistics() *SearchStatsInnerStatistics`

NewSearchStatsInnerStatistics instantiates a new SearchStatsInnerStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchStatsInnerStatisticsWithDefaults

`func NewSearchStatsInnerStatisticsWithDefaults() *SearchStatsInnerStatistics`

NewSearchStatsInnerStatisticsWithDefaults instantiates a new SearchStatsInnerStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDurationMs

`func (o *SearchStatsInnerStatistics) GetDurationMs() int32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *SearchStatsInnerStatistics) GetDurationMsOk() (*int32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *SearchStatsInnerStatistics) SetDurationMs(v int32)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *SearchStatsInnerStatistics) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.

### GetBytesAll

`func (o *SearchStatsInnerStatistics) GetBytesAll() int32`

GetBytesAll returns the BytesAll field if non-nil, zero value otherwise.

### GetBytesAllOk

`func (o *SearchStatsInnerStatistics) GetBytesAllOk() (*int32, bool)`

GetBytesAllOk returns a tuple with the BytesAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesAll

`func (o *SearchStatsInnerStatistics) SetBytesAll(v int32)`

SetBytesAll sets BytesAll field to given value.

### HasBytesAll

`func (o *SearchStatsInnerStatistics) HasBytesAll() bool`

HasBytesAll returns a boolean if a field has been set.

### GetBytesChecked

`func (o *SearchStatsInnerStatistics) GetBytesChecked() int32`

GetBytesChecked returns the BytesChecked field if non-nil, zero value otherwise.

### GetBytesCheckedOk

`func (o *SearchStatsInnerStatistics) GetBytesCheckedOk() (*int32, bool)`

GetBytesCheckedOk returns a tuple with the BytesChecked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesChecked

`func (o *SearchStatsInnerStatistics) SetBytesChecked(v int32)`

SetBytesChecked sets BytesChecked field to given value.

### HasBytesChecked

`func (o *SearchStatsInnerStatistics) HasBytesChecked() bool`

HasBytesChecked returns a boolean if a field has been set.

### GetEventsAll

`func (o *SearchStatsInnerStatistics) GetEventsAll() int32`

GetEventsAll returns the EventsAll field if non-nil, zero value otherwise.

### GetEventsAllOk

`func (o *SearchStatsInnerStatistics) GetEventsAllOk() (*int32, bool)`

GetEventsAllOk returns a tuple with the EventsAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsAll

`func (o *SearchStatsInnerStatistics) SetEventsAll(v int32)`

SetEventsAll sets EventsAll field to given value.

### HasEventsAll

`func (o *SearchStatsInnerStatistics) HasEventsAll() bool`

HasEventsAll returns a boolean if a field has been set.

### GetEventsChecked

`func (o *SearchStatsInnerStatistics) GetEventsChecked() int32`

GetEventsChecked returns the EventsChecked field if non-nil, zero value otherwise.

### GetEventsCheckedOk

`func (o *SearchStatsInnerStatistics) GetEventsCheckedOk() (*int32, bool)`

GetEventsCheckedOk returns a tuple with the EventsChecked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsChecked

`func (o *SearchStatsInnerStatistics) SetEventsChecked(v int32)`

SetEventsChecked sets EventsChecked field to given value.

### HasEventsChecked

`func (o *SearchStatsInnerStatistics) HasEventsChecked() bool`

HasEventsChecked returns a boolean if a field has been set.

### GetEventsMatched

`func (o *SearchStatsInnerStatistics) GetEventsMatched() int32`

GetEventsMatched returns the EventsMatched field if non-nil, zero value otherwise.

### GetEventsMatchedOk

`func (o *SearchStatsInnerStatistics) GetEventsMatchedOk() (*int32, bool)`

GetEventsMatchedOk returns a tuple with the EventsMatched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsMatched

`func (o *SearchStatsInnerStatistics) SetEventsMatched(v int32)`

SetEventsMatched sets EventsMatched field to given value.

### HasEventsMatched

`func (o *SearchStatsInnerStatistics) HasEventsMatched() bool`

HasEventsMatched returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


