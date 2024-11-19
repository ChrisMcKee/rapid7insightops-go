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

// GetQueryLogs200Response - struct for GetQueryLogs200Response
type GetQueryLogs200Response struct {
	EventsResponse      *EventsResponse
	StatisticalResponse *StatisticalResponse
}

// EventsResponseAsGetQueryLogs200Response is a convenience function that returns EventsResponse wrapped in GetQueryLogs200Response
func EventsResponseAsGetQueryLogs200Response(v *EventsResponse) GetQueryLogs200Response {
	return GetQueryLogs200Response{
		EventsResponse: v,
	}
}

// StatisticalResponseAsGetQueryLogs200Response is a convenience function that returns StatisticalResponse wrapped in GetQueryLogs200Response
func StatisticalResponseAsGetQueryLogs200Response(v *StatisticalResponse) GetQueryLogs200Response {
	return GetQueryLogs200Response{
		StatisticalResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetQueryLogs200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into EventsResponse
	err = newStrictDecoder(data).Decode(&dst.EventsResponse)
	if err == nil {
		jsonEventsResponse, _ := json.Marshal(dst.EventsResponse)
		if string(jsonEventsResponse) == "{}" { // empty struct
			dst.EventsResponse = nil
		} else {
			if err = validator.Validate(dst.EventsResponse); err != nil {
				dst.EventsResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.EventsResponse = nil
	}

	// try to unmarshal data into StatisticalResponse
	err = newStrictDecoder(data).Decode(&dst.StatisticalResponse)
	if err == nil {
		jsonStatisticalResponse, _ := json.Marshal(dst.StatisticalResponse)
		if string(jsonStatisticalResponse) == "{}" { // empty struct
			dst.StatisticalResponse = nil
		} else {
			if err = validator.Validate(dst.StatisticalResponse); err != nil {
				dst.StatisticalResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.StatisticalResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.EventsResponse = nil
		dst.StatisticalResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetQueryLogs200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetQueryLogs200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetQueryLogs200Response) MarshalJSON() ([]byte, error) {
	if src.EventsResponse != nil {
		return json.Marshal(&src.EventsResponse)
	}

	if src.StatisticalResponse != nil {
		return json.Marshal(&src.StatisticalResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetQueryLogs200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.EventsResponse != nil {
		return obj.EventsResponse
	}

	if obj.StatisticalResponse != nil {
		return obj.StatisticalResponse
	}

	// all schemas are nil
	return nil
}

type NullableGetQueryLogs200Response struct {
	value *GetQueryLogs200Response
	isSet bool
}

func (v NullableGetQueryLogs200Response) Get() *GetQueryLogs200Response {
	return v.value
}

func (v *NullableGetQueryLogs200Response) Set(val *GetQueryLogs200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetQueryLogs200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetQueryLogs200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetQueryLogs200Response(val *GetQueryLogs200Response) *NullableGetQueryLogs200Response {
	return &NullableGetQueryLogs200Response{value: val, isSet: true}
}

func (v NullableGetQueryLogs200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetQueryLogs200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
