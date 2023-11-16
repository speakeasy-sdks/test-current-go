# Authentication
(*Authentication*)

## Overview

The authentication endpoints.

### Available Operations

* [Authenticate](#authenticate) - Authenticate with the API by providing a username and password.

## Authenticate

Authenticate with the API by providing a username and password.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/test-current-go/pkg/models/shared"
	testcurrentgo "github.com/speakeasy-sdks/test-current-go"
	"context"
	"github.com/speakeasy-sdks/test-current-go/pkg/models/operations"
	"log"
)

func main() {
    s := testcurrentgo.New(
        testcurrentgo.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Authentication.Authenticate(ctx, operations.AuthenticateRequestBody{})
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.AuthenticateRequestBody](../../pkg/models/operations/authenticaterequestbody.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.AuthenticateResponse](../../pkg/models/operations/authenticateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 5XX                | application/json   |
| sdkerrors.SDKError | 400-600            | */*                |
