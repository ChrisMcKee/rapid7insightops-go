# SearchStatsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryId** | Pointer to **string** | The unique ID of the query. | [optional] 
**Date** | Pointer to **int32** | The date when the query was started, as a UNIX timestamp in milliseconds. | [optional] 
**Source** | Pointer to **string** | May contain some information on the source of the query. For example, if the value is &#x60;\&quot;dashboard\&quot;&#x60;, then the query was launched by a Log Search dashboard rather than explicitly by a user through the Log Search UI. | [optional] 
**Leql** | Pointer to [**SearchStatsInnerLeql**](SearchStatsInnerLeql.md) |  | [optional] 
**Statistics** | Pointer to [**SearchStatsInnerStatistics**](SearchStatsInnerStatistics.md) |  | [optional] 

## Methods

### NewSearchStatsInner

`func NewSearchStatsInner() *SearchStatsInner`

NewSearchStatsInner instantiates a new SearchStatsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchStatsInnerWithDefaults

`func NewSearchStatsInnerWithDefaults() *SearchStatsInner`

NewSearchStatsInnerWithDefaults instantiates a new SearchStatsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryId

`func (o *SearchStatsInner) GetQueryId() string`

GetQueryId returns the QueryId field if non-nil, zero value otherwise.

### GetQueryIdOk

`func (o *SearchStatsInner) GetQueryIdOk() (*string, bool)`

GetQueryIdOk returns a tuple with the QueryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryId

`func (o *SearchStatsInner) SetQueryId(v string)`

SetQueryId sets QueryId field to given value.

### HasQueryId

`func (o *SearchStatsInner) HasQueryId() bool`

HasQueryId returns a boolean if a field has been set.

### GetDate

`func (o *SearchStatsInner) GetDate() int32`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *SearchStatsInner) GetDateOk() (*int32, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *SearchStatsInner) SetDate(v int32)`

SetDate sets Date field to given value.

### HasDate

`func (o *SearchStatsInner) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetSource

`func (o *SearchStatsInner) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SearchStatsInner) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SearchStatsInner) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *SearchStatsInner) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetLeql

`func (o *SearchStatsInner) GetLeql() SearchStatsInnerLeql`

GetLeql returns the Leql field if non-nil, zero value otherwise.

### GetLeqlOk

`func (o *SearchStatsInner) GetLeqlOk() (*SearchStatsInnerLeql, bool)`

GetLeqlOk returns a tuple with the Leql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeql

`func (o *SearchStatsInner) SetLeql(v SearchStatsInnerLeql)`

SetLeql sets Leql field to given value.

### HasLeql

`func (o *SearchStatsInner) HasLeql() bool`

HasLeql returns a boolean if a field has been set.

### GetStatistics

`func (o *SearchStatsInner) GetStatistics() SearchStatsInnerStatistics`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *SearchStatsInner) GetStatisticsOk() (*SearchStatsInnerStatistics, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *SearchStatsInner) SetStatistics(v SearchStatsInnerStatistics)`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *SearchStatsInner) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


