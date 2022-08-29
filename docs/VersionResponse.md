# VersionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Component** | **string** | Name of a versioned component of Pulp | 
**Version** | **string** | Version of the component (e.g. 3.0.0) | 
**Package** | **string** | Python package name providing the component | 

## Methods

### NewVersionResponse

`func NewVersionResponse(component string, version string, package_ string, ) *VersionResponse`

NewVersionResponse instantiates a new VersionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionResponseWithDefaults

`func NewVersionResponseWithDefaults() *VersionResponse`

NewVersionResponseWithDefaults instantiates a new VersionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponent

`func (o *VersionResponse) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *VersionResponse) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *VersionResponse) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetVersion

`func (o *VersionResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VersionResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VersionResponse) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetPackage

`func (o *VersionResponse) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *VersionResponse) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *VersionResponse) SetPackage(v string)`

SetPackage sets Package field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


