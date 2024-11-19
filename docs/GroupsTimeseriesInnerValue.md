# GroupsTimeseriesInnerValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupTimeseries** | [**[]map[string]GroupsTimeseriesInnerValue**](map[string]GroupsTimeseriesInnerValue.md) | For &#39;groupby&#39; queries, holds the timeseries object for each group. | 
**Series** | [**[]Stats**](Stats.md) |  | 
**Totals** | [**Stats**](Stats.md) |  | 

## Methods

### NewGroupsTimeseriesInnerValue

`func NewGroupsTimeseriesInnerValue(groupTimeseries []map[string]GroupsTimeseriesInnerValue, series []Stats, totals Stats, ) *GroupsTimeseriesInnerValue`

NewGroupsTimeseriesInnerValue instantiates a new GroupsTimeseriesInnerValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupsTimeseriesInnerValueWithDefaults

`func NewGroupsTimeseriesInnerValueWithDefaults() *GroupsTimeseriesInnerValue`

NewGroupsTimeseriesInnerValueWithDefaults instantiates a new GroupsTimeseriesInnerValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupTimeseries

`func (o *GroupsTimeseriesInnerValue) GetGroupTimeseries() []map[string]GroupsTimeseriesInnerValue`

GetGroupTimeseries returns the GroupTimeseries field if non-nil, zero value otherwise.

### GetGroupTimeseriesOk

`func (o *GroupsTimeseriesInnerValue) GetGroupTimeseriesOk() (*[]map[string]GroupsTimeseriesInnerValue, bool)`

GetGroupTimeseriesOk returns a tuple with the GroupTimeseries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupTimeseries

`func (o *GroupsTimeseriesInnerValue) SetGroupTimeseries(v []map[string]GroupsTimeseriesInnerValue)`

SetGroupTimeseries sets GroupTimeseries field to given value.


### GetSeries

`func (o *GroupsTimeseriesInnerValue) GetSeries() []Stats`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *GroupsTimeseriesInnerValue) GetSeriesOk() (*[]Stats, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *GroupsTimeseriesInnerValue) SetSeries(v []Stats)`

SetSeries sets Series field to given value.


### GetTotals

`func (o *GroupsTimeseriesInnerValue) GetTotals() Stats`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *GroupsTimeseriesInnerValue) GetTotalsOk() (*Stats, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *GroupsTimeseriesInnerValue) SetTotals(v Stats)`

SetTotals sets Totals field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


