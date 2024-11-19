# S3ArchivingSetupS3setup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique id of the S3 Archiving Setup. | 
**BucketName** | **string** | The name of the AWS S3 Bucket. | 
**Enabled** | **bool** | Boolean indicating whether the S3 setup is enabled. | 
**Compression** | **string** | The type of compression used for archiving. This is either \&quot;GZIP\&quot; or \&quot;BZIP2\&quot;. | 

## Methods

### NewS3ArchivingSetupS3setup

`func NewS3ArchivingSetupS3setup(id string, bucketName string, enabled bool, compression string, ) *S3ArchivingSetupS3setup`

NewS3ArchivingSetupS3setup instantiates a new S3ArchivingSetupS3setup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3ArchivingSetupS3setupWithDefaults

`func NewS3ArchivingSetupS3setupWithDefaults() *S3ArchivingSetupS3setup`

NewS3ArchivingSetupS3setupWithDefaults instantiates a new S3ArchivingSetupS3setup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *S3ArchivingSetupS3setup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *S3ArchivingSetupS3setup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *S3ArchivingSetupS3setup) SetId(v string)`

SetId sets Id field to given value.


### GetBucketName

`func (o *S3ArchivingSetupS3setup) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *S3ArchivingSetupS3setup) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *S3ArchivingSetupS3setup) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### GetEnabled

`func (o *S3ArchivingSetupS3setup) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *S3ArchivingSetupS3setup) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *S3ArchivingSetupS3setup) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetCompression

`func (o *S3ArchivingSetupS3setup) GetCompression() string`

GetCompression returns the Compression field if non-nil, zero value otherwise.

### GetCompressionOk

`func (o *S3ArchivingSetupS3setup) GetCompressionOk() (*string, bool)`

GetCompressionOk returns a tuple with the Compression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompression

`func (o *S3ArchivingSetupS3setup) SetCompression(v string)`

SetCompression sets Compression field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


