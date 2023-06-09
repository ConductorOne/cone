package main

import (
	"github.com/spf13/cobra"

	"github.com/conductorone/cone/internal/c1api"
	"github.com/conductorone/cone/pkg/client"
	"github.com/conductorone/cone/pkg/output"
)

func denyTasksCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deny <task-id>",
		Short: "",
		RunE:  denyTasksRun,
		Args:  cobra.ExactArgs(1),
	}

	addWaitFlag(cmd)
	return cmd
}

func denyTasksRun(cmd *cobra.Command, args []string) error {
	ctx, c, v, err := cmdContext(cmd)
	if err != nil {
		return err
	}

	taskId := args[0]
	comment := v.GetString(commentFlag)

	taskResp, err := c.GetTask(ctx, taskId)
	if err != nil {
		return err
	}

	task := taskResp.GetTaskView().Task
	policyId := task.GetPolicy().Current.Id

	denyResp, err := c.DenyTask(ctx, client.StringFromPtr(task.Id), comment, client.StringFromPtr(policyId))
	if err != nil {
		return err
	}

	resp := C1ApiTaskV1TaskActionsServiceDenyResponse(*denyResp)
	outputManager := output.NewManager(ctx, v)
	err = outputManager.Output(ctx, &resp)
	if err != nil {
		return err
	}

	if wait, _ := cmd.Flags().GetBool("wait"); wait {
		err = handleWaitBehavior(ctx, c, resp.TaskView.Task, outputManager)
		if err != nil {
			return err
		}
	}

	return nil
}

type C1ApiTaskV1TaskActionsServiceDenyResponse c1api.C1ApiTaskV1TaskActionsServiceDenyResponse

func (r *C1ApiTaskV1TaskActionsServiceDenyResponse) Header() []string {
	return []string{
		"Id",
		"Name",
		"State",
		"Processing",
		"Created At",
	}
}

func (r *C1ApiTaskV1TaskActionsServiceDenyResponse) Rows() [][]string {
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
