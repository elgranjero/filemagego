# AuditAPI

All URIs are relative to *https://sftp.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListLogs**](AuditAPI.md#ListLogs) | **Get** /logs/ | Get audit log entries



## ListLogs

> []AuditLog ListLogs(ctx).Start(start).End(end).Path(path).Operation(operation).User(user).Execute()

Get audit log entries



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

func fm_ListLogs(client fm.AuditAPI) {
    start := "start_example" // string | RFC 3339 timestamp of first event to return. (optional)
    end := "end_example" // string | RFC 3339 timestamp of last event to return. (optional)
    path := "path_example" // string | Show only events for the object at this path. (optional)
    operation := "operation_example" // string | The operation type to filter by. Valid values are 'cd', 'mkdir', 'ls','get','put','stat','rmdir','rm','mv' (optional)
    user := "user_example" // string | The username to filter by. (optional)

    resp, _, err := client.ListLogs(context.Background()).Start(start).End(end).Path(path).Operation(operation).User(user).Execute()
    if err != nil {
        log.Errorf("ListLogs: %s", err.Error())
        return
    }
    // response from `ListLogs`: []AuditLog
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

    fm_ListLogs(client.AuditAPI)
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

Other parameters are passed through a pointer to a apiListLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** | RFC 3339 timestamp of first event to return. | 
 **end** | **string** | RFC 3339 timestamp of last event to return. | 
 **path** | **string** | Show only events for the object at this path. | 
 **operation** | **string** | The operation type to filter by. Valid values are &#39;cd&#39;, &#39;mkdir&#39;, &#39;ls&#39;,&#39;get&#39;,&#39;put&#39;,&#39;stat&#39;,&#39;rmdir&#39;,&#39;rm&#39;,&#39;mv&#39; | 
 **user** | **string** | The username to filter by. | 

### Return type

[**[]AuditLog**](AuditLog.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

