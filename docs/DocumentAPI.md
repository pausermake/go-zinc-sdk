# \DocumentAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Bulk**](DocumentAPI.md#Bulk) | **Post** /api/_bulk | Bulk documents
[**Bulkv2**](DocumentAPI.md#Bulkv2) | **Post** /api/_bulkv2 | Bulkv2 documents
[**DeleteDocument**](DocumentAPI.md#DeleteDocument) | **Delete** /api/{index}/_doc/{id} | Delete document
[**ESBulk**](DocumentAPI.md#ESBulk) | **Post** /es/_bulk | ES bulk documents
[**GetDocument**](DocumentAPI.md#GetDocument) | **Get** /api/{index}/_doc/{id} | get document with id
[**IndexDocument**](DocumentAPI.md#IndexDocument) | **Post** /api/{index}/_doc | Create or update document
[**IndexDocumentWithID**](DocumentAPI.md#IndexDocumentWithID) | **Put** /api/{index}/_doc/{id} | Create or update document with id
[**Multi**](DocumentAPI.md#Multi) | **Post** /api/{index}/_multi | Multi documents
[**UpdateDocument**](DocumentAPI.md#UpdateDocument) | **Post** /api/{index}/_update/{id} | Update document with id



## Bulk

> MetaHTTPResponseRecordCount Bulk(ctx).Query(query).Execute()

Bulk documents

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	query := "query_example" // string | Query

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentAPI.Bulk(context.Background()).Query(query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentAPI.Bulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Bulk`: MetaHTTPResponseRecordCount
	fmt.Fprintf(os.Stdout, "Response from `DocumentAPI.Bulk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Query | 

### Return type

[**MetaHTTPResponseRecordCount**](MetaHTTPResponseRecordCount.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Bulkv2

> MetaHTTPResponseRecordCount Bulkv2(ctx).Query(query).Execute()

Bulkv2 documents

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	query := *openapiclient.NewMetaJSONIngest() // MetaJSONIngest | Query

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentAPI.Bulkv2(context.Background()).Query(query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentAPI.Bulkv2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Bulkv2`: MetaHTTPResponseRecordCount
	fmt.Fprintf(os.Stdout, "Response from `DocumentAPI.Bulkv2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkv2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | [**MetaJSONIngest**](MetaJSONIngest.md) | Query | 

### Return type

[**MetaHTTPResponseRecordCount**](MetaHTTPResponseRecordCount.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDocument

> MetaHTTPResponseDocument DeleteDocument(ctx, index, id).Execute()

Delete document

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	index := "index_example" // string | Index
	id := "id_example" // string | ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentAPI.DeleteDocument(context.Background(), index, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentAPI.DeleteDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDocument`: MetaHTTPResponseDocument
	fmt.Fprintf(os.Stdout, "Response from `DocumentAPI.DeleteDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | Index | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MetaHTTPResponseDocument**](MetaHTTPResponseDocument.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ESBulk

> map[string]interface{} ESBulk(ctx).Query(query).Execute()

ES bulk documents

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	query := "query_example" // string | Query

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentAPI.ESBulk(context.Background()).Query(query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentAPI.ESBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ESBulk`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DocumentAPI.ESBulk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiESBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Query | 

### Return type

**map[string]interface{}**

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocument

> MetaHit GetDocument(ctx, index, id).Execute()

get document with id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	index := "index_example" // string | Index
	id := "id_example" // string | ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentAPI.GetDocument(context.Background(), index, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentAPI.GetDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDocument`: MetaHit
	fmt.Fprintf(os.Stdout, "Response from `DocumentAPI.GetDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | Index | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MetaHit**](MetaHit.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IndexDocument

> MetaHTTPResponseID IndexDocument(ctx, index).Document(document).Execute()

Create or update document

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	index := "index_example" // string | Index
	document := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | Document

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentAPI.IndexDocument(context.Background(), index).Document(document).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentAPI.IndexDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IndexDocument`: MetaHTTPResponseID
	fmt.Fprintf(os.Stdout, "Response from `DocumentAPI.IndexDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | Index | 

### Other Parameters

Other parameters are passed through a pointer to a apiIndexDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **document** | **map[string]interface{}** | Document | 

### Return type

[**MetaHTTPResponseID**](MetaHTTPResponseID.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IndexDocumentWithID

> MetaHTTPResponseID IndexDocumentWithID(ctx, index, id).Document(document).Execute()

Create or update document with id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	index := "index_example" // string | Index
	id := "id_example" // string | ID
	document := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | Document

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentAPI.IndexDocumentWithID(context.Background(), index, id).Document(document).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentAPI.IndexDocumentWithID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IndexDocumentWithID`: MetaHTTPResponseID
	fmt.Fprintf(os.Stdout, "Response from `DocumentAPI.IndexDocumentWithID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | Index | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiIndexDocumentWithIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **document** | **map[string]interface{}** | Document | 

### Return type

[**MetaHTTPResponseID**](MetaHTTPResponseID.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Multi

> MetaHTTPResponseRecordCount Multi(ctx, index).Query(query).Execute()

Multi documents

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	index := "index_example" // string | Index
	query := "query_example" // string | Query

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentAPI.Multi(context.Background(), index).Query(query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentAPI.Multi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Multi`: MetaHTTPResponseRecordCount
	fmt.Fprintf(os.Stdout, "Response from `DocumentAPI.Multi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | Index | 

### Other Parameters

Other parameters are passed through a pointer to a apiMultiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | **string** | Query | 

### Return type

[**MetaHTTPResponseRecordCount**](MetaHTTPResponseRecordCount.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDocument

> MetaHTTPResponseID UpdateDocument(ctx, index, id).Document(document).Execute()

Update document with id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	index := "index_example" // string | Index
	id := "id_example" // string | ID
	document := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | Document

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentAPI.UpdateDocument(context.Background(), index, id).Document(document).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentAPI.UpdateDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDocument`: MetaHTTPResponseID
	fmt.Fprintf(os.Stdout, "Response from `DocumentAPI.UpdateDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | Index | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **document** | **map[string]interface{}** | Document | 

### Return type

[**MetaHTTPResponseID**](MetaHTTPResponseID.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

