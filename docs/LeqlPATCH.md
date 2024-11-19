# LeqlPATCH

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**During** | Pointer to [**LeqlPATCHDuring**](LeqlPATCHDuring.md) |  | [optional] 
**Statement** | Pointer to **string** | \\ The LEQL query run against the log(s). If empty, the query retrieves all log entries in the specified time range.  Cannot be empty for saved queries.  | [optional] 

## Methods

### NewLeqlPATCH

`func NewLeqlPATCH() *LeqlPATCH`

NewLeqlPATCH instantiates a new LeqlPATCH object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeqlPATCHWithDefaults

`func NewLeqlPATCHWithDefaults() *LeqlPATCH`

NewLeqlPATCHWithDefaults instantiates a new LeqlPATCH object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuring

`func (o *LeqlPATCH) GetDuring() LeqlPATCHDuring`

GetDuring returns the During field if non-nil, zero value otherwise.

### GetDuringOk

`func (o *LeqlPATCH) GetDuringOk() (*LeqlPATCHDuring, bool)`

GetDuringOk returns a tuple with the During field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuring

`func (o *LeqlPATCH) SetDuring(v LeqlPATCHDuring)`

SetDuring sets During field to given value.

### HasDuring

`func (o *LeqlPATCH) HasDuring() bool`

HasDuring returns a boolean if a field has been set.

### GetStatement

`func (o *LeqlPATCH) GetStatement() string`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *LeqlPATCH) GetStatementOk() (*string, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *LeqlPATCH) SetStatement(v string)`

SetStatement sets Statement field to given value.

### HasStatement

`func (o *LeqlPATCH) HasStatement() bool`

HasStatement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


