# MetricsInputMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to [**String**](String.md) | The name of the pre-computed query. | [optional] 
**Description** | Pointer to [**String**](String.md) | The description of the pre-computed query. | [optional] 
**Enabled** | Pointer to **interface{}** | This indicates whether the pre-computed query is enabled or not. | [optional] 
**Logs** | Pointer to [**MetricsInputMetricLogs**](MetricsInputMetricLogs.md) |  | [optional] 
**Leql** | Pointer to [**LeqlMetrics**](LeqlMetrics.md) |  | [optional] 
**Resolution** | Pointer to **int32** | The time window in seconds, that each PCQ datapoint corresponds to. | [optional] 
**Retention** | Pointer to **int32** | The length of time in seconds, that the PCQ will be stored for. | [optional] 

## Methods

### NewMetricsInputMetric

`func NewMetricsInputMetric() *MetricsInputMetric`

NewMetricsInputMetric instantiates a new MetricsInputMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsInputMetricWithDefaults

`func NewMetricsInputMetricWithDefaults() *MetricsInputMetric`

NewMetricsInputMetricWithDefaults instantiates a new MetricsInputMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MetricsInputMetric) GetName() String`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetricsInputMetric) GetNameOk() (*String, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetricsInputMetric) SetName(v String)`

SetName sets Name field to given value.

### HasName

`func (o *MetricsInputMetric) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *MetricsInputMetric) GetDescription() String`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetricsInputMetric) GetDescriptionOk() (*String, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetricsInputMetric) SetDescription(v String)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MetricsInputMetric) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *MetricsInputMetric) GetEnabled() interface{}`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MetricsInputMetric) GetEnabledOk() (*interface{}, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MetricsInputMetric) SetEnabled(v interface{})`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MetricsInputMetric) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *MetricsInputMetric) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *MetricsInputMetric) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetLogs

`func (o *MetricsInputMetric) GetLogs() MetricsInputMetricLogs`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *MetricsInputMetric) GetLogsOk() (*MetricsInputMetricLogs, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *MetricsInputMetric) SetLogs(v MetricsInputMetricLogs)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *MetricsInputMetric) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetLeql

`func (o *MetricsInputMetric) GetLeql() LeqlMetrics`

GetLeql returns the Leql field if non-nil, zero value otherwise.

### GetLeqlOk

`func (o *MetricsInputMetric) GetLeqlOk() (*LeqlMetrics, bool)`

GetLeqlOk returns a tuple with the Leql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeql

`func (o *MetricsInputMetric) SetLeql(v LeqlMetrics)`

SetLeql sets Leql field to given value.

### HasLeql

`func (o *MetricsInputMetric) HasLeql() bool`

HasLeql returns a boolean if a field has been set.

### GetResolution

`func (o *MetricsInputMetric) GetResolution() int32`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *MetricsInputMetric) GetResolutionOk() (*int32, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *MetricsInputMetric) SetResolution(v int32)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *MetricsInputMetric) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetRetention

`func (o *MetricsInputMetric) GetRetention() int32`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *MetricsInputMetric) GetRetentionOk() (*int32, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *MetricsInputMetric) SetRetention(v int32)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *MetricsInputMetric) HasRetention() bool`

HasRetention returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


