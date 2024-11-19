# LabelsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**[]LinksInner**](LinksInner.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 

## Methods

### NewLabelsInner

`func NewLabelsInner() *LabelsInner`

NewLabelsInner instantiates a new LabelsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelsInnerWithDefaults

`func NewLabelsInnerWithDefaults() *LabelsInner`

NewLabelsInnerWithDefaults instantiates a new LabelsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *LabelsInner) GetLinks() []LinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *LabelsInner) GetLinksOk() (*[]LinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *LabelsInner) SetLinks(v []LinksInner)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *LabelsInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *LabelsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LabelsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LabelsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LabelsInner) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


