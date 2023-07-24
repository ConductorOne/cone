package main

import (
	"github.com/spf13/cobra"

	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/cone/pkg/client"
	"github.com/conductorone/cone/pkg/output"
)

func escalateTasksCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "escalate <task-id>",
		Short: "Escalate a task  blah blah blah emergency access",
		RunE:  runEscalateTasks,
		Args:  cobra.ExactArgs(1),
	}
	return cmd
}

func runEscalateTasks(cmd *cobra.Command, args []string) error {
	ctx, c, v, err := cmdContext(cmd)
	if err != nil {
		return err
	}

	taskId := args[0]

	userResp, err := c.EscalateTask(ctx, taskId)
	if err != nil {
		return err
	}

	resp := TaskServiceActionResponse(*userResp)
	outputManager := output.NewManager(ctx, v)
	err = outputManager.Output(ctx, &resp)
	if err != nil {
		return err
	}

	return nil
}

type TaskServiceActionResponse shared.TaskServiceActionResponse

func (t *TaskServiceActionResponse) Header() []string {
	return []string{
		"Id",
		"Display Name",
		"State",
	}
}

func (t *TaskServiceActionResponse) Rows() [][]string {
	return [][]string{
		{
			client.StringFromPtr(t.TaskView.Task.NumericID),
			client.StringFromPtr(t.TaskView.Task.DisplayName),
			taskStateToString[*t.TaskView.Task.State],
		},
	}
}