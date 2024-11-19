# PatchBasicTagTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the detection rule. | [optional] 
**Patterns** | Pointer to **[]string** | Use the &#x60;leql&#x60; parameter instead of this parameter. | [optional] 
**Leql** | Pointer to [**BasicLeql**](BasicLeql.md) |  | [optional] 
**Sources** | Pointer to [**[]SourcesIdArrayInner**](SourcesIdArrayInner.md) | The IDs of the logs that the detection rule operates on. | [optional] 
**Type** | Pointer to **string** | Always set to \&quot;Alert\&quot;. | [optional] 
**Description** | Pointer to **string** | The description of the detection rule. | [optional] 
**Actions** | Pointer to [**[]CreateTagAction**](CreateTagAction.md) | The notifications attached to the detection rule. | [optional] 
**Labels** | Pointer to [**[]LabelResponse**](LabelResponse.md) | The labels attached to the detection rule. | [optional] 
**UserData** | Pointer to **map[string]interface{}** | A list of key-value pairs that may indicate some auxiliary information about the change detection rule. | [optional] 
**Priority** | Pointer to **int32** | This ensures investigations are ordered by priority in Investigation Management. Must be &gt;&#x3D;0 | [optional] 

## Methods

### NewPatchBasicTagTag

`func NewPatchBasicTagTag() *PatchBasicTagTag`

NewPatchBasicTagTag instantiates a new PatchBasicTagTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchBasicTagTagWithDefaults

`func NewPatchBasicTagTagWithDefaults() *PatchBasicTagTag`

NewPatchBasicTagTagWithDefaults instantiates a new PatchBasicTagTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchBasicTagTag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchBasicTagTag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchBasicTagTag) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchBasicTagTag) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPatterns

`func (o *PatchBasicTagTag) GetPatterns() []string`

GetPatterns returns the Patterns field if non-nil, zero value otherwise.

### GetPatternsOk

`func (o *PatchBasicTagTag) GetPatternsOk() (*[]string, bool)`

GetPatternsOk returns a tuple with the Patterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatterns

`func (o *PatchBasicTagTag) SetPatterns(v []string)`

SetPatterns sets Patterns field to given value.

### HasPatterns

`func (o *PatchBasicTagTag) HasPatterns() bool`

HasPatterns returns a boolean if a field has been set.

### GetLeql

`func (o *PatchBasicTagTag) GetLeql() BasicLeql`

GetLeql returns the Leql field if non-nil, zero value otherwise.

### GetLeqlOk

`func (o *PatchBasicTagTag) GetLeqlOk() (*BasicLeql, bool)`

GetLeqlOk returns a tuple with the Leql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeql

`func (o *PatchBasicTagTag) SetLeql(v BasicLeql)`

SetLeql sets Leql field to given value.

### HasLeql

`func (o *PatchBasicTagTag) HasLeql() bool`

HasLeql returns a boolean if a field has been set.

### GetSources

`func (o *PatchBasicTagTag) GetSources() []SourcesIdArrayInner`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *PatchBasicTagTag) GetSourcesOk() (*[]SourcesIdArrayInner, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *PatchBasicTagTag) SetSources(v []SourcesIdArrayInner)`

SetSources sets Sources field to given value.

### HasSources

`func (o *PatchBasicTagTag) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetType

`func (o *PatchBasicTagTag) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchBasicTagTag) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchBasicTagTag) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PatchBasicTagTag) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *PatchBasicTagTag) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchBasicTagTag) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchBasicTagTag) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchBasicTagTag) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActions

`func (o *PatchBasicTagTag) GetActions() []CreateTagAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *PatchBasicTagTag) GetActionsOk() (*[]CreateTagAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *PatchBasicTagTag) SetActions(v []CreateTagAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *PatchBasicTagTag) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetLabels

`func (o *PatchBasicTagTag) GetLabels() []LabelResponse`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *PatchBasicTagTag) GetLabelsOk() (*[]LabelResponse, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *PatchBasicTagTag) SetLabels(v []LabelResponse)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *PatchBasicTagTag) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetUserData

`func (o *PatchBasicTagTag) GetUserData() map[string]interface{}`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *PatchBasicTagTag) GetUserDataOk() (*map[string]interface{}, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *PatchBasicTagTag) SetUserData(v map[string]interface{})`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *PatchBasicTagTag) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### GetPriority

`func (o *PatchBasicTagTag) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *PatchBasicTagTag) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *PatchBasicTagTag) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *PatchBasicTagTag) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


