# TagAnomalyAlertResponse

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
**UserData** | Pointer to **map[string]interface{}** | A list of key-value pairs that may indicate some auxiliary information about the change detection rule. | [optional] 
**ScheduledQueryId** | Pointer to **string** | ID of the scheduled query associated with this change detection rule. | [optional] 
**SubType** | Pointer to **string** | Always set to \&quot;AnomalyAlert\&quot;. | [optional] 
**ScheduledQuery** | Pointer to [**ScheduledQueryResponse**](ScheduledQueryResponse.md) |  | [optional] 
**ThresholdValue** | Pointer to **int32** | The percentage that defines when to trigger notifications. The percentage can be positive or negative. For example, if the percentage is set to -50, then notifications will be triggered when the query result for the time range defined by the &#x60;time_value&#x60; and &#x60;time_period&#x60; parameters is 50% smaller when compared to the same query over the time range defined by the &#x60;timeframe_value&#x60; and &#x60;timeframe_period&#x60; parameters. | [optional] 
**Priority** | Pointer to **int32** | This ensures investigations are ordered by priority in Investigation Management. Must be &gt;&#x3D;0 | [optional] 

## Methods

### NewTagAnomalyAlertResponse

`func NewTagAnomalyAlertResponse() *TagAnomalyAlertResponse`

NewTagAnomalyAlertResponse instantiates a new TagAnomalyAlertResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagAnomalyAlertResponseWithDefaults

`func NewTagAnomalyAlertResponseWithDefaults() *TagAnomalyAlertResponse`

NewTagAnomalyAlertResponseWithDefaults instantiates a new TagAnomalyAlertResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TagAnomalyAlertResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TagAnomalyAlertResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TagAnomalyAlertResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TagAnomalyAlertResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *TagAnomalyAlertResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TagAnomalyAlertResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TagAnomalyAlertResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TagAnomalyAlertResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TagAnomalyAlertResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TagAnomalyAlertResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TagAnomalyAlertResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TagAnomalyAlertResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *TagAnomalyAlertResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TagAnomalyAlertResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TagAnomalyAlertResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TagAnomalyAlertResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSources

`func (o *TagAnomalyAlertResponse) GetSources() []TagSourceResponse`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *TagAnomalyAlertResponse) GetSourcesOk() (*[]TagSourceResponse, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *TagAnomalyAlertResponse) SetSources(v []TagSourceResponse)`

SetSources sets Sources field to given value.

### HasSources

`func (o *TagAnomalyAlertResponse) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetActions

`func (o *TagAnomalyAlertResponse) GetActions() []ActionResponseAnomaly`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *TagAnomalyAlertResponse) GetActionsOk() (*[]ActionResponseAnomaly, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *TagAnomalyAlertResponse) SetActions(v []ActionResponseAnomaly)`

SetActions sets Actions field to given value.

### HasActions

`func (o *TagAnomalyAlertResponse) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetPatterns

`func (o *TagAnomalyAlertResponse) GetPatterns() []string`

GetPatterns returns the Patterns field if non-nil, zero value otherwise.

### GetPatternsOk

`func (o *TagAnomalyAlertResponse) GetPatternsOk() (*[]string, bool)`

GetPatternsOk returns a tuple with the Patterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatterns

`func (o *TagAnomalyAlertResponse) SetPatterns(v []string)`

SetPatterns sets Patterns field to given value.

### HasPatterns

`func (o *TagAnomalyAlertResponse) HasPatterns() bool`

HasPatterns returns a boolean if a field has been set.

### GetUserData

`func (o *TagAnomalyAlertResponse) GetUserData() map[string]interface{}`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *TagAnomalyAlertResponse) GetUserDataOk() (*map[string]interface{}, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *TagAnomalyAlertResponse) SetUserData(v map[string]interface{})`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *TagAnomalyAlertResponse) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### GetScheduledQueryId

`func (o *TagAnomalyAlertResponse) GetScheduledQueryId() string`

GetScheduledQueryId returns the ScheduledQueryId field if non-nil, zero value otherwise.

### GetScheduledQueryIdOk

`func (o *TagAnomalyAlertResponse) GetScheduledQueryIdOk() (*string, bool)`

GetScheduledQueryIdOk returns a tuple with the ScheduledQueryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledQueryId

`func (o *TagAnomalyAlertResponse) SetScheduledQueryId(v string)`

SetScheduledQueryId sets ScheduledQueryId field to given value.

### HasScheduledQueryId

`func (o *TagAnomalyAlertResponse) HasScheduledQueryId() bool`

HasScheduledQueryId returns a boolean if a field has been set.

### GetSubType

`func (o *TagAnomalyAlertResponse) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *TagAnomalyAlertResponse) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *TagAnomalyAlertResponse) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *TagAnomalyAlertResponse) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetScheduledQuery

`func (o *TagAnomalyAlertResponse) GetScheduledQuery() ScheduledQueryResponse`

GetScheduledQuery returns the ScheduledQuery field if non-nil, zero value otherwise.

### GetScheduledQueryOk

`func (o *TagAnomalyAlertResponse) GetScheduledQueryOk() (*ScheduledQueryResponse, bool)`

GetScheduledQueryOk returns a tuple with the ScheduledQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledQuery

`func (o *TagAnomalyAlertResponse) SetScheduledQuery(v ScheduledQueryResponse)`

SetScheduledQuery sets ScheduledQuery field to given value.

### HasScheduledQuery

`func (o *TagAnomalyAlertResponse) HasScheduledQuery() bool`

HasScheduledQuery returns a boolean if a field has been set.

### GetThresholdValue

`func (o *TagAnomalyAlertResponse) GetThresholdValue() int32`

GetThresholdValue returns the ThresholdValue field if non-nil, zero value otherwise.

### GetThresholdValueOk

`func (o *TagAnomalyAlertResponse) GetThresholdValueOk() (*int32, bool)`

GetThresholdValueOk returns a tuple with the ThresholdValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdValue

`func (o *TagAnomalyAlertResponse) SetThresholdValue(v int32)`

SetThresholdValue sets ThresholdValue field to given value.

### HasThresholdValue

`func (o *TagAnomalyAlertResponse) HasThresholdValue() bool`

HasThresholdValue returns a boolean if a field has been set.

### GetPriority

`func (o *TagAnomalyAlertResponse) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TagAnomalyAlertResponse) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TagAnomalyAlertResponse) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *TagAnomalyAlertResponse) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


