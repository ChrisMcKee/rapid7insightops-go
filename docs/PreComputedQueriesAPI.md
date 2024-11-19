# invoker\PreComputedQueriesAPI

All URIs are relative to *https://eu.rest.logs.insight.rapid7.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MetricsCreate**](PreComputedQueriesAPI.md#MetricsCreate) | **Post** /management/metrics | Create A Pre-Computed Query
[**MetricsDelete**](PreComputedQueriesAPI.md#MetricsDelete) | **Delete** /management/metrics/{metric_id} | Delete A Pre-Computed Query
[**MetricsPut**](PreComputedQueriesAPI.md#MetricsPut) | **Put** /management/metrics/{metric_id} | Replace a Pre-Computed Query
[**MetricsRetrieve**](PreComputedQueriesAPI.md#MetricsRetrieve) | **Get** /management/metrics/{metric_id} | Retrieve A Pre-Computed Query
[**MetricsRetrieveAll**](PreComputedQueriesAPI.md#MetricsRetrieveAll) | **Get** /management/metrics | List All Pre-Computed Queries
[**QueryMetrics**](PreComputedQueriesAPI.md#QueryMetrics) | **Get** /query/metrics/{id} | Fetch Query Results



## MetricsCreate

> MetricsResponseSingle MetricsCreate(ctx).MetricsCreateRequest(metricsCreateRequest).Execute()

Create A Pre-Computed Query



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
	metricsCreateRequest := openapiclient.metrics_create_request{MetricsInput: openapiclient.NewMetricsInput()} // MetricsCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PreComputedQueriesAPI.MetricsCreate(context.Background()).MetricsCreateRequest(metricsCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PreComputedQueriesAPI.MetricsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetricsCreate`: MetricsResponseSingle
	fmt.Fprintf(os.Stdout, "Response from `PreComputedQueriesAPI.MetricsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetricsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metricsCreateRequest** | [**MetricsCreateRequest**](MetricsCreateRequest.md) |  | 

### Return type

[**MetricsResponseSingle**](MetricsResponseSingle.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricsDelete

> MetricsDelete(ctx, metricId).Execute()

Delete A Pre-Computed Query



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
	metricId := "9aff96bd-ef12-4887-93b5-00931se12f1a" // string | \\ The UUID of the pre-computed query. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PreComputedQueriesAPI.MetricsDelete(context.Background(), metricId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PreComputedQueriesAPI.MetricsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricId** | **string** | \\ The UUID of the pre-computed query.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMetricsDeleteRequest struct via the builder pattern


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


## MetricsPut

> MetricsResponseSingle MetricsPut(ctx, metricId).MetricsCreateRequest(metricsCreateRequest).Execute()

Replace a Pre-Computed Query



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
	metricId := "9aff96bd-ef12-4887-93b5-00931se12f1a" // string | \\ The UUID of the pre-computed query. 
	metricsCreateRequest := openapiclient.metrics_create_request{MetricsInput: openapiclient.NewMetricsInput()} // MetricsCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PreComputedQueriesAPI.MetricsPut(context.Background(), metricId).MetricsCreateRequest(metricsCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PreComputedQueriesAPI.MetricsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetricsPut`: MetricsResponseSingle
	fmt.Fprintf(os.Stdout, "Response from `PreComputedQueriesAPI.MetricsPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricId** | **string** | \\ The UUID of the pre-computed query.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMetricsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metricsCreateRequest** | [**MetricsCreateRequest**](MetricsCreateRequest.md) |  | 

### Return type

[**MetricsResponseSingle**](MetricsResponseSingle.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricsRetrieve

> MetricsResponseSingle MetricsRetrieve(ctx, metricId).Execute()

Retrieve A Pre-Computed Query



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
	metricId := "9aff96bd-ef12-4887-93b5-00931se12f1a" // string | \\ The UUID of the pre-computed query. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PreComputedQueriesAPI.MetricsRetrieve(context.Background(), metricId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PreComputedQueriesAPI.MetricsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetricsRetrieve`: MetricsResponseSingle
	fmt.Fprintf(os.Stdout, "Response from `PreComputedQueriesAPI.MetricsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricId** | **string** | \\ The UUID of the pre-computed query.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMetricsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MetricsResponseSingle**](MetricsResponseSingle.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricsRetrieveAll

> MetricsResponseList MetricsRetrieveAll(ctx).Execute()

List All Pre-Computed Queries



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
	resp, r, err := apiClient.PreComputedQueriesAPI.MetricsRetrieveAll(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PreComputedQueriesAPI.MetricsRetrieveAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetricsRetrieveAll`: MetricsResponseList
	fmt.Fprintf(os.Stdout, "Response from `PreComputedQueriesAPI.MetricsRetrieveAll`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMetricsRetrieveAllRequest struct via the builder pattern


### Return type

[**MetricsResponseList**](MetricsResponseList.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryMetrics

> MetricsResponse QueryMetrics(ctx, metricId).From(from).To(to).TimeRange(timeRange).Execute()

Fetch Query Results



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
	metricId := "9aff96bd-ef12-4887-93b5-00931se12f1a" // string | \\ The UUID of the pre-computed query. 
	from := int32(1450557004000) // int32 | \\ The start of the time range for the query, as a UNIX timestamp in milliseconds. 
	to := int32(1460557604000) // int32 | \\ The end of the time range for the query, as a UNIX timestamp in milliseconds. 
	timeRange := "last 5 minutes" // string | \\ An alternative to the `from` and `to` query parameters. Supported values: * `yesterday` * `today` * `last x timeunits` where x is the number of time unit back from the current server time. Supported time units (case insensitive):     - min(s) or minute(s)     - hr(s) or hour(s)     - day(s)     - week(s)     - month(s)     - year(s)  If `time_range` is used, then the `from` and `to` query parameters must not be used.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PreComputedQueriesAPI.QueryMetrics(context.Background(), metricId).From(from).To(to).TimeRange(timeRange).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PreComputedQueriesAPI.QueryMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryMetrics`: MetricsResponse
	fmt.Fprintf(os.Stdout, "Response from `PreComputedQueriesAPI.QueryMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricId** | **string** | \\ The UUID of the pre-computed query.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **int32** | \\ The start of the time range for the query, as a UNIX timestamp in milliseconds.  | 
 **to** | **int32** | \\ The end of the time range for the query, as a UNIX timestamp in milliseconds.  | 
 **timeRange** | **string** | \\ An alternative to the &#x60;from&#x60; and &#x60;to&#x60; query parameters. Supported values: * &#x60;yesterday&#x60; * &#x60;today&#x60; * &#x60;last x timeunits&#x60; where x is the number of time unit back from the current server time. Supported time units (case insensitive):     - min(s) or minute(s)     - hr(s) or hour(s)     - day(s)     - week(s)     - month(s)     - year(s)  If &#x60;time_range&#x60; is used, then the &#x60;from&#x60; and &#x60;to&#x60; query parameters must not be used.  | 

### Return type

[**MetricsResponse**](MetricsResponse.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

