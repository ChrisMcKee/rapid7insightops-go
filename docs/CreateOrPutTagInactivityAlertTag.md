# CreateOrPutTagInactivityAlertTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the detection rule. | 
**Patterns** | Pointer to **[]string** | Use the &#x60;leql&#x60; parameter instead of this parameter. | [optional] 
**Sources** | [**[]SourcesIdArrayInner**](SourcesIdArrayInner.md) | The IDs of the logs that the detection rule operates on. | 
**Type** | **string** | The Always set to \&quot;AlertNotify\&quot;. | 
**SubType** | **string** | Always set to \&quot;InactivityAlert\&quot;. | 
**Description** | Pointer to **string** | The description of the detection rule. | [optional] 
**Actions** | Pointer to [**[]CreateTagActionInactivity**](CreateTagActionInactivity.md) | The notifications attached to the detection rule. | [optional] 
**UserData** | Pointer to **map[string]interface{}** | A list of key-value pairs that may indicate some auxiliary information about the change detection rule. | [optional] 
**TimeframeValue** | Pointer to **int32** | Defines the duration of inactivity before an alert triggers along with the &#x60;timeframe_period&#x60; parameter. | [optional] 
**TimeframePeriod** | Pointer to [**TimeframePeriod**](TimeframePeriod.md) |  | [optional] 
**Leql** | [**InactivityLeql**](InactivityLeql.md) |  | 
**Priority** | Pointer to **int32** | This ensures investigations are ordered by priority in Investigation Management. Must be &gt;&#x3D;0 | [optional] 

## Methods

### NewCreateOrPutTagInactivityAlertTag

`func NewCreateOrPutTagInactivityAlertTag(name string, sources []SourcesIdArrayInner, type_ string, subType string, leql InactivityLeql, ) *CreateOrPutTagInactivityAlertTag`

NewCreateOrPutTagInactivityAlertTag instantiates a new CreateOrPutTagInactivityAlertTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrPutTagInactivityAlertTagWithDefaults

`func NewCreateOrPutTagInactivityAlertTagWithDefaults() *CreateOrPutTagInactivityAlertTag`

NewCreateOrPutTagInactivityAlertTagWithDefaults instantiates a new CreateOrPutTagInactivityAlertTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateOrPutTagInactivityAlertTag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOrPutTagInactivityAlertTag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOrPutTagInactivityAlertTag) SetName(v string)`

SetName sets Name field to given value.


### GetPatterns

`func (o *CreateOrPutTagInactivityAlertTag) GetPatterns() []string`

GetPatterns returns the Patterns field if non-nil, zero value otherwise.

### GetPatternsOk

`func (o *CreateOrPutTagInactivityAlertTag) GetPatternsOk() (*[]string, bool)`

GetPatternsOk returns a tuple with the Patterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatterns

`func (o *CreateOrPutTagInactivityAlertTag) SetPatterns(v []string)`

SetPatterns sets Patterns field to given value.

### HasPatterns

`func (o *CreateOrPutTagInactivityAlertTag) HasPatterns() bool`

HasPatterns returns a boolean if a field has been set.

### GetSources

`func (o *CreateOrPutTagInactivityAlertTag) GetSources() []SourcesIdArrayInner`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *CreateOrPutTagInactivityAlertTag) GetSourcesOk() (*[]SourcesIdArrayInner, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *CreateOrPutTagInactivityAlertTag) SetSources(v []SourcesIdArrayInner)`

SetSources sets Sources field to given value.


### GetType

`func (o *CreateOrPutTagInactivityAlertTag) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateOrPutTagInactivityAlertTag) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateOrPutTagInactivityAlertTag) SetType(v string)`

SetType sets Type field to given value.


### GetSubType

`func (o *CreateOrPutTagInactivityAlertTag) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *CreateOrPutTagInactivityAlertTag) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *CreateOrPutTagInactivityAlertTag) SetSubType(v string)`

SetSubType sets SubType field to given value.


### GetDescription

`func (o *CreateOrPutTagInactivityAlertTag) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateOrPutTagInactivityAlertTag) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateOrPutTagInactivityAlertTag) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateOrPutTagInactivityAlertTag) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActions

`func (o *CreateOrPutTagInactivityAlertTag) GetActions() []CreateTagActionInactivity`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *CreateOrPutTagInactivityAlertTag) GetActionsOk() (*[]CreateTagActionInactivity, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *CreateOrPutTagInactivityAlertTag) SetActions(v []CreateTagActionInactivity)`

SetActions sets Actions field to given value.

### HasActions

`func (o *CreateOrPutTagInactivityAlertTag) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetUserData

`func (o *CreateOrPutTagInactivityAlertTag) GetUserData() map[string]interface{}`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *CreateOrPutTagInactivityAlertTag) GetUserDataOk() (*map[string]interface{}, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *CreateOrPutTagInactivityAlertTag) SetUserData(v map[string]interface{})`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *CreateOrPutTagInactivityAlertTag) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### GetTimeframeValue

`func (o *CreateOrPutTagInactivityAlertTag) GetTimeframeValue() int32`

GetTimeframeValue returns the TimeframeValue field if non-nil, zero value otherwise.

### GetTimeframeValueOk

`func (o *CreateOrPutTagInactivityAlertTag) GetTimeframeValueOk() (*int32, bool)`

GetTimeframeValueOk returns a tuple with the TimeframeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframeValue

`func (o *CreateOrPutTagInactivityAlertTag) SetTimeframeValue(v int32)`

SetTimeframeValue sets TimeframeValue field to given value.

### HasTimeframeValue

`func (o *CreateOrPutTagInactivityAlertTag) HasTimeframeValue() bool`

HasTimeframeValue returns a boolean if a field has been set.

### GetTimeframePeriod

`func (o *CreateOrPutTagInactivityAlertTag) GetTimeframePeriod() TimeframePeriod`

GetTimeframePeriod returns the TimeframePeriod field if non-nil, zero value otherwise.

### GetTimeframePeriodOk

`func (o *CreateOrPutTagInactivityAlertTag) GetTimeframePeriodOk() (*TimeframePeriod, bool)`

GetTimeframePeriodOk returns a tuple with the TimeframePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframePeriod

`func (o *CreateOrPutTagInactivityAlertTag) SetTimeframePeriod(v TimeframePeriod)`

SetTimeframePeriod sets TimeframePeriod field to given value.

### HasTimeframePeriod

`func (o *CreateOrPutTagInactivityAlertTag) HasTimeframePeriod() bool`

HasTimeframePeriod returns a boolean if a field has been set.

### GetLeql

`func (o *CreateOrPutTagInactivityAlertTag) GetLeql() InactivityLeql`

GetLeql returns the Leql field if non-nil, zero value otherwise.

### GetLeqlOk

`func (o *CreateOrPutTagInactivityAlertTag) GetLeqlOk() (*InactivityLeql, bool)`

GetLeqlOk returns a tuple with the Leql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeql

`func (o *CreateOrPutTagInactivityAlertTag) SetLeql(v InactivityLeql)`

SetLeql sets Leql field to given value.


### GetPriority

`func (o *CreateOrPutTagInactivityAlertTag) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CreateOrPutTagInactivityAlertTag) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CreateOrPutTagInactivityAlertTag) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CreateOrPutTagInactivityAlertTag) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


