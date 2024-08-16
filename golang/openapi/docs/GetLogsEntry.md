# GetLogsEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** | log message | [optional] 

## Methods

### NewGetLogsEntry

`func NewGetLogsEntry() *GetLogsEntry`

NewGetLogsEntry instantiates a new GetLogsEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLogsEntryWithDefaults

`func NewGetLogsEntryWithDefaults() *GetLogsEntry`

NewGetLogsEntryWithDefaults instantiates a new GetLogsEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *GetLogsEntry) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetLogsEntry) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetLogsEntry) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *GetLogsEntry) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetMessage

`func (o *GetLogsEntry) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetLogsEntry) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetLogsEntry) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetLogsEntry) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


