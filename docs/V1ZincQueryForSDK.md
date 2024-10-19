# V1ZincQueryForSDK

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | Pointer to **[]string** |  | [optional] 
**Aggs** | Pointer to [**map[string]V1AggregationParams**](V1AggregationParams.md) |  | [optional] 
**Explain** | Pointer to **bool** |  | [optional] 
**From** | Pointer to **int32** |  | [optional] 
**Highlight** | Pointer to [**MetaHighlight**](MetaHighlight.md) |  | [optional] 
**MaxResults** | Pointer to **int32** |  | [optional] 
**Query** | Pointer to [**V1QueryParams**](V1QueryParams.md) |  | [optional] 
**SearchType** | Pointer to **string** | SearchType is the type of search to perform. Can be match, match_phrase, query_string, etc | [optional] 
**SortFields** | Pointer to **[]string** |  | [optional] 

## Methods

### NewV1ZincQueryForSDK

`func NewV1ZincQueryForSDK() *V1ZincQueryForSDK`

NewV1ZincQueryForSDK instantiates a new V1ZincQueryForSDK object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ZincQueryForSDKWithDefaults

`func NewV1ZincQueryForSDKWithDefaults() *V1ZincQueryForSDK`

NewV1ZincQueryForSDKWithDefaults instantiates a new V1ZincQueryForSDK object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *V1ZincQueryForSDK) GetSource() []string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *V1ZincQueryForSDK) GetSourceOk() (*[]string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *V1ZincQueryForSDK) SetSource(v []string)`

SetSource sets Source field to given value.

### HasSource

`func (o *V1ZincQueryForSDK) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetAggs

`func (o *V1ZincQueryForSDK) GetAggs() map[string]V1AggregationParams`

GetAggs returns the Aggs field if non-nil, zero value otherwise.

### GetAggsOk

`func (o *V1ZincQueryForSDK) GetAggsOk() (*map[string]V1AggregationParams, bool)`

GetAggsOk returns a tuple with the Aggs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggs

`func (o *V1ZincQueryForSDK) SetAggs(v map[string]V1AggregationParams)`

SetAggs sets Aggs field to given value.

### HasAggs

`func (o *V1ZincQueryForSDK) HasAggs() bool`

HasAggs returns a boolean if a field has been set.

### GetExplain

`func (o *V1ZincQueryForSDK) GetExplain() bool`

GetExplain returns the Explain field if non-nil, zero value otherwise.

### GetExplainOk

`func (o *V1ZincQueryForSDK) GetExplainOk() (*bool, bool)`

GetExplainOk returns a tuple with the Explain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplain

`func (o *V1ZincQueryForSDK) SetExplain(v bool)`

SetExplain sets Explain field to given value.

### HasExplain

`func (o *V1ZincQueryForSDK) HasExplain() bool`

HasExplain returns a boolean if a field has been set.

### GetFrom

`func (o *V1ZincQueryForSDK) GetFrom() int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *V1ZincQueryForSDK) GetFromOk() (*int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *V1ZincQueryForSDK) SetFrom(v int32)`

SetFrom sets From field to given value.

### HasFrom

`func (o *V1ZincQueryForSDK) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetHighlight

`func (o *V1ZincQueryForSDK) GetHighlight() MetaHighlight`

GetHighlight returns the Highlight field if non-nil, zero value otherwise.

### GetHighlightOk

`func (o *V1ZincQueryForSDK) GetHighlightOk() (*MetaHighlight, bool)`

GetHighlightOk returns a tuple with the Highlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlight

`func (o *V1ZincQueryForSDK) SetHighlight(v MetaHighlight)`

SetHighlight sets Highlight field to given value.

### HasHighlight

`func (o *V1ZincQueryForSDK) HasHighlight() bool`

HasHighlight returns a boolean if a field has been set.

### GetMaxResults

`func (o *V1ZincQueryForSDK) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *V1ZincQueryForSDK) GetMaxResultsOk() (*int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *V1ZincQueryForSDK) SetMaxResults(v int32)`

SetMaxResults sets MaxResults field to given value.

### HasMaxResults

`func (o *V1ZincQueryForSDK) HasMaxResults() bool`

HasMaxResults returns a boolean if a field has been set.

### GetQuery

`func (o *V1ZincQueryForSDK) GetQuery() V1QueryParams`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *V1ZincQueryForSDK) GetQueryOk() (*V1QueryParams, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *V1ZincQueryForSDK) SetQuery(v V1QueryParams)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *V1ZincQueryForSDK) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetSearchType

`func (o *V1ZincQueryForSDK) GetSearchType() string`

GetSearchType returns the SearchType field if non-nil, zero value otherwise.

### GetSearchTypeOk

`func (o *V1ZincQueryForSDK) GetSearchTypeOk() (*string, bool)`

GetSearchTypeOk returns a tuple with the SearchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchType

`func (o *V1ZincQueryForSDK) SetSearchType(v string)`

SetSearchType sets SearchType field to given value.

### HasSearchType

`func (o *V1ZincQueryForSDK) HasSearchType() bool`

HasSearchType returns a boolean if a field has been set.

### GetSortFields

`func (o *V1ZincQueryForSDK) GetSortFields() []string`

GetSortFields returns the SortFields field if non-nil, zero value otherwise.

### GetSortFieldsOk

`func (o *V1ZincQueryForSDK) GetSortFieldsOk() (*[]string, bool)`

GetSortFieldsOk returns a tuple with the SortFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortFields

`func (o *V1ZincQueryForSDK) SetSortFields(v []string)`

SetSortFields sets SortFields field to given value.

### HasSortFields

`func (o *V1ZincQueryForSDK) HasSortFields() bool`

HasSortFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


