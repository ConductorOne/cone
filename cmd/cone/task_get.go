package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/conductorone/cone/internal/c1api"
	"github.com/conductorone/cone/pkg/client"
	"github.com/conductorone/cone/pkg/output"
)

func getTasksCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Gets a task by id",
		RunE:  getTaskRun,
	}

	addTaskIdFlag(cmd)
	return cmd
}

func getTaskRun(cmd *cobra.Command, args []string) error {
	ctx, c, v, err := cmdContext(cmd)
	if err != nil {
		return err
	}

	if len(args) != 1 {
		return fmt.Errorf("expected 1 argument, got %d", len(args))
	}

	taskID := args[0]

	taskResp, err := c.GetTask(ctx, taskID)
	if err != nil {
		return err
	}

	resp := C1ApiTaskV1TaskServiceGetResponse(*taskResp)
	outputManager := output.NewManager(ctx, v)
	err = outputManager.Output(ctx, &resp)
	if err != nil {
		return err
	}

	return nil
}

type C1ApiTaskV1TaskServiceGetResponse c1api.C1ApiTaskV1TaskServiceGetResponse

func (r *C1ApiTaskV1TaskServiceGetResponse) Header() []string {
	return []string{
		"Id",
		"Name",
		"State",
		"Processing",
		"Created At",
	}
}

func (r *C1ApiTaskV1TaskServiceGetResponse) Rows() [][]string {
	return [][]string{
		{
			client.StringFromPtr(r.TaskView.Task.NumericId),
			client.StringFromPtr(r.TaskView.Task.DisplayName),
			client.StringFromPtr(r.TaskView.Task.State),
			client.StringFromPtr(r.TaskView.Task.Processing),
			output.FormatTime(r.TaskView.Task.CreatedAt),
		},
	}
}
