# invoker\ExploreTheSizeOfYourLogDataAPI

All URIs are relative to *https://eu.rest.logs.insight.rapid7.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationsLogUsage**](ExploreTheSizeOfYourLogDataAPI.md#GetOrganizationsLogUsage) | **Get** /usage/organizations/logs/{log_key} | Data Size For a Specific Log
[**GetOrganizationsLogsUsage**](ExploreTheSizeOfYourLogDataAPI.md#GetOrganizationsLogsUsage) | **Get** /usage/organizations/logs | Data Size Broken Down By Log
[**GetOrganizationsUsage**](ExploreTheSizeOfYourLogDataAPI.md#GetOrganizationsUsage) | **Get** /usage/organizations | Total Data Size Across All Logs



## GetOrganizationsLogUsage

> LogUsageResponse GetOrganizationsLogUsage(ctx, logKey).From(from).To(to).Execute()

Data Size For a Specific Log



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
	logKey := "565c1b7b-c08b-4c87-a42a-ab08bad56071" // string | \\ The key of the log to be queried. 
	from := "2023-09-14" // string | \\ The start of the time range as a string in the format \"YYYY-MM-DD\". 
	to := "2023-09-20" // string | \\ The end of the time range as a string in the format \"YYYY-MM-DD\". 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExploreTheSizeOfYourLogDataAPI.GetOrganizationsLogUsage(context.Background(), logKey).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExploreTheSizeOfYourLogDataAPI.GetOrganizationsLogUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationsLogUsage`: LogUsageResponse
	fmt.Fprintf(os.Stdout, "Response from `ExploreTheSizeOfYourLogDataAPI.GetOrganizationsLogUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logKey** | **string** | \\ The key of the log to be queried.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationsLogUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | \\ The start of the time range as a string in the format \&quot;YYYY-MM-DD\&quot;.  | 
 **to** | **string** | \\ The end of the time range as a string in the format \&quot;YYYY-MM-DD\&quot;.  | 

### Return type

[**LogUsageResponse**](LogUsageResponse.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationsLogsUsage

> LogsUsageResponse GetOrganizationsLogsUsage(ctx).From(from).To(to).TimeRange(timeRange).Execute()

Data Size Broken Down By Log



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
	from := "2023-09-14" // string | \\ The start of the time range as a string in the format \"YYYY-MM-DD\". 
	to := "2023-09-20" // string | \\ The end of the time range as a string in the format \"YYYY-MM-DD\". 
	timeRange := "last 5 minutes" // string | \\ An alternative to the `from` and `to` query parameters. Supported values: * `yesterday` * `today` * `last x timeunits` where x is the number of time unit back from the current server time. Supported time units (case insensitive):     - min(s) or minute(s)     - hr(s) or hour(s)     - day(s)     - week(s)     - month(s)     - year(s)  If `time_range` is used, then the `from` and `to` query parameters must not be used.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExploreTheSizeOfYourLogDataAPI.GetOrganizationsLogsUsage(context.Background()).From(from).To(to).TimeRange(timeRange).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExploreTheSizeOfYourLogDataAPI.GetOrganizationsLogsUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationsLogsUsage`: LogsUsageResponse
	fmt.Fprintf(os.Stdout, "Response from `ExploreTheSizeOfYourLogDataAPI.GetOrganizationsLogsUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationsLogsUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | \\ The start of the time range as a string in the format \&quot;YYYY-MM-DD\&quot;.  | 
 **to** | **string** | \\ The end of the time range as a string in the format \&quot;YYYY-MM-DD\&quot;.  | 
 **timeRange** | **string** | \\ An alternative to the &#x60;from&#x60; and &#x60;to&#x60; query parameters. Supported values: * &#x60;yesterday&#x60; * &#x60;today&#x60; * &#x60;last x timeunits&#x60; where x is the number of time unit back from the current server time. Supported time units (case insensitive):     - min(s) or minute(s)     - hr(s) or hour(s)     - day(s)     - week(s)     - month(s)     - year(s)  If &#x60;time_range&#x60; is used, then the &#x60;from&#x60; and &#x60;to&#x60; query parameters must not be used.  | 

### Return type

[**LogsUsageResponse**](LogsUsageResponse.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationsUsage

> AccountsUsageResponse GetOrganizationsUsage(ctx).From(from).To(to).Execute()

Total Data Size Across All Logs



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
	from := "2023-09-14" // string | \\ The start of the time range as a string in the format \"YYYY-MM-DD\". 
	to := "2023-09-20" // string | \\ The end of the time range as a string in the format \"YYYY-MM-DD\". 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExploreTheSizeOfYourLogDataAPI.GetOrganizationsUsage(context.Background()).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExploreTheSizeOfYourLogDataAPI.GetOrganizationsUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationsUsage`: AccountsUsageResponse
	fmt.Fprintf(os.Stdout, "Response from `ExploreTheSizeOfYourLogDataAPI.GetOrganizationsUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationsUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | \\ The start of the time range as a string in the format \&quot;YYYY-MM-DD\&quot;.  | 
 **to** | **string** | \\ The end of the time range as a string in the format \&quot;YYYY-MM-DD\&quot;.  | 

### Return type

[**AccountsUsageResponse**](AccountsUsageResponse.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

