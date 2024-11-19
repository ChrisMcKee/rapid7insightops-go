# CreateTagAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | \\ When set to true, the notification is enabled. When set to false, the notification is disabled.  | 
**MinMatchesCount** | **int32** | The number of log entries that the pattern detection rule needs to match for a notification to be triggered (within a window of time, controlled by the &#x60;min_matches_period&#x60; parameter). | [default to 0]
**MinMatchesPeriod** | **string** | The time window used in combination with the &#x60;min_matches_count&#x60; parameter to determine if enough log entries have matched a pattern detection rule for an alert notification to be sent. For example, if &#x60;min_matches_count&#x3D;2&#x60;, and &#x60;min_matches_period&#x3D;\&quot;5Minute\&quot;&#x60;, then this notification will be triggered whenever two log entries match the pattern detection rule within 5 minutes. | 
**MinReportCount** | **int32** | The maximum number of notifications that will be triggered within a window of time, controlled by the &#x60;min_report_period&#x60; parameter. | [default to 1]
**MinReportPeriod** | [**PeriodEnum**](PeriodEnum.md) |  | [default to HOUR]
**Targets** | [**[]CreateTagTarget**](CreateTagTarget.md) | The list of notification targets for this notification. | 
**Type** | Pointer to **string** | Always set to \&quot;Alert\&quot;. | [optional] 

## Methods

### NewCreateTagAction

`func NewCreateTagAction(enabled bool, minMatchesCount int32, minMatchesPeriod string, minReportCount int32, minReportPeriod PeriodEnum, targets []CreateTagTarget, ) *CreateTagAction`

NewCreateTagAction instantiates a new CreateTagAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTagActionWithDefaults

`func NewCreateTagActionWithDefaults() *CreateTagAction`

NewCreateTagActionWithDefaults instantiates a new CreateTagAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *CreateTagAction) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateTagAction) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateTagAction) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMinMatchesCount

`func (o *CreateTagAction) GetMinMatchesCount() int32`

GetMinMatchesCount returns the MinMatchesCount field if non-nil, zero value otherwise.

### GetMinMatchesCountOk

`func (o *CreateTagAction) GetMinMatchesCountOk() (*int32, bool)`

GetMinMatchesCountOk returns a tuple with the MinMatchesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinMatchesCount

`func (o *CreateTagAction) SetMinMatchesCount(v int32)`

SetMinMatchesCount sets MinMatchesCount field to given value.


### GetMinMatchesPeriod

`func (o *CreateTagAction) GetMinMatchesPeriod() string`

GetMinMatchesPeriod returns the MinMatchesPeriod field if non-nil, zero value otherwise.

### GetMinMatchesPeriodOk

`func (o *CreateTagAction) GetMinMatchesPeriodOk() (*string, bool)`

GetMinMatchesPeriodOk returns a tuple with the MinMatchesPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinMatchesPeriod

`func (o *CreateTagAction) SetMinMatchesPeriod(v string)`

SetMinMatchesPeriod sets MinMatchesPeriod field to given value.


### GetMinReportCount

`func (o *CreateTagAction) GetMinReportCount() int32`

GetMinReportCount returns the MinReportCount field if non-nil, zero value otherwise.

### GetMinReportCountOk

`func (o *CreateTagAction) GetMinReportCountOk() (*int32, bool)`

GetMinReportCountOk returns a tuple with the MinReportCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReportCount

`func (o *CreateTagAction) SetMinReportCount(v int32)`

SetMinReportCount sets MinReportCount field to given value.


### GetMinReportPeriod

`func (o *CreateTagAction) GetMinReportPeriod() PeriodEnum`

GetMinReportPeriod returns the MinReportPeriod field if non-nil, zero value otherwise.

### GetMinReportPeriodOk

`func (o *CreateTagAction) GetMinReportPeriodOk() (*PeriodEnum, bool)`

GetMinReportPeriodOk returns a tuple with the MinReportPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReportPeriod

`func (o *CreateTagAction) SetMinReportPeriod(v PeriodEnum)`

SetMinReportPeriod sets MinReportPeriod field to given value.


### GetTargets

`func (o *CreateTagAction) GetTargets() []CreateTagTarget`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *CreateTagAction) GetTargetsOk() (*[]CreateTagTarget, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *CreateTagAction) SetTargets(v []CreateTagTarget)`

SetTargets sets Targets field to given value.


### GetType

`func (o *CreateTagAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateTagAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateTagAction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateTagAction) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


