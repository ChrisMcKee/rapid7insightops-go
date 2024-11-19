# StatisticalResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | **[]string** | The log keys of the logs which the query is run against. | 
**Leql** | [**Leql**](Leql.md) |  | 
**Statistics** | [**Statistics**](Statistics.md) |  | 
**SearchStats** | [**SearchStats**](SearchStats.md) |  | 

## Methods

### NewStatisticalResponse

`func NewStatisticalResponse(logs []string, leql Leql, statistics Statistics, searchStats SearchStats, ) *StatisticalResponse`

NewStatisticalResponse instantiates a new StatisticalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatisticalResponseWithDefaults

`func NewStatisticalResponseWithDefaults() *StatisticalResponse`

NewStatisticalResponseWithDefaults instantiates a new StatisticalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *StatisticalResponse) GetLogs() []string`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *StatisticalResponse) GetLogsOk() (*[]string, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *StatisticalResponse) SetLogs(v []string)`

SetLogs sets Logs field to given value.


### GetLeql

`func (o *StatisticalResponse) GetLeql() Leql`

GetLeql returns the Leql field if non-nil, zero value otherwise.

### GetLeqlOk

`func (o *StatisticalResponse) GetLeqlOk() (*Leql, bool)`

GetLeqlOk returns a tuple with the Leql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeql

`func (o *StatisticalResponse) SetLeql(v Leql)`

SetLeql sets Leql field to given value.


### GetStatistics

`func (o *StatisticalResponse) GetStatistics() Statistics`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *StatisticalResponse) GetStatisticsOk() (*Statistics, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *StatisticalResponse) SetStatistics(v Statistics)`

SetStatistics sets Statistics field to given value.


### GetSearchStats

`func (o *StatisticalResponse) GetSearchStats() SearchStats`

GetSearchStats returns the SearchStats field if non-nil, zero value otherwise.

### GetSearchStatsOk

`func (o *StatisticalResponse) GetSearchStatsOk() (*SearchStats, bool)`

GetSearchStatsOk returns a tuple with the SearchStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchStats

`func (o *StatisticalResponse) SetSearchStats(v SearchStats)`

SetSearchStats sets SearchStats field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


