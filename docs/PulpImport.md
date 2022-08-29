# PulpImport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | Path to export that will be imported. | [optional] 
**Toc** | Pointer to **string** | Path to a table-of-contents file describing chunks to be validated, reassembled, and imported. | [optional] 

## Methods

### NewPulpImport

`func NewPulpImport() *PulpImport`

NewPulpImport instantiates a new PulpImport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPulpImportWithDefaults

`func NewPulpImportWithDefaults() *PulpImport`

NewPulpImportWithDefaults instantiates a new PulpImport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *PulpImport) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PulpImport) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PulpImport) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *PulpImport) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetToc

`func (o *PulpImport) GetToc() string`

GetToc returns the Toc field if non-nil, zero value otherwise.

### GetTocOk

`func (o *PulpImport) GetTocOk() (*string, bool)`

GetTocOk returns a tuple with the Toc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToc

`func (o *PulpImport) SetToc(v string)`

SetToc sets Toc field to given value.

### HasToc

`func (o *PulpImport) HasToc() bool`

HasToc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


