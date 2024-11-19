# AuditLogList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | Pointer to [**[]AuditLog**](AuditLog.md) |  | [optional] 

## Methods

### NewAuditLogList

`func NewAuditLogList() *AuditLogList`

NewAuditLogList instantiates a new AuditLogList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogListWithDefaults

`func NewAuditLogListWithDefaults() *AuditLogList`

NewAuditLogListWithDefaults instantiates a new AuditLogList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *AuditLogList) GetLogs() []AuditLog`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *AuditLogList) GetLogsOk() (*[]AuditLog, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *AuditLogList) SetLogs(v []AuditLog)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *AuditLogList) HasLogs() bool`

HasLogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


