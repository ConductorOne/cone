package main

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/conductorone/cone/internal/c1api"
	"github.com/conductorone/cone/pkg/client"
	"github.com/conductorone/cone/pkg/output"
)

func approveTasksCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "approve <task-id>",
		Short: "",
		RunE:  runApproveTasks,
		Args:  cobra.ExactArgs(1),
	}

	addCommentFlag(cmd)
	addWaitFlag(cmd)
	return cmd
}

func denyTasksCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deny <task-id>",
		Short: "",
		RunE:  runDenyTasks,
		Args:  cobra.ExactArgs(1),
	}

	addCommentFlag(cmd)
	addWaitFlag(cmd)
	return cmd
}

func runApproveTasks(cmd *cobra.Command, args []string) error {
	return runApproveDeny(cmd, args, func(c client.C1Client, ctx context.Context, taskId string, comment string, policyId string) (*c1api.C1ApiTaskV1Task, error) {
		approveResp, err := c.ApproveTask(ctx, taskId, comment, policyId)
		if err != nil {
			return nil, err
		}
		return approveResp.TaskView.Task, nil
	})
}

func runDenyTasks(cmd *cobra.Command, args []string) error {
	return runApproveDeny(cmd, args, func(c client.C1Client, ctx context.Context, taskId string, comment string, policyId string) (*c1api.C1ApiTaskV1Task, error) {
		approveResp, err := c.DenyTask(ctx, taskId, comment, policyId)
		if err != nil {
			return nil, err
		}
		return approveResp.TaskView.Task, nil
	})
}

func runApproveDeny(
	cmd *cobra.Command,
	args []string,
	run func(c client.C1Client, ctx context.Context, taskId string, comment string, policyId string) (*c1api.C1ApiTaskV1Task, error),
) error {
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

	task, err := run(c, ctx, taskId, comment, client.StringFromPtr(policyId))
	if err != nil {
		return err
	}

	resp := C1ApiTaskV1Task{task: task, client: c}
	outputManager := output.NewManager(ctx, v)
	err = outputManager.Output(ctx, &resp)
	if err != nil {
		return err
	}

	if wait, _ := cmd.Flags().GetBool("wait"); wait {
		err = handleWaitBehavior(ctx, c, resp.task, outputManager)
		if err != nil {
			return err
		}
	}

	return nil
}
