# PulpImporterResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**Name** | **string** | Unique name of the Importer. | 
**RepoMapping** | Pointer to **map[string]string** | Mapping of repo names in an export file to the repo names in Pulp. For example, if the export has a repo named &#39;foo&#39; and the repo to import content into was &#39;bar&#39;, the mapping would be \&quot;{&#39;foo&#39;: &#39;bar&#39;}\&quot;. | [optional] 

## Methods

### NewPulpImporterResponse

`func NewPulpImporterResponse(name string, ) *PulpImporterResponse`

NewPulpImporterResponse instantiates a new PulpImporterResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPulpImporterResponseWithDefaults

`func NewPulpImporterResponseWithDefaults() *PulpImporterResponse`

NewPulpImporterResponseWithDefaults instantiates a new PulpImporterResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *PulpImporterResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *PulpImporterResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *PulpImporterResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *PulpImporterResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *PulpImporterResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *PulpImporterResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *PulpImporterResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *PulpImporterResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetName

`func (o *PulpImporterResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PulpImporterResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PulpImporterResponse) SetName(v string)`

SetName sets Name field to given value.


### GetRepoMapping

`func (o *PulpImporterResponse) GetRepoMapping() map[string]string`

GetRepoMapping returns the RepoMapping field if non-nil, zero value otherwise.

### GetRepoMappingOk

`func (o *PulpImporterResponse) GetRepoMappingOk() (*map[string]string, bool)`

GetRepoMappingOk returns a tuple with the RepoMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoMapping

`func (o *PulpImporterResponse) SetRepoMapping(v map[string]string)`

SetRepoMapping sets RepoMapping field to given value.

### HasRepoMapping

`func (o *PulpImporterResponse) HasRepoMapping() bool`

HasRepoMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


