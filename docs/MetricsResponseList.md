# MetricsResponseList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metrics** | Pointer to [**[]MetricsResponseSingle**](MetricsResponseSingle.md) |  | [optional] 

## Methods

### NewMetricsResponseList

`func NewMetricsResponseList() *MetricsResponseList`

NewMetricsResponseList instantiates a new MetricsResponseList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsResponseListWithDefaults

`func NewMetricsResponseListWithDefaults() *MetricsResponseList`

NewMetricsResponseListWithDefaults instantiates a new MetricsResponseList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetrics

`func (o *MetricsResponseList) GetMetrics() []MetricsResponseSingle`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *MetricsResponseList) GetMetricsOk() (*[]MetricsResponseSingle, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *MetricsResponseList) SetMetrics(v []MetricsResponseSingle)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *MetricsResponseList) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


