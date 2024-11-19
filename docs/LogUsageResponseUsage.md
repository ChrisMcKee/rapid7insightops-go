# LogUsageResponseUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the log (UUID). | [optional] 
**Period** | Pointer to [**AccountsUsageResponsePeriod**](AccountsUsageResponsePeriod.md) |  | [optional] 
**DailyUsage** | Pointer to [**LogUsageResponseUsageDailyUsage**](LogUsageResponseUsageDailyUsage.md) |  | [optional] 

## Methods

### NewLogUsageResponseUsage

`func NewLogUsageResponseUsage() *LogUsageResponseUsage`

NewLogUsageResponseUsage instantiates a new LogUsageResponseUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogUsageResponseUsageWithDefaults

`func NewLogUsageResponseUsageWithDefaults() *LogUsageResponseUsage`

NewLogUsageResponseUsageWithDefaults instantiates a new LogUsageResponseUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LogUsageResponseUsage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogUsageResponseUsage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogUsageResponseUsage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LogUsageResponseUsage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPeriod

`func (o *LogUsageResponseUsage) GetPeriod() AccountsUsageResponsePeriod`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *LogUsageResponseUsage) GetPeriodOk() (*AccountsUsageResponsePeriod, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *LogUsageResponseUsage) SetPeriod(v AccountsUsageResponsePeriod)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *LogUsageResponseUsage) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetDailyUsage

`func (o *LogUsageResponseUsage) GetDailyUsage() LogUsageResponseUsageDailyUsage`

GetDailyUsage returns the DailyUsage field if non-nil, zero value otherwise.

### GetDailyUsageOk

`func (o *LogUsageResponseUsage) GetDailyUsageOk() (*LogUsageResponseUsageDailyUsage, bool)`

GetDailyUsageOk returns a tuple with the DailyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyUsage

`func (o *LogUsageResponseUsage) SetDailyUsage(v LogUsageResponseUsageDailyUsage)`

SetDailyUsage sets DailyUsage field to given value.

### HasDailyUsage

`func (o *LogUsageResponseUsage) HasDailyUsage() bool`

HasDailyUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


