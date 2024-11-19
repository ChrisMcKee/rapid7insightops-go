# Stats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Average** | Pointer to **float32** |  | [optional] 
**Min** | Pointer to **float32** |  | [optional] 
**Max** | Pointer to **float32** |  | [optional] 
**Median** | Pointer to **float32** |  | [optional] 
**Percentile1To99** | Pointer to **float32** |  | [optional] 
**Count** | Pointer to **float32** |  | [optional] 
**Sum** | Pointer to **float32** |  | [optional] 
**Standarddeviation** | Pointer to **float32** |  | [optional] 
**Unique** | Pointer to **float32** |  | [optional] 
**Bytes** | Pointer to **float32** |  | [optional] 

## Methods

### NewStats

`func NewStats() *Stats`

NewStats instantiates a new Stats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatsWithDefaults

`func NewStatsWithDefaults() *Stats`

NewStatsWithDefaults instantiates a new Stats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverage

`func (o *Stats) GetAverage() float32`

GetAverage returns the Average field if non-nil, zero value otherwise.

### GetAverageOk

`func (o *Stats) GetAverageOk() (*float32, bool)`

GetAverageOk returns a tuple with the Average field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverage

`func (o *Stats) SetAverage(v float32)`

SetAverage sets Average field to given value.

### HasAverage

`func (o *Stats) HasAverage() bool`

HasAverage returns a boolean if a field has been set.

### GetMin

`func (o *Stats) GetMin() float32`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *Stats) GetMinOk() (*float32, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *Stats) SetMin(v float32)`

SetMin sets Min field to given value.

### HasMin

`func (o *Stats) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetMax

`func (o *Stats) GetMax() float32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *Stats) GetMaxOk() (*float32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *Stats) SetMax(v float32)`

SetMax sets Max field to given value.

### HasMax

`func (o *Stats) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMedian

`func (o *Stats) GetMedian() float32`

GetMedian returns the Median field if non-nil, zero value otherwise.

### GetMedianOk

`func (o *Stats) GetMedianOk() (*float32, bool)`

GetMedianOk returns a tuple with the Median field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedian

`func (o *Stats) SetMedian(v float32)`

SetMedian sets Median field to given value.

### HasMedian

`func (o *Stats) HasMedian() bool`

HasMedian returns a boolean if a field has been set.

### GetPercentile1To99

`func (o *Stats) GetPercentile1To99() float32`

GetPercentile1To99 returns the Percentile1To99 field if non-nil, zero value otherwise.

### GetPercentile1To99Ok

`func (o *Stats) GetPercentile1To99Ok() (*float32, bool)`

GetPercentile1To99Ok returns a tuple with the Percentile1To99 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentile1To99

`func (o *Stats) SetPercentile1To99(v float32)`

SetPercentile1To99 sets Percentile1To99 field to given value.

### HasPercentile1To99

`func (o *Stats) HasPercentile1To99() bool`

HasPercentile1To99 returns a boolean if a field has been set.

### GetCount

`func (o *Stats) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *Stats) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *Stats) SetCount(v float32)`

SetCount sets Count field to given value.

### HasCount

`func (o *Stats) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetSum

`func (o *Stats) GetSum() float32`

GetSum returns the Sum field if non-nil, zero value otherwise.

### GetSumOk

`func (o *Stats) GetSumOk() (*float32, bool)`

GetSumOk returns a tuple with the Sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSum

`func (o *Stats) SetSum(v float32)`

SetSum sets Sum field to given value.

### HasSum

`func (o *Stats) HasSum() bool`

HasSum returns a boolean if a field has been set.

### GetStandarddeviation

`func (o *Stats) GetStandarddeviation() float32`

GetStandarddeviation returns the Standarddeviation field if non-nil, zero value otherwise.

### GetStandarddeviationOk

`func (o *Stats) GetStandarddeviationOk() (*float32, bool)`

GetStandarddeviationOk returns a tuple with the Standarddeviation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandarddeviation

`func (o *Stats) SetStandarddeviation(v float32)`

SetStandarddeviation sets Standarddeviation field to given value.

### HasStandarddeviation

`func (o *Stats) HasStandarddeviation() bool`

HasStandarddeviation returns a boolean if a field has been set.

### GetUnique

`func (o *Stats) GetUnique() float32`

GetUnique returns the Unique field if non-nil, zero value otherwise.

### GetUniqueOk

`func (o *Stats) GetUniqueOk() (*float32, bool)`

GetUniqueOk returns a tuple with the Unique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnique

`func (o *Stats) SetUnique(v float32)`

SetUnique sets Unique field to given value.

### HasUnique

`func (o *Stats) HasUnique() bool`

HasUnique returns a boolean if a field has been set.

### GetBytes

`func (o *Stats) GetBytes() float32`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *Stats) GetBytesOk() (*float32, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *Stats) SetBytes(v float32)`

SetBytes sets Bytes field to given value.

### HasBytes

`func (o *Stats) HasBytes() bool`

HasBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


