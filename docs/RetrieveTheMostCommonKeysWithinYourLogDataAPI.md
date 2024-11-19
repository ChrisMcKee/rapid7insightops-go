# invoker\RetrieveTheMostCommonKeysWithinYourLogDataAPI

All URIs are relative to *https://eu.rest.logs.insight.rapid7.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLogTopkeys**](RetrieveTheMostCommonKeysWithinYourLogDataAPI.md#GetLogTopkeys) | **Get** /management/logs/{id}/topkeys | Retrieve Most Common Keys For A Log



## GetLogTopkeys

> LogTopkeysResponse GetLogTopkeys(ctx, id).Execute()

Retrieve Most Common Keys For A Log



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
	resp, r, err := apiClient.RetrieveTheMostCommonKeysWithinYourLogDataAPI.GetLogTopkeys(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RetrieveTheMostCommonKeysWithinYourLogDataAPI.GetLogTopkeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogTopkeys`: LogTopkeysResponse
	fmt.Fprintf(os.Stdout, "Response from `RetrieveTheMostCommonKeysWithinYourLogDataAPI.GetLogTopkeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | \\ The id of the log to be queried.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogTopkeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LogTopkeysResponse**](LogTopkeysResponse.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

