// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// TaskServiceCreateGrantResponse - The TaskServiceCreateGrantResponse message.
type TaskServiceCreateGrantResponse struct {
	// The TaskView message.
	TaskView *TaskView `json:"taskView,omitempty"`
	// The expanded field.
	Expanded []map[string]interface{} `json:"expanded,omitempty"`
}

func (o *TaskServiceCreateGrantResponse) GetTaskView() *TaskView {
	if o == nil {
		return nil
	}
	return o.TaskView
}

func (o *TaskServiceCreateGrantResponse) GetExpanded() []map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.Expanded
}
