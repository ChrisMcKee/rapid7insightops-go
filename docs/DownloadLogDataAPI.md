# invoker\DownloadLogDataAPI

All URIs are relative to *https://eu.rest.logs.insight.rapid7.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDownloadLogs**](DownloadLogDataAPI.md#GetDownloadLogs) | **Get** /download/logs | Download Log Data



## GetDownloadLogs

> GetDownloadLogs(ctx, ids).From(from).To(to).TimeRange(timeRange).Query(query).Limit(limit).Execute()

Download Log Data



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
	ids := "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa:bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb:cccccccc-cccc-cccc-cccc-cccccccccccc" // string | \\ The UUIDs of the logs separated by colons \":\" or semicolons \";\". For example \"UUID1:UUID2:UUID3\" or \"UUID1;UUID2;UUID3\". 
	from := int32(1450557004000) // int32 | \\ The start of the time range for the query, as a UNIX timestamp in milliseconds. 
	to := int32(1460557604000) // int32 | \\ The end of the time range for the query, as a UNIX timestamp in milliseconds. 
	timeRange := "last 5 minutes" // string | \\ An alternative to the `from` and `to` query parameters. Supported values: * `yesterday` * `today` * `last x timeunits` where x is the number of time unit back from the current server time. Supported time units (case insensitive):     - min(s) or minute(s)     - hr(s) or hour(s)     - day(s)     - week(s)     - month(s)     - year(s)  If `time_range` is used, then the `from` and `to` query parameters must not be used.  (optional)
	query := "where(foo=bar)" // string | \\ A valid LEQL query to run against the log. If omitted, the query retrieves all log entries in the specified time range.  (optional)
	limit := int32(10) // int32 | \\ The maximum, and default, number of log entries to download is 500,000,000 log entries. You can specify this parameter to limit the number of log entries in the download.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DownloadLogDataAPI.GetDownloadLogs(context.Background(), ids).From(from).To(to).TimeRange(timeRange).Query(query).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DownloadLogDataAPI.GetDownloadLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ids** | **string** | \\ The UUIDs of the logs separated by colons \&quot;:\&quot; or semicolons \&quot;;\&quot;. For example \&quot;UUID1:UUID2:UUID3\&quot; or \&quot;UUID1;UUID2;UUID3\&quot;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDownloadLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **int32** | \\ The start of the time range for the query, as a UNIX timestamp in milliseconds.  | 
 **to** | **int32** | \\ The end of the time range for the query, as a UNIX timestamp in milliseconds.  | 
 **timeRange** | **string** | \\ An alternative to the &#x60;from&#x60; and &#x60;to&#x60; query parameters. Supported values: * &#x60;yesterday&#x60; * &#x60;today&#x60; * &#x60;last x timeunits&#x60; where x is the number of time unit back from the current server time. Supported time units (case insensitive):     - min(s) or minute(s)     - hr(s) or hour(s)     - day(s)     - week(s)     - month(s)     - year(s)  If &#x60;time_range&#x60; is used, then the &#x60;from&#x60; and &#x60;to&#x60; query parameters must not be used.  | 
 **query** | **string** | \\ A valid LEQL query to run against the log. If omitted, the query retrieves all log entries in the specified time range.  | 
 **limit** | **int32** | \\ The maximum, and default, number of log entries to download is 500,000,000 log entries. You can specify this parameter to limit the number of log entries in the download.  | 

### Return type

 (empty response body)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

