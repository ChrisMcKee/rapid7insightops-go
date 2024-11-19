# InactivityLeql

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statement** | **string** | The &#x60;where()&#x60; clause of the [LEQL](https://docs.rapid7.com/insightidr/log-search/#write-a-leql-query) query, for example, &#x60;where(foo)&#x60;. If empty, the detection rule matches all log entries. | 

## Methods

### NewInactivityLeql

`func NewInactivityLeql(statement string, ) *InactivityLeql`

NewInactivityLeql instantiates a new InactivityLeql object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInactivityLeqlWithDefaults

`func NewInactivityLeqlWithDefaults() *InactivityLeql`

NewInactivityLeqlWithDefaults instantiates a new InactivityLeql object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatement

`func (o *InactivityLeql) GetStatement() string`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *InactivityLeql) GetStatementOk() (*string, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *InactivityLeql) SetStatement(v string)`

SetStatement sets Statement field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


