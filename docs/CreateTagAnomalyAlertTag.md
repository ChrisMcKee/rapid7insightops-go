# CreateTagAnomalyAlertTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the detection rule. | 
**Sources** | [**[]SourcesIdArrayInner**](SourcesIdArrayInner.md) | The IDs of the logs that the detection rule operates on. | 
**Type** | **string** | Always set to \&quot;AlertNotify\&quot;. | 
**SubType** | **string** | Always set to \&quot;AnomalyAlert\&quot;. | 
**Description** | Pointer to **string** | The description of the detection rule. | [optional] 
**Actions** | Pointer to [**[]CreateTagActionAnomaly**](CreateTagActionAnomaly.md) | The notifications attached to the detection rule. | [optional] 
**UserData** | Pointer to **map[string]interface{}** | A list of key-value pairs that may indicate some auxiliary information about the change detection rule. | [optional] 
**ScheduledQuery** | [**CreateScheduledQuery**](CreateScheduledQuery.md) |  | 
**ThresholdValue** | **int32** | The percentage that defines when to trigger notifications. The percentage can be positive or negative. For example, if the percentage is set to -50, then notifications will be triggered when the query result for the time range defined by the &#x60;time_value&#x60; and &#x60;time_period&#x60; parameters is 50% smaller when compared to the same query over the time range defined by the &#x60;timeframe_value&#x60; and &#x60;timeframe_period&#x60; parameters. | 
**Priority** | Pointer to **int32** | This ensures investigations are ordered by priority in Investigation Management. Must be &gt;&#x3D;0 | [optional] 

## Methods

### NewCreateTagAnomalyAlertTag

`func NewCreateTagAnomalyAlertTag(name string, sources []SourcesIdArrayInner, type_ string, subType string, scheduledQuery CreateScheduledQuery, thresholdValue int32, ) *CreateTagAnomalyAlertTag`

NewCreateTagAnomalyAlertTag instantiates a new CreateTagAnomalyAlertTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTagAnomalyAlertTagWithDefaults

`func NewCreateTagAnomalyAlertTagWithDefaults() *CreateTagAnomalyAlertTag`

NewCreateTagAnomalyAlertTagWithDefaults instantiates a new CreateTagAnomalyAlertTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateTagAnomalyAlertTag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTagAnomalyAlertTag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTagAnomalyAlertTag) SetName(v string)`

SetName sets Name field to given value.


### GetSources

`func (o *CreateTagAnomalyAlertTag) GetSources() []SourcesIdArrayInner`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *CreateTagAnomalyAlertTag) GetSourcesOk() (*[]SourcesIdArrayInner, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *CreateTagAnomalyAlertTag) SetSources(v []SourcesIdArrayInner)`

SetSources sets Sources field to given value.


### GetType

`func (o *CreateTagAnomalyAlertTag) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateTagAnomalyAlertTag) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateTagAnomalyAlertTag) SetType(v string)`

SetType sets Type field to given value.


### GetSubType

`func (o *CreateTagAnomalyAlertTag) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *CreateTagAnomalyAlertTag) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *CreateTagAnomalyAlertTag) SetSubType(v string)`

SetSubType sets SubType field to given value.


### GetDescription

`func (o *CreateTagAnomalyAlertTag) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateTagAnomalyAlertTag) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateTagAnomalyAlertTag) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateTagAnomalyAlertTag) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActions

`func (o *CreateTagAnomalyAlertTag) GetActions() []CreateTagActionAnomaly`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *CreateTagAnomalyAlertTag) GetActionsOk() (*[]CreateTagActionAnomaly, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *CreateTagAnomalyAlertTag) SetActions(v []CreateTagActionAnomaly)`

SetActions sets Actions field to given value.

### HasActions

`func (o *CreateTagAnomalyAlertTag) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetUserData

`func (o *CreateTagAnomalyAlertTag) GetUserData() map[string]interface{}`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *CreateTagAnomalyAlertTag) GetUserDataOk() (*map[string]interface{}, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *CreateTagAnomalyAlertTag) SetUserData(v map[string]interface{})`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *CreateTagAnomalyAlertTag) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### GetScheduledQuery

`func (o *CreateTagAnomalyAlertTag) GetScheduledQuery() CreateScheduledQuery`

GetScheduledQuery returns the ScheduledQuery field if non-nil, zero value otherwise.

### GetScheduledQueryOk

`func (o *CreateTagAnomalyAlertTag) GetScheduledQueryOk() (*CreateScheduledQuery, bool)`

GetScheduledQueryOk returns a tuple with the ScheduledQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledQuery

`func (o *CreateTagAnomalyAlertTag) SetScheduledQuery(v CreateScheduledQuery)`

SetScheduledQuery sets ScheduledQuery field to given value.


### GetThresholdValue

`func (o *CreateTagAnomalyAlertTag) GetThresholdValue() int32`

GetThresholdValue returns the ThresholdValue field if non-nil, zero value otherwise.

### GetThresholdValueOk

`func (o *CreateTagAnomalyAlertTag) GetThresholdValueOk() (*int32, bool)`

GetThresholdValueOk returns a tuple with the ThresholdValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdValue

`func (o *CreateTagAnomalyAlertTag) SetThresholdValue(v int32)`

SetThresholdValue sets ThresholdValue field to given value.


### GetPriority

`func (o *CreateTagAnomalyAlertTag) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CreateTagAnomalyAlertTag) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CreateTagAnomalyAlertTag) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CreateTagAnomalyAlertTag) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


