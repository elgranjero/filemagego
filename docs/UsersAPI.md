# UsersAPI

All URIs are relative to *https://filemage.ppfa.org:8443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatUser**](UsersAPI.md#CreatUser) | **Post** /users/ | Create user
[**DeleteUser**](UsersAPI.md#DeleteUser) | **Delete** /users/{id}/ | Delete user
[**GetUser**](UsersAPI.md#GetUser) | **Get** /users/{id}/ | Get user
[**ListUsers**](UsersAPI.md#ListUsers) | **Get** /users/ | List users
[**UpdateUser**](UsersAPI.md#UpdateUser) | **Put** /users/{id}/ | Update user



## CreatUser

> UserId CreatUser(ctx).Body(body).Execute()

Create user

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

func fm_CreatUser(client fm.UsersAPI) {
    body := *fm.NewUser() // User | User settings

    resp, _, err := client.CreatUser(context.Background()).Body(body).Execute()
    if err != nil {
        log.Errorf("CreatUser: %s", err.Error())
        return
    }
    // response from `CreatUser`: UserId
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

    fm_CreatUser(client.UsersAPI)
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

Other parameters are passed through a pointer to a apiCreatUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**User**](User.md) | User settings | 

### Return type

[**UserId**](UserId.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> DeleteUser(ctx, id).Execute()

Delete user

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

func fm_DeleteUser(client fm.UsersAPI) {
    id := int32(56) // int32 | ID of user to delete

    _, err := client.DeleteUser(context.Background(), id).Execute()
    if err != nil {
        log.Errorf("DeleteUser: %s", err.Error())
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

    fm_DeleteUser(client.UsersAPI)
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
**id** | **int32** | ID of user to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


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


## GetUser

> UserDetail GetUser(ctx, id).Execute()

Get user



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

func fm_GetUser(client fm.UsersAPI) {
    id := int32(56) // int32 | ID of user to return

    resp, _, err := client.GetUser(context.Background(), id).Execute()
    if err != nil {
        log.Errorf("GetUser: %s", err.Error())
        return
    }
    // response from `GetUser`: UserDetail
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

    fm_GetUser(client.UsersAPI)
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
**id** | **int32** | ID of user to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserDetail**](UserDetail.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsers

> []UserInfo ListUsers(ctx).Execute()

List users



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

func fm_ListUsers(client fm.UsersAPI) {

    resp, _, err := client.ListUsers(context.Background()).Execute()
    if err != nil {
        log.Errorf("ListUsers: %s", err.Error())
        return
    }
    // response from `ListUsers`: []UserInfo
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

    fm_ListUsers(client.UsersAPI)
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

Other parameters are passed through a pointer to a apiListUsersRequest struct via the builder pattern


### Return type

[**[]UserInfo**](UserInfo.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> UpdateUser(ctx, id).Body(body).Execute()

Update user



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

func fm_UpdateUser(client fm.UsersAPI) {
    id := int32(56) // int32 | ID of user to update
    body := *fm.NewUser() // User | User configuration

    _, err := client.UpdateUser(context.Background(), id).Body(body).Execute()
    if err != nil {
        log.Errorf("UpdateUser: %s", err.Error())
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

    fm_UpdateUser(client.UsersAPI)
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
**id** | **int32** | ID of user to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**User**](User.md) | User configuration | 

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

