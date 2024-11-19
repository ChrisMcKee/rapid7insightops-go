# TargetResponseParamsSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direct** | **string** | The list of emails to send the notification to. The emails are separated by commas. | 
**ServiceKey** | **string** | The service key of the PagerDuty. | 
**Description** | Pointer to **string** | The description of the PagerDuty target. | [optional] 
**Url** | **string** | The URL to the hook to send the notification to. | 
**DestinationApp** | **string** | The product code for any license with LOG_SEARCH e.g. OPS, IDR. | 
**DestinationQueue** | **string** | The name for the destination queue e.g. IDR_investigation_queue. This is not the URL. | 
**WorkflowId** | **string** | The ID of the basic detection rule trigger type of workflow in InsightConnect. | 

## Methods

### NewTargetResponseParamsSet

`func NewTargetResponseParamsSet(direct string, serviceKey string, url string, destinationApp string, destinationQueue string, workflowId string, ) *TargetResponseParamsSet`

NewTargetResponseParamsSet instantiates a new TargetResponseParamsSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetResponseParamsSetWithDefaults

`func NewTargetResponseParamsSetWithDefaults() *TargetResponseParamsSet`

NewTargetResponseParamsSetWithDefaults instantiates a new TargetResponseParamsSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirect

`func (o *TargetResponseParamsSet) GetDirect() string`

GetDirect returns the Direct field if non-nil, zero value otherwise.

### GetDirectOk

`func (o *TargetResponseParamsSet) GetDirectOk() (*string, bool)`

GetDirectOk returns a tuple with the Direct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirect

`func (o *TargetResponseParamsSet) SetDirect(v string)`

SetDirect sets Direct field to given value.


### GetServiceKey

`func (o *TargetResponseParamsSet) GetServiceKey() string`

GetServiceKey returns the ServiceKey field if non-nil, zero value otherwise.

### GetServiceKeyOk

`func (o *TargetResponseParamsSet) GetServiceKeyOk() (*string, bool)`

GetServiceKeyOk returns a tuple with the ServiceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKey

`func (o *TargetResponseParamsSet) SetServiceKey(v string)`

SetServiceKey sets ServiceKey field to given value.


### GetDescription

`func (o *TargetResponseParamsSet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TargetResponseParamsSet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TargetResponseParamsSet) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TargetResponseParamsSet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUrl

`func (o *TargetResponseParamsSet) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TargetResponseParamsSet) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TargetResponseParamsSet) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDestinationApp

`func (o *TargetResponseParamsSet) GetDestinationApp() string`

GetDestinationApp returns the DestinationApp field if non-nil, zero value otherwise.

### GetDestinationAppOk

`func (o *TargetResponseParamsSet) GetDestinationAppOk() (*string, bool)`

GetDestinationAppOk returns a tuple with the DestinationApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationApp

`func (o *TargetResponseParamsSet) SetDestinationApp(v string)`

SetDestinationApp sets DestinationApp field to given value.


### GetDestinationQueue

`func (o *TargetResponseParamsSet) GetDestinationQueue() string`

GetDestinationQueue returns the DestinationQueue field if non-nil, zero value otherwise.

### GetDestinationQueueOk

`func (o *TargetResponseParamsSet) GetDestinationQueueOk() (*string, bool)`

GetDestinationQueueOk returns a tuple with the DestinationQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationQueue

`func (o *TargetResponseParamsSet) SetDestinationQueue(v string)`

SetDestinationQueue sets DestinationQueue field to given value.


### GetWorkflowId

`func (o *TargetResponseParamsSet) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *TargetResponseParamsSet) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *TargetResponseParamsSet) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


