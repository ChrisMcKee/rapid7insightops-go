# AuditLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the log. | [optional] 
**Name** | Pointer to **string** | The name of the log. | [optional] 
**Rrn** | Pointer to **string** | The Rapid7 Resource Name of the log (unique identifier across the Rapid7 Platform). | [optional] 
**LogsetsInfo** | Pointer to [**[]MemberInfoInner**](MemberInfoInner.md) | The information on each log set that this log is part of.  | [optional] 

## Methods

### NewAuditLog

`func NewAuditLog() *AuditLog`

NewAuditLog instantiates a new AuditLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogWithDefaults

`func NewAuditLogWithDefaults() *AuditLog`

NewAuditLogWithDefaults instantiates a new AuditLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuditLog) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditLog) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditLog) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuditLog) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AuditLog) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuditLog) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuditLog) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuditLog) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRrn

`func (o *AuditLog) GetRrn() string`

GetRrn returns the Rrn field if non-nil, zero value otherwise.

### GetRrnOk

`func (o *AuditLog) GetRrnOk() (*string, bool)`

GetRrnOk returns a tuple with the Rrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrn

`func (o *AuditLog) SetRrn(v string)`

SetRrn sets Rrn field to given value.

### HasRrn

`func (o *AuditLog) HasRrn() bool`

HasRrn returns a boolean if a field has been set.

### GetLogsetsInfo

`func (o *AuditLog) GetLogsetsInfo() []MemberInfoInner`

GetLogsetsInfo returns the LogsetsInfo field if non-nil, zero value otherwise.

### GetLogsetsInfoOk

`func (o *AuditLog) GetLogsetsInfoOk() (*[]MemberInfoInner, bool)`

GetLogsetsInfoOk returns a tuple with the LogsetsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsetsInfo

`func (o *AuditLog) SetLogsetsInfo(v []MemberInfoInner)`

SetLogsetsInfo sets LogsetsInfo field to given value.

### HasLogsetsInfo

`func (o *AuditLog) HasLogsetsInfo() bool`

HasLogsetsInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


