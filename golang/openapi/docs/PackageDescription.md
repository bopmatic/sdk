# PackageDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageId** | Pointer to **string** |  | [optional] 
**ProjId** | Pointer to **string** |  | [optional] 
**State** | Pointer to [**PackageState**](PackageState.md) |  | [optional] [default to UPLOADING]
**UploadTime** | Pointer to **string** |  | [optional] 
**PackageSize** | Pointer to **string** |  | [optional] 

## Methods

### NewPackageDescription

`func NewPackageDescription() *PackageDescription`

NewPackageDescription instantiates a new PackageDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageDescriptionWithDefaults

`func NewPackageDescriptionWithDefaults() *PackageDescription`

NewPackageDescriptionWithDefaults instantiates a new PackageDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageId

`func (o *PackageDescription) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *PackageDescription) GetPackageIdOk() (*string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageId

`func (o *PackageDescription) SetPackageId(v string)`

SetPackageId sets PackageId field to given value.

### HasPackageId

`func (o *PackageDescription) HasPackageId() bool`

HasPackageId returns a boolean if a field has been set.

### GetProjId

`func (o *PackageDescription) GetProjId() string`

GetProjId returns the ProjId field if non-nil, zero value otherwise.

### GetProjIdOk

`func (o *PackageDescription) GetProjIdOk() (*string, bool)`

GetProjIdOk returns a tuple with the ProjId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjId

`func (o *PackageDescription) SetProjId(v string)`

SetProjId sets ProjId field to given value.

### HasProjId

`func (o *PackageDescription) HasProjId() bool`

HasProjId returns a boolean if a field has been set.

### GetState

`func (o *PackageDescription) GetState() PackageState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PackageDescription) GetStateOk() (*PackageState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PackageDescription) SetState(v PackageState)`

SetState sets State field to given value.

### HasState

`func (o *PackageDescription) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUploadTime

`func (o *PackageDescription) GetUploadTime() string`

GetUploadTime returns the UploadTime field if non-nil, zero value otherwise.

### GetUploadTimeOk

`func (o *PackageDescription) GetUploadTimeOk() (*string, bool)`

GetUploadTimeOk returns a tuple with the UploadTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadTime

`func (o *PackageDescription) SetUploadTime(v string)`

SetUploadTime sets UploadTime field to given value.

### HasUploadTime

`func (o *PackageDescription) HasUploadTime() bool`

HasUploadTime returns a boolean if a field has been set.

### GetPackageSize

`func (o *PackageDescription) GetPackageSize() string`

GetPackageSize returns the PackageSize field if non-nil, zero value otherwise.

### GetPackageSizeOk

`func (o *PackageDescription) GetPackageSizeOk() (*string, bool)`

GetPackageSizeOk returns a tuple with the PackageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageSize

`func (o *PackageDescription) SetPackageSize(v string)`

SetPackageSize sets PackageSize field to given value.

### HasPackageSize

`func (o *PackageDescription) HasPackageSize() bool`

HasPackageSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


