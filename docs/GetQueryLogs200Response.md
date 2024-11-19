# GetQueryLogs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | **[]string** | The log keys of the logs which the query is run against. | 
**Leql** | [**Leql**](Leql.md) |  | 
**Events** | [**[]EventQuery**](EventQuery.md) | The result of the query - the matching log entries (with metadata). It is empty for 202 responses.  | 
**Links** | Pointer to [**[]QueryApiLinksInner**](QueryApiLinksInner.md) |  | [optional] 
**Statistics** | [**Statistics**](Statistics.md) |  | 
**SearchStats** | [**SearchStats**](SearchStats.md) |  | 

## Methods

### NewGetQueryLogs200Response

`func NewGetQueryLogs200Response(logs []string, leql Leql, events []EventQuery, statistics Statistics, searchStats SearchStats, ) *GetQueryLogs200Response`

NewGetQueryLogs200Response instantiates a new GetQueryLogs200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetQueryLogs200ResponseWithDefaults

`func NewGetQueryLogs200ResponseWithDefaults() *GetQueryLogs200Response`

NewGetQueryLogs200ResponseWithDefaults instantiates a new GetQueryLogs200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *GetQueryLogs200Response) GetLogs() []string`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *GetQueryLogs200Response) GetLogsOk() (*[]string, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *GetQueryLogs200Response) SetLogs(v []string)`

SetLogs sets Logs field to given value.


### GetLeql

`func (o *GetQueryLogs200Response) GetLeql() Leql`

GetLeql returns the Leql field if non-nil, zero value otherwise.

### GetLeqlOk

`func (o *GetQueryLogs200Response) GetLeqlOk() (*Leql, bool)`

GetLeqlOk returns a tuple with the Leql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeql

`func (o *GetQueryLogs200Response) SetLeql(v Leql)`

SetLeql sets Leql field to given value.


### GetEvents

`func (o *GetQueryLogs200Response) GetEvents() []EventQuery`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *GetQueryLogs200Response) GetEventsOk() (*[]EventQuery, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *GetQueryLogs200Response) SetEvents(v []EventQuery)`

SetEvents sets Events field to given value.


### GetLinks

`func (o *GetQueryLogs200Response) GetLinks() []QueryApiLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetQueryLogs200Response) GetLinksOk() (*[]QueryApiLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetQueryLogs200Response) SetLinks(v []QueryApiLinksInner)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetQueryLogs200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetStatistics

`func (o *GetQueryLogs200Response) GetStatistics() Statistics`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *GetQueryLogs200Response) GetStatisticsOk() (*Statistics, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *GetQueryLogs200Response) SetStatistics(v Statistics)`

SetStatistics sets Statistics field to given value.


### GetSearchStats

`func (o *GetQueryLogs200Response) GetSearchStats() SearchStats`

GetSearchStats returns the SearchStats field if non-nil, zero value otherwise.

### GetSearchStatsOk

`func (o *GetQueryLogs200Response) GetSearchStatsOk() (*SearchStats, bool)`

GetSearchStatsOk returns a tuple with the SearchStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchStats

`func (o *GetQueryLogs200Response) SetSearchStats(v SearchStats)`

SetSearchStats sets SearchStats field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


