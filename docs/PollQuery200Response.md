# PollQuery200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | **[]string** | The log keys of the logs which the query is run against. | 
**Leql** | [**Leql**](Leql.md) |  | 
**Progress** | **int32** | The completion of the query in percent (0 to 99). | 
**Id** | **string** | The query ID which can be used to poll the query on the **_/query/{id}** endpoint. | 
**Events** | **[]map[string]interface{}** | This array is always empty. | 
**Links** | [**[]QueryApiLinksInner**](QueryApiLinksInner.md) |  | 
**Partial** | [**Statistics**](Statistics.md) |  | 
**Statistics** | [**Statistics**](Statistics.md) |  | 
**SearchStats** | [**SearchStats**](SearchStats.md) |  | 

## Methods

### NewPollQuery200Response

`func NewPollQuery200Response(logs []string, leql Leql, progress int32, id string, events []map[string]interface{}, links []QueryApiLinksInner, partial Statistics, statistics Statistics, searchStats SearchStats, ) *PollQuery200Response`

NewPollQuery200Response instantiates a new PollQuery200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPollQuery200ResponseWithDefaults

`func NewPollQuery200ResponseWithDefaults() *PollQuery200Response`

NewPollQuery200ResponseWithDefaults instantiates a new PollQuery200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *PollQuery200Response) GetLogs() []string`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *PollQuery200Response) GetLogsOk() (*[]string, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *PollQuery200Response) SetLogs(v []string)`

SetLogs sets Logs field to given value.


### GetLeql

`func (o *PollQuery200Response) GetLeql() Leql`

GetLeql returns the Leql field if non-nil, zero value otherwise.

### GetLeqlOk

`func (o *PollQuery200Response) GetLeqlOk() (*Leql, bool)`

GetLeqlOk returns a tuple with the Leql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeql

`func (o *PollQuery200Response) SetLeql(v Leql)`

SetLeql sets Leql field to given value.


### GetProgress

`func (o *PollQuery200Response) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *PollQuery200Response) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *PollQuery200Response) SetProgress(v int32)`

SetProgress sets Progress field to given value.


### GetId

`func (o *PollQuery200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PollQuery200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PollQuery200Response) SetId(v string)`

SetId sets Id field to given value.


### GetEvents

`func (o *PollQuery200Response) GetEvents() []map[string]interface{}`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *PollQuery200Response) GetEventsOk() (*[]map[string]interface{}, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *PollQuery200Response) SetEvents(v []map[string]interface{})`

SetEvents sets Events field to given value.


### GetLinks

`func (o *PollQuery200Response) GetLinks() []QueryApiLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PollQuery200Response) GetLinksOk() (*[]QueryApiLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PollQuery200Response) SetLinks(v []QueryApiLinksInner)`

SetLinks sets Links field to given value.


### GetPartial

`func (o *PollQuery200Response) GetPartial() Statistics`

GetPartial returns the Partial field if non-nil, zero value otherwise.

### GetPartialOk

`func (o *PollQuery200Response) GetPartialOk() (*Statistics, bool)`

GetPartialOk returns a tuple with the Partial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartial

`func (o *PollQuery200Response) SetPartial(v Statistics)`

SetPartial sets Partial field to given value.


### GetStatistics

`func (o *PollQuery200Response) GetStatistics() Statistics`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *PollQuery200Response) GetStatisticsOk() (*Statistics, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *PollQuery200Response) SetStatistics(v Statistics)`

SetStatistics sets Statistics field to given value.


### GetSearchStats

`func (o *PollQuery200Response) GetSearchStats() SearchStats`

GetSearchStats returns the SearchStats field if non-nil, zero value otherwise.

### GetSearchStatsOk

`func (o *PollQuery200Response) GetSearchStatsOk() (*SearchStats, bool)`

GetSearchStatsOk returns a tuple with the SearchStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchStats

`func (o *PollQuery200Response) SetSearchStats(v SearchStats)`

SetSearchStats sets SearchStats field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


