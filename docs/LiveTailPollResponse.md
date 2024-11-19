# LiveTailPollResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | **[]string** | The log keys of the logs for the live tail feed. | 
**Leql** | [**LeqlLiveTail**](LeqlLiveTail.md) |  | 
**Events** | [**[]EventLiveTail**](EventLiveTail.md) | The log entries (with metadata) ingested which match the query.  | 
**Links** | [**[]LinksInner**](LinksInner.md) |  | 

## Methods

### NewLiveTailPollResponse

`func NewLiveTailPollResponse(logs []string, leql LeqlLiveTail, events []EventLiveTail, links []LinksInner, ) *LiveTailPollResponse`

NewLiveTailPollResponse instantiates a new LiveTailPollResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveTailPollResponseWithDefaults

`func NewLiveTailPollResponseWithDefaults() *LiveTailPollResponse`

NewLiveTailPollResponseWithDefaults instantiates a new LiveTailPollResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *LiveTailPollResponse) GetLogs() []string`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *LiveTailPollResponse) GetLogsOk() (*[]string, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *LiveTailPollResponse) SetLogs(v []string)`

SetLogs sets Logs field to given value.


### GetLeql

`func (o *LiveTailPollResponse) GetLeql() LeqlLiveTail`

GetLeql returns the Leql field if non-nil, zero value otherwise.

### GetLeqlOk

`func (o *LiveTailPollResponse) GetLeqlOk() (*LeqlLiveTail, bool)`

GetLeqlOk returns a tuple with the Leql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeql

`func (o *LiveTailPollResponse) SetLeql(v LeqlLiveTail)`

SetLeql sets Leql field to given value.


### GetEvents

`func (o *LiveTailPollResponse) GetEvents() []EventLiveTail`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *LiveTailPollResponse) GetEventsOk() (*[]EventLiveTail, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *LiveTailPollResponse) SetEvents(v []EventLiveTail)`

SetEvents sets Events field to given value.


### GetLinks

`func (o *LiveTailPollResponse) GetLinks() []LinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *LiveTailPollResponse) GetLinksOk() (*[]LinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *LiveTailPollResponse) SetLinks(v []LinksInner)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


