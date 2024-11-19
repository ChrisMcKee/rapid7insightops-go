# LogTopkeysResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Topkeys** | [**[]LogTopkeysResponseTopkeysInner**](LogTopkeysResponseTopkeysInner.md) | The list of commonly occurring keys for this log. | 

## Methods

### NewLogTopkeysResponse

`func NewLogTopkeysResponse(topkeys []LogTopkeysResponseTopkeysInner, ) *LogTopkeysResponse`

NewLogTopkeysResponse instantiates a new LogTopkeysResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogTopkeysResponseWithDefaults

`func NewLogTopkeysResponseWithDefaults() *LogTopkeysResponse`

NewLogTopkeysResponseWithDefaults instantiates a new LogTopkeysResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopkeys

`func (o *LogTopkeysResponse) GetTopkeys() []LogTopkeysResponseTopkeysInner`

GetTopkeys returns the Topkeys field if non-nil, zero value otherwise.

### GetTopkeysOk

`func (o *LogTopkeysResponse) GetTopkeysOk() (*[]LogTopkeysResponseTopkeysInner, bool)`

GetTopkeysOk returns a tuple with the Topkeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopkeys

`func (o *LogTopkeysResponse) SetTopkeys(v []LogTopkeysResponseTopkeysInner)`

SetTopkeys sets Topkeys field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


