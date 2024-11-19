# EventsContinuation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | **[]string** | The log keys of the logs which the query is run against. | 
**Leql** | [**Leql**](Leql.md) |  | 
**Progress** | **int32** | The completion of the query in percent (0 to 99). | 
**Id** | **string** | The query ID which can be used to poll the query on the **_/query/{id}** endpoint. | 
**Events** | [**[]EventQuery**](EventQuery.md) | The result of the query - the matching log entries (with metadata). It is empty for 202 responses.  | 
**Links** | [**[]QueryApiLinksInner**](QueryApiLinksInner.md) |  | 

## Methods

### NewEventsContinuation

`func NewEventsContinuation(logs []string, leql Leql, progress int32, id string, events []EventQuery, links []QueryApiLinksInner, ) *EventsContinuation`

NewEventsContinuation instantiates a new EventsContinuation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsContinuationWithDefaults

`func NewEventsContinuationWithDefaults() *EventsContinuation`

NewEventsContinuationWithDefaults instantiates a new EventsContinuation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *EventsContinuation) GetLogs() []string`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *EventsContinuation) GetLogsOk() (*[]string, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *EventsContinuation) SetLogs(v []string)`

SetLogs sets Logs field to given value.


### GetLeql

`func (o *EventsContinuation) GetLeql() Leql`

GetLeql returns the Leql field if non-nil, zero value otherwise.

### GetLeqlOk

`func (o *EventsContinuation) GetLeqlOk() (*Leql, bool)`

GetLeqlOk returns a tuple with the Leql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeql

`func (o *EventsContinuation) SetLeql(v Leql)`

SetLeql sets Leql field to given value.


### GetProgress

`func (o *EventsContinuation) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *EventsContinuation) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *EventsContinuation) SetProgress(v int32)`

SetProgress sets Progress field to given value.


### GetId

`func (o *EventsContinuation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventsContinuation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventsContinuation) SetId(v string)`

SetId sets Id field to given value.


### GetEvents

`func (o *EventsContinuation) GetEvents() []EventQuery`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *EventsContinuation) GetEventsOk() (*[]EventQuery, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *EventsContinuation) SetEvents(v []EventQuery)`

SetEvents sets Events field to given value.


### GetLinks

`func (o *EventsContinuation) GetLinks() []QueryApiLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EventsContinuation) GetLinksOk() (*[]QueryApiLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EventsContinuation) SetLinks(v []QueryApiLinksInner)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


