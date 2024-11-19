# LabelResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the label. | 
**Sn** | **int32** | A property which exists for legacy reasons and is no longer used for anything. | 
**Name** | **string** | The name of the label. | 
**Color** | **string** | The color of the label in HEX code. | 
**Reserved** | **bool** | When set to true, the label is read-only and cannot be modified. | 

## Methods

### NewLabelResponse

`func NewLabelResponse(id string, sn int32, name string, color string, reserved bool, ) *LabelResponse`

NewLabelResponse instantiates a new LabelResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelResponseWithDefaults

`func NewLabelResponseWithDefaults() *LabelResponse`

NewLabelResponseWithDefaults instantiates a new LabelResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LabelResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LabelResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LabelResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSn

`func (o *LabelResponse) GetSn() int32`

GetSn returns the Sn field if non-nil, zero value otherwise.

### GetSnOk

`func (o *LabelResponse) GetSnOk() (*int32, bool)`

GetSnOk returns a tuple with the Sn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSn

`func (o *LabelResponse) SetSn(v int32)`

SetSn sets Sn field to given value.


### GetName

`func (o *LabelResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LabelResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LabelResponse) SetName(v string)`

SetName sets Name field to given value.


### GetColor

`func (o *LabelResponse) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *LabelResponse) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *LabelResponse) SetColor(v string)`

SetColor sets Color field to given value.


### GetReserved

`func (o *LabelResponse) GetReserved() bool`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *LabelResponse) GetReservedOk() (*bool, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *LabelResponse) SetReserved(v bool)`

SetReserved sets Reserved field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


