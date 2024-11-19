/*
InsightOps REST API

### Overview  Our REST API lets you use InsightOps through HTTP requests. Currently, the REST API allows you to perform the majority of the actions available through the UI, and has some additional functionality that is not available through the UI. You may use this API to automate common tasks (for example, via shell scripts), and to generally interact with InsightOps programmatically.  This page precisely describes the REST API and serves as a reference for the API. Each HTTP method and each URL endpoint is documented in a self-contained unit so that users only need to read about the HTTP methods relevant to them.  ### Terminology  * A **log entry** is an individual log event. * A **log** is a collection of log entries, or a single log stream. * A **log set** is a logical-only collection of logs, i.e. logs can be in multiple logsets and deleting a logset only deletes the relation between logs, not the logs themselves. * [Log Entry Query Language](https://docs.rapid7.com/insightops/log-search) (**LEQL**) is the query language used in Insight Ops to search log data.

API version: latest
Contact: support@rapid7.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package insightops

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// WatchIncomingLogDataInRealtimeAPIService WatchIncomingLogDataInRealtimeAPI service
type WatchIncomingLogDataInRealtimeAPIService service

type ApiGetLiveLogsRequest struct {
	ctx        context.Context
	ApiService *WatchIncomingLogDataInRealtimeAPIService
	logKeys    string
	query      *string
	mode       *string
	label      *string
	labels     *string
	perPage    *int32
}

// \\ A valid LEQL query for the Live Tail feed. If omitted, the query matches all log entries.  *Live Tail does not support &#39;calculate&#39; or &#39;groupby&#39; queries.*
func (r ApiGetLiveLogsRequest) Query(query string) ApiGetLiveLogsRequest {
	r.query = &query
	return r
}

// \\ The order the log entries are returned in. Possible values: * &#x60;tail&#x60; (which returns entries in the order: newest to oldest.) * &#x60;head&#x60; (which returns entries in the order: oldest to newest.)
func (r ApiGetLiveLogsRequest) Mode(mode string) ApiGetLiveLogsRequest {
	r.mode = &mode
	return r
}

// \\ Only entries which have a label with this UUID will be returned. This only works with non-statistical queries (i.e. no &#39;groupby&#39; or &#39;calculate&#39; clauses).
func (r ApiGetLiveLogsRequest) Label(label string) ApiGetLiveLogsRequest {
	r.label = &label
	return r
}

// \\ A set of &#39;:&#39; separated label UUIDs. Only entries which have a label matching one of these UUIDs will be returned. This only works with non-statistical queries (i.e. no &#39;groupby&#39; or &#39;calculate&#39; clauses). Takes precedence over the &#39;label&#39; parameter.
func (r ApiGetLiveLogsRequest) Labels(labels string) ApiGetLiveLogsRequest {
	r.labels = &labels
	return r
}

// \\ The maximum number of log entries to return at a time, up to 500(the maximum allowed). If the number of results exceed this value, the next \&quot;page\&quot; of results is returned the next time that the Live Tail feed is polled.
func (r ApiGetLiveLogsRequest) PerPage(perPage int32) ApiGetLiveLogsRequest {
	r.perPage = &perPage
	return r
}

func (r ApiGetLiveLogsRequest) Execute() (*LiveTailResponse, *http.Response, error) {
	return r.ApiService.GetLiveLogsExecute(r)
}

/*
GetLiveLogs Start A Live Tail Feed (Individual Logs)

Starts a Live Tail feed on the logs with those **{log_keys}**, for the Query specified in the URL parameters.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param logKeys \\ The keys of the logs for the Live Tail Feed.  Supplying multiple ':' separated log keys is deprecated; users should instead 'POST' the query as a JSON payload to [/query/live/logs](#operation/postLiveLogs).
	@return ApiGetLiveLogsRequest
*/
func (a *WatchIncomingLogDataInRealtimeAPIService) GetLiveLogs(ctx context.Context, logKeys string) ApiGetLiveLogsRequest {
	return ApiGetLiveLogsRequest{
		ApiService: a,
		ctx:        ctx,
		logKeys:    logKeys,
	}
}

// Execute executes the request
//
//	@return LiveTailResponse
func (a *WatchIncomingLogDataInRealtimeAPIService) GetLiveLogsExecute(r ApiGetLiveLogsRequest) (*LiveTailResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *LiveTailResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WatchIncomingLogDataInRealtimeAPIService.GetLiveLogs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/query/live/logs/{log_keys}"
	localVarPath = strings.Replace(localVarPath, "{"+"log_keys"+"}", url.PathEscape(parameterValueTostring(r.logKeys, "logKeys")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.query != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "query", r.query, "", "")
	}
	if r.mode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mode", r.mode, "", "")
	} else {
		var defaultValue string = "tail"
		r.mode = &defaultValue
	}
	if r.label != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "label", r.label, "", "")
	}
	if r.labels != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "labels", r.labels, "", "")
	}
	if r.perPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "per_page", r.perPage, "", "")
	} else {
		var defaultValue int32 = 50
		r.perPage = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Api Key Authentication"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiLiveSavedQueryRequest struct {
	ctx          context.Context
	ApiService   *WatchIncomingLogDataInRealtimeAPIService
	savedQueryId string
	perPage      *int32
	mode         *string
}

// \\ The maximum number of log entries to return at a time, up to 500(the maximum allowed). If the number of results exceed this value, the next \&quot;page\&quot; of results is returned the next time that the Live Tail feed is polled.
func (r ApiLiveSavedQueryRequest) PerPage(perPage int32) ApiLiveSavedQueryRequest {
	r.perPage = &perPage
	return r
}

// \\ The order the log entries are returned in. Possible values: * &#x60;tail&#x60; (which returns entries in the order: newest to oldest.) * &#x60;head&#x60; (which returns entries in the order: oldest to newest.)
func (r ApiLiveSavedQueryRequest) Mode(mode string) ApiLiveSavedQueryRequest {
	r.mode = &mode
	return r
}

func (r ApiLiveSavedQueryRequest) Execute() (*LiveTailResponse, *http.Response, error) {
	return r.ApiService.LiveSavedQueryExecute(r)
}

/*
LiveSavedQuery Start a Live Tail Feed (Saved Query - with logs)

Starts a Live Tail Feed using the saved query which has id **{saved_query_id}**.

This endpoint can only be used for saved queries which *have* the logs for the live tail feed already specified.
For saved queries which *don't have* the logs already specified, use the
[/query/live/logs/{log_keys}/{saved_query_id}](#operation/liveSavedQueryNoLogs) endpoint.

Saved queries which have a *time range* already specified can still be used (the time range is ignored).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param savedQueryId \\ The id of the saved query.
	@return ApiLiveSavedQueryRequest
*/
func (a *WatchIncomingLogDataInRealtimeAPIService) LiveSavedQuery(ctx context.Context, savedQueryId string) ApiLiveSavedQueryRequest {
	return ApiLiveSavedQueryRequest{
		ApiService:   a,
		ctx:          ctx,
		savedQueryId: savedQueryId,
	}
}

// Execute executes the request
//
//	@return LiveTailResponse
func (a *WatchIncomingLogDataInRealtimeAPIService) LiveSavedQueryExecute(r ApiLiveSavedQueryRequest) (*LiveTailResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *LiveTailResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WatchIncomingLogDataInRealtimeAPIService.LiveSavedQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/query/live/saved_query/{saved_query_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"saved_query_id"+"}", url.PathEscape(parameterValueTostring(r.savedQueryId, "savedQueryId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.perPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "per_page", r.perPage, "", "")
	} else {
		var defaultValue int32 = 50
		r.perPage = &defaultValue
	}
	if r.mode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mode", r.mode, "", "")
	} else {
		var defaultValue string = "tail"
		r.mode = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Api Key Authentication"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiLiveSavedQueryNoLogsRequest struct {
	ctx          context.Context
	ApiService   *WatchIncomingLogDataInRealtimeAPIService
	logKeys      string
	savedQueryId string
	perPage      *int32
	mode         *string
}

// \\ The maximum number of log entries to return at a time, up to 500(the maximum allowed). If the number of results exceed this value, the next \&quot;page\&quot; of results is returned the next time that the Live Tail feed is polled.
func (r ApiLiveSavedQueryNoLogsRequest) PerPage(perPage int32) ApiLiveSavedQueryNoLogsRequest {
	r.perPage = &perPage
	return r
}

// \\ The order the log entries are returned in. Possible values: * &#x60;tail&#x60; (which returns entries in the order: newest to oldest.) * &#x60;head&#x60; (which returns entries in the order: oldest to newest.)
func (r ApiLiveSavedQueryNoLogsRequest) Mode(mode string) ApiLiveSavedQueryNoLogsRequest {
	r.mode = &mode
	return r
}

func (r ApiLiveSavedQueryNoLogsRequest) Execute() (*LiveTailResponse, *http.Response, error) {
	return r.ApiService.LiveSavedQueryNoLogsExecute(r)
}

/*
LiveSavedQueryNoLogs Start a Live Tail Feed (Saved Query - without logs)

Starts a Live Tail Feed using the saved query which has the id **{saved_query_id}**
on the logs matching the **{log_keys}**.

This endpoint can only be used for saved queries which *don't have* the logs for the query already specified.
For saved queries which *have* the logs already specified, use the
[/query/live/saved_query/{saved_query_id}](#operation/liveSavedQuery) endpoint.

Saved queries which have a *time range* already specified can still be used (the time range is ignored).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param logKeys \\ The keys of the logs for the Live Tail Feed, separated by the ':' character.
	@param savedQueryId \\ The id of the saved query.
	@return ApiLiveSavedQueryNoLogsRequest
*/
func (a *WatchIncomingLogDataInRealtimeAPIService) LiveSavedQueryNoLogs(ctx context.Context, logKeys string, savedQueryId string) ApiLiveSavedQueryNoLogsRequest {
	return ApiLiveSavedQueryNoLogsRequest{
		ApiService:   a,
		ctx:          ctx,
		logKeys:      logKeys,
		savedQueryId: savedQueryId,
	}
}

// Execute executes the request
//
//	@return LiveTailResponse
func (a *WatchIncomingLogDataInRealtimeAPIService) LiveSavedQueryNoLogsExecute(r ApiLiveSavedQueryNoLogsRequest) (*LiveTailResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *LiveTailResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WatchIncomingLogDataInRealtimeAPIService.LiveSavedQueryNoLogs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/query/live/logs/{log_keys}/{saved_query_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"log_keys"+"}", url.PathEscape(parameterValueTostring(r.logKeys, "logKeys")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"saved_query_id"+"}", url.PathEscape(parameterValueTostring(r.savedQueryId, "savedQueryId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.perPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "per_page", r.perPage, "", "")
	} else {
		var defaultValue int32 = 50
		r.perPage = &defaultValue
	}
	if r.mode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mode", r.mode, "", "")
	} else {
		var defaultValue string = "tail"
		r.mode = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Api Key Authentication"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPollLiveTailRequest struct {
	ctx        context.Context
	ApiService *WatchIncomingLogDataInRealtimeAPIService
	id         string
	query      *string
	logKeys    *string
}

// \\ A valid LEQL query for the Live Tail feed. If omitted, the query matches all log entries.  *Live Tail does not support &#39;calculate&#39; or &#39;groupby&#39; queries.*
func (r ApiPollLiveTailRequest) Query(query string) ApiPollLiveTailRequest {
	r.query = &query
	return r
}

// \\ The keys of the logs for the Live Tail Feed, separated by the &#39;:&#39; character.
func (r ApiPollLiveTailRequest) LogKeys(logKeys string) ApiPollLiveTailRequest {
	r.logKeys = &logKeys
	return r
}

func (r ApiPollLiveTailRequest) Execute() (*LiveTailPollResponse, *http.Response, error) {
	return r.ApiService.PollLiveTailExecute(r)
}

/*
PollLiveTail Poll a Live Tail Feed

Polls the Live Tail Feed with the corresponding **{id}**, returning all the log entries for the Feed since the
since last time the Feed was polled. If the number of log entries exceeds 50 (by default),
only 50 log entries will be returned, and the next "page" is returned on the next poll.

The `log_keys` and `query` metadata information for the Live Tail Feed can only be tracked by the client,
by providing it as a query parameter with every polling request, which will be returned as given in the response
(does not affect the query execution - only metadata information).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id \\ The polling id generated when the Live Tail Feed started.
	@return ApiPollLiveTailRequest
*/
func (a *WatchIncomingLogDataInRealtimeAPIService) PollLiveTail(ctx context.Context, id string) ApiPollLiveTailRequest {
	return ApiPollLiveTailRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return LiveTailPollResponse
func (a *WatchIncomingLogDataInRealtimeAPIService) PollLiveTailExecute(r ApiPollLiveTailRequest) (*LiveTailPollResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *LiveTailPollResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WatchIncomingLogDataInRealtimeAPIService.PollLiveTail")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/query/live/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueTostring(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.query != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "query", r.query, "", "")
	}
	if r.logKeys != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "log_keys", r.logKeys, "", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Api Key Authentication"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPostLiveLogsRequest struct {
	ctx                 context.Context
	ApiService          *WatchIncomingLogDataInRealtimeAPIService
	postLiveLogsRequest *PostLiveLogsRequest
	labels              *string
}

func (r ApiPostLiveLogsRequest) PostLiveLogsRequest(postLiveLogsRequest PostLiveLogsRequest) ApiPostLiveLogsRequest {
	r.postLiveLogsRequest = &postLiveLogsRequest
	return r
}

// \\ A set of &#39;:&#39; separated label UUIDs. Only entries which have a label matching one of these UUIDs will be returned. This only works with non-statistical queries (i.e. no &#39;groupby&#39; or &#39;calculate&#39; clauses). Takes precedence over the &#39;label&#39; parameter.
func (r ApiPostLiveLogsRequest) Labels(labels string) ApiPostLiveLogsRequest {
	r.labels = &labels
	return r
}

func (r ApiPostLiveLogsRequest) Execute() (*LiveTailResponse, *http.Response, error) {
	return r.ApiService.PostLiveLogsExecute(r)
}

/*
PostLiveLogs Start A Live Tail Feed (Multiple Logs)

Starts a Live Tail feed, where the logs and the query for the feed are specified in the (JSON) request body.

An arbitrary collection of logs can be specified in the request body.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPostLiveLogsRequest
*/
func (a *WatchIncomingLogDataInRealtimeAPIService) PostLiveLogs(ctx context.Context) ApiPostLiveLogsRequest {
	return ApiPostLiveLogsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return LiveTailResponse
func (a *WatchIncomingLogDataInRealtimeAPIService) PostLiveLogsExecute(r ApiPostLiveLogsRequest) (*LiveTailResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *LiveTailResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WatchIncomingLogDataInRealtimeAPIService.PostLiveLogs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/query/live/logs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.postLiveLogsRequest == nil {
		return localVarReturnValue, nil, reportError("postLiveLogsRequest is required and must be specified")
	}

	if r.labels != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "labels", r.labels, "", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.postLiveLogsRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Api Key Authentication"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}