# invoker\AuditAPIAPI

All URIs are relative to *https://eu.rest.logs.insight.rapid7.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuditGetExportJob**](AuditAPIAPI.md#AuditGetExportJob) | **Get** /audit/exports/{id} | Retrieve An Export Job By Id
[**AuditGetExportJobs**](AuditAPIAPI.md#AuditGetExportJobs) | **Get** /audit/exports | Retrieve All Export Jobs
[**AuditGetQueryEndpoints**](AuditAPIAPI.md#AuditGetQueryEndpoints) | **Get** /audit/query | List All Endpoints
[**AuditGetQueryLogs**](AuditAPIAPI.md#AuditGetQueryLogs) | **Get** /audit/logs/{log_keys} | Query Individual Logs Using LEQL
[**AuditPollQuery**](AuditAPIAPI.md#AuditPollQuery) | **Get** /audit/query/{id} | Poll a Query In Progress
[**AuditPostQueryLogs**](AuditAPIAPI.md#AuditPostQueryLogs) | **Post** /audit/query/logs | Query Multiple Logs Using LEQL
[**GetAllAuditLogs**](AuditAPIAPI.md#GetAllAuditLogs) | **Get** /audit/management/logs | Retrieve All Audit Logs
[**GetAuditLog**](AuditAPIAPI.md#GetAuditLog) | **Get** /audit/management/logs/{id} | Retrieve An Audit Log



## AuditGetExportJob

> ExportJobResponse AuditGetExportJob(ctx).Id(id).Execute()

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
	resp, r, err := apiClient.AuditAPIAPI.AuditGetExportJob(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditAPIAPI.AuditGetExportJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditGetExportJob`: ExportJobResponse
	fmt.Fprintf(os.Stdout, "Response from `AuditAPIAPI.AuditGetExportJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuditGetExportJobRequest struct via the builder pattern


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


## AuditGetExportJobs

> ExportJobsResponse AuditGetExportJobs(ctx).Execute()

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
	resp, r, err := apiClient.AuditAPIAPI.AuditGetExportJobs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditAPIAPI.AuditGetExportJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditGetExportJobs`: ExportJobsResponse
	fmt.Fprintf(os.Stdout, "Response from `AuditAPIAPI.AuditGetExportJobs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuditGetExportJobsRequest struct via the builder pattern


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


## AuditGetQueryEndpoints

> map[string]ListEndpointsResponseValue AuditGetQueryEndpoints(ctx).Execute()

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
	resp, r, err := apiClient.AuditAPIAPI.AuditGetQueryEndpoints(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditAPIAPI.AuditGetQueryEndpoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditGetQueryEndpoints`: map[string]ListEndpointsResponseValue
	fmt.Fprintf(os.Stdout, "Response from `AuditAPIAPI.AuditGetQueryEndpoints`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuditGetQueryEndpointsRequest struct via the builder pattern


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


## AuditGetQueryLogs

> GetQueryLogs200Response AuditGetQueryLogs(ctx, logKeys).From(from).To(to).Query(query).TimeRange(timeRange).Label(label).Labels(labels).PerPage(perPage).ExportFormat(exportFormat).KvpInfo(kvpInfo).MostRecentFirst(mostRecentFirst).SequenceNumber(sequenceNumber).Execute()

Query Individual Logs Using LEQL



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
	resp, r, err := apiClient.AuditAPIAPI.AuditGetQueryLogs(context.Background(), logKeys).From(from).To(to).Query(query).TimeRange(timeRange).Label(label).Labels(labels).PerPage(perPage).ExportFormat(exportFormat).KvpInfo(kvpInfo).MostRecentFirst(mostRecentFirst).SequenceNumber(sequenceNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditAPIAPI.AuditGetQueryLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditGetQueryLogs`: GetQueryLogs200Response
	fmt.Fprintf(os.Stdout, "Response from `AuditAPIAPI.AuditGetQueryLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logKeys** | **string** | \\ The key of the log to be queried.  Supplying multiple &#39;:&#39; separated log keys is deprecated; users should instead &#39;POST&#39; the query as a JSON payload to [/query/logs](#operation/postQueryLogs).  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditGetQueryLogsRequest struct via the builder pattern


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


## AuditPollQuery

> PollQuery200Response AuditPollQuery(ctx, id).TimeRange(timeRange).Execute()

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
	resp, r, err := apiClient.AuditAPIAPI.AuditPollQuery(context.Background(), id).TimeRange(timeRange).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditAPIAPI.AuditPollQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditPollQuery`: PollQuery200Response
	fmt.Fprintf(os.Stdout, "Response from `AuditAPIAPI.AuditPollQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | \\ The continuation id generated when the query started.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditPollQueryRequest struct via the builder pattern


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


## AuditPostQueryLogs

> GetQueryLogs200Response AuditPostQueryLogs(ctx).PostQueryLogsRequest(postQueryLogsRequest).PerPage(perPage).ExportFormat(exportFormat).KvpInfo(kvpInfo).MostRecentFirst(mostRecentFirst).Labels(labels).SequenceNumber(sequenceNumber).Execute()

Query Multiple Logs Using LEQL



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
	postQueryLogsRequest := *openapiclient.NewPostQueryLogsRequest() // PostQueryLogsRequest | The JSON payload must contain: * a non-empty array of the log keys for the logs that will be queried * a `leql` object, containing a value for either of the following (but **not** for both):   - `leql.during.time_range`   - `leql.during.from` and `leql.during.to`  The `leql.statement` can be empty or omitted. 
	perPage := int32(100) // int32 | \\ Number of log entries to return per page, up to 500(the maximum allowed).  (optional) (default to 50)
	exportFormat := "csv" // string | \\ If included, the query results will be exported to the given format.  * Currently only `csv` is supported. * This parameter is only for non-statistical search queries (i.e. no 'calculate' and/or 'groupby' clauses). * Results are limited to the first 1 million log entries, and only one export job may run per account at a time.  The response will be a 202, and the response body will contain a link for polling the export job (on the **_/exports/{id}** endpoint).  (optional)
	kvpInfo := true // bool | \\ When set to true, the `events` object that is returned will additionally contain information about all the key-value pairs in each returned log entry.  (optional)
	mostRecentFirst := true // bool | \\ When set to `true`, the query returns the most recent events first. When set to `false`, it returns the oldest events first.  (optional) (default to false)
	labels := "00000000-0000-0000-0000-000000000001:21b21bb8-8869-4e2c-98df-684892e4e112" // string | \\ A set of ':' separated label UUIDs. Only entries which have a label matching one of these UUIDs will be returned. This only works with non-statistical queries (i.e. no 'groupby' or 'calculate' clauses). Takes precedence over the 'label' parameter.  (optional)
	sequenceNumber := int32(56) // int32 | \\ If this query parameter is included, the query results will additionally include all log entries received in the `from` millisecond which have sequence numbers larger than the one specified.  Sequence numbers are identifiers used to distinguish between log entries received in the same millisecond. If a log entry was split up into several log entries during ingestion, then those chunks are ordered by sequence number.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditAPIAPI.AuditPostQueryLogs(context.Background()).PostQueryLogsRequest(postQueryLogsRequest).PerPage(perPage).ExportFormat(exportFormat).KvpInfo(kvpInfo).MostRecentFirst(mostRecentFirst).Labels(labels).SequenceNumber(sequenceNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditAPIAPI.AuditPostQueryLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditPostQueryLogs`: GetQueryLogs200Response
	fmt.Fprintf(os.Stdout, "Response from `AuditAPIAPI.AuditPostQueryLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuditPostQueryLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postQueryLogsRequest** | [**PostQueryLogsRequest**](PostQueryLogsRequest.md) | The JSON payload must contain: * a non-empty array of the log keys for the logs that will be queried * a &#x60;leql&#x60; object, containing a value for either of the following (but **not** for both):   - &#x60;leql.during.time_range&#x60;   - &#x60;leql.during.from&#x60; and &#x60;leql.during.to&#x60;  The &#x60;leql.statement&#x60; can be empty or omitted.  | 
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


## GetAllAuditLogs

> AuditLogList GetAllAuditLogs(ctx).Execute()

Retrieve All Audit Logs



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
	resp, r, err := apiClient.AuditAPIAPI.GetAllAuditLogs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditAPIAPI.GetAllAuditLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllAuditLogs`: AuditLogList
	fmt.Fprintf(os.Stdout, "Response from `AuditAPIAPI.GetAllAuditLogs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllAuditLogsRequest struct via the builder pattern


### Return type

[**AuditLogList**](AuditLogList.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuditLog

> AuditLogSingle GetAuditLog(ctx, id).Execute()

Retrieve An Audit Log



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
	resp, r, err := apiClient.AuditAPIAPI.GetAuditLog(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditAPIAPI.GetAuditLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuditLog`: AuditLogSingle
	fmt.Fprintf(os.Stdout, "Response from `AuditAPIAPI.GetAuditLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | \\ The id of the log to be queried.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuditLogSingle**](AuditLogSingle.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

