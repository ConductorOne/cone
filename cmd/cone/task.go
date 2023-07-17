package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/cone/pkg/client"
	"github.com/conductorone/cone/pkg/output"
)

func tasksCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "task",
		Short: "A group of commands related to interacting with tasks directly.",
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

func (r *Task) Pretext() string {
	return fmt.Sprintf("Ticket URL: %s/task/%s", r.client.BaseURL(), client.StringFromPtr(r.task.NumericID))
}

type Task struct {
	task   *shared.Task
	client client.C1Client
}

func (r *Task) Header() []string {
	return []string{"Id", "Display Name", "State", "Processing"}
}

func (r *Task) WideHeader() []string {
	return append(r.Header(), "Emergency Access")
}

func (r *Task) rows() []string {
	return []string{
		client.StringFromPtr(r.task.NumericID),
		client.StringFromPtr(r.task.DisplayName),
		taskStateToString[*r.task.State],
		processStateToString[*r.task.Processing],
	}
}
func (r *Task) Rows() [][]string {
	return [][]string{r.rows()}
}
func (r *Task) WideRows() [][]string {
	var emergencyAccess string
	if r.task.EmergencyAccess != nil && *r.task.EmergencyAccess {
		emergencyAccess = output.Checkmark
	} else {
		emergencyAccess = output.Unchecked
	}
	return [][]string{append(r.rows(), emergencyAccess)}
}
