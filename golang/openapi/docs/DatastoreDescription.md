# DatastoreDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatastoreHeader** | Pointer to [**DatastoreHeader**](DatastoreHeader.md) |  | [optional] 
**NumObjects** | Pointer to **string** |  | [optional] 
**CapacityConsumedInBytes** | Pointer to **string** |  | [optional] 

## Methods

### NewDatastoreDescription

`func NewDatastoreDescription() *DatastoreDescription`

NewDatastoreDescription instantiates a new DatastoreDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatastoreDescriptionWithDefaults

`func NewDatastoreDescriptionWithDefaults() *DatastoreDescription`

NewDatastoreDescriptionWithDefaults instantiates a new DatastoreDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatastoreHeader

`func (o *DatastoreDescription) GetDatastoreHeader() DatastoreHeader`

GetDatastoreHeader returns the DatastoreHeader field if non-nil, zero value otherwise.

### GetDatastoreHeaderOk

`func (o *DatastoreDescription) GetDatastoreHeaderOk() (*DatastoreHeader, bool)`

GetDatastoreHeaderOk returns a tuple with the DatastoreHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreHeader

`func (o *DatastoreDescription) SetDatastoreHeader(v DatastoreHeader)`

SetDatastoreHeader sets DatastoreHeader field to given value.

### HasDatastoreHeader

`func (o *DatastoreDescription) HasDatastoreHeader() bool`

HasDatastoreHeader returns a boolean if a field has been set.

### GetNumObjects

`func (o *DatastoreDescription) GetNumObjects() string`

GetNumObjects returns the NumObjects field if non-nil, zero value otherwise.

### GetNumObjectsOk

`func (o *DatastoreDescription) GetNumObjectsOk() (*string, bool)`

GetNumObjectsOk returns a tuple with the NumObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumObjects

`func (o *DatastoreDescription) SetNumObjects(v string)`

SetNumObjects sets NumObjects field to given value.

### HasNumObjects

`func (o *DatastoreDescription) HasNumObjects() bool`

HasNumObjects returns a boolean if a field has been set.

### GetCapacityConsumedInBytes

`func (o *DatastoreDescription) GetCapacityConsumedInBytes() string`

GetCapacityConsumedInBytes returns the CapacityConsumedInBytes field if non-nil, zero value otherwise.

### GetCapacityConsumedInBytesOk

`func (o *DatastoreDescription) GetCapacityConsumedInBytesOk() (*string, bool)`

GetCapacityConsumedInBytesOk returns a tuple with the CapacityConsumedInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityConsumedInBytes

`func (o *DatastoreDescription) SetCapacityConsumedInBytes(v string)`

SetCapacityConsumedInBytes sets CapacityConsumedInBytes field to given value.

### HasCapacityConsumedInBytes

`func (o *DatastoreDescription) HasCapacityConsumedInBytes() bool`

HasCapacityConsumedInBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


