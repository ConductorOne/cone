package main

import (
	"github.com/spf13/cobra"

	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/cone/pkg/client"
	"github.com/conductorone/cone/pkg/output"
)

func getTasksCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get <task-id>",
		Short: "Gets a task by id",
		RunE:  getTaskRun,
	}

	return cmd
}

func getTaskRun(cmd *cobra.Command, args []string) error {
	ctx, c, v, err := cmdContext(cmd)
	if err != nil {
		return err
	}

	if err := validateArgLenth(1, args, cmd); err != nil {
		return err
	}

	taskId := args[0]
	taskResp, err := c.GetTask(ctx, taskId)
	if err != nil {
		return err
	}

	resp := TaskGetResponse(*taskResp.TaskView.Task)
	outputManager := output.NewManager(ctx, v)
	err = outputManager.Output(ctx, &resp, output.WithTransposeTable())
	if err != nil {
		return err
	}

	return nil
}

type TaskGetResponse shared.Task

func (r *TaskGetResponse) Header() []string {
	return []string{
		"Id",
		"Name",
		"State",
		"Processing",
		"Created At",
	}
}

func (r *TaskGetResponse) WideHeader() []string {
	return append(r.Header(), "Emergency Access")
}
func (r *TaskGetResponse) rows() []string {
	return []string{
		client.StringFromPtr(r.NumericID),
		client.StringFromPtr(r.DisplayName),
		taskStateToString[*r.State],
		processStateToString[*r.Processing],
		output.FormatTime(r.CreatedAt),
	}
}
func (r *TaskGetResponse) Rows() [][]string {
	return [][]string{r.rows()}
}
func (r *TaskGetResponse) WideRows() [][]string {
	var emergencyAccess string
	if r.EmergencyAccess != nil && *r.EmergencyAccess {
		emergencyAccess = output.Checkmark
	} else {
		emergencyAccess = output.Unchecked
	}
	return [][]string{append(r.rows(), emergencyAccess)}
}
