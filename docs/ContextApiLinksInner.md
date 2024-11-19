# ContextApiLinksInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rel** | **string** | If equal to **Self**, then href contains a link to poll the query that&#39;s in progress. If equal to **Next**, then href contains a link to retrieve the next page of results. If equal to **Prev**, then href contains a link to retrieve the previous page of results.  | 
**Href** | **string** | A http link to which a &#x60;GET&#x60; request can be made. | 

## Methods

### NewContextApiLinksInner

`func NewContextApiLinksInner(rel string, href string, ) *ContextApiLinksInner`

NewContextApiLinksInner instantiates a new ContextApiLinksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextApiLinksInnerWithDefaults

`func NewContextApiLinksInnerWithDefaults() *ContextApiLinksInner`

NewContextApiLinksInnerWithDefaults instantiates a new ContextApiLinksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRel

`func (o *ContextApiLinksInner) GetRel() string`

GetRel returns the Rel field if non-nil, zero value otherwise.

### GetRelOk

`func (o *ContextApiLinksInner) GetRelOk() (*string, bool)`

GetRelOk returns a tuple with the Rel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRel

`func (o *ContextApiLinksInner) SetRel(v string)`

SetRel sets Rel field to given value.


### GetHref

`func (o *ContextApiLinksInner) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ContextApiLinksInner) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ContextApiLinksInner) SetHref(v string)`

SetHref sets Href field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


