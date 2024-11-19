# CreateTagTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertContentSet** | Pointer to [**TargetAlertContentSet**](TargetAlertContentSet.md) |  | [optional] 
**ParamsSet** | [**TargetResponseParamsSet**](TargetResponseParamsSet.md) |  | 
**Type** | [**NotificationTypeEnum**](NotificationTypeEnum.md) |  | 

## Methods

### NewCreateTagTarget

`func NewCreateTagTarget(paramsSet TargetResponseParamsSet, type_ NotificationTypeEnum, ) *CreateTagTarget`

NewCreateTagTarget instantiates a new CreateTagTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTagTargetWithDefaults

`func NewCreateTagTargetWithDefaults() *CreateTagTarget`

NewCreateTagTargetWithDefaults instantiates a new CreateTagTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertContentSet

`func (o *CreateTagTarget) GetAlertContentSet() TargetAlertContentSet`

GetAlertContentSet returns the AlertContentSet field if non-nil, zero value otherwise.

### GetAlertContentSetOk

`func (o *CreateTagTarget) GetAlertContentSetOk() (*TargetAlertContentSet, bool)`

GetAlertContentSetOk returns a tuple with the AlertContentSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertContentSet

`func (o *CreateTagTarget) SetAlertContentSet(v TargetAlertContentSet)`

SetAlertContentSet sets AlertContentSet field to given value.

### HasAlertContentSet

`func (o *CreateTagTarget) HasAlertContentSet() bool`

HasAlertContentSet returns a boolean if a field has been set.

### GetParamsSet

`func (o *CreateTagTarget) GetParamsSet() TargetResponseParamsSet`

GetParamsSet returns the ParamsSet field if non-nil, zero value otherwise.

### GetParamsSetOk

`func (o *CreateTagTarget) GetParamsSetOk() (*TargetResponseParamsSet, bool)`

GetParamsSetOk returns a tuple with the ParamsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamsSet

`func (o *CreateTagTarget) SetParamsSet(v TargetResponseParamsSet)`

SetParamsSet sets ParamsSet field to given value.


### GetType

`func (o *CreateTagTarget) GetType() NotificationTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateTagTarget) GetTypeOk() (*NotificationTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateTagTarget) SetType(v NotificationTypeEnum)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


