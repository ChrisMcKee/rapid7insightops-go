# ScheduledQueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the scheduled query associated with this change detection rule. | 
**Name** | Pointer to **string** | The name of the scheduled query. | [optional] 
**Leql** | Pointer to [**AnomalyLeql**](AnomalyLeql.md) |  | [optional] 
**TimeValue** | Pointer to **int32** | Defines the timerange for the two queries, along with the &#x60;time_period&#x60; parameter.  Two queries will be run: the timerange of the first, more recent query, will end at the current system time;  the timerange of the older query will end at the current system time minus the time specified by the &#x60;timeframe_value&#x60; &amp; &#x60;timeframe_period&#x60; parameters.  | [optional] 
**TimePeriod** | Pointer to **string** | Defines the timerange for the two queries, along with the &#x60;time_value&#x60; parameter. | [optional] 
**TimeframeValue** | Pointer to **int32** | Defines the \&quot;offset\&quot; for the older of the two queries along with the &#x60;timeframe_period&#x60; parameter. For example, if the &#x60;timeframe_value&#x60; is set to 1 and &#x60;timeframe_period&#x60; is set to &#x60;Week&#x60;, then the time range for the older of the two queries will end at the current system time minus 1 week. | [optional] 
**TimeframePeriod** | Pointer to **string** | Defines the \&quot;offset\&quot; for the older of the two queries, along with the &#x60;timeframe_value&#x60; parameter. | [optional] 

## Methods

### NewScheduledQueryResponse

`func NewScheduledQueryResponse(id string, ) *ScheduledQueryResponse`

NewScheduledQueryResponse instantiates a new ScheduledQueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduledQueryResponseWithDefaults

`func NewScheduledQueryResponseWithDefaults() *ScheduledQueryResponse`

NewScheduledQueryResponseWithDefaults instantiates a new ScheduledQueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScheduledQueryResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScheduledQueryResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScheduledQueryResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ScheduledQueryResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScheduledQueryResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScheduledQueryResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ScheduledQueryResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLeql

`func (o *ScheduledQueryResponse) GetLeql() AnomalyLeql`

GetLeql returns the Leql field if non-nil, zero value otherwise.

### GetLeqlOk

`func (o *ScheduledQueryResponse) GetLeqlOk() (*AnomalyLeql, bool)`

GetLeqlOk returns a tuple with the Leql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeql

`func (o *ScheduledQueryResponse) SetLeql(v AnomalyLeql)`

SetLeql sets Leql field to given value.

### HasLeql

`func (o *ScheduledQueryResponse) HasLeql() bool`

HasLeql returns a boolean if a field has been set.

### GetTimeValue

`func (o *ScheduledQueryResponse) GetTimeValue() int32`

GetTimeValue returns the TimeValue field if non-nil, zero value otherwise.

### GetTimeValueOk

`func (o *ScheduledQueryResponse) GetTimeValueOk() (*int32, bool)`

GetTimeValueOk returns a tuple with the TimeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeValue

`func (o *ScheduledQueryResponse) SetTimeValue(v int32)`

SetTimeValue sets TimeValue field to given value.

### HasTimeValue

`func (o *ScheduledQueryResponse) HasTimeValue() bool`

HasTimeValue returns a boolean if a field has been set.

### GetTimePeriod

`func (o *ScheduledQueryResponse) GetTimePeriod() string`

GetTimePeriod returns the TimePeriod field if non-nil, zero value otherwise.

### GetTimePeriodOk

`func (o *ScheduledQueryResponse) GetTimePeriodOk() (*string, bool)`

GetTimePeriodOk returns a tuple with the TimePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePeriod

`func (o *ScheduledQueryResponse) SetTimePeriod(v string)`

SetTimePeriod sets TimePeriod field to given value.

### HasTimePeriod

`func (o *ScheduledQueryResponse) HasTimePeriod() bool`

HasTimePeriod returns a boolean if a field has been set.

### GetTimeframeValue

`func (o *ScheduledQueryResponse) GetTimeframeValue() int32`

GetTimeframeValue returns the TimeframeValue field if non-nil, zero value otherwise.

### GetTimeframeValueOk

`func (o *ScheduledQueryResponse) GetTimeframeValueOk() (*int32, bool)`

GetTimeframeValueOk returns a tuple with the TimeframeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframeValue

`func (o *ScheduledQueryResponse) SetTimeframeValue(v int32)`

SetTimeframeValue sets TimeframeValue field to given value.

### HasTimeframeValue

`func (o *ScheduledQueryResponse) HasTimeframeValue() bool`

HasTimeframeValue returns a boolean if a field has been set.

### GetTimeframePeriod

`func (o *ScheduledQueryResponse) GetTimeframePeriod() string`

GetTimeframePeriod returns the TimeframePeriod field if non-nil, zero value otherwise.

### GetTimeframePeriodOk

`func (o *ScheduledQueryResponse) GetTimeframePeriodOk() (*string, bool)`

GetTimeframePeriodOk returns a tuple with the TimeframePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframePeriod

`func (o *ScheduledQueryResponse) SetTimeframePeriod(v string)`

SetTimeframePeriod sets TimeframePeriod field to given value.

### HasTimeframePeriod

`func (o *ScheduledQueryResponse) HasTimeframePeriod() bool`

HasTimeframePeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


