# SearchStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DurationMs** | Pointer to **int32** | The time taken to complete the query, in milliseconds. | [optional] 
**BytesAll** | Pointer to **int32** | The total number of bytes. | [optional] 
**BytesChecked** | Pointer to **int32** | The number of bytes checked. | [optional] 
**EventsAll** | Pointer to **int32** | The total number of events. | [optional] 
**EventsChecked** | Pointer to **int32** | The number of events checked. | [optional] 
**EventsMatched** | Pointer to **int32** | The number of events matched. | [optional] 
**IndexFactor** | Pointer to **float32** | This shows if the query was able to use our indexing technology to help reduce the query duration. * Value of 0: Indexing could not be used. * Value of 1: Full indexing was achieved to reduce the query duration.  | [optional] 

## Methods

### NewSearchStats

`func NewSearchStats() *SearchStats`

NewSearchStats instantiates a new SearchStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchStatsWithDefaults

`func NewSearchStatsWithDefaults() *SearchStats`

NewSearchStatsWithDefaults instantiates a new SearchStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDurationMs

`func (o *SearchStats) GetDurationMs() int32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *SearchStats) GetDurationMsOk() (*int32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *SearchStats) SetDurationMs(v int32)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *SearchStats) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.

### GetBytesAll

`func (o *SearchStats) GetBytesAll() int32`

GetBytesAll returns the BytesAll field if non-nil, zero value otherwise.

### GetBytesAllOk

`func (o *SearchStats) GetBytesAllOk() (*int32, bool)`

GetBytesAllOk returns a tuple with the BytesAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesAll

`func (o *SearchStats) SetBytesAll(v int32)`

SetBytesAll sets BytesAll field to given value.

### HasBytesAll

`func (o *SearchStats) HasBytesAll() bool`

HasBytesAll returns a boolean if a field has been set.

### GetBytesChecked

`func (o *SearchStats) GetBytesChecked() int32`

GetBytesChecked returns the BytesChecked field if non-nil, zero value otherwise.

### GetBytesCheckedOk

`func (o *SearchStats) GetBytesCheckedOk() (*int32, bool)`

GetBytesCheckedOk returns a tuple with the BytesChecked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesChecked

`func (o *SearchStats) SetBytesChecked(v int32)`

SetBytesChecked sets BytesChecked field to given value.

### HasBytesChecked

`func (o *SearchStats) HasBytesChecked() bool`

HasBytesChecked returns a boolean if a field has been set.

### GetEventsAll

`func (o *SearchStats) GetEventsAll() int32`

GetEventsAll returns the EventsAll field if non-nil, zero value otherwise.

### GetEventsAllOk

`func (o *SearchStats) GetEventsAllOk() (*int32, bool)`

GetEventsAllOk returns a tuple with the EventsAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsAll

`func (o *SearchStats) SetEventsAll(v int32)`

SetEventsAll sets EventsAll field to given value.

### HasEventsAll

`func (o *SearchStats) HasEventsAll() bool`

HasEventsAll returns a boolean if a field has been set.

### GetEventsChecked

`func (o *SearchStats) GetEventsChecked() int32`

GetEventsChecked returns the EventsChecked field if non-nil, zero value otherwise.

### GetEventsCheckedOk

`func (o *SearchStats) GetEventsCheckedOk() (*int32, bool)`

GetEventsCheckedOk returns a tuple with the EventsChecked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsChecked

`func (o *SearchStats) SetEventsChecked(v int32)`

SetEventsChecked sets EventsChecked field to given value.

### HasEventsChecked

`func (o *SearchStats) HasEventsChecked() bool`

HasEventsChecked returns a boolean if a field has been set.

### GetEventsMatched

`func (o *SearchStats) GetEventsMatched() int32`

GetEventsMatched returns the EventsMatched field if non-nil, zero value otherwise.

### GetEventsMatchedOk

`func (o *SearchStats) GetEventsMatchedOk() (*int32, bool)`

GetEventsMatchedOk returns a tuple with the EventsMatched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsMatched

`func (o *SearchStats) SetEventsMatched(v int32)`

SetEventsMatched sets EventsMatched field to given value.

### HasEventsMatched

`func (o *SearchStats) HasEventsMatched() bool`

HasEventsMatched returns a boolean if a field has been set.

### GetIndexFactor

`func (o *SearchStats) GetIndexFactor() float32`

GetIndexFactor returns the IndexFactor field if non-nil, zero value otherwise.

### GetIndexFactorOk

`func (o *SearchStats) GetIndexFactorOk() (*float32, bool)`

GetIndexFactorOk returns a tuple with the IndexFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexFactor

`func (o *SearchStats) SetIndexFactor(v float32)`

SetIndexFactor sets IndexFactor field to given value.

### HasIndexFactor

`func (o *SearchStats) HasIndexFactor() bool`

HasIndexFactor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


