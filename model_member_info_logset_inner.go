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

// checks if the MemberInfoLogsetInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MemberInfoLogsetInner{}

// MemberInfoLogsetInner struct for MemberInfoLogsetInner
type MemberInfoLogsetInner struct {
	// The UUID of the log.
	Id string `json:"id"`
	// The name of the log.
	Name string `json:"name"`
	// The Rapid7 Resource Name (RRN) of the log. The RRN is a unique identifier across the Rapid7 Platform.
	Rrn   *string      `json:"rrn,omitempty"`
	Links []LinksInner `json:"links,omitempty"`
}

type _MemberInfoLogsetInner MemberInfoLogsetInner

// NewMemberInfoLogsetInner instantiates a new MemberInfoLogsetInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemberInfoLogsetInner(id string, name string) *MemberInfoLogsetInner {
	this := MemberInfoLogsetInner{}
	this.Id = id
	this.Name = name
	return &this
}

// NewMemberInfoLogsetInnerWithDefaults instantiates a new MemberInfoLogsetInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemberInfoLogsetInnerWithDefaults() *MemberInfoLogsetInner {
	this := MemberInfoLogsetInner{}
	return &this
}

// GetId returns the Id field value
func (o *MemberInfoLogsetInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MemberInfoLogsetInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MemberInfoLogsetInner) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *MemberInfoLogsetInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MemberInfoLogsetInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MemberInfoLogsetInner) SetName(v string) {
	o.Name = v
}

// GetRrn returns the Rrn field value if set, zero value otherwise.
func (o *MemberInfoLogsetInner) GetRrn() string {
	if o == nil || IsNil(o.Rrn) {
		var ret string
		return ret
	}
	return *o.Rrn
}

// GetRrnOk returns a tuple with the Rrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberInfoLogsetInner) GetRrnOk() (*string, bool) {
	if o == nil || IsNil(o.Rrn) {
		return nil, false
	}
	return o.Rrn, true
}

// HasRrn returns a boolean if a field has been set.
func (o *MemberInfoLogsetInner) HasRrn() bool {
	if o != nil && !IsNil(o.Rrn) {
		return true
	}

	return false
}

// SetRrn gets a reference to the given string and assigns it to the Rrn field.
func (o *MemberInfoLogsetInner) SetRrn(v string) {
	o.Rrn = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *MemberInfoLogsetInner) GetLinks() []LinksInner {
	if o == nil || IsNil(o.Links) {
		var ret []LinksInner
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberInfoLogsetInner) GetLinksOk() ([]LinksInner, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *MemberInfoLogsetInner) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []LinksInner and assigns it to the Links field.
func (o *MemberInfoLogsetInner) SetLinks(v []LinksInner) {
	o.Links = v
}

func (o MemberInfoLogsetInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MemberInfoLogsetInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.Rrn) {
		toSerialize["rrn"] = o.Rrn
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

func (o *MemberInfoLogsetInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varMemberInfoLogsetInner := _MemberInfoLogsetInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMemberInfoLogsetInner)

	if err != nil {
		return err
	}

	*o = MemberInfoLogsetInner(varMemberInfoLogsetInner)

	return err
}

type NullableMemberInfoLogsetInner struct {
	value *MemberInfoLogsetInner
	isSet bool
}

func (v NullableMemberInfoLogsetInner) Get() *MemberInfoLogsetInner {
	return v.value
}

func (v *NullableMemberInfoLogsetInner) Set(val *MemberInfoLogsetInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMemberInfoLogsetInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMemberInfoLogsetInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemberInfoLogsetInner(val *MemberInfoLogsetInner) *NullableMemberInfoLogsetInner {
	return &NullableMemberInfoLogsetInner{value: val, isSet: true}
}

func (v NullableMemberInfoLogsetInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemberInfoLogsetInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
