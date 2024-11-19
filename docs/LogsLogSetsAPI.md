# invoker\LogsLogSetsAPI

All URIs are relative to *https://eu.rest.logs.insight.rapid7.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteLog**](LogsLogSetsAPI.md#DeleteLog) | **Delete** /management/logs/{id} | Delete a Log
[**DeleteManagementLogsetById**](LogsLogSetsAPI.md#DeleteManagementLogsetById) | **Delete** /management/logsets/{logset_id} | Delete a Log Set
[**GetLog**](LogsLogSetsAPI.md#GetLog) | **Get** /management/logs/{id} | Retrieve a Log
[**GetLogs**](LogsLogSetsAPI.md#GetLogs) | **Get** /management/logs | Retrieve All Logs
[**GetManagementLogsetById**](LogsLogSetsAPI.md#GetManagementLogsetById) | **Get** /management/logsets/{logset_id} | Retrieve a Log Set
[**GetManagementLogsets**](LogsLogSetsAPI.md#GetManagementLogsets) | **Get** /management/logsets | List All Log Sets
[**PostLog**](LogsLogSetsAPI.md#PostLog) | **Post** /management/logs | Create a Log
[**PostManagementLogsets**](LogsLogSetsAPI.md#PostManagementLogsets) | **Post** /management/logsets | Create a Log Set
[**PutLog**](LogsLogSetsAPI.md#PutLog) | **Put** /management/logs/{id} | Replace a Log
[**PutManagementLogsetById**](LogsLogSetsAPI.md#PutManagementLogsetById) | **Put** /management/logsets/{logset_id} | Replace a Log Set



## DeleteLog

> DeleteLog(ctx, id).Execute()

Delete a Log



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
	id := "a53578a2-e667-423b-9694-3e989c388186" // string | \\ The id of the log to be queried. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogsLogSetsAPI.DeleteLog(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsLogSetsAPI.DeleteLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | \\ The id of the log to be queried.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## DeleteManagementLogsetById

> DeleteManagementLogsetById(ctx, logsetId).Execute()

Delete a Log Set



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
	logsetId := "4a8869fa-0843-4edd-85f5-bf7d6b80d6c5" // string | \\ The id of the log set to be queried. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogsLogSetsAPI.DeleteManagementLogsetById(context.Background(), logsetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsLogSetsAPI.DeleteManagementLogsetById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logsetId** | **string** | \\ The id of the log set to be queried.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteManagementLogsetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetLog

> LogResponse GetLog(ctx, id).Execute()

Retrieve a Log



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
	id := "a53578a2-e667-423b-9694-3e989c388186" // string | \\ The id of the log to be queried. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogsLogSetsAPI.GetLog(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsLogSetsAPI.GetLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLog`: LogResponse
	fmt.Fprintf(os.Stdout, "Response from `LogsLogSetsAPI.GetLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | \\ The id of the log to be queried.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LogResponse**](LogResponse.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogs

> LogsResponse GetLogs(ctx).Execute()

Retrieve All Logs



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
	resp, r, err := apiClient.LogsLogSetsAPI.GetLogs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsLogSetsAPI.GetLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogs`: LogsResponse
	fmt.Fprintf(os.Stdout, "Response from `LogsLogSetsAPI.GetLogs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogsRequest struct via the builder pattern


### Return type

[**LogsResponse**](LogsResponse.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagementLogsetById

> LogsetResponse GetManagementLogsetById(ctx, logsetId).Execute()

Retrieve a Log Set



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
	logsetId := "4a8869fa-0843-4edd-85f5-bf7d6b80d6c5" // string | \\ The id of the log set to be queried. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogsLogSetsAPI.GetManagementLogsetById(context.Background(), logsetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsLogSetsAPI.GetManagementLogsetById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManagementLogsetById`: LogsetResponse
	fmt.Fprintf(os.Stdout, "Response from `LogsLogSetsAPI.GetManagementLogsetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logsetId** | **string** | \\ The id of the log set to be queried.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetManagementLogsetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LogsetResponse**](LogsetResponse.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagementLogsets

> LogsetResponse GetManagementLogsets(ctx).Execute()

List All Log Sets



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
	resp, r, err := apiClient.LogsLogSetsAPI.GetManagementLogsets(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsLogSetsAPI.GetManagementLogsets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManagementLogsets`: LogsetResponse
	fmt.Fprintf(os.Stdout, "Response from `LogsLogSetsAPI.GetManagementLogsets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetManagementLogsetsRequest struct via the builder pattern


### Return type

[**LogsetResponse**](LogsetResponse.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLog

> LogResponse PostLog(ctx).CreateOrPutLog(createOrPutLog).Execute()

Create a Log



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
	createOrPutLog := *openapiclient.NewCreateOrPutLog(*openapiclient.NewCreateOrPutLogLog(interface{}(123))) // CreateOrPutLog | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogsLogSetsAPI.PostLog(context.Background()).CreateOrPutLog(createOrPutLog).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsLogSetsAPI.PostLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLog`: LogResponse
	fmt.Fprintf(os.Stdout, "Response from `LogsLogSetsAPI.PostLog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createOrPutLog** | [**CreateOrPutLog**](CreateOrPutLog.md) |  | 

### Return type

[**LogResponse**](LogResponse.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostManagementLogsets

> LogsetResponse PostManagementLogsets(ctx).CreateLogset(createLogset).Execute()

Create a Log Set



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
	createLogset := *openapiclient.NewCreateLogset() // CreateLogset | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogsLogSetsAPI.PostManagementLogsets(context.Background()).CreateLogset(createLogset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsLogSetsAPI.PostManagementLogsets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostManagementLogsets`: LogsetResponse
	fmt.Fprintf(os.Stdout, "Response from `LogsLogSetsAPI.PostManagementLogsets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostManagementLogsetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createLogset** | [**CreateLogset**](CreateLogset.md) |  | 

### Return type

[**LogsetResponse**](LogsetResponse.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutLog

> LogResponse PutLog(ctx, id).CreateOrPutLog(createOrPutLog).Execute()

Replace a Log



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
	id := "a53578a2-e667-423b-9694-3e989c388186" // string | \\ The id of the log to be queried. 
	createOrPutLog := *openapiclient.NewCreateOrPutLog(*openapiclient.NewCreateOrPutLogLog(interface{}(123))) // CreateOrPutLog | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogsLogSetsAPI.PutLog(context.Background(), id).CreateOrPutLog(createOrPutLog).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsLogSetsAPI.PutLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutLog`: LogResponse
	fmt.Fprintf(os.Stdout, "Response from `LogsLogSetsAPI.PutLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | \\ The id of the log to be queried.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrPutLog** | [**CreateOrPutLog**](CreateOrPutLog.md) |  | 

### Return type

[**LogResponse**](LogResponse.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutManagementLogsetById

> LogsetResponse PutManagementLogsetById(ctx, logsetId).LogsetResponse(logsetResponse).Execute()

Replace a Log Set



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
	logsetId := "4a8869fa-0843-4edd-85f5-bf7d6b80d6c5" // string | \\ The id of the log set to be queried. 
	logsetResponse := *openapiclient.NewLogsetResponse() // LogsetResponse | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogsLogSetsAPI.PutManagementLogsetById(context.Background(), logsetId).LogsetResponse(logsetResponse).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsLogSetsAPI.PutManagementLogsetById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutManagementLogsetById`: LogsetResponse
	fmt.Fprintf(os.Stdout, "Response from `LogsLogSetsAPI.PutManagementLogsetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logsetId** | **string** | \\ The id of the log set to be queried.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutManagementLogsetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **logsetResponse** | [**LogsetResponse**](LogsetResponse.md) |  | 

### Return type

[**LogsetResponse**](LogsetResponse.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

