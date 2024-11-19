# invoker\BackupYourLogDataToS3API

All URIs are relative to *https://eu.rest.logs.insight.rapid7.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteArchivingS3Setup**](BackupYourLogDataToS3API.md#DeleteArchivingS3Setup) | **Delete** /management/archiving/s3setup | Disable Daily Archiving
[**GetArchivingS3Setup**](BackupYourLogDataToS3API.md#GetArchivingS3Setup) | **Get** /management/archiving/s3setup | View Daily Archiving Settings
[**PatchArchivingS3Setup**](BackupYourLogDataToS3API.md#PatchArchivingS3Setup) | **Patch** /management/archiving/s3setup | Modify Daily Archiving Settings
[**PostArchivingS3Setup**](BackupYourLogDataToS3API.md#PostArchivingS3Setup) | **Post** /management/archiving/s3setup | Enable Daily Archiving
[**PutArchivingS3Setup**](BackupYourLogDataToS3API.md#PutArchivingS3Setup) | **Put** /management/archiving/s3setup | Replace Daily Archiving Settings



## DeleteArchivingS3Setup

> DeleteArchivingS3Setup(ctx).Execute()

Disable Daily Archiving



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
	r, err := apiClient.BackupYourLogDataToS3API.DeleteArchivingS3Setup(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupYourLogDataToS3API.DeleteArchivingS3Setup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteArchivingS3SetupRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArchivingS3Setup

> S3ArchivingSetup GetArchivingS3Setup(ctx).Execute()

View Daily Archiving Settings



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
	resp, r, err := apiClient.BackupYourLogDataToS3API.GetArchivingS3Setup(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupYourLogDataToS3API.GetArchivingS3Setup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetArchivingS3Setup`: S3ArchivingSetup
	fmt.Fprintf(os.Stdout, "Response from `BackupYourLogDataToS3API.GetArchivingS3Setup`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetArchivingS3SetupRequest struct via the builder pattern


### Return type

[**S3ArchivingSetup**](S3ArchivingSetup.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchArchivingS3Setup

> S3ArchivingSetup PatchArchivingS3Setup(ctx).PatchS3ArchivingSetup(patchS3ArchivingSetup).Execute()

Modify Daily Archiving Settings



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
	patchS3ArchivingSetup := *openapiclient.NewPatchS3ArchivingSetup() // PatchS3ArchivingSetup | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupYourLogDataToS3API.PatchArchivingS3Setup(context.Background()).PatchS3ArchivingSetup(patchS3ArchivingSetup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupYourLogDataToS3API.PatchArchivingS3Setup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchArchivingS3Setup`: S3ArchivingSetup
	fmt.Fprintf(os.Stdout, "Response from `BackupYourLogDataToS3API.PatchArchivingS3Setup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchArchivingS3SetupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchS3ArchivingSetup** | [**PatchS3ArchivingSetup**](PatchS3ArchivingSetup.md) |  | 

### Return type

[**S3ArchivingSetup**](S3ArchivingSetup.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostArchivingS3Setup

> S3ArchivingSetup PostArchivingS3Setup(ctx).CreateS3ArchivingSetup(createS3ArchivingSetup).Execute()

Enable Daily Archiving



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
	createS3ArchivingSetup := *openapiclient.NewCreateS3ArchivingSetup() // CreateS3ArchivingSetup | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupYourLogDataToS3API.PostArchivingS3Setup(context.Background()).CreateS3ArchivingSetup(createS3ArchivingSetup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupYourLogDataToS3API.PostArchivingS3Setup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostArchivingS3Setup`: S3ArchivingSetup
	fmt.Fprintf(os.Stdout, "Response from `BackupYourLogDataToS3API.PostArchivingS3Setup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostArchivingS3SetupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createS3ArchivingSetup** | [**CreateS3ArchivingSetup**](CreateS3ArchivingSetup.md) |  | 

### Return type

[**S3ArchivingSetup**](S3ArchivingSetup.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutArchivingS3Setup

> S3ArchivingSetup PutArchivingS3Setup(ctx).CreateS3ArchivingSetup(createS3ArchivingSetup).Execute()

Replace Daily Archiving Settings



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
	createS3ArchivingSetup := *openapiclient.NewCreateS3ArchivingSetup() // CreateS3ArchivingSetup | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupYourLogDataToS3API.PutArchivingS3Setup(context.Background()).CreateS3ArchivingSetup(createS3ArchivingSetup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupYourLogDataToS3API.PutArchivingS3Setup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutArchivingS3Setup`: S3ArchivingSetup
	fmt.Fprintf(os.Stdout, "Response from `BackupYourLogDataToS3API.PutArchivingS3Setup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutArchivingS3SetupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createS3ArchivingSetup** | [**CreateS3ArchivingSetup**](CreateS3ArchivingSetup.md) |  | 

### Return type

[**S3ArchivingSetup**](S3ArchivingSetup.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

