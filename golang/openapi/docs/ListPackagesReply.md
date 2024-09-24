# ListPackagesReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**ServiceRunnerResult**](ServiceRunnerResult.md) |  | [optional] 
**Items** | Pointer to [**[]ListPackagesReplyListPackagesItem**](ListPackagesReplyListPackagesItem.md) |  | [optional] 

## Methods

### NewListPackagesReply

`func NewListPackagesReply() *ListPackagesReply`

NewListPackagesReply instantiates a new ListPackagesReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPackagesReplyWithDefaults

`func NewListPackagesReplyWithDefaults() *ListPackagesReply`

NewListPackagesReplyWithDefaults instantiates a new ListPackagesReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ListPackagesReply) GetResult() ServiceRunnerResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ListPackagesReply) GetResultOk() (*ServiceRunnerResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ListPackagesReply) SetResult(v ServiceRunnerResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *ListPackagesReply) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetItems

`func (o *ListPackagesReply) GetItems() []ListPackagesReplyListPackagesItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListPackagesReply) GetItemsOk() (*[]ListPackagesReplyListPackagesItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListPackagesReply) SetItems(v []ListPackagesReplyListPackagesItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *ListPackagesReply) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


