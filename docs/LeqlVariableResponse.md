# LeqlVariableResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique variable ID. UUID. | 
**Name** | **string** | Variable name. | 
**Description** | Pointer to **NullableString** | Variable description. | [optional] 
**Value** | **string** | Variable contents. | 
**UserId** | Pointer to **string** | The ID of the user that created the variable. | [optional] 
**Created** | Pointer to **string** | The date/time when the variable was created. | [optional] 
**Updated** | Pointer to **string** | The date/time when the variable was updated. | [optional] 
**Reserved** | Pointer to **bool** | Reserved variables cannot be modified. | [optional] 

## Methods

### NewLeqlVariableResponse

`func NewLeqlVariableResponse(id string, name string, value string, ) *LeqlVariableResponse`

NewLeqlVariableResponse instantiates a new LeqlVariableResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeqlVariableResponseWithDefaults

`func NewLeqlVariableResponseWithDefaults() *LeqlVariableResponse`

NewLeqlVariableResponseWithDefaults instantiates a new LeqlVariableResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LeqlVariableResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LeqlVariableResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LeqlVariableResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *LeqlVariableResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LeqlVariableResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LeqlVariableResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *LeqlVariableResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LeqlVariableResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LeqlVariableResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LeqlVariableResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *LeqlVariableResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *LeqlVariableResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetValue

`func (o *LeqlVariableResponse) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *LeqlVariableResponse) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *LeqlVariableResponse) SetValue(v string)`

SetValue sets Value field to given value.


### GetUserId

`func (o *LeqlVariableResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *LeqlVariableResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *LeqlVariableResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *LeqlVariableResponse) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetCreated

`func (o *LeqlVariableResponse) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *LeqlVariableResponse) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *LeqlVariableResponse) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *LeqlVariableResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *LeqlVariableResponse) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *LeqlVariableResponse) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *LeqlVariableResponse) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *LeqlVariableResponse) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetReserved

`func (o *LeqlVariableResponse) GetReserved() bool`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *LeqlVariableResponse) GetReservedOk() (*bool, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *LeqlVariableResponse) SetReserved(v bool)`

SetReserved sets Reserved field to given value.

### HasReserved

`func (o *LeqlVariableResponse) HasReserved() bool`

HasReserved returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


