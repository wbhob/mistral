# ListDocumentOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**PaginationInfo**](PaginationInfo.md) |  | 
**Data** | [**[]DocumentOut**](DocumentOut.md) |  | 

## Methods

### NewListDocumentOut

`func NewListDocumentOut(pagination PaginationInfo, data []DocumentOut, ) *ListDocumentOut`

NewListDocumentOut instantiates a new ListDocumentOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDocumentOutWithDefaults

`func NewListDocumentOutWithDefaults() *ListDocumentOut`

NewListDocumentOutWithDefaults instantiates a new ListDocumentOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *ListDocumentOut) GetPagination() PaginationInfo`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListDocumentOut) GetPaginationOk() (*PaginationInfo, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListDocumentOut) SetPagination(v PaginationInfo)`

SetPagination sets Pagination field to given value.


### GetData

`func (o *ListDocumentOut) GetData() []DocumentOut`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListDocumentOut) GetDataOk() (*[]DocumentOut, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListDocumentOut) SetData(v []DocumentOut)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


