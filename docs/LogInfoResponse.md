# LogInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The UUID of the log. | 
**Name** | **interface{}** | The name of the log. | 
**Tokens** | Pointer to **[]string** | The log token(s) used for writing to the log. This only applies to token type logs (view the &#x60;source_type&#x60; parameter). | [optional] 
**Structures** | Pointer to **[]string** | Structures are internal entities which may apply some additional processing to log entries written to this log.  | [optional] 
**IpAddress** | Pointer to **string** | The IP address that the Log Search system receives log entries from. This only applies to syslog type logs (view the &#x60;source_type&#x60; parameter). | [optional] 
**LogsetsInfo** | Pointer to [**[]MemberInfoInner**](MemberInfoInner.md) | The information on each log set that this log is part of.  | [optional] 
**UserData** | Pointer to **map[string]string** | A list of key-values pairs that may indicate some auxiliary information about the log. | [optional] 
**SourceType** | Pointer to **string** | A categorization of logs which defines how log entries are received by a server. * Syslog type logs are associated with an IP address and a port (the values appear in the &#x60;user_data&#x60; field). Any log entries received by the server on that port and from that IP address will belong to that log. * Token type logs are associated with a log token. Any log entries received by the server on one of the [standard ports](https://docs.rapid7.com/insightops/token-tcp) containing that log token will belong to that log (with the log token removed from the log entry). * Internal type logs use internal mechanisms to receive log entries, for example any of the logs in the Internal Logs log set.  | [optional] [default to "token"]
**Rrn** | Pointer to **string** | The Rapid7 Resource Name (RRN) of the log. The RRN is a unique identifier across the Rapid7 Platform. | [optional] 
**TokenSeed** | Pointer to **interface{}** | The seed used to generate the log token (if the log&#39;s source type is &#x60;\&quot;token\&quot;&#x60;). | [optional] 
**RetentionPeriod** | Pointer to [**RetentionPeriodEnum**](RetentionPeriodEnum.md) |  | [optional] 
**Links** | [**[]LinksInner**](LinksInner.md) |  | 

## Methods

### NewLogInfoResponse

`func NewLogInfoResponse(id string, name interface{}, links []LinksInner, ) *LogInfoResponse`

NewLogInfoResponse instantiates a new LogInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogInfoResponseWithDefaults

`func NewLogInfoResponseWithDefaults() *LogInfoResponse`

NewLogInfoResponseWithDefaults instantiates a new LogInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LogInfoResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogInfoResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogInfoResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *LogInfoResponse) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogInfoResponse) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LogInfoResponse) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *LogInfoResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *LogInfoResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTokens

`func (o *LogInfoResponse) GetTokens() []string`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *LogInfoResponse) GetTokensOk() (*[]string, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *LogInfoResponse) SetTokens(v []string)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *LogInfoResponse) HasTokens() bool`

HasTokens returns a boolean if a field has been set.

### GetStructures

`func (o *LogInfoResponse) GetStructures() []string`

GetStructures returns the Structures field if non-nil, zero value otherwise.

### GetStructuresOk

`func (o *LogInfoResponse) GetStructuresOk() (*[]string, bool)`

GetStructuresOk returns a tuple with the Structures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructures

`func (o *LogInfoResponse) SetStructures(v []string)`

SetStructures sets Structures field to given value.

### HasStructures

`func (o *LogInfoResponse) HasStructures() bool`

HasStructures returns a boolean if a field has been set.

### GetIpAddress

`func (o *LogInfoResponse) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *LogInfoResponse) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *LogInfoResponse) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *LogInfoResponse) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetLogsetsInfo

`func (o *LogInfoResponse) GetLogsetsInfo() []MemberInfoInner`

GetLogsetsInfo returns the LogsetsInfo field if non-nil, zero value otherwise.

### GetLogsetsInfoOk

`func (o *LogInfoResponse) GetLogsetsInfoOk() (*[]MemberInfoInner, bool)`

GetLogsetsInfoOk returns a tuple with the LogsetsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsetsInfo

`func (o *LogInfoResponse) SetLogsetsInfo(v []MemberInfoInner)`

SetLogsetsInfo sets LogsetsInfo field to given value.

### HasLogsetsInfo

`func (o *LogInfoResponse) HasLogsetsInfo() bool`

HasLogsetsInfo returns a boolean if a field has been set.

### GetUserData

`func (o *LogInfoResponse) GetUserData() map[string]string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *LogInfoResponse) GetUserDataOk() (*map[string]string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *LogInfoResponse) SetUserData(v map[string]string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *LogInfoResponse) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### GetSourceType

`func (o *LogInfoResponse) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *LogInfoResponse) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *LogInfoResponse) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *LogInfoResponse) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetRrn

`func (o *LogInfoResponse) GetRrn() string`

GetRrn returns the Rrn field if non-nil, zero value otherwise.

### GetRrnOk

`func (o *LogInfoResponse) GetRrnOk() (*string, bool)`

GetRrnOk returns a tuple with the Rrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrn

`func (o *LogInfoResponse) SetRrn(v string)`

SetRrn sets Rrn field to given value.

### HasRrn

`func (o *LogInfoResponse) HasRrn() bool`

HasRrn returns a boolean if a field has been set.

### GetTokenSeed

`func (o *LogInfoResponse) GetTokenSeed() interface{}`

GetTokenSeed returns the TokenSeed field if non-nil, zero value otherwise.

### GetTokenSeedOk

`func (o *LogInfoResponse) GetTokenSeedOk() (*interface{}, bool)`

GetTokenSeedOk returns a tuple with the TokenSeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenSeed

`func (o *LogInfoResponse) SetTokenSeed(v interface{})`

SetTokenSeed sets TokenSeed field to given value.

### HasTokenSeed

`func (o *LogInfoResponse) HasTokenSeed() bool`

HasTokenSeed returns a boolean if a field has been set.

### SetTokenSeedNil

`func (o *LogInfoResponse) SetTokenSeedNil(b bool)`

 SetTokenSeedNil sets the value for TokenSeed to be an explicit nil

### UnsetTokenSeed
`func (o *LogInfoResponse) UnsetTokenSeed()`

UnsetTokenSeed ensures that no value is present for TokenSeed, not even an explicit nil
### GetRetentionPeriod

`func (o *LogInfoResponse) GetRetentionPeriod() RetentionPeriodEnum`

GetRetentionPeriod returns the RetentionPeriod field if non-nil, zero value otherwise.

### GetRetentionPeriodOk

`func (o *LogInfoResponse) GetRetentionPeriodOk() (*RetentionPeriodEnum, bool)`

GetRetentionPeriodOk returns a tuple with the RetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriod

`func (o *LogInfoResponse) SetRetentionPeriod(v RetentionPeriodEnum)`

SetRetentionPeriod sets RetentionPeriod field to given value.

### HasRetentionPeriod

`func (o *LogInfoResponse) HasRetentionPeriod() bool`

HasRetentionPeriod returns a boolean if a field has been set.

### GetLinks

`func (o *LogInfoResponse) GetLinks() []LinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *LogInfoResponse) GetLinksOk() (*[]LinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *LogInfoResponse) SetLinks(v []LinksInner)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


