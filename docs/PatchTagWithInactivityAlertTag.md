# PatchTagWithInactivityAlertTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the detection rule. | [optional] 
**Patterns** | Pointer to **[]string** | Use the &#x60;leql&#x60; parameter instead of this parameter. | [optional] 
**Leql** | Pointer to [**InactivityLeql**](InactivityLeql.md) |  | [optional] 
**Sources** | Pointer to [**[]SourcesIdArrayInner**](SourcesIdArrayInner.md) | The IDs of the logs that the detection rule operates on. | [optional] 
**Type** | **string** | Always set to \&quot;AlertNotify\&quot;. | 
**SubType** | Pointer to **string** | Always set to \&quot;InactivityAlert\&quot;. | [optional] 
**Description** | Pointer to **string** | The description of the detection rule. | [optional] 
**Actions** | Pointer to [**[]CreateTagActionInactivity**](CreateTagActionInactivity.md) | The notifications attached to the detection rule. | [optional] 
**UserData** | Pointer to **map[string]interface{}** | A list of key-value pairs that may indicate some auxiliary information about the change detection rule. | [optional] 
**TimeframeValue** | Pointer to **int32** | Defines the duration of inactivity before an alert triggers along with the &#x60;timeframe_period&#x60; parameter. | [optional] 
**TimeframePeriod** | Pointer to [**TimeframePeriod**](TimeframePeriod.md) |  | [optional] 
**Priority** | Pointer to **int32** | This ensures investigations are ordered by priority in Investigation Management. Must be &gt;&#x3D;0 | [optional] 

## Methods

### NewPatchTagWithInactivityAlertTag

`func NewPatchTagWithInactivityAlertTag(type_ string, ) *PatchTagWithInactivityAlertTag`

NewPatchTagWithInactivityAlertTag instantiates a new PatchTagWithInactivityAlertTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchTagWithInactivityAlertTagWithDefaults

`func NewPatchTagWithInactivityAlertTagWithDefaults() *PatchTagWithInactivityAlertTag`

NewPatchTagWithInactivityAlertTagWithDefaults instantiates a new PatchTagWithInactivityAlertTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchTagWithInactivityAlertTag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchTagWithInactivityAlertTag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchTagWithInactivityAlertTag) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchTagWithInactivityAlertTag) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPatterns

`func (o *PatchTagWithInactivityAlertTag) GetPatterns() []string`

GetPatterns returns the Patterns field if non-nil, zero value otherwise.

### GetPatternsOk

`func (o *PatchTagWithInactivityAlertTag) GetPatternsOk() (*[]string, bool)`

GetPatternsOk returns a tuple with the Patterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatterns

`func (o *PatchTagWithInactivityAlertTag) SetPatterns(v []string)`

SetPatterns sets Patterns field to given value.

### HasPatterns

`func (o *PatchTagWithInactivityAlertTag) HasPatterns() bool`

HasPatterns returns a boolean if a field has been set.

### GetLeql

`func (o *PatchTagWithInactivityAlertTag) GetLeql() InactivityLeql`

GetLeql returns the Leql field if non-nil, zero value otherwise.

### GetLeqlOk

`func (o *PatchTagWithInactivityAlertTag) GetLeqlOk() (*InactivityLeql, bool)`

GetLeqlOk returns a tuple with the Leql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeql

`func (o *PatchTagWithInactivityAlertTag) SetLeql(v InactivityLeql)`

SetLeql sets Leql field to given value.

### HasLeql

`func (o *PatchTagWithInactivityAlertTag) HasLeql() bool`

HasLeql returns a boolean if a field has been set.

### GetSources

`func (o *PatchTagWithInactivityAlertTag) GetSources() []SourcesIdArrayInner`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *PatchTagWithInactivityAlertTag) GetSourcesOk() (*[]SourcesIdArrayInner, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *PatchTagWithInactivityAlertTag) SetSources(v []SourcesIdArrayInner)`

SetSources sets Sources field to given value.

### HasSources

`func (o *PatchTagWithInactivityAlertTag) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetType

`func (o *PatchTagWithInactivityAlertTag) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchTagWithInactivityAlertTag) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchTagWithInactivityAlertTag) SetType(v string)`

SetType sets Type field to given value.


### GetSubType

`func (o *PatchTagWithInactivityAlertTag) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *PatchTagWithInactivityAlertTag) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *PatchTagWithInactivityAlertTag) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *PatchTagWithInactivityAlertTag) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetDescription

`func (o *PatchTagWithInactivityAlertTag) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchTagWithInactivityAlertTag) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchTagWithInactivityAlertTag) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchTagWithInactivityAlertTag) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActions

`func (o *PatchTagWithInactivityAlertTag) GetActions() []CreateTagActionInactivity`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *PatchTagWithInactivityAlertTag) GetActionsOk() (*[]CreateTagActionInactivity, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *PatchTagWithInactivityAlertTag) SetActions(v []CreateTagActionInactivity)`

SetActions sets Actions field to given value.

### HasActions

`func (o *PatchTagWithInactivityAlertTag) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetUserData

`func (o *PatchTagWithInactivityAlertTag) GetUserData() map[string]interface{}`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *PatchTagWithInactivityAlertTag) GetUserDataOk() (*map[string]interface{}, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *PatchTagWithInactivityAlertTag) SetUserData(v map[string]interface{})`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *PatchTagWithInactivityAlertTag) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### GetTimeframeValue

`func (o *PatchTagWithInactivityAlertTag) GetTimeframeValue() int32`

GetTimeframeValue returns the TimeframeValue field if non-nil, zero value otherwise.

### GetTimeframeValueOk

`func (o *PatchTagWithInactivityAlertTag) GetTimeframeValueOk() (*int32, bool)`

GetTimeframeValueOk returns a tuple with the TimeframeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframeValue

`func (o *PatchTagWithInactivityAlertTag) SetTimeframeValue(v int32)`

SetTimeframeValue sets TimeframeValue field to given value.

### HasTimeframeValue

`func (o *PatchTagWithInactivityAlertTag) HasTimeframeValue() bool`

HasTimeframeValue returns a boolean if a field has been set.

### GetTimeframePeriod

`func (o *PatchTagWithInactivityAlertTag) GetTimeframePeriod() TimeframePeriod`

GetTimeframePeriod returns the TimeframePeriod field if non-nil, zero value otherwise.

### GetTimeframePeriodOk

`func (o *PatchTagWithInactivityAlertTag) GetTimeframePeriodOk() (*TimeframePeriod, bool)`

GetTimeframePeriodOk returns a tuple with the TimeframePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframePeriod

`func (o *PatchTagWithInactivityAlertTag) SetTimeframePeriod(v TimeframePeriod)`

SetTimeframePeriod sets TimeframePeriod field to given value.

### HasTimeframePeriod

`func (o *PatchTagWithInactivityAlertTag) HasTimeframePeriod() bool`

HasTimeframePeriod returns a boolean if a field has been set.

### GetPriority

`func (o *PatchTagWithInactivityAlertTag) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *PatchTagWithInactivityAlertTag) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *PatchTagWithInactivityAlertTag) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *PatchTagWithInactivityAlertTag) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


