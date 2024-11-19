# LogsUsageResponsePerDayUsageUsageInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interval** | Pointer to **string** | The date in the format \&quot;YYYY-MM-DD\&quot;. | [optional] 
**LogUsage** | Pointer to [**[]LogsUsageResponsePerDayUsageUsageInnerLogUsageInner**](LogsUsageResponsePerDayUsageUsageInnerLogUsageInner.md) | The total number of bytes written to each log. | [optional] 

## Methods

### NewLogsUsageResponsePerDayUsageUsageInner

`func NewLogsUsageResponsePerDayUsageUsageInner() *LogsUsageResponsePerDayUsageUsageInner`

NewLogsUsageResponsePerDayUsageUsageInner instantiates a new LogsUsageResponsePerDayUsageUsageInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsUsageResponsePerDayUsageUsageInnerWithDefaults

`func NewLogsUsageResponsePerDayUsageUsageInnerWithDefaults() *LogsUsageResponsePerDayUsageUsageInner`

NewLogsUsageResponsePerDayUsageUsageInnerWithDefaults instantiates a new LogsUsageResponsePerDayUsageUsageInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterval

`func (o *LogsUsageResponsePerDayUsageUsageInner) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *LogsUsageResponsePerDayUsageUsageInner) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *LogsUsageResponsePerDayUsageUsageInner) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *LogsUsageResponsePerDayUsageUsageInner) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetLogUsage

`func (o *LogsUsageResponsePerDayUsageUsageInner) GetLogUsage() []LogsUsageResponsePerDayUsageUsageInnerLogUsageInner`

GetLogUsage returns the LogUsage field if non-nil, zero value otherwise.

### GetLogUsageOk

`func (o *LogsUsageResponsePerDayUsageUsageInner) GetLogUsageOk() (*[]LogsUsageResponsePerDayUsageUsageInnerLogUsageInner, bool)`

GetLogUsageOk returns a tuple with the LogUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogUsage

`func (o *LogsUsageResponsePerDayUsageUsageInner) SetLogUsage(v []LogsUsageResponsePerDayUsageUsageInnerLogUsageInner)`

SetLogUsage sets LogUsage field to given value.

### HasLogUsage

`func (o *LogsUsageResponsePerDayUsageUsageInner) HasLogUsage() bool`

HasLogUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


