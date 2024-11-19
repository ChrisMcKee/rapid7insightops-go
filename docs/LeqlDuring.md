# LeqlDuring

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **int32** | The start of the time range for the query, as a UNIX timestamp in milliseconds. | [optional] 
**To** | Pointer to **int32** | The end of the time range for the query, as a UNIX timestamp in milliseconds. | [optional] 
**TimeRange** | Pointer to **string** | \\ Relative time range (instead of absolute &#x60;from&#x60; + &#x60;to&#x60; time range). Possible values: * &#x60;yesterday&#x60; * &#x60;today&#x60; * &#x60;last x timeunits&#x60; where x is the number of time unit back from the current server time. Supported time units (case insensitive):     - min(s) or minute(s)     - hr(s) or hour(s)     - day(s)     - week(s)     - month(s)     - year(s)  | [optional] 

## Methods

### NewLeqlDuring

`func NewLeqlDuring() *LeqlDuring`

NewLeqlDuring instantiates a new LeqlDuring object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeqlDuringWithDefaults

`func NewLeqlDuringWithDefaults() *LeqlDuring`

NewLeqlDuringWithDefaults instantiates a new LeqlDuring object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *LeqlDuring) GetFrom() int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *LeqlDuring) GetFromOk() (*int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *LeqlDuring) SetFrom(v int32)`

SetFrom sets From field to given value.

### HasFrom

`func (o *LeqlDuring) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *LeqlDuring) GetTo() int32`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *LeqlDuring) GetToOk() (*int32, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *LeqlDuring) SetTo(v int32)`

SetTo sets To field to given value.

### HasTo

`func (o *LeqlDuring) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetTimeRange

`func (o *LeqlDuring) GetTimeRange() string`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *LeqlDuring) GetTimeRangeOk() (*string, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *LeqlDuring) SetTimeRange(v string)`

SetTimeRange sets TimeRange field to given value.

### HasTimeRange

`func (o *LeqlDuring) HasTimeRange() bool`

HasTimeRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

