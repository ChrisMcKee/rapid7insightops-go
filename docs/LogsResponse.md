# LogsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | [**[]LogInfoResponse**](LogInfoResponse.md) |  | 

## Methods

### NewLogsResponse

`func NewLogsResponse(logs []LogInfoResponse, ) *LogsResponse`

NewLogsResponse instantiates a new LogsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsResponseWithDefaults

`func NewLogsResponseWithDefaults() *LogsResponse`

NewLogsResponseWithDefaults instantiates a new LogsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *LogsResponse) GetLogs() []LogInfoResponse`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *LogsResponse) GetLogsOk() (*[]LogInfoResponse, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *LogsResponse) SetLogs(v []LogInfoResponse)`

SetLogs sets Logs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


