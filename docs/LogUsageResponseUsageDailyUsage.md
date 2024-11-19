# LogUsageResponseUsageDailyUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Usage** | Pointer to [**Ints**](ints.md) | The total number of bytes written to the log. | [optional] 
**Day** | Pointer to **string** | The date in the format \&quot;YYYY-MM-DD\&quot;. | [optional] 

## Methods

### NewLogUsageResponseUsageDailyUsage

`func NewLogUsageResponseUsageDailyUsage() *LogUsageResponseUsageDailyUsage`

NewLogUsageResponseUsageDailyUsage instantiates a new LogUsageResponseUsageDailyUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogUsageResponseUsageDailyUsageWithDefaults

`func NewLogUsageResponseUsageDailyUsageWithDefaults() *LogUsageResponseUsageDailyUsage`

NewLogUsageResponseUsageDailyUsageWithDefaults instantiates a new LogUsageResponseUsageDailyUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsage

`func (o *LogUsageResponseUsageDailyUsage) GetUsage() Ints`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *LogUsageResponseUsageDailyUsage) GetUsageOk() (*Ints, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *LogUsageResponseUsageDailyUsage) SetUsage(v Ints)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *LogUsageResponseUsageDailyUsage) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetDay

`func (o *LogUsageResponseUsageDailyUsage) GetDay() string`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *LogUsageResponseUsageDailyUsage) GetDayOk() (*string, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *LogUsageResponseUsageDailyUsage) SetDay(v string)`

SetDay sets Day field to given value.

### HasDay

`func (o *LogUsageResponseUsageDailyUsage) HasDay() bool`

HasDay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


