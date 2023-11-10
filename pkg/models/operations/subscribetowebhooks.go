// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/test-current-go/pkg/models/shared"
	"net/http"
)

type Webhook string

const (
	WebhookStockUpdate Webhook = "stockUpdate"
)

func (e Webhook) ToPointer() *Webhook {
	return &e
}

func (e *Webhook) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "stockUpdate":
		*e = Webhook(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Webhook: %v", v)
	}
}

type RequestBody struct {
	URL     *string  `json:"url,omitempty"`
	Webhook *Webhook `json:"webhook,omitempty"`
}

func (o *RequestBody) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *RequestBody) GetWebhook() *Webhook {
	if o == nil {
		return nil
	}
	return o.Webhook
}

type SubscribeToWebhooksResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// An unknown error occurred interacting with the API.
	Error *shared.Error
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *SubscribeToWebhooksResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SubscribeToWebhooksResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *SubscribeToWebhooksResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SubscribeToWebhooksResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
