# LabelsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | Pointer to [**[]LabelResponse**](LabelResponse.md) | The labels attached to the detection rule. | [optional] 

## Methods

### NewLabelsResponse

`func NewLabelsResponse() *LabelsResponse`

NewLabelsResponse instantiates a new LabelsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelsResponseWithDefaults

`func NewLabelsResponseWithDefaults() *LabelsResponse`

NewLabelsResponseWithDefaults instantiates a new LabelsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *LabelsResponse) GetLabels() []LabelResponse`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *LabelsResponse) GetLabelsOk() (*[]LabelResponse, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *LabelsResponse) SetLabels(v []LabelResponse)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *LabelsResponse) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


