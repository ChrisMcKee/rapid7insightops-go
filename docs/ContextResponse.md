# ContextResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | **[]string** | The log key of the log which contains the contextual log entries. | 
**Leql** | [**LeqlContext**](LeqlContext.md) |  | 
**Events** | [**[]EventQuery**](EventQuery.md) | The contextual log entries for the given log entry. It is empty for a 202 response.  | 
**Links** | Pointer to [**[]ContextApiLinksInner**](ContextApiLinksInner.md) |  | [optional] 

## Methods

### NewContextResponse

`func NewContextResponse(logs []string, leql LeqlContext, events []EventQuery, ) *ContextResponse`

NewContextResponse instantiates a new ContextResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextResponseWithDefaults

`func NewContextResponseWithDefaults() *ContextResponse`

NewContextResponseWithDefaults instantiates a new ContextResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *ContextResponse) GetLogs() []string`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *ContextResponse) GetLogsOk() (*[]string, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *ContextResponse) SetLogs(v []string)`

SetLogs sets Logs field to given value.


### GetLeql

`func (o *ContextResponse) GetLeql() LeqlContext`

GetLeql returns the Leql field if non-nil, zero value otherwise.

### GetLeqlOk

`func (o *ContextResponse) GetLeqlOk() (*LeqlContext, bool)`

GetLeqlOk returns a tuple with the Leql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeql

`func (o *ContextResponse) SetLeql(v LeqlContext)`

SetLeql sets Leql field to given value.


### GetEvents

`func (o *ContextResponse) GetEvents() []EventQuery`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ContextResponse) GetEventsOk() (*[]EventQuery, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ContextResponse) SetEvents(v []EventQuery)`

SetEvents sets Events field to given value.


### GetLinks

`func (o *ContextResponse) GetLinks() []ContextApiLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ContextResponse) GetLinksOk() (*[]ContextApiLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ContextResponse) SetLinks(v []ContextApiLinksInner)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ContextResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


