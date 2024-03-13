# PermissionsAPI

All URIs are relative to *https://filemage.ppfa.org:8443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DetailPermissions**](PermissionsAPI.md#DetailPermissions) | **Get** /permissions/query/ | Folder permission detail
[**ViewPermissions**](PermissionsAPI.md#ViewPermissions) | **Get** /permissions/ | Folder permissions overview



## DetailPermissions

> []PermissionDetail DetailPermissions(ctx).Path(path).Execute()

Folder permission detail



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

func fm_DetailPermissions(client fm.PermissionsAPI) {
    path := int32(56) // int32 | Folder path to list permissions of.

    resp, _, err := client.DetailPermissions(context.Background()).Path(path).Execute()
    if err != nil {
        log.Errorf("DetailPermissions: %s", err.Error())
        return
    }
    // response from `DetailPermissions`: []PermissionDetail
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

    fm_DetailPermissions(client.PermissionsAPI)
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



### Other Parameters

Other parameters are passed through a pointer to a apiDetailPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **int32** | Folder path to list permissions of. | 

### Return type

[**[]PermissionDetail**](PermissionDetail.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ViewPermissions

> []PermissionOverview ViewPermissions(ctx).Execute()

Folder permissions overview



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

func fm_ViewPermissions(client fm.PermissionsAPI) {

    resp, _, err := client.ViewPermissions(context.Background()).Execute()
    if err != nil {
        log.Errorf("ViewPermissions: %s", err.Error())
        return
    }
    // response from `ViewPermissions`: []PermissionOverview
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

    fm_ViewPermissions(client.PermissionsAPI)
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

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiViewPermissionsRequest struct via the builder pattern


### Return type

[**[]PermissionOverview**](PermissionOverview.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

