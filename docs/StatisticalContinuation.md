# StatisticalContinuation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | **[]string** | The log keys of the logs which the query is run against. | 
**Leql** | [**Leql**](Leql.md) |  | 
**Progress** | **int32** | The completion of the query in percent (0 to 99). | 
**Id** | **string** | The query ID which can be used to poll the query on the **_/query/{id}** endpoint. | 
**Events** | **[]map[string]interface{}** | This array is always empty. | 
**Partial** | [**Statistics**](Statistics.md) |  | 
**Links** | [**[]QueryApiLinksInner**](QueryApiLinksInner.md) |  | 

## Methods

### NewStatisticalContinuation

`func NewStatisticalContinuation(logs []string, leql Leql, progress int32, id string, events []map[string]interface{}, partial Statistics, links []QueryApiLinksInner, ) *StatisticalContinuation`

NewStatisticalContinuation instantiates a new StatisticalContinuation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatisticalContinuationWithDefaults

`func NewStatisticalContinuationWithDefaults() *StatisticalContinuation`

NewStatisticalContinuationWithDefaults instantiates a new StatisticalContinuation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *StatisticalContinuation) GetLogs() []string`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *StatisticalContinuation) GetLogsOk() (*[]string, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *StatisticalContinuation) SetLogs(v []string)`

SetLogs sets Logs field to given value.


### GetLeql

`func (o *StatisticalContinuation) GetLeql() Leql`

GetLeql returns the Leql field if non-nil, zero value otherwise.

### GetLeqlOk

`func (o *StatisticalContinuation) GetLeqlOk() (*Leql, bool)`

GetLeqlOk returns a tuple with the Leql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeql

`func (o *StatisticalContinuation) SetLeql(v Leql)`

SetLeql sets Leql field to given value.


### GetProgress

`func (o *StatisticalContinuation) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *StatisticalContinuation) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *StatisticalContinuation) SetProgress(v int32)`

SetProgress sets Progress field to given value.


### GetId

`func (o *StatisticalContinuation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StatisticalContinuation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StatisticalContinuation) SetId(v string)`

SetId sets Id field to given value.


### GetEvents

`func (o *StatisticalContinuation) GetEvents() []map[string]interface{}`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *StatisticalContinuation) GetEventsOk() (*[]map[string]interface{}, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *StatisticalContinuation) SetEvents(v []map[string]interface{})`

SetEvents sets Events field to given value.


### GetPartial

`func (o *StatisticalContinuation) GetPartial() Statistics`

GetPartial returns the Partial field if non-nil, zero value otherwise.

### GetPartialOk

`func (o *StatisticalContinuation) GetPartialOk() (*Statistics, bool)`

GetPartialOk returns a tuple with the Partial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartial

`func (o *StatisticalContinuation) SetPartial(v Statistics)`

SetPartial sets Partial field to given value.


### GetLinks

`func (o *StatisticalContinuation) GetLinks() []QueryApiLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StatisticalContinuation) GetLinksOk() (*[]QueryApiLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StatisticalContinuation) SetLinks(v []QueryApiLinksInner)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


