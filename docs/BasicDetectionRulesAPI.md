# invoker\BasicDetectionRulesAPI

All URIs are relative to *https://eu.rest.logs.insight.rapid7.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteLabelById**](BasicDetectionRulesAPI.md#DeleteLabelById) | **Delete** /management/labels/{label_id} | Delete a Label
[**DeleteManagementTagActionById**](BasicDetectionRulesAPI.md#DeleteManagementTagActionById) | **Delete** /management/actions/{action_id} | Delete a Notification
[**DeleteManagementTagById**](BasicDetectionRulesAPI.md#DeleteManagementTagById) | **Delete** /management/tags/{rule_id} | Delete a Basic Detection Rule
[**DeleteManagementTargetById**](BasicDetectionRulesAPI.md#DeleteManagementTargetById) | **Delete** /management/targets/{target_id} | Delete a Target
[**GetLabelById**](BasicDetectionRulesAPI.md#GetLabelById) | **Get** /management/labels/{label_id} | Retrieve a Label
[**GetLabels**](BasicDetectionRulesAPI.md#GetLabels) | **Get** /management/labels | List all Labels
[**GetManagementAlertNotificationSettings**](BasicDetectionRulesAPI.md#GetManagementAlertNotificationSettings) | **Get** /management/actions | List All Notifications
[**GetManagementTagActionById**](BasicDetectionRulesAPI.md#GetManagementTagActionById) | **Get** /management/actions/{action_id} | Retrieve a Notification
[**GetManagementTagActionTargets**](BasicDetectionRulesAPI.md#GetManagementTagActionTargets) | **Get** /management/actions/{action_id}/targets | List All Targets Attached To A Notification
[**GetManagementTagById**](BasicDetectionRulesAPI.md#GetManagementTagById) | **Get** /management/tags/{rule_id} | Retrieve a Basic Detection Rule
[**GetManagementTags**](BasicDetectionRulesAPI.md#GetManagementTags) | **Get** /management/tags | List All Tags
[**GetManagementTargetById**](BasicDetectionRulesAPI.md#GetManagementTargetById) | **Get** /management/targets/{target_id} | Retrieve a Target
[**GetManagementTargets**](BasicDetectionRulesAPI.md#GetManagementTargets) | **Get** /management/targets | List All Targets
[**PatchLabelById**](BasicDetectionRulesAPI.md#PatchLabelById) | **Patch** /management/labels/{label_id} | Update a Label
[**PatchManagementTagActionById**](BasicDetectionRulesAPI.md#PatchManagementTagActionById) | **Patch** /management/actions/{action_id} | Modify a Notification
[**PatchManagementTagActionTargetById**](BasicDetectionRulesAPI.md#PatchManagementTagActionTargetById) | **Patch** /management/actions/{action_id}/targets | Modify the Targets Attached To A Notification
[**PatchManagementTagById**](BasicDetectionRulesAPI.md#PatchManagementTagById) | **Patch** /management/tags/{rule_id} | Modify a Basic Detection Rule
[**PostLabels**](BasicDetectionRulesAPI.md#PostLabels) | **Post** /management/labels | Create a Label
[**PostManagementTagActions**](BasicDetectionRulesAPI.md#PostManagementTagActions) | **Post** /management/actions | Create a Notification
[**PostManagementTags**](BasicDetectionRulesAPI.md#PostManagementTags) | **Post** /management/tags | Create A Basic Detection Rule
[**PostManagementTargets**](BasicDetectionRulesAPI.md#PostManagementTargets) | **Post** /management/targets | Create a Target
[**PutLabelById**](BasicDetectionRulesAPI.md#PutLabelById) | **Put** /management/labels/{label_id} | Replace a Label
[**PutManagementTagActionById**](BasicDetectionRulesAPI.md#PutManagementTagActionById) | **Put** /management/actions/{action_id} | Replace a Notification
[**PutManagementTagById**](BasicDetectionRulesAPI.md#PutManagementTagById) | **Put** /management/tags/{rule_id} | Replace a Basic Detection Rule
[**PutManagementTargetById**](BasicDetectionRulesAPI.md#PutManagementTargetById) | **Put** /management/targets/{target_id} | Replace a Target



## DeleteLabelById

> DeleteLabelById(ctx, labelId).Execute()

Delete a Label



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
	labelId := "4a8869fa-0843-4edd-85f5-bf7d6b80d6c5" // string | \\ The ID of the label to be queried. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BasicDetectionRulesAPI.DeleteLabelById(context.Background(), labelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicDetectionRulesAPI.DeleteLabelById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**labelId** | **string** | \\ The ID of the label to be queried.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLabelByIdRequest struct via the builder pattern


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


## DeleteManagementTagActionById

> DeleteManagementTagActionById(ctx, actionId).Execute()

Delete a Notification



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
	actionId := "00000000-0000-00cf-0000-000000000000" // string | \\ The id of the notification. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BasicDetectionRulesAPI.DeleteManagementTagActionById(context.Background(), actionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicDetectionRulesAPI.DeleteManagementTagActionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actionId** | **string** | \\ The id of the notification.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteManagementTagActionByIdRequest struct via the builder pattern


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


## DeleteManagementTagById

> DeleteManagementTagById(ctx, tagId).Execute()

Delete a Basic Detection Rule



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
	tagId := "00000000-0000-00cf-0000-000000000000" // string | \\ The id of the change detection rule. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BasicDetectionRulesAPI.DeleteManagementTagById(context.Background(), tagId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicDetectionRulesAPI.DeleteManagementTagById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **string** | \\ The id of the change detection rule.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteManagementTagByIdRequest struct via the builder pattern


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


## DeleteManagementTargetById

> DeleteManagementTargetById(ctx, targetId).Execute()

Delete a Target



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
	targetId := "4a8869fa-0843-4edd-85f5-bf7d6b80d6c5" // string | \\ The id of the notification target. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BasicDetectionRulesAPI.DeleteManagementTargetById(context.Background(), targetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicDetectionRulesAPI.DeleteManagementTargetById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetId** | **string** | \\ The id of the notification target.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteManagementTargetByIdRequest struct via the builder pattern


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


## GetLabelById

> LabelResponse GetLabelById(ctx, labelId).Execute()

Retrieve a Label



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
	labelId := "4a8869fa-0843-4edd-85f5-bf7d6b80d6c5" // string | \\ The ID of the label to be queried. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicDetectionRulesAPI.GetLabelById(context.Background(), labelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicDetectionRulesAPI.GetLabelById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLabelById`: LabelResponse
	fmt.Fprintf(os.Stdout, "Response from `BasicDetectionRulesAPI.GetLabelById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**labelId** | **string** | \\ The ID of the label to be queried.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLabelByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LabelResponse**](LabelResponse.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLabels

> LabelsResponse GetLabels(ctx).Execute()

List all Labels



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
	resp, r, err := apiClient.BasicDetectionRulesAPI.GetLabels(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicDetectionRulesAPI.GetLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLabels`: LabelsResponse
	fmt.Fprintf(os.Stdout, "Response from `BasicDetectionRulesAPI.GetLabels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLabelsRequest struct via the builder pattern


### Return type

[**LabelsResponse**](LabelsResponse.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagementAlertNotificationSettings

> GetManagementAlertNotificationSettings200Response GetManagementAlertNotificationSettings(ctx).Execute()

List All Notifications



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
	resp, r, err := apiClient.BasicDetectionRulesAPI.GetManagementAlertNotificationSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicDetectionRulesAPI.GetManagementAlertNotificationSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManagementAlertNotificationSettings`: GetManagementAlertNotificationSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `BasicDetectionRulesAPI.GetManagementAlertNotificationSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetManagementAlertNotificationSettingsRequest struct via the builder pattern


### Return type

[**GetManagementAlertNotificationSettings200Response**](GetManagementAlertNotificationSettings200Response.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagementTagActionById

> PostManagementTagActions201Response GetManagementTagActionById(ctx, actionId).Execute()

Retrieve a Notification



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
	actionId := "00000000-0000-00cf-0000-000000000000" // string | \\ The id of the notification. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicDetectionRulesAPI.GetManagementTagActionById(context.Background(), actionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicDetectionRulesAPI.GetManagementTagActionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManagementTagActionById`: PostManagementTagActions201Response
	fmt.Fprintf(os.Stdout, "Response from `BasicDetectionRulesAPI.GetManagementTagActionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actionId** | **string** | \\ The id of the notification.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetManagementTagActionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PostManagementTagActions201Response**](PostManagementTagActions201Response.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagementTagActionTargets

> GetManagementTagActionTargets200Response GetManagementTagActionTargets(ctx, actionId).Execute()

List All Targets Attached To A Notification



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
	actionId := "00000000-0000-00cf-0000-000000000000" // string | \\ The id of the notification. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicDetectionRulesAPI.GetManagementTagActionTargets(context.Background(), actionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicDetectionRulesAPI.GetManagementTagActionTargets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManagementTagActionTargets`: GetManagementTagActionTargets200Response
	fmt.Fprintf(os.Stdout, "Response from `BasicDetectionRulesAPI.GetManagementTagActionTargets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actionId** | **string** | \\ The id of the notification.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetManagementTagActionTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetManagementTagActionTargets200Response**](GetManagementTagActionTargets200Response.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagementTagById

> PostManagementTags201Response GetManagementTagById(ctx, tagId).Execute()

Retrieve a Basic Detection Rule



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
	tagId := "00000000-0000-00cf-0000-000000000000" // string | \\ The id of the change detection rule. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicDetectionRulesAPI.GetManagementTagById(context.Background(), tagId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicDetectionRulesAPI.GetManagementTagById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManagementTagById`: PostManagementTags201Response
	fmt.Fprintf(os.Stdout, "Response from `BasicDetectionRulesAPI.GetManagementTagById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **string** | \\ The id of the change detection rule.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetManagementTagByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PostManagementTags201Response**](PostManagementTags201Response.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagementTags

> GetManagementTags200Response GetManagementTags(ctx).Execute()

List All Tags



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
	resp, r, err := apiClient.BasicDetectionRulesAPI.GetManagementTags(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicDetectionRulesAPI.GetManagementTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManagementTags`: GetManagementTags200Response
	fmt.Fprintf(os.Stdout, "Response from `BasicDetectionRulesAPI.GetManagementTags`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetManagementTagsRequest struct via the builder pattern


### Return type

[**GetManagementTags200Response**](GetManagementTags200Response.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagementTargetById

> TargetResponse GetManagementTargetById(ctx, targetId).Execute()

Retrieve a Target



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
	targetId := "4a8869fa-0843-4edd-85f5-bf7d6b80d6c5" // string | \\ The id of the notification target. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicDetectionRulesAPI.GetManagementTargetById(context.Background(), targetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicDetectionRulesAPI.GetManagementTargetById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManagementTargetById`: TargetResponse
	fmt.Fprintf(os.Stdout, "Response from `BasicDetectionRulesAPI.GetManagementTargetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetId** | **string** | \\ The id of the notification target.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetManagementTargetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TargetResponse**](TargetResponse.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagementTargets

> TargetResponse GetManagementTargets(ctx).Execute()

List All Targets



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
	resp, r, err := apiClient.BasicDetectionRulesAPI.GetManagementTargets(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicDetectionRulesAPI.GetManagementTargets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManagementTargets`: TargetResponse
	fmt.Fprintf(os.Stdout, "Response from `BasicDetectionRulesAPI.GetManagementTargets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetManagementTargetsRequest struct via the builder pattern


### Return type

[**TargetResponse**](TargetResponse.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchLabelById

> LabelResponse PatchLabelById(ctx, labelId).PatchLabel(patchLabel).Execute()

Update a Label



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
	labelId := "4a8869fa-0843-4edd-85f5-bf7d6b80d6c5" // string | \\ The ID of the label to be queried. 
	patchLabel := *openapiclient.NewPatchLabel() // PatchLabel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicDetectionRulesAPI.PatchLabelById(context.Background(), labelId).PatchLabel(patchLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicDetectionRulesAPI.PatchLabelById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchLabelById`: LabelResponse
	fmt.Fprintf(os.Stdout, "Response from `BasicDetectionRulesAPI.PatchLabelById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**labelId** | **string** | \\ The ID of the label to be queried.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchLabelByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchLabel** | [**PatchLabel**](PatchLabel.md) |  | 

### Return type

[**LabelResponse**](LabelResponse.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchManagementTagActionById

> PostManagementTagActions201Response PatchManagementTagActionById(ctx, actionId).PostManagementTagActionsRequest(postManagementTagActionsRequest).Execute()

Modify a Notification



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
	actionId := "00000000-0000-00cf-0000-000000000000" // string | \\ The id of the notification. 
	postManagementTagActionsRequest := *openapiclient.NewPostManagementTagActionsRequest() // PostManagementTagActionsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicDetectionRulesAPI.PatchManagementTagActionById(context.Background(), actionId).PostManagementTagActionsRequest(postManagementTagActionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicDetectionRulesAPI.PatchManagementTagActionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchManagementTagActionById`: PostManagementTagActions201Response
	fmt.Fprintf(os.Stdout, "Response from `BasicDetectionRulesAPI.PatchManagementTagActionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actionId** | **string** | \\ The id of the notification.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchManagementTagActionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postManagementTagActionsRequest** | [**PostManagementTagActionsRequest**](PostManagementTagActionsRequest.md) |  | 

### Return type

[**PostManagementTagActions201Response**](PostManagementTagActions201Response.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchManagementTagActionTargetById

> PatchManagementTagActionTargetByIdRequest PatchManagementTagActionTargetById(ctx, actionId).PatchManagementTagActionTargetByIdRequest(patchManagementTagActionTargetByIdRequest).Execute()

Modify the Targets Attached To A Notification



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
	actionId := "00000000-0000-00cf-0000-000000000000" // string | \\ The id of the notification. 
	patchManagementTagActionTargetByIdRequest := *openapiclient.NewPatchManagementTagActionTargetByIdRequest() // PatchManagementTagActionTargetByIdRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicDetectionRulesAPI.PatchManagementTagActionTargetById(context.Background(), actionId).PatchManagementTagActionTargetByIdRequest(patchManagementTagActionTargetByIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicDetectionRulesAPI.PatchManagementTagActionTargetById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchManagementTagActionTargetById`: PatchManagementTagActionTargetByIdRequest
	fmt.Fprintf(os.Stdout, "Response from `BasicDetectionRulesAPI.PatchManagementTagActionTargetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actionId** | **string** | \\ The id of the notification.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchManagementTagActionTargetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchManagementTagActionTargetByIdRequest** | [**PatchManagementTagActionTargetByIdRequest**](PatchManagementTagActionTargetByIdRequest.md) |  | 

### Return type

[**PatchManagementTagActionTargetByIdRequest**](PatchManagementTagActionTargetByIdRequest.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchManagementTagById

> PostManagementTags201Response PatchManagementTagById(ctx, tagId).PatchManagementTagByIdRequest(patchManagementTagByIdRequest).Execute()

Modify a Basic Detection Rule



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
	tagId := "00000000-0000-00cf-0000-000000000000" // string | \\ The id of the change detection rule. 
	patchManagementTagByIdRequest := *openapiclient.NewPatchManagementTagByIdRequest() // PatchManagementTagByIdRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicDetectionRulesAPI.PatchManagementTagById(context.Background(), tagId).PatchManagementTagByIdRequest(patchManagementTagByIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicDetectionRulesAPI.PatchManagementTagById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchManagementTagById`: PostManagementTags201Response
	fmt.Fprintf(os.Stdout, "Response from `BasicDetectionRulesAPI.PatchManagementTagById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **string** | \\ The id of the change detection rule.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchManagementTagByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchManagementTagByIdRequest** | [**PatchManagementTagByIdRequest**](PatchManagementTagByIdRequest.md) |  | 

### Return type

[**PostManagementTags201Response**](PostManagementTags201Response.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLabels

> LabelResponse PostLabels(ctx).CreateLabel(createLabel).Execute()

Create a Label



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
	createLabel := *openapiclient.NewCreateLabel() // CreateLabel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicDetectionRulesAPI.PostLabels(context.Background()).CreateLabel(createLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicDetectionRulesAPI.PostLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLabels`: LabelResponse
	fmt.Fprintf(os.Stdout, "Response from `BasicDetectionRulesAPI.PostLabels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createLabel** | [**CreateLabel**](CreateLabel.md) |  | 

### Return type

[**LabelResponse**](LabelResponse.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostManagementTagActions

> PostManagementTagActions201Response PostManagementTagActions(ctx).PostManagementTagActionsRequest(postManagementTagActionsRequest).Execute()

Create a Notification



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
	postManagementTagActionsRequest := *openapiclient.NewPostManagementTagActionsRequest() // PostManagementTagActionsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicDetectionRulesAPI.PostManagementTagActions(context.Background()).PostManagementTagActionsRequest(postManagementTagActionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicDetectionRulesAPI.PostManagementTagActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostManagementTagActions`: PostManagementTagActions201Response
	fmt.Fprintf(os.Stdout, "Response from `BasicDetectionRulesAPI.PostManagementTagActions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostManagementTagActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postManagementTagActionsRequest** | [**PostManagementTagActionsRequest**](PostManagementTagActionsRequest.md) |  | 

### Return type

[**PostManagementTagActions201Response**](PostManagementTagActions201Response.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostManagementTags

> PostManagementTags201Response PostManagementTags(ctx).PostManagementTagsRequest(postManagementTagsRequest).Execute()

Create A Basic Detection Rule



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
	postManagementTagsRequest := openapiclient.postManagementTags_request{CreateBasicTag: openapiclient.NewCreateBasicTag()} // PostManagementTagsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicDetectionRulesAPI.PostManagementTags(context.Background()).PostManagementTagsRequest(postManagementTagsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicDetectionRulesAPI.PostManagementTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostManagementTags`: PostManagementTags201Response
	fmt.Fprintf(os.Stdout, "Response from `BasicDetectionRulesAPI.PostManagementTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostManagementTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postManagementTagsRequest** | [**PostManagementTagsRequest**](PostManagementTagsRequest.md) |  | 

### Return type

[**PostManagementTags201Response**](PostManagementTags201Response.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostManagementTargets

> TargetResponse PostManagementTargets(ctx).CreatePutTarget(createPutTarget).Execute()

Create a Target



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
	createPutTarget := *openapiclient.NewCreatePutTarget() // CreatePutTarget | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicDetectionRulesAPI.PostManagementTargets(context.Background()).CreatePutTarget(createPutTarget).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicDetectionRulesAPI.PostManagementTargets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostManagementTargets`: TargetResponse
	fmt.Fprintf(os.Stdout, "Response from `BasicDetectionRulesAPI.PostManagementTargets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostManagementTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPutTarget** | [**CreatePutTarget**](CreatePutTarget.md) |  | 

### Return type

[**TargetResponse**](TargetResponse.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutLabelById

> LabelResponse PutLabelById(ctx, labelId).PutLabel(putLabel).Execute()

Replace a Label



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
	labelId := "4a8869fa-0843-4edd-85f5-bf7d6b80d6c5" // string | \\ The ID of the label to be queried. 
	putLabel := *openapiclient.NewPutLabel() // PutLabel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicDetectionRulesAPI.PutLabelById(context.Background(), labelId).PutLabel(putLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicDetectionRulesAPI.PutLabelById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutLabelById`: LabelResponse
	fmt.Fprintf(os.Stdout, "Response from `BasicDetectionRulesAPI.PutLabelById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**labelId** | **string** | \\ The ID of the label to be queried.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutLabelByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putLabel** | [**PutLabel**](PutLabel.md) |  | 

### Return type

[**LabelResponse**](LabelResponse.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutManagementTagActionById

> PostManagementTagActions201Response PutManagementTagActionById(ctx, actionId).PostManagementTagActionsRequest(postManagementTagActionsRequest).Execute()

Replace a Notification



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
	actionId := "00000000-0000-00cf-0000-000000000000" // string | \\ The id of the notification. 
	postManagementTagActionsRequest := *openapiclient.NewPostManagementTagActionsRequest() // PostManagementTagActionsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicDetectionRulesAPI.PutManagementTagActionById(context.Background(), actionId).PostManagementTagActionsRequest(postManagementTagActionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicDetectionRulesAPI.PutManagementTagActionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutManagementTagActionById`: PostManagementTagActions201Response
	fmt.Fprintf(os.Stdout, "Response from `BasicDetectionRulesAPI.PutManagementTagActionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actionId** | **string** | \\ The id of the notification.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutManagementTagActionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postManagementTagActionsRequest** | [**PostManagementTagActionsRequest**](PostManagementTagActionsRequest.md) |  | 

### Return type

[**PostManagementTagActions201Response**](PostManagementTagActions201Response.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutManagementTagById

> PostManagementTags201Response PutManagementTagById(ctx, tagId).PutManagementTagByIdRequest(putManagementTagByIdRequest).Execute()

Replace a Basic Detection Rule



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
	tagId := "00000000-0000-00cf-0000-000000000000" // string | \\ The id of the change detection rule. 
	putManagementTagByIdRequest := openapiclient.putManagementTagById_request{CreateOrPutTagInactivityAlert: openapiclient.NewCreateOrPutTagInactivityAlert()} // PutManagementTagByIdRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicDetectionRulesAPI.PutManagementTagById(context.Background(), tagId).PutManagementTagByIdRequest(putManagementTagByIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicDetectionRulesAPI.PutManagementTagById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutManagementTagById`: PostManagementTags201Response
	fmt.Fprintf(os.Stdout, "Response from `BasicDetectionRulesAPI.PutManagementTagById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **string** | \\ The id of the change detection rule.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutManagementTagByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putManagementTagByIdRequest** | [**PutManagementTagByIdRequest**](PutManagementTagByIdRequest.md) |  | 

### Return type

[**PostManagementTags201Response**](PostManagementTags201Response.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutManagementTargetById

> TargetResponse PutManagementTargetById(ctx, targetId).CreatePutTarget(createPutTarget).Execute()

Replace a Target



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
	targetId := "4a8869fa-0843-4edd-85f5-bf7d6b80d6c5" // string | \\ The id of the notification target. 
	createPutTarget := *openapiclient.NewCreatePutTarget() // CreatePutTarget | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicDetectionRulesAPI.PutManagementTargetById(context.Background(), targetId).CreatePutTarget(createPutTarget).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicDetectionRulesAPI.PutManagementTargetById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutManagementTargetById`: TargetResponse
	fmt.Fprintf(os.Stdout, "Response from `BasicDetectionRulesAPI.PutManagementTargetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetId** | **string** | \\ The id of the notification target.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutManagementTargetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createPutTarget** | [**CreatePutTarget**](CreatePutTarget.md) |  | 

### Return type

[**TargetResponse**](TargetResponse.md)

### Authorization

[Api Key Authentication](../README.md#Api Key Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

