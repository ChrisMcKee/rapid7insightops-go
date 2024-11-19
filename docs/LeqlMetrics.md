# LeqlMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statement** | Pointer to **string** | The &#39;where&#39; statement of the pre-computed query, for example, &#x60;\&quot;where(foo&gt;10 AND bar!&#x3D;0 OR /some_reg[ex]*_/)\&quot;&#x60;. | [optional] 
**Function** | Pointer to **string** | The &#39;calculate&#39; statement of the pre-computed query. Possible values: * &#x60;\&quot;calculate(count)\&quot;&#x60; * &#x60;\&quot;calculate(min: &lt;key&gt;)\&quot;&#x60; * &#x60;\&quot;calculate(max: &lt;key&gt;)\&quot;&#x60; * &#x60;\&quot;calculate(sum: &lt;key&gt;)\&quot;&#x60; * &#x60;\&quot;calculate(average: &lt;key&gt;)\&quot;&#x60; * &#x60;\&quot;calculate(bytes)\&quot;&#x60;  | [optional] 

## Methods

### NewLeqlMetrics

`func NewLeqlMetrics() *LeqlMetrics`

NewLeqlMetrics instantiates a new LeqlMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeqlMetricsWithDefaults

`func NewLeqlMetricsWithDefaults() *LeqlMetrics`

NewLeqlMetricsWithDefaults instantiates a new LeqlMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatement

`func (o *LeqlMetrics) GetStatement() string`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *LeqlMetrics) GetStatementOk() (*string, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *LeqlMetrics) SetStatement(v string)`

SetStatement sets Statement field to given value.

### HasStatement

`func (o *LeqlMetrics) HasStatement() bool`

HasStatement returns a boolean if a field has been set.

### GetFunction

`func (o *LeqlMetrics) GetFunction() string`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *LeqlMetrics) GetFunctionOk() (*string, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *LeqlMetrics) SetFunction(v string)`

SetFunction sets Function field to given value.

### HasFunction

`func (o *LeqlMetrics) HasFunction() bool`

HasFunction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


