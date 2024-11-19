# TargetAlertContentSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LeContext** | Pointer to **string** | When set to \&quot;true\&quot;, the log context is set to appear in the notification. When set to \&quot;false\&quot;, the log context is set to not appear in the notification. | [optional] 
**LeTriggerEvent** | Pointer to **[]string** | When set to \&quot;true\&quot;, the log entry which triggered the alert will appear in the notification. When set to \&quot;false\&quot;, the log entry which triggered the alert will not appear in the notification. | [optional] 

## Methods

### NewTargetAlertContentSet

`func NewTargetAlertContentSet() *TargetAlertContentSet`

NewTargetAlertContentSet instantiates a new TargetAlertContentSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetAlertContentSetWithDefaults

`func NewTargetAlertContentSetWithDefaults() *TargetAlertContentSet`

NewTargetAlertContentSetWithDefaults instantiates a new TargetAlertContentSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLeContext

`func (o *TargetAlertContentSet) GetLeContext() string`

GetLeContext returns the LeContext field if non-nil, zero value otherwise.

### GetLeContextOk

`func (o *TargetAlertContentSet) GetLeContextOk() (*string, bool)`

GetLeContextOk returns a tuple with the LeContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeContext

`func (o *TargetAlertContentSet) SetLeContext(v string)`

SetLeContext sets LeContext field to given value.

### HasLeContext

`func (o *TargetAlertContentSet) HasLeContext() bool`

HasLeContext returns a boolean if a field has been set.

### GetLeTriggerEvent

`func (o *TargetAlertContentSet) GetLeTriggerEvent() []string`

GetLeTriggerEvent returns the LeTriggerEvent field if non-nil, zero value otherwise.

### GetLeTriggerEventOk

`func (o *TargetAlertContentSet) GetLeTriggerEventOk() (*[]string, bool)`

GetLeTriggerEventOk returns a tuple with the LeTriggerEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeTriggerEvent

`func (o *TargetAlertContentSet) SetLeTriggerEvent(v []string)`

SetLeTriggerEvent sets LeTriggerEvent field to given value.

### HasLeTriggerEvent

`func (o *TargetAlertContentSet) HasLeTriggerEvent() bool`

HasLeTriggerEvent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


