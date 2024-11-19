# LeqlContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**During** | Pointer to [**LeqlContextDuring**](LeqlContextDuring.md) |  | [optional] 
**Statement** | Pointer to **string** | Always null. | [optional] 

## Methods

### NewLeqlContext

`func NewLeqlContext() *LeqlContext`

NewLeqlContext instantiates a new LeqlContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeqlContextWithDefaults

`func NewLeqlContextWithDefaults() *LeqlContext`

NewLeqlContextWithDefaults instantiates a new LeqlContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuring

`func (o *LeqlContext) GetDuring() LeqlContextDuring`

GetDuring returns the During field if non-nil, zero value otherwise.

### GetDuringOk

`func (o *LeqlContext) GetDuringOk() (*LeqlContextDuring, bool)`

GetDuringOk returns a tuple with the During field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuring

`func (o *LeqlContext) SetDuring(v LeqlContextDuring)`

SetDuring sets During field to given value.

### HasDuring

`func (o *LeqlContext) HasDuring() bool`

HasDuring returns a boolean if a field has been set.

### GetStatement

`func (o *LeqlContext) GetStatement() string`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *LeqlContext) GetStatementOk() (*string, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *LeqlContext) SetStatement(v string)`

SetStatement sets Statement field to given value.

### HasStatement

`func (o *LeqlContext) HasStatement() bool`

HasStatement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


