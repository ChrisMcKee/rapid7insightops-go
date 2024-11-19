# SavedQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique id of the saved query. | 
**Name** | **string** | The name of the saved query. | 
**Leql** | Pointer to [**Leql**](Leql.md) |  | [optional] 
**Logs** | Pointer to **[]string** | The log keys of the logs which the query is run against. | [optional] 

## Methods

### NewSavedQuery

`func NewSavedQuery(id string, name string, ) *SavedQuery`

NewSavedQuery instantiates a new SavedQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSavedQueryWithDefaults

`func NewSavedQueryWithDefaults() *SavedQuery`

NewSavedQueryWithDefaults instantiates a new SavedQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SavedQuery) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SavedQuery) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SavedQuery) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SavedQuery) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SavedQuery) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SavedQuery) SetName(v string)`

SetName sets Name field to given value.


### GetLeql

`func (o *SavedQuery) GetLeql() Leql`

GetLeql returns the Leql field if non-nil, zero value otherwise.

### GetLeqlOk

`func (o *SavedQuery) GetLeqlOk() (*Leql, bool)`

GetLeqlOk returns a tuple with the Leql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeql

`func (o *SavedQuery) SetLeql(v Leql)`

SetLeql sets Leql field to given value.

### HasLeql

`func (o *SavedQuery) HasLeql() bool`

HasLeql returns a boolean if a field has been set.

### GetLogs

`func (o *SavedQuery) GetLogs() []string`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *SavedQuery) GetLogsOk() (*[]string, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *SavedQuery) SetLogs(v []string)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *SavedQuery) HasLogs() bool`

HasLogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


