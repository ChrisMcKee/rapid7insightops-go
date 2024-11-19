# ActionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the notification. | [optional] 
**MinMatchesCount** | Pointer to **int32** | The number of log entries that the pattern detection rule needs to match for a notification to be triggered (within a window of time, controlled by the &#x60;min_matches_period&#x60; parameter). | [optional] [default to 0]
**MinMatchesPeriod** | Pointer to **string** | The time window used in combination with the &#x60;min_matches_count&#x60; parameter to determine if enough log entries have matched a pattern detection rule for an alert notification to be sent. For example, if &#x60;min_matches_count&#x3D;2&#x60;, and &#x60;min_matches_period&#x3D;\&quot;5Minute\&quot;&#x60;, then this notification will be triggered whenever two log entries match the pattern detection rule within 5 Minutes. | [optional] 
**MinReportCount** | Pointer to **int32** | The maximum number of notifications that will be triggered within a window of time, controlled by the &#x60;min_report_period&#x60; parameter. | [optional] [default to 1]
**MinReportPeriod** | Pointer to [**PeriodEnum**](PeriodEnum.md) |  | [optional] [default to HOUR]
**Targets** | Pointer to [**[]TargetResponse**](TargetResponse.md) | The list of notification targets for this notification. | [optional] 
**Enabled** | Pointer to **bool** | \\ When set to true, the notification is enabled. When set to false, the notification is disabled.  | [optional] 
**Type** | Pointer to **string** | Always set to \&quot;Alert\&quot;. | [optional] 

## Methods

### NewActionResponse

`func NewActionResponse() *ActionResponse`

NewActionResponse instantiates a new ActionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionResponseWithDefaults

`func NewActionResponseWithDefaults() *ActionResponse`

NewActionResponseWithDefaults instantiates a new ActionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActionResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ActionResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMinMatchesCount

`func (o *ActionResponse) GetMinMatchesCount() int32`

GetMinMatchesCount returns the MinMatchesCount field if non-nil, zero value otherwise.

### GetMinMatchesCountOk

`func (o *ActionResponse) GetMinMatchesCountOk() (*int32, bool)`

GetMinMatchesCountOk returns a tuple with the MinMatchesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinMatchesCount

`func (o *ActionResponse) SetMinMatchesCount(v int32)`

SetMinMatchesCount sets MinMatchesCount field to given value.

### HasMinMatchesCount

`func (o *ActionResponse) HasMinMatchesCount() bool`

HasMinMatchesCount returns a boolean if a field has been set.

### GetMinMatchesPeriod

`func (o *ActionResponse) GetMinMatchesPeriod() string`

GetMinMatchesPeriod returns the MinMatchesPeriod field if non-nil, zero value otherwise.

### GetMinMatchesPeriodOk

`func (o *ActionResponse) GetMinMatchesPeriodOk() (*string, bool)`

GetMinMatchesPeriodOk returns a tuple with the MinMatchesPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinMatchesPeriod

`func (o *ActionResponse) SetMinMatchesPeriod(v string)`

SetMinMatchesPeriod sets MinMatchesPeriod field to given value.

### HasMinMatchesPeriod

`func (o *ActionResponse) HasMinMatchesPeriod() bool`

HasMinMatchesPeriod returns a boolean if a field has been set.

### GetMinReportCount

`func (o *ActionResponse) GetMinReportCount() int32`

GetMinReportCount returns the MinReportCount field if non-nil, zero value otherwise.

### GetMinReportCountOk

`func (o *ActionResponse) GetMinReportCountOk() (*int32, bool)`

GetMinReportCountOk returns a tuple with the MinReportCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReportCount

`func (o *ActionResponse) SetMinReportCount(v int32)`

SetMinReportCount sets MinReportCount field to given value.

### HasMinReportCount

`func (o *ActionResponse) HasMinReportCount() bool`

HasMinReportCount returns a boolean if a field has been set.

### GetMinReportPeriod

`func (o *ActionResponse) GetMinReportPeriod() PeriodEnum`

GetMinReportPeriod returns the MinReportPeriod field if non-nil, zero value otherwise.

### GetMinReportPeriodOk

`func (o *ActionResponse) GetMinReportPeriodOk() (*PeriodEnum, bool)`

GetMinReportPeriodOk returns a tuple with the MinReportPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReportPeriod

`func (o *ActionResponse) SetMinReportPeriod(v PeriodEnum)`

SetMinReportPeriod sets MinReportPeriod field to given value.

### HasMinReportPeriod

`func (o *ActionResponse) HasMinReportPeriod() bool`

HasMinReportPeriod returns a boolean if a field has been set.

### GetTargets

`func (o *ActionResponse) GetTargets() []TargetResponse`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *ActionResponse) GetTargetsOk() (*[]TargetResponse, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *ActionResponse) SetTargets(v []TargetResponse)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *ActionResponse) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetEnabled

`func (o *ActionResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ActionResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ActionResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ActionResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetType

`func (o *ActionResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActionResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActionResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ActionResponse) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


