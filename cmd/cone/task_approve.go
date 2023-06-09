package main

import (
	"github.com/spf13/cobra"

	"github.com/conductorone/cone/internal/c1api"
	"github.com/conductorone/cone/pkg/client"
	"github.com/conductorone/cone/pkg/output"
)

func approveTasksCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "approve <task-id>",
		Short: "",
		RunE:  approveTasksRun,
		Args:  cobra.ExactArgs(1),
	}

	addTaskIdFlag(cmd)
	addCommentFlag(cmd)
	addWaitFlag(cmd)
	return cmd
}

func approveTasksRun(cmd *cobra.Command, args []string) error {
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

	policyId := taskResp.GetTaskView().Task.GetPolicy().Current.Id

	approveResp, err := c.ApproveTask(ctx, taskId, comment, client.StringFromPtr(policyId))
	if err != nil {
		return err
	}

	resp := C1ApiTaskV1TaskServiceGetResponse(*approveResp)
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

type C1ApiTaskV1TaskActionsServiceApproveResponse c1api.C1ApiTaskV1TaskActionsServiceApproveResponse

func (r *C1ApiTaskV1TaskActionsServiceApproveResponse) Header() []string {
	return []string{
		"Id",
		"Name",
		"State",
		"Processing",
		"Created At",
	}
}

func (r *C1ApiTaskV1TaskActionsServiceApproveResponse) Rows() [][]string {
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
