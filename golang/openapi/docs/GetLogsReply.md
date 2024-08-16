# GetLogsReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entries** | Pointer to [**[]GetLogsEntry**](GetLogsEntry.md) |  | [optional] 

## Methods

### NewGetLogsReply

`func NewGetLogsReply() *GetLogsReply`

NewGetLogsReply instantiates a new GetLogsReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLogsReplyWithDefaults

`func NewGetLogsReplyWithDefaults() *GetLogsReply`

NewGetLogsReplyWithDefaults instantiates a new GetLogsReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntries

`func (o *GetLogsReply) GetEntries() []GetLogsEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *GetLogsReply) GetEntriesOk() (*[]GetLogsEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *GetLogsReply) SetEntries(v []GetLogsEntry)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *GetLogsReply) HasEntries() bool`

HasEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

