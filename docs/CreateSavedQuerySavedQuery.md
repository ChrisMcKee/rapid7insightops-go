# CreateSavedQuerySavedQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name for the Saved Query (doesn&#39;t need to be unique). | 
**Leql** | Pointer to [**Leql**](Leql.md) |  | [optional] 
**Logs** | Pointer to **[]string** | The log keys of the logs which the query is run against. | [optional] 

## Methods

### NewCreateSavedQuerySavedQuery

`func NewCreateSavedQuerySavedQuery(name string, ) *CreateSavedQuerySavedQuery`

NewCreateSavedQuerySavedQuery instantiates a new CreateSavedQuerySavedQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSavedQuerySavedQueryWithDefaults

`func NewCreateSavedQuerySavedQueryWithDefaults() *CreateSavedQuerySavedQuery`

NewCreateSavedQuerySavedQueryWithDefaults instantiates a new CreateSavedQuerySavedQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateSavedQuerySavedQuery) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSavedQuerySavedQuery) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSavedQuerySavedQuery) SetName(v string)`

SetName sets Name field to given value.


### GetLeql

`func (o *CreateSavedQuerySavedQuery) GetLeql() Leql`

GetLeql returns the Leql field if non-nil, zero value otherwise.

### GetLeqlOk

`func (o *CreateSavedQuerySavedQuery) GetLeqlOk() (*Leql, bool)`

GetLeqlOk returns a tuple with the Leql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeql

`func (o *CreateSavedQuerySavedQuery) SetLeql(v Leql)`

SetLeql sets Leql field to given value.

### HasLeql

`func (o *CreateSavedQuerySavedQuery) HasLeql() bool`

HasLeql returns a boolean if a field has been set.

### GetLogs

`func (o *CreateSavedQuerySavedQuery) GetLogs() []string`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *CreateSavedQuerySavedQuery) GetLogsOk() (*[]string, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *CreateSavedQuerySavedQuery) SetLogs(v []string)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *CreateSavedQuerySavedQuery) HasLogs() bool`

HasLogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


