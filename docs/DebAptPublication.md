# DebAptPublication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RepositoryVersion** | Pointer to **string** |  | [optional] 
**Repository** | Pointer to **string** | A URI of the repository to be published. | [optional] 
**Simple** | Pointer to **bool** | Activate simple publishing mode (all packages in one release component). | [optional] [default to false]
**Structured** | Pointer to **bool** | Activate structured publishing mode. | [optional] [default to false]
**SigningService** | Pointer to **string** | Sign Release files with this signing key | [optional] 

## Methods

### NewDebAptPublication

`func NewDebAptPublication() *DebAptPublication`

NewDebAptPublication instantiates a new DebAptPublication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebAptPublicationWithDefaults

`func NewDebAptPublicationWithDefaults() *DebAptPublication`

NewDebAptPublicationWithDefaults instantiates a new DebAptPublication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepositoryVersion

`func (o *DebAptPublication) GetRepositoryVersion() string`

GetRepositoryVersion returns the RepositoryVersion field if non-nil, zero value otherwise.

### GetRepositoryVersionOk

`func (o *DebAptPublication) GetRepositoryVersionOk() (*string, bool)`

GetRepositoryVersionOk returns a tuple with the RepositoryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryVersion

`func (o *DebAptPublication) SetRepositoryVersion(v string)`

SetRepositoryVersion sets RepositoryVersion field to given value.

### HasRepositoryVersion

`func (o *DebAptPublication) HasRepositoryVersion() bool`

HasRepositoryVersion returns a boolean if a field has been set.

### GetRepository

`func (o *DebAptPublication) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *DebAptPublication) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *DebAptPublication) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *DebAptPublication) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetSimple

`func (o *DebAptPublication) GetSimple() bool`

GetSimple returns the Simple field if non-nil, zero value otherwise.

### GetSimpleOk

`func (o *DebAptPublication) GetSimpleOk() (*bool, bool)`

GetSimpleOk returns a tuple with the Simple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimple

`func (o *DebAptPublication) SetSimple(v bool)`

SetSimple sets Simple field to given value.

### HasSimple

`func (o *DebAptPublication) HasSimple() bool`

HasSimple returns a boolean if a field has been set.

### GetStructured

`func (o *DebAptPublication) GetStructured() bool`

GetStructured returns the Structured field if non-nil, zero value otherwise.

### GetStructuredOk

`func (o *DebAptPublication) GetStructuredOk() (*bool, bool)`

GetStructuredOk returns a tuple with the Structured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructured

`func (o *DebAptPublication) SetStructured(v bool)`

SetStructured sets Structured field to given value.

### HasStructured

`func (o *DebAptPublication) HasStructured() bool`

HasStructured returns a boolean if a field has been set.

### GetSigningService

`func (o *DebAptPublication) GetSigningService() string`

GetSigningService returns the SigningService field if non-nil, zero value otherwise.

### GetSigningServiceOk

`func (o *DebAptPublication) GetSigningServiceOk() (*string, bool)`

GetSigningServiceOk returns a tuple with the SigningService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningService

`func (o *DebAptPublication) SetSigningService(v string)`

SetSigningService sets SigningService field to given value.

### HasSigningService

`func (o *DebAptPublication) HasSigningService() bool`

HasSigningService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


