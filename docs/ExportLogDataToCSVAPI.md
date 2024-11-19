# invoker\ExportLogDataToCSVAPI

All URIs are relative to *https://eu.rest.logs.insight.rapid7.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteExportJob**](ExportLogDataToCSVAPI.md#DeleteExportJob) | **Delete** /exports/{id} | Delete An Export Job By Id
[**GetExportJob**](ExportLogDataToCSVAPI.md#GetExportJob) | **Get** /exports/{id} | Retrieve An Export Job By Id
[**GetExportJobs**](ExportLogDataToCSVAPI.md#GetExportJobs) | **Get** /exports | Retrieve All Export Jobs



## DeleteExportJob

> DeleteExportJob(ctx).Id(id).Execute()

Delete An Export Job By Id



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/chrismckee/rapid7insightops-go"
)

func main() {
	id := "00000000-0000-0000-0000-000000000000" // string | \\ The UUID of the export job. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExportLogDataToCSVAPI.DeleteExportJob(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportLogDataToCSVAPI.DeleteExportJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteExportJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | \\ The UUID of the export job.  | 

### Return type

 (empty response body)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExportJob

> ExportJobResponse GetExportJob(ctx).Id(id).Execute()

Retrieve An Export Job By Id



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/chrismckee/rapid7insightops-go"
)

func main() {
	id := "00000000-0000-0000-0000-000000000000" // string | \\ The UUID of the export job. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExportLogDataToCSVAPI.GetExportJob(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportLogDataToCSVAPI.GetExportJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExportJob`: ExportJobResponse
	fmt.Fprintf(os.Stdout, "Response from `ExportLogDataToCSVAPI.GetExportJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExportJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | \\ The UUID of the export job.  | 

### Return type

[**ExportJobResponse**](ExportJobResponse.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExportJobs

> ExportJobsResponse GetExportJobs(ctx).Execute()

Retrieve All Export Jobs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/chrismckee/rapid7insightops-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExportLogDataToCSVAPI.GetExportJobs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportLogDataToCSVAPI.GetExportJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExportJobs`: ExportJobsResponse
	fmt.Fprintf(os.Stdout, "Response from `ExportLogDataToCSVAPI.GetExportJobs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetExportJobsRequest struct via the builder pattern


### Return type

[**ExportJobsResponse**](ExportJobsResponse.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

