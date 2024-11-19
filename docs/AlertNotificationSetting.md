# AlertNotificationSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the notification. | [optional] 
**MinMatchesCount** | Pointer to **int32** | The number of log entries that the pattern detection rule needs to match for a notification to be triggered (within a window of time, controlled by the &#x60;min_matches_period&#x60; parameter). Only for pattern detection rules. This parameter is ignored when used with change detection rules or inactivity detection rules. | [optional] [default to 0]
**MinMatchesPeriod** | Pointer to **string** | The time window used in combination with the &#x60;min_matches_count&#x60; parameter to determine if enough log entries have matched a pattern detection rule for an alert notification to be sent. For example, if &#x60;min_matches_count&#x3D;2&#x60;, and &#x60;min_matches_period&#x3D;\&quot;5Minute\&quot;&#x60;, then this notification will be triggered whenever two log entries match the pattern detection rule within 5 minutes. Only for pattern detection rules. This parameter is ignored when used with change detection rules or inactivity detection rules. | [optional] 
**MinReportCount** | Pointer to **int32** | The maximum number of notifications that will be triggered within a window of time, controlled by the &#x60;min_report_period&#x60; parameter. | [optional] [default to 1]
**MinReportPeriod** | Pointer to [**PeriodEnum**](PeriodEnum.md) |  | [optional] [default to HOUR]
**Targets** | Pointer to [**[]TargetResponse**](TargetResponse.md) | The list of targets for the action. | [optional] 
**Enabled** | Pointer to **bool** | \\ When set to true, the action is enabled. When set to false, the action is disabled.  | [optional] 
**Type** | Pointer to **string** | Always set to \&quot;Alert\&quot;. | [optional] 

## Methods

### NewAlertNotificationSetting

`func NewAlertNotificationSetting() *AlertNotificationSetting`

NewAlertNotificationSetting instantiates a new AlertNotificationSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertNotificationSettingWithDefaults

`func NewAlertNotificationSettingWithDefaults() *AlertNotificationSetting`

NewAlertNotificationSettingWithDefaults instantiates a new AlertNotificationSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AlertNotificationSetting) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlertNotificationSetting) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlertNotificationSetting) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AlertNotificationSetting) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMinMatchesCount

`func (o *AlertNotificationSetting) GetMinMatchesCount() int32`

GetMinMatchesCount returns the MinMatchesCount field if non-nil, zero value otherwise.

### GetMinMatchesCountOk

`func (o *AlertNotificationSetting) GetMinMatchesCountOk() (*int32, bool)`

GetMinMatchesCountOk returns a tuple with the MinMatchesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinMatchesCount

`func (o *AlertNotificationSetting) SetMinMatchesCount(v int32)`

SetMinMatchesCount sets MinMatchesCount field to given value.

### HasMinMatchesCount

`func (o *AlertNotificationSetting) HasMinMatchesCount() bool`

HasMinMatchesCount returns a boolean if a field has been set.

### GetMinMatchesPeriod

`func (o *AlertNotificationSetting) GetMinMatchesPeriod() string`

GetMinMatchesPeriod returns the MinMatchesPeriod field if non-nil, zero value otherwise.

### GetMinMatchesPeriodOk

`func (o *AlertNotificationSetting) GetMinMatchesPeriodOk() (*string, bool)`

GetMinMatchesPeriodOk returns a tuple with the MinMatchesPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinMatchesPeriod

`func (o *AlertNotificationSetting) SetMinMatchesPeriod(v string)`

SetMinMatchesPeriod sets MinMatchesPeriod field to given value.

### HasMinMatchesPeriod

`func (o *AlertNotificationSetting) HasMinMatchesPeriod() bool`

HasMinMatchesPeriod returns a boolean if a field has been set.

### GetMinReportCount

`func (o *AlertNotificationSetting) GetMinReportCount() int32`

GetMinReportCount returns the MinReportCount field if non-nil, zero value otherwise.

### GetMinReportCountOk

`func (o *AlertNotificationSetting) GetMinReportCountOk() (*int32, bool)`

GetMinReportCountOk returns a tuple with the MinReportCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReportCount

`func (o *AlertNotificationSetting) SetMinReportCount(v int32)`

SetMinReportCount sets MinReportCount field to given value.

### HasMinReportCount

`func (o *AlertNotificationSetting) HasMinReportCount() bool`

HasMinReportCount returns a boolean if a field has been set.

### GetMinReportPeriod

`func (o *AlertNotificationSetting) GetMinReportPeriod() PeriodEnum`

GetMinReportPeriod returns the MinReportPeriod field if non-nil, zero value otherwise.

### GetMinReportPeriodOk

`func (o *AlertNotificationSetting) GetMinReportPeriodOk() (*PeriodEnum, bool)`

GetMinReportPeriodOk returns a tuple with the MinReportPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReportPeriod

`func (o *AlertNotificationSetting) SetMinReportPeriod(v PeriodEnum)`

SetMinReportPeriod sets MinReportPeriod field to given value.

### HasMinReportPeriod

`func (o *AlertNotificationSetting) HasMinReportPeriod() bool`

HasMinReportPeriod returns a boolean if a field has been set.

### GetTargets

`func (o *AlertNotificationSetting) GetTargets() []TargetResponse`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *AlertNotificationSetting) GetTargetsOk() (*[]TargetResponse, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *AlertNotificationSetting) SetTargets(v []TargetResponse)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *AlertNotificationSetting) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetEnabled

`func (o *AlertNotificationSetting) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AlertNotificationSetting) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AlertNotificationSetting) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AlertNotificationSetting) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetType

`func (o *AlertNotificationSetting) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AlertNotificationSetting) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AlertNotificationSetting) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AlertNotificationSetting) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


