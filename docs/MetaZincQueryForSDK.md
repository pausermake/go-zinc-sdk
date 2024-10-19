# MetaZincQueryForSDK

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | Pointer to **[]string** | true, false, [\&quot;field1\&quot;, \&quot;field2.*\&quot;] | [optional] 
**Aggs** | Pointer to [**map[string]MetaAggregations**](MetaAggregations.md) |  | [optional] 
**Explain** | Pointer to **bool** |  | [optional] 
**Fields** | Pointer to **[]string** | [\&quot;field1\&quot;, \&quot;field2.*\&quot;, {\&quot;field\&quot;: \&quot;fieldName\&quot;, \&quot;format\&quot;: \&quot;epoch_millis\&quot;}] | [optional] 
**From** | Pointer to **int32** |  | [optional] 
**Highlight** | Pointer to [**MetaHighlight**](MetaHighlight.md) |  | [optional] 
**Query** | Pointer to [**MetaQueryForSDK**](MetaQueryForSDK.md) |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Sort** | Pointer to **[]string** | \&quot;_score\&quot;, [\&quot;+Year\&quot;,\&quot;-Year\&quot;, {\&quot;Year\&quot;: \&quot;desc\&quot;}, \&quot;Date\&quot;: {\&quot;order\&quot;: \&quot;asc\&quot;\&quot;, \&quot;format\&quot;: \&quot;yyyy-MM-dd\&quot;}}\&quot;}] | [optional] 
**Timeout** | Pointer to **int32** |  | [optional] 
**TrackTotalHits** | Pointer to **bool** |  | [optional] 

## Methods

### NewMetaZincQueryForSDK

`func NewMetaZincQueryForSDK() *MetaZincQueryForSDK`

NewMetaZincQueryForSDK instantiates a new MetaZincQueryForSDK object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaZincQueryForSDKWithDefaults

`func NewMetaZincQueryForSDKWithDefaults() *MetaZincQueryForSDK`

NewMetaZincQueryForSDKWithDefaults instantiates a new MetaZincQueryForSDK object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *MetaZincQueryForSDK) GetSource() []string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *MetaZincQueryForSDK) GetSourceOk() (*[]string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *MetaZincQueryForSDK) SetSource(v []string)`

SetSource sets Source field to given value.

### HasSource

`func (o *MetaZincQueryForSDK) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetAggs

`func (o *MetaZincQueryForSDK) GetAggs() map[string]MetaAggregations`

GetAggs returns the Aggs field if non-nil, zero value otherwise.

### GetAggsOk

`func (o *MetaZincQueryForSDK) GetAggsOk() (*map[string]MetaAggregations, bool)`

GetAggsOk returns a tuple with the Aggs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggs

`func (o *MetaZincQueryForSDK) SetAggs(v map[string]MetaAggregations)`

SetAggs sets Aggs field to given value.

### HasAggs

`func (o *MetaZincQueryForSDK) HasAggs() bool`

HasAggs returns a boolean if a field has been set.

### GetExplain

`func (o *MetaZincQueryForSDK) GetExplain() bool`

GetExplain returns the Explain field if non-nil, zero value otherwise.

### GetExplainOk

`func (o *MetaZincQueryForSDK) GetExplainOk() (*bool, bool)`

GetExplainOk returns a tuple with the Explain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplain

`func (o *MetaZincQueryForSDK) SetExplain(v bool)`

SetExplain sets Explain field to given value.

### HasExplain

`func (o *MetaZincQueryForSDK) HasExplain() bool`

HasExplain returns a boolean if a field has been set.

### GetFields

`func (o *MetaZincQueryForSDK) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *MetaZincQueryForSDK) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *MetaZincQueryForSDK) SetFields(v []string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *MetaZincQueryForSDK) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFrom

`func (o *MetaZincQueryForSDK) GetFrom() int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *MetaZincQueryForSDK) GetFromOk() (*int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *MetaZincQueryForSDK) SetFrom(v int32)`

SetFrom sets From field to given value.

### HasFrom

`func (o *MetaZincQueryForSDK) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetHighlight

`func (o *MetaZincQueryForSDK) GetHighlight() MetaHighlight`

GetHighlight returns the Highlight field if non-nil, zero value otherwise.

### GetHighlightOk

`func (o *MetaZincQueryForSDK) GetHighlightOk() (*MetaHighlight, bool)`

GetHighlightOk returns a tuple with the Highlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlight

`func (o *MetaZincQueryForSDK) SetHighlight(v MetaHighlight)`

SetHighlight sets Highlight field to given value.

### HasHighlight

`func (o *MetaZincQueryForSDK) HasHighlight() bool`

HasHighlight returns a boolean if a field has been set.

### GetQuery

`func (o *MetaZincQueryForSDK) GetQuery() MetaQueryForSDK`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *MetaZincQueryForSDK) GetQueryOk() (*MetaQueryForSDK, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *MetaZincQueryForSDK) SetQuery(v MetaQueryForSDK)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *MetaZincQueryForSDK) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetSize

`func (o *MetaZincQueryForSDK) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MetaZincQueryForSDK) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MetaZincQueryForSDK) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *MetaZincQueryForSDK) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSort

`func (o *MetaZincQueryForSDK) GetSort() []string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *MetaZincQueryForSDK) GetSortOk() (*[]string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *MetaZincQueryForSDK) SetSort(v []string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *MetaZincQueryForSDK) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetTimeout

`func (o *MetaZincQueryForSDK) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *MetaZincQueryForSDK) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *MetaZincQueryForSDK) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *MetaZincQueryForSDK) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetTrackTotalHits

`func (o *MetaZincQueryForSDK) GetTrackTotalHits() bool`

GetTrackTotalHits returns the TrackTotalHits field if non-nil, zero value otherwise.

### GetTrackTotalHitsOk

`func (o *MetaZincQueryForSDK) GetTrackTotalHitsOk() (*bool, bool)`

GetTrackTotalHitsOk returns a tuple with the TrackTotalHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackTotalHits

`func (o *MetaZincQueryForSDK) SetTrackTotalHits(v bool)`

SetTrackTotalHits sets TrackTotalHits field to given value.

### HasTrackTotalHits

`func (o *MetaZincQueryForSDK) HasTrackTotalHits() bool`

HasTrackTotalHits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


