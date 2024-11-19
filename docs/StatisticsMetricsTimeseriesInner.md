# StatisticsMetricsTimeseriesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to **int32** | The end of the time range that the pre-computed query is applied to. The time is displayed as a UNIX timestamp in milliseconds. | [optional] 
**Value** | Pointer to **float32** | The pre-computed query value for this slice of the time range. | [optional] 

## Methods

### NewStatisticsMetricsTimeseriesInner

`func NewStatisticsMetricsTimeseriesInner() *StatisticsMetricsTimeseriesInner`

NewStatisticsMetricsTimeseriesInner instantiates a new StatisticsMetricsTimeseriesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatisticsMetricsTimeseriesInnerWithDefaults

`func NewStatisticsMetricsTimeseriesInnerWithDefaults() *StatisticsMetricsTimeseriesInner`

NewStatisticsMetricsTimeseriesInnerWithDefaults instantiates a new StatisticsMetricsTimeseriesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *StatisticsMetricsTimeseriesInner) GetTime() int32`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *StatisticsMetricsTimeseriesInner) GetTimeOk() (*int32, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *StatisticsMetricsTimeseriesInner) SetTime(v int32)`

SetTime sets Time field to given value.

### HasTime

`func (o *StatisticsMetricsTimeseriesInner) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetValue

`func (o *StatisticsMetricsTimeseriesInner) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *StatisticsMetricsTimeseriesInner) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *StatisticsMetricsTimeseriesInner) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *StatisticsMetricsTimeseriesInner) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


