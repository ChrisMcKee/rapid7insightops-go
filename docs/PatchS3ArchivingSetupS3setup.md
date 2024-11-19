# PatchS3ArchivingSetupS3setup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | Pointer to **string** | The name of the S3 Bucket. | [optional] 
**Enabled** | Pointer to **bool** | Boolean indicating whether the S3 setup is enabled. | [optional] 
**Compression** | Pointer to **string** | The type of compression used for archiving. This is either \&quot;GZIP\&quot; or \&quot;BZIP2\&quot;. | [optional] 

## Methods

### NewPatchS3ArchivingSetupS3setup

`func NewPatchS3ArchivingSetupS3setup() *PatchS3ArchivingSetupS3setup`

NewPatchS3ArchivingSetupS3setup instantiates a new PatchS3ArchivingSetupS3setup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchS3ArchivingSetupS3setupWithDefaults

`func NewPatchS3ArchivingSetupS3setupWithDefaults() *PatchS3ArchivingSetupS3setup`

NewPatchS3ArchivingSetupS3setupWithDefaults instantiates a new PatchS3ArchivingSetupS3setup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *PatchS3ArchivingSetupS3setup) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *PatchS3ArchivingSetupS3setup) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *PatchS3ArchivingSetupS3setup) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *PatchS3ArchivingSetupS3setup) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchS3ArchivingSetupS3setup) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchS3ArchivingSetupS3setup) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchS3ArchivingSetupS3setup) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchS3ArchivingSetupS3setup) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetCompression

`func (o *PatchS3ArchivingSetupS3setup) GetCompression() string`

GetCompression returns the Compression field if non-nil, zero value otherwise.

### GetCompressionOk

`func (o *PatchS3ArchivingSetupS3setup) GetCompressionOk() (*string, bool)`

GetCompressionOk returns a tuple with the Compression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompression

`func (o *PatchS3ArchivingSetupS3setup) SetCompression(v string)`

SetCompression sets Compression field to given value.

### HasCompression

`func (o *PatchS3ArchivingSetupS3setup) HasCompression() bool`

HasCompression returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


