# LeqlVariableRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Variable name. | 
**Description** | Pointer to **NullableString** | Variable description. | [optional] 
**Value** | **string** | Variable contents. | 

## Methods

### NewLeqlVariableRequest

`func NewLeqlVariableRequest(name string, value string, ) *LeqlVariableRequest`

NewLeqlVariableRequest instantiates a new LeqlVariableRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeqlVariableRequestWithDefaults

`func NewLeqlVariableRequestWithDefaults() *LeqlVariableRequest`

NewLeqlVariableRequestWithDefaults instantiates a new LeqlVariableRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LeqlVariableRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LeqlVariableRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LeqlVariableRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *LeqlVariableRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LeqlVariableRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LeqlVariableRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LeqlVariableRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *LeqlVariableRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *LeqlVariableRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetValue

`func (o *LeqlVariableRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *LeqlVariableRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *LeqlVariableRequest) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


