# EndpointsAPI

All URIs are relative to *https://filemage.ppfa.org:8443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEndpoint**](EndpointsAPI.md#CreateEndpoint) | **Post** /endpoints/ | Create endpoint
[**DeleteEndpoint**](EndpointsAPI.md#DeleteEndpoint) | **Delete** /endpoints/{id}/ | Delete endpoint
[**GetEndpoint**](EndpointsAPI.md#GetEndpoint) | **Get** /endpoints/{id}/ | Get endpoint
[**ListEndpoints**](EndpointsAPI.md#ListEndpoints) | **Get** /endpoints/ | List endpoints
[**UpdateEndpoint**](EndpointsAPI.md#UpdateEndpoint) | **Put** /endpoints/{id}/ | Update endpoint



## CreateEndpoint

> EndpointCreated CreateEndpoint(ctx).Body(body).Execute()

Create endpoint



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

func fm_CreateEndpoint(client fm.EndpointsAPI) {
    body := fm.createEndpoint_request{BlobEndpoint: fm.NewBlobEndpoint()} // CreateEndpointRequest | Endpoint configuration. Check cloud provider specific schemas for config parameters.

    resp, _, err := client.CreateEndpoint(context.Background()).Body(body).Execute()
    if err != nil {
        log.Errorf("CreateEndpoint: %s", err.Error())
        return
    }
    // response from `CreateEndpoint`: EndpointCreated
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

    fm_CreateEndpoint(client.EndpointsAPI)
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

Other parameters are passed through a pointer to a apiCreateEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateEndpointRequest**](CreateEndpointRequest.md) | Endpoint configuration. Check cloud provider specific schemas for config parameters. | 

### Return type

[**EndpointCreated**](EndpointCreated.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEndpoint

> DeleteEndpoint(ctx, id).Execute()

Delete endpoint

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

func fm_DeleteEndpoint(client fm.EndpointsAPI) {
    id := int32(56) // int32 | ID of endpoint to delete

    _, err := client.DeleteEndpoint(context.Background(), id).Execute()
    if err != nil {
        log.Errorf("DeleteEndpoint: %s", err.Error())
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

    fm_DeleteEndpoint(client.EndpointsAPI)
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
**id** | **int32** | ID of endpoint to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEndpointRequest struct via the builder pattern


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


## GetEndpoint

> CreateEndpointRequest GetEndpoint(ctx, id).Execute()

Get endpoint



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

func fm_GetEndpoint(client fm.EndpointsAPI) {
    id := int32(56) // int32 | ID of endpoint to return

    resp, _, err := client.GetEndpoint(context.Background(), id).Execute()
    if err != nil {
        log.Errorf("GetEndpoint: %s", err.Error())
        return
    }
    // response from `GetEndpoint`: CreateEndpointRequest
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

    fm_GetEndpoint(client.EndpointsAPI)
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
**id** | **int32** | ID of endpoint to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateEndpointRequest**](CreateEndpointRequest.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEndpoints

> []EndpointInfo ListEndpoints(ctx).Execute()

List endpoints



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

func fm_ListEndpoints(client fm.EndpointsAPI) {

    resp, _, err := client.ListEndpoints(context.Background()).Execute()
    if err != nil {
        log.Errorf("ListEndpoints: %s", err.Error())
        return
    }
    // response from `ListEndpoints`: []EndpointInfo
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

    fm_ListEndpoints(client.EndpointsAPI)
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

Other parameters are passed through a pointer to a apiListEndpointsRequest struct via the builder pattern


### Return type

[**[]EndpointInfo**](EndpointInfo.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEndpoint

> UpdateEndpoint(ctx, id).Body(body).Execute()

Update endpoint



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

func fm_UpdateEndpoint(client fm.EndpointsAPI) {
    id := int32(56) // int32 | ID of endpoint to update
    body := fm.createEndpoint_request{BlobEndpoint: fm.NewBlobEndpoint()} // CreateEndpointRequest | Endpoint configuration. Check cloud provider specific schemas for config parameters.

    _, err := client.UpdateEndpoint(context.Background(), id).Body(body).Execute()
    if err != nil {
        log.Errorf("UpdateEndpoint: %s", err.Error())
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

    fm_UpdateEndpoint(client.EndpointsAPI)
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
**id** | **int32** | ID of endpoint to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateEndpointRequest**](CreateEndpointRequest.md) | Endpoint configuration. Check cloud provider specific schemas for config parameters. | 

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

