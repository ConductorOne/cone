// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The TaskAuditWebhookTriggered message.
type TaskAuditWebhookTriggered struct {
	// The webhookId field.
	WebhookID *string `json:"webhookId,omitempty"`
	// The webhookInstanceId field.
	WebhookInstanceID *string `json:"webhookInstanceId,omitempty"`
	// The webhookName field.
	WebhookName *string `json:"webhookName,omitempty"`
	// The webhookUrl field.
	WebhookURL *string `json:"webhookUrl,omitempty"`
}

func (o *TaskAuditWebhookTriggered) GetWebhookID() *string {
	if o == nil {
		return nil
	}
	return o.WebhookID
}

func (o *TaskAuditWebhookTriggered) GetWebhookInstanceID() *string {
	if o == nil {
		return nil
	}
	return o.WebhookInstanceID
}

func (o *TaskAuditWebhookTriggered) GetWebhookName() *string {
	if o == nil {
		return nil
	}
	return o.WebhookName
}

func (o *TaskAuditWebhookTriggered) GetWebhookURL() *string {
	if o == nil {
		return nil
	}
	return o.WebhookURL
}
