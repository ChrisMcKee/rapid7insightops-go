# LeqlLiveTail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statement** | Pointer to **string** | \\ The LEQL query that incoming log entries are matched against. If empty, the query matches all incoming log entries.  Cannot be a statistical query (&#39;calculate&#39; or &#39;groupby&#39; query).  | [optional] 

## Methods

### NewLeqlLiveTail

`func NewLeqlLiveTail() *LeqlLiveTail`

NewLeqlLiveTail instantiates a new LeqlLiveTail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeqlLiveTailWithDefaults

`func NewLeqlLiveTailWithDefaults() *LeqlLiveTail`

NewLeqlLiveTailWithDefaults instantiates a new LeqlLiveTail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatement

`func (o *LeqlLiveTail) GetStatement() string`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *LeqlLiveTail) GetStatementOk() (*string, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *LeqlLiveTail) SetStatement(v string)`

SetStatement sets Statement field to given value.

### HasStatement

`func (o *LeqlLiveTail) HasStatement() bool`

HasStatement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


