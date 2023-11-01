# Ingredients
(*.Ingredients*)

## Overview

The ingredients endpoints.

### Available Operations

* [ListIngredients](#listingredients) - Get a list of ingredients.

## ListIngredients

Get a list of ingredients, if authenticated this will include stock levels and product codes otherwise it will only include public information.

### Example Usage

```go
package main

import(
	"context"
	"log"
	testcurrentgo "github.com/speakeasy-sdks/test-current-go"
	"github.com/speakeasy-sdks/test-current-go/pkg/models/shared"
	"github.com/speakeasy-sdks/test-current-go/pkg/models/operations"
)

func main() {
    s := testcurrentgo.New(
        testcurrentgo.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Ingredients.ListIngredients(ctx, operations.ListIngredientsRequest{
        Ingredients: []string{
            "string",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListIngredientsRequest](../../models/operations/listingredientsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.ListIngredientsResponse](../../models/operations/listingredientsresponse.md), error**
