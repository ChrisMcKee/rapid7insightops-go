# GetManagementTags200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to [**[]GetManagementTags200ResponseTagsInner**](GetManagementTags200ResponseTagsInner.md) | All the basic detection rules for this account. | [optional] 

## Methods

### NewGetManagementTags200Response

`func NewGetManagementTags200Response() *GetManagementTags200Response`

NewGetManagementTags200Response instantiates a new GetManagementTags200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetManagementTags200ResponseWithDefaults

`func NewGetManagementTags200ResponseWithDefaults() *GetManagementTags200Response`

NewGetManagementTags200ResponseWithDefaults instantiates a new GetManagementTags200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *GetManagementTags200Response) GetTags() []GetManagementTags200ResponseTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetManagementTags200Response) GetTagsOk() (*[]GetManagementTags200ResponseTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetManagementTags200Response) SetTags(v []GetManagementTags200ResponseTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetManagementTags200Response) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


