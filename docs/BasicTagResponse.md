# BasicTagResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Always set to \&quot;Alert\&quot;. | [optional] 
**Id** | Pointer to **string** | The unique id of the detection rule. | [optional] 
**Name** | Pointer to **string** | The name of the detection rule. | [optional] 
**Description** | Pointer to **string** | The description of the detection rule. | [optional] 
**Sources** | Pointer to [**[]TagSourceResponse**](TagSourceResponse.md) | The collection of logs the detection rule operates on. | [optional] 
**Actions** | Pointer to [**[]ActionResponse**](ActionResponse.md) | The notifications attached to the detection rule. | [optional] 
**Patterns** | Pointer to **[]string** | Use the &#x60;leql&#x60; parameter instead of this parameter. | [optional] 
**Leql** | Pointer to [**BasicLeql**](BasicLeql.md) |  | [optional] 
**Priority** | Pointer to **int32** | This ensures investigations are ordered by priority in Investigation Management. Must be &gt;&#x3D;0 | [optional] 
**UserData** | Pointer to **map[string]interface{}** | A list of key-value pairs that may indicate some auxiliary information about the change detection rule. | [optional] 
**Labels** | Pointer to [**[]LabelResponse**](LabelResponse.md) | The labels attached to the detection rule. | [optional] 

## Methods

### NewBasicTagResponse

`func NewBasicTagResponse() *BasicTagResponse`

NewBasicTagResponse instantiates a new BasicTagResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasicTagResponseWithDefaults

`func NewBasicTagResponseWithDefaults() *BasicTagResponse`

NewBasicTagResponseWithDefaults instantiates a new BasicTagResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BasicTagResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BasicTagResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BasicTagResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BasicTagResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *BasicTagResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BasicTagResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BasicTagResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BasicTagResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BasicTagResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BasicTagResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BasicTagResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BasicTagResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *BasicTagResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BasicTagResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BasicTagResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BasicTagResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSources

`func (o *BasicTagResponse) GetSources() []TagSourceResponse`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *BasicTagResponse) GetSourcesOk() (*[]TagSourceResponse, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *BasicTagResponse) SetSources(v []TagSourceResponse)`

SetSources sets Sources field to given value.

### HasSources

`func (o *BasicTagResponse) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetActions

`func (o *BasicTagResponse) GetActions() []ActionResponse`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *BasicTagResponse) GetActionsOk() (*[]ActionResponse, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *BasicTagResponse) SetActions(v []ActionResponse)`

SetActions sets Actions field to given value.

### HasActions

`func (o *BasicTagResponse) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetPatterns

`func (o *BasicTagResponse) GetPatterns() []string`

GetPatterns returns the Patterns field if non-nil, zero value otherwise.

### GetPatternsOk

`func (o *BasicTagResponse) GetPatternsOk() (*[]string, bool)`

GetPatternsOk returns a tuple with the Patterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatterns

`func (o *BasicTagResponse) SetPatterns(v []string)`

SetPatterns sets Patterns field to given value.

### HasPatterns

`func (o *BasicTagResponse) HasPatterns() bool`

HasPatterns returns a boolean if a field has been set.

### GetLeql

`func (o *BasicTagResponse) GetLeql() BasicLeql`

GetLeql returns the Leql field if non-nil, zero value otherwise.

### GetLeqlOk

`func (o *BasicTagResponse) GetLeqlOk() (*BasicLeql, bool)`

GetLeqlOk returns a tuple with the Leql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeql

`func (o *BasicTagResponse) SetLeql(v BasicLeql)`

SetLeql sets Leql field to given value.

### HasLeql

`func (o *BasicTagResponse) HasLeql() bool`

HasLeql returns a boolean if a field has been set.

### GetPriority

`func (o *BasicTagResponse) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *BasicTagResponse) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *BasicTagResponse) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *BasicTagResponse) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetUserData

`func (o *BasicTagResponse) GetUserData() map[string]interface{}`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *BasicTagResponse) GetUserDataOk() (*map[string]interface{}, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *BasicTagResponse) SetUserData(v map[string]interface{})`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *BasicTagResponse) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### GetLabels

`func (o *BasicTagResponse) GetLabels() []LabelResponse`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *BasicTagResponse) GetLabelsOk() (*[]LabelResponse, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *BasicTagResponse) SetLabels(v []LabelResponse)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *BasicTagResponse) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


