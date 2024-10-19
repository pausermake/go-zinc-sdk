# MetaQueryForSDK

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bool** | Pointer to [**MetaBoolQueryForSDK**](MetaBoolQueryForSDK.md) | . | [optional] 
**Exists** | Pointer to [**MetaExistsQuery**](MetaExistsQuery.md) | . | [optional] 
**Fuzzy** | Pointer to [**map[string]MetaFuzzyQuery**](MetaFuzzyQuery.md) | simple, PrefixQuery | [optional] 
**Ids** | Pointer to [**MetaIdsQuery**](MetaIdsQuery.md) | . | [optional] 
**Match** | Pointer to [**map[string]MetaMatchQuery**](MetaMatchQuery.md) | simple, MatchQuery | [optional] 
**MatchAll** | Pointer to **map[string]interface{}** | just set or null | [optional] 
**MatchBoolPrefix** | Pointer to [**map[string]MetaMatchBoolPrefixQuery**](MetaMatchBoolPrefixQuery.md) | simple, MatchBoolPrefixQuery | [optional] 
**MatchNone** | Pointer to **map[string]interface{}** | just set or null | [optional] 
**MatchPhrase** | Pointer to [**map[string]MetaMatchPhraseQuery**](MetaMatchPhraseQuery.md) | simple, MatchPhraseQuery | [optional] 
**MatchPhrasePrefix** | Pointer to [**map[string]MetaMatchPhrasePrefixQuery**](MetaMatchPhrasePrefixQuery.md) | simple, MatchPhrasePrefixQuery | [optional] 
**MultiMatch** | Pointer to [**MetaMultiMatchQuery**](MetaMultiMatchQuery.md) | . | [optional] 
**Prefix** | Pointer to [**map[string]MetaPrefixQuery**](MetaPrefixQuery.md) | . | [optional] 
**QueryString** | Pointer to [**MetaQueryStringQuery**](MetaQueryStringQuery.md) | . | [optional] 
**Range** | Pointer to [**map[string]MetaRangeQueryForSDK**](MetaRangeQueryForSDK.md) | simple, FuzzyQuery | [optional] 
**Regexp** | Pointer to [**map[string]MetaRegexpQuery**](MetaRegexpQuery.md) | simple, FuzzyQuery | [optional] 
**SimpleQueryString** | Pointer to [**MetaSimpleQueryStringQuery**](MetaSimpleQueryStringQuery.md) | . | [optional] 
**Term** | Pointer to [**map[string]MetaTermQueryForSDK**](MetaTermQueryForSDK.md) | simple, TermQuery | [optional] 
**Terms** | Pointer to **map[string]map[string]interface{}** | . | [optional] 
**Wildcard** | Pointer to [**map[string]MetaWildcardQuery**](MetaWildcardQuery.md) | simple, WildcardQuery | [optional] 

## Methods

### NewMetaQueryForSDK

`func NewMetaQueryForSDK() *MetaQueryForSDK`

NewMetaQueryForSDK instantiates a new MetaQueryForSDK object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaQueryForSDKWithDefaults

`func NewMetaQueryForSDKWithDefaults() *MetaQueryForSDK`

NewMetaQueryForSDKWithDefaults instantiates a new MetaQueryForSDK object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBool

`func (o *MetaQueryForSDK) GetBool() MetaBoolQueryForSDK`

GetBool returns the Bool field if non-nil, zero value otherwise.

### GetBoolOk

`func (o *MetaQueryForSDK) GetBoolOk() (*MetaBoolQueryForSDK, bool)`

GetBoolOk returns a tuple with the Bool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBool

`func (o *MetaQueryForSDK) SetBool(v MetaBoolQueryForSDK)`

SetBool sets Bool field to given value.

### HasBool

`func (o *MetaQueryForSDK) HasBool() bool`

HasBool returns a boolean if a field has been set.

### GetExists

`func (o *MetaQueryForSDK) GetExists() MetaExistsQuery`

GetExists returns the Exists field if non-nil, zero value otherwise.

### GetExistsOk

`func (o *MetaQueryForSDK) GetExistsOk() (*MetaExistsQuery, bool)`

GetExistsOk returns a tuple with the Exists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExists

`func (o *MetaQueryForSDK) SetExists(v MetaExistsQuery)`

SetExists sets Exists field to given value.

### HasExists

`func (o *MetaQueryForSDK) HasExists() bool`

HasExists returns a boolean if a field has been set.

### GetFuzzy

`func (o *MetaQueryForSDK) GetFuzzy() map[string]MetaFuzzyQuery`

GetFuzzy returns the Fuzzy field if non-nil, zero value otherwise.

### GetFuzzyOk

`func (o *MetaQueryForSDK) GetFuzzyOk() (*map[string]MetaFuzzyQuery, bool)`

GetFuzzyOk returns a tuple with the Fuzzy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuzzy

`func (o *MetaQueryForSDK) SetFuzzy(v map[string]MetaFuzzyQuery)`

SetFuzzy sets Fuzzy field to given value.

### HasFuzzy

`func (o *MetaQueryForSDK) HasFuzzy() bool`

HasFuzzy returns a boolean if a field has been set.

### GetIds

`func (o *MetaQueryForSDK) GetIds() MetaIdsQuery`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *MetaQueryForSDK) GetIdsOk() (*MetaIdsQuery, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *MetaQueryForSDK) SetIds(v MetaIdsQuery)`

SetIds sets Ids field to given value.

### HasIds

`func (o *MetaQueryForSDK) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetMatch

`func (o *MetaQueryForSDK) GetMatch() map[string]MetaMatchQuery`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *MetaQueryForSDK) GetMatchOk() (*map[string]MetaMatchQuery, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *MetaQueryForSDK) SetMatch(v map[string]MetaMatchQuery)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *MetaQueryForSDK) HasMatch() bool`

HasMatch returns a boolean if a field has been set.

### GetMatchAll

`func (o *MetaQueryForSDK) GetMatchAll() map[string]interface{}`

GetMatchAll returns the MatchAll field if non-nil, zero value otherwise.

### GetMatchAllOk

`func (o *MetaQueryForSDK) GetMatchAllOk() (*map[string]interface{}, bool)`

GetMatchAllOk returns a tuple with the MatchAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchAll

`func (o *MetaQueryForSDK) SetMatchAll(v map[string]interface{})`

SetMatchAll sets MatchAll field to given value.

### HasMatchAll

`func (o *MetaQueryForSDK) HasMatchAll() bool`

HasMatchAll returns a boolean if a field has been set.

### GetMatchBoolPrefix

`func (o *MetaQueryForSDK) GetMatchBoolPrefix() map[string]MetaMatchBoolPrefixQuery`

GetMatchBoolPrefix returns the MatchBoolPrefix field if non-nil, zero value otherwise.

### GetMatchBoolPrefixOk

`func (o *MetaQueryForSDK) GetMatchBoolPrefixOk() (*map[string]MetaMatchBoolPrefixQuery, bool)`

GetMatchBoolPrefixOk returns a tuple with the MatchBoolPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchBoolPrefix

`func (o *MetaQueryForSDK) SetMatchBoolPrefix(v map[string]MetaMatchBoolPrefixQuery)`

SetMatchBoolPrefix sets MatchBoolPrefix field to given value.

### HasMatchBoolPrefix

`func (o *MetaQueryForSDK) HasMatchBoolPrefix() bool`

HasMatchBoolPrefix returns a boolean if a field has been set.

### GetMatchNone

`func (o *MetaQueryForSDK) GetMatchNone() map[string]interface{}`

GetMatchNone returns the MatchNone field if non-nil, zero value otherwise.

### GetMatchNoneOk

`func (o *MetaQueryForSDK) GetMatchNoneOk() (*map[string]interface{}, bool)`

GetMatchNoneOk returns a tuple with the MatchNone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchNone

`func (o *MetaQueryForSDK) SetMatchNone(v map[string]interface{})`

SetMatchNone sets MatchNone field to given value.

### HasMatchNone

`func (o *MetaQueryForSDK) HasMatchNone() bool`

HasMatchNone returns a boolean if a field has been set.

### GetMatchPhrase

`func (o *MetaQueryForSDK) GetMatchPhrase() map[string]MetaMatchPhraseQuery`

GetMatchPhrase returns the MatchPhrase field if non-nil, zero value otherwise.

### GetMatchPhraseOk

`func (o *MetaQueryForSDK) GetMatchPhraseOk() (*map[string]MetaMatchPhraseQuery, bool)`

GetMatchPhraseOk returns a tuple with the MatchPhrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchPhrase

`func (o *MetaQueryForSDK) SetMatchPhrase(v map[string]MetaMatchPhraseQuery)`

SetMatchPhrase sets MatchPhrase field to given value.

### HasMatchPhrase

`func (o *MetaQueryForSDK) HasMatchPhrase() bool`

HasMatchPhrase returns a boolean if a field has been set.

### GetMatchPhrasePrefix

`func (o *MetaQueryForSDK) GetMatchPhrasePrefix() map[string]MetaMatchPhrasePrefixQuery`

GetMatchPhrasePrefix returns the MatchPhrasePrefix field if non-nil, zero value otherwise.

### GetMatchPhrasePrefixOk

`func (o *MetaQueryForSDK) GetMatchPhrasePrefixOk() (*map[string]MetaMatchPhrasePrefixQuery, bool)`

GetMatchPhrasePrefixOk returns a tuple with the MatchPhrasePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchPhrasePrefix

`func (o *MetaQueryForSDK) SetMatchPhrasePrefix(v map[string]MetaMatchPhrasePrefixQuery)`

SetMatchPhrasePrefix sets MatchPhrasePrefix field to given value.

### HasMatchPhrasePrefix

`func (o *MetaQueryForSDK) HasMatchPhrasePrefix() bool`

HasMatchPhrasePrefix returns a boolean if a field has been set.

### GetMultiMatch

`func (o *MetaQueryForSDK) GetMultiMatch() MetaMultiMatchQuery`

GetMultiMatch returns the MultiMatch field if non-nil, zero value otherwise.

### GetMultiMatchOk

`func (o *MetaQueryForSDK) GetMultiMatchOk() (*MetaMultiMatchQuery, bool)`

GetMultiMatchOk returns a tuple with the MultiMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiMatch

`func (o *MetaQueryForSDK) SetMultiMatch(v MetaMultiMatchQuery)`

SetMultiMatch sets MultiMatch field to given value.

### HasMultiMatch

`func (o *MetaQueryForSDK) HasMultiMatch() bool`

HasMultiMatch returns a boolean if a field has been set.

### GetPrefix

`func (o *MetaQueryForSDK) GetPrefix() map[string]MetaPrefixQuery`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *MetaQueryForSDK) GetPrefixOk() (*map[string]MetaPrefixQuery, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *MetaQueryForSDK) SetPrefix(v map[string]MetaPrefixQuery)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *MetaQueryForSDK) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetQueryString

`func (o *MetaQueryForSDK) GetQueryString() MetaQueryStringQuery`

GetQueryString returns the QueryString field if non-nil, zero value otherwise.

### GetQueryStringOk

`func (o *MetaQueryForSDK) GetQueryStringOk() (*MetaQueryStringQuery, bool)`

GetQueryStringOk returns a tuple with the QueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryString

`func (o *MetaQueryForSDK) SetQueryString(v MetaQueryStringQuery)`

SetQueryString sets QueryString field to given value.

### HasQueryString

`func (o *MetaQueryForSDK) HasQueryString() bool`

HasQueryString returns a boolean if a field has been set.

### GetRange

`func (o *MetaQueryForSDK) GetRange() map[string]MetaRangeQueryForSDK`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *MetaQueryForSDK) GetRangeOk() (*map[string]MetaRangeQueryForSDK, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *MetaQueryForSDK) SetRange(v map[string]MetaRangeQueryForSDK)`

SetRange sets Range field to given value.

### HasRange

`func (o *MetaQueryForSDK) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetRegexp

`func (o *MetaQueryForSDK) GetRegexp() map[string]MetaRegexpQuery`

GetRegexp returns the Regexp field if non-nil, zero value otherwise.

### GetRegexpOk

`func (o *MetaQueryForSDK) GetRegexpOk() (*map[string]MetaRegexpQuery, bool)`

GetRegexpOk returns a tuple with the Regexp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegexp

`func (o *MetaQueryForSDK) SetRegexp(v map[string]MetaRegexpQuery)`

SetRegexp sets Regexp field to given value.

### HasRegexp

`func (o *MetaQueryForSDK) HasRegexp() bool`

HasRegexp returns a boolean if a field has been set.

### GetSimpleQueryString

`func (o *MetaQueryForSDK) GetSimpleQueryString() MetaSimpleQueryStringQuery`

GetSimpleQueryString returns the SimpleQueryString field if non-nil, zero value otherwise.

### GetSimpleQueryStringOk

`func (o *MetaQueryForSDK) GetSimpleQueryStringOk() (*MetaSimpleQueryStringQuery, bool)`

GetSimpleQueryStringOk returns a tuple with the SimpleQueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimpleQueryString

`func (o *MetaQueryForSDK) SetSimpleQueryString(v MetaSimpleQueryStringQuery)`

SetSimpleQueryString sets SimpleQueryString field to given value.

### HasSimpleQueryString

`func (o *MetaQueryForSDK) HasSimpleQueryString() bool`

HasSimpleQueryString returns a boolean if a field has been set.

### GetTerm

`func (o *MetaQueryForSDK) GetTerm() map[string]MetaTermQueryForSDK`

GetTerm returns the Term field if non-nil, zero value otherwise.

### GetTermOk

`func (o *MetaQueryForSDK) GetTermOk() (*map[string]MetaTermQueryForSDK, bool)`

GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerm

`func (o *MetaQueryForSDK) SetTerm(v map[string]MetaTermQueryForSDK)`

SetTerm sets Term field to given value.

### HasTerm

`func (o *MetaQueryForSDK) HasTerm() bool`

HasTerm returns a boolean if a field has been set.

### GetTerms

`func (o *MetaQueryForSDK) GetTerms() map[string]map[string]interface{}`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *MetaQueryForSDK) GetTermsOk() (*map[string]map[string]interface{}, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *MetaQueryForSDK) SetTerms(v map[string]map[string]interface{})`

SetTerms sets Terms field to given value.

### HasTerms

`func (o *MetaQueryForSDK) HasTerms() bool`

HasTerms returns a boolean if a field has been set.

### GetWildcard

`func (o *MetaQueryForSDK) GetWildcard() map[string]MetaWildcardQuery`

GetWildcard returns the Wildcard field if non-nil, zero value otherwise.

### GetWildcardOk

`func (o *MetaQueryForSDK) GetWildcardOk() (*map[string]MetaWildcardQuery, bool)`

GetWildcardOk returns a tuple with the Wildcard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWildcard

`func (o *MetaQueryForSDK) SetWildcard(v map[string]MetaWildcardQuery)`

SetWildcard sets Wildcard field to given value.

### HasWildcard

`func (o *MetaQueryForSDK) HasWildcard() bool`

HasWildcard returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


