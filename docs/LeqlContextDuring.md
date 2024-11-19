# LeqlContextDuring

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **int32** | The timestamp of the first (earliest) log entry on the current page. | [optional] 
**To** | Pointer to **int32** | The timestamp of the last (latest) log entry on the current page. | [optional] 

## Methods

### NewLeqlContextDuring

`func NewLeqlContextDuring() *LeqlContextDuring`

NewLeqlContextDuring instantiates a new LeqlContextDuring object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeqlContextDuringWithDefaults

`func NewLeqlContextDuringWithDefaults() *LeqlContextDuring`

NewLeqlContextDuringWithDefaults instantiates a new LeqlContextDuring object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *LeqlContextDuring) GetFrom() int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *LeqlContextDuring) GetFromOk() (*int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *LeqlContextDuring) SetFrom(v int32)`

SetFrom sets From field to given value.

### HasFrom

`func (o *LeqlContextDuring) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *LeqlContextDuring) GetTo() int32`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *LeqlContextDuring) GetToOk() (*int32, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *LeqlContextDuring) SetTo(v int32)`

SetTo sets To field to given value.

### HasTo

`func (o *LeqlContextDuring) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


