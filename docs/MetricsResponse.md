# MetricsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The UUID of the pre-computed query. | [optional] 
**Logs** | Pointer to [**[]LogInfoResponse**](LogInfoResponse.md) | The information of the logs used in the pre-computed query. | [optional] 
**Logsets** | Pointer to [**[]LogsetInfo**](LogsetInfo.md) | The information on the log sets used in the pre-computed query. | [optional] 
**Leql** | Pointer to [**LeqlMetrics**](LeqlMetrics.md) |  | [optional] 
**Statistics** | Pointer to [**StatisticsMetrics**](StatisticsMetrics.md) |  | [optional] 

## Methods

### NewMetricsResponse

`func NewMetricsResponse() *MetricsResponse`

NewMetricsResponse instantiates a new MetricsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsResponseWithDefaults

`func NewMetricsResponseWithDefaults() *MetricsResponse`

NewMetricsResponseWithDefaults instantiates a new MetricsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MetricsResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetricsResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetricsResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MetricsResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLogs

`func (o *MetricsResponse) GetLogs() []LogInfoResponse`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *MetricsResponse) GetLogsOk() (*[]LogInfoResponse, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *MetricsResponse) SetLogs(v []LogInfoResponse)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *MetricsResponse) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetLogsets

`func (o *MetricsResponse) GetLogsets() []LogsetInfo`

GetLogsets returns the Logsets field if non-nil, zero value otherwise.

### GetLogsetsOk

`func (o *MetricsResponse) GetLogsetsOk() (*[]LogsetInfo, bool)`

GetLogsetsOk returns a tuple with the Logsets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsets

`func (o *MetricsResponse) SetLogsets(v []LogsetInfo)`

SetLogsets sets Logsets field to given value.

### HasLogsets

`func (o *MetricsResponse) HasLogsets() bool`

HasLogsets returns a boolean if a field has been set.

### GetLeql

`func (o *MetricsResponse) GetLeql() LeqlMetrics`

GetLeql returns the Leql field if non-nil, zero value otherwise.

### GetLeqlOk

`func (o *MetricsResponse) GetLeqlOk() (*LeqlMetrics, bool)`

GetLeqlOk returns a tuple with the Leql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeql

`func (o *MetricsResponse) SetLeql(v LeqlMetrics)`

SetLeql sets Leql field to given value.

### HasLeql

`func (o *MetricsResponse) HasLeql() bool`

HasLeql returns a boolean if a field has been set.

### GetStatistics

`func (o *MetricsResponse) GetStatistics() StatisticsMetrics`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *MetricsResponse) GetStatisticsOk() (*StatisticsMetrics, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *MetricsResponse) SetStatistics(v StatisticsMetrics)`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *MetricsResponse) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


