// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The WebhookRef message.
type WebhookRef struct {
	// The id field.
	ID *string `json:"id,omitempty"`
}

func (o *WebhookRef) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}
