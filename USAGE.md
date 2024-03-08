<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	testcurrentgo "github.com/speakeasy-sdks/test-current-go"
	"github.com/speakeasy-sdks/test-current-go/pkg/models/operations"
	"log"
)

func main() {
	s := testcurrentgo.New()

	ctx := context.Background()
	res, err := s.Drinks.ListDrinks(ctx, operations.ListDrinksRequest{})
	if err != nil {
		log.Fatal(err)
	}
	if res.Classes != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->