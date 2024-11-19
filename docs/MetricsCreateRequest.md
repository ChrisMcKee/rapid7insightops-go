# MetricsCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | Pointer to [**MetricsInputLogsetsByNameMetric**](MetricsInputLogsetsByNameMetric.md) |  | [optional] 

## Methods

### NewMetricsCreateRequest

`func NewMetricsCreateRequest() *MetricsCreateRequest`

NewMetricsCreateRequest instantiates a new MetricsCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsCreateRequestWithDefaults

`func NewMetricsCreateRequestWithDefaults() *MetricsCreateRequest`

NewMetricsCreateRequestWithDefaults instantiates a new MetricsCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *MetricsCreateRequest) GetMetric() MetricsInputLogsetsByNameMetric`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *MetricsCreateRequest) GetMetricOk() (*MetricsInputLogsetsByNameMetric, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *MetricsCreateRequest) SetMetric(v MetricsInputLogsetsByNameMetric)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *MetricsCreateRequest) HasMetric() bool`

HasMetric returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


