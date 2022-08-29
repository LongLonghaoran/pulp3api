# SigningServiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**Name** | **string** | A unique name used to recognize a script. | 
**PublicKey** | **string** | The value of a public key used for the repository verification. | 
**PubkeyFingerprint** | **string** | The fingerprint of the public key. | 
**Script** | **string** | An absolute path to a script which is going to be used for the signing. | 

## Methods

### NewSigningServiceResponse

`func NewSigningServiceResponse(name string, publicKey string, pubkeyFingerprint string, script string, ) *SigningServiceResponse`

NewSigningServiceResponse instantiates a new SigningServiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSigningServiceResponseWithDefaults

`func NewSigningServiceResponseWithDefaults() *SigningServiceResponse`

NewSigningServiceResponseWithDefaults instantiates a new SigningServiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *SigningServiceResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *SigningServiceResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *SigningServiceResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *SigningServiceResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *SigningServiceResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *SigningServiceResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *SigningServiceResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *SigningServiceResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetName

`func (o *SigningServiceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SigningServiceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SigningServiceResponse) SetName(v string)`

SetName sets Name field to given value.


### GetPublicKey

`func (o *SigningServiceResponse) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *SigningServiceResponse) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *SigningServiceResponse) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### GetPubkeyFingerprint

`func (o *SigningServiceResponse) GetPubkeyFingerprint() string`

GetPubkeyFingerprint returns the PubkeyFingerprint field if non-nil, zero value otherwise.

### GetPubkeyFingerprintOk

`func (o *SigningServiceResponse) GetPubkeyFingerprintOk() (*string, bool)`

GetPubkeyFingerprintOk returns a tuple with the PubkeyFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubkeyFingerprint

`func (o *SigningServiceResponse) SetPubkeyFingerprint(v string)`

SetPubkeyFingerprint sets PubkeyFingerprint field to given value.


### GetScript

`func (o *SigningServiceResponse) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *SigningServiceResponse) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *SigningServiceResponse) SetScript(v string)`

SetScript sets Script field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


