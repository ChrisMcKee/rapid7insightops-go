# invoker\SavedQueriesAPI

All URIs are relative to *https://eu.rest.logs.insight.rapid7.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSavedQueryId**](SavedQueriesAPI.md#DeleteSavedQueryId) | **Delete** /query/saved_queries/{saved_query_id} | Delete A Saved Query
[**GetSavedQueryId**](SavedQueriesAPI.md#GetSavedQueryId) | **Get** /query/saved_queries/{saved_query_id} | Retrieve A Saved Query
[**ListSavedQueries**](SavedQueriesAPI.md#ListSavedQueries) | **Get** /query/saved_queries | List All Saved Queries
[**PatchSavedQueryId**](SavedQueriesAPI.md#PatchSavedQueryId) | **Patch** /query/saved_queries/{saved_query_id} | Modify A Saved Query
[**PostSavedQueryRoot**](SavedQueriesAPI.md#PostSavedQueryRoot) | **Post** /query/saved_queries | Create A Saved Query
[**PutSavedQueryId**](SavedQueriesAPI.md#PutSavedQueryId) | **Put** /query/saved_queries/{saved_query_id} | Replace A Saved Query
[**UseSavedQuery**](SavedQueriesAPI.md#UseSavedQuery) | **Get** /query/saved_query/{saved_query_id} | Use a Saved Query (logs specified)
[**UseSavedQueryNoLogs**](SavedQueriesAPI.md#UseSavedQueryNoLogs) | **Get** /query/logs/{log_keys}/{saved_query_id} | Use a Saved Query (logs unspecified)



## DeleteSavedQueryId

> DeleteSavedQueryId(ctx, savedQueryId).Execute()

Delete A Saved Query



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
	savedQueryId := "00000000-0000-00cf-0000-000000000000" // string | \\ The id of the saved query. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SavedQueriesAPI.DeleteSavedQueryId(context.Background(), savedQueryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SavedQueriesAPI.DeleteSavedQueryId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**savedQueryId** | **string** | \\ The id of the saved query.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSavedQueryIdRequest struct via the builder pattern


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


## GetSavedQueryId

> SavedQuery GetSavedQueryId(ctx, savedQueryId).Execute()

Retrieve A Saved Query



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
	savedQueryId := "00000000-0000-00cf-0000-000000000000" // string | \\ The id of the saved query. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SavedQueriesAPI.GetSavedQueryId(context.Background(), savedQueryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SavedQueriesAPI.GetSavedQueryId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSavedQueryId`: SavedQuery
	fmt.Fprintf(os.Stdout, "Response from `SavedQueriesAPI.GetSavedQueryId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**savedQueryId** | **string** | \\ The id of the saved query.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSavedQueryIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SavedQuery**](SavedQuery.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSavedQueries

> ListSavedQueries200Response ListSavedQueries(ctx).Execute()

List All Saved Queries



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
	resp, r, err := apiClient.SavedQueriesAPI.ListSavedQueries(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SavedQueriesAPI.ListSavedQueries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSavedQueries`: ListSavedQueries200Response
	fmt.Fprintf(os.Stdout, "Response from `SavedQueriesAPI.ListSavedQueries`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSavedQueriesRequest struct via the builder pattern


### Return type

[**ListSavedQueries200Response**](ListSavedQueries200Response.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSavedQueryId

> SavedQuery PatchSavedQueryId(ctx, savedQueryId).PatchSavedQuery(patchSavedQuery).Execute()

Modify A Saved Query



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
	savedQueryId := "00000000-0000-00cf-0000-000000000000" // string | \\ The id of the saved query. 
	patchSavedQuery := *openapiclient.NewPatchSavedQuery() // PatchSavedQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SavedQueriesAPI.PatchSavedQueryId(context.Background(), savedQueryId).PatchSavedQuery(patchSavedQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SavedQueriesAPI.PatchSavedQueryId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSavedQueryId`: SavedQuery
	fmt.Fprintf(os.Stdout, "Response from `SavedQueriesAPI.PatchSavedQueryId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**savedQueryId** | **string** | \\ The id of the saved query.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSavedQueryIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchSavedQuery** | [**PatchSavedQuery**](PatchSavedQuery.md) |  | 

### Return type

[**SavedQuery**](SavedQuery.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSavedQueryRoot

> SavedQuery PostSavedQueryRoot(ctx).CreateSavedQuery(createSavedQuery).Execute()

Create A Saved Query



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
	createSavedQuery := *openapiclient.NewCreateSavedQuery() // CreateSavedQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SavedQueriesAPI.PostSavedQueryRoot(context.Background()).CreateSavedQuery(createSavedQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SavedQueriesAPI.PostSavedQueryRoot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSavedQueryRoot`: SavedQuery
	fmt.Fprintf(os.Stdout, "Response from `SavedQueriesAPI.PostSavedQueryRoot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSavedQueryRootRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSavedQuery** | [**CreateSavedQuery**](CreateSavedQuery.md) |  | 

### Return type

[**SavedQuery**](SavedQuery.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSavedQueryId

> SavedQuery PutSavedQueryId(ctx, savedQueryId).CreateSavedQuery(createSavedQuery).Execute()

Replace A Saved Query



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
	savedQueryId := "00000000-0000-00cf-0000-000000000000" // string | \\ The id of the saved query. 
	createSavedQuery := *openapiclient.NewCreateSavedQuery() // CreateSavedQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SavedQueriesAPI.PutSavedQueryId(context.Background(), savedQueryId).CreateSavedQuery(createSavedQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SavedQueriesAPI.PutSavedQueryId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSavedQueryId`: SavedQuery
	fmt.Fprintf(os.Stdout, "Response from `SavedQueriesAPI.PutSavedQueryId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**savedQueryId** | **string** | \\ The id of the saved query.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSavedQueryIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createSavedQuery** | [**CreateSavedQuery**](CreateSavedQuery.md) |  | 

### Return type

[**SavedQuery**](SavedQuery.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UseSavedQuery

> GetQueryLogs200Response UseSavedQuery(ctx, savedQueryId).From(from).To(to).TimeRange(timeRange).PerPage(perPage).KvpInfo(kvpInfo).MostRecentFirst(mostRecentFirst).Execute()

Use a Saved Query (logs specified)



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
	savedQueryId := "00000000-0000-00cf-0000-000000000000" // string | \\ The id of the saved query. 
	from := int32(1450557004000) // int32 | \\ The start of the time range for the query, as a UNIX timestamp in milliseconds. 
	to := int32(1460557604000) // int32 | \\ The end of the time range for the query, as a UNIX timestamp in milliseconds. 
	timeRange := "last 5 minutes" // string | \\ An alternative to the `from` and `to` query parameters. Supported values: * `yesterday` * `today` * `last x timeunits` where x is the number of time unit back from the current server time. Supported time units (case insensitive):     - min(s) or minute(s)     - hr(s) or hour(s)     - day(s)     - week(s)     - month(s)     - year(s)  If `time_range` is used, then the `from` and `to` query parameters must not be used.  (optional)
	perPage := int32(100) // int32 | \\ Number of log entries to return per page, up to 500(the maximum allowed).  (optional) (default to 50)
	kvpInfo := true // bool | \\ When set to true, the `events` object that is returned will additionally contain information about all the key-value pairs in each returned log entry.  (optional)
	mostRecentFirst := true // bool | \\ When set to `true`, the query returns the most recent events first. When set to `false`, it returns the oldest events first.  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SavedQueriesAPI.UseSavedQuery(context.Background(), savedQueryId).From(from).To(to).TimeRange(timeRange).PerPage(perPage).KvpInfo(kvpInfo).MostRecentFirst(mostRecentFirst).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SavedQueriesAPI.UseSavedQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UseSavedQuery`: GetQueryLogs200Response
	fmt.Fprintf(os.Stdout, "Response from `SavedQueriesAPI.UseSavedQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**savedQueryId** | **string** | \\ The id of the saved query.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUseSavedQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **int32** | \\ The start of the time range for the query, as a UNIX timestamp in milliseconds.  | 
 **to** | **int32** | \\ The end of the time range for the query, as a UNIX timestamp in milliseconds.  | 
 **timeRange** | **string** | \\ An alternative to the &#x60;from&#x60; and &#x60;to&#x60; query parameters. Supported values: * &#x60;yesterday&#x60; * &#x60;today&#x60; * &#x60;last x timeunits&#x60; where x is the number of time unit back from the current server time. Supported time units (case insensitive):     - min(s) or minute(s)     - hr(s) or hour(s)     - day(s)     - week(s)     - month(s)     - year(s)  If &#x60;time_range&#x60; is used, then the &#x60;from&#x60; and &#x60;to&#x60; query parameters must not be used.  | 
 **perPage** | **int32** | \\ Number of log entries to return per page, up to 500(the maximum allowed).  | [default to 50]
 **kvpInfo** | **bool** | \\ When set to true, the &#x60;events&#x60; object that is returned will additionally contain information about all the key-value pairs in each returned log entry.  | 
 **mostRecentFirst** | **bool** | \\ When set to &#x60;true&#x60;, the query returns the most recent events first. When set to &#x60;false&#x60;, it returns the oldest events first.  | [default to false]

### Return type

[**GetQueryLogs200Response**](GetQueryLogs200Response.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UseSavedQueryNoLogs

> GetQueryLogs200Response UseSavedQueryNoLogs(ctx, logKeys, savedQueryId).From(from).To(to).TimeRange(timeRange).PerPage(perPage).KvpInfo(kvpInfo).MostRecentFirst(mostRecentFirst).Execute()

Use a Saved Query (logs unspecified)



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
	logKeys := "/query/logs/565c1b7b-c08b-4c87-a42a-ab08bad56071" // string | \\ The keys of the logs to be queried, separated by the ':' character. 
	savedQueryId := "00000000-0000-00cf-0000-000000000000" // string | \\ The id of the saved query. 
	from := int32(1450557004000) // int32 | \\ The start of the time range for the query, as a UNIX timestamp in milliseconds. 
	to := int32(1460557604000) // int32 | \\ The end of the time range for the query, as a UNIX timestamp in milliseconds. 
	timeRange := "last 5 minutes" // string | \\ An alternative to the `from` and `to` query parameters. Supported values: * `yesterday` * `today` * `last x timeunits` where x is the number of time unit back from the current server time. Supported time units (case insensitive):     - min(s) or minute(s)     - hr(s) or hour(s)     - day(s)     - week(s)     - month(s)     - year(s)  If `time_range` is used, then the `from` and `to` query parameters must not be used.  (optional)
	perPage := int32(100) // int32 | \\ Number of log entries to return per page, up to 500(the maximum allowed).  (optional) (default to 50)
	kvpInfo := true // bool | \\ When set to true, the `events` object that is returned will additionally contain information about all the key-value pairs in each returned log entry.  (optional)
	mostRecentFirst := true // bool | \\ When set to `true`, the query returns the most recent events first. When set to `false`, it returns the oldest events first.  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SavedQueriesAPI.UseSavedQueryNoLogs(context.Background(), logKeys, savedQueryId).From(from).To(to).TimeRange(timeRange).PerPage(perPage).KvpInfo(kvpInfo).MostRecentFirst(mostRecentFirst).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SavedQueriesAPI.UseSavedQueryNoLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UseSavedQueryNoLogs`: GetQueryLogs200Response
	fmt.Fprintf(os.Stdout, "Response from `SavedQueriesAPI.UseSavedQueryNoLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logKeys** | **string** | \\ The keys of the logs to be queried, separated by the &#39;:&#39; character.  | 
**savedQueryId** | **string** | \\ The id of the saved query.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUseSavedQueryNoLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **from** | **int32** | \\ The start of the time range for the query, as a UNIX timestamp in milliseconds.  | 
 **to** | **int32** | \\ The end of the time range for the query, as a UNIX timestamp in milliseconds.  | 
 **timeRange** | **string** | \\ An alternative to the &#x60;from&#x60; and &#x60;to&#x60; query parameters. Supported values: * &#x60;yesterday&#x60; * &#x60;today&#x60; * &#x60;last x timeunits&#x60; where x is the number of time unit back from the current server time. Supported time units (case insensitive):     - min(s) or minute(s)     - hr(s) or hour(s)     - day(s)     - week(s)     - month(s)     - year(s)  If &#x60;time_range&#x60; is used, then the &#x60;from&#x60; and &#x60;to&#x60; query parameters must not be used.  | 
 **perPage** | **int32** | \\ Number of log entries to return per page, up to 500(the maximum allowed).  | [default to 50]
 **kvpInfo** | **bool** | \\ When set to true, the &#x60;events&#x60; object that is returned will additionally contain information about all the key-value pairs in each returned log entry.  | 
 **mostRecentFirst** | **bool** | \\ When set to &#x60;true&#x60;, the query returns the most recent events first. When set to &#x60;false&#x60;, it returns the oldest events first.  | [default to false]

### Return type

[**GetQueryLogs200Response**](GetQueryLogs200Response.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

