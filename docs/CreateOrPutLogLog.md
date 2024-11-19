# CreateOrPutLogLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **interface{}** | The name of the log. | 
**Tokens** | Pointer to **[]string** | The log token(s) used for writing to the log. This only applies to token type logs (view the &#x60;source_type&#x60; parameter). | [optional] 
**Structures** | Pointer to **[]string** | Structures are internal entities which may apply some additional processing to log entries written to this this log.  | [optional] 
**IpAddress** | Pointer to **string** | The IP address that the Log Search system receives log entries from. This only applies to syslog type logs (view the &#x60;source_type&#x60; parameter). | [optional] 
**LogsetsInfo** | Pointer to [**[]CreateOrPutLogLogLogsetsInfoInner**](CreateOrPutLogLogLogsetsInfoInner.md) | The information on each log set that this log is part of.  | [optional] 
**UserData** | Pointer to **map[string]interface{}** | A list of key-value pairs that may indicate some auxiliary information about the log. | [optional] 
**SourceType** | Pointer to **string** | A categorization of logs that defines how log entries are received by servers. * &#x60;syslog&#x60; type logs are associated with an IP address and a port (the values appear in the &#x60;user_data&#x60; field). Any log entries received by the server on that port and from that IP address will belong to that log. * &#x60;token&#x60; type logs are associated with a log token. Any log entries received by the server on one of the [standard ports](https://docs.rapid7.com/insightops/token-tcp) containing that log token will belong to that log (with the log token removed from the log entry). * &#x60;internal&#x60; type logs use internal mechanisms to receive log entries, for example any of the logs in the Internal Logs log set.  | [optional] [default to "token"]
**TokenSeed** | Pointer to **string** | The seed used to generate the log token (for token type logs). | [optional] 

## Methods

### NewCreateOrPutLogLog

`func NewCreateOrPutLogLog(name interface{}, ) *CreateOrPutLogLog`

NewCreateOrPutLogLog instantiates a new CreateOrPutLogLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrPutLogLogWithDefaults

`func NewCreateOrPutLogLogWithDefaults() *CreateOrPutLogLog`

NewCreateOrPutLogLogWithDefaults instantiates a new CreateOrPutLogLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateOrPutLogLog) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOrPutLogLog) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOrPutLogLog) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *CreateOrPutLogLog) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateOrPutLogLog) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTokens

`func (o *CreateOrPutLogLog) GetTokens() []string`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *CreateOrPutLogLog) GetTokensOk() (*[]string, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *CreateOrPutLogLog) SetTokens(v []string)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *CreateOrPutLogLog) HasTokens() bool`

HasTokens returns a boolean if a field has been set.

### GetStructures

`func (o *CreateOrPutLogLog) GetStructures() []string`

GetStructures returns the Structures field if non-nil, zero value otherwise.

### GetStructuresOk

`func (o *CreateOrPutLogLog) GetStructuresOk() (*[]string, bool)`

GetStructuresOk returns a tuple with the Structures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructures

`func (o *CreateOrPutLogLog) SetStructures(v []string)`

SetStructures sets Structures field to given value.

### HasStructures

`func (o *CreateOrPutLogLog) HasStructures() bool`

HasStructures returns a boolean if a field has been set.

### GetIpAddress

`func (o *CreateOrPutLogLog) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *CreateOrPutLogLog) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *CreateOrPutLogLog) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *CreateOrPutLogLog) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetLogsetsInfo

`func (o *CreateOrPutLogLog) GetLogsetsInfo() []CreateOrPutLogLogLogsetsInfoInner`

GetLogsetsInfo returns the LogsetsInfo field if non-nil, zero value otherwise.

### GetLogsetsInfoOk

`func (o *CreateOrPutLogLog) GetLogsetsInfoOk() (*[]CreateOrPutLogLogLogsetsInfoInner, bool)`

GetLogsetsInfoOk returns a tuple with the LogsetsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsetsInfo

`func (o *CreateOrPutLogLog) SetLogsetsInfo(v []CreateOrPutLogLogLogsetsInfoInner)`

SetLogsetsInfo sets LogsetsInfo field to given value.

### HasLogsetsInfo

`func (o *CreateOrPutLogLog) HasLogsetsInfo() bool`

HasLogsetsInfo returns a boolean if a field has been set.

### GetUserData

`func (o *CreateOrPutLogLog) GetUserData() map[string]interface{}`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *CreateOrPutLogLog) GetUserDataOk() (*map[string]interface{}, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *CreateOrPutLogLog) SetUserData(v map[string]interface{})`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *CreateOrPutLogLog) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### GetSourceType

`func (o *CreateOrPutLogLog) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *CreateOrPutLogLog) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *CreateOrPutLogLog) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *CreateOrPutLogLog) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetTokenSeed

`func (o *CreateOrPutLogLog) GetTokenSeed() string`

GetTokenSeed returns the TokenSeed field if non-nil, zero value otherwise.

### GetTokenSeedOk

`func (o *CreateOrPutLogLog) GetTokenSeedOk() (*string, bool)`

GetTokenSeedOk returns a tuple with the TokenSeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenSeed

`func (o *CreateOrPutLogLog) SetTokenSeed(v string)`

SetTokenSeed sets TokenSeed field to given value.

### HasTokenSeed

`func (o *CreateOrPutLogLog) HasTokenSeed() bool`

HasTokenSeed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


