# EventLiveTail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | Pointer to [**[]LabelsInner**](LabelsInner.md) | the UUIDs of all labels of the log entry (and a link to **_/management/labels/{label_id}**). | [optional] 
**Timestamp** | **int32** | The UNIX timestamp in milliseconds, of when the log entry was received. | 
**SequenceNumber** | **int32** | The sequence number of the log entry (a unique identifier used to distinguish between log entries received in the same millisecond). | 
**LogId** | **string** | The key of the log that the entry belongs to. | 
**Message** | **string** | The log entry itself. | 
**SequenceNumberStr** | **string** | The sequence number as a string, rather than as an integer. | 

## Methods

### NewEventLiveTail

`func NewEventLiveTail(timestamp int32, sequenceNumber int32, logId string, message string, sequenceNumberStr string, ) *EventLiveTail`

NewEventLiveTail instantiates a new EventLiveTail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventLiveTailWithDefaults

`func NewEventLiveTailWithDefaults() *EventLiveTail`

NewEventLiveTailWithDefaults instantiates a new EventLiveTail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *EventLiveTail) GetLabels() []LabelsInner`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *EventLiveTail) GetLabelsOk() (*[]LabelsInner, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *EventLiveTail) SetLabels(v []LabelsInner)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *EventLiveTail) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetTimestamp

`func (o *EventLiveTail) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *EventLiveTail) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *EventLiveTail) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetSequenceNumber

`func (o *EventLiveTail) GetSequenceNumber() int32`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *EventLiveTail) GetSequenceNumberOk() (*int32, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *EventLiveTail) SetSequenceNumber(v int32)`

SetSequenceNumber sets SequenceNumber field to given value.


### GetLogId

`func (o *EventLiveTail) GetLogId() string`

GetLogId returns the LogId field if non-nil, zero value otherwise.

### GetLogIdOk

`func (o *EventLiveTail) GetLogIdOk() (*string, bool)`

GetLogIdOk returns a tuple with the LogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogId

`func (o *EventLiveTail) SetLogId(v string)`

SetLogId sets LogId field to given value.


### GetMessage

`func (o *EventLiveTail) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EventLiveTail) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EventLiveTail) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetSequenceNumberStr

`func (o *EventLiveTail) GetSequenceNumberStr() string`

GetSequenceNumberStr returns the SequenceNumberStr field if non-nil, zero value otherwise.

### GetSequenceNumberStrOk

`func (o *EventLiveTail) GetSequenceNumberStrOk() (*string, bool)`

GetSequenceNumberStrOk returns a tuple with the SequenceNumberStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumberStr

`func (o *EventLiveTail) SetSequenceNumberStr(v string)`

SetSequenceNumberStr sets SequenceNumberStr field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


