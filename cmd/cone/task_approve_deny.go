package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

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
	return runApproveDeny(cmd, args, func(c client.C1Client, ctx context.Context, taskId string, comment string, policyId string, v *viper.Viper) (*shared.Task, error) {
		// Only show info message for non-JSON output to avoid corrupting structured output
		outputFormat := v.GetString("output")
		if outputFormat != output.JSON && outputFormat != output.JSONPretty {
			pterm.Info.Printf("Starting task approval process for task %s\n", taskId)
		}

		taskResp, err := c.GetTask(ctx, taskId)
		if err != nil {
			return nil, err
		}
		// Only show debug message for non-JSON output to avoid corrupting structured output
		if outputFormat != output.JSON && outputFormat != output.JSONPretty {
			pterm.Debug.Printf("Got task details: %+v\n", taskResp.TaskView.Task)
		}

		var appID, entitlementID string
		switch {
		case taskResp.TaskView.Task.TaskType.TaskTypeGrant != nil:
			// Handle grant task
			grantTask := taskResp.TaskView.Task.TaskType.TaskTypeGrant
			if grantTask.AppID == nil || grantTask.AppEntitlementID == nil {
				return nil, fmt.Errorf("grant task is missing required AppID or AppEntitlementID")
			}
			appID = *grantTask.AppID
			entitlementID = *grantTask.AppEntitlementID
		case taskResp.TaskView.Task.TaskType.TaskTypeRevoke != nil:
			// Handle revoke task
			revokeTask := taskResp.TaskView.Task.TaskType.TaskTypeRevoke
			if revokeTask.AppID == nil || revokeTask.AppEntitlementID == nil {
				return nil, fmt.Errorf("revoke task is missing required AppID or AppEntitlementID")
			}
			appID = *revokeTask.AppID
			entitlementID = *revokeTask.AppEntitlementID
		default:
			return nil, fmt.Errorf("unsupported task type")
		}
		pterm.Debug.Printf("App ID: %s, Entitlement ID: %s\n", appID, entitlementID)

		if appID == "" || entitlementID == "" {
			return nil, fmt.Errorf("could not determine app ID or entitlement ID from task")
		}

		entitlement, err := c.GetEntitlement(ctx, appID, entitlementID)
		if err != nil {
			return nil, err
		}
		pterm.Debug.Printf("Got entitlement details: %+v\n", entitlement)

		approveResp, err := c.ApproveTask(ctx, taskId, comment, policyId)
		if err != nil {
			return nil, err
		}
		pterm.Success.Println("Task approved successfully")

		// Check for nil pointers before dereferencing
		if entitlement.AppResourceTypeID == nil {
			pterm.Warning.Println("Entitlement AppResourceTypeID is nil, skipping AWS SSO profile creation")
			return approveResp.TaskView.Task, nil
		}

		resourceType, err := c.GetResourceType(ctx, appID, *entitlement.AppResourceTypeID)
		if err != nil {
			pterm.Warning.Printf("Failed to get resource type details: %v\n", err)
			return approveResp.TaskView.Task, nil
		}
		pterm.Debug.Printf("Got resource type details: %+v\n", resourceType)

		if client.IsAWSPermissionSet(entitlement, resourceType) {
			pterm.Info.Println("Detected AWS permission set, getting resource details...")

			if entitlement.AppResourceID == nil {
				pterm.Warning.Println("Entitlement AppResourceID is nil, cannot create AWS SSO profile")
				return approveResp.TaskView.Task, nil
			}

			resource, err := c.GetResource(ctx, appID, *entitlement.AppResourceTypeID, *entitlement.AppResourceID)
			if err != nil {
				pterm.Warning.Printf("Failed to get resource details: %v\n", err)
				return approveResp.TaskView.Task, nil
			}

			if err := client.CreateAWSSSOProfile(entitlement, resource); err != nil {
				pterm.Warning.Printf("Failed to create AWS SSO profile: %v\n", err)
			} else {
				pterm.Success.Printf("Successfully created AWS SSO profile for entitlement %s\n", *entitlement.DisplayName)
			}
		} else {
			pterm.Debug.Println("Not an AWS permission set")
		}

		return approveResp.TaskView.Task, nil
	})
}

func runDenyTasks(cmd *cobra.Command, args []string) error {
	return runApproveDeny(cmd, args, func(c client.C1Client, ctx context.Context, taskId string, comment string, policyId string, v *viper.Viper) (*shared.Task, error) {
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
	run func(c client.C1Client, ctx context.Context, taskId string, comment string, policyId string, v *viper.Viper) (*shared.Task, error),
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

	task, err := run(c, ctx, taskId, comment, client.StringFromPtr(taskResp.TaskView.Task.PolicyInstance.PolicyStepInstance.ID), v)
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
