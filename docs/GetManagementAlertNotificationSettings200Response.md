# GetManagementAlertNotificationSettings200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]AlertNotificationSetting**](AlertNotificationSetting.md) | All the notifications for this account that can be attached to basic detection rules. | [optional] 

## Methods

### NewGetManagementAlertNotificationSettings200Response

`func NewGetManagementAlertNotificationSettings200Response() *GetManagementAlertNotificationSettings200Response`

NewGetManagementAlertNotificationSettings200Response instantiates a new GetManagementAlertNotificationSettings200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetManagementAlertNotificationSettings200ResponseWithDefaults

`func NewGetManagementAlertNotificationSettings200ResponseWithDefaults() *GetManagementAlertNotificationSettings200Response`

NewGetManagementAlertNotificationSettings200ResponseWithDefaults instantiates a new GetManagementAlertNotificationSettings200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *GetManagementAlertNotificationSettings200Response) GetActions() []AlertNotificationSetting`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *GetManagementAlertNotificationSettings200Response) GetActionsOk() (*[]AlertNotificationSetting, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *GetManagementAlertNotificationSettings200Response) SetActions(v []AlertNotificationSetting)`

SetActions sets Actions field to given value.

### HasActions

`func (o *GetManagementAlertNotificationSettings200Response) HasActions() bool`

HasActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


