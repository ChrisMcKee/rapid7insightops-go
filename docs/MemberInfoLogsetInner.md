# MemberInfoLogsetInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The UUID of the log. | 
**Name** | **string** | The name of the log. | 
**Rrn** | Pointer to **string** | The Rapid7 Resource Name (RRN) of the log. The RRN is a unique identifier across the Rapid7 Platform. | [optional] 
**Links** | Pointer to [**[]LinksInner**](LinksInner.md) |  | [optional] 

## Methods

### NewMemberInfoLogsetInner

`func NewMemberInfoLogsetInner(id string, name string, ) *MemberInfoLogsetInner`

NewMemberInfoLogsetInner instantiates a new MemberInfoLogsetInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberInfoLogsetInnerWithDefaults

`func NewMemberInfoLogsetInnerWithDefaults() *MemberInfoLogsetInner`

NewMemberInfoLogsetInnerWithDefaults instantiates a new MemberInfoLogsetInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MemberInfoLogsetInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MemberInfoLogsetInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MemberInfoLogsetInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *MemberInfoLogsetInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MemberInfoLogsetInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MemberInfoLogsetInner) SetName(v string)`

SetName sets Name field to given value.


### GetRrn

`func (o *MemberInfoLogsetInner) GetRrn() string`

GetRrn returns the Rrn field if non-nil, zero value otherwise.

### GetRrnOk

`func (o *MemberInfoLogsetInner) GetRrnOk() (*string, bool)`

GetRrnOk returns a tuple with the Rrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrn

`func (o *MemberInfoLogsetInner) SetRrn(v string)`

SetRrn sets Rrn field to given value.

### HasRrn

`func (o *MemberInfoLogsetInner) HasRrn() bool`

HasRrn returns a boolean if a field has been set.

### GetLinks

`func (o *MemberInfoLogsetInner) GetLinks() []LinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MemberInfoLogsetInner) GetLinksOk() (*[]LinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MemberInfoLogsetInner) SetLinks(v []LinksInner)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MemberInfoLogsetInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


