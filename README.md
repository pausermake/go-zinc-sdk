# Go API client for openapi

Zinc Search engine API documents https://zincsearch-docs.zinc.dev

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Generator version: 7.9.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://www.zincsearch.com](https://www.zincsearch.com)

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `openapi.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `openapi.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `openapi.ContextOperationServerIndices` and `openapi.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultAPI* | [**Healthz**](docs/DefaultAPI.md#healthz) | **Get** /healthz | Get healthz
*DefaultAPI* | [**Version**](docs/DefaultAPI.md#version) | **Get** /version | Get version
*DocumentAPI* | [**Bulk**](docs/DocumentAPI.md#bulk) | **Post** /api/_bulk | Bulk documents
*DocumentAPI* | [**Bulkv2**](docs/DocumentAPI.md#bulkv2) | **Post** /api/_bulkv2 | Bulkv2 documents
*DocumentAPI* | [**DeleteDocument**](docs/DocumentAPI.md#deletedocument) | **Delete** /api/{index}/_doc/{id} | Delete document
*DocumentAPI* | [**ESBulk**](docs/DocumentAPI.md#esbulk) | **Post** /es/_bulk | ES bulk documents
*DocumentAPI* | [**GetDocument**](docs/DocumentAPI.md#getdocument) | **Get** /api/{index}/_doc/{id} | get document with id
*DocumentAPI* | [**IndexDocument**](docs/DocumentAPI.md#indexdocument) | **Post** /api/{index}/_doc | Create or update document
*DocumentAPI* | [**IndexDocumentWithID**](docs/DocumentAPI.md#indexdocumentwithid) | **Put** /api/{index}/_doc/{id} | Create or update document with id
*DocumentAPI* | [**Multi**](docs/DocumentAPI.md#multi) | **Post** /api/{index}/_multi | Multi documents
*DocumentAPI* | [**UpdateDocument**](docs/DocumentAPI.md#updatedocument) | **Post** /api/{index}/_update/{id} | Update document with id
*IndexAPI* | [**AddOrRemoveESAlias**](docs/IndexAPI.md#addorremoveesalias) | **Post** /es/_aliases | Add or remove index alias for compatible ES
*IndexAPI* | [**Analyze**](docs/IndexAPI.md#analyze) | **Post** /api/_analyze | Analyze
*IndexAPI* | [**AnalyzeIndex**](docs/IndexAPI.md#analyzeindex) | **Post** /api/{index}/_analyze | Analyze
*IndexAPI* | [**CreateIndex**](docs/IndexAPI.md#createindex) | **Post** /api/index | Create index
*IndexAPI* | [**CreateTemplate**](docs/IndexAPI.md#createtemplate) | **Post** /es/_index_template | Create update index template
*IndexAPI* | [**DeleteIndex**](docs/IndexAPI.md#deleteindex) | **Delete** /api/index/{index} | Delete index
*IndexAPI* | [**DeleteTemplate**](docs/IndexAPI.md#deletetemplate) | **Delete** /es/_index_template/{name} | Delete template
*IndexAPI* | [**ESCreateIndex**](docs/IndexAPI.md#escreateindex) | **Put** /es/{index} | Create index for compatible ES
*IndexAPI* | [**ESGetMapping**](docs/IndexAPI.md#esgetmapping) | **Get** /es/{index}/_mapping | Get index mappings for compatible ES
*IndexAPI* | [**EsExists**](docs/IndexAPI.md#esexists) | **Head** /es/{index} | Checks if the index exists for compatible ES
*IndexAPI* | [**Exists**](docs/IndexAPI.md#exists) | **Head** /api/index/{index} | Checks if the index exists
*IndexAPI* | [**GetESAliases**](docs/IndexAPI.md#getesaliases) | **Get** /es/{target}/_alias/{target_alias} | Get index alias for compatible ES
*IndexAPI* | [**GetIndex**](docs/IndexAPI.md#getindex) | **Get** /api/index/{index} | Get index metadata
*IndexAPI* | [**GetMapping**](docs/IndexAPI.md#getmapping) | **Get** /api/{index}/_mapping | Get index mappings
*IndexAPI* | [**GetSettings**](docs/IndexAPI.md#getsettings) | **Get** /api/{index}/_settings | Get index settings
*IndexAPI* | [**GetTemplate**](docs/IndexAPI.md#gettemplate) | **Get** /es/_index_template/{name} | Get index template
*IndexAPI* | [**IndexNameList**](docs/IndexAPI.md#indexnamelist) | **Get** /api/index_name | List index Name
*IndexAPI* | [**ListIndexes**](docs/IndexAPI.md#listindexes) | **Get** /api/index | List indexes
*IndexAPI* | [**ListTemplates**](docs/IndexAPI.md#listtemplates) | **Get** /es/_index_template | List index teplates
*IndexAPI* | [**Refresh**](docs/IndexAPI.md#refresh) | **Post** /api/index/{index}/refresh | Resfresh index
*IndexAPI* | [**SetMapping**](docs/IndexAPI.md#setmapping) | **Put** /api/{index}/_mapping | Set index mappings
*IndexAPI* | [**SetSettings**](docs/IndexAPI.md#setsettings) | **Put** /api/{index}/_settings | Set index Settings
*IndexAPI* | [**UpdateTemplate**](docs/IndexAPI.md#updatetemplate) | **Put** /es/_index_template/{name} | Create update index template
*PermissionAPI* | [**ListPermissions**](docs/PermissionAPI.md#listpermissions) | **Get** /api/permissions | List permissions
*RoleAPI* | [**CreateRole**](docs/RoleAPI.md#createrole) | **Post** /api/role | Create role
*RoleAPI* | [**DeleteRole**](docs/RoleAPI.md#deleterole) | **Delete** /api/role/{id} | Delete role
*RoleAPI* | [**ListRoles**](docs/RoleAPI.md#listroles) | **Get** /api/role | List role
*SearchAPI* | [**DeleteByQuery**](docs/SearchAPI.md#deletebyquery) | **Post** /es/{index}/_delete_by_query | Searches the index and deletes all matched documents
*SearchAPI* | [**MSearch**](docs/SearchAPI.md#msearch) | **Post** /es/_msearch | Search V2 MultipleSearch for compatible ES
*SearchAPI* | [**Search**](docs/SearchAPI.md#search) | **Post** /es/{index}/_search | Search V2 DSL for compatible ES
*SearchAPI* | [**SearchV1**](docs/SearchAPI.md#searchv1) | **Post** /api/{index}/_search | Search V1
*UserAPI* | [**CreateUser**](docs/UserAPI.md#createuser) | **Post** /api/user | Create user
*UserAPI* | [**DeleteUser**](docs/UserAPI.md#deleteuser) | **Delete** /api/user/{id} | Delete user
*UserAPI* | [**ListUsers**](docs/UserAPI.md#listusers) | **Get** /api/user | List user
*UserAPI* | [**Login**](docs/UserAPI.md#login) | **Post** /api/login | Login
*UserAPI* | [**UpdateUser**](docs/UserAPI.md#updateuser) | **Put** /api/user | Update user


## Documentation For Models

 - [AggregationHistogramBound](docs/AggregationHistogramBound.md)
 - [AuthLoginRequest](docs/AuthLoginRequest.md)
 - [AuthLoginResponse](docs/AuthLoginResponse.md)
 - [AuthLoginUser](docs/AuthLoginUser.md)
 - [IndexAnalyzeResponse](docs/IndexAnalyzeResponse.md)
 - [IndexAnalyzeResponseToken](docs/IndexAnalyzeResponseToken.md)
 - [IndexIndexListResponse](docs/IndexIndexListResponse.md)
 - [MetaAggregationAutoDateHistogram](docs/MetaAggregationAutoDateHistogram.md)
 - [MetaAggregationDateHistogram](docs/MetaAggregationDateHistogram.md)
 - [MetaAggregationDateRange](docs/MetaAggregationDateRange.md)
 - [MetaAggregationHistogram](docs/MetaAggregationHistogram.md)
 - [MetaAggregationIPRange](docs/MetaAggregationIPRange.md)
 - [MetaAggregationMetric](docs/MetaAggregationMetric.md)
 - [MetaAggregationRange](docs/MetaAggregationRange.md)
 - [MetaAggregationResponse](docs/MetaAggregationResponse.md)
 - [MetaAggregations](docs/MetaAggregations.md)
 - [MetaAggregationsTerms](docs/MetaAggregationsTerms.md)
 - [MetaAnalyzer](docs/MetaAnalyzer.md)
 - [MetaBoolQueryForSDK](docs/MetaBoolQueryForSDK.md)
 - [MetaDateRange](docs/MetaDateRange.md)
 - [MetaExistsQuery](docs/MetaExistsQuery.md)
 - [MetaFuzzyQuery](docs/MetaFuzzyQuery.md)
 - [MetaHTTPResponse](docs/MetaHTTPResponse.md)
 - [MetaHTTPResponseDeleteByQuery](docs/MetaHTTPResponseDeleteByQuery.md)
 - [MetaHTTPResponseDocument](docs/MetaHTTPResponseDocument.md)
 - [MetaHTTPResponseError](docs/MetaHTTPResponseError.md)
 - [MetaHTTPResponseID](docs/MetaHTTPResponseID.md)
 - [MetaHTTPResponseIndex](docs/MetaHTTPResponseIndex.md)
 - [MetaHTTPResponseRecordCount](docs/MetaHTTPResponseRecordCount.md)
 - [MetaHTTPResponseTemplate](docs/MetaHTTPResponseTemplate.md)
 - [MetaHealthzResponse](docs/MetaHealthzResponse.md)
 - [MetaHighlight](docs/MetaHighlight.md)
 - [MetaHit](docs/MetaHit.md)
 - [MetaHits](docs/MetaHits.md)
 - [MetaHttpRetriesResponse](docs/MetaHttpRetriesResponse.md)
 - [MetaIPRange](docs/MetaIPRange.md)
 - [MetaIdsQuery](docs/MetaIdsQuery.md)
 - [MetaIndexAnalysis](docs/MetaIndexAnalysis.md)
 - [MetaIndexSettings](docs/MetaIndexSettings.md)
 - [MetaIndexSimple](docs/MetaIndexSimple.md)
 - [MetaIndexTemplate](docs/MetaIndexTemplate.md)
 - [MetaJSONIngest](docs/MetaJSONIngest.md)
 - [MetaMappings](docs/MetaMappings.md)
 - [MetaMatchBoolPrefixQuery](docs/MetaMatchBoolPrefixQuery.md)
 - [MetaMatchPhrasePrefixQuery](docs/MetaMatchPhrasePrefixQuery.md)
 - [MetaMatchPhraseQuery](docs/MetaMatchPhraseQuery.md)
 - [MetaMatchQuery](docs/MetaMatchQuery.md)
 - [MetaMultiMatchQuery](docs/MetaMultiMatchQuery.md)
 - [MetaPage](docs/MetaPage.md)
 - [MetaPrefixQuery](docs/MetaPrefixQuery.md)
 - [MetaProperty](docs/MetaProperty.md)
 - [MetaQueryForSDK](docs/MetaQueryForSDK.md)
 - [MetaQueryStringQuery](docs/MetaQueryStringQuery.md)
 - [MetaRange](docs/MetaRange.md)
 - [MetaRangeQueryForSDK](docs/MetaRangeQueryForSDK.md)
 - [MetaRegexpQuery](docs/MetaRegexpQuery.md)
 - [MetaRole](docs/MetaRole.md)
 - [MetaSearchResponse](docs/MetaSearchResponse.md)
 - [MetaShards](docs/MetaShards.md)
 - [MetaSimpleQueryStringQuery](docs/MetaSimpleQueryStringQuery.md)
 - [MetaTemplate](docs/MetaTemplate.md)
 - [MetaTemplateTemplate](docs/MetaTemplateTemplate.md)
 - [MetaTermQueryForSDK](docs/MetaTermQueryForSDK.md)
 - [MetaTotal](docs/MetaTotal.md)
 - [MetaUser](docs/MetaUser.md)
 - [MetaVersionResponse](docs/MetaVersionResponse.md)
 - [MetaWildcardQuery](docs/MetaWildcardQuery.md)
 - [MetaZincQueryForSDK](docs/MetaZincQueryForSDK.md)
 - [V1AggregationDateRange](docs/V1AggregationDateRange.md)
 - [V1AggregationNumberRange](docs/V1AggregationNumberRange.md)
 - [V1AggregationParams](docs/V1AggregationParams.md)
 - [V1QueryParams](docs/V1QueryParams.md)
 - [V1ZincQueryForSDK](docs/V1ZincQueryForSDK.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### BasicAuth

- **Type**: HTTP basic authentication

Example

```go
auth := context.WithValue(context.Background(), openapi.ContextBasicAuth, openapi.BasicAuth{
	UserName: "username",
	Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



