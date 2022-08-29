# DebPackageReleaseComponentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**Package** | **string** | Package that is contained in release_comonent. | 
**ReleaseComponent** | **string** | ReleaseComponent this package is contained in. | 

## Methods

### NewDebPackageReleaseComponentResponse

`func NewDebPackageReleaseComponentResponse(package_ string, releaseComponent string, ) *DebPackageReleaseComponentResponse`

NewDebPackageReleaseComponentResponse instantiates a new DebPackageReleaseComponentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebPackageReleaseComponentResponseWithDefaults

`func NewDebPackageReleaseComponentResponseWithDefaults() *DebPackageReleaseComponentResponse`

NewDebPackageReleaseComponentResponseWithDefaults instantiates a new DebPackageReleaseComponentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *DebPackageReleaseComponentResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *DebPackageReleaseComponentResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *DebPackageReleaseComponentResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *DebPackageReleaseComponentResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *DebPackageReleaseComponentResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *DebPackageReleaseComponentResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *DebPackageReleaseComponentResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *DebPackageReleaseComponentResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetPackage

`func (o *DebPackageReleaseComponentResponse) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *DebPackageReleaseComponentResponse) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *DebPackageReleaseComponentResponse) SetPackage(v string)`

SetPackage sets Package field to given value.


### GetReleaseComponent

`func (o *DebPackageReleaseComponentResponse) GetReleaseComponent() string`

GetReleaseComponent returns the ReleaseComponent field if non-nil, zero value otherwise.

### GetReleaseComponentOk

`func (o *DebPackageReleaseComponentResponse) GetReleaseComponentOk() (*string, bool)`

GetReleaseComponentOk returns a tuple with the ReleaseComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseComponent

`func (o *DebPackageReleaseComponentResponse) SetReleaseComponent(v string)`

SetReleaseComponent sets ReleaseComponent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


