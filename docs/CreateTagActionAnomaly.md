# CreateTagActionAnomaly

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | \\ When set to true, the notification is enabled. When set to false, the notification is disabled.  | 
**MinReportCount** | **int32** | The maximum number of notifications that will be triggered within a window of time, controlled by the &#x60;min_report_period&#x60; parameter. | [default to 1]
**MinReportPeriod** | [**PeriodEnum**](PeriodEnum.md) |  | [default to HOUR]
**Targets** | [**[]CreateTagTarget**](CreateTagTarget.md) | The list of notification targets for this notification. | 
**Type** | Pointer to **string** | Always set to \&quot;Alert\&quot;. | [optional] 

## Methods

### NewCreateTagActionAnomaly

`func NewCreateTagActionAnomaly(enabled bool, minReportCount int32, minReportPeriod PeriodEnum, targets []CreateTagTarget, ) *CreateTagActionAnomaly`

NewCreateTagActionAnomaly instantiates a new CreateTagActionAnomaly object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTagActionAnomalyWithDefaults

`func NewCreateTagActionAnomalyWithDefaults() *CreateTagActionAnomaly`

NewCreateTagActionAnomalyWithDefaults instantiates a new CreateTagActionAnomaly object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *CreateTagActionAnomaly) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateTagActionAnomaly) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateTagActionAnomaly) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMinReportCount

`func (o *CreateTagActionAnomaly) GetMinReportCount() int32`

GetMinReportCount returns the MinReportCount field if non-nil, zero value otherwise.

### GetMinReportCountOk

`func (o *CreateTagActionAnomaly) GetMinReportCountOk() (*int32, bool)`

GetMinReportCountOk returns a tuple with the MinReportCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReportCount

`func (o *CreateTagActionAnomaly) SetMinReportCount(v int32)`

SetMinReportCount sets MinReportCount field to given value.


### GetMinReportPeriod

`func (o *CreateTagActionAnomaly) GetMinReportPeriod() PeriodEnum`

GetMinReportPeriod returns the MinReportPeriod field if non-nil, zero value otherwise.

### GetMinReportPeriodOk

`func (o *CreateTagActionAnomaly) GetMinReportPeriodOk() (*PeriodEnum, bool)`

GetMinReportPeriodOk returns a tuple with the MinReportPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReportPeriod

`func (o *CreateTagActionAnomaly) SetMinReportPeriod(v PeriodEnum)`

SetMinReportPeriod sets MinReportPeriod field to given value.


### GetTargets

`func (o *CreateTagActionAnomaly) GetTargets() []CreateTagTarget`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *CreateTagActionAnomaly) GetTargetsOk() (*[]CreateTagTarget, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *CreateTagActionAnomaly) SetTargets(v []CreateTagTarget)`

SetTargets sets Targets field to given value.


### GetType

`func (o *CreateTagActionAnomaly) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateTagActionAnomaly) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateTagActionAnomaly) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateTagActionAnomaly) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


