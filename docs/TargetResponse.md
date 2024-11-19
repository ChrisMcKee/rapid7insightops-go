# TargetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the notification target. | 
**Name** | **string** | The name of the notification target. | 
**Description** | **string** | The description of the notification target. | 
**Type** | [**NotificationTypeEnum**](NotificationTypeEnum.md) |  | 
**ParamsSet** | [**TargetResponseParamsSet**](TargetResponseParamsSet.md) |  | 
**UserData** | **map[string]interface{}** | A list of key-value pairs that may indicate some auxiliary information about the notification target. | 
**AlertContentSet** | [**TargetAlertContentSet**](TargetAlertContentSet.md) |  | 

## Methods

### NewTargetResponse

`func NewTargetResponse(id string, name string, description string, type_ NotificationTypeEnum, paramsSet TargetResponseParamsSet, userData map[string]interface{}, alertContentSet TargetAlertContentSet, ) *TargetResponse`

NewTargetResponse instantiates a new TargetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetResponseWithDefaults

`func NewTargetResponseWithDefaults() *TargetResponse`

NewTargetResponseWithDefaults instantiates a new TargetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TargetResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TargetResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TargetResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TargetResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TargetResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TargetResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TargetResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *TargetResponse) GetType() NotificationTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TargetResponse) GetTypeOk() (*NotificationTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TargetResponse) SetType(v NotificationTypeEnum)`

SetType sets Type field to given value.


### GetParamsSet

`func (o *TargetResponse) GetParamsSet() TargetResponseParamsSet`

GetParamsSet returns the ParamsSet field if non-nil, zero value otherwise.

### GetParamsSetOk

`func (o *TargetResponse) GetParamsSetOk() (*TargetResponseParamsSet, bool)`

GetParamsSetOk returns a tuple with the ParamsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamsSet

`func (o *TargetResponse) SetParamsSet(v TargetResponseParamsSet)`

SetParamsSet sets ParamsSet field to given value.


### GetUserData

`func (o *TargetResponse) GetUserData() map[string]interface{}`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *TargetResponse) GetUserDataOk() (*map[string]interface{}, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *TargetResponse) SetUserData(v map[string]interface{})`

SetUserData sets UserData field to given value.


### GetAlertContentSet

`func (o *TargetResponse) GetAlertContentSet() TargetAlertContentSet`

GetAlertContentSet returns the AlertContentSet field if non-nil, zero value otherwise.

### GetAlertContentSetOk

`func (o *TargetResponse) GetAlertContentSetOk() (*TargetAlertContentSet, bool)`

GetAlertContentSetOk returns a tuple with the AlertContentSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertContentSet

`func (o *TargetResponse) SetAlertContentSet(v TargetAlertContentSet)`

SetAlertContentSet sets AlertContentSet field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


