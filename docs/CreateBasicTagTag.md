# CreateBasicTagTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the detection rule. | 
**Patterns** | Pointer to **[]string** | Use the &#x60;leql&#x60; parameter instead of this parameter. | [optional] 
**Sources** | [**[]SourcesIdArrayInner**](SourcesIdArrayInner.md) | The IDs of the logs that the detection rule operates on. | 
**Type** | Pointer to **string** | Always set to \&quot;Alert\&quot;. | [optional] 
**Description** | Pointer to **string** | The description of the detection rule. | [optional] 
**Actions** | Pointer to [**[]CreateTagAction**](CreateTagAction.md) | The notifications attached to the detection rule. | [optional] 
**Labels** | Pointer to [**[]LabelResponse**](LabelResponse.md) | The labels attached to the detection rule. | [optional] 
**UserData** | Pointer to **map[string]interface{}** | A list of key-value pairs that may indicate some auxiliary information about the change detection rule. | [optional] 
**Leql** | [**BasicLeql**](BasicLeql.md) |  | 
**Priority** | Pointer to **int32** | This ensures investigations are ordered by priority in Investigation Management. Must be &gt;&#x3D;0 | [optional] 

## Methods

### NewCreateBasicTagTag

`func NewCreateBasicTagTag(name string, sources []SourcesIdArrayInner, leql BasicLeql, ) *CreateBasicTagTag`

NewCreateBasicTagTag instantiates a new CreateBasicTagTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBasicTagTagWithDefaults

`func NewCreateBasicTagTagWithDefaults() *CreateBasicTagTag`

NewCreateBasicTagTagWithDefaults instantiates a new CreateBasicTagTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateBasicTagTag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateBasicTagTag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateBasicTagTag) SetName(v string)`

SetName sets Name field to given value.


### GetPatterns

`func (o *CreateBasicTagTag) GetPatterns() []string`

GetPatterns returns the Patterns field if non-nil, zero value otherwise.

### GetPatternsOk

`func (o *CreateBasicTagTag) GetPatternsOk() (*[]string, bool)`

GetPatternsOk returns a tuple with the Patterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatterns

`func (o *CreateBasicTagTag) SetPatterns(v []string)`

SetPatterns sets Patterns field to given value.

### HasPatterns

`func (o *CreateBasicTagTag) HasPatterns() bool`

HasPatterns returns a boolean if a field has been set.

### GetSources

`func (o *CreateBasicTagTag) GetSources() []SourcesIdArrayInner`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *CreateBasicTagTag) GetSourcesOk() (*[]SourcesIdArrayInner, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *CreateBasicTagTag) SetSources(v []SourcesIdArrayInner)`

SetSources sets Sources field to given value.


### GetType

`func (o *CreateBasicTagTag) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateBasicTagTag) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateBasicTagTag) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateBasicTagTag) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *CreateBasicTagTag) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateBasicTagTag) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateBasicTagTag) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateBasicTagTag) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActions

`func (o *CreateBasicTagTag) GetActions() []CreateTagAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *CreateBasicTagTag) GetActionsOk() (*[]CreateTagAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *CreateBasicTagTag) SetActions(v []CreateTagAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *CreateBasicTagTag) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetLabels

`func (o *CreateBasicTagTag) GetLabels() []LabelResponse`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CreateBasicTagTag) GetLabelsOk() (*[]LabelResponse, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CreateBasicTagTag) SetLabels(v []LabelResponse)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CreateBasicTagTag) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetUserData

`func (o *CreateBasicTagTag) GetUserData() map[string]interface{}`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *CreateBasicTagTag) GetUserDataOk() (*map[string]interface{}, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *CreateBasicTagTag) SetUserData(v map[string]interface{})`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *CreateBasicTagTag) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### GetLeql

`func (o *CreateBasicTagTag) GetLeql() BasicLeql`

GetLeql returns the Leql field if non-nil, zero value otherwise.

### GetLeqlOk

`func (o *CreateBasicTagTag) GetLeqlOk() (*BasicLeql, bool)`

GetLeqlOk returns a tuple with the Leql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeql

`func (o *CreateBasicTagTag) SetLeql(v BasicLeql)`

SetLeql sets Leql field to given value.


### GetPriority

`func (o *CreateBasicTagTag) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CreateBasicTagTag) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CreateBasicTagTag) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CreateBasicTagTag) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


