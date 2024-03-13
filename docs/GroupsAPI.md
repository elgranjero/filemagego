# GroupsAPI

All URIs are relative to *https://sftp.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGroup**](GroupsAPI.md#CreateGroup) | **Post** /groups/ | Create group
[**DeleteGroup**](GroupsAPI.md#DeleteGroup) | **Delete** /groups/{id}/ | Delete group
[**GetGroup**](GroupsAPI.md#GetGroup) | **Get** /groups/{id}/ | Get group
[**ListGroups**](GroupsAPI.md#ListGroups) | **Get** /groups/ | List groups
[**UpdateGroup**](GroupsAPI.md#UpdateGroup) | **Put** /groups/{id}/ | Update group



## CreateGroup

> GroupCreated CreateGroup(ctx).Body(body).Execute()

Create group

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

func fm_CreateGroup(client fm.GroupsAPI) {
    body := *fm.NewGroup() // Group | Group settings

    resp, _, err := client.CreateGroup(context.Background()).Body(body).Execute()
    if err != nil {
        log.Errorf("CreateGroup: %s", err.Error())
        return
    }
    // response from `CreateGroup`: GroupCreated
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

    fm_CreateGroup(client.GroupsAPI)
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

Other parameters are passed through a pointer to a apiCreateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Group**](Group.md) | Group settings | 

### Return type

[**GroupCreated**](GroupCreated.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroup

> DeleteGroup(ctx, id).Execute()

Delete group

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

func fm_DeleteGroup(client fm.GroupsAPI) {
    id := int32(56) // int32 | ID of group to delete

    _, err := client.DeleteGroup(context.Background(), id).Execute()
    if err != nil {
        log.Errorf("DeleteGroup: %s", err.Error())
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

    fm_DeleteGroup(client.GroupsAPI)
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
**id** | **int32** | ID of group to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupRequest struct via the builder pattern


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


## GetGroup

> Group GetGroup(ctx, id).Execute()

Get group



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

func fm_GetGroup(client fm.GroupsAPI) {
    id := int32(56) // int32 | ID of group to return

    resp, _, err := client.GetGroup(context.Background(), id).Execute()
    if err != nil {
        log.Errorf("GetGroup: %s", err.Error())
        return
    }
    // response from `GetGroup`: Group
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

    fm_GetGroup(client.GroupsAPI)
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
**id** | **int32** | ID of group to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Group**](Group.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGroups

> []GroupInfo ListGroups(ctx).Execute()

List groups



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

func fm_ListGroups(client fm.GroupsAPI) {

    resp, _, err := client.ListGroups(context.Background()).Execute()
    if err != nil {
        log.Errorf("ListGroups: %s", err.Error())
        return
    }
    // response from `ListGroups`: []GroupInfo
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

    fm_ListGroups(client.GroupsAPI)
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

Other parameters are passed through a pointer to a apiListGroupsRequest struct via the builder pattern


### Return type

[**[]GroupInfo**](GroupInfo.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroup

> UpdateGroup(ctx, id).Body(body).Execute()

Update group



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

func fm_UpdateGroup(client fm.GroupsAPI) {
    id := int32(56) // int32 | ID of group to update
    body := *fm.NewGroup() // Group | User configuration

    _, err := client.UpdateGroup(context.Background(), id).Body(body).Execute()
    if err != nil {
        log.Errorf("UpdateGroup: %s", err.Error())
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

    fm_UpdateGroup(client.GroupsAPI)
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
**id** | **int32** | ID of group to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**Group**](Group.md) | User configuration | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

