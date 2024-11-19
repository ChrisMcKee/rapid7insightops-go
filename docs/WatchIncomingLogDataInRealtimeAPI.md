# invoker\WatchIncomingLogDataInRealtimeAPI

All URIs are relative to *https://eu.rest.logs.insight.rapid7.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLiveLogs**](WatchIncomingLogDataInRealtimeAPI.md#GetLiveLogs) | **Get** /query/live/logs/{log_keys} | Start A Live Tail Feed (Individual Logs)
[**LiveSavedQuery**](WatchIncomingLogDataInRealtimeAPI.md#LiveSavedQuery) | **Get** /query/live/saved_query/{saved_query_id} | Start a Live Tail Feed (Saved Query - with logs)
[**LiveSavedQueryNoLogs**](WatchIncomingLogDataInRealtimeAPI.md#LiveSavedQueryNoLogs) | **Get** /query/live/logs/{log_keys}/{saved_query_id} | Start a Live Tail Feed (Saved Query - without logs)
[**PollLiveTail**](WatchIncomingLogDataInRealtimeAPI.md#PollLiveTail) | **Get** /query/live/{id} | Poll a Live Tail Feed
[**PostLiveLogs**](WatchIncomingLogDataInRealtimeAPI.md#PostLiveLogs) | **Post** /query/live/logs | Start A Live Tail Feed (Multiple Logs)



## GetLiveLogs

> LiveTailResponse GetLiveLogs(ctx, logKeys).Query(query).Mode(mode).Label(label).Labels(labels).PerPage(perPage).Execute()

Start A Live Tail Feed (Individual Logs)



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
	logKeys := "/query/live/logs/565c1b7b-c08b-4c87-a42a-ab08bad56071" // string | \\ The keys of the logs for the Live Tail Feed.  Supplying multiple ':' separated log keys is deprecated; users should instead 'POST' the query as a JSON payload to [/query/live/logs](#operation/postLiveLogs). 
	query := "where(foo=bar)" // string | \\ A valid LEQL query for the Live Tail feed. If omitted, the query matches all log entries.  *Live Tail does not support 'calculate' or 'groupby' queries.*  (optional)
	mode := "head" // string | \\ The order the log entries are returned in. Possible values: * `tail` (which returns entries in the order: newest to oldest.) * `head` (which returns entries in the order: oldest to newest.)  (optional) (default to "tail")
	label := "00000000-0000-0000-0000-000000000001" // string | \\ Only entries which have a label with this UUID will be returned. This only works with non-statistical queries (i.e. no 'groupby' or 'calculate' clauses).  (optional)
	labels := "00000000-0000-0000-0000-000000000001:21b21bb8-8869-4e2c-98df-684892e4e112" // string | \\ A set of ':' separated label UUIDs. Only entries which have a label matching one of these UUIDs will be returned. This only works with non-statistical queries (i.e. no 'groupby' or 'calculate' clauses). Takes precedence over the 'label' parameter.  (optional)
	perPage := int32(100) // int32 | \\ The maximum number of log entries to return at a time, up to 500(the maximum allowed). If the number of results exceed this value, the next \"page\" of results is returned the next time that the Live Tail feed is polled.  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WatchIncomingLogDataInRealtimeAPI.GetLiveLogs(context.Background(), logKeys).Query(query).Mode(mode).Label(label).Labels(labels).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WatchIncomingLogDataInRealtimeAPI.GetLiveLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLiveLogs`: LiveTailResponse
	fmt.Fprintf(os.Stdout, "Response from `WatchIncomingLogDataInRealtimeAPI.GetLiveLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logKeys** | **string** | \\ The keys of the logs for the Live Tail Feed.  Supplying multiple &#39;:&#39; separated log keys is deprecated; users should instead &#39;POST&#39; the query as a JSON payload to [/query/live/logs](#operation/postLiveLogs).  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLiveLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | **string** | \\ A valid LEQL query for the Live Tail feed. If omitted, the query matches all log entries.  *Live Tail does not support &#39;calculate&#39; or &#39;groupby&#39; queries.*  | 
 **mode** | **string** | \\ The order the log entries are returned in. Possible values: * &#x60;tail&#x60; (which returns entries in the order: newest to oldest.) * &#x60;head&#x60; (which returns entries in the order: oldest to newest.)  | [default to &quot;tail&quot;]
 **label** | **string** | \\ Only entries which have a label with this UUID will be returned. This only works with non-statistical queries (i.e. no &#39;groupby&#39; or &#39;calculate&#39; clauses).  | 
 **labels** | **string** | \\ A set of &#39;:&#39; separated label UUIDs. Only entries which have a label matching one of these UUIDs will be returned. This only works with non-statistical queries (i.e. no &#39;groupby&#39; or &#39;calculate&#39; clauses). Takes precedence over the &#39;label&#39; parameter.  | 
 **perPage** | **int32** | \\ The maximum number of log entries to return at a time, up to 500(the maximum allowed). If the number of results exceed this value, the next \&quot;page\&quot; of results is returned the next time that the Live Tail feed is polled.  | [default to 50]

### Return type

[**LiveTailResponse**](LiveTailResponse.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LiveSavedQuery

> LiveTailResponse LiveSavedQuery(ctx, savedQueryId).PerPage(perPage).Mode(mode).Execute()

Start a Live Tail Feed (Saved Query - with logs)



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
	perPage := int32(100) // int32 | \\ The maximum number of log entries to return at a time, up to 500(the maximum allowed). If the number of results exceed this value, the next \"page\" of results is returned the next time that the Live Tail feed is polled.  (optional) (default to 50)
	mode := "head" // string | \\ The order the log entries are returned in. Possible values: * `tail` (which returns entries in the order: newest to oldest.) * `head` (which returns entries in the order: oldest to newest.)  (optional) (default to "tail")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WatchIncomingLogDataInRealtimeAPI.LiveSavedQuery(context.Background(), savedQueryId).PerPage(perPage).Mode(mode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WatchIncomingLogDataInRealtimeAPI.LiveSavedQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LiveSavedQuery`: LiveTailResponse
	fmt.Fprintf(os.Stdout, "Response from `WatchIncomingLogDataInRealtimeAPI.LiveSavedQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**savedQueryId** | **string** | \\ The id of the saved query.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLiveSavedQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | \\ The maximum number of log entries to return at a time, up to 500(the maximum allowed). If the number of results exceed this value, the next \&quot;page\&quot; of results is returned the next time that the Live Tail feed is polled.  | [default to 50]
 **mode** | **string** | \\ The order the log entries are returned in. Possible values: * &#x60;tail&#x60; (which returns entries in the order: newest to oldest.) * &#x60;head&#x60; (which returns entries in the order: oldest to newest.)  | [default to &quot;tail&quot;]

### Return type

[**LiveTailResponse**](LiveTailResponse.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LiveSavedQueryNoLogs

> LiveTailResponse LiveSavedQueryNoLogs(ctx, logKeys, savedQueryId).PerPage(perPage).Mode(mode).Execute()

Start a Live Tail Feed (Saved Query - without logs)



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
	logKeys := "/query/live/logs/565c1b7b-c08b-4c87-a42a-ab08bad56071" // string | \\ The keys of the logs for the Live Tail Feed, separated by the ':' character. 
	savedQueryId := "00000000-0000-00cf-0000-000000000000" // string | \\ The id of the saved query. 
	perPage := int32(100) // int32 | \\ The maximum number of log entries to return at a time, up to 500(the maximum allowed). If the number of results exceed this value, the next \"page\" of results is returned the next time that the Live Tail feed is polled.  (optional) (default to 50)
	mode := "head" // string | \\ The order the log entries are returned in. Possible values: * `tail` (which returns entries in the order: newest to oldest.) * `head` (which returns entries in the order: oldest to newest.)  (optional) (default to "tail")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WatchIncomingLogDataInRealtimeAPI.LiveSavedQueryNoLogs(context.Background(), logKeys, savedQueryId).PerPage(perPage).Mode(mode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WatchIncomingLogDataInRealtimeAPI.LiveSavedQueryNoLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LiveSavedQueryNoLogs`: LiveTailResponse
	fmt.Fprintf(os.Stdout, "Response from `WatchIncomingLogDataInRealtimeAPI.LiveSavedQueryNoLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logKeys** | **string** | \\ The keys of the logs for the Live Tail Feed, separated by the &#39;:&#39; character.  | 
**savedQueryId** | **string** | \\ The id of the saved query.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLiveSavedQueryNoLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | \\ The maximum number of log entries to return at a time, up to 500(the maximum allowed). If the number of results exceed this value, the next \&quot;page\&quot; of results is returned the next time that the Live Tail feed is polled.  | [default to 50]
 **mode** | **string** | \\ The order the log entries are returned in. Possible values: * &#x60;tail&#x60; (which returns entries in the order: newest to oldest.) * &#x60;head&#x60; (which returns entries in the order: oldest to newest.)  | [default to &quot;tail&quot;]

### Return type

[**LiveTailResponse**](LiveTailResponse.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PollLiveTail

> LiveTailPollResponse PollLiveTail(ctx, id).Query(query).LogKeys(logKeys).Execute()

Poll a Live Tail Feed



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
	id := "c9d059da-7b1b-48a9-aaeb-edf2aa3e3fd9:2:::0447214e1b4e5b2041bf62bd7421cb9b7069065f:" // string | \\ The polling id generated when the Live Tail Feed started. 
	query := "where(foo=bar)" // string | \\ A valid LEQL query for the Live Tail feed. If omitted, the query matches all log entries.  *Live Tail does not support 'calculate' or 'groupby' queries.*  (optional)
	logKeys := "565c1b7b-c08b-4c87-a42a-ab08bad56071" // string | \\ The keys of the logs for the Live Tail Feed, separated by the ':' character.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WatchIncomingLogDataInRealtimeAPI.PollLiveTail(context.Background(), id).Query(query).LogKeys(logKeys).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WatchIncomingLogDataInRealtimeAPI.PollLiveTail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PollLiveTail`: LiveTailPollResponse
	fmt.Fprintf(os.Stdout, "Response from `WatchIncomingLogDataInRealtimeAPI.PollLiveTail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | \\ The polling id generated when the Live Tail Feed started.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPollLiveTailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | **string** | \\ A valid LEQL query for the Live Tail feed. If omitted, the query matches all log entries.  *Live Tail does not support &#39;calculate&#39; or &#39;groupby&#39; queries.*  | 
 **logKeys** | **string** | \\ The keys of the logs for the Live Tail Feed, separated by the &#39;:&#39; character.  | 

### Return type

[**LiveTailPollResponse**](LiveTailPollResponse.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLiveLogs

> LiveTailResponse PostLiveLogs(ctx).PostLiveLogsRequest(postLiveLogsRequest).Labels(labels).Execute()

Start A Live Tail Feed (Multiple Logs)



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
	postLiveLogsRequest := *openapiclient.NewPostLiveLogsRequest([]string{"Logs_example"}) // PostLiveLogsRequest | 
	labels := "00000000-0000-0000-0000-000000000001:21b21bb8-8869-4e2c-98df-684892e4e112" // string | \\ A set of ':' separated label UUIDs. Only entries which have a label matching one of these UUIDs will be returned. This only works with non-statistical queries (i.e. no 'groupby' or 'calculate' clauses). Takes precedence over the 'label' parameter.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WatchIncomingLogDataInRealtimeAPI.PostLiveLogs(context.Background()).PostLiveLogsRequest(postLiveLogsRequest).Labels(labels).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WatchIncomingLogDataInRealtimeAPI.PostLiveLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLiveLogs`: LiveTailResponse
	fmt.Fprintf(os.Stdout, "Response from `WatchIncomingLogDataInRealtimeAPI.PostLiveLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLiveLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postLiveLogsRequest** | [**PostLiveLogsRequest**](PostLiveLogsRequest.md) |  | 
 **labels** | **string** | \\ A set of &#39;:&#39; separated label UUIDs. Only entries which have a label matching one of these UUIDs will be returned. This only works with non-statistical queries (i.e. no &#39;groupby&#39; or &#39;calculate&#39; clauses). Takes precedence over the &#39;label&#39; parameter.  | 

### Return type

[**LiveTailResponse**](LiveTailResponse.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

