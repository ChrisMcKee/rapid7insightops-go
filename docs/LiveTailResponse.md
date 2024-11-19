# LiveTailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | **[]string** | The log keys of the logs for the live tail feed. | 
**Leql** | [**LeqlLiveTail**](LeqlLiveTail.md) |  | 
**Events** | [**[]EventLiveTail**](EventLiveTail.md) | Always empty.  | 
**Links** | [**[]LinksInner**](LinksInner.md) |  | 

## Methods

### NewLiveTailResponse

`func NewLiveTailResponse(logs []string, leql LeqlLiveTail, events []EventLiveTail, links []LinksInner, ) *LiveTailResponse`

NewLiveTailResponse instantiates a new LiveTailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveTailResponseWithDefaults

`func NewLiveTailResponseWithDefaults() *LiveTailResponse`

NewLiveTailResponseWithDefaults instantiates a new LiveTailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *LiveTailResponse) GetLogs() []string`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *LiveTailResponse) GetLogsOk() (*[]string, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *LiveTailResponse) SetLogs(v []string)`

SetLogs sets Logs field to given value.


### GetLeql

`func (o *LiveTailResponse) GetLeql() LeqlLiveTail`

GetLeql returns the Leql field if non-nil, zero value otherwise.

### GetLeqlOk

`func (o *LiveTailResponse) GetLeqlOk() (*LeqlLiveTail, bool)`

GetLeqlOk returns a tuple with the Leql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeql

`func (o *LiveTailResponse) SetLeql(v LeqlLiveTail)`

SetLeql sets Leql field to given value.


### GetEvents

`func (o *LiveTailResponse) GetEvents() []EventLiveTail`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *LiveTailResponse) GetEventsOk() (*[]EventLiveTail, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *LiveTailResponse) SetEvents(v []EventLiveTail)`

SetEvents sets Events field to given value.


### GetLinks

`func (o *LiveTailResponse) GetLinks() []LinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *LiveTailResponse) GetLinksOk() (*[]LinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *LiveTailResponse) SetLinks(v []LinksInner)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


