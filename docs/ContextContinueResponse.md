# ContextContinueResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | **[]string** | The log key of the log which contains the contextual log entries. | 
**Leql** | [**LeqlContext**](LeqlContext.md) |  | 
**Progress** | **int32** | The completion of the query in percent (0 to 99). | 
**Id** | **string** | The query ID which can be used to poll the query on the **_/query/{id}** endpoint. | 
**Events** | [**[]EventQuery**](EventQuery.md) | The contextual log entries for the given log entry. It is empty for a 202 response.  | 
**Links** | Pointer to [**[]ContextApiLinksInner**](ContextApiLinksInner.md) |  | [optional] 

## Methods

### NewContextContinueResponse

`func NewContextContinueResponse(logs []string, leql LeqlContext, progress int32, id string, events []EventQuery, ) *ContextContinueResponse`

NewContextContinueResponse instantiates a new ContextContinueResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextContinueResponseWithDefaults

`func NewContextContinueResponseWithDefaults() *ContextContinueResponse`

NewContextContinueResponseWithDefaults instantiates a new ContextContinueResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *ContextContinueResponse) GetLogs() []string`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *ContextContinueResponse) GetLogsOk() (*[]string, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *ContextContinueResponse) SetLogs(v []string)`

SetLogs sets Logs field to given value.


### GetLeql

`func (o *ContextContinueResponse) GetLeql() LeqlContext`

GetLeql returns the Leql field if non-nil, zero value otherwise.

### GetLeqlOk

`func (o *ContextContinueResponse) GetLeqlOk() (*LeqlContext, bool)`

GetLeqlOk returns a tuple with the Leql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeql

`func (o *ContextContinueResponse) SetLeql(v LeqlContext)`

SetLeql sets Leql field to given value.


### GetProgress

`func (o *ContextContinueResponse) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *ContextContinueResponse) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *ContextContinueResponse) SetProgress(v int32)`

SetProgress sets Progress field to given value.


### GetId

`func (o *ContextContinueResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContextContinueResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContextContinueResponse) SetId(v string)`

SetId sets Id field to given value.


### GetEvents

`func (o *ContextContinueResponse) GetEvents() []EventQuery`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ContextContinueResponse) GetEventsOk() (*[]EventQuery, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ContextContinueResponse) SetEvents(v []EventQuery)`

SetEvents sets Events field to given value.


### GetLinks

`func (o *ContextContinueResponse) GetLinks() []ContextApiLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ContextContinueResponse) GetLinksOk() (*[]ContextApiLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ContextContinueResponse) SetLinks(v []ContextApiLinksInner)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ContextContinueResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


