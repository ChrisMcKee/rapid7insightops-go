# TagInactivityAlertResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Always set to \&quot;AlertNotify\&quot;. | [optional] 
**Id** | Pointer to **string** | The unique ID of the detection rule. | [optional] 
**Name** | Pointer to **string** | The name of the detection rule. | [optional] 
**Description** | Pointer to **string** | The description of the detection rule. | [optional] 
**Sources** | Pointer to [**[]TagSourceResponse**](TagSourceResponse.md) | The collection of logs the detection rule operates on. | [optional] 
**Actions** | Pointer to [**[]ActionResponseInactivity**](ActionResponseInactivity.md) | The notifications attached to the detection rule. | [optional] 
**Patterns** | Pointer to **[]string** | Use the &#x60;leql&#x60; parameter instead of this parameter. | [optional] 
**Leql** | Pointer to [**InactivityLeql**](InactivityLeql.md) |  | [optional] 
**UserData** | Pointer to **map[string]interface{}** | A list of key-value pairs that may indicate some auxiliary information about the change detection rule. | [optional] 
**TimeframeValue** | Pointer to **int32** | Defines the duration of inactivity before an alert triggers along with the &#x60;timeframe_period&#x60; parameter. | [optional] 
**TimeframePeriod** | Pointer to [**TimeframePeriod**](TimeframePeriod.md) |  | [optional] 
**Priority** | Pointer to **int32** | This ensures investigations are ordered by priority in Investigation Management. Must be &gt;&#x3D;0 | [optional] 
**SubType** | Pointer to **string** | Always set to \&quot;InactivityAlert\&quot;. | [optional] 

## Methods

### NewTagInactivityAlertResponse

`func NewTagInactivityAlertResponse() *TagInactivityAlertResponse`

NewTagInactivityAlertResponse instantiates a new TagInactivityAlertResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagInactivityAlertResponseWithDefaults

`func NewTagInactivityAlertResponseWithDefaults() *TagInactivityAlertResponse`

NewTagInactivityAlertResponseWithDefaults instantiates a new TagInactivityAlertResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TagInactivityAlertResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TagInactivityAlertResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TagInactivityAlertResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TagInactivityAlertResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *TagInactivityAlertResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TagInactivityAlertResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TagInactivityAlertResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TagInactivityAlertResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TagInactivityAlertResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TagInactivityAlertResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TagInactivityAlertResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TagInactivityAlertResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *TagInactivityAlertResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TagInactivityAlertResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TagInactivityAlertResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TagInactivityAlertResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSources

`func (o *TagInactivityAlertResponse) GetSources() []TagSourceResponse`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *TagInactivityAlertResponse) GetSourcesOk() (*[]TagSourceResponse, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *TagInactivityAlertResponse) SetSources(v []TagSourceResponse)`

SetSources sets Sources field to given value.

### HasSources

`func (o *TagInactivityAlertResponse) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetActions

`func (o *TagInactivityAlertResponse) GetActions() []ActionResponseInactivity`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *TagInactivityAlertResponse) GetActionsOk() (*[]ActionResponseInactivity, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *TagInactivityAlertResponse) SetActions(v []ActionResponseInactivity)`

SetActions sets Actions field to given value.

### HasActions

`func (o *TagInactivityAlertResponse) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetPatterns

`func (o *TagInactivityAlertResponse) GetPatterns() []string`

GetPatterns returns the Patterns field if non-nil, zero value otherwise.

### GetPatternsOk

`func (o *TagInactivityAlertResponse) GetPatternsOk() (*[]string, bool)`

GetPatternsOk returns a tuple with the Patterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatterns

`func (o *TagInactivityAlertResponse) SetPatterns(v []string)`

SetPatterns sets Patterns field to given value.

### HasPatterns

`func (o *TagInactivityAlertResponse) HasPatterns() bool`

HasPatterns returns a boolean if a field has been set.

### GetLeql

`func (o *TagInactivityAlertResponse) GetLeql() InactivityLeql`

GetLeql returns the Leql field if non-nil, zero value otherwise.

### GetLeqlOk

`func (o *TagInactivityAlertResponse) GetLeqlOk() (*InactivityLeql, bool)`

GetLeqlOk returns a tuple with the Leql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeql

`func (o *TagInactivityAlertResponse) SetLeql(v InactivityLeql)`

SetLeql sets Leql field to given value.

### HasLeql

`func (o *TagInactivityAlertResponse) HasLeql() bool`

HasLeql returns a boolean if a field has been set.

### GetUserData

`func (o *TagInactivityAlertResponse) GetUserData() map[string]interface{}`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *TagInactivityAlertResponse) GetUserDataOk() (*map[string]interface{}, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *TagInactivityAlertResponse) SetUserData(v map[string]interface{})`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *TagInactivityAlertResponse) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### GetTimeframeValue

`func (o *TagInactivityAlertResponse) GetTimeframeValue() int32`

GetTimeframeValue returns the TimeframeValue field if non-nil, zero value otherwise.

### GetTimeframeValueOk

`func (o *TagInactivityAlertResponse) GetTimeframeValueOk() (*int32, bool)`

GetTimeframeValueOk returns a tuple with the TimeframeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframeValue

`func (o *TagInactivityAlertResponse) SetTimeframeValue(v int32)`

SetTimeframeValue sets TimeframeValue field to given value.

### HasTimeframeValue

`func (o *TagInactivityAlertResponse) HasTimeframeValue() bool`

HasTimeframeValue returns a boolean if a field has been set.

### GetTimeframePeriod

`func (o *TagInactivityAlertResponse) GetTimeframePeriod() TimeframePeriod`

GetTimeframePeriod returns the TimeframePeriod field if non-nil, zero value otherwise.

### GetTimeframePeriodOk

`func (o *TagInactivityAlertResponse) GetTimeframePeriodOk() (*TimeframePeriod, bool)`

GetTimeframePeriodOk returns a tuple with the TimeframePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframePeriod

`func (o *TagInactivityAlertResponse) SetTimeframePeriod(v TimeframePeriod)`

SetTimeframePeriod sets TimeframePeriod field to given value.

### HasTimeframePeriod

`func (o *TagInactivityAlertResponse) HasTimeframePeriod() bool`

HasTimeframePeriod returns a boolean if a field has been set.

### GetPriority

`func (o *TagInactivityAlertResponse) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TagInactivityAlertResponse) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TagInactivityAlertResponse) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *TagInactivityAlertResponse) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetSubType

`func (o *TagInactivityAlertResponse) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *TagInactivityAlertResponse) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *TagInactivityAlertResponse) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *TagInactivityAlertResponse) HasSubType() bool`

HasSubType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


