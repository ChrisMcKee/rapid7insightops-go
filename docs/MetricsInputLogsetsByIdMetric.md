# MetricsInputLogsetsByIdMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to [**String**](String.md) | The name of the pre-computed query. | [optional] 
**Description** | Pointer to [**String**](String.md) | The description of the pre-computed query. | [optional] 
**Enabled** | Pointer to **interface{}** | This indicates whether the pre-computed query is enabled or not. | [optional] 
**Logsets** | Pointer to [**MetricsInputLogsetsByIdMetricLogsets**](MetricsInputLogsetsByIdMetricLogsets.md) |  | [optional] 
**Leql** | Pointer to [**LeqlMetricsEndpoints**](LeqlMetricsEndpoints.md) |  | [optional] 
**Resolution** | Pointer to **int32** | The time window in seconds, that each PCQ datapoint corresponds to. | [optional] 
**Retention** | Pointer to **int32** | The length of time in seconds, that the PCQ will be stored for. | [optional] 

## Methods

### NewMetricsInputLogsetsByIdMetric

`func NewMetricsInputLogsetsByIdMetric() *MetricsInputLogsetsByIdMetric`

NewMetricsInputLogsetsByIdMetric instantiates a new MetricsInputLogsetsByIdMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsInputLogsetsByIdMetricWithDefaults

`func NewMetricsInputLogsetsByIdMetricWithDefaults() *MetricsInputLogsetsByIdMetric`

NewMetricsInputLogsetsByIdMetricWithDefaults instantiates a new MetricsInputLogsetsByIdMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MetricsInputLogsetsByIdMetric) GetName() String`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetricsInputLogsetsByIdMetric) GetNameOk() (*String, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetricsInputLogsetsByIdMetric) SetName(v String)`

SetName sets Name field to given value.

### HasName

`func (o *MetricsInputLogsetsByIdMetric) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *MetricsInputLogsetsByIdMetric) GetDescription() String`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetricsInputLogsetsByIdMetric) GetDescriptionOk() (*String, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetricsInputLogsetsByIdMetric) SetDescription(v String)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MetricsInputLogsetsByIdMetric) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *MetricsInputLogsetsByIdMetric) GetEnabled() interface{}`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MetricsInputLogsetsByIdMetric) GetEnabledOk() (*interface{}, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MetricsInputLogsetsByIdMetric) SetEnabled(v interface{})`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MetricsInputLogsetsByIdMetric) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *MetricsInputLogsetsByIdMetric) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *MetricsInputLogsetsByIdMetric) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetLogsets

`func (o *MetricsInputLogsetsByIdMetric) GetLogsets() MetricsInputLogsetsByIdMetricLogsets`

GetLogsets returns the Logsets field if non-nil, zero value otherwise.

### GetLogsetsOk

`func (o *MetricsInputLogsetsByIdMetric) GetLogsetsOk() (*MetricsInputLogsetsByIdMetricLogsets, bool)`

GetLogsetsOk returns a tuple with the Logsets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsets

`func (o *MetricsInputLogsetsByIdMetric) SetLogsets(v MetricsInputLogsetsByIdMetricLogsets)`

SetLogsets sets Logsets field to given value.

### HasLogsets

`func (o *MetricsInputLogsetsByIdMetric) HasLogsets() bool`

HasLogsets returns a boolean if a field has been set.

### GetLeql

`func (o *MetricsInputLogsetsByIdMetric) GetLeql() LeqlMetricsEndpoints`

GetLeql returns the Leql field if non-nil, zero value otherwise.

### GetLeqlOk

`func (o *MetricsInputLogsetsByIdMetric) GetLeqlOk() (*LeqlMetricsEndpoints, bool)`

GetLeqlOk returns a tuple with the Leql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeql

`func (o *MetricsInputLogsetsByIdMetric) SetLeql(v LeqlMetricsEndpoints)`

SetLeql sets Leql field to given value.

### HasLeql

`func (o *MetricsInputLogsetsByIdMetric) HasLeql() bool`

HasLeql returns a boolean if a field has been set.

### GetResolution

`func (o *MetricsInputLogsetsByIdMetric) GetResolution() int32`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *MetricsInputLogsetsByIdMetric) GetResolutionOk() (*int32, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *MetricsInputLogsetsByIdMetric) SetResolution(v int32)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *MetricsInputLogsetsByIdMetric) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetRetention

`func (o *MetricsInputLogsetsByIdMetric) GetRetention() int32`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *MetricsInputLogsetsByIdMetric) GetRetentionOk() (*int32, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *MetricsInputLogsetsByIdMetric) SetRetention(v int32)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *MetricsInputLogsetsByIdMetric) HasRetention() bool`

HasRetention returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


