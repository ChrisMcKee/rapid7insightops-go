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
)

// PeriodEnum the model 'PeriodEnum'
type PeriodEnum string

// List of period_enum
const (
	_5_MINUTE  PeriodEnum = "5Minute"
	_10_MINUTE PeriodEnum = "10Minute"
	_15_MINUTE PeriodEnum = "15Minute"
	_30_MINUTE PeriodEnum = "30Minute"
)

// All allowed values of PeriodEnum enum
var AllowedPeriodEnumEnumValues = []PeriodEnum{
	"Hour",
	"Day",
	"Minute",
	"5Minute",
	"10Minute",
	"15Minute",
	"30Minute",
}

func (v *PeriodEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PeriodEnum(value)
	for _, existing := range AllowedPeriodEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PeriodEnum", value)
}

// NewPeriodEnumFromValue returns a pointer to a valid PeriodEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPeriodEnumFromValue(v string) (*PeriodEnum, error) {
	ev := PeriodEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PeriodEnum: valid values are %v", v, AllowedPeriodEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PeriodEnum) IsValid() bool {
	for _, existing := range AllowedPeriodEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to period_enum value
func (v PeriodEnum) Ptr() *PeriodEnum {
	return &v
}

type NullablePeriodEnum struct {
	value *PeriodEnum
	isSet bool
}

func (v NullablePeriodEnum) Get() *PeriodEnum {
	return v.value
}

func (v *NullablePeriodEnum) Set(val *PeriodEnum) {
	v.value = val
	v.isSet = true
}

func (v NullablePeriodEnum) IsSet() bool {
	return v.isSet
}

func (v *NullablePeriodEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePeriodEnum(val *PeriodEnum) *NullablePeriodEnum {
	return &NullablePeriodEnum{value: val, isSet: true}
}

func (v NullablePeriodEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePeriodEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}