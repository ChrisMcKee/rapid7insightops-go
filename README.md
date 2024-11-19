# Go API client for insightops

### Overview

Our REST API lets you use InsightOps through HTTP requests.
Currently, the REST API allows you to perform the majority of the actions available through the UI, and has some
additional functionality that is not available through the UI.
You may use this API to automate common tasks (for example, via shell scripts), and to generally interact with InsightOps
programmatically.

This page precisely describes the REST API and serves as a reference for the API.
Each HTTP method and each URL endpoint is documented in a self-contained unit
so that users only need to read about the HTTP methods relevant to them.

### Terminology

* A **log entry** is an individual log event.
* A **log** is a collection of log entries, or a single log stream.
* A **log set** is a logical-only collection of logs, i.e. logs can be in multiple logsets and deleting a logset only deletes the relation between logs, not the logs themselves.
* [Log Entry Query Language](https://docs.rapid7.com/insightops/log-search) (**LEQL**) is the query language used in
Insight Ops to search log data.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: latest
- Package version: 1.0.0
- Generator version: 7.10.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import insightops "github.com/chrismckee/rapid7insightops-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `insightops.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), insightops.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `insightops.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), insightops.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `insightops.ContextOperationServerIndices` and `insightops.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), insightops.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), insightops.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://eu.rest.logs.insight.rapid7.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AuditAPIAPI* | [**AuditGetExportJob**](docs/AuditAPIAPI.md#auditgetexportjob) | **Get** /audit/exports/{id} | Retrieve An Export Job By Id
*AuditAPIAPI* | [**AuditGetExportJobs**](docs/AuditAPIAPI.md#auditgetexportjobs) | **Get** /audit/exports | Retrieve All Export Jobs
*AuditAPIAPI* | [**AuditGetQueryEndpoints**](docs/AuditAPIAPI.md#auditgetqueryendpoints) | **Get** /audit/query | List All Endpoints
*AuditAPIAPI* | [**AuditGetQueryLogs**](docs/AuditAPIAPI.md#auditgetquerylogs) | **Get** /audit/logs/{log_keys} | Query Individual Logs Using LEQL
*AuditAPIAPI* | [**AuditPollQuery**](docs/AuditAPIAPI.md#auditpollquery) | **Get** /audit/query/{id} | Poll a Query In Progress
*AuditAPIAPI* | [**AuditPostQueryLogs**](docs/AuditAPIAPI.md#auditpostquerylogs) | **Post** /audit/query/logs | Query Multiple Logs Using LEQL
*AuditAPIAPI* | [**GetAllAuditLogs**](docs/AuditAPIAPI.md#getallauditlogs) | **Get** /audit/management/logs | Retrieve All Audit Logs
*AuditAPIAPI* | [**GetAuditLog**](docs/AuditAPIAPI.md#getauditlog) | **Get** /audit/management/logs/{id} | Retrieve An Audit Log
*BackupYourLogDataToS3API* | [**DeleteArchivingS3Setup**](docs/BackupYourLogDataToS3API.md#deletearchivings3setup) | **Delete** /management/archiving/s3setup | Disable Daily Archiving
*BackupYourLogDataToS3API* | [**GetArchivingS3Setup**](docs/BackupYourLogDataToS3API.md#getarchivings3setup) | **Get** /management/archiving/s3setup | View Daily Archiving Settings
*BackupYourLogDataToS3API* | [**PatchArchivingS3Setup**](docs/BackupYourLogDataToS3API.md#patcharchivings3setup) | **Patch** /management/archiving/s3setup | Modify Daily Archiving Settings
*BackupYourLogDataToS3API* | [**PostArchivingS3Setup**](docs/BackupYourLogDataToS3API.md#postarchivings3setup) | **Post** /management/archiving/s3setup | Enable Daily Archiving
*BackupYourLogDataToS3API* | [**PutArchivingS3Setup**](docs/BackupYourLogDataToS3API.md#putarchivings3setup) | **Put** /management/archiving/s3setup | Replace Daily Archiving Settings
*BasicDetectionRulesAPI* | [**DeleteLabelById**](docs/BasicDetectionRulesAPI.md#deletelabelbyid) | **Delete** /management/labels/{label_id} | Delete a Label
*BasicDetectionRulesAPI* | [**DeleteManagementTagActionById**](docs/BasicDetectionRulesAPI.md#deletemanagementtagactionbyid) | **Delete** /management/actions/{action_id} | Delete a Notification
*BasicDetectionRulesAPI* | [**DeleteManagementTagById**](docs/BasicDetectionRulesAPI.md#deletemanagementtagbyid) | **Delete** /management/tags/{rule_id} | Delete a Basic Detection Rule
*BasicDetectionRulesAPI* | [**DeleteManagementTargetById**](docs/BasicDetectionRulesAPI.md#deletemanagementtargetbyid) | **Delete** /management/targets/{target_id} | Delete a Target
*BasicDetectionRulesAPI* | [**GetLabelById**](docs/BasicDetectionRulesAPI.md#getlabelbyid) | **Get** /management/labels/{label_id} | Retrieve a Label
*BasicDetectionRulesAPI* | [**GetLabels**](docs/BasicDetectionRulesAPI.md#getlabels) | **Get** /management/labels | List all Labels
*BasicDetectionRulesAPI* | [**GetManagementAlertNotificationSettings**](docs/BasicDetectionRulesAPI.md#getmanagementalertnotificationsettings) | **Get** /management/actions | List All Notifications
*BasicDetectionRulesAPI* | [**GetManagementTagActionById**](docs/BasicDetectionRulesAPI.md#getmanagementtagactionbyid) | **Get** /management/actions/{action_id} | Retrieve a Notification
*BasicDetectionRulesAPI* | [**GetManagementTagActionTargets**](docs/BasicDetectionRulesAPI.md#getmanagementtagactiontargets) | **Get** /management/actions/{action_id}/targets | List All Targets Attached To A Notification
*BasicDetectionRulesAPI* | [**GetManagementTagById**](docs/BasicDetectionRulesAPI.md#getmanagementtagbyid) | **Get** /management/tags/{rule_id} | Retrieve a Basic Detection Rule
*BasicDetectionRulesAPI* | [**GetManagementTags**](docs/BasicDetectionRulesAPI.md#getmanagementtags) | **Get** /management/tags | List All Tags
*BasicDetectionRulesAPI* | [**GetManagementTargetById**](docs/BasicDetectionRulesAPI.md#getmanagementtargetbyid) | **Get** /management/targets/{target_id} | Retrieve a Target
*BasicDetectionRulesAPI* | [**GetManagementTargets**](docs/BasicDetectionRulesAPI.md#getmanagementtargets) | **Get** /management/targets | List All Targets
*BasicDetectionRulesAPI* | [**PatchLabelById**](docs/BasicDetectionRulesAPI.md#patchlabelbyid) | **Patch** /management/labels/{label_id} | Update a Label
*BasicDetectionRulesAPI* | [**PatchManagementTagActionById**](docs/BasicDetectionRulesAPI.md#patchmanagementtagactionbyid) | **Patch** /management/actions/{action_id} | Modify a Notification
*BasicDetectionRulesAPI* | [**PatchManagementTagActionTargetById**](docs/BasicDetectionRulesAPI.md#patchmanagementtagactiontargetbyid) | **Patch** /management/actions/{action_id}/targets | Modify the Targets Attached To A Notification
*BasicDetectionRulesAPI* | [**PatchManagementTagById**](docs/BasicDetectionRulesAPI.md#patchmanagementtagbyid) | **Patch** /management/tags/{rule_id} | Modify a Basic Detection Rule
*BasicDetectionRulesAPI* | [**PostLabels**](docs/BasicDetectionRulesAPI.md#postlabels) | **Post** /management/labels | Create a Label
*BasicDetectionRulesAPI* | [**PostManagementTagActions**](docs/BasicDetectionRulesAPI.md#postmanagementtagactions) | **Post** /management/actions | Create a Notification
*BasicDetectionRulesAPI* | [**PostManagementTags**](docs/BasicDetectionRulesAPI.md#postmanagementtags) | **Post** /management/tags | Create A Basic Detection Rule
*BasicDetectionRulesAPI* | [**PostManagementTargets**](docs/BasicDetectionRulesAPI.md#postmanagementtargets) | **Post** /management/targets | Create a Target
*BasicDetectionRulesAPI* | [**PutLabelById**](docs/BasicDetectionRulesAPI.md#putlabelbyid) | **Put** /management/labels/{label_id} | Replace a Label
*BasicDetectionRulesAPI* | [**PutManagementTagActionById**](docs/BasicDetectionRulesAPI.md#putmanagementtagactionbyid) | **Put** /management/actions/{action_id} | Replace a Notification
*BasicDetectionRulesAPI* | [**PutManagementTagById**](docs/BasicDetectionRulesAPI.md#putmanagementtagbyid) | **Put** /management/tags/{rule_id} | Replace a Basic Detection Rule
*BasicDetectionRulesAPI* | [**PutManagementTargetById**](docs/BasicDetectionRulesAPI.md#putmanagementtargetbyid) | **Put** /management/targets/{target_id} | Replace a Target
*DownloadLogDataAPI* | [**GetDownloadLogs**](docs/DownloadLogDataAPI.md#getdownloadlogs) | **Get** /download/logs | Download Log Data
*ExploreTheSizeOfYourLogDataAPI* | [**GetOrganizationsLogUsage**](docs/ExploreTheSizeOfYourLogDataAPI.md#getorganizationslogusage) | **Get** /usage/organizations/logs/{log_key} | Data Size For a Specific Log
*ExploreTheSizeOfYourLogDataAPI* | [**GetOrganizationsLogsUsage**](docs/ExploreTheSizeOfYourLogDataAPI.md#getorganizationslogsusage) | **Get** /usage/organizations/logs | Data Size Broken Down By Log
*ExploreTheSizeOfYourLogDataAPI* | [**GetOrganizationsUsage**](docs/ExploreTheSizeOfYourLogDataAPI.md#getorganizationsusage) | **Get** /usage/organizations | Total Data Size Across All Logs
*ExportLogDataToCSVAPI* | [**DeleteExportJob**](docs/ExportLogDataToCSVAPI.md#deleteexportjob) | **Delete** /exports/{id} | Delete An Export Job By Id
*ExportLogDataToCSVAPI* | [**GetExportJob**](docs/ExportLogDataToCSVAPI.md#getexportjob) | **Get** /exports/{id} | Retrieve An Export Job By Id
*ExportLogDataToCSVAPI* | [**GetExportJobs**](docs/ExportLogDataToCSVAPI.md#getexportjobs) | **Get** /exports | Retrieve All Export Jobs
*LEQLVariablesAPI* | [**CreateVariable**](docs/LEQLVariablesAPI.md#createvariable) | **Post** /query/variables | Create a LEQL variable
*LEQLVariablesAPI* | [**DeleteVariable**](docs/LEQLVariablesAPI.md#deletevariable) | **Delete** /query/variables/{id} | Delete a LEQL variable
*LEQLVariablesAPI* | [**ListVariables**](docs/LEQLVariablesAPI.md#listvariables) | **Get** /query/variables | List all LEQL variables
*LEQLVariablesAPI* | [**RetrieveVariable**](docs/LEQLVariablesAPI.md#retrievevariable) | **Get** /query/variables/{id} | Retrieve a LEQL variable
*LEQLVariablesAPI* | [**UpdateVariable**](docs/LEQLVariablesAPI.md#updatevariable) | **Put** /query/variables/{id} | Update a LEQL variable
*LogsLogSetsAPI* | [**DeleteLog**](docs/LogsLogSetsAPI.md#deletelog) | **Delete** /management/logs/{id} | Delete a Log
*LogsLogSetsAPI* | [**DeleteManagementLogsetById**](docs/LogsLogSetsAPI.md#deletemanagementlogsetbyid) | **Delete** /management/logsets/{logset_id} | Delete a Log Set
*LogsLogSetsAPI* | [**GetLog**](docs/LogsLogSetsAPI.md#getlog) | **Get** /management/logs/{id} | Retrieve a Log
*LogsLogSetsAPI* | [**GetLogs**](docs/LogsLogSetsAPI.md#getlogs) | **Get** /management/logs | Retrieve All Logs
*LogsLogSetsAPI* | [**GetManagementLogsetById**](docs/LogsLogSetsAPI.md#getmanagementlogsetbyid) | **Get** /management/logsets/{logset_id} | Retrieve a Log Set
*LogsLogSetsAPI* | [**GetManagementLogsets**](docs/LogsLogSetsAPI.md#getmanagementlogsets) | **Get** /management/logsets | List All Log Sets
*LogsLogSetsAPI* | [**PostLog**](docs/LogsLogSetsAPI.md#postlog) | **Post** /management/logs | Create a Log
*LogsLogSetsAPI* | [**PostManagementLogsets**](docs/LogsLogSetsAPI.md#postmanagementlogsets) | **Post** /management/logsets | Create a Log Set
*LogsLogSetsAPI* | [**PutLog**](docs/LogsLogSetsAPI.md#putlog) | **Put** /management/logs/{id} | Replace a Log
*LogsLogSetsAPI* | [**PutManagementLogsetById**](docs/LogsLogSetsAPI.md#putmanagementlogsetbyid) | **Put** /management/logsets/{logset_id} | Replace a Log Set
*PreComputedQueriesAPI* | [**MetricsCreate**](docs/PreComputedQueriesAPI.md#metricscreate) | **Post** /management/metrics | Create A Pre-Computed Query
*PreComputedQueriesAPI* | [**MetricsDelete**](docs/PreComputedQueriesAPI.md#metricsdelete) | **Delete** /management/metrics/{metric_id} | Delete A Pre-Computed Query
*PreComputedQueriesAPI* | [**MetricsPut**](docs/PreComputedQueriesAPI.md#metricsput) | **Put** /management/metrics/{metric_id} | Replace a Pre-Computed Query
*PreComputedQueriesAPI* | [**MetricsRetrieve**](docs/PreComputedQueriesAPI.md#metricsretrieve) | **Get** /management/metrics/{metric_id} | Retrieve A Pre-Computed Query
*PreComputedQueriesAPI* | [**MetricsRetrieveAll**](docs/PreComputedQueriesAPI.md#metricsretrieveall) | **Get** /management/metrics | List All Pre-Computed Queries
*PreComputedQueriesAPI* | [**QueryMetrics**](docs/PreComputedQueriesAPI.md#querymetrics) | **Get** /query/metrics/{id} | Fetch Query Results
*QueryLogDataAPI* | [**GetContextEvents**](docs/QueryLogDataAPI.md#getcontextevents) | **Get** /query/context/{id} | Retrieve The Log Data Immediately Before And After A Specific Log Line
*QueryLogDataAPI* | [**GetQueryEndpoints**](docs/QueryLogDataAPI.md#getqueryendpoints) | **Get** /query | List All Endpoints
*QueryLogDataAPI* | [**GetQueryLogs**](docs/QueryLogDataAPI.md#getquerylogs) | **Get** /query/logs/{log_keys} | Query Individual Logs
*QueryLogDataAPI* | [**GetSearchStats**](docs/QueryLogDataAPI.md#getsearchstats) | **Get** /search-stats | View Statistics On Past Queries
*QueryLogDataAPI* | [**PollQuery**](docs/QueryLogDataAPI.md#pollquery) | **Get** /query/{id} | Poll a Query In Progress
*QueryLogDataAPI* | [**PostQueryLogs**](docs/QueryLogDataAPI.md#postquerylogs) | **Post** /query/logs | Query Multiple Logs
*QueryLogDataAPI* | [**QueryLogsetsById**](docs/QueryLogDataAPI.md#querylogsetsbyid) | **Get** /query/logsets/{id} | Query Individual Log Sets
*QueryLogDataAPI* | [**QueryLogsetsByName**](docs/QueryLogDataAPI.md#querylogsetsbyname) | **Get** /query/logsets | Query Multiple Log Sets
*RetrieveTheMostCommonKeysWithinYourLogDataAPI* | [**GetLogTopkeys**](docs/RetrieveTheMostCommonKeysWithinYourLogDataAPI.md#getlogtopkeys) | **Get** /management/logs/{id}/topkeys | Retrieve Most Common Keys For A Log
*SavedQueriesAPI* | [**DeleteSavedQueryId**](docs/SavedQueriesAPI.md#deletesavedqueryid) | **Delete** /query/saved_queries/{saved_query_id} | Delete A Saved Query
*SavedQueriesAPI* | [**GetSavedQueryId**](docs/SavedQueriesAPI.md#getsavedqueryid) | **Get** /query/saved_queries/{saved_query_id} | Retrieve A Saved Query
*SavedQueriesAPI* | [**ListSavedQueries**](docs/SavedQueriesAPI.md#listsavedqueries) | **Get** /query/saved_queries | List All Saved Queries
*SavedQueriesAPI* | [**PatchSavedQueryId**](docs/SavedQueriesAPI.md#patchsavedqueryid) | **Patch** /query/saved_queries/{saved_query_id} | Modify A Saved Query
*SavedQueriesAPI* | [**PostSavedQueryRoot**](docs/SavedQueriesAPI.md#postsavedqueryroot) | **Post** /query/saved_queries | Create A Saved Query
*SavedQueriesAPI* | [**PutSavedQueryId**](docs/SavedQueriesAPI.md#putsavedqueryid) | **Put** /query/saved_queries/{saved_query_id} | Replace A Saved Query
*SavedQueriesAPI* | [**UseSavedQuery**](docs/SavedQueriesAPI.md#usesavedquery) | **Get** /query/saved_query/{saved_query_id} | Use a Saved Query (logs specified)
*SavedQueriesAPI* | [**UseSavedQueryNoLogs**](docs/SavedQueriesAPI.md#usesavedquerynologs) | **Get** /query/logs/{log_keys}/{saved_query_id} | Use a Saved Query (logs unspecified)
*WatchIncomingLogDataInRealtimeAPI* | [**GetLiveLogs**](docs/WatchIncomingLogDataInRealtimeAPI.md#getlivelogs) | **Get** /query/live/logs/{log_keys} | Start A Live Tail Feed (Individual Logs)
*WatchIncomingLogDataInRealtimeAPI* | [**LiveSavedQuery**](docs/WatchIncomingLogDataInRealtimeAPI.md#livesavedquery) | **Get** /query/live/saved_query/{saved_query_id} | Start a Live Tail Feed (Saved Query - with logs)
*WatchIncomingLogDataInRealtimeAPI* | [**LiveSavedQueryNoLogs**](docs/WatchIncomingLogDataInRealtimeAPI.md#livesavedquerynologs) | **Get** /query/live/logs/{log_keys}/{saved_query_id} | Start a Live Tail Feed (Saved Query - without logs)
*WatchIncomingLogDataInRealtimeAPI* | [**PollLiveTail**](docs/WatchIncomingLogDataInRealtimeAPI.md#polllivetail) | **Get** /query/live/{id} | Poll a Live Tail Feed
*WatchIncomingLogDataInRealtimeAPI* | [**PostLiveLogs**](docs/WatchIncomingLogDataInRealtimeAPI.md#postlivelogs) | **Post** /query/live/logs | Start A Live Tail Feed (Multiple Logs)


## Documentation For Models

 - [AccountsUsageResponse](docs/AccountsUsageResponse.md)
 - [AccountsUsageResponseDailyUsage](docs/AccountsUsageResponseDailyUsage.md)
 - [AccountsUsageResponsePeriod](docs/AccountsUsageResponsePeriod.md)
 - [ActionResponse](docs/ActionResponse.md)
 - [ActionResponseAnomaly](docs/ActionResponseAnomaly.md)
 - [ActionResponseInactivity](docs/ActionResponseInactivity.md)
 - [AlertNotificationSetting](docs/AlertNotificationSetting.md)
 - [AlertNotificationSettingRequest](docs/AlertNotificationSettingRequest.md)
 - [AnomalyLeql](docs/AnomalyLeql.md)
 - [AuditLog](docs/AuditLog.md)
 - [AuditLogList](docs/AuditLogList.md)
 - [AuditLogSingle](docs/AuditLogSingle.md)
 - [BasicLeql](docs/BasicLeql.md)
 - [BasicTagResponse](docs/BasicTagResponse.md)
 - [ContextApiLinksInner](docs/ContextApiLinksInner.md)
 - [ContextContinueResponse](docs/ContextContinueResponse.md)
 - [ContextResponse](docs/ContextResponse.md)
 - [CreateBasicTag](docs/CreateBasicTag.md)
 - [CreateBasicTagTag](docs/CreateBasicTagTag.md)
 - [CreateLabel](docs/CreateLabel.md)
 - [CreateLabelLabel](docs/CreateLabelLabel.md)
 - [CreateLogset](docs/CreateLogset.md)
 - [CreateLogsetLogset](docs/CreateLogsetLogset.md)
 - [CreateOrPutLog](docs/CreateOrPutLog.md)
 - [CreateOrPutLogLog](docs/CreateOrPutLogLog.md)
 - [CreateOrPutLogLogLogsetsInfoInner](docs/CreateOrPutLogLogLogsetsInfoInner.md)
 - [CreateOrPutTagInactivityAlert](docs/CreateOrPutTagInactivityAlert.md)
 - [CreateOrPutTagInactivityAlertTag](docs/CreateOrPutTagInactivityAlertTag.md)
 - [CreatePatchTarget](docs/CreatePatchTarget.md)
 - [CreatePatchTargetParamsSet](docs/CreatePatchTargetParamsSet.md)
 - [CreatePutTarget](docs/CreatePutTarget.md)
 - [CreatePutTargetTarget](docs/CreatePutTargetTarget.md)
 - [CreateS3ArchivingSetup](docs/CreateS3ArchivingSetup.md)
 - [CreateS3ArchivingSetupS3setup](docs/CreateS3ArchivingSetupS3setup.md)
 - [CreateSavedQuery](docs/CreateSavedQuery.md)
 - [CreateSavedQuerySavedQuery](docs/CreateSavedQuerySavedQuery.md)
 - [CreateScheduledQuery](docs/CreateScheduledQuery.md)
 - [CreateTagAction](docs/CreateTagAction.md)
 - [CreateTagActionAnomaly](docs/CreateTagActionAnomaly.md)
 - [CreateTagActionInactivity](docs/CreateTagActionInactivity.md)
 - [CreateTagAnomalyAlert](docs/CreateTagAnomalyAlert.md)
 - [CreateTagAnomalyAlertTag](docs/CreateTagAnomalyAlertTag.md)
 - [CreateTagTarget](docs/CreateTagTarget.md)
 - [CreateVariable201Response](docs/CreateVariable201Response.md)
 - [CreateVariableRequest](docs/CreateVariableRequest.md)
 - [EventLiveTail](docs/EventLiveTail.md)
 - [EventQuery](docs/EventQuery.md)
 - [EventsContinuation](docs/EventsContinuation.md)
 - [EventsResponse](docs/EventsResponse.md)
 - [ExportJob](docs/ExportJob.md)
 - [ExportJobLinksInner](docs/ExportJobLinksInner.md)
 - [ExportJobResponse](docs/ExportJobResponse.md)
 - [ExportJobsResponse](docs/ExportJobsResponse.md)
 - [FunctionDef](docs/FunctionDef.md)
 - [GetManagementAlertNotificationSettings200Response](docs/GetManagementAlertNotificationSettings200Response.md)
 - [GetManagementTagActionTargets200Response](docs/GetManagementTagActionTargets200Response.md)
 - [GetManagementTags200Response](docs/GetManagementTags200Response.md)
 - [GetManagementTags200ResponseTagsInner](docs/GetManagementTags200ResponseTagsInner.md)
 - [GetQueryLogs200Response](docs/GetQueryLogs200Response.md)
 - [GetQueryLogs202Response](docs/GetQueryLogs202Response.md)
 - [GetTargetAction](docs/GetTargetAction.md)
 - [GroupsTimeseriesInnerValue](docs/GroupsTimeseriesInnerValue.md)
 - [InactivityLeql](docs/InactivityLeql.md)
 - [KvpInfoInner](docs/KvpInfoInner.md)
 - [KvpInfoInnerKey](docs/KvpInfoInnerKey.md)
 - [KvpInfoInnerValue](docs/KvpInfoInnerValue.md)
 - [LabelResponse](docs/LabelResponse.md)
 - [LabelsInner](docs/LabelsInner.md)
 - [LabelsResponse](docs/LabelsResponse.md)
 - [Leql](docs/Leql.md)
 - [LeqlContext](docs/LeqlContext.md)
 - [LeqlContextDuring](docs/LeqlContextDuring.md)
 - [LeqlDuring](docs/LeqlDuring.md)
 - [LeqlLiveTail](docs/LeqlLiveTail.md)
 - [LeqlMetrics](docs/LeqlMetrics.md)
 - [LeqlMetricsEndpoints](docs/LeqlMetricsEndpoints.md)
 - [LeqlPATCH](docs/LeqlPATCH.md)
 - [LeqlPATCHDuring](docs/LeqlPATCHDuring.md)
 - [LeqlVariableRequest](docs/LeqlVariableRequest.md)
 - [LeqlVariableResponse](docs/LeqlVariableResponse.md)
 - [LinksInner](docs/LinksInner.md)
 - [ListEndpointsResponseValue](docs/ListEndpointsResponseValue.md)
 - [ListSavedQueries200Response](docs/ListSavedQueries200Response.md)
 - [ListVariables200Response](docs/ListVariables200Response.md)
 - [LiveTailPollResponse](docs/LiveTailPollResponse.md)
 - [LiveTailResponse](docs/LiveTailResponse.md)
 - [LogInfoResponse](docs/LogInfoResponse.md)
 - [LogResponse](docs/LogResponse.md)
 - [LogTopkeysResponse](docs/LogTopkeysResponse.md)
 - [LogTopkeysResponseTopkeysInner](docs/LogTopkeysResponseTopkeysInner.md)
 - [LogUsageResponse](docs/LogUsageResponse.md)
 - [LogUsageResponseUsage](docs/LogUsageResponseUsage.md)
 - [LogUsageResponseUsageDailyUsage](docs/LogUsageResponseUsageDailyUsage.md)
 - [LogsResponse](docs/LogsResponse.md)
 - [LogsUsageResponse](docs/LogsUsageResponse.md)
 - [LogsUsageResponsePerDayUsage](docs/LogsUsageResponsePerDayUsage.md)
 - [LogsUsageResponsePerDayUsageUsageInner](docs/LogsUsageResponsePerDayUsageUsageInner.md)
 - [LogsUsageResponsePerDayUsageUsageInnerLogUsageInner](docs/LogsUsageResponsePerDayUsageUsageInnerLogUsageInner.md)
 - [LogsetInfo](docs/LogsetInfo.md)
 - [LogsetResponse](docs/LogsetResponse.md)
 - [MemberInfoInner](docs/MemberInfoInner.md)
 - [MemberInfoLogsetInner](docs/MemberInfoLogsetInner.md)
 - [MetricsCreateRequest](docs/MetricsCreateRequest.md)
 - [MetricsInput](docs/MetricsInput.md)
 - [MetricsInputLogsetsById](docs/MetricsInputLogsetsById.md)
 - [MetricsInputLogsetsByIdMetric](docs/MetricsInputLogsetsByIdMetric.md)
 - [MetricsInputLogsetsByIdMetricLogsets](docs/MetricsInputLogsetsByIdMetricLogsets.md)
 - [MetricsInputLogsetsByName](docs/MetricsInputLogsetsByName.md)
 - [MetricsInputLogsetsByNameMetric](docs/MetricsInputLogsetsByNameMetric.md)
 - [MetricsInputLogsetsByNameMetricLogsets](docs/MetricsInputLogsetsByNameMetricLogsets.md)
 - [MetricsInputMetric](docs/MetricsInputMetric.md)
 - [MetricsInputMetricLogs](docs/MetricsInputMetricLogs.md)
 - [MetricsResponse](docs/MetricsResponse.md)
 - [MetricsResponseList](docs/MetricsResponseList.md)
 - [MetricsResponseSingle](docs/MetricsResponseSingle.md)
 - [MetricsResponseSingleMetric](docs/MetricsResponseSingleMetric.md)
 - [NotificationTypeEnum](docs/NotificationTypeEnum.md)
 - [PatchBasicTag](docs/PatchBasicTag.md)
 - [PatchBasicTagTag](docs/PatchBasicTagTag.md)
 - [PatchLabel](docs/PatchLabel.md)
 - [PatchLabelLabel](docs/PatchLabelLabel.md)
 - [PatchManagementTagActionTargetByIdRequest](docs/PatchManagementTagActionTargetByIdRequest.md)
 - [PatchManagementTagByIdRequest](docs/PatchManagementTagByIdRequest.md)
 - [PatchManagementTagByIdRequestTag](docs/PatchManagementTagByIdRequestTag.md)
 - [PatchS3ArchivingSetup](docs/PatchS3ArchivingSetup.md)
 - [PatchS3ArchivingSetupS3setup](docs/PatchS3ArchivingSetupS3setup.md)
 - [PatchSavedQuery](docs/PatchSavedQuery.md)
 - [PatchSavedQuerySavedQuery](docs/PatchSavedQuerySavedQuery.md)
 - [PatchTagWithAnomalyAlert](docs/PatchTagWithAnomalyAlert.md)
 - [PatchTagWithAnomalyAlertTag](docs/PatchTagWithAnomalyAlertTag.md)
 - [PatchTagWithInactivityAlert](docs/PatchTagWithInactivityAlert.md)
 - [PatchTagWithInactivityAlertTag](docs/PatchTagWithInactivityAlertTag.md)
 - [PeriodEnum](docs/PeriodEnum.md)
 - [PollQuery200Response](docs/PollQuery200Response.md)
 - [PostLiveLogsRequest](docs/PostLiveLogsRequest.md)
 - [PostManagementTagActions201Response](docs/PostManagementTagActions201Response.md)
 - [PostManagementTagActionsRequest](docs/PostManagementTagActionsRequest.md)
 - [PostManagementTags201Response](docs/PostManagementTags201Response.md)
 - [PostManagementTagsRequest](docs/PostManagementTagsRequest.md)
 - [PostQueryLogsRequest](docs/PostQueryLogsRequest.md)
 - [PutBasicTag](docs/PutBasicTag.md)
 - [PutBasicTagTag](docs/PutBasicTagTag.md)
 - [PutLabel](docs/PutLabel.md)
 - [PutManagementTagByIdRequest](docs/PutManagementTagByIdRequest.md)
 - [PutTagAnomalyAlert](docs/PutTagAnomalyAlert.md)
 - [PutTagAnomalyAlertTag](docs/PutTagAnomalyAlertTag.md)
 - [QueryApiLinksInner](docs/QueryApiLinksInner.md)
 - [RetentionPeriodEnum](docs/RetentionPeriodEnum.md)
 - [S3ArchivingSetup](docs/S3ArchivingSetup.md)
 - [S3ArchivingSetupS3setup](docs/S3ArchivingSetupS3setup.md)
 - [SavedQuery](docs/SavedQuery.md)
 - [ScheduledQueryResponse](docs/ScheduledQueryResponse.md)
 - [SearchStats](docs/SearchStats.md)
 - [SearchStatsInner](docs/SearchStatsInner.md)
 - [SearchStatsInnerLeql](docs/SearchStatsInnerLeql.md)
 - [SearchStatsInnerStatistics](docs/SearchStatsInnerStatistics.md)
 - [SourcesIdArrayInner](docs/SourcesIdArrayInner.md)
 - [StatisticalContinuation](docs/StatisticalContinuation.md)
 - [StatisticalResponse](docs/StatisticalResponse.md)
 - [Statistics](docs/Statistics.md)
 - [StatisticsMetrics](docs/StatisticsMetrics.md)
 - [StatisticsMetricsTimeseriesInner](docs/StatisticsMetricsTimeseriesInner.md)
 - [Stats](docs/Stats.md)
 - [TagAnomalyAlertResponse](docs/TagAnomalyAlertResponse.md)
 - [TagInactivityAlertResponse](docs/TagInactivityAlertResponse.md)
 - [TagSourceResponse](docs/TagSourceResponse.md)
 - [TargetAlertContentSet](docs/TargetAlertContentSet.md)
 - [TargetParamsSetIconworkflow](docs/TargetParamsSetIconworkflow.md)
 - [TargetParamsSetMailto](docs/TargetParamsSetMailto.md)
 - [TargetParamsSetPagerduty](docs/TargetParamsSetPagerduty.md)
 - [TargetParamsSetSlack](docs/TargetParamsSetSlack.md)
 - [TargetParamsSetSqs](docs/TargetParamsSetSqs.md)
 - [TargetParamsSetWebhook](docs/TargetParamsSetWebhook.md)
 - [TargetResponse](docs/TargetResponse.md)
 - [TargetResponseParamsSet](docs/TargetResponseParamsSet.md)
 - [TimeframePeriod](docs/TimeframePeriod.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### Api Key Authentication

- **Type**: API key
- **API key parameter name**: x-api-key
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Api Key Authentication and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		insightops.ContextAPIKeys,
		map[string]insightops.APIKey{
			"Api Key Authentication": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@rapid7.com
