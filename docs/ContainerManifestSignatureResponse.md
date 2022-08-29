# ContainerManifestSignatureResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**Name** | **string** | Signature name in the format of &#x60;digest_algo:manifest_digest@random_32_chars&#x60; | 
**Digest** | **string** | sha256 digest of the signature blob | 
**Type** | **string** | Container signature type, e.g. &#39;atomic&#39; | 
**KeyId** | **string** | Signing key ID | 
**Timestamp** | **int32** | Timestamp of a signature | 
**Creator** | **string** | Signature creator | 
**SignedManifest** | **string** | Manifest that is signed | 

## Methods

### NewContainerManifestSignatureResponse

`func NewContainerManifestSignatureResponse(name string, digest string, type_ string, keyId string, timestamp int32, creator string, signedManifest string, ) *ContainerManifestSignatureResponse`

NewContainerManifestSignatureResponse instantiates a new ContainerManifestSignatureResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerManifestSignatureResponseWithDefaults

`func NewContainerManifestSignatureResponseWithDefaults() *ContainerManifestSignatureResponse`

NewContainerManifestSignatureResponseWithDefaults instantiates a new ContainerManifestSignatureResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *ContainerManifestSignatureResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *ContainerManifestSignatureResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *ContainerManifestSignatureResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *ContainerManifestSignatureResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *ContainerManifestSignatureResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *ContainerManifestSignatureResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *ContainerManifestSignatureResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *ContainerManifestSignatureResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetName

`func (o *ContainerManifestSignatureResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContainerManifestSignatureResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContainerManifestSignatureResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDigest

`func (o *ContainerManifestSignatureResponse) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *ContainerManifestSignatureResponse) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *ContainerManifestSignatureResponse) SetDigest(v string)`

SetDigest sets Digest field to given value.


### GetType

`func (o *ContainerManifestSignatureResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContainerManifestSignatureResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContainerManifestSignatureResponse) SetType(v string)`

SetType sets Type field to given value.


### GetKeyId

`func (o *ContainerManifestSignatureResponse) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *ContainerManifestSignatureResponse) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *ContainerManifestSignatureResponse) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.


### GetTimestamp

`func (o *ContainerManifestSignatureResponse) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ContainerManifestSignatureResponse) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ContainerManifestSignatureResponse) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetCreator

`func (o *ContainerManifestSignatureResponse) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *ContainerManifestSignatureResponse) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *ContainerManifestSignatureResponse) SetCreator(v string)`

SetCreator sets Creator field to given value.


### GetSignedManifest

`func (o *ContainerManifestSignatureResponse) GetSignedManifest() string`

GetSignedManifest returns the SignedManifest field if non-nil, zero value otherwise.

### GetSignedManifestOk

`func (o *ContainerManifestSignatureResponse) GetSignedManifestOk() (*string, bool)`

GetSignedManifestOk returns a tuple with the SignedManifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedManifest

`func (o *ContainerManifestSignatureResponse) SetSignedManifest(v string)`

SetSignedManifest sets SignedManifest field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


