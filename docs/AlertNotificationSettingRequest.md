# AlertNotificationSettingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinMatchesCount** | **int32** | The number of log entries that the pattern detection rule needs to match for a notification to be triggered (within a window of time, controlled by the &#x60;min_matches_period&#x60; parameter). Only for pattern detection rules. This parameter is ignored when used with change detection rules or inactivity detection rules. | [default to 0]
**MinMatchesPeriod** | **string** | The time window used in combination with the &#x60;min_matches_count&#x60; parameter to determine if enough log entries have matched a pattern detection rule for an alert notification to be sent. For example, if &#x60;min_matches_count&#x3D;2&#x60;, and &#x60;min_matches_period&#x3D;\&quot;5Minute\&quot;&#x60;, then this notification will be triggered whenever two log entries match the pattern detection rule within 5 minutes. Only for pattern detection rules. This parameter is ignored when used with change detection rules or inactivity detection rules. | 
**MinReportCount** | **int32** | The maximum number of notifications that will be triggered within a window of time, controlled by the &#x60;min_report_period&#x60; parameter. | [default to 1]
**MinReportPeriod** | [**PeriodEnum**](PeriodEnum.md) |  | [default to HOUR]
**Targets** | [**[]TargetResponse**](TargetResponse.md) | The list of targets for the action. | 
**Enabled** | **bool** | \\ When set to true, the action is enabled. When set to false, the action is disabled.  | 
**Type** | **string** | Always set to \&quot;Alert\&quot;. | 

## Methods

### NewAlertNotificationSettingRequest

`func NewAlertNotificationSettingRequest(minMatchesCount int32, minMatchesPeriod string, minReportCount int32, minReportPeriod PeriodEnum, targets []TargetResponse, enabled bool, type_ string, ) *AlertNotificationSettingRequest`

NewAlertNotificationSettingRequest instantiates a new AlertNotificationSettingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertNotificationSettingRequestWithDefaults

`func NewAlertNotificationSettingRequestWithDefaults() *AlertNotificationSettingRequest`

NewAlertNotificationSettingRequestWithDefaults instantiates a new AlertNotificationSettingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinMatchesCount

`func (o *AlertNotificationSettingRequest) GetMinMatchesCount() int32`

GetMinMatchesCount returns the MinMatchesCount field if non-nil, zero value otherwise.

### GetMinMatchesCountOk

`func (o *AlertNotificationSettingRequest) GetMinMatchesCountOk() (*int32, bool)`

GetMinMatchesCountOk returns a tuple with the MinMatchesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinMatchesCount

`func (o *AlertNotificationSettingRequest) SetMinMatchesCount(v int32)`

SetMinMatchesCount sets MinMatchesCount field to given value.


### GetMinMatchesPeriod

`func (o *AlertNotificationSettingRequest) GetMinMatchesPeriod() string`

GetMinMatchesPeriod returns the MinMatchesPeriod field if non-nil, zero value otherwise.

### GetMinMatchesPeriodOk

`func (o *AlertNotificationSettingRequest) GetMinMatchesPeriodOk() (*string, bool)`

GetMinMatchesPeriodOk returns a tuple with the MinMatchesPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinMatchesPeriod

`func (o *AlertNotificationSettingRequest) SetMinMatchesPeriod(v string)`

SetMinMatchesPeriod sets MinMatchesPeriod field to given value.


### GetMinReportCount

`func (o *AlertNotificationSettingRequest) GetMinReportCount() int32`

GetMinReportCount returns the MinReportCount field if non-nil, zero value otherwise.

### GetMinReportCountOk

`func (o *AlertNotificationSettingRequest) GetMinReportCountOk() (*int32, bool)`

GetMinReportCountOk returns a tuple with the MinReportCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReportCount

`func (o *AlertNotificationSettingRequest) SetMinReportCount(v int32)`

SetMinReportCount sets MinReportCount field to given value.


### GetMinReportPeriod

`func (o *AlertNotificationSettingRequest) GetMinReportPeriod() PeriodEnum`

GetMinReportPeriod returns the MinReportPeriod field if non-nil, zero value otherwise.

### GetMinReportPeriodOk

`func (o *AlertNotificationSettingRequest) GetMinReportPeriodOk() (*PeriodEnum, bool)`

GetMinReportPeriodOk returns a tuple with the MinReportPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReportPeriod

`func (o *AlertNotificationSettingRequest) SetMinReportPeriod(v PeriodEnum)`

SetMinReportPeriod sets MinReportPeriod field to given value.


### GetTargets

`func (o *AlertNotificationSettingRequest) GetTargets() []TargetResponse`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *AlertNotificationSettingRequest) GetTargetsOk() (*[]TargetResponse, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *AlertNotificationSettingRequest) SetTargets(v []TargetResponse)`

SetTargets sets Targets field to given value.


### GetEnabled

`func (o *AlertNotificationSettingRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AlertNotificationSettingRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AlertNotificationSettingRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetType

`func (o *AlertNotificationSettingRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AlertNotificationSettingRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AlertNotificationSettingRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


