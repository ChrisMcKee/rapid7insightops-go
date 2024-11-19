# PostLiveLogsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | **[]string** | The log keys of the logs for the live tail feed. | 
**Leql** | Pointer to [**LeqlLiveTail**](LeqlLiveTail.md) |  | [optional] 

## Methods

### NewPostLiveLogsRequest

`func NewPostLiveLogsRequest(logs []string, ) *PostLiveLogsRequest`

NewPostLiveLogsRequest instantiates a new PostLiveLogsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostLiveLogsRequestWithDefaults

`func NewPostLiveLogsRequestWithDefaults() *PostLiveLogsRequest`

NewPostLiveLogsRequestWithDefaults instantiates a new PostLiveLogsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *PostLiveLogsRequest) GetLogs() []string`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *PostLiveLogsRequest) GetLogsOk() (*[]string, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *PostLiveLogsRequest) SetLogs(v []string)`

SetLogs sets Logs field to given value.


### GetLeql

`func (o *PostLiveLogsRequest) GetLeql() LeqlLiveTail`

GetLeql returns the Leql field if non-nil, zero value otherwise.

### GetLeqlOk

`func (o *PostLiveLogsRequest) GetLeqlOk() (*LeqlLiveTail, bool)`

GetLeqlOk returns a tuple with the Leql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeql

`func (o *PostLiveLogsRequest) SetLeql(v LeqlLiveTail)`

SetLeql sets Leql field to given value.

### HasLeql

`func (o *PostLiveLogsRequest) HasLeql() bool`

HasLeql returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


