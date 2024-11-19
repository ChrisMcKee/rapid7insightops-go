# invoker\LEQLVariablesAPI

All URIs are relative to *https://eu.rest.logs.insight.rapid7.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVariable**](LEQLVariablesAPI.md#CreateVariable) | **Post** /query/variables | Create a LEQL variable
[**DeleteVariable**](LEQLVariablesAPI.md#DeleteVariable) | **Delete** /query/variables/{id} | Delete a LEQL variable
[**ListVariables**](LEQLVariablesAPI.md#ListVariables) | **Get** /query/variables | List all LEQL variables
[**RetrieveVariable**](LEQLVariablesAPI.md#RetrieveVariable) | **Get** /query/variables/{id} | Retrieve a LEQL variable
[**UpdateVariable**](LEQLVariablesAPI.md#UpdateVariable) | **Put** /query/variables/{id} | Update a LEQL variable



## CreateVariable

> CreateVariable201Response CreateVariable(ctx).CreateVariableRequest(createVariableRequest).Execute()

Create a LEQL variable



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
	createVariableRequest := *openapiclient.NewCreateVariableRequest() // CreateVariableRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LEQLVariablesAPI.CreateVariable(context.Background()).CreateVariableRequest(createVariableRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LEQLVariablesAPI.CreateVariable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVariable`: CreateVariable201Response
	fmt.Fprintf(os.Stdout, "Response from `LEQLVariablesAPI.CreateVariable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createVariableRequest** | [**CreateVariableRequest**](CreateVariableRequest.md) |  | 

### Return type

[**CreateVariable201Response**](CreateVariable201Response.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVariable

> DeleteVariable(ctx).Execute()

Delete a LEQL variable



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
	r, err := apiClient.LEQLVariablesAPI.DeleteVariable(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LEQLVariablesAPI.DeleteVariable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVariableRequest struct via the builder pattern


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


## ListVariables

> ListVariables200Response ListVariables(ctx).Execute()

List all LEQL variables



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
	resp, r, err := apiClient.LEQLVariablesAPI.ListVariables(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LEQLVariablesAPI.ListVariables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVariables`: ListVariables200Response
	fmt.Fprintf(os.Stdout, "Response from `LEQLVariablesAPI.ListVariables`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVariablesRequest struct via the builder pattern


### Return type

[**ListVariables200Response**](ListVariables200Response.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveVariable

> CreateVariable201Response RetrieveVariable(ctx).Execute()

Retrieve a LEQL variable



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
	resp, r, err := apiClient.LEQLVariablesAPI.RetrieveVariable(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LEQLVariablesAPI.RetrieveVariable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveVariable`: CreateVariable201Response
	fmt.Fprintf(os.Stdout, "Response from `LEQLVariablesAPI.RetrieveVariable`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveVariableRequest struct via the builder pattern


### Return type

[**CreateVariable201Response**](CreateVariable201Response.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVariable

> CreateVariable201Response UpdateVariable(ctx).CreateVariableRequest(createVariableRequest).Execute()

Update a LEQL variable



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
	createVariableRequest := *openapiclient.NewCreateVariableRequest() // CreateVariableRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LEQLVariablesAPI.UpdateVariable(context.Background()).CreateVariableRequest(createVariableRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LEQLVariablesAPI.UpdateVariable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateVariable`: CreateVariable201Response
	fmt.Fprintf(os.Stdout, "Response from `LEQLVariablesAPI.UpdateVariable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createVariableRequest** | [**CreateVariableRequest**](CreateVariableRequest.md) |  | 

### Return type

[**CreateVariable201Response**](CreateVariable201Response.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

