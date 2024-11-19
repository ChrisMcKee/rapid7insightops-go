# PatchSavedQuerySavedQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name for the Saved Query (doesn&#39;t need to be unique). | [optional] 
**Leql** | Pointer to [**LeqlPATCH**](LeqlPATCH.md) |  | [optional] 
**Logs** | Pointer to **[]string** | The log keys of the logs which the query is run against. | [optional] 

## Methods

### NewPatchSavedQuerySavedQuery

`func NewPatchSavedQuerySavedQuery() *PatchSavedQuerySavedQuery`

NewPatchSavedQuerySavedQuery instantiates a new PatchSavedQuerySavedQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchSavedQuerySavedQueryWithDefaults

`func NewPatchSavedQuerySavedQueryWithDefaults() *PatchSavedQuerySavedQuery`

NewPatchSavedQuerySavedQueryWithDefaults instantiates a new PatchSavedQuerySavedQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchSavedQuerySavedQuery) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchSavedQuerySavedQuery) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchSavedQuerySavedQuery) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchSavedQuerySavedQuery) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLeql

`func (o *PatchSavedQuerySavedQuery) GetLeql() LeqlPATCH`

GetLeql returns the Leql field if non-nil, zero value otherwise.

### GetLeqlOk

`func (o *PatchSavedQuerySavedQuery) GetLeqlOk() (*LeqlPATCH, bool)`

GetLeqlOk returns a tuple with the Leql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeql

`func (o *PatchSavedQuerySavedQuery) SetLeql(v LeqlPATCH)`

SetLeql sets Leql field to given value.

### HasLeql

`func (o *PatchSavedQuerySavedQuery) HasLeql() bool`

HasLeql returns a boolean if a field has been set.

### GetLogs

`func (o *PatchSavedQuerySavedQuery) GetLogs() []string`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *PatchSavedQuerySavedQuery) GetLogsOk() (*[]string, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *PatchSavedQuerySavedQuery) SetLogs(v []string)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *PatchSavedQuerySavedQuery) HasLogs() bool`

HasLogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


