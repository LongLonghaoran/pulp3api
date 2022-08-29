# RemoveSignaturesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignedWithKeyId** | **string** | key_id of the key the signatures were produced with | 

## Methods

### NewRemoveSignaturesResponse

`func NewRemoveSignaturesResponse(signedWithKeyId string, ) *RemoveSignaturesResponse`

NewRemoveSignaturesResponse instantiates a new RemoveSignaturesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveSignaturesResponseWithDefaults

`func NewRemoveSignaturesResponseWithDefaults() *RemoveSignaturesResponse`

NewRemoveSignaturesResponseWithDefaults instantiates a new RemoveSignaturesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignedWithKeyId

`func (o *RemoveSignaturesResponse) GetSignedWithKeyId() string`

GetSignedWithKeyId returns the SignedWithKeyId field if non-nil, zero value otherwise.

### GetSignedWithKeyIdOk

`func (o *RemoveSignaturesResponse) GetSignedWithKeyIdOk() (*string, bool)`

GetSignedWithKeyIdOk returns a tuple with the SignedWithKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedWithKeyId

`func (o *RemoveSignaturesResponse) SetSignedWithKeyId(v string)`

SetSignedWithKeyId sets SignedWithKeyId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


