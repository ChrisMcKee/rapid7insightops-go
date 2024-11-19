# SearchStatsInnerLeql

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statement** | Pointer to **string** | The LEQL query run against the log(s), for example, &#x60;where(foo&gt;10 AND bar!&#x3D;0 OR /some_reg[ex]*_/)&#x60;. If empty, the query retrieves all log entries in the specified time range. | [optional] 
**From** | Pointer to **int32** | The start of the time range for the query, as a UNIX timestamp in milliseconds. | [optional] 
**To** | Pointer to **int32** | The end of the time range for the query, as a UNIX timestamp in milliseconds. | [optional] 
**NumOfLogs** | Pointer to **int32** | The number of logs searched by this query. | [optional] 

## Methods

### NewSearchStatsInnerLeql

`func NewSearchStatsInnerLeql() *SearchStatsInnerLeql`

NewSearchStatsInnerLeql instantiates a new SearchStatsInnerLeql object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchStatsInnerLeqlWithDefaults

`func NewSearchStatsInnerLeqlWithDefaults() *SearchStatsInnerLeql`

NewSearchStatsInnerLeqlWithDefaults instantiates a new SearchStatsInnerLeql object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatement

`func (o *SearchStatsInnerLeql) GetStatement() string`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *SearchStatsInnerLeql) GetStatementOk() (*string, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *SearchStatsInnerLeql) SetStatement(v string)`

SetStatement sets Statement field to given value.

### HasStatement

`func (o *SearchStatsInnerLeql) HasStatement() bool`

HasStatement returns a boolean if a field has been set.

### GetFrom

`func (o *SearchStatsInnerLeql) GetFrom() int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *SearchStatsInnerLeql) GetFromOk() (*int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *SearchStatsInnerLeql) SetFrom(v int32)`

SetFrom sets From field to given value.

### HasFrom

`func (o *SearchStatsInnerLeql) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *SearchStatsInnerLeql) GetTo() int32`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *SearchStatsInnerLeql) GetToOk() (*int32, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *SearchStatsInnerLeql) SetTo(v int32)`

SetTo sets To field to given value.

### HasTo

`func (o *SearchStatsInnerLeql) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetNumOfLogs

`func (o *SearchStatsInnerLeql) GetNumOfLogs() int32`

GetNumOfLogs returns the NumOfLogs field if non-nil, zero value otherwise.

### GetNumOfLogsOk

`func (o *SearchStatsInnerLeql) GetNumOfLogsOk() (*int32, bool)`

GetNumOfLogsOk returns a tuple with the NumOfLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfLogs

`func (o *SearchStatsInnerLeql) SetNumOfLogs(v int32)`

SetNumOfLogs sets NumOfLogs field to given value.

### HasNumOfLogs

`func (o *SearchStatsInnerLeql) HasNumOfLogs() bool`

HasNumOfLogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


