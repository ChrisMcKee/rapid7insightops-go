# PostQueryLogsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | Pointer to **[]string** | The log keys of the logs which the query is run against. | [optional] 
**Leql** | Pointer to [**Leql**](Leql.md) |  | [optional] 

## Methods

### NewPostQueryLogsRequest

`func NewPostQueryLogsRequest() *PostQueryLogsRequest`

NewPostQueryLogsRequest instantiates a new PostQueryLogsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostQueryLogsRequestWithDefaults

`func NewPostQueryLogsRequestWithDefaults() *PostQueryLogsRequest`

NewPostQueryLogsRequestWithDefaults instantiates a new PostQueryLogsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *PostQueryLogsRequest) GetLogs() []string`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *PostQueryLogsRequest) GetLogsOk() (*[]string, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *PostQueryLogsRequest) SetLogs(v []string)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *PostQueryLogsRequest) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetLeql

`func (o *PostQueryLogsRequest) GetLeql() Leql`

GetLeql returns the Leql field if non-nil, zero value otherwise.

### GetLeqlOk

`func (o *PostQueryLogsRequest) GetLeqlOk() (*Leql, bool)`

GetLeqlOk returns a tuple with the Leql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeql

`func (o *PostQueryLogsRequest) SetLeql(v Leql)`

SetLeql sets Leql field to given value.

### HasLeql

`func (o *PostQueryLogsRequest) HasLeql() bool`

HasLeql returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


