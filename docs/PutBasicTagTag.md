# PutBasicTagTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the detection rule. | 
**Patterns** | Pointer to **[]string** | Use the &#x60;leql&#x60; parameter instead of this parameter. | [optional] 
**Leql** | [**BasicLeql**](BasicLeql.md) |  | 
**Sources** | [**[]SourcesIdArrayInner**](SourcesIdArrayInner.md) | The IDs of the logs that the detection rule operates on. | 
**Type** | Pointer to **string** | Always set to \&quot;Alert\&quot;. | [optional] 
**Description** | Pointer to **string** | The description of the detection rule. | [optional] 
**Actions** | Pointer to [**[]CreateTagAction**](CreateTagAction.md) | The notifications attached to the detection rule. | [optional] 
**Labels** | Pointer to [**[]LabelResponse**](LabelResponse.md) | The labels attached to the detection rule. | [optional] 
**UserData** | Pointer to **map[string]interface{}** | A list of key-value pairs that may indicate some auxiliary information about the change detection rule. | [optional] 
**Priority** | Pointer to **int32** | This ensures investigations are ordered by priority in Investigation Management. Must be &gt;&#x3D;0 | [optional] 

## Methods

### NewPutBasicTagTag

`func NewPutBasicTagTag(name string, leql BasicLeql, sources []SourcesIdArrayInner, ) *PutBasicTagTag`

NewPutBasicTagTag instantiates a new PutBasicTagTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutBasicTagTagWithDefaults

`func NewPutBasicTagTagWithDefaults() *PutBasicTagTag`

NewPutBasicTagTagWithDefaults instantiates a new PutBasicTagTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PutBasicTagTag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PutBasicTagTag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PutBasicTagTag) SetName(v string)`

SetName sets Name field to given value.


### GetPatterns

`func (o *PutBasicTagTag) GetPatterns() []string`

GetPatterns returns the Patterns field if non-nil, zero value otherwise.

### GetPatternsOk

`func (o *PutBasicTagTag) GetPatternsOk() (*[]string, bool)`

GetPatternsOk returns a tuple with the Patterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatterns

`func (o *PutBasicTagTag) SetPatterns(v []string)`

SetPatterns sets Patterns field to given value.

### HasPatterns

`func (o *PutBasicTagTag) HasPatterns() bool`

HasPatterns returns a boolean if a field has been set.

### GetLeql

`func (o *PutBasicTagTag) GetLeql() BasicLeql`

GetLeql returns the Leql field if non-nil, zero value otherwise.

### GetLeqlOk

`func (o *PutBasicTagTag) GetLeqlOk() (*BasicLeql, bool)`

GetLeqlOk returns a tuple with the Leql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeql

`func (o *PutBasicTagTag) SetLeql(v BasicLeql)`

SetLeql sets Leql field to given value.


### GetSources

`func (o *PutBasicTagTag) GetSources() []SourcesIdArrayInner`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *PutBasicTagTag) GetSourcesOk() (*[]SourcesIdArrayInner, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *PutBasicTagTag) SetSources(v []SourcesIdArrayInner)`

SetSources sets Sources field to given value.


### GetType

`func (o *PutBasicTagTag) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PutBasicTagTag) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PutBasicTagTag) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PutBasicTagTag) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *PutBasicTagTag) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PutBasicTagTag) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PutBasicTagTag) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PutBasicTagTag) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActions

`func (o *PutBasicTagTag) GetActions() []CreateTagAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *PutBasicTagTag) GetActionsOk() (*[]CreateTagAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *PutBasicTagTag) SetActions(v []CreateTagAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *PutBasicTagTag) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetLabels

`func (o *PutBasicTagTag) GetLabels() []LabelResponse`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *PutBasicTagTag) GetLabelsOk() (*[]LabelResponse, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *PutBasicTagTag) SetLabels(v []LabelResponse)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *PutBasicTagTag) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetUserData

`func (o *PutBasicTagTag) GetUserData() map[string]interface{}`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *PutBasicTagTag) GetUserDataOk() (*map[string]interface{}, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *PutBasicTagTag) SetUserData(v map[string]interface{})`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *PutBasicTagTag) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### GetPriority

`func (o *PutBasicTagTag) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *PutBasicTagTag) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *PutBasicTagTag) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *PutBasicTagTag) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


