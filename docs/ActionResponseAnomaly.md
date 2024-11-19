# ActionResponseAnomaly

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the notification. | [optional] 
**MinReportCount** | Pointer to **int32** | The maximum number of notifications that will be triggered within a window of time, controlled by the &#x60;min_report_period&#x60; parameter. | [optional] [default to 1]
**MinReportPeriod** | Pointer to [**PeriodEnum**](PeriodEnum.md) |  | [optional] [default to HOUR]
**Targets** | Pointer to [**[]TargetResponse**](TargetResponse.md) | The list of notification targets for this notification. | [optional] 
**Enabled** | Pointer to **bool** | \\ When set to true, the notification is enabled. When set to false, the notification is disabled.  | [optional] 
**Type** | Pointer to **string** | Always set to \&quot;Alert\&quot;. | [optional] 

## Methods

### NewActionResponseAnomaly

`func NewActionResponseAnomaly() *ActionResponseAnomaly`

NewActionResponseAnomaly instantiates a new ActionResponseAnomaly object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionResponseAnomalyWithDefaults

`func NewActionResponseAnomalyWithDefaults() *ActionResponseAnomaly`

NewActionResponseAnomalyWithDefaults instantiates a new ActionResponseAnomaly object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActionResponseAnomaly) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActionResponseAnomaly) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActionResponseAnomaly) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ActionResponseAnomaly) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMinReportCount

`func (o *ActionResponseAnomaly) GetMinReportCount() int32`

GetMinReportCount returns the MinReportCount field if non-nil, zero value otherwise.

### GetMinReportCountOk

`func (o *ActionResponseAnomaly) GetMinReportCountOk() (*int32, bool)`

GetMinReportCountOk returns a tuple with the MinReportCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReportCount

`func (o *ActionResponseAnomaly) SetMinReportCount(v int32)`

SetMinReportCount sets MinReportCount field to given value.

### HasMinReportCount

`func (o *ActionResponseAnomaly) HasMinReportCount() bool`

HasMinReportCount returns a boolean if a field has been set.

### GetMinReportPeriod

`func (o *ActionResponseAnomaly) GetMinReportPeriod() PeriodEnum`

GetMinReportPeriod returns the MinReportPeriod field if non-nil, zero value otherwise.

### GetMinReportPeriodOk

`func (o *ActionResponseAnomaly) GetMinReportPeriodOk() (*PeriodEnum, bool)`

GetMinReportPeriodOk returns a tuple with the MinReportPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReportPeriod

`func (o *ActionResponseAnomaly) SetMinReportPeriod(v PeriodEnum)`

SetMinReportPeriod sets MinReportPeriod field to given value.

### HasMinReportPeriod

`func (o *ActionResponseAnomaly) HasMinReportPeriod() bool`

HasMinReportPeriod returns a boolean if a field has been set.

### GetTargets

`func (o *ActionResponseAnomaly) GetTargets() []TargetResponse`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *ActionResponseAnomaly) GetTargetsOk() (*[]TargetResponse, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *ActionResponseAnomaly) SetTargets(v []TargetResponse)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *ActionResponseAnomaly) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetEnabled

`func (o *ActionResponseAnomaly) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ActionResponseAnomaly) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ActionResponseAnomaly) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ActionResponseAnomaly) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetType

`func (o *ActionResponseAnomaly) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActionResponseAnomaly) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActionResponseAnomaly) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ActionResponseAnomaly) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


