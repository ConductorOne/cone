package main

import (
	"github.com/conductorone/cone/internal/c1api"
	"github.com/conductorone/cone/pkg/client"
	"github.com/conductorone/cone/pkg/output"
	"github.com/spf13/cobra"
)

func tasksCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "task",
		Short: "",
		RunE:  taskRun,
	}

	cmd.AddCommand(getTasksCmd())

	return cmd
}

func taskRun(cmd *cobra.Command, _ []string) error {
	return cmd.Help()
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
