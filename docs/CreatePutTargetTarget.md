# CreatePutTargetTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the target. | 
**Description** | Pointer to **string** | The description of the target. | [optional] 
**Type** | [**NotificationTypeEnum**](NotificationTypeEnum.md) |  | 
**ParamsSet** | [**CreatePatchTargetParamsSet**](CreatePatchTargetParamsSet.md) |  | 
**AlertContentSet** | [**TargetAlertContentSet**](TargetAlertContentSet.md) |  | 
**UserData** | **map[string]interface{}** | A list of key-value pairs that may indicate some auxiliary information about the notification target. | 

## Methods

### NewCreatePutTargetTarget

`func NewCreatePutTargetTarget(name string, type_ NotificationTypeEnum, paramsSet CreatePatchTargetParamsSet, alertContentSet TargetAlertContentSet, userData map[string]interface{}, ) *CreatePutTargetTarget`

NewCreatePutTargetTarget instantiates a new CreatePutTargetTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePutTargetTargetWithDefaults

`func NewCreatePutTargetTargetWithDefaults() *CreatePutTargetTarget`

NewCreatePutTargetTargetWithDefaults instantiates a new CreatePutTargetTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreatePutTargetTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePutTargetTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePutTargetTarget) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreatePutTargetTarget) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreatePutTargetTarget) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreatePutTargetTarget) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreatePutTargetTarget) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *CreatePutTargetTarget) GetType() NotificationTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreatePutTargetTarget) GetTypeOk() (*NotificationTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreatePutTargetTarget) SetType(v NotificationTypeEnum)`

SetType sets Type field to given value.


### GetParamsSet

`func (o *CreatePutTargetTarget) GetParamsSet() CreatePatchTargetParamsSet`

GetParamsSet returns the ParamsSet field if non-nil, zero value otherwise.

### GetParamsSetOk

`func (o *CreatePutTargetTarget) GetParamsSetOk() (*CreatePatchTargetParamsSet, bool)`

GetParamsSetOk returns a tuple with the ParamsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamsSet

`func (o *CreatePutTargetTarget) SetParamsSet(v CreatePatchTargetParamsSet)`

SetParamsSet sets ParamsSet field to given value.


### GetAlertContentSet

`func (o *CreatePutTargetTarget) GetAlertContentSet() TargetAlertContentSet`

GetAlertContentSet returns the AlertContentSet field if non-nil, zero value otherwise.

### GetAlertContentSetOk

`func (o *CreatePutTargetTarget) GetAlertContentSetOk() (*TargetAlertContentSet, bool)`

GetAlertContentSetOk returns a tuple with the AlertContentSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertContentSet

`func (o *CreatePutTargetTarget) SetAlertContentSet(v TargetAlertContentSet)`

SetAlertContentSet sets AlertContentSet field to given value.


### GetUserData

`func (o *CreatePutTargetTarget) GetUserData() map[string]interface{}`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *CreatePutTargetTarget) GetUserDataOk() (*map[string]interface{}, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *CreatePutTargetTarget) SetUserData(v map[string]interface{})`

SetUserData sets UserData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


