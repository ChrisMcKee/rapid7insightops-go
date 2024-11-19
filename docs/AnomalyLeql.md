# AnomalyLeql

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statement** | Pointer to **string** | The &#x60;where()&#x60; clause of the [LEQL](https://docs.rapid7.com/insightidr/log-search/#write-a-leql-query) query, for example, &#x60;where(foo)&#x60;. If empty, the query matches all log entries. | [optional] 
**Function** | [**FunctionDef**](FunctionDef.md) |  | 

## Methods

### NewAnomalyLeql

`func NewAnomalyLeql(function FunctionDef, ) *AnomalyLeql`

NewAnomalyLeql instantiates a new AnomalyLeql object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnomalyLeqlWithDefaults

`func NewAnomalyLeqlWithDefaults() *AnomalyLeql`

NewAnomalyLeqlWithDefaults instantiates a new AnomalyLeql object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatement

`func (o *AnomalyLeql) GetStatement() string`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *AnomalyLeql) GetStatementOk() (*string, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *AnomalyLeql) SetStatement(v string)`

SetStatement sets Statement field to given value.

### HasStatement

`func (o *AnomalyLeql) HasStatement() bool`

HasStatement returns a boolean if a field has been set.

### GetFunction

`func (o *AnomalyLeql) GetFunction() FunctionDef`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *AnomalyLeql) GetFunctionOk() (*FunctionDef, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *AnomalyLeql) SetFunction(v FunctionDef)`

SetFunction sets Function field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


