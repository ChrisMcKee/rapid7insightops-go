# KvpInfoInnerValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** | the data associated with the key. | 
**Start** | **int32** | the position in the log entry of the first character of the value (first character in the log entry is at position 0). | 
**End** | **int32** | the position in the log entry of the next character after the last character of the value. | 

## Methods

### NewKvpInfoInnerValue

`func NewKvpInfoInnerValue(text string, start int32, end int32, ) *KvpInfoInnerValue`

NewKvpInfoInnerValue instantiates a new KvpInfoInnerValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvpInfoInnerValueWithDefaults

`func NewKvpInfoInnerValueWithDefaults() *KvpInfoInnerValue`

NewKvpInfoInnerValueWithDefaults instantiates a new KvpInfoInnerValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *KvpInfoInnerValue) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *KvpInfoInnerValue) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *KvpInfoInnerValue) SetText(v string)`

SetText sets Text field to given value.


### GetStart

`func (o *KvpInfoInnerValue) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *KvpInfoInnerValue) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *KvpInfoInnerValue) SetStart(v int32)`

SetStart sets Start field to given value.


### GetEnd

`func (o *KvpInfoInnerValue) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *KvpInfoInnerValue) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *KvpInfoInnerValue) SetEnd(v int32)`

SetEnd sets End field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


