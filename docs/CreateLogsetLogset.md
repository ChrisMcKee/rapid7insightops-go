# CreateLogsetLogset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The UUID of the log set. | [optional] 
**Name** | **interface{}** | The name of the log set. | 
**Description** | Pointer to **string** | The description of the log set. | [optional] 
**UserData** | Pointer to **map[string]interface{}** | A list of key-value pairs that may indicate some auxiliary information about the log set. | [optional] 
**LogsInfo** | Pointer to [**[]MemberInfoLogsetInner**](MemberInfoLogsetInner.md) | Information on each log that this log set contains.  | [optional] 
**TokenSeed** | Pointer to **string** | The string that will be used in combination with the log set name to generate the log set ID. | [optional] 
**Rrn** | Pointer to **interface{}** | The Rapid7 Resource Name (RRN) of the log set. The RRN is a unique identifier across the Rapid7 Platform. | [optional] 

## Methods

### NewCreateLogsetLogset

`func NewCreateLogsetLogset(name interface{}, ) *CreateLogsetLogset`

NewCreateLogsetLogset instantiates a new CreateLogsetLogset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLogsetLogsetWithDefaults

`func NewCreateLogsetLogsetWithDefaults() *CreateLogsetLogset`

NewCreateLogsetLogsetWithDefaults instantiates a new CreateLogsetLogset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateLogsetLogset) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateLogsetLogset) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateLogsetLogset) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateLogsetLogset) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CreateLogsetLogset) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateLogsetLogset) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateLogsetLogset) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *CreateLogsetLogset) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateLogsetLogset) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *CreateLogsetLogset) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateLogsetLogset) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateLogsetLogset) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateLogsetLogset) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUserData

`func (o *CreateLogsetLogset) GetUserData() map[string]interface{}`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *CreateLogsetLogset) GetUserDataOk() (*map[string]interface{}, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *CreateLogsetLogset) SetUserData(v map[string]interface{})`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *CreateLogsetLogset) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### GetLogsInfo

`func (o *CreateLogsetLogset) GetLogsInfo() []MemberInfoLogsetInner`

GetLogsInfo returns the LogsInfo field if non-nil, zero value otherwise.

### GetLogsInfoOk

`func (o *CreateLogsetLogset) GetLogsInfoOk() (*[]MemberInfoLogsetInner, bool)`

GetLogsInfoOk returns a tuple with the LogsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsInfo

`func (o *CreateLogsetLogset) SetLogsInfo(v []MemberInfoLogsetInner)`

SetLogsInfo sets LogsInfo field to given value.

### HasLogsInfo

`func (o *CreateLogsetLogset) HasLogsInfo() bool`

HasLogsInfo returns a boolean if a field has been set.

### GetTokenSeed

`func (o *CreateLogsetLogset) GetTokenSeed() string`

GetTokenSeed returns the TokenSeed field if non-nil, zero value otherwise.

### GetTokenSeedOk

`func (o *CreateLogsetLogset) GetTokenSeedOk() (*string, bool)`

GetTokenSeedOk returns a tuple with the TokenSeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenSeed

`func (o *CreateLogsetLogset) SetTokenSeed(v string)`

SetTokenSeed sets TokenSeed field to given value.

### HasTokenSeed

`func (o *CreateLogsetLogset) HasTokenSeed() bool`

HasTokenSeed returns a boolean if a field has been set.

### GetRrn

`func (o *CreateLogsetLogset) GetRrn() interface{}`

GetRrn returns the Rrn field if non-nil, zero value otherwise.

### GetRrnOk

`func (o *CreateLogsetLogset) GetRrnOk() (*interface{}, bool)`

GetRrnOk returns a tuple with the Rrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrn

`func (o *CreateLogsetLogset) SetRrn(v interface{})`

SetRrn sets Rrn field to given value.

### HasRrn

`func (o *CreateLogsetLogset) HasRrn() bool`

HasRrn returns a boolean if a field has been set.

### SetRrnNil

`func (o *CreateLogsetLogset) SetRrnNil(b bool)`

 SetRrnNil sets the value for Rrn to be an explicit nil

### UnsetRrn
`func (o *CreateLogsetLogset) UnsetRrn()`

UnsetRrn ensures that no value is present for Rrn, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


