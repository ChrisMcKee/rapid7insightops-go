# CreatePatchTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the notification target. | 
**Description** | **string** | The description of the notification target. | 
**Type** | [**NotificationTypeEnum**](NotificationTypeEnum.md) |  | 
**ParamsSet** | [**CreatePatchTargetParamsSet**](CreatePatchTargetParamsSet.md) |  | 
**AlertContentSet** | [**TargetAlertContentSet**](TargetAlertContentSet.md) |  | 
**UserData** | **map[string]interface{}** | A list of key-value pairs that may indicate some auxiliary information about the notification target. | 

## Methods

### NewCreatePatchTarget

`func NewCreatePatchTarget(name string, description string, type_ NotificationTypeEnum, paramsSet CreatePatchTargetParamsSet, alertContentSet TargetAlertContentSet, userData map[string]interface{}, ) *CreatePatchTarget`

NewCreatePatchTarget instantiates a new CreatePatchTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePatchTargetWithDefaults

`func NewCreatePatchTargetWithDefaults() *CreatePatchTarget`

NewCreatePatchTargetWithDefaults instantiates a new CreatePatchTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreatePatchTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePatchTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePatchTarget) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreatePatchTarget) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreatePatchTarget) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreatePatchTarget) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *CreatePatchTarget) GetType() NotificationTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreatePatchTarget) GetTypeOk() (*NotificationTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreatePatchTarget) SetType(v NotificationTypeEnum)`

SetType sets Type field to given value.


### GetParamsSet

`func (o *CreatePatchTarget) GetParamsSet() CreatePatchTargetParamsSet`

GetParamsSet returns the ParamsSet field if non-nil, zero value otherwise.

### GetParamsSetOk

`func (o *CreatePatchTarget) GetParamsSetOk() (*CreatePatchTargetParamsSet, bool)`

GetParamsSetOk returns a tuple with the ParamsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamsSet

`func (o *CreatePatchTarget) SetParamsSet(v CreatePatchTargetParamsSet)`

SetParamsSet sets ParamsSet field to given value.


### GetAlertContentSet

`func (o *CreatePatchTarget) GetAlertContentSet() TargetAlertContentSet`

GetAlertContentSet returns the AlertContentSet field if non-nil, zero value otherwise.

### GetAlertContentSetOk

`func (o *CreatePatchTarget) GetAlertContentSetOk() (*TargetAlertContentSet, bool)`

GetAlertContentSetOk returns a tuple with the AlertContentSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertContentSet

`func (o *CreatePatchTarget) SetAlertContentSet(v TargetAlertContentSet)`

SetAlertContentSet sets AlertContentSet field to given value.


### GetUserData

`func (o *CreatePatchTarget) GetUserData() map[string]interface{}`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *CreatePatchTarget) GetUserDataOk() (*map[string]interface{}, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *CreatePatchTarget) SetUserData(v map[string]interface{})`

SetUserData sets UserData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


