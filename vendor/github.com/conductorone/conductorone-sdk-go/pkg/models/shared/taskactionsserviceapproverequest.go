// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The TaskActionsServiceApproveRequest object lets you approve a task.
type TaskActionsServiceApproveRequest struct {
	// The task expand mask is an array of strings that specifes the related objects the requester wishes to have returned when making a request where the expand mask is part of the input. Use '*' to view all possible responses.
	TaskExpandMask *TaskExpandMask `json:"expandMask,omitempty"`
	// The comment attached to the request.
	Comment *string `json:"comment,omitempty"`
	// The ID of the policy step on the given task to approve.
	PolicyStepID string `json:"policyStepId"`
}

func (o *TaskActionsServiceApproveRequest) GetTaskExpandMask() *TaskExpandMask {
	if o == nil {
		return nil
	}
	return o.TaskExpandMask
}

func (o *TaskActionsServiceApproveRequest) GetComment() *string {
	if o == nil {
		return nil
	}
	return o.Comment
}

func (o *TaskActionsServiceApproveRequest) GetPolicyStepID() string {
	if o == nil {
		return ""
	}
	return o.PolicyStepID
}
