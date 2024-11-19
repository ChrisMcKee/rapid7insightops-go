# MetricsInputLogsetsByNameMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to [**String**](String.md) | The name of the pre-computed query. | [optional] 
**Description** | Pointer to [**String**](String.md) | The description of the pre-computed query. | [optional] 
**Enabled** | Pointer to **interface{}** | This indicates whether the pre-computed query is enabled or not. | [optional] 
**Logsets** | Pointer to [**MetricsInputLogsetsByNameMetricLogsets**](MetricsInputLogsetsByNameMetricLogsets.md) |  | [optional] 
**Leql** | Pointer to [**LeqlMetrics**](LeqlMetrics.md) |  | [optional] 
**Resolution** | Pointer to **int32** | The time window in seconds, that each PCQ datapoint corresponds to. | [optional] 
**Retention** | Pointer to **int32** | The length of time in seconds, that the PCQ will be stored for. | [optional] 

## Methods

### NewMetricsInputLogsetsByNameMetric

`func NewMetricsInputLogsetsByNameMetric() *MetricsInputLogsetsByNameMetric`

NewMetricsInputLogsetsByNameMetric instantiates a new MetricsInputLogsetsByNameMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsInputLogsetsByNameMetricWithDefaults

`func NewMetricsInputLogsetsByNameMetricWithDefaults() *MetricsInputLogsetsByNameMetric`

NewMetricsInputLogsetsByNameMetricWithDefaults instantiates a new MetricsInputLogsetsByNameMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MetricsInputLogsetsByNameMetric) GetName() String`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetricsInputLogsetsByNameMetric) GetNameOk() (*String, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetricsInputLogsetsByNameMetric) SetName(v String)`

SetName sets Name field to given value.

### HasName

`func (o *MetricsInputLogsetsByNameMetric) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *MetricsInputLogsetsByNameMetric) GetDescription() String`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetricsInputLogsetsByNameMetric) GetDescriptionOk() (*String, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetricsInputLogsetsByNameMetric) SetDescription(v String)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MetricsInputLogsetsByNameMetric) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *MetricsInputLogsetsByNameMetric) GetEnabled() interface{}`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MetricsInputLogsetsByNameMetric) GetEnabledOk() (*interface{}, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MetricsInputLogsetsByNameMetric) SetEnabled(v interface{})`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MetricsInputLogsetsByNameMetric) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *MetricsInputLogsetsByNameMetric) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *MetricsInputLogsetsByNameMetric) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetLogsets

`func (o *MetricsInputLogsetsByNameMetric) GetLogsets() MetricsInputLogsetsByNameMetricLogsets`

GetLogsets returns the Logsets field if non-nil, zero value otherwise.

### GetLogsetsOk

`func (o *MetricsInputLogsetsByNameMetric) GetLogsetsOk() (*MetricsInputLogsetsByNameMetricLogsets, bool)`

GetLogsetsOk returns a tuple with the Logsets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsets

`func (o *MetricsInputLogsetsByNameMetric) SetLogsets(v MetricsInputLogsetsByNameMetricLogsets)`

SetLogsets sets Logsets field to given value.

### HasLogsets

`func (o *MetricsInputLogsetsByNameMetric) HasLogsets() bool`

HasLogsets returns a boolean if a field has been set.

### GetLeql

`func (o *MetricsInputLogsetsByNameMetric) GetLeql() LeqlMetrics`

GetLeql returns the Leql field if non-nil, zero value otherwise.

### GetLeqlOk

`func (o *MetricsInputLogsetsByNameMetric) GetLeqlOk() (*LeqlMetrics, bool)`

GetLeqlOk returns a tuple with the Leql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeql

`func (o *MetricsInputLogsetsByNameMetric) SetLeql(v LeqlMetrics)`

SetLeql sets Leql field to given value.

### HasLeql

`func (o *MetricsInputLogsetsByNameMetric) HasLeql() bool`

HasLeql returns a boolean if a field has been set.

### GetResolution

`func (o *MetricsInputLogsetsByNameMetric) GetResolution() int32`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *MetricsInputLogsetsByNameMetric) GetResolutionOk() (*int32, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *MetricsInputLogsetsByNameMetric) SetResolution(v int32)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *MetricsInputLogsetsByNameMetric) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetRetention

`func (o *MetricsInputLogsetsByNameMetric) GetRetention() int32`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *MetricsInputLogsetsByNameMetric) GetRetentionOk() (*int32, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *MetricsInputLogsetsByNameMetric) SetRetention(v int32)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *MetricsInputLogsetsByNameMetric) HasRetention() bool`

HasRetention returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


