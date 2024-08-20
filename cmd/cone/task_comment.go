package main

import (
	"github.com/spf13/cobra"

	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"

	"github.com/conductorone/cone/pkg/client"
	"github.com/conductorone/cone/pkg/output"
)

func tasksCommentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "comment <task-id> <comment>",
		Short: "Adds the specified comment to a task",
		RunE:  tasksCommentRun,
	}

	return cmd
}

func tasksCommentRun(cmd *cobra.Command, args []string) error {
	ctx, c, v, err := cmdContext(cmd)
	if err != nil {
		return err
	}

	if err := validateArgLenth(2, args, cmd); err != nil {
		return err
	}

	taskId := args[0]
	comment := args[1]

	userResp, err := c.CommentOnTask(ctx, taskId, comment)
	if err != nil {
		return err
	}

	resp := TaskCommentResponse(*userResp)
	outputManager := output.NewManager(ctx, v)
	err = outputManager.Output(ctx, &resp)
	if err != nil {
		return err
	}

	return nil
}

type TaskCommentResponse shared.TaskActionsServiceCommentResponse

func (t *TaskCommentResponse) Header() []string {
	return []string{
		"Id",
		"Display Name",
		"State",
	}
}

func (t *TaskCommentResponse) Rows() [][]string {
	return [][]string{
		{
			client.StringFromIntPtr(t.TaskView.Task.NumericID),
			client.StringFromPtr(t.TaskView.Task.DisplayName),
			taskStateToString[*t.TaskView.Task.State],
		},
	}
}
