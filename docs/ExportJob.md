# ExportJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The UUID of the job. | 
**Created** | **int32** | The Unix timestamp of when the job was created in milliseconds. | 
**Format** | **string** | The export format. This can only be &#39;csv&#39;. | 
**Percentage** | **int32** | The percentage progress of the export job. | 
**Status** | **string** | The status of the export job.   * &#x60;Started&#x60;: The job has been accepted and registered.   * &#x60;Retrieving&#x60;: The query is being executed and the events are being retrieved and streamed to disk.   * &#x60;Converting&#x60;: The events are being run through the converter pipeline.   * &#x60;Uploading&#x60;: The results are being uploaded to the S3 bucket.   * &#x60;Completed&#x60;: The job has been completed and the direct download link is available.   * &#x60;Failed&#x60;: The export job has failed. Report this to support for analysis and resolution.  | 
**Expires** | **int32** | The Unix timestamp of when the job result will be deleted. | 
**Links** | [**[]ExportJobLinksInner**](ExportJobLinksInner.md) |  | 

## Methods

### NewExportJob

`func NewExportJob(id string, created int32, format string, percentage int32, status string, expires int32, links []ExportJobLinksInner, ) *ExportJob`

NewExportJob instantiates a new ExportJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportJobWithDefaults

`func NewExportJobWithDefaults() *ExportJob`

NewExportJobWithDefaults instantiates a new ExportJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExportJob) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExportJob) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExportJob) SetId(v string)`

SetId sets Id field to given value.


### GetCreated

`func (o *ExportJob) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ExportJob) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ExportJob) SetCreated(v int32)`

SetCreated sets Created field to given value.


### GetFormat

`func (o *ExportJob) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ExportJob) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ExportJob) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetPercentage

`func (o *ExportJob) GetPercentage() int32`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *ExportJob) GetPercentageOk() (*int32, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *ExportJob) SetPercentage(v int32)`

SetPercentage sets Percentage field to given value.


### GetStatus

`func (o *ExportJob) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExportJob) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExportJob) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetExpires

`func (o *ExportJob) GetExpires() int32`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *ExportJob) GetExpiresOk() (*int32, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *ExportJob) SetExpires(v int32)`

SetExpires sets Expires field to given value.


### GetLinks

`func (o *ExportJob) GetLinks() []ExportJobLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExportJob) GetLinksOk() (*[]ExportJobLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExportJob) SetLinks(v []ExportJobLinksInner)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


