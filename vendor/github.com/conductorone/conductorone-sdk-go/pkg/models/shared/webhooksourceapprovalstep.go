// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The WebhookSourceApprovalStep message.
type WebhookSourceApprovalStep struct {
	// The ticketId field.
	TicketID *string `json:"ticketId,omitempty"`
}

func (o *WebhookSourceApprovalStep) GetTicketID() *string {
	if o == nil {
		return nil
	}
	return o.TicketID
}
