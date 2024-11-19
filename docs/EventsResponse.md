# EventsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | **[]string** | The log keys of the logs which the query is run against. | 
**Leql** | [**Leql**](Leql.md) |  | 
**Events** | [**[]EventQuery**](EventQuery.md) | The result of the query - the matching log entries (with metadata). It is empty for 202 responses.  | 
**Links** | Pointer to [**[]QueryApiLinksInner**](QueryApiLinksInner.md) |  | [optional] 

## Methods

### NewEventsResponse

`func NewEventsResponse(logs []string, leql Leql, events []EventQuery, ) *EventsResponse`

NewEventsResponse instantiates a new EventsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsResponseWithDefaults

`func NewEventsResponseWithDefaults() *EventsResponse`

NewEventsResponseWithDefaults instantiates a new EventsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *EventsResponse) GetLogs() []string`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *EventsResponse) GetLogsOk() (*[]string, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *EventsResponse) SetLogs(v []string)`

SetLogs sets Logs field to given value.


### GetLeql

`func (o *EventsResponse) GetLeql() Leql`

GetLeql returns the Leql field if non-nil, zero value otherwise.

### GetLeqlOk

`func (o *EventsResponse) GetLeqlOk() (*Leql, bool)`

GetLeqlOk returns a tuple with the Leql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeql

`func (o *EventsResponse) SetLeql(v Leql)`

SetLeql sets Leql field to given value.


### GetEvents

`func (o *EventsResponse) GetEvents() []EventQuery`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *EventsResponse) GetEventsOk() (*[]EventQuery, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *EventsResponse) SetEvents(v []EventQuery)`

SetEvents sets Events field to given value.


### GetLinks

`func (o *EventsResponse) GetLinks() []QueryApiLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EventsResponse) GetLinksOk() (*[]QueryApiLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EventsResponse) SetLinks(v []QueryApiLinksInner)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EventsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


