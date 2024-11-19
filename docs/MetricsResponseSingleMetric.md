# MetricsResponseSingleMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The UUID of the pre-computed query. | [optional] 
**Name** | Pointer to [**String**](String.md) | The name of the pre-computed query. | [optional] 
**Description** | Pointer to [**String**](String.md) | The description of the pre-computed query. | [optional] 
**Enabled** | Pointer to **interface{}** | This indicates whether the pre-computed query is enabled or not. | [optional] 
**Logs** | Pointer to [**[]LogInfoResponse**](LogInfoResponse.md) | The information of the logs used in the pre-computed query. | [optional] 
**Logsets** | Pointer to [**[]LogsetInfo**](LogsetInfo.md) | The information on the log sets used in the pre-computed query. | [optional] 
**Leql** | Pointer to [**LeqlMetricsEndpoints**](LeqlMetricsEndpoints.md) |  | [optional] 
**ResolutionSecs** | Pointer to **int32** | The time window in seconds, that each PCQ datapoint corresponds to. | [optional] 
**RetentionSecs** | Pointer to **int32** | The length of time in seconds, that the PCQ will be stored for. | [optional] 
**EnabledSince** | Pointer to **string** | The time the PCQ was enabled. | [optional] 
**Created** | Pointer to **string** | The time the PCQ was created. | [optional] 

## Methods

### NewMetricsResponseSingleMetric

`func NewMetricsResponseSingleMetric() *MetricsResponseSingleMetric`

NewMetricsResponseSingleMetric instantiates a new MetricsResponseSingleMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsResponseSingleMetricWithDefaults

`func NewMetricsResponseSingleMetricWithDefaults() *MetricsResponseSingleMetric`

NewMetricsResponseSingleMetricWithDefaults instantiates a new MetricsResponseSingleMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MetricsResponseSingleMetric) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetricsResponseSingleMetric) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetricsResponseSingleMetric) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MetricsResponseSingleMetric) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MetricsResponseSingleMetric) GetName() String`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetricsResponseSingleMetric) GetNameOk() (*String, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetricsResponseSingleMetric) SetName(v String)`

SetName sets Name field to given value.

### HasName

`func (o *MetricsResponseSingleMetric) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *MetricsResponseSingleMetric) GetDescription() String`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetricsResponseSingleMetric) GetDescriptionOk() (*String, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetricsResponseSingleMetric) SetDescription(v String)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MetricsResponseSingleMetric) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *MetricsResponseSingleMetric) GetEnabled() interface{}`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MetricsResponseSingleMetric) GetEnabledOk() (*interface{}, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MetricsResponseSingleMetric) SetEnabled(v interface{})`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MetricsResponseSingleMetric) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *MetricsResponseSingleMetric) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *MetricsResponseSingleMetric) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetLogs

`func (o *MetricsResponseSingleMetric) GetLogs() []LogInfoResponse`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *MetricsResponseSingleMetric) GetLogsOk() (*[]LogInfoResponse, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *MetricsResponseSingleMetric) SetLogs(v []LogInfoResponse)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *MetricsResponseSingleMetric) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetLogsets

`func (o *MetricsResponseSingleMetric) GetLogsets() []LogsetInfo`

GetLogsets returns the Logsets field if non-nil, zero value otherwise.

### GetLogsetsOk

`func (o *MetricsResponseSingleMetric) GetLogsetsOk() (*[]LogsetInfo, bool)`

GetLogsetsOk returns a tuple with the Logsets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsets

`func (o *MetricsResponseSingleMetric) SetLogsets(v []LogsetInfo)`

SetLogsets sets Logsets field to given value.

### HasLogsets

`func (o *MetricsResponseSingleMetric) HasLogsets() bool`

HasLogsets returns a boolean if a field has been set.

### GetLeql

`func (o *MetricsResponseSingleMetric) GetLeql() LeqlMetricsEndpoints`

GetLeql returns the Leql field if non-nil, zero value otherwise.

### GetLeqlOk

`func (o *MetricsResponseSingleMetric) GetLeqlOk() (*LeqlMetricsEndpoints, bool)`

GetLeqlOk returns a tuple with the Leql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeql

`func (o *MetricsResponseSingleMetric) SetLeql(v LeqlMetricsEndpoints)`

SetLeql sets Leql field to given value.

### HasLeql

`func (o *MetricsResponseSingleMetric) HasLeql() bool`

HasLeql returns a boolean if a field has been set.

### GetResolutionSecs

`func (o *MetricsResponseSingleMetric) GetResolutionSecs() int32`

GetResolutionSecs returns the ResolutionSecs field if non-nil, zero value otherwise.

### GetResolutionSecsOk

`func (o *MetricsResponseSingleMetric) GetResolutionSecsOk() (*int32, bool)`

GetResolutionSecsOk returns a tuple with the ResolutionSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionSecs

`func (o *MetricsResponseSingleMetric) SetResolutionSecs(v int32)`

SetResolutionSecs sets ResolutionSecs field to given value.

### HasResolutionSecs

`func (o *MetricsResponseSingleMetric) HasResolutionSecs() bool`

HasResolutionSecs returns a boolean if a field has been set.

### GetRetentionSecs

`func (o *MetricsResponseSingleMetric) GetRetentionSecs() int32`

GetRetentionSecs returns the RetentionSecs field if non-nil, zero value otherwise.

### GetRetentionSecsOk

`func (o *MetricsResponseSingleMetric) GetRetentionSecsOk() (*int32, bool)`

GetRetentionSecsOk returns a tuple with the RetentionSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionSecs

`func (o *MetricsResponseSingleMetric) SetRetentionSecs(v int32)`

SetRetentionSecs sets RetentionSecs field to given value.

### HasRetentionSecs

`func (o *MetricsResponseSingleMetric) HasRetentionSecs() bool`

HasRetentionSecs returns a boolean if a field has been set.

### GetEnabledSince

`func (o *MetricsResponseSingleMetric) GetEnabledSince() string`

GetEnabledSince returns the EnabledSince field if non-nil, zero value otherwise.

### GetEnabledSinceOk

`func (o *MetricsResponseSingleMetric) GetEnabledSinceOk() (*string, bool)`

GetEnabledSinceOk returns a tuple with the EnabledSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledSince

`func (o *MetricsResponseSingleMetric) SetEnabledSince(v string)`

SetEnabledSince sets EnabledSince field to given value.

### HasEnabledSince

`func (o *MetricsResponseSingleMetric) HasEnabledSince() bool`

HasEnabledSince returns a boolean if a field has been set.

### GetCreated

`func (o *MetricsResponseSingleMetric) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *MetricsResponseSingleMetric) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *MetricsResponseSingleMetric) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *MetricsResponseSingleMetric) HasCreated() bool`

HasCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


