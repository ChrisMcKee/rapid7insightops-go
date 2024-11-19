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
)

// checks if the PatchS3ArchivingSetup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchS3ArchivingSetup{}

// PatchS3ArchivingSetup struct for PatchS3ArchivingSetup
type PatchS3ArchivingSetup struct {
	S3setup *PatchS3ArchivingSetupS3setup `json:"s3setup,omitempty"`
}

// NewPatchS3ArchivingSetup instantiates a new PatchS3ArchivingSetup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchS3ArchivingSetup() *PatchS3ArchivingSetup {
	this := PatchS3ArchivingSetup{}
	return &this
}

// NewPatchS3ArchivingSetupWithDefaults instantiates a new PatchS3ArchivingSetup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchS3ArchivingSetupWithDefaults() *PatchS3ArchivingSetup {
	this := PatchS3ArchivingSetup{}
	return &this
}

// GetS3setup returns the S3setup field value if set, zero value otherwise.
func (o *PatchS3ArchivingSetup) GetS3setup() PatchS3ArchivingSetupS3setup {
	if o == nil || IsNil(o.S3setup) {
		var ret PatchS3ArchivingSetupS3setup
		return ret
	}
	return *o.S3setup
}

// GetS3setupOk returns a tuple with the S3setup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchS3ArchivingSetup) GetS3setupOk() (*PatchS3ArchivingSetupS3setup, bool) {
	if o == nil || IsNil(o.S3setup) {
		return nil, false
	}
	return o.S3setup, true
}

// HasS3setup returns a boolean if a field has been set.
func (o *PatchS3ArchivingSetup) HasS3setup() bool {
	if o != nil && !IsNil(o.S3setup) {
		return true
	}

	return false
}

// SetS3setup gets a reference to the given PatchS3ArchivingSetupS3setup and assigns it to the S3setup field.
func (o *PatchS3ArchivingSetup) SetS3setup(v PatchS3ArchivingSetupS3setup) {
	o.S3setup = &v
}

func (o PatchS3ArchivingSetup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchS3ArchivingSetup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.S3setup) {
		toSerialize["s3setup"] = o.S3setup
	}
	return toSerialize, nil
}

type NullablePatchS3ArchivingSetup struct {
	value *PatchS3ArchivingSetup
	isSet bool
}

func (v NullablePatchS3ArchivingSetup) Get() *PatchS3ArchivingSetup {
	return v.value
}

func (v *NullablePatchS3ArchivingSetup) Set(val *PatchS3ArchivingSetup) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchS3ArchivingSetup) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchS3ArchivingSetup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchS3ArchivingSetup(val *PatchS3ArchivingSetup) *NullablePatchS3ArchivingSetup {
	return &NullablePatchS3ArchivingSetup{value: val, isSet: true}
}

func (v NullablePatchS3ArchivingSetup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchS3ArchivingSetup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}