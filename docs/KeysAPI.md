# KeysAPI

All URIs are relative to *https://filemage.ppfa.org:8443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKey**](KeysAPI.md#CreateKey) | **Post** /users/{userId}/keys/ | Add key
[**DeleteKey**](KeysAPI.md#DeleteKey) | **Delete** /users/{userId}/keys/{keyId}/ | Delete key



## CreateKey

> string CreateKey(ctx, userId).Body(body).Execute()

Add key



### Example

```go
package main

import (
    "bytes"
    "context"
    "encoding/json"
    "fmt"
    "os"
    log "github.com/sirupsen/logrus"
    fm "github.com/elgranjero/filemagego"
)

func fm_CreateKey(client fm.KeysAPI) {
    userId := int32(56) // int32 | ID of user to associate key with
    body := *fm.NewKey() // Key | Key settings

    resp, _, err := client.CreateKey(context.Background(), userId).Body(body).Execute()
    if err != nil {
        log.Errorf("CreateKey: %s", err.Error())
        return
    }
    // response from `CreateKey`: string
    // fmt.Printf("%v\n", resp)
    b, _ := json.Marshal(resp)
    fmt.Fprintf(os.Stdout, "%s\n", jsonPrettyPrint(string(b)))
}

func main() {
    cfg := fm.NewConfiguration()
    cfg.AddDefaultHeader("filemage-api-token", os.Getenv("FILEMAGE_API_TOKEN"))
    host := os.Getenv("FILEMAGE_HOST")
    if len(host) > 0 {
        cfg.Servers[0].URL = host
    }
    client := fm.NewAPIClient(cfg)

    fm_CreateKey(client.KeysAPI)
}

func jsonPrettyPrint(in string) string {
        var out bytes.Buffer
        err := json.Indent(&out, []byte(in), "", "  ")
        if err != nil {
                return in
        }
        return out.String()
}

```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | ID of user to associate key with | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**Key**](Key.md) | Key settings | 

### Return type

**string**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/x-pem-file

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKey

> DeleteKey(ctx, userId, keyId).Execute()

Delete key

### Example

```go
package main

import (
    "bytes"
    "context"
    "encoding/json"
    "fmt"
    "os"
    log "github.com/sirupsen/logrus"
    fm "github.com/elgranjero/filemagego"
)

func fm_DeleteKey(client fm.KeysAPI) {
    userId := int32(56) // int32 | ID of user to remove key from
    keyId := int32(56) // int32 | ID of key to remove

    _, err := client.DeleteKey(context.Background(), userId, keyId).Execute()
    if err != nil {
        log.Errorf("DeleteKey: %s", err.Error())
        return
    }
    b, _ := json.Marshal(resp)
    fmt.Fprintf(os.Stdout, "%s\n", jsonPrettyPrint(string(b)))
}

func main() {
    cfg := fm.NewConfiguration()
    cfg.AddDefaultHeader("filemage-api-token", os.Getenv("FILEMAGE_API_TOKEN"))
    host := os.Getenv("FILEMAGE_HOST")
    if len(host) > 0 {
        cfg.Servers[0].URL = host
    }
    client := fm.NewAPIClient(cfg)

    fm_DeleteKey(client.KeysAPI)
}

func jsonPrettyPrint(in string) string {
        var out bytes.Buffer
        err := json.Indent(&out, []byte(in), "", "  ")
        if err != nil {
                return in
        }
        return out.String()
}

```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | ID of user to remove key from | 
**keyId** | **int32** | ID of key to remove | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

