# ListSavedQueries200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SavedQueries** | Pointer to [**[]SavedQuery**](SavedQuery.md) | All the saved queries for this account. | [optional] 

## Methods

### NewListSavedQueries200Response

`func NewListSavedQueries200Response() *ListSavedQueries200Response`

NewListSavedQueries200Response instantiates a new ListSavedQueries200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSavedQueries200ResponseWithDefaults

`func NewListSavedQueries200ResponseWithDefaults() *ListSavedQueries200Response`

NewListSavedQueries200ResponseWithDefaults instantiates a new ListSavedQueries200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSavedQueries

`func (o *ListSavedQueries200Response) GetSavedQueries() []SavedQuery`

GetSavedQueries returns the SavedQueries field if non-nil, zero value otherwise.

### GetSavedQueriesOk

`func (o *ListSavedQueries200Response) GetSavedQueriesOk() (*[]SavedQuery, bool)`

GetSavedQueriesOk returns a tuple with the SavedQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavedQueries

`func (o *ListSavedQueries200Response) SetSavedQueries(v []SavedQuery)`

SetSavedQueries sets SavedQueries field to given value.

### HasSavedQueries

`func (o *ListSavedQueries200Response) HasSavedQueries() bool`

HasSavedQueries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


