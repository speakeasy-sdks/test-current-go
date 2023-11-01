// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package callbacks

import (
	"github.com/speakeasy-sdks/test-current-go/pkg/models/shared"
	"net/http"
)

type CreateOrderOrderUpdateResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// An unknown error occurred interacting with the API.
	Error *shared.Error
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateOrderOrderUpdateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateOrderOrderUpdateResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *CreateOrderOrderUpdateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateOrderOrderUpdateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

type CreateOrderOrderUpdateRequestBody struct {
	// An order for a drink or ingredient.
	Order *shared.OrderInput `json:"order,omitempty"`
}

func (o *CreateOrderOrderUpdateRequestBody) GetOrder() *shared.OrderInput {
	if o == nil {
		return nil
	}
	return o.Order
}