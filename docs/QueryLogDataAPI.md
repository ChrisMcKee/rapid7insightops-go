# invoker\QueryLogDataAPI

All URIs are relative to *https://eu.rest.logs.insight.rapid7.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetContextEvents**](QueryLogDataAPI.md#GetContextEvents) | **Get** /query/context/{id} | Retrieve The Log Data Immediately Before And After A Specific Log Line
[**GetQueryEndpoints**](QueryLogDataAPI.md#GetQueryEndpoints) | **Get** /query | List All Endpoints
[**GetQueryLogs**](QueryLogDataAPI.md#GetQueryLogs) | **Get** /query/logs/{log_keys} | Query Individual Logs
[**GetSearchStats**](QueryLogDataAPI.md#GetSearchStats) | **Get** /search-stats | View Statistics On Past Queries
[**PollQuery**](QueryLogDataAPI.md#PollQuery) | **Get** /query/{id} | Poll a Query In Progress
[**PostQueryLogs**](QueryLogDataAPI.md#PostQueryLogs) | **Post** /query/logs | Query Multiple Logs
[**QueryLogsetsById**](QueryLogDataAPI.md#QueryLogsetsById) | **Get** /query/logsets/{id} | Query Individual Log Sets
[**QueryLogsetsByName**](QueryLogDataAPI.md#QueryLogsetsByName) | **Get** /query/logsets | Query Multiple Log Sets



## GetContextEvents

> ContextResponse GetContextEvents(ctx, idSequenceNumber).Timestamp(timestamp).LogKeys(logKeys).ContextType(contextType).PerPage(perPage).KvpInfo(kvpInfo).MostRecentFirst(mostRecentFirst).Execute()

Retrieve The Log Data Immediately Before And After A Specific Log Line



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
	idSequenceNumber := int32(2134329353399137024) // int32 | \\ The sequence number of the log entry to fetch contextual log entries for (a unique identifier used to distinguish between log entries received in the same millisecond).  The sequence number can be found next to the log entry in the response body for an Event Search query. For more context on creating an Event Search query [view the introduction to this section](#tag/Query-Log-Data). 
	timestamp := "1609607464185" // string | \\ The timestamp of the log entry to fetch contextual events for. 
	logKeys := "565c1b7b-c08b-4c87-a42a-ab08bad56071" // string | \\ The key of the log containing the log entry to fetch the contextual events for. 
	contextType := "SURROUND" // string | Possible values: * `AFTER`  -  returns the log entries immediately *after* the given log entry (doesn't include the given log entry). * `BEFORE` -  returns the log entries immediately *before* the given log entry (doesn't include the given log entry). * `SURROUND` -  returns the log entries immediately *before and after* the given log entry(*does* include the given log entry). 
	perPage := int32(100) // int32 | \\ Number of log entries to return per page, up to 500(the maximum allowed).  (optional) (default to 50)
	kvpInfo := true // bool | \\ When set to true, the `events` object that is returned will additionally contain information about all the key-value pairs in each returned log entry.  (optional)
	mostRecentFirst := true // bool | \\ When set to `true`, the query returns the most recent events first. When set to `false`, it returns the oldest events first.  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueryLogDataAPI.GetContextEvents(context.Background(), idSequenceNumber).Timestamp(timestamp).LogKeys(logKeys).ContextType(contextType).PerPage(perPage).KvpInfo(kvpInfo).MostRecentFirst(mostRecentFirst).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueryLogDataAPI.GetContextEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContextEvents`: ContextResponse
	fmt.Fprintf(os.Stdout, "Response from `QueryLogDataAPI.GetContextEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idSequenceNumber** | **int32** | \\ The sequence number of the log entry to fetch contextual log entries for (a unique identifier used to distinguish between log entries received in the same millisecond).  The sequence number can be found next to the log entry in the response body for an Event Search query. For more context on creating an Event Search query [view the introduction to this section](#tag/Query-Log-Data).  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContextEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timestamp** | **string** | \\ The timestamp of the log entry to fetch contextual events for.  | 
 **logKeys** | **string** | \\ The key of the log containing the log entry to fetch the contextual events for.  | 
 **contextType** | **string** | Possible values: * &#x60;AFTER&#x60;  -  returns the log entries immediately *after* the given log entry (doesn&#39;t include the given log entry). * &#x60;BEFORE&#x60; -  returns the log entries immediately *before* the given log entry (doesn&#39;t include the given log entry). * &#x60;SURROUND&#x60; -  returns the log entries immediately *before and after* the given log entry(*does* include the given log entry).  | 
 **perPage** | **int32** | \\ Number of log entries to return per page, up to 500(the maximum allowed).  | [default to 50]
 **kvpInfo** | **bool** | \\ When set to true, the &#x60;events&#x60; object that is returned will additionally contain information about all the key-value pairs in each returned log entry.  | 
 **mostRecentFirst** | **bool** | \\ When set to &#x60;true&#x60;, the query returns the most recent events first. When set to &#x60;false&#x60;, it returns the oldest events first.  | [default to false]

### Return type

[**ContextResponse**](ContextResponse.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQueryEndpoints

> map[string]ListEndpointsResponseValue GetQueryEndpoints(ctx).Execute()

List All Endpoints

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
	resp, r, err := apiClient.QueryLogDataAPI.GetQueryEndpoints(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueryLogDataAPI.GetQueryEndpoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQueryEndpoints`: map[string]ListEndpointsResponseValue
	fmt.Fprintf(os.Stdout, "Response from `QueryLogDataAPI.GetQueryEndpoints`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetQueryEndpointsRequest struct via the builder pattern


### Return type

[**map[string]ListEndpointsResponseValue**](ListEndpointsResponseValue.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQueryLogs

> GetQueryLogs200Response GetQueryLogs(ctx, logKeys).From(from).To(to).Query(query).TimeRange(timeRange).Label(label).Labels(labels).PerPage(perPage).ExportFormat(exportFormat).KvpInfo(kvpInfo).MostRecentFirst(mostRecentFirst).SequenceNumber(sequenceNumber).Execute()

Query Individual Logs



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
	logKeys := "/query/logs/565c1b7b-c08b-4c87-a42a-ab08bad56071" // string | \\ The key of the log to be queried.  Supplying multiple ':' separated log keys is deprecated; users should instead 'POST' the query as a JSON payload to [/query/logs](#operation/postQueryLogs). 
	from := int32(1450557004000) // int32 | \\ The start of the time range for the query, as a UNIX timestamp in milliseconds. 
	to := int32(1460557604000) // int32 | \\ The end of the time range for the query, as a UNIX timestamp in milliseconds. 
	query := "where(foo=bar)" // string | \\ A valid LEQL query to run against the log. If omitted, the query retrieves all log entries in the specified time range.  (optional)
	timeRange := "last 5 minutes" // string | \\ An alternative to the `from` and `to` query parameters. Supported values: * `yesterday` * `today` * `last x timeunits` where x is the number of time unit back from the current server time. Supported time units (case insensitive):     - min(s) or minute(s)     - hr(s) or hour(s)     - day(s)     - week(s)     - month(s)     - year(s)  If `time_range` is used, then the `from` and `to` query parameters must not be used.  (optional)
	label := "00000000-0000-0000-0000-000000000001" // string | \\ Only entries which have a label with this UUID will be returned. This only works with non-statistical queries (i.e. no 'groupby' or 'calculate' clauses).  (optional)
	labels := "00000000-0000-0000-0000-000000000001:21b21bb8-8869-4e2c-98df-684892e4e112" // string | \\ A set of ':' separated label UUIDs. Only entries which have a label matching one of these UUIDs will be returned. This only works with non-statistical queries (i.e. no 'groupby' or 'calculate' clauses). Takes precedence over the 'label' parameter.  (optional)
	perPage := int32(100) // int32 | \\ Number of log entries to return per page, up to 500(the maximum allowed).  (optional) (default to 50)
	exportFormat := "csv" // string | \\ If included, the query results will be exported to the given format.  * Currently only `csv` is supported. * This parameter is only for non-statistical search queries (i.e. no 'calculate' and/or 'groupby' clauses). * Results are limited to the first 1 million log entries, and only one export job may run per account at a time.  The response will be a 202, and the response body will contain a link for polling the export job (on the **_/exports/{id}** endpoint).  (optional)
	kvpInfo := true // bool | \\ When set to true, the `events` object that is returned will additionally contain information about all the key-value pairs in each returned log entry.  (optional)
	mostRecentFirst := true // bool | \\ When set to `true`, the query returns the most recent events first. When set to `false`, it returns the oldest events first.  (optional) (default to false)
	sequenceNumber := int32(56) // int32 | \\ If this query parameter is included, the query results will additionally include all log entries received in the `from` millisecond which have sequence numbers larger than the one specified.  Sequence numbers are identifiers used to distinguish between log entries received in the same millisecond. If a log entry was split up into several log entries during ingestion, then those chunks are ordered by sequence number.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueryLogDataAPI.GetQueryLogs(context.Background(), logKeys).From(from).To(to).Query(query).TimeRange(timeRange).Label(label).Labels(labels).PerPage(perPage).ExportFormat(exportFormat).KvpInfo(kvpInfo).MostRecentFirst(mostRecentFirst).SequenceNumber(sequenceNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueryLogDataAPI.GetQueryLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQueryLogs`: GetQueryLogs200Response
	fmt.Fprintf(os.Stdout, "Response from `QueryLogDataAPI.GetQueryLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logKeys** | **string** | \\ The key of the log to be queried.  Supplying multiple &#39;:&#39; separated log keys is deprecated; users should instead &#39;POST&#39; the query as a JSON payload to [/query/logs](#operation/postQueryLogs).  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQueryLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **int32** | \\ The start of the time range for the query, as a UNIX timestamp in milliseconds.  | 
 **to** | **int32** | \\ The end of the time range for the query, as a UNIX timestamp in milliseconds.  | 
 **query** | **string** | \\ A valid LEQL query to run against the log. If omitted, the query retrieves all log entries in the specified time range.  | 
 **timeRange** | **string** | \\ An alternative to the &#x60;from&#x60; and &#x60;to&#x60; query parameters. Supported values: * &#x60;yesterday&#x60; * &#x60;today&#x60; * &#x60;last x timeunits&#x60; where x is the number of time unit back from the current server time. Supported time units (case insensitive):     - min(s) or minute(s)     - hr(s) or hour(s)     - day(s)     - week(s)     - month(s)     - year(s)  If &#x60;time_range&#x60; is used, then the &#x60;from&#x60; and &#x60;to&#x60; query parameters must not be used.  | 
 **label** | **string** | \\ Only entries which have a label with this UUID will be returned. This only works with non-statistical queries (i.e. no &#39;groupby&#39; or &#39;calculate&#39; clauses).  | 
 **labels** | **string** | \\ A set of &#39;:&#39; separated label UUIDs. Only entries which have a label matching one of these UUIDs will be returned. This only works with non-statistical queries (i.e. no &#39;groupby&#39; or &#39;calculate&#39; clauses). Takes precedence over the &#39;label&#39; parameter.  | 
 **perPage** | **int32** | \\ Number of log entries to return per page, up to 500(the maximum allowed).  | [default to 50]
 **exportFormat** | **string** | \\ If included, the query results will be exported to the given format.  * Currently only &#x60;csv&#x60; is supported. * This parameter is only for non-statistical search queries (i.e. no &#39;calculate&#39; and/or &#39;groupby&#39; clauses). * Results are limited to the first 1 million log entries, and only one export job may run per account at a time.  The response will be a 202, and the response body will contain a link for polling the export job (on the **_/exports/{id}** endpoint).  | 
 **kvpInfo** | **bool** | \\ When set to true, the &#x60;events&#x60; object that is returned will additionally contain information about all the key-value pairs in each returned log entry.  | 
 **mostRecentFirst** | **bool** | \\ When set to &#x60;true&#x60;, the query returns the most recent events first. When set to &#x60;false&#x60;, it returns the oldest events first.  | [default to false]
 **sequenceNumber** | **int32** | \\ If this query parameter is included, the query results will additionally include all log entries received in the &#x60;from&#x60; millisecond which have sequence numbers larger than the one specified.  Sequence numbers are identifiers used to distinguish between log entries received in the same millisecond. If a log entry was split up into several log entries during ingestion, then those chunks are ordered by sequence number.  | 

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


## GetSearchStats

> []SearchStatsInner GetSearchStats(ctx).Execute()

View Statistics On Past Queries



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
	resp, r, err := apiClient.QueryLogDataAPI.GetSearchStats(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueryLogDataAPI.GetSearchStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSearchStats`: []SearchStatsInner
	fmt.Fprintf(os.Stdout, "Response from `QueryLogDataAPI.GetSearchStats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSearchStatsRequest struct via the builder pattern


### Return type

[**[]SearchStatsInner**](SearchStatsInner.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PollQuery

> PollQuery200Response PollQuery(ctx, id).TimeRange(timeRange).Execute()

Poll a Query In Progress



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
	id := "0bnc17f4-bbbe-46b1-a843-163c074bd1ad:0:c5be1c97f925ff772263661182974903b83ef305:50:f89f7d5cbeb270abeb18b1c28262cee66047d21d:" // string | \\ The continuation id generated when the query started. 
	timeRange := "last 5 minutes" // string | \\ An alternative to the `from` and `to` query parameters. Supported values: * `yesterday` * `today` * `last x timeunits` where x is the number of time unit back from the current server time. Supported time units (case insensitive):     - min(s) or minute(s)     - hr(s) or hour(s)     - day(s)     - week(s)     - month(s)     - year(s)  If `time_range` is used, then the `from` and `to` query parameters must not be used.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueryLogDataAPI.PollQuery(context.Background(), id).TimeRange(timeRange).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueryLogDataAPI.PollQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PollQuery`: PollQuery200Response
	fmt.Fprintf(os.Stdout, "Response from `QueryLogDataAPI.PollQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | \\ The continuation id generated when the query started.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPollQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timeRange** | **string** | \\ An alternative to the &#x60;from&#x60; and &#x60;to&#x60; query parameters. Supported values: * &#x60;yesterday&#x60; * &#x60;today&#x60; * &#x60;last x timeunits&#x60; where x is the number of time unit back from the current server time. Supported time units (case insensitive):     - min(s) or minute(s)     - hr(s) or hour(s)     - day(s)     - week(s)     - month(s)     - year(s)  If &#x60;time_range&#x60; is used, then the &#x60;from&#x60; and &#x60;to&#x60; query parameters must not be used.  | 

### Return type

[**PollQuery200Response**](PollQuery200Response.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostQueryLogs

> GetQueryLogs200Response PostQueryLogs(ctx).PostQueryLogsRequest(postQueryLogsRequest).PerPage(perPage).ExportFormat(exportFormat).KvpInfo(kvpInfo).MostRecentFirst(mostRecentFirst).Labels(labels).SequenceNumber(sequenceNumber).Execute()

Query Multiple Logs



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
	postQueryLogsRequest := *openapiclient.NewPostQueryLogsRequest() // PostQueryLogsRequest | The JSON payload must contain: * a non-empty array of the log keys for the logs that will be queried * a `leql` object, containing a value for either of the following (but **not** for both):   - `leql.during.time_range`   - `leql.during.from` and `leql.during.to`    The `leql.statement` can be empty or omitted. 
	perPage := int32(100) // int32 | \\ Number of log entries to return per page, up to 500(the maximum allowed).  (optional) (default to 50)
	exportFormat := "csv" // string | \\ If included, the query results will be exported to the given format.  * Currently only `csv` is supported. * This parameter is only for non-statistical search queries (i.e. no 'calculate' and/or 'groupby' clauses). * Results are limited to the first 1 million log entries, and only one export job may run per account at a time.  The response will be a 202, and the response body will contain a link for polling the export job (on the **_/exports/{id}** endpoint).  (optional)
	kvpInfo := true // bool | \\ When set to true, the `events` object that is returned will additionally contain information about all the key-value pairs in each returned log entry.  (optional)
	mostRecentFirst := true // bool | \\ When set to `true`, the query returns the most recent events first. When set to `false`, it returns the oldest events first.  (optional) (default to false)
	labels := "00000000-0000-0000-0000-000000000001:21b21bb8-8869-4e2c-98df-684892e4e112" // string | \\ A set of ':' separated label UUIDs. Only entries which have a label matching one of these UUIDs will be returned. This only works with non-statistical queries (i.e. no 'groupby' or 'calculate' clauses). Takes precedence over the 'label' parameter.  (optional)
	sequenceNumber := int32(56) // int32 | \\ If this query parameter is included, the query results will additionally include all log entries received in the `from` millisecond which have sequence numbers larger than the one specified.  Sequence numbers are identifiers used to distinguish between log entries received in the same millisecond. If a log entry was split up into several log entries during ingestion, then those chunks are ordered by sequence number.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueryLogDataAPI.PostQueryLogs(context.Background()).PostQueryLogsRequest(postQueryLogsRequest).PerPage(perPage).ExportFormat(exportFormat).KvpInfo(kvpInfo).MostRecentFirst(mostRecentFirst).Labels(labels).SequenceNumber(sequenceNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueryLogDataAPI.PostQueryLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostQueryLogs`: GetQueryLogs200Response
	fmt.Fprintf(os.Stdout, "Response from `QueryLogDataAPI.PostQueryLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostQueryLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postQueryLogsRequest** | [**PostQueryLogsRequest**](PostQueryLogsRequest.md) | The JSON payload must contain: * a non-empty array of the log keys for the logs that will be queried * a &#x60;leql&#x60; object, containing a value for either of the following (but **not** for both):   - &#x60;leql.during.time_range&#x60;   - &#x60;leql.during.from&#x60; and &#x60;leql.during.to&#x60;    The &#x60;leql.statement&#x60; can be empty or omitted.  | 
 **perPage** | **int32** | \\ Number of log entries to return per page, up to 500(the maximum allowed).  | [default to 50]
 **exportFormat** | **string** | \\ If included, the query results will be exported to the given format.  * Currently only &#x60;csv&#x60; is supported. * This parameter is only for non-statistical search queries (i.e. no &#39;calculate&#39; and/or &#39;groupby&#39; clauses). * Results are limited to the first 1 million log entries, and only one export job may run per account at a time.  The response will be a 202, and the response body will contain a link for polling the export job (on the **_/exports/{id}** endpoint).  | 
 **kvpInfo** | **bool** | \\ When set to true, the &#x60;events&#x60; object that is returned will additionally contain information about all the key-value pairs in each returned log entry.  | 
 **mostRecentFirst** | **bool** | \\ When set to &#x60;true&#x60;, the query returns the most recent events first. When set to &#x60;false&#x60;, it returns the oldest events first.  | [default to false]
 **labels** | **string** | \\ A set of &#39;:&#39; separated label UUIDs. Only entries which have a label matching one of these UUIDs will be returned. This only works with non-statistical queries (i.e. no &#39;groupby&#39; or &#39;calculate&#39; clauses). Takes precedence over the &#39;label&#39; parameter.  | 
 **sequenceNumber** | **int32** | \\ If this query parameter is included, the query results will additionally include all log entries received in the &#x60;from&#x60; millisecond which have sequence numbers larger than the one specified.  Sequence numbers are identifiers used to distinguish between log entries received in the same millisecond. If a log entry was split up into several log entries during ingestion, then those chunks are ordered by sequence number.  | 

### Return type

[**GetQueryLogs200Response**](GetQueryLogs200Response.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryLogsetsById

> GetQueryLogs200Response QueryLogsetsById(ctx, logsetId).From(from).To(to).Query(query).TimeRange(timeRange).Label(label).Labels(labels).PerPage(perPage).KvpInfo(kvpInfo).MostRecentFirst(mostRecentFirst).SequenceNumber(sequenceNumber).Execute()

Query Individual Log Sets



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
	from := int32(1450557004000) // int32 | \\ The start of the time range for the query, as a UNIX timestamp in milliseconds. 
	to := int32(1460557604000) // int32 | \\ The end of the time range for the query, as a UNIX timestamp in milliseconds. 
	query := "where(foo=bar)" // string | \\ A valid LEQL query to run against the log. If omitted, the query retrieves all log entries in the specified time range.  (optional)
	timeRange := "last 5 minutes" // string | \\ An alternative to the `from` and `to` query parameters. Supported values: * `yesterday` * `today` * `last x timeunits` where x is the number of time unit back from the current server time. Supported time units (case insensitive):     - min(s) or minute(s)     - hr(s) or hour(s)     - day(s)     - week(s)     - month(s)     - year(s)  If `time_range` is used, then the `from` and `to` query parameters must not be used.  (optional)
	label := "00000000-0000-0000-0000-000000000001" // string | \\ Only entries which have a label with this UUID will be returned. This only works with non-statistical queries (i.e. no 'groupby' or 'calculate' clauses).  (optional)
	labels := "00000000-0000-0000-0000-000000000001:21b21bb8-8869-4e2c-98df-684892e4e112" // string | \\ A set of ':' separated label UUIDs. Only entries which have a label matching one of these UUIDs will be returned. This only works with non-statistical queries (i.e. no 'groupby' or 'calculate' clauses). Takes precedence over the 'label' parameter.  (optional)
	perPage := int32(100) // int32 | \\ Number of log entries to return per page, up to 500(the maximum allowed).  (optional) (default to 50)
	kvpInfo := true // bool | \\ When set to true, the `events` object that is returned will additionally contain information about all the key-value pairs in each returned log entry.  (optional)
	mostRecentFirst := true // bool | \\ When set to `true`, the query returns the most recent events first. When set to `false`, it returns the oldest events first.  (optional) (default to false)
	sequenceNumber := int32(56) // int32 | \\ If this query parameter is included, the query results will additionally include all log entries received in the `from` millisecond which have sequence numbers larger than the one specified.  Sequence numbers are identifiers used to distinguish between log entries received in the same millisecond. If a log entry was split up into several log entries during ingestion, then those chunks are ordered by sequence number.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueryLogDataAPI.QueryLogsetsById(context.Background(), logsetId).From(from).To(to).Query(query).TimeRange(timeRange).Label(label).Labels(labels).PerPage(perPage).KvpInfo(kvpInfo).MostRecentFirst(mostRecentFirst).SequenceNumber(sequenceNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueryLogDataAPI.QueryLogsetsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryLogsetsById`: GetQueryLogs200Response
	fmt.Fprintf(os.Stdout, "Response from `QueryLogDataAPI.QueryLogsetsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logsetId** | **string** | \\ The id of the log set to be queried.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryLogsetsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **int32** | \\ The start of the time range for the query, as a UNIX timestamp in milliseconds.  | 
 **to** | **int32** | \\ The end of the time range for the query, as a UNIX timestamp in milliseconds.  | 
 **query** | **string** | \\ A valid LEQL query to run against the log. If omitted, the query retrieves all log entries in the specified time range.  | 
 **timeRange** | **string** | \\ An alternative to the &#x60;from&#x60; and &#x60;to&#x60; query parameters. Supported values: * &#x60;yesterday&#x60; * &#x60;today&#x60; * &#x60;last x timeunits&#x60; where x is the number of time unit back from the current server time. Supported time units (case insensitive):     - min(s) or minute(s)     - hr(s) or hour(s)     - day(s)     - week(s)     - month(s)     - year(s)  If &#x60;time_range&#x60; is used, then the &#x60;from&#x60; and &#x60;to&#x60; query parameters must not be used.  | 
 **label** | **string** | \\ Only entries which have a label with this UUID will be returned. This only works with non-statistical queries (i.e. no &#39;groupby&#39; or &#39;calculate&#39; clauses).  | 
 **labels** | **string** | \\ A set of &#39;:&#39; separated label UUIDs. Only entries which have a label matching one of these UUIDs will be returned. This only works with non-statistical queries (i.e. no &#39;groupby&#39; or &#39;calculate&#39; clauses). Takes precedence over the &#39;label&#39; parameter.  | 
 **perPage** | **int32** | \\ Number of log entries to return per page, up to 500(the maximum allowed).  | [default to 50]
 **kvpInfo** | **bool** | \\ When set to true, the &#x60;events&#x60; object that is returned will additionally contain information about all the key-value pairs in each returned log entry.  | 
 **mostRecentFirst** | **bool** | \\ When set to &#x60;true&#x60;, the query returns the most recent events first. When set to &#x60;false&#x60;, it returns the oldest events first.  | [default to false]
 **sequenceNumber** | **int32** | \\ If this query parameter is included, the query results will additionally include all log entries received in the &#x60;from&#x60; millisecond which have sequence numbers larger than the one specified.  Sequence numbers are identifiers used to distinguish between log entries received in the same millisecond. If a log entry was split up into several log entries during ingestion, then those chunks are ordered by sequence number.  | 

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


## QueryLogsetsByName

> GetQueryLogs200Response QueryLogsetsByName(ctx).LogsetName(logsetName).From(from).To(to).Query(query).TimeRange(timeRange).Label(label).Labels(labels).PerPage(perPage).KvpInfo(kvpInfo).MostRecentFirst(mostRecentFirst).SequenceNumber(sequenceNumber).Execute()

Query Multiple Log Sets



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
	logsetName := "Internal Logs" // string | \\ The name of the log set to be queried. Multiple log sets can be queried by providing this parameter multiple times, which results in the query running on the *union* of all contained logs. If any of the provided log set names do not exist, a 404 error response is returned. 
	from := int32(1450557004000) // int32 | \\ The start of the time range for the query, as a UNIX timestamp in milliseconds. 
	to := int32(1460557604000) // int32 | \\ The end of the time range for the query, as a UNIX timestamp in milliseconds. 
	query := "where(foo=bar)" // string | \\ A valid LEQL query to run against the log. If omitted, the query retrieves all log entries in the specified time range.  (optional)
	timeRange := "last 5 minutes" // string | \\ An alternative to the `from` and `to` query parameters. Supported values: * `yesterday` * `today` * `last x timeunits` where x is the number of time unit back from the current server time. Supported time units (case insensitive):     - min(s) or minute(s)     - hr(s) or hour(s)     - day(s)     - week(s)     - month(s)     - year(s)  If `time_range` is used, then the `from` and `to` query parameters must not be used.  (optional)
	label := "00000000-0000-0000-0000-000000000001" // string | \\ Only entries which have a label with this UUID will be returned. This only works with non-statistical queries (i.e. no 'groupby' or 'calculate' clauses).  (optional)
	labels := "00000000-0000-0000-0000-000000000001:21b21bb8-8869-4e2c-98df-684892e4e112" // string | \\ A set of ':' separated label UUIDs. Only entries which have a label matching one of these UUIDs will be returned. This only works with non-statistical queries (i.e. no 'groupby' or 'calculate' clauses). Takes precedence over the 'label' parameter.  (optional)
	perPage := int32(100) // int32 | \\ Number of log entries to return per page, up to 500(the maximum allowed).  (optional) (default to 50)
	kvpInfo := true // bool | \\ When set to true, the `events` object that is returned will additionally contain information about all the key-value pairs in each returned log entry.  (optional)
	mostRecentFirst := true // bool | \\ When set to `true`, the query returns the most recent events first. When set to `false`, it returns the oldest events first.  (optional) (default to false)
	sequenceNumber := int32(56) // int32 | \\ If this query parameter is included, the query results will additionally include all log entries received in the `from` millisecond which have sequence numbers larger than the one specified.  Sequence numbers are identifiers used to distinguish between log entries received in the same millisecond. If a log entry was split up into several log entries during ingestion, then those chunks are ordered by sequence number.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueryLogDataAPI.QueryLogsetsByName(context.Background()).LogsetName(logsetName).From(from).To(to).Query(query).TimeRange(timeRange).Label(label).Labels(labels).PerPage(perPage).KvpInfo(kvpInfo).MostRecentFirst(mostRecentFirst).SequenceNumber(sequenceNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueryLogDataAPI.QueryLogsetsByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryLogsetsByName`: GetQueryLogs200Response
	fmt.Fprintf(os.Stdout, "Response from `QueryLogDataAPI.QueryLogsetsByName`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryLogsetsByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **logsetName** | **string** | \\ The name of the log set to be queried. Multiple log sets can be queried by providing this parameter multiple times, which results in the query running on the *union* of all contained logs. If any of the provided log set names do not exist, a 404 error response is returned.  | 
 **from** | **int32** | \\ The start of the time range for the query, as a UNIX timestamp in milliseconds.  | 
 **to** | **int32** | \\ The end of the time range for the query, as a UNIX timestamp in milliseconds.  | 
 **query** | **string** | \\ A valid LEQL query to run against the log. If omitted, the query retrieves all log entries in the specified time range.  | 
 **timeRange** | **string** | \\ An alternative to the &#x60;from&#x60; and &#x60;to&#x60; query parameters. Supported values: * &#x60;yesterday&#x60; * &#x60;today&#x60; * &#x60;last x timeunits&#x60; where x is the number of time unit back from the current server time. Supported time units (case insensitive):     - min(s) or minute(s)     - hr(s) or hour(s)     - day(s)     - week(s)     - month(s)     - year(s)  If &#x60;time_range&#x60; is used, then the &#x60;from&#x60; and &#x60;to&#x60; query parameters must not be used.  | 
 **label** | **string** | \\ Only entries which have a label with this UUID will be returned. This only works with non-statistical queries (i.e. no &#39;groupby&#39; or &#39;calculate&#39; clauses).  | 
 **labels** | **string** | \\ A set of &#39;:&#39; separated label UUIDs. Only entries which have a label matching one of these UUIDs will be returned. This only works with non-statistical queries (i.e. no &#39;groupby&#39; or &#39;calculate&#39; clauses). Takes precedence over the &#39;label&#39; parameter.  | 
 **perPage** | **int32** | \\ Number of log entries to return per page, up to 500(the maximum allowed).  | [default to 50]
 **kvpInfo** | **bool** | \\ When set to true, the &#x60;events&#x60; object that is returned will additionally contain information about all the key-value pairs in each returned log entry.  | 
 **mostRecentFirst** | **bool** | \\ When set to &#x60;true&#x60;, the query returns the most recent events first. When set to &#x60;false&#x60;, it returns the oldest events first.  | [default to false]
 **sequenceNumber** | **int32** | \\ If this query parameter is included, the query results will additionally include all log entries received in the &#x60;from&#x60; millisecond which have sequence numbers larger than the one specified.  Sequence numbers are identifiers used to distinguish between log entries received in the same millisecond. If a log entry was split up into several log entries during ingestion, then those chunks are ordered by sequence number.  | 

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

