# AccountsUsageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DailyUsage** | Pointer to [**AccountsUsageResponseDailyUsage**](AccountsUsageResponseDailyUsage.md) |  | [optional] 
**PeriodUsage** | Pointer to **int** | The total number of bytes written to every log, for the entire time range. | [optional] 
**Period** | Pointer to [**AccountsUsageResponsePeriod**](AccountsUsageResponsePeriod.md) |  | [optional] 
**Id** | Pointer to **string** | The account resource ID (UUID). A unique identifier for your organization. | [optional] 
**Name** | Pointer to **string** | The name of your organization. | [optional] 

## Methods

### NewAccountsUsageResponse

`func NewAccountsUsageResponse() *AccountsUsageResponse`

NewAccountsUsageResponse instantiates a new AccountsUsageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountsUsageResponseWithDefaults

`func NewAccountsUsageResponseWithDefaults() *AccountsUsageResponse`

NewAccountsUsageResponseWithDefaults instantiates a new AccountsUsageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDailyUsage

`func (o *AccountsUsageResponse) GetDailyUsage() AccountsUsageResponseDailyUsage`

GetDailyUsage returns the DailyUsage field if non-nil, zero value otherwise.

### GetDailyUsageOk

`func (o *AccountsUsageResponse) GetDailyUsageOk() (*AccountsUsageResponseDailyUsage, bool)`

GetDailyUsageOk returns a tuple with the DailyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyUsage

`func (o *AccountsUsageResponse) SetDailyUsage(v AccountsUsageResponseDailyUsage)`

SetDailyUsage sets DailyUsage field to given value.

### HasDailyUsage

`func (o *AccountsUsageResponse) HasDailyUsage() bool`

HasDailyUsage returns a boolean if a field has been set.

### GetPeriodUsage

`func (o *AccountsUsageResponse) GetPeriodUsage() int`

GetPeriodUsage returns the PeriodUsage field if non-nil, zero value otherwise.

### GetPeriodUsageOk

`func (o *AccountsUsageResponse) GetPeriodUsageOk() (*int, bool)`

GetPeriodUsageOk returns a tuple with the PeriodUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodUsage

`func (o *AccountsUsageResponse) SetPeriodUsage(v int)`

SetPeriodUsage sets PeriodUsage field to given value.

### HasPeriodUsage

`func (o *AccountsUsageResponse) HasPeriodUsage() bool`

HasPeriodUsage returns a boolean if a field has been set.

### GetPeriod

`func (o *AccountsUsageResponse) GetPeriod() AccountsUsageResponsePeriod`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *AccountsUsageResponse) GetPeriodOk() (*AccountsUsageResponsePeriod, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *AccountsUsageResponse) SetPeriod(v AccountsUsageResponsePeriod)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *AccountsUsageResponse) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetId

`func (o *AccountsUsageResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountsUsageResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountsUsageResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountsUsageResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AccountsUsageResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountsUsageResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountsUsageResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountsUsageResponse) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


