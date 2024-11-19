# EventQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | Pointer to [**[]LabelsInner**](LabelsInner.md) | the UUIDs of all labels of the log entry (and a link to **_/management/labels/{label_id}**). | [optional] 
**Timestamp** | **int32** | The UNIX timestamp in milliseconds, of when the log entry was received. | 
**SequenceNumber** | **int32** | The sequence number of the log entry (a unique identifier used to distinguish between log entries received in the same millisecond). | 
**LogId** | **string** | The key of the log that the entry belongs to. | 
**Message** | **string** | The log entry itself. | 
**Links** | Pointer to [**[]LinksInner**](LinksInner.md) |  | [optional] 
**SequenceNumberStr** | **string** | The sequence number as a string, rather than as an integer. | 
**KvpInfo** | Pointer to [**[]KvpInfoInner**](KvpInfoInner.md) | Information about the KVPs in the log entry that were parsed when they were received. | [optional] 

## Methods

### NewEventQuery

`func NewEventQuery(timestamp int32, sequenceNumber int32, logId string, message string, sequenceNumberStr string, ) *EventQuery`

NewEventQuery instantiates a new EventQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventQueryWithDefaults

`func NewEventQueryWithDefaults() *EventQuery`

NewEventQueryWithDefaults instantiates a new EventQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *EventQuery) GetLabels() []LabelsInner`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *EventQuery) GetLabelsOk() (*[]LabelsInner, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *EventQuery) SetLabels(v []LabelsInner)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *EventQuery) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetTimestamp

`func (o *EventQuery) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *EventQuery) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *EventQuery) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetSequenceNumber

`func (o *EventQuery) GetSequenceNumber() int32`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *EventQuery) GetSequenceNumberOk() (*int32, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *EventQuery) SetSequenceNumber(v int32)`

SetSequenceNumber sets SequenceNumber field to given value.


### GetLogId

`func (o *EventQuery) GetLogId() string`

GetLogId returns the LogId field if non-nil, zero value otherwise.

### GetLogIdOk

`func (o *EventQuery) GetLogIdOk() (*string, bool)`

GetLogIdOk returns a tuple with the LogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogId

`func (o *EventQuery) SetLogId(v string)`

SetLogId sets LogId field to given value.


### GetMessage

`func (o *EventQuery) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EventQuery) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EventQuery) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetLinks

`func (o *EventQuery) GetLinks() []LinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EventQuery) GetLinksOk() (*[]LinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EventQuery) SetLinks(v []LinksInner)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EventQuery) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetSequenceNumberStr

`func (o *EventQuery) GetSequenceNumberStr() string`

GetSequenceNumberStr returns the SequenceNumberStr field if non-nil, zero value otherwise.

### GetSequenceNumberStrOk

`func (o *EventQuery) GetSequenceNumberStrOk() (*string, bool)`

GetSequenceNumberStrOk returns a tuple with the SequenceNumberStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumberStr

`func (o *EventQuery) SetSequenceNumberStr(v string)`

SetSequenceNumberStr sets SequenceNumberStr field to given value.


### GetKvpInfo

`func (o *EventQuery) GetKvpInfo() []KvpInfoInner`

GetKvpInfo returns the KvpInfo field if non-nil, zero value otherwise.

### GetKvpInfoOk

`func (o *EventQuery) GetKvpInfoOk() (*[]KvpInfoInner, bool)`

GetKvpInfoOk returns a tuple with the KvpInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvpInfo

`func (o *EventQuery) SetKvpInfo(v []KvpInfoInner)`

SetKvpInfo sets KvpInfo field to given value.

### HasKvpInfo

`func (o *EventQuery) HasKvpInfo() bool`

HasKvpInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


