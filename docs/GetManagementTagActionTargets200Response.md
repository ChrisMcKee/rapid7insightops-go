# GetManagementTagActionTargets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Targets** | Pointer to [**[]GetTargetAction**](GetTargetAction.md) | All the notification targets for this account. | [optional] 

## Methods

### NewGetManagementTagActionTargets200Response

`func NewGetManagementTagActionTargets200Response() *GetManagementTagActionTargets200Response`

NewGetManagementTagActionTargets200Response instantiates a new GetManagementTagActionTargets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetManagementTagActionTargets200ResponseWithDefaults

`func NewGetManagementTagActionTargets200ResponseWithDefaults() *GetManagementTagActionTargets200Response`

NewGetManagementTagActionTargets200ResponseWithDefaults instantiates a new GetManagementTagActionTargets200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargets

`func (o *GetManagementTagActionTargets200Response) GetTargets() []GetTargetAction`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *GetManagementTagActionTargets200Response) GetTargetsOk() (*[]GetTargetAction, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *GetManagementTagActionTargets200Response) SetTargets(v []GetTargetAction)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *GetManagementTagActionTargets200Response) HasTargets() bool`

HasTargets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


