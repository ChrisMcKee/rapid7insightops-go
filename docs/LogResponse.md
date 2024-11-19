# LogResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The UUID of the log. | 
**Name** | **interface{}** | The name of the log. | 
**Tokens** | **[]string** | The log token(s) used for writing to the log. This only applies to token type logs (view &#x60;source_type&#x60; parameter). | 
**Structures** | **[]string** | Structures are internal entities which may apply some additional processing to log entries written to this this log.  | 
**IpAddress** | Pointer to **string** | The IP address that the Log Search system receives log entries from. This only applies to syslog type logs (view the &#x60;source_type&#x60; parameter).  | [optional] 
**LogsetsInfo** | [**[]MemberInfoInner**](MemberInfoInner.md) | The information on each log set that this log is part of.  | 
**UserData** | **map[string]string** | A list of key-values pairs that may indicate some auxiliary information about the log. | 
**SourceType** | **string** | A categorization of logs which defines how log entries are received by a server. * Syslog type logs are associated with an IP address and a port (the values appear in the &#x60;user_data&#x60; field). Any log entries received by the server on that port and from that IP address will belong to that log. * Token type logs are associated with a log token. Any log entries received by the server on one of the [standard ports](https://docs.rapid7.com/insightops/token-tcp) containing that log token will belong to that log (with the log token removed from the log entry). * Internal type logs use internal mechanisms to receive log entries, for example any of the logs in the Internal Logs log set.  | [default to "token"]
**Rrn** | **string** | The Rapid7 Resource Name (RRN) of the log. The RRN is a unique identifier across the Rapid7 Platform. | 
**TokenSeed** | **interface{}** | The seed used to generate the log token if the log&#39;s &#x60;source_type&#x60; is &#x60;token&#x60;. | 
**RetentionPeriod** | [**RetentionPeriodEnum**](RetentionPeriodEnum.md) |  | 
**Links** | [**[]LinksInner**](LinksInner.md) |  | 

## Methods

### NewLogResponse

`func NewLogResponse(id string, name interface{}, tokens []string, structures []string, logsetsInfo []MemberInfoInner, userData map[string]string, sourceType string, rrn string, tokenSeed interface{}, retentionPeriod RetentionPeriodEnum, links []LinksInner, ) *LogResponse`

NewLogResponse instantiates a new LogResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogResponseWithDefaults

`func NewLogResponseWithDefaults() *LogResponse`

NewLogResponseWithDefaults instantiates a new LogResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LogResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *LogResponse) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogResponse) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LogResponse) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *LogResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *LogResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTokens

`func (o *LogResponse) GetTokens() []string`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *LogResponse) GetTokensOk() (*[]string, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *LogResponse) SetTokens(v []string)`

SetTokens sets Tokens field to given value.


### GetStructures

`func (o *LogResponse) GetStructures() []string`

GetStructures returns the Structures field if non-nil, zero value otherwise.

### GetStructuresOk

`func (o *LogResponse) GetStructuresOk() (*[]string, bool)`

GetStructuresOk returns a tuple with the Structures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructures

`func (o *LogResponse) SetStructures(v []string)`

SetStructures sets Structures field to given value.


### GetIpAddress

`func (o *LogResponse) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *LogResponse) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *LogResponse) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *LogResponse) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetLogsetsInfo

`func (o *LogResponse) GetLogsetsInfo() []MemberInfoInner`

GetLogsetsInfo returns the LogsetsInfo field if non-nil, zero value otherwise.

### GetLogsetsInfoOk

`func (o *LogResponse) GetLogsetsInfoOk() (*[]MemberInfoInner, bool)`

GetLogsetsInfoOk returns a tuple with the LogsetsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsetsInfo

`func (o *LogResponse) SetLogsetsInfo(v []MemberInfoInner)`

SetLogsetsInfo sets LogsetsInfo field to given value.


### GetUserData

`func (o *LogResponse) GetUserData() map[string]string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *LogResponse) GetUserDataOk() (*map[string]string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *LogResponse) SetUserData(v map[string]string)`

SetUserData sets UserData field to given value.


### GetSourceType

`func (o *LogResponse) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *LogResponse) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *LogResponse) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### GetRrn

`func (o *LogResponse) GetRrn() string`

GetRrn returns the Rrn field if non-nil, zero value otherwise.

### GetRrnOk

`func (o *LogResponse) GetRrnOk() (*string, bool)`

GetRrnOk returns a tuple with the Rrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrn

`func (o *LogResponse) SetRrn(v string)`

SetRrn sets Rrn field to given value.


### GetTokenSeed

`func (o *LogResponse) GetTokenSeed() interface{}`

GetTokenSeed returns the TokenSeed field if non-nil, zero value otherwise.

### GetTokenSeedOk

`func (o *LogResponse) GetTokenSeedOk() (*interface{}, bool)`

GetTokenSeedOk returns a tuple with the TokenSeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenSeed

`func (o *LogResponse) SetTokenSeed(v interface{})`

SetTokenSeed sets TokenSeed field to given value.


### SetTokenSeedNil

`func (o *LogResponse) SetTokenSeedNil(b bool)`

 SetTokenSeedNil sets the value for TokenSeed to be an explicit nil

### UnsetTokenSeed
`func (o *LogResponse) UnsetTokenSeed()`

UnsetTokenSeed ensures that no value is present for TokenSeed, not even an explicit nil
### GetRetentionPeriod

`func (o *LogResponse) GetRetentionPeriod() RetentionPeriodEnum`

GetRetentionPeriod returns the RetentionPeriod field if non-nil, zero value otherwise.

### GetRetentionPeriodOk

`func (o *LogResponse) GetRetentionPeriodOk() (*RetentionPeriodEnum, bool)`

GetRetentionPeriodOk returns a tuple with the RetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriod

`func (o *LogResponse) SetRetentionPeriod(v RetentionPeriodEnum)`

SetRetentionPeriod sets RetentionPeriod field to given value.


### GetLinks

`func (o *LogResponse) GetLinks() []LinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *LogResponse) GetLinksOk() (*[]LinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *LogResponse) SetLinks(v []LinksInner)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


