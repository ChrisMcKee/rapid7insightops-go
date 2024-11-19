# StatisticsMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of calculation performed.  | [optional] 
**Key** | Pointer to **string** | The key that the calculation was performed on. | [optional] 
**Resolution** | Pointer to **int32** | The length of each element of the time series, in seconds. | [optional] 
**Timeseries** | Pointer to [**[]StatisticsMetricsTimeseriesInner**](StatisticsMetricsTimeseriesInner.md) |  | [optional] 
**Result** | Pointer to **float32** | The total of all the values in the timeseries. | [optional] 

## Methods

### NewStatisticsMetrics

`func NewStatisticsMetrics() *StatisticsMetrics`

NewStatisticsMetrics instantiates a new StatisticsMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatisticsMetricsWithDefaults

`func NewStatisticsMetricsWithDefaults() *StatisticsMetrics`

NewStatisticsMetricsWithDefaults instantiates a new StatisticsMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *StatisticsMetrics) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StatisticsMetrics) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StatisticsMetrics) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StatisticsMetrics) HasType() bool`

HasType returns a boolean if a field has been set.

### GetKey

`func (o *StatisticsMetrics) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *StatisticsMetrics) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *StatisticsMetrics) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *StatisticsMetrics) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetResolution

`func (o *StatisticsMetrics) GetResolution() int32`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *StatisticsMetrics) GetResolutionOk() (*int32, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *StatisticsMetrics) SetResolution(v int32)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *StatisticsMetrics) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetTimeseries

`func (o *StatisticsMetrics) GetTimeseries() []StatisticsMetricsTimeseriesInner`

GetTimeseries returns the Timeseries field if non-nil, zero value otherwise.

### GetTimeseriesOk

`func (o *StatisticsMetrics) GetTimeseriesOk() (*[]StatisticsMetricsTimeseriesInner, bool)`

GetTimeseriesOk returns a tuple with the Timeseries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeseries

`func (o *StatisticsMetrics) SetTimeseries(v []StatisticsMetricsTimeseriesInner)`

SetTimeseries sets Timeseries field to given value.

### HasTimeseries

`func (o *StatisticsMetrics) HasTimeseries() bool`

HasTimeseries returns a boolean if a field has been set.

### GetResult

`func (o *StatisticsMetrics) GetResult() float32`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *StatisticsMetrics) GetResultOk() (*float32, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *StatisticsMetrics) SetResult(v float32)`

SetResult sets Result field to given value.

### HasResult

`func (o *StatisticsMetrics) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


