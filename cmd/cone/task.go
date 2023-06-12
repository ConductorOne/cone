package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/conductorone/cone/internal/c1api"
	"github.com/conductorone/cone/pkg/client"
)

func tasksCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "task",
		Short: "",
		RunE:  taskRun,
	}

	cmd.AddCommand(getTasksCmd())
	cmd.AddCommand(searchTasksCmd())
	cmd.AddCommand(tasksCommentCmd())
	cmd.AddCommand(approveTasksCmd())
	cmd.AddCommand(denyTasksCmd())

	return cmd
}

func taskRun(cmd *cobra.Command, _ []string) error {
	return cmd.Help()
}

func (r *C1ApiTaskV1Task) Pretext() string {
	return fmt.Sprintf("Ticket URL: %s/task/%s", r.client.BaseURL(), client.StringFromPtr(r.task.NumericId))
}

type C1ApiTaskV1Task struct {
	task   *c1api.C1ApiTaskV1Task
	client client.C1Client
}

func (r *C1ApiTaskV1Task) Header() []string {
	return []string{"Id", "Display Name", "State", "Processing"}
}
func (r *C1ApiTaskV1Task) Rows() [][]string {
	return [][]string{{
		client.StringFromPtr(r.task.NumericId),
		client.StringFromPtr(r.task.DisplayName),
		taskStateToString[client.StringFromPtr(r.task.State)],
		processStateToString[client.StringFromPtr(r.task.Processing)],
	}}
}
