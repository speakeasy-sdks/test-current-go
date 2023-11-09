# github.com/speakeasy-sdks/test-current-go

<div align="left">
    <a href="https://speakeasyapi.dev/"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://github.com/speakeasy-sdks/test-current-go.git/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/test-current-go/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
    
</div>


## 🏗 **Welcome to your new SDK!** 🏗

It has been generated successfully based on your OpenAPI spec. However, it is not yet ready for production use. Here are some next steps:
- [ ] 🛠 Make your SDK feel handcrafted by [customizing it](https://www.speakeasyapi.dev/docs/customize-sdks)
- [ ] ♻️ Refine your SDK quickly by iterating locally with the [Speakeasy CLI](https://github.com/speakeasy-api/speakeasy)
- [ ] 🎁 Publish your SDK to package managers by [configuring automatic publishing](https://www.speakeasyapi.dev/docs/productionize-sdks/publish-sdks)
- [ ] ✨ When ready to productionize, delete this section from the README
<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/test-current-go
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import (
	"context"
	testcurrentgo "github.com/speakeasy-sdks/test-current-go"
	"github.com/speakeasy-sdks/test-current-go/pkg/models/operations"
	"github.com/speakeasy-sdks/test-current-go/pkg/models/shared"
	"log"
)

func main() {
	s := testcurrentgo.New(
		testcurrentgo.WithSecurity(""),
	)

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
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Authentication](docs/sdks/authentication/README.md)

* [Authenticate](docs/sdks/authentication/README.md#authenticate) - Authenticate with the API by providing a username and password.

### [Drinks](docs/sdks/drinks/README.md)

* [GetDrink](docs/sdks/drinks/README.md#getdrink) - Get a drink.
* [ListDrinks](docs/sdks/drinks/README.md#listdrinks) - Get a list of drinks.

### [Ingredients](docs/sdks/ingredients/README.md)

* [ListIngredients](docs/sdks/ingredients/README.md#listingredients) - Get a list of ingredients.

### [Orders](docs/sdks/orders/README.md)

* [CreateOrder](docs/sdks/orders/README.md#createorder) - Create an order.

### [Config](docs/sdks/config/README.md)

* [SubscribeToWebhooks](docs/sdks/config/README.md#subscribetowebhooks) - Subscribe to webhooks.
<!-- End SDK Available Operations -->

<!-- Start Dev Containers -->

<!-- End Dev Containers -->

<!-- Start Error Handling -->
# Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 5XX                | application/json   |
| sdkerrors.SDKError | 400-600            | */*                |


## Example

```go
package main

import (
	"context"
	testcurrentgo "github.com/speakeasy-sdks/test-current-go"
	"github.com/speakeasy-sdks/test-current-go/pkg/models/operations"
	"github.com/speakeasy-sdks/test-current-go/pkg/models/shared"
	"log"
)

func main() {
	s := testcurrentgo.New(
		testcurrentgo.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.Authentication.Authenticate(ctx, operations.AuthenticateRequestBody{})
	if err != nil {

		var e *sdkerrors.APIError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling -->

<!-- Start Server Selection -->
# Server Selection

## Select Server by Name

You can override the default server globally using the `WithServer` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the names associated with the available servers:

| Name | Server | Variables |
| ----- | ------ | --------- |
| `prod` | `https://speakeasy.bar` | None |
| `staging` | `https://staging.speakeasy.bar` | None |
| `customer` | `https://{organization}.{environment}.speakeasy.bar` | `environment` (default is `prod`), `organization` (default is `api`) |


Some of the server options above contain variables. If you want to set the values of those variables, the following options are provided for doing so:
 * `WithEnvironment testcurrentgo.ServerEnvironment`

 * `WithOrganization string`

For example:

```go
package main

import (
	"context"
	testcurrentgo "github.com/speakeasy-sdks/test-current-go"
	"github.com/speakeasy-sdks/test-current-go/pkg/models/operations"
	"github.com/speakeasy-sdks/test-current-go/pkg/models/shared"
	"log"
)

func main() {
	s := testcurrentgo.New(
		testcurrentgo.WithServer("customer"),
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


## Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:

```go
package main

import (
	"context"
	testcurrentgo "github.com/speakeasy-sdks/test-current-go"
	"github.com/speakeasy-sdks/test-current-go/pkg/models/operations"
	"github.com/speakeasy-sdks/test-current-go/pkg/models/shared"
	"log"
)

func main() {
	s := testcurrentgo.New(
		testcurrentgo.WithServerURL("https://speakeasy.bar"),
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
<!-- End Server Selection -->

<!-- Start Custom HTTP Client -->
# Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client -->

<!-- Start Go Types -->

<!-- End Go Types -->



<!-- Start Authentication -->
# Authentication

## Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name     | Type     | Scheme   |
| -------- | -------- | -------- |
| `APIKey` | apiKey   | API key  |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:

```go
package main

import (
	"context"
	testcurrentgo "github.com/speakeasy-sdks/test-current-go"
	"github.com/speakeasy-sdks/test-current-go/pkg/models/operations"
	"github.com/speakeasy-sdks/test-current-go/pkg/models/shared"
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
<!-- End Authentication -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
