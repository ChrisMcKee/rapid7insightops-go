# LogsetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logsets** | Pointer to [**[]LogsetInfo**](LogsetInfo.md) | A list of the log sets.  | [optional] 

## Methods

### NewLogsetResponse

`func NewLogsetResponse() *LogsetResponse`

NewLogsetResponse instantiates a new LogsetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsetResponseWithDefaults

`func NewLogsetResponseWithDefaults() *LogsetResponse`

NewLogsetResponseWithDefaults instantiates a new LogsetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogsets

`func (o *LogsetResponse) GetLogsets() []LogsetInfo`

GetLogsets returns the Logsets field if non-nil, zero value otherwise.

### GetLogsetsOk

`func (o *LogsetResponse) GetLogsetsOk() (*[]LogsetInfo, bool)`

GetLogsetsOk returns a tuple with the Logsets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsets

`func (o *LogsetResponse) SetLogsets(v []LogsetInfo)`

SetLogsets sets Logsets field to given value.

### HasLogsets

`func (o *LogsetResponse) HasLogsets() bool`

HasLogsets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


