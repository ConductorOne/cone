// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The WebhooksServiceTestResponse message.
type WebhooksServiceTestResponse struct {
	// The WebhookInstance message.
	WebhookInstance *WebhookInstance `json:"webhook,omitempty"`
}

func (o *WebhooksServiceTestResponse) GetWebhookInstance() *WebhookInstance {
	if o == nil {
		return nil
	}
	return o.WebhookInstance
}
