# TagSourceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The UUID of the log. | [optional] 
**RetentionPeriod** | Pointer to [**RetentionPeriodEnum**](RetentionPeriodEnum.md) |  | [optional] 
**Token** | Pointer to **string** | The log token(s) used for writing to the log. This only applies to token type logs (view the &#x60;source_type&#x60; parameter). | [optional] 
**Name** | Pointer to **string** | The name of the log. | [optional] 
**StoredDays** | Pointer to **interface{}** | The number of days stored. | [optional] 

## Methods

### NewTagSourceResponse

`func NewTagSourceResponse() *TagSourceResponse`

NewTagSourceResponse instantiates a new TagSourceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagSourceResponseWithDefaults

`func NewTagSourceResponseWithDefaults() *TagSourceResponse`

NewTagSourceResponseWithDefaults instantiates a new TagSourceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TagSourceResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TagSourceResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TagSourceResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TagSourceResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRetentionPeriod

`func (o *TagSourceResponse) GetRetentionPeriod() RetentionPeriodEnum`

GetRetentionPeriod returns the RetentionPeriod field if non-nil, zero value otherwise.

### GetRetentionPeriodOk

`func (o *TagSourceResponse) GetRetentionPeriodOk() (*RetentionPeriodEnum, bool)`

GetRetentionPeriodOk returns a tuple with the RetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriod

`func (o *TagSourceResponse) SetRetentionPeriod(v RetentionPeriodEnum)`

SetRetentionPeriod sets RetentionPeriod field to given value.

### HasRetentionPeriod

`func (o *TagSourceResponse) HasRetentionPeriod() bool`

HasRetentionPeriod returns a boolean if a field has been set.

### GetToken

`func (o *TagSourceResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TagSourceResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TagSourceResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TagSourceResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetName

`func (o *TagSourceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TagSourceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TagSourceResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TagSourceResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStoredDays

`func (o *TagSourceResponse) GetStoredDays() interface{}`

GetStoredDays returns the StoredDays field if non-nil, zero value otherwise.

### GetStoredDaysOk

`func (o *TagSourceResponse) GetStoredDaysOk() (*interface{}, bool)`

GetStoredDaysOk returns a tuple with the StoredDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoredDays

`func (o *TagSourceResponse) SetStoredDays(v interface{})`

SetStoredDays sets StoredDays field to given value.

### HasStoredDays

`func (o *TagSourceResponse) HasStoredDays() bool`

HasStoredDays returns a boolean if a field has been set.

### SetStoredDaysNil

`func (o *TagSourceResponse) SetStoredDaysNil(b bool)`

 SetStoredDaysNil sets the value for StoredDays to be an explicit nil

### UnsetStoredDays
`func (o *TagSourceResponse) UnsetStoredDays()`

UnsetStoredDays ensures that no value is present for StoredDays, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


