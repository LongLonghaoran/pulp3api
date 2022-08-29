# DebReleaseArchitecture

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Architecture** | **string** | Name of the architecture. | 
**Release** | **string** | Release this architecture is contained in. | 

## Methods

### NewDebReleaseArchitecture

`func NewDebReleaseArchitecture(architecture string, release string, ) *DebReleaseArchitecture`

NewDebReleaseArchitecture instantiates a new DebReleaseArchitecture object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebReleaseArchitectureWithDefaults

`func NewDebReleaseArchitectureWithDefaults() *DebReleaseArchitecture`

NewDebReleaseArchitectureWithDefaults instantiates a new DebReleaseArchitecture object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchitecture

`func (o *DebReleaseArchitecture) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *DebReleaseArchitecture) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *DebReleaseArchitecture) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.


### GetRelease

`func (o *DebReleaseArchitecture) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *DebReleaseArchitecture) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *DebReleaseArchitecture) SetRelease(v string)`

SetRelease sets Release field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


