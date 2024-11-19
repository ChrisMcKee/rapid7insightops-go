# LogsUsageResponsePerDayUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Period** | Pointer to [**AccountsUsageResponsePeriod**](AccountsUsageResponsePeriod.md) |  | [optional] 
**ReportInterval** | Pointer to **string** | Always set to \&quot;day\&quot; | [optional] 
**UsageUnits** | Pointer to **string** | Always set to \&quot;bytes\&quot;. | [optional] 
**Usage** | Pointer to [**[]LogsUsageResponsePerDayUsageUsageInner**](LogsUsageResponsePerDayUsageUsageInner.md) | The total number of bytes written to each log each day. | [optional] 

## Methods

### NewLogsUsageResponsePerDayUsage

`func NewLogsUsageResponsePerDayUsage() *LogsUsageResponsePerDayUsage`

NewLogsUsageResponsePerDayUsage instantiates a new LogsUsageResponsePerDayUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsUsageResponsePerDayUsageWithDefaults

`func NewLogsUsageResponsePerDayUsageWithDefaults() *LogsUsageResponsePerDayUsage`

NewLogsUsageResponsePerDayUsageWithDefaults instantiates a new LogsUsageResponsePerDayUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriod

`func (o *LogsUsageResponsePerDayUsage) GetPeriod() AccountsUsageResponsePeriod`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *LogsUsageResponsePerDayUsage) GetPeriodOk() (*AccountsUsageResponsePeriod, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *LogsUsageResponsePerDayUsage) SetPeriod(v AccountsUsageResponsePeriod)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *LogsUsageResponsePerDayUsage) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetReportInterval

`func (o *LogsUsageResponsePerDayUsage) GetReportInterval() string`

GetReportInterval returns the ReportInterval field if non-nil, zero value otherwise.

### GetReportIntervalOk

`func (o *LogsUsageResponsePerDayUsage) GetReportIntervalOk() (*string, bool)`

GetReportIntervalOk returns a tuple with the ReportInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportInterval

`func (o *LogsUsageResponsePerDayUsage) SetReportInterval(v string)`

SetReportInterval sets ReportInterval field to given value.

### HasReportInterval

`func (o *LogsUsageResponsePerDayUsage) HasReportInterval() bool`

HasReportInterval returns a boolean if a field has been set.

### GetUsageUnits

`func (o *LogsUsageResponsePerDayUsage) GetUsageUnits() string`

GetUsageUnits returns the UsageUnits field if non-nil, zero value otherwise.

### GetUsageUnitsOk

`func (o *LogsUsageResponsePerDayUsage) GetUsageUnitsOk() (*string, bool)`

GetUsageUnitsOk returns a tuple with the UsageUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageUnits

`func (o *LogsUsageResponsePerDayUsage) SetUsageUnits(v string)`

SetUsageUnits sets UsageUnits field to given value.

### HasUsageUnits

`func (o *LogsUsageResponsePerDayUsage) HasUsageUnits() bool`

HasUsageUnits returns a boolean if a field has been set.

### GetUsage

`func (o *LogsUsageResponsePerDayUsage) GetUsage() []LogsUsageResponsePerDayUsageUsageInner`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *LogsUsageResponsePerDayUsage) GetUsageOk() (*[]LogsUsageResponsePerDayUsageUsageInner, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *LogsUsageResponsePerDayUsage) SetUsage(v []LogsUsageResponsePerDayUsageUsageInner)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *LogsUsageResponsePerDayUsage) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


