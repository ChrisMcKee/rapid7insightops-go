# CreateScheduledQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the scheduled query. | 
**Leql** | [**AnomalyLeql**](AnomalyLeql.md) |  | 
**TimeValue** | **int32** | Defines the time range for the two queries, along with the &#x60;time_period&#x60; parameter. Two queries will run: the time range of the first, more recent query, will end at the current system time; the time range of the older query will end at the current system time minus the time specified by the &#x60;timeframe_value&#x60; and &#x60;timeframe_period&#x60; parameters.  | 
**TimePeriod** | **string** | Defines the time range for the two queries, along with the &#x60;time_value&#x60; parameter. | 
**TimeframeValue** | **int32** | Defines the \&quot;offset\&quot; for the older of the two queries along with the &#x60;timeframe_period&#x60; parameter. For example, if the &#x60;timeframe_value&#x60; is set to 1 and &#x60;timeframe_period&#x60; is set to &#x60;Week&#x60;, then the time range for the older of the two queries will end at the current system time minus 1 week. | 
**TimeframePeriod** | **string** | Defines the \&quot;offset\&quot; for the older of the two queries, along with the &#x60;timeframe_value&#x60; parameter. | 

## Methods

### NewCreateScheduledQuery

`func NewCreateScheduledQuery(name string, leql AnomalyLeql, timeValue int32, timePeriod string, timeframeValue int32, timeframePeriod string, ) *CreateScheduledQuery`

NewCreateScheduledQuery instantiates a new CreateScheduledQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateScheduledQueryWithDefaults

`func NewCreateScheduledQueryWithDefaults() *CreateScheduledQuery`

NewCreateScheduledQueryWithDefaults instantiates a new CreateScheduledQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateScheduledQuery) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateScheduledQuery) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateScheduledQuery) SetName(v string)`

SetName sets Name field to given value.


### GetLeql

`func (o *CreateScheduledQuery) GetLeql() AnomalyLeql`

GetLeql returns the Leql field if non-nil, zero value otherwise.

### GetLeqlOk

`func (o *CreateScheduledQuery) GetLeqlOk() (*AnomalyLeql, bool)`

GetLeqlOk returns a tuple with the Leql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeql

`func (o *CreateScheduledQuery) SetLeql(v AnomalyLeql)`

SetLeql sets Leql field to given value.


### GetTimeValue

`func (o *CreateScheduledQuery) GetTimeValue() int32`

GetTimeValue returns the TimeValue field if non-nil, zero value otherwise.

### GetTimeValueOk

`func (o *CreateScheduledQuery) GetTimeValueOk() (*int32, bool)`

GetTimeValueOk returns a tuple with the TimeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeValue

`func (o *CreateScheduledQuery) SetTimeValue(v int32)`

SetTimeValue sets TimeValue field to given value.


### GetTimePeriod

`func (o *CreateScheduledQuery) GetTimePeriod() string`

GetTimePeriod returns the TimePeriod field if non-nil, zero value otherwise.

### GetTimePeriodOk

`func (o *CreateScheduledQuery) GetTimePeriodOk() (*string, bool)`

GetTimePeriodOk returns a tuple with the TimePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePeriod

`func (o *CreateScheduledQuery) SetTimePeriod(v string)`

SetTimePeriod sets TimePeriod field to given value.


### GetTimeframeValue

`func (o *CreateScheduledQuery) GetTimeframeValue() int32`

GetTimeframeValue returns the TimeframeValue field if non-nil, zero value otherwise.

### GetTimeframeValueOk

`func (o *CreateScheduledQuery) GetTimeframeValueOk() (*int32, bool)`

GetTimeframeValueOk returns a tuple with the TimeframeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframeValue

`func (o *CreateScheduledQuery) SetTimeframeValue(v int32)`

SetTimeframeValue sets TimeframeValue field to given value.


### GetTimeframePeriod

`func (o *CreateScheduledQuery) GetTimeframePeriod() string`

GetTimeframePeriod returns the TimeframePeriod field if non-nil, zero value otherwise.

### GetTimeframePeriodOk

`func (o *CreateScheduledQuery) GetTimeframePeriodOk() (*string, bool)`

GetTimeframePeriodOk returns a tuple with the TimeframePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframePeriod

`func (o *CreateScheduledQuery) SetTimeframePeriod(v string)`

SetTimeframePeriod sets TimeframePeriod field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


