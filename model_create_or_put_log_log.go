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
	"encoding/json"
	"fmt"
)

// checks if the CreateOrPutLogLog type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrPutLogLog{}

// CreateOrPutLogLog struct for CreateOrPutLogLog
type CreateOrPutLogLog struct {
	// The name of the log.
	Name interface{} `json:"name"`
	// The log token(s) used for writing to the log. This only applies to token type logs (view the `source_type` parameter).
	Tokens []string `json:"tokens,omitempty"`
	// Structures are internal entities which may apply some additional processing to log entries written to this this log.
	Structures []string `json:"structures,omitempty"`
	// The IP address that the Log Search system receives log entries from. This only applies to syslog type logs (view the `source_type` parameter).
	IpAddress *string `json:"ip_address,omitempty"`
	// The information on each log set that this log is part of.
	LogsetsInfo []CreateOrPutLogLogLogsetsInfoInner `json:"logsets_info,omitempty"`
	// A list of key-value pairs that may indicate some auxiliary information about the log.
	UserData map[string]interface{} `json:"user_data,omitempty"`
	// A categorization of logs that defines how log entries are received by servers. * `syslog` type logs are associated with an IP address and a port (the values appear in the `user_data` field). Any log entries received by the server on that port and from that IP address will belong to that log. * `token` type logs are associated with a log token. Any log entries received by the server on one of the [standard ports](https://docs.rapid7.com/insightops/token-tcp) containing that log token will belong to that log (with the log token removed from the log entry). * `internal` type logs use internal mechanisms to receive log entries, for example any of the logs in the Internal Logs log set.
	SourceType *string `json:"source_type,omitempty"`
	// The seed used to generate the log token (for token type logs).
	TokenSeed *string `json:"token_seed,omitempty"`
}

type _CreateOrPutLogLog CreateOrPutLogLog

// NewCreateOrPutLogLog instantiates a new CreateOrPutLogLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrPutLogLog(name interface{}) *CreateOrPutLogLog {
	this := CreateOrPutLogLog{}
	this.Name = name
	var sourceType string = "token"
	this.SourceType = &sourceType
	return &this
}

// NewCreateOrPutLogLogWithDefaults instantiates a new CreateOrPutLogLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrPutLogLogWithDefaults() *CreateOrPutLogLog {
	this := CreateOrPutLogLog{}
	var sourceType string = "token"
	this.SourceType = &sourceType
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *CreateOrPutLogLog) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOrPutLogLog) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateOrPutLogLog) SetName(v interface{}) {
	o.Name = v
}

// GetTokens returns the Tokens field value if set, zero value otherwise.
func (o *CreateOrPutLogLog) GetTokens() []string {
	if o == nil || IsNil(o.Tokens) {
		var ret []string
		return ret
	}
	return o.Tokens
}

// GetTokensOk returns a tuple with the Tokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrPutLogLog) GetTokensOk() ([]string, bool) {
	if o == nil || IsNil(o.Tokens) {
		return nil, false
	}
	return o.Tokens, true
}

// HasTokens returns a boolean if a field has been set.
func (o *CreateOrPutLogLog) HasTokens() bool {
	if o != nil && !IsNil(o.Tokens) {
		return true
	}

	return false
}

// SetTokens gets a reference to the given []string and assigns it to the Tokens field.
func (o *CreateOrPutLogLog) SetTokens(v []string) {
	o.Tokens = v
}

// GetStructures returns the Structures field value if set, zero value otherwise.
func (o *CreateOrPutLogLog) GetStructures() []string {
	if o == nil || IsNil(o.Structures) {
		var ret []string
		return ret
	}
	return o.Structures
}

// GetStructuresOk returns a tuple with the Structures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrPutLogLog) GetStructuresOk() ([]string, bool) {
	if o == nil || IsNil(o.Structures) {
		return nil, false
	}
	return o.Structures, true
}

// HasStructures returns a boolean if a field has been set.
func (o *CreateOrPutLogLog) HasStructures() bool {
	if o != nil && !IsNil(o.Structures) {
		return true
	}

	return false
}

// SetStructures gets a reference to the given []string and assigns it to the Structures field.
func (o *CreateOrPutLogLog) SetStructures(v []string) {
	o.Structures = v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *CreateOrPutLogLog) GetIpAddress() string {
	if o == nil || IsNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrPutLogLog) GetIpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.IpAddress) {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *CreateOrPutLogLog) HasIpAddress() bool {
	if o != nil && !IsNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *CreateOrPutLogLog) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetLogsetsInfo returns the LogsetsInfo field value if set, zero value otherwise.
func (o *CreateOrPutLogLog) GetLogsetsInfo() []CreateOrPutLogLogLogsetsInfoInner {
	if o == nil || IsNil(o.LogsetsInfo) {
		var ret []CreateOrPutLogLogLogsetsInfoInner
		return ret
	}
	return o.LogsetsInfo
}

// GetLogsetsInfoOk returns a tuple with the LogsetsInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrPutLogLog) GetLogsetsInfoOk() ([]CreateOrPutLogLogLogsetsInfoInner, bool) {
	if o == nil || IsNil(o.LogsetsInfo) {
		return nil, false
	}
	return o.LogsetsInfo, true
}

// HasLogsetsInfo returns a boolean if a field has been set.
func (o *CreateOrPutLogLog) HasLogsetsInfo() bool {
	if o != nil && !IsNil(o.LogsetsInfo) {
		return true
	}

	return false
}

// SetLogsetsInfo gets a reference to the given []CreateOrPutLogLogLogsetsInfoInner and assigns it to the LogsetsInfo field.
func (o *CreateOrPutLogLog) SetLogsetsInfo(v []CreateOrPutLogLogLogsetsInfoInner) {
	o.LogsetsInfo = v
}

// GetUserData returns the UserData field value if set, zero value otherwise.
func (o *CreateOrPutLogLog) GetUserData() map[string]interface{} {
	if o == nil || IsNil(o.UserData) {
		var ret map[string]interface{}
		return ret
	}
	return o.UserData
}

// GetUserDataOk returns a tuple with the UserData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrPutLogLog) GetUserDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.UserData) {
		return map[string]interface{}{}, false
	}
	return o.UserData, true
}

// HasUserData returns a boolean if a field has been set.
func (o *CreateOrPutLogLog) HasUserData() bool {
	if o != nil && !IsNil(o.UserData) {
		return true
	}

	return false
}

// SetUserData gets a reference to the given map[string]interface{} and assigns it to the UserData field.
func (o *CreateOrPutLogLog) SetUserData(v map[string]interface{}) {
	o.UserData = v
}

// GetSourceType returns the SourceType field value if set, zero value otherwise.
func (o *CreateOrPutLogLog) GetSourceType() string {
	if o == nil || IsNil(o.SourceType) {
		var ret string
		return ret
	}
	return *o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrPutLogLog) GetSourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SourceType) {
		return nil, false
	}
	return o.SourceType, true
}

// HasSourceType returns a boolean if a field has been set.
func (o *CreateOrPutLogLog) HasSourceType() bool {
	if o != nil && !IsNil(o.SourceType) {
		return true
	}

	return false
}

// SetSourceType gets a reference to the given string and assigns it to the SourceType field.
func (o *CreateOrPutLogLog) SetSourceType(v string) {
	o.SourceType = &v
}

// GetTokenSeed returns the TokenSeed field value if set, zero value otherwise.
func (o *CreateOrPutLogLog) GetTokenSeed() string {
	if o == nil || IsNil(o.TokenSeed) {
		var ret string
		return ret
	}
	return *o.TokenSeed
}

// GetTokenSeedOk returns a tuple with the TokenSeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrPutLogLog) GetTokenSeedOk() (*string, bool) {
	if o == nil || IsNil(o.TokenSeed) {
		return nil, false
	}
	return o.TokenSeed, true
}

// HasTokenSeed returns a boolean if a field has been set.
func (o *CreateOrPutLogLog) HasTokenSeed() bool {
	if o != nil && !IsNil(o.TokenSeed) {
		return true
	}

	return false
}

// SetTokenSeed gets a reference to the given string and assigns it to the TokenSeed field.
func (o *CreateOrPutLogLog) SetTokenSeed(v string) {
	o.TokenSeed = &v
}

func (o CreateOrPutLogLog) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrPutLogLog) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Tokens) {
		toSerialize["tokens"] = o.Tokens
	}
	if !IsNil(o.Structures) {
		toSerialize["structures"] = o.Structures
	}
	if !IsNil(o.IpAddress) {
		toSerialize["ip_address"] = o.IpAddress
	}
	if !IsNil(o.LogsetsInfo) {
		toSerialize["logsets_info"] = o.LogsetsInfo
	}
	if !IsNil(o.UserData) {
		toSerialize["user_data"] = o.UserData
	}
	if !IsNil(o.SourceType) {
		toSerialize["source_type"] = o.SourceType
	}
	if !IsNil(o.TokenSeed) {
		toSerialize["token_seed"] = o.TokenSeed
	}
	return toSerialize, nil
}

func (o *CreateOrPutLogLog) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateOrPutLogLog := _CreateOrPutLogLog{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateOrPutLogLog)

	if err != nil {
		return err
	}

	*o = CreateOrPutLogLog(varCreateOrPutLogLog)

	return err
}

type NullableCreateOrPutLogLog struct {
	value *CreateOrPutLogLog
	isSet bool
}

func (v NullableCreateOrPutLogLog) Get() *CreateOrPutLogLog {
	return v.value
}

func (v *NullableCreateOrPutLogLog) Set(val *CreateOrPutLogLog) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrPutLogLog) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrPutLogLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrPutLogLog(val *CreateOrPutLogLog) *NullableCreateOrPutLogLog {
	return &NullableCreateOrPutLogLog{value: val, isSet: true}
}

func (v NullableCreateOrPutLogLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrPutLogLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}