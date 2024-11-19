# PutTagAnomalyAlertTag

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
**ScheduledQueryId** | Pointer to **string** | ID of the scheduled query associated with this change detection rule. | [optional] 
**ScheduledQuery** | [**ScheduledQueryResponse**](ScheduledQueryResponse.md) |  | 
**ThresholdValue** | **int32** | The percentage that defines when to trigger notifications. The percentage can be positive or negative. For example, if the percentage is set to -50, then notifications will be triggered when the query result for the time range defined by the &#x60;time_value&#x60; and &#x60;time_period&#x60; parameters is 50% smaller when compared to the same query over the time range defined by the &#x60;timeframe_value&#x60; and &#x60;timeframe_period&#x60; parameters. | 
**Priority** | Pointer to **int32** | This ensures investigations are ordered by priority in Investigation Management. Must be &gt;&#x3D;0 | [optional] 

## Methods

### NewPutTagAnomalyAlertTag

`func NewPutTagAnomalyAlertTag(name string, sources []SourcesIdArrayInner, type_ string, subType string, scheduledQuery ScheduledQueryResponse, thresholdValue int32, ) *PutTagAnomalyAlertTag`

NewPutTagAnomalyAlertTag instantiates a new PutTagAnomalyAlertTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutTagAnomalyAlertTagWithDefaults

`func NewPutTagAnomalyAlertTagWithDefaults() *PutTagAnomalyAlertTag`

NewPutTagAnomalyAlertTagWithDefaults instantiates a new PutTagAnomalyAlertTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PutTagAnomalyAlertTag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PutTagAnomalyAlertTag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PutTagAnomalyAlertTag) SetName(v string)`

SetName sets Name field to given value.


### GetSources

`func (o *PutTagAnomalyAlertTag) GetSources() []SourcesIdArrayInner`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *PutTagAnomalyAlertTag) GetSourcesOk() (*[]SourcesIdArrayInner, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *PutTagAnomalyAlertTag) SetSources(v []SourcesIdArrayInner)`

SetSources sets Sources field to given value.


### GetType

`func (o *PutTagAnomalyAlertTag) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PutTagAnomalyAlertTag) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PutTagAnomalyAlertTag) SetType(v string)`

SetType sets Type field to given value.


### GetSubType

`func (o *PutTagAnomalyAlertTag) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *PutTagAnomalyAlertTag) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *PutTagAnomalyAlertTag) SetSubType(v string)`

SetSubType sets SubType field to given value.


### GetDescription

`func (o *PutTagAnomalyAlertTag) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PutTagAnomalyAlertTag) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PutTagAnomalyAlertTag) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PutTagAnomalyAlertTag) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActions

`func (o *PutTagAnomalyAlertTag) GetActions() []CreateTagActionAnomaly`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *PutTagAnomalyAlertTag) GetActionsOk() (*[]CreateTagActionAnomaly, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *PutTagAnomalyAlertTag) SetActions(v []CreateTagActionAnomaly)`

SetActions sets Actions field to given value.

### HasActions

`func (o *PutTagAnomalyAlertTag) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetUserData

`func (o *PutTagAnomalyAlertTag) GetUserData() map[string]interface{}`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *PutTagAnomalyAlertTag) GetUserDataOk() (*map[string]interface{}, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *PutTagAnomalyAlertTag) SetUserData(v map[string]interface{})`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *PutTagAnomalyAlertTag) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### GetScheduledQueryId

`func (o *PutTagAnomalyAlertTag) GetScheduledQueryId() string`

GetScheduledQueryId returns the ScheduledQueryId field if non-nil, zero value otherwise.

### GetScheduledQueryIdOk

`func (o *PutTagAnomalyAlertTag) GetScheduledQueryIdOk() (*string, bool)`

GetScheduledQueryIdOk returns a tuple with the ScheduledQueryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledQueryId

`func (o *PutTagAnomalyAlertTag) SetScheduledQueryId(v string)`

SetScheduledQueryId sets ScheduledQueryId field to given value.

### HasScheduledQueryId

`func (o *PutTagAnomalyAlertTag) HasScheduledQueryId() bool`

HasScheduledQueryId returns a boolean if a field has been set.

### GetScheduledQuery

`func (o *PutTagAnomalyAlertTag) GetScheduledQuery() ScheduledQueryResponse`

GetScheduledQuery returns the ScheduledQuery field if non-nil, zero value otherwise.

### GetScheduledQueryOk

`func (o *PutTagAnomalyAlertTag) GetScheduledQueryOk() (*ScheduledQueryResponse, bool)`

GetScheduledQueryOk returns a tuple with the ScheduledQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledQuery

`func (o *PutTagAnomalyAlertTag) SetScheduledQuery(v ScheduledQueryResponse)`

SetScheduledQuery sets ScheduledQuery field to given value.


### GetThresholdValue

`func (o *PutTagAnomalyAlertTag) GetThresholdValue() int32`

GetThresholdValue returns the ThresholdValue field if non-nil, zero value otherwise.

### GetThresholdValueOk

`func (o *PutTagAnomalyAlertTag) GetThresholdValueOk() (*int32, bool)`

GetThresholdValueOk returns a tuple with the ThresholdValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdValue

`func (o *PutTagAnomalyAlertTag) SetThresholdValue(v int32)`

SetThresholdValue sets ThresholdValue field to given value.


### GetPriority

`func (o *PutTagAnomalyAlertTag) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *PutTagAnomalyAlertTag) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *PutTagAnomalyAlertTag) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *PutTagAnomalyAlertTag) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


