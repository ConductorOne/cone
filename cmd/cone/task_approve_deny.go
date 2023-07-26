package main

import (
	"context"
	"errors"

	"github.com/spf13/cobra"

	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/cone/pkg/client"
	"github.com/conductorone/cone/pkg/output"
)

func approveTasksCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "approve <task-id>",
		Short: "Mark a task as approved",
		RunE:  runApproveTasks,
	}

	addCommentFlag(cmd)
	addWaitFlag(cmd)
	return cmd
}

func denyTasksCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deny <task-id>",
		Short: "Mark a task as denied",
		RunE:  runDenyTasks,
	}

	addCommentFlag(cmd)
	addWaitFlag(cmd)
	return cmd
}

func runApproveTasks(cmd *cobra.Command, args []string) error {
	return runApproveDeny(cmd, args, func(c client.C1Client, ctx context.Context, taskId string, comment string, policyId string) (*shared.Task, error) {
		approveResp, err := c.ApproveTask(ctx, taskId, comment, policyId)
		if err != nil {
			return nil, err
		}
		return approveResp.TaskView.Task, nil
	})
}

func runDenyTasks(cmd *cobra.Command, args []string) error {
	return runApproveDeny(cmd, args, func(c client.C1Client, ctx context.Context, taskId string, comment string, policyId string) (*shared.Task, error) {
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
	run func(c client.C1Client, ctx context.Context, taskId string, comment string, policyId string) (*shared.Task, error),
) error {
	ctx, c, v, err := cmdContext(cmd)
	if err != nil {
		return err
	}

	if err := validateArgLenth(1, args, cmd); err != nil {
		return err
	}

	taskId := args[0]
	comment := v.GetString(commentFlag)

	taskResp, err := c.GetTask(ctx, taskId)
	if err != nil {
		return err
	}

	if taskResp.TaskView.Task.PolicyInstance.PolicyStepInstance == nil {
		return errors.New("task does not have a current policy step id and cannot be approved or denied")
	}

	task, err := run(c, ctx, taskId, comment, client.StringFromPtr(taskResp.TaskView.Task.PolicyInstance.PolicyStepInstance.ID))
	if err != nil {
		return err
	}

	resp := Task{task: task, client: c}
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
