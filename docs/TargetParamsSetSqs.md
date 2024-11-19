# TargetParamsSetSqs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationApp** | **string** | The product code for any license with LOG_SEARCH e.g. OPS, IDR. | 
**DestinationQueue** | **string** | The name for the destination queue e.g. IDR_investigation_queue. This is not the URL. | 

## Methods

### NewTargetParamsSetSqs

`func NewTargetParamsSetSqs(destinationApp string, destinationQueue string, ) *TargetParamsSetSqs`

NewTargetParamsSetSqs instantiates a new TargetParamsSetSqs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetParamsSetSqsWithDefaults

`func NewTargetParamsSetSqsWithDefaults() *TargetParamsSetSqs`

NewTargetParamsSetSqsWithDefaults instantiates a new TargetParamsSetSqs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationApp

`func (o *TargetParamsSetSqs) GetDestinationApp() string`

GetDestinationApp returns the DestinationApp field if non-nil, zero value otherwise.

### GetDestinationAppOk

`func (o *TargetParamsSetSqs) GetDestinationAppOk() (*string, bool)`

GetDestinationAppOk returns a tuple with the DestinationApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationApp

`func (o *TargetParamsSetSqs) SetDestinationApp(v string)`

SetDestinationApp sets DestinationApp field to given value.


### GetDestinationQueue

`func (o *TargetParamsSetSqs) GetDestinationQueue() string`

GetDestinationQueue returns the DestinationQueue field if non-nil, zero value otherwise.

### GetDestinationQueueOk

`func (o *TargetParamsSetSqs) GetDestinationQueueOk() (*string, bool)`

GetDestinationQueueOk returns a tuple with the DestinationQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationQueue

`func (o *TargetParamsSetSqs) SetDestinationQueue(v string)`

SetDestinationQueue sets DestinationQueue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


