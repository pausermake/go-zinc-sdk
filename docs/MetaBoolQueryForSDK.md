# MetaBoolQueryForSDK

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**[]MetaQueryForSDK**](MetaQueryForSDK.md) | query, [query1, query2] | [optional] 
**MinimumShouldMatch** | Pointer to **float32** | only for should | [optional] 
**Must** | Pointer to [**[]MetaQueryForSDK**](MetaQueryForSDK.md) | query, [query1, query2] | [optional] 
**MustNot** | Pointer to [**[]MetaQueryForSDK**](MetaQueryForSDK.md) | query, [query1, query2] | [optional] 
**Should** | Pointer to [**[]MetaQueryForSDK**](MetaQueryForSDK.md) | query, [query1, query2] | [optional] 

## Methods

### NewMetaBoolQueryForSDK

`func NewMetaBoolQueryForSDK() *MetaBoolQueryForSDK`

NewMetaBoolQueryForSDK instantiates a new MetaBoolQueryForSDK object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaBoolQueryForSDKWithDefaults

`func NewMetaBoolQueryForSDKWithDefaults() *MetaBoolQueryForSDK`

NewMetaBoolQueryForSDKWithDefaults instantiates a new MetaBoolQueryForSDK object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *MetaBoolQueryForSDK) GetFilter() []MetaQueryForSDK`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *MetaBoolQueryForSDK) GetFilterOk() (*[]MetaQueryForSDK, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *MetaBoolQueryForSDK) SetFilter(v []MetaQueryForSDK)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *MetaBoolQueryForSDK) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetMinimumShouldMatch

`func (o *MetaBoolQueryForSDK) GetMinimumShouldMatch() float32`

GetMinimumShouldMatch returns the MinimumShouldMatch field if non-nil, zero value otherwise.

### GetMinimumShouldMatchOk

`func (o *MetaBoolQueryForSDK) GetMinimumShouldMatchOk() (*float32, bool)`

GetMinimumShouldMatchOk returns a tuple with the MinimumShouldMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumShouldMatch

`func (o *MetaBoolQueryForSDK) SetMinimumShouldMatch(v float32)`

SetMinimumShouldMatch sets MinimumShouldMatch field to given value.

### HasMinimumShouldMatch

`func (o *MetaBoolQueryForSDK) HasMinimumShouldMatch() bool`

HasMinimumShouldMatch returns a boolean if a field has been set.

### GetMust

`func (o *MetaBoolQueryForSDK) GetMust() []MetaQueryForSDK`

GetMust returns the Must field if non-nil, zero value otherwise.

### GetMustOk

`func (o *MetaBoolQueryForSDK) GetMustOk() (*[]MetaQueryForSDK, bool)`

GetMustOk returns a tuple with the Must field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMust

`func (o *MetaBoolQueryForSDK) SetMust(v []MetaQueryForSDK)`

SetMust sets Must field to given value.

### HasMust

`func (o *MetaBoolQueryForSDK) HasMust() bool`

HasMust returns a boolean if a field has been set.

### GetMustNot

`func (o *MetaBoolQueryForSDK) GetMustNot() []MetaQueryForSDK`

GetMustNot returns the MustNot field if non-nil, zero value otherwise.

### GetMustNotOk

`func (o *MetaBoolQueryForSDK) GetMustNotOk() (*[]MetaQueryForSDK, bool)`

GetMustNotOk returns a tuple with the MustNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMustNot

`func (o *MetaBoolQueryForSDK) SetMustNot(v []MetaQueryForSDK)`

SetMustNot sets MustNot field to given value.

### HasMustNot

`func (o *MetaBoolQueryForSDK) HasMustNot() bool`

HasMustNot returns a boolean if a field has been set.

### GetShould

`func (o *MetaBoolQueryForSDK) GetShould() []MetaQueryForSDK`

GetShould returns the Should field if non-nil, zero value otherwise.

### GetShouldOk

`func (o *MetaBoolQueryForSDK) GetShouldOk() (*[]MetaQueryForSDK, bool)`

GetShouldOk returns a tuple with the Should field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShould

`func (o *MetaBoolQueryForSDK) SetShould(v []MetaQueryForSDK)`

SetShould sets Should field to given value.

### HasShould

`func (o *MetaBoolQueryForSDK) HasShould() bool`

HasShould returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


