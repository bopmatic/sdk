# UploadPackageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjId** | Pointer to **string** |  | [optional] 
**PackageXsum** | Pointer to **string** |  | [optional] 
**PackageTarballData** | Pointer to **string** | package tarball content in .tar.xz format (limited to 6MiB); | [optional] 
**PackageTarballURL** | Pointer to **string** |  | [optional] 

## Methods

### NewUploadPackageRequest

`func NewUploadPackageRequest() *UploadPackageRequest`

NewUploadPackageRequest instantiates a new UploadPackageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadPackageRequestWithDefaults

`func NewUploadPackageRequestWithDefaults() *UploadPackageRequest`

NewUploadPackageRequestWithDefaults instantiates a new UploadPackageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjId

`func (o *UploadPackageRequest) GetProjId() string`

GetProjId returns the ProjId field if non-nil, zero value otherwise.

### GetProjIdOk

`func (o *UploadPackageRequest) GetProjIdOk() (*string, bool)`

GetProjIdOk returns a tuple with the ProjId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjId

`func (o *UploadPackageRequest) SetProjId(v string)`

SetProjId sets ProjId field to given value.

### HasProjId

`func (o *UploadPackageRequest) HasProjId() bool`

HasProjId returns a boolean if a field has been set.

### GetPackageXsum

`func (o *UploadPackageRequest) GetPackageXsum() string`

GetPackageXsum returns the PackageXsum field if non-nil, zero value otherwise.

### GetPackageXsumOk

`func (o *UploadPackageRequest) GetPackageXsumOk() (*string, bool)`

GetPackageXsumOk returns a tuple with the PackageXsum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageXsum

`func (o *UploadPackageRequest) SetPackageXsum(v string)`

SetPackageXsum sets PackageXsum field to given value.

### HasPackageXsum

`func (o *UploadPackageRequest) HasPackageXsum() bool`

HasPackageXsum returns a boolean if a field has been set.

### GetPackageTarballData

`func (o *UploadPackageRequest) GetPackageTarballData() string`

GetPackageTarballData returns the PackageTarballData field if non-nil, zero value otherwise.

### GetPackageTarballDataOk

`func (o *UploadPackageRequest) GetPackageTarballDataOk() (*string, bool)`

GetPackageTarballDataOk returns a tuple with the PackageTarballData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageTarballData

`func (o *UploadPackageRequest) SetPackageTarballData(v string)`

SetPackageTarballData sets PackageTarballData field to given value.

### HasPackageTarballData

`func (o *UploadPackageRequest) HasPackageTarballData() bool`

HasPackageTarballData returns a boolean if a field has been set.

### GetPackageTarballURL

`func (o *UploadPackageRequest) GetPackageTarballURL() string`

GetPackageTarballURL returns the PackageTarballURL field if non-nil, zero value otherwise.

### GetPackageTarballURLOk

`func (o *UploadPackageRequest) GetPackageTarballURLOk() (*string, bool)`

GetPackageTarballURLOk returns a tuple with the PackageTarballURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageTarballURL

`func (o *UploadPackageRequest) SetPackageTarballURL(v string)`

SetPackageTarballURL sets PackageTarballURL field to given value.

### HasPackageTarballURL

`func (o *UploadPackageRequest) HasPackageTarballURL() bool`

HasPackageTarballURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


