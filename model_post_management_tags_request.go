/*
InsightOps REST API

### Overview  Our REST API lets you use InsightOps through HTTP requests. Currently, the REST API allows you to perform the majority of the actions available through the UI, and has some additional functionality that is not available through the UI. You may use this API to automate common tasks (for example, via shell scripts), and to generally interact with InsightOps programmatically.  This page precisely describes the REST API and serves as a reference for the API. Each HTTP method and each URL endpoint is documented in a self-contained unit so that users only need to read about the HTTP methods relevant to them.  ### Terminology  * A **log entry** is an individual log event. * A **log** is a collection of log entries, or a single log stream. * A **log set** is a logical-only collection of logs, i.e. logs can be in multiple logsets and deleting a logset only deletes the relation between logs, not the logs themselves. * [Log Entry Query Language](https://docs.rapid7.com/insightops/log-search) (**LEQL**) is the query language used in Insight Ops to search log data.

API version: latest
Contact: support@rapid7.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package insightops

import (
	"encoding/json"
	"fmt"
	"gopkg.in/validator.v2"
)

// PostManagementTagsRequest - struct for PostManagementTagsRequest
type PostManagementTagsRequest struct {
	CreateBasicTag                *CreateBasicTag
	CreateOrPutTagInactivityAlert *CreateOrPutTagInactivityAlert
	CreateTagAnomalyAlert         *CreateTagAnomalyAlert
}

// CreateBasicTagAsPostManagementTagsRequest is a convenience function that returns CreateBasicTag wrapped in PostManagementTagsRequest
func CreateBasicTagAsPostManagementTagsRequest(v *CreateBasicTag) PostManagementTagsRequest {
	return PostManagementTagsRequest{
		CreateBasicTag: v,
	}
}

// CreateOrPutTagInactivityAlertAsPostManagementTagsRequest is a convenience function that returns CreateOrPutTagInactivityAlert wrapped in PostManagementTagsRequest
func CreateOrPutTagInactivityAlertAsPostManagementTagsRequest(v *CreateOrPutTagInactivityAlert) PostManagementTagsRequest {
	return PostManagementTagsRequest{
		CreateOrPutTagInactivityAlert: v,
	}
}

// CreateTagAnomalyAlertAsPostManagementTagsRequest is a convenience function that returns CreateTagAnomalyAlert wrapped in PostManagementTagsRequest
func CreateTagAnomalyAlertAsPostManagementTagsRequest(v *CreateTagAnomalyAlert) PostManagementTagsRequest {
	return PostManagementTagsRequest{
		CreateTagAnomalyAlert: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *PostManagementTagsRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CreateBasicTag
	err = newStrictDecoder(data).Decode(&dst.CreateBasicTag)
	if err == nil {
		jsonCreateBasicTag, _ := json.Marshal(dst.CreateBasicTag)
		if string(jsonCreateBasicTag) == "{}" { // empty struct
			dst.CreateBasicTag = nil
		} else {
			if err = validator.Validate(dst.CreateBasicTag); err != nil {
				dst.CreateBasicTag = nil
			} else {
				match++
			}
		}
	} else {
		dst.CreateBasicTag = nil
	}

	// try to unmarshal data into CreateOrPutTagInactivityAlert
	err = newStrictDecoder(data).Decode(&dst.CreateOrPutTagInactivityAlert)
	if err == nil {
		jsonCreateOrPutTagInactivityAlert, _ := json.Marshal(dst.CreateOrPutTagInactivityAlert)
		if string(jsonCreateOrPutTagInactivityAlert) == "{}" { // empty struct
			dst.CreateOrPutTagInactivityAlert = nil
		} else {
			if err = validator.Validate(dst.CreateOrPutTagInactivityAlert); err != nil {
				dst.CreateOrPutTagInactivityAlert = nil
			} else {
				match++
			}
		}
	} else {
		dst.CreateOrPutTagInactivityAlert = nil
	}

	// try to unmarshal data into CreateTagAnomalyAlert
	err = newStrictDecoder(data).Decode(&dst.CreateTagAnomalyAlert)
	if err == nil {
		jsonCreateTagAnomalyAlert, _ := json.Marshal(dst.CreateTagAnomalyAlert)
		if string(jsonCreateTagAnomalyAlert) == "{}" { // empty struct
			dst.CreateTagAnomalyAlert = nil
		} else {
			if err = validator.Validate(dst.CreateTagAnomalyAlert); err != nil {
				dst.CreateTagAnomalyAlert = nil
			} else {
				match++
			}
		}
	} else {
		dst.CreateTagAnomalyAlert = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CreateBasicTag = nil
		dst.CreateOrPutTagInactivityAlert = nil
		dst.CreateTagAnomalyAlert = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PostManagementTagsRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PostManagementTagsRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PostManagementTagsRequest) MarshalJSON() ([]byte, error) {
	if src.CreateBasicTag != nil {
		return json.Marshal(&src.CreateBasicTag)
	}

	if src.CreateOrPutTagInactivityAlert != nil {
		return json.Marshal(&src.CreateOrPutTagInactivityAlert)
	}

	if src.CreateTagAnomalyAlert != nil {
		return json.Marshal(&src.CreateTagAnomalyAlert)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PostManagementTagsRequest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CreateBasicTag != nil {
		return obj.CreateBasicTag
	}

	if obj.CreateOrPutTagInactivityAlert != nil {
		return obj.CreateOrPutTagInactivityAlert
	}

	if obj.CreateTagAnomalyAlert != nil {
		return obj.CreateTagAnomalyAlert
	}

	// all schemas are nil
	return nil
}

type NullablePostManagementTagsRequest struct {
	value *PostManagementTagsRequest
	isSet bool
}

func (v NullablePostManagementTagsRequest) Get() *PostManagementTagsRequest {
	return v.value
}

func (v *NullablePostManagementTagsRequest) Set(val *PostManagementTagsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostManagementTagsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostManagementTagsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostManagementTagsRequest(val *PostManagementTagsRequest) *NullablePostManagementTagsRequest {
	return &NullablePostManagementTagsRequest{value: val, isSet: true}
}

func (v NullablePostManagementTagsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostManagementTagsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
