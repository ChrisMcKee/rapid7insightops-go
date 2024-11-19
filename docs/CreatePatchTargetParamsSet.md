# CreatePatchTargetParamsSet

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

### NewCreatePatchTargetParamsSet

`func NewCreatePatchTargetParamsSet(direct string, serviceKey string, url string, destinationApp string, destinationQueue string, workflowId string, ) *CreatePatchTargetParamsSet`

NewCreatePatchTargetParamsSet instantiates a new CreatePatchTargetParamsSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePatchTargetParamsSetWithDefaults

`func NewCreatePatchTargetParamsSetWithDefaults() *CreatePatchTargetParamsSet`

NewCreatePatchTargetParamsSetWithDefaults instantiates a new CreatePatchTargetParamsSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirect

`func (o *CreatePatchTargetParamsSet) GetDirect() string`

GetDirect returns the Direct field if non-nil, zero value otherwise.

### GetDirectOk

`func (o *CreatePatchTargetParamsSet) GetDirectOk() (*string, bool)`

GetDirectOk returns a tuple with the Direct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirect

`func (o *CreatePatchTargetParamsSet) SetDirect(v string)`

SetDirect sets Direct field to given value.


### GetServiceKey

`func (o *CreatePatchTargetParamsSet) GetServiceKey() string`

GetServiceKey returns the ServiceKey field if non-nil, zero value otherwise.

### GetServiceKeyOk

`func (o *CreatePatchTargetParamsSet) GetServiceKeyOk() (*string, bool)`

GetServiceKeyOk returns a tuple with the ServiceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKey

`func (o *CreatePatchTargetParamsSet) SetServiceKey(v string)`

SetServiceKey sets ServiceKey field to given value.


### GetDescription

`func (o *CreatePatchTargetParamsSet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreatePatchTargetParamsSet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreatePatchTargetParamsSet) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreatePatchTargetParamsSet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUrl

`func (o *CreatePatchTargetParamsSet) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreatePatchTargetParamsSet) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreatePatchTargetParamsSet) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDestinationApp

`func (o *CreatePatchTargetParamsSet) GetDestinationApp() string`

GetDestinationApp returns the DestinationApp field if non-nil, zero value otherwise.

### GetDestinationAppOk

`func (o *CreatePatchTargetParamsSet) GetDestinationAppOk() (*string, bool)`

GetDestinationAppOk returns a tuple with the DestinationApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationApp

`func (o *CreatePatchTargetParamsSet) SetDestinationApp(v string)`

SetDestinationApp sets DestinationApp field to given value.


### GetDestinationQueue

`func (o *CreatePatchTargetParamsSet) GetDestinationQueue() string`

GetDestinationQueue returns the DestinationQueue field if non-nil, zero value otherwise.

### GetDestinationQueueOk

`func (o *CreatePatchTargetParamsSet) GetDestinationQueueOk() (*string, bool)`

GetDestinationQueueOk returns a tuple with the DestinationQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationQueue

`func (o *CreatePatchTargetParamsSet) SetDestinationQueue(v string)`

SetDestinationQueue sets DestinationQueue field to given value.


### GetWorkflowId

`func (o *CreatePatchTargetParamsSet) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *CreatePatchTargetParamsSet) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *CreatePatchTargetParamsSet) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


