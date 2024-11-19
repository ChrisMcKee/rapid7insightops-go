# PatchTagWithAnomalyAlertTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the detection rule. | [optional] 
**Sources** | Pointer to [**[]SourcesIdArrayInner**](SourcesIdArrayInner.md) | The IDs of the logs that the detection rule operates on. | [optional] 
**Type** | **string** | Always set to \&quot;AlertNotify\&quot;. | 
**SubType** | Pointer to **string** | Always set to \&quot;AnomalyAlert\&quot;. | [optional] 
**Description** | Pointer to **string** | The description of the detection rule. | [optional] 
**Actions** | Pointer to [**[]CreateTagActionAnomaly**](CreateTagActionAnomaly.md) | The notifications attached to the detection rule. | [optional] 
**UserData** | Pointer to **map[string]interface{}** | A list of key-value pairs that may indicate some auxiliary information about the change detection rule. | [optional] 
**ScheduledQuery** | Pointer to [**CreateScheduledQuery**](CreateScheduledQuery.md) |  | [optional] 
**ThresholdValue** | Pointer to **int32** | The percentage that defines when to trigger notifications. The percentage can be positive or negative. For example, if the percentage is set to -50, then notifications will be triggered when the query result for the time range defined by the &#x60;time_value&#x60; and &#x60;time_period&#x60; parameters is 50% smaller when compared to the same query over the time range defined by the &#x60;timeframe_value&#x60; and &#x60;timeframe_period&#x60; parameters. | [optional] 
**Priority** | Pointer to **int32** | This ensures investigations are ordered by priority in Investigation Management. Must be &gt;&#x3D;0 | [optional] 

## Methods

### NewPatchTagWithAnomalyAlertTag

`func NewPatchTagWithAnomalyAlertTag(type_ string, ) *PatchTagWithAnomalyAlertTag`

NewPatchTagWithAnomalyAlertTag instantiates a new PatchTagWithAnomalyAlertTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchTagWithAnomalyAlertTagWithDefaults

`func NewPatchTagWithAnomalyAlertTagWithDefaults() *PatchTagWithAnomalyAlertTag`

NewPatchTagWithAnomalyAlertTagWithDefaults instantiates a new PatchTagWithAnomalyAlertTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchTagWithAnomalyAlertTag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchTagWithAnomalyAlertTag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchTagWithAnomalyAlertTag) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchTagWithAnomalyAlertTag) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSources

`func (o *PatchTagWithAnomalyAlertTag) GetSources() []SourcesIdArrayInner`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *PatchTagWithAnomalyAlertTag) GetSourcesOk() (*[]SourcesIdArrayInner, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *PatchTagWithAnomalyAlertTag) SetSources(v []SourcesIdArrayInner)`

SetSources sets Sources field to given value.

### HasSources

`func (o *PatchTagWithAnomalyAlertTag) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetType

`func (o *PatchTagWithAnomalyAlertTag) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchTagWithAnomalyAlertTag) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchTagWithAnomalyAlertTag) SetType(v string)`

SetType sets Type field to given value.


### GetSubType

`func (o *PatchTagWithAnomalyAlertTag) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *PatchTagWithAnomalyAlertTag) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *PatchTagWithAnomalyAlertTag) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *PatchTagWithAnomalyAlertTag) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetDescription

`func (o *PatchTagWithAnomalyAlertTag) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchTagWithAnomalyAlertTag) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchTagWithAnomalyAlertTag) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchTagWithAnomalyAlertTag) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActions

`func (o *PatchTagWithAnomalyAlertTag) GetActions() []CreateTagActionAnomaly`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *PatchTagWithAnomalyAlertTag) GetActionsOk() (*[]CreateTagActionAnomaly, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *PatchTagWithAnomalyAlertTag) SetActions(v []CreateTagActionAnomaly)`

SetActions sets Actions field to given value.

### HasActions

`func (o *PatchTagWithAnomalyAlertTag) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetUserData

`func (o *PatchTagWithAnomalyAlertTag) GetUserData() map[string]interface{}`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *PatchTagWithAnomalyAlertTag) GetUserDataOk() (*map[string]interface{}, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *PatchTagWithAnomalyAlertTag) SetUserData(v map[string]interface{})`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *PatchTagWithAnomalyAlertTag) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### GetScheduledQuery

`func (o *PatchTagWithAnomalyAlertTag) GetScheduledQuery() CreateScheduledQuery`

GetScheduledQuery returns the ScheduledQuery field if non-nil, zero value otherwise.

### GetScheduledQueryOk

`func (o *PatchTagWithAnomalyAlertTag) GetScheduledQueryOk() (*CreateScheduledQuery, bool)`

GetScheduledQueryOk returns a tuple with the ScheduledQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledQuery

`func (o *PatchTagWithAnomalyAlertTag) SetScheduledQuery(v CreateScheduledQuery)`

SetScheduledQuery sets ScheduledQuery field to given value.

### HasScheduledQuery

`func (o *PatchTagWithAnomalyAlertTag) HasScheduledQuery() bool`

HasScheduledQuery returns a boolean if a field has been set.

### GetThresholdValue

`func (o *PatchTagWithAnomalyAlertTag) GetThresholdValue() int32`

GetThresholdValue returns the ThresholdValue field if non-nil, zero value otherwise.

### GetThresholdValueOk

`func (o *PatchTagWithAnomalyAlertTag) GetThresholdValueOk() (*int32, bool)`

GetThresholdValueOk returns a tuple with the ThresholdValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdValue

`func (o *PatchTagWithAnomalyAlertTag) SetThresholdValue(v int32)`

SetThresholdValue sets ThresholdValue field to given value.

### HasThresholdValue

`func (o *PatchTagWithAnomalyAlertTag) HasThresholdValue() bool`

HasThresholdValue returns a boolean if a field has been set.

### GetPriority

`func (o *PatchTagWithAnomalyAlertTag) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *PatchTagWithAnomalyAlertTag) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *PatchTagWithAnomalyAlertTag) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *PatchTagWithAnomalyAlertTag) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


