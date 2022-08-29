# Repair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerifyChecksums** | Pointer to **bool** | Will verify that the checksum of all stored files matches what saved in the database. Otherwise only the existence of the files will be checked. Enabled by default | [optional] [default to true]

## Methods

### NewRepair

`func NewRepair() *Repair`

NewRepair instantiates a new Repair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepairWithDefaults

`func NewRepairWithDefaults() *Repair`

NewRepairWithDefaults instantiates a new Repair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerifyChecksums

`func (o *Repair) GetVerifyChecksums() bool`

GetVerifyChecksums returns the VerifyChecksums field if non-nil, zero value otherwise.

### GetVerifyChecksumsOk

`func (o *Repair) GetVerifyChecksumsOk() (*bool, bool)`

GetVerifyChecksumsOk returns a tuple with the VerifyChecksums field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyChecksums

`func (o *Repair) SetVerifyChecksums(v bool)`

SetVerifyChecksums sets VerifyChecksums field to given value.

### HasVerifyChecksums

`func (o *Repair) HasVerifyChecksums() bool`

HasVerifyChecksums returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


