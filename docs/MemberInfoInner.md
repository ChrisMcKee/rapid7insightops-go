# MemberInfoInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The UUID of the log set. | 
**Name** | **string** | The name of the log set. | 
**Rrn** | Pointer to **string** | The Rapid7 Resource Name (RRN) of the Log Set. The RRN is a unique identifier across the Rapid7 Platform. | [optional] 
**Links** | Pointer to [**[]LinksInner**](LinksInner.md) |  | [optional] 

## Methods

### NewMemberInfoInner

`func NewMemberInfoInner(id string, name string, ) *MemberInfoInner`

NewMemberInfoInner instantiates a new MemberInfoInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberInfoInnerWithDefaults

`func NewMemberInfoInnerWithDefaults() *MemberInfoInner`

NewMemberInfoInnerWithDefaults instantiates a new MemberInfoInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MemberInfoInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MemberInfoInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MemberInfoInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *MemberInfoInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MemberInfoInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MemberInfoInner) SetName(v string)`

SetName sets Name field to given value.


### GetRrn

`func (o *MemberInfoInner) GetRrn() string`

GetRrn returns the Rrn field if non-nil, zero value otherwise.

### GetRrnOk

`func (o *MemberInfoInner) GetRrnOk() (*string, bool)`

GetRrnOk returns a tuple with the Rrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrn

`func (o *MemberInfoInner) SetRrn(v string)`

SetRrn sets Rrn field to given value.

### HasRrn

`func (o *MemberInfoInner) HasRrn() bool`

HasRrn returns a boolean if a field has been set.

### GetLinks

`func (o *MemberInfoInner) GetLinks() []LinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MemberInfoInner) GetLinksOk() (*[]LinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MemberInfoInner) SetLinks(v []LinksInner)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MemberInfoInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


