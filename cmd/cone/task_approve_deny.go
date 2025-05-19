package main

import (
	"context"
	"errors"
	"fmt"

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
		fmt.Printf("\nStarting task approval process for task %s\n", taskId)

		taskResp, err := c.GetTask(ctx, taskId)
		if err != nil {
			return nil, err
		}
		fmt.Printf("Got task details: %+v\n", taskResp.TaskView.Task)

		var appID, entitlementID string
		if taskResp.TaskView.Task.TaskType != nil {
			if taskResp.TaskView.Task.TaskType.TaskTypeGrant != nil {
				appID = *taskResp.TaskView.Task.TaskType.TaskTypeGrant.AppID
				entitlementID = *taskResp.TaskView.Task.TaskType.TaskTypeGrant.AppEntitlementID
			} else if taskResp.TaskView.Task.TaskType.TaskTypeRevoke != nil {
				appID = *taskResp.TaskView.Task.TaskType.TaskTypeRevoke.AppID
				entitlementID = *taskResp.TaskView.Task.TaskType.TaskTypeRevoke.AppEntitlementID
			} else if taskResp.TaskView.Task.TaskType.TaskTypeCertify != nil {
				appID = *taskResp.TaskView.Task.TaskType.TaskTypeCertify.AppID
				entitlementID = *taskResp.TaskView.Task.TaskType.TaskTypeCertify.AppEntitlementID
			}
		}
		fmt.Printf("App ID: %s, Entitlement ID: %s\n", appID, entitlementID)

		if appID == "" || entitlementID == "" {
			return nil, fmt.Errorf("could not determine app ID or entitlement ID from task")
		}

		entitlement, err := c.GetEntitlement(ctx, appID, entitlementID)
		if err != nil {
			return nil, err
		}
		fmt.Printf("Got entitlement details: %+v\n", entitlement)

		approveResp, err := c.ApproveTask(ctx, taskId, comment, policyId)
		if err != nil {
			return nil, err
		}
		fmt.Printf("Task approved successfully\n")

		resourceType, err := c.GetResourceType(ctx, appID, *entitlement.AppResourceTypeID)
		if err != nil {
			fmt.Printf("Warning: Failed to get resource type details: %v\n", err)
			return approveResp.TaskView.Task, nil
		}
		fmt.Printf("Got resource type details: %+v\n", resourceType)

		if client.IsAWSPermissionSet(entitlement, resourceType) {
			fmt.Printf("Detected AWS permission set, getting resource details...\n")
			resource, err := c.GetResource(ctx, appID, *entitlement.AppResourceTypeID, *entitlement.AppResourceID)
			if err != nil {
				fmt.Printf("Warning: Failed to get resource details: %v\n", err)
				return approveResp.TaskView.Task, nil
			}

			if err := client.CreateAWSSSOProfile(entitlement, resource); err != nil {
				fmt.Printf("Warning: Failed to create AWS SSO profile: %v\n", err)
			} else {
				fmt.Printf("Successfully created AWS SSO profile for entitlement %s\n", *entitlement.DisplayName)
			}
		} else {
			fmt.Printf("Not an AWS permission set\n")
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
