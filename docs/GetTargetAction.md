# GetTargetAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the notification target. | [optional] 
**Description** | Pointer to **string** | The description of the notifiaction target. | [optional] 
**Type** | Pointer to [**NotificationTypeEnum**](NotificationTypeEnum.md) |  | [optional] 
**ParamsSet** | Pointer to [**TargetResponseParamsSet**](TargetResponseParamsSet.md) |  | [optional] 
**AlertContentSet** | Pointer to [**TargetAlertContentSet**](TargetAlertContentSet.md) |  | [optional] 
**UserData** | Pointer to **map[string]interface{}** | A list of key-value pairs that may indicate some auxiliary information about the notification target. | [optional] 

## Methods

### NewGetTargetAction

`func NewGetTargetAction() *GetTargetAction`

NewGetTargetAction instantiates a new GetTargetAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTargetActionWithDefaults

`func NewGetTargetActionWithDefaults() *GetTargetAction`

NewGetTargetActionWithDefaults instantiates a new GetTargetAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetTargetAction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetTargetAction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetTargetAction) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetTargetAction) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetTargetAction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetTargetAction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetTargetAction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetTargetAction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *GetTargetAction) GetType() NotificationTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetTargetAction) GetTypeOk() (*NotificationTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetTargetAction) SetType(v NotificationTypeEnum)`

SetType sets Type field to given value.

### HasType

`func (o *GetTargetAction) HasType() bool`

HasType returns a boolean if a field has been set.

### GetParamsSet

`func (o *GetTargetAction) GetParamsSet() TargetResponseParamsSet`

GetParamsSet returns the ParamsSet field if non-nil, zero value otherwise.

### GetParamsSetOk

`func (o *GetTargetAction) GetParamsSetOk() (*TargetResponseParamsSet, bool)`

GetParamsSetOk returns a tuple with the ParamsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamsSet

`func (o *GetTargetAction) SetParamsSet(v TargetResponseParamsSet)`

SetParamsSet sets ParamsSet field to given value.

### HasParamsSet

`func (o *GetTargetAction) HasParamsSet() bool`

HasParamsSet returns a boolean if a field has been set.

### GetAlertContentSet

`func (o *GetTargetAction) GetAlertContentSet() TargetAlertContentSet`

GetAlertContentSet returns the AlertContentSet field if non-nil, zero value otherwise.

### GetAlertContentSetOk

`func (o *GetTargetAction) GetAlertContentSetOk() (*TargetAlertContentSet, bool)`

GetAlertContentSetOk returns a tuple with the AlertContentSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertContentSet

`func (o *GetTargetAction) SetAlertContentSet(v TargetAlertContentSet)`

SetAlertContentSet sets AlertContentSet field to given value.

### HasAlertContentSet

`func (o *GetTargetAction) HasAlertContentSet() bool`

HasAlertContentSet returns a boolean if a field has been set.

### GetUserData

`func (o *GetTargetAction) GetUserData() map[string]interface{}`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *GetTargetAction) GetUserDataOk() (*map[string]interface{}, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *GetTargetAction) SetUserData(v map[string]interface{})`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *GetTargetAction) HasUserData() bool`

HasUserData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


