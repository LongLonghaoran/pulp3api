# OrphansCleanup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentHrefs** | Pointer to **[]interface{}** | Will delete specified content and associated Artifacts if they are orphans. | [optional] 
**OrphanProtectionTime** | Pointer to **NullableInt32** | The time in minutes for how long Pulp will hold orphan Content and Artifacts before they become candidates for deletion by this orphan cleanup task. This should ideally be longer than your longest running task otherwise any content created during that task could be cleaned up before the task finishes. If not specified, a default value is taken from the setting ORPHAN_PROTECTION_TIME. | [optional] 

## Methods

### NewOrphansCleanup

`func NewOrphansCleanup() *OrphansCleanup`

NewOrphansCleanup instantiates a new OrphansCleanup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrphansCleanupWithDefaults

`func NewOrphansCleanupWithDefaults() *OrphansCleanup`

NewOrphansCleanupWithDefaults instantiates a new OrphansCleanup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentHrefs

`func (o *OrphansCleanup) GetContentHrefs() []interface{}`

GetContentHrefs returns the ContentHrefs field if non-nil, zero value otherwise.

### GetContentHrefsOk

`func (o *OrphansCleanup) GetContentHrefsOk() (*[]interface{}, bool)`

GetContentHrefsOk returns a tuple with the ContentHrefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentHrefs

`func (o *OrphansCleanup) SetContentHrefs(v []interface{})`

SetContentHrefs sets ContentHrefs field to given value.

### HasContentHrefs

`func (o *OrphansCleanup) HasContentHrefs() bool`

HasContentHrefs returns a boolean if a field has been set.

### GetOrphanProtectionTime

`func (o *OrphansCleanup) GetOrphanProtectionTime() int32`

GetOrphanProtectionTime returns the OrphanProtectionTime field if non-nil, zero value otherwise.

### GetOrphanProtectionTimeOk

`func (o *OrphansCleanup) GetOrphanProtectionTimeOk() (*int32, bool)`

GetOrphanProtectionTimeOk returns a tuple with the OrphanProtectionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrphanProtectionTime

`func (o *OrphansCleanup) SetOrphanProtectionTime(v int32)`

SetOrphanProtectionTime sets OrphanProtectionTime field to given value.

### HasOrphanProtectionTime

`func (o *OrphansCleanup) HasOrphanProtectionTime() bool`

HasOrphanProtectionTime returns a boolean if a field has been set.

### SetOrphanProtectionTimeNil

`func (o *OrphansCleanup) SetOrphanProtectionTimeNil(b bool)`

 SetOrphanProtectionTimeNil sets the value for OrphanProtectionTime to be an explicit nil

### UnsetOrphanProtectionTime
`func (o *OrphansCleanup) UnsetOrphanProtectionTime()`

UnsetOrphanProtectionTime ensures that no value is present for OrphanProtectionTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


