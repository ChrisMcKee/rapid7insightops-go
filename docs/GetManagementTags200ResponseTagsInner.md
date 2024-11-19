# GetManagementTags200ResponseTagsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Always set to \&quot;AlertNotify\&quot;. | [optional] 
**Id** | Pointer to **string** | The unique ID of the detection rule. | [optional] 
**Name** | Pointer to **string** | The name of the detection rule. | [optional] 
**Description** | Pointer to **string** | The description of the detection rule. | [optional] 
**Sources** | Pointer to [**[]TagSourceResponse**](TagSourceResponse.md) | The collection of logs the detection rule operates on. | [optional] 
**Actions** | Pointer to [**[]ActionResponseAnomaly**](ActionResponseAnomaly.md) | The notifications attached to the detection rule. | [optional] 
**Patterns** | Pointer to **[]string** | Use the &#x60;leql&#x60; parameter instead of this parameter. | [optional] 
**Leql** | Pointer to [**InactivityLeql**](InactivityLeql.md) |  | [optional] 
**Priority** | Pointer to **int32** | This ensures investigations are ordered by priority in Investigation Management. Must be &gt;&#x3D;0 | [optional] 
**UserData** | Pointer to **map[string]interface{}** | A list of key-value pairs that may indicate some auxiliary information about the change detection rule. | [optional] 
**Labels** | Pointer to [**[]LabelResponse**](LabelResponse.md) | The labels attached to the detection rule. | [optional] 
**TimeframeValue** | Pointer to **int32** | Defines the duration of inactivity before an alert triggers along with the &#x60;timeframe_period&#x60; parameter. | [optional] 
**TimeframePeriod** | Pointer to [**TimeframePeriod**](TimeframePeriod.md) |  | [optional] 
**SubType** | Pointer to **string** | Always set to \&quot;AnomalyAlert\&quot;. | [optional] 
**ScheduledQueryId** | Pointer to **string** | ID of the scheduled query associated with this change detection rule. | [optional] 
**ScheduledQuery** | Pointer to [**ScheduledQueryResponse**](ScheduledQueryResponse.md) |  | [optional] 
**ThresholdValue** | Pointer to **int32** | The percentage that defines when to trigger notifications. The percentage can be positive or negative. For example, if the percentage is set to -50, then notifications will be triggered when the query result for the time range defined by the &#x60;time_value&#x60; and &#x60;time_period&#x60; parameters is 50% smaller when compared to the same query over the time range defined by the &#x60;timeframe_value&#x60; and &#x60;timeframe_period&#x60; parameters. | [optional] 

## Methods

### NewGetManagementTags200ResponseTagsInner

`func NewGetManagementTags200ResponseTagsInner() *GetManagementTags200ResponseTagsInner`

NewGetManagementTags200ResponseTagsInner instantiates a new GetManagementTags200ResponseTagsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetManagementTags200ResponseTagsInnerWithDefaults

`func NewGetManagementTags200ResponseTagsInnerWithDefaults() *GetManagementTags200ResponseTagsInner`

NewGetManagementTags200ResponseTagsInnerWithDefaults instantiates a new GetManagementTags200ResponseTagsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GetManagementTags200ResponseTagsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetManagementTags200ResponseTagsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetManagementTags200ResponseTagsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetManagementTags200ResponseTagsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *GetManagementTags200ResponseTagsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetManagementTags200ResponseTagsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetManagementTags200ResponseTagsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetManagementTags200ResponseTagsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetManagementTags200ResponseTagsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetManagementTags200ResponseTagsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetManagementTags200ResponseTagsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetManagementTags200ResponseTagsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetManagementTags200ResponseTagsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetManagementTags200ResponseTagsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetManagementTags200ResponseTagsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetManagementTags200ResponseTagsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSources

`func (o *GetManagementTags200ResponseTagsInner) GetSources() []TagSourceResponse`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *GetManagementTags200ResponseTagsInner) GetSourcesOk() (*[]TagSourceResponse, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *GetManagementTags200ResponseTagsInner) SetSources(v []TagSourceResponse)`

SetSources sets Sources field to given value.

### HasSources

`func (o *GetManagementTags200ResponseTagsInner) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetActions

`func (o *GetManagementTags200ResponseTagsInner) GetActions() []ActionResponseAnomaly`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *GetManagementTags200ResponseTagsInner) GetActionsOk() (*[]ActionResponseAnomaly, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *GetManagementTags200ResponseTagsInner) SetActions(v []ActionResponseAnomaly)`

SetActions sets Actions field to given value.

### HasActions

`func (o *GetManagementTags200ResponseTagsInner) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetPatterns

`func (o *GetManagementTags200ResponseTagsInner) GetPatterns() []string`

GetPatterns returns the Patterns field if non-nil, zero value otherwise.

### GetPatternsOk

`func (o *GetManagementTags200ResponseTagsInner) GetPatternsOk() (*[]string, bool)`

GetPatternsOk returns a tuple with the Patterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatterns

`func (o *GetManagementTags200ResponseTagsInner) SetPatterns(v []string)`

SetPatterns sets Patterns field to given value.

### HasPatterns

`func (o *GetManagementTags200ResponseTagsInner) HasPatterns() bool`

HasPatterns returns a boolean if a field has been set.

### GetLeql

`func (o *GetManagementTags200ResponseTagsInner) GetLeql() InactivityLeql`

GetLeql returns the Leql field if non-nil, zero value otherwise.

### GetLeqlOk

`func (o *GetManagementTags200ResponseTagsInner) GetLeqlOk() (*InactivityLeql, bool)`

GetLeqlOk returns a tuple with the Leql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeql

`func (o *GetManagementTags200ResponseTagsInner) SetLeql(v InactivityLeql)`

SetLeql sets Leql field to given value.

### HasLeql

`func (o *GetManagementTags200ResponseTagsInner) HasLeql() bool`

HasLeql returns a boolean if a field has been set.

### GetPriority

`func (o *GetManagementTags200ResponseTagsInner) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *GetManagementTags200ResponseTagsInner) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *GetManagementTags200ResponseTagsInner) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *GetManagementTags200ResponseTagsInner) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetUserData

`func (o *GetManagementTags200ResponseTagsInner) GetUserData() map[string]interface{}`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *GetManagementTags200ResponseTagsInner) GetUserDataOk() (*map[string]interface{}, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *GetManagementTags200ResponseTagsInner) SetUserData(v map[string]interface{})`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *GetManagementTags200ResponseTagsInner) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### GetLabels

`func (o *GetManagementTags200ResponseTagsInner) GetLabels() []LabelResponse`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetManagementTags200ResponseTagsInner) GetLabelsOk() (*[]LabelResponse, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetManagementTags200ResponseTagsInner) SetLabels(v []LabelResponse)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GetManagementTags200ResponseTagsInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetTimeframeValue

`func (o *GetManagementTags200ResponseTagsInner) GetTimeframeValue() int32`

GetTimeframeValue returns the TimeframeValue field if non-nil, zero value otherwise.

### GetTimeframeValueOk

`func (o *GetManagementTags200ResponseTagsInner) GetTimeframeValueOk() (*int32, bool)`

GetTimeframeValueOk returns a tuple with the TimeframeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframeValue

`func (o *GetManagementTags200ResponseTagsInner) SetTimeframeValue(v int32)`

SetTimeframeValue sets TimeframeValue field to given value.

### HasTimeframeValue

`func (o *GetManagementTags200ResponseTagsInner) HasTimeframeValue() bool`

HasTimeframeValue returns a boolean if a field has been set.

### GetTimeframePeriod

`func (o *GetManagementTags200ResponseTagsInner) GetTimeframePeriod() TimeframePeriod`

GetTimeframePeriod returns the TimeframePeriod field if non-nil, zero value otherwise.

### GetTimeframePeriodOk

`func (o *GetManagementTags200ResponseTagsInner) GetTimeframePeriodOk() (*TimeframePeriod, bool)`

GetTimeframePeriodOk returns a tuple with the TimeframePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframePeriod

`func (o *GetManagementTags200ResponseTagsInner) SetTimeframePeriod(v TimeframePeriod)`

SetTimeframePeriod sets TimeframePeriod field to given value.

### HasTimeframePeriod

`func (o *GetManagementTags200ResponseTagsInner) HasTimeframePeriod() bool`

HasTimeframePeriod returns a boolean if a field has been set.

### GetSubType

`func (o *GetManagementTags200ResponseTagsInner) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *GetManagementTags200ResponseTagsInner) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *GetManagementTags200ResponseTagsInner) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *GetManagementTags200ResponseTagsInner) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetScheduledQueryId

`func (o *GetManagementTags200ResponseTagsInner) GetScheduledQueryId() string`

GetScheduledQueryId returns the ScheduledQueryId field if non-nil, zero value otherwise.

### GetScheduledQueryIdOk

`func (o *GetManagementTags200ResponseTagsInner) GetScheduledQueryIdOk() (*string, bool)`

GetScheduledQueryIdOk returns a tuple with the ScheduledQueryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledQueryId

`func (o *GetManagementTags200ResponseTagsInner) SetScheduledQueryId(v string)`

SetScheduledQueryId sets ScheduledQueryId field to given value.

### HasScheduledQueryId

`func (o *GetManagementTags200ResponseTagsInner) HasScheduledQueryId() bool`

HasScheduledQueryId returns a boolean if a field has been set.

### GetScheduledQuery

`func (o *GetManagementTags200ResponseTagsInner) GetScheduledQuery() ScheduledQueryResponse`

GetScheduledQuery returns the ScheduledQuery field if non-nil, zero value otherwise.

### GetScheduledQueryOk

`func (o *GetManagementTags200ResponseTagsInner) GetScheduledQueryOk() (*ScheduledQueryResponse, bool)`

GetScheduledQueryOk returns a tuple with the ScheduledQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledQuery

`func (o *GetManagementTags200ResponseTagsInner) SetScheduledQuery(v ScheduledQueryResponse)`

SetScheduledQuery sets ScheduledQuery field to given value.

### HasScheduledQuery

`func (o *GetManagementTags200ResponseTagsInner) HasScheduledQuery() bool`

HasScheduledQuery returns a boolean if a field has been set.

### GetThresholdValue

`func (o *GetManagementTags200ResponseTagsInner) GetThresholdValue() int32`

GetThresholdValue returns the ThresholdValue field if non-nil, zero value otherwise.

### GetThresholdValueOk

`func (o *GetManagementTags200ResponseTagsInner) GetThresholdValueOk() (*int32, bool)`

GetThresholdValueOk returns a tuple with the ThresholdValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdValue

`func (o *GetManagementTags200ResponseTagsInner) SetThresholdValue(v int32)`

SetThresholdValue sets ThresholdValue field to given value.

### HasThresholdValue

`func (o *GetManagementTags200ResponseTagsInner) HasThresholdValue() bool`

HasThresholdValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


