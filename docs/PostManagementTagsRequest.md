# PostManagementTagsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tag** | Pointer to [**CreateTagAnomalyAlertTag**](CreateTagAnomalyAlertTag.md) |  | [optional] 

## Methods

### NewPostManagementTagsRequest

`func NewPostManagementTagsRequest() *PostManagementTagsRequest`

NewPostManagementTagsRequest instantiates a new PostManagementTagsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostManagementTagsRequestWithDefaults

`func NewPostManagementTagsRequestWithDefaults() *PostManagementTagsRequest`

NewPostManagementTagsRequestWithDefaults instantiates a new PostManagementTagsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTag

`func (o *PostManagementTagsRequest) GetTag() CreateTagAnomalyAlertTag`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *PostManagementTagsRequest) GetTagOk() (*CreateTagAnomalyAlertTag, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *PostManagementTagsRequest) SetTag(v CreateTagAnomalyAlertTag)`

SetTag sets Tag field to given value.

### HasTag

`func (o *PostManagementTagsRequest) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


