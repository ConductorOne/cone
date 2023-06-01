package main

import (
	"context"
	"fmt"
	"time"

	"github.com/conductorone/cone/internal/c1api"
	"github.com/conductorone/cone/pkg/client"
	"github.com/conductorone/cone/pkg/output"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

func getCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Create an access request for an entitlement by slug",
		RunE:  runGet,
	}

	return taskCmd(cmd)
}

func dropCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "drop",
		Short: "Create a revoke access ticket for an entitlement by slug",
		RunE:  runDrop,
	}

	return taskCmd(cmd)
}

func taskCmd(cmd *cobra.Command) *cobra.Command {
	addWaitFlag(cmd)
	addAppIdFlag(cmd)
	addEntitlementIdFlag(cmd)
	return cmd
}

func runGet(cmd *cobra.Command, args []string) error {
	return runTask(cmd, args, func(c client.C1Client, ctx context.Context, appId string, entitlementId string, userId string) (*c1api.C1ApiTaskV1Task, error) {
		accessRequest, err := c.CreateGrantTask(ctx, appId, entitlementId, userId)
		if err != nil {
			return nil, err
		}
		return accessRequest.TaskView.Task, nil
	})
}

func runDrop(cmd *cobra.Command, args []string) error {
	return runTask(cmd, args, func(c client.C1Client, ctx context.Context, appId string, entitlementId string, userId string) (*c1api.C1ApiTaskV1Task, error) {
		accessRequest, err := c.CreateRevokeTask(ctx, appId, entitlementId, userId)
		if err != nil {
			return nil, err
		}
		return accessRequest.TaskView.Task, nil
	})
}

func runTask(cmd *cobra.Command, args []string, run func(c client.C1Client, ctx context.Context, appId string, entitlementId string, userId string) (*c1api.C1ApiTaskV1Task, error)) error {
	ctx := context.Background()
	alias := ""

	v, err := getSubViperForProfile(cmd)
	if err != nil {
		return err
	}

	clientId, clientSecret, err := getCredentials(v)
	if err != nil {
		return err
	}

	entitlementId := v.GetString(entitlementIdFlag)
	appId := v.GetString(appIdFlag)

	if len(args) == 1 {
		alias = args[0]
	}

	if alias == "" && (appId == "" || entitlementId == "") {
		return fmt.Errorf("must provide either an alias or an entitlement id and app id")
	}

	if alias != "" && (appId != "" || entitlementId != "") {
		return fmt.Errorf("cannot provide both an alias and an entitlement id and app id")
	}

	c, err := client.New(ctx, clientId, clientSecret, client.WithDebug(v.GetBool("debug")))
	if err != nil {
		return err
	}

	if alias != "" {
		entitlement, err := c.SearchEntitlements(ctx, &client.SearchEntitlementsFilter{EntitlementAlias: alias})
		if err != nil {
			return err
		}
		if len(entitlement.List) == 0 {
			return fmt.Errorf("no entitlement found with alias %s", alias)
		}
		if len(entitlement.List) > 1 {
			// TODO: this should show a list and prompt for input.
			return fmt.Errorf("multiple entitlements found with alias %s", alias)
		}
		entitlementId = client.StringFromPtr(entitlement.List[0].AppEntitlement.Id)
		appId = client.StringFromPtr(entitlement.List[0].AppEntitlement.AppId)
	}

	resp, err := c.WhoAmI(ctx)
	if err != nil {
		return err
	}

	task, err := run(c, ctx, appId, entitlementId, client.StringFromPtr(resp.UserId))
	if err != nil {
		return err
	}

	outputManager := output.NewManager(ctx, v)
	taskResp := C1ApiTaskV1Task{task: task, client: c}
	err = outputManager.Output(ctx, &taskResp)
	if err != nil {
		return err
	}

	if wait, _ := cmd.Flags().GetBool("wait"); wait {
		spinner, _ := pterm.DefaultSpinner.Start("Waiting for ticket to close.")
		for {
			time.Sleep(1 * time.Second)
			task, err := c.GetTask(ctx, client.StringFromPtr(task.Id))
			if err != nil {
				return err
			}

			taskItem := task.TaskView.Task
			taskResp = C1ApiTaskV1Task{task: taskItem, client: c}
			err = outputManager.Output(ctx, &taskResp)
			if err != nil {
				return err
			}
			if client.StringFromPtr(taskItem.State) == "TASK_STATE_CLOSED" {
				break
			}
		}
		spinner.Success("Ticket closed.")
	}

	return nil
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
		client.StringFromPtr(r.task.State),
		client.StringFromPtr(r.task.Processing),
	}}
}

func (r *C1ApiTaskV1Task) Pretext() string {
	return fmt.Sprintf("Ticket URL: %s/task/%s", r.client.BaseURL(), client.StringFromPtr(r.task.NumericId))
}
