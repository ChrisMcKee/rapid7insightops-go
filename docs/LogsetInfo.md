# LogsetInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The UUID of the log set. | 
**Name** | **interface{}** | The name of the log set. | 
**Description** | **string** | The description of the log set. | 
**UserData** | **map[string]interface{}** | A list of key-value pairs that may indicate some auxiliary information about the log set. | 
**LogsInfo** | [**[]MemberInfoLogsetInner**](MemberInfoLogsetInner.md) | Information on each log that this log set contains.  | 
**Rrn** | **interface{}** | The Rapid7 Resource Name (RRN) of the log set. The RRN is a unique identifier across the Rapid7 Platform. | 

## Methods

### NewLogsetInfo

`func NewLogsetInfo(id string, name interface{}, description string, userData map[string]interface{}, logsInfo []MemberInfoLogsetInner, rrn interface{}, ) *LogsetInfo`

NewLogsetInfo instantiates a new LogsetInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsetInfoWithDefaults

`func NewLogsetInfoWithDefaults() *LogsetInfo`

NewLogsetInfoWithDefaults instantiates a new LogsetInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LogsetInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogsetInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogsetInfo) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *LogsetInfo) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogsetInfo) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LogsetInfo) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *LogsetInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *LogsetInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *LogsetInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LogsetInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LogsetInfo) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetUserData

`func (o *LogsetInfo) GetUserData() map[string]interface{}`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *LogsetInfo) GetUserDataOk() (*map[string]interface{}, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *LogsetInfo) SetUserData(v map[string]interface{})`

SetUserData sets UserData field to given value.


### GetLogsInfo

`func (o *LogsetInfo) GetLogsInfo() []MemberInfoLogsetInner`

GetLogsInfo returns the LogsInfo field if non-nil, zero value otherwise.

### GetLogsInfoOk

`func (o *LogsetInfo) GetLogsInfoOk() (*[]MemberInfoLogsetInner, bool)`

GetLogsInfoOk returns a tuple with the LogsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsInfo

`func (o *LogsetInfo) SetLogsInfo(v []MemberInfoLogsetInner)`

SetLogsInfo sets LogsInfo field to given value.


### GetRrn

`func (o *LogsetInfo) GetRrn() interface{}`

GetRrn returns the Rrn field if non-nil, zero value otherwise.

### GetRrnOk

`func (o *LogsetInfo) GetRrnOk() (*interface{}, bool)`

GetRrnOk returns a tuple with the Rrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrn

`func (o *LogsetInfo) SetRrn(v interface{})`

SetRrn sets Rrn field to given value.


### SetRrnNil

`func (o *LogsetInfo) SetRrnNil(b bool)`

 SetRrnNil sets the value for Rrn to be an explicit nil

### UnsetRrn
`func (o *LogsetInfo) UnsetRrn()`

UnsetRrn ensures that no value is present for Rrn, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


