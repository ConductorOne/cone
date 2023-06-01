package main

import (
	"context"
	"fmt"

	"github.com/conductorone/cone/pkg/client"
	"github.com/conductorone/cone/pkg/output"
	"github.com/spf13/cobra"
)

func dropCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "drop",
		Short: "Create a revoke access ticket for an entitlement by slug",
		RunE:  runDrop,
	}

	addWaitFlag(cmd)
	addAppIdFlag(cmd)
	addEntitlementIdFlag(cmd)

	return cmd
}

func runDrop(cmd *cobra.Command, args []string) error {
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

	accessRequest, err := c.CreateRevokeTask(ctx, appId, entitlementId, client.StringFromPtr(resp.UserId))
	if err != nil {
		return err
	}

	taskResp := C1ApiTaskV1Task(*accessRequest.TaskView.Task)
	outputManager := output.NewManager(ctx, v)
	err = outputManager.Output(ctx, &taskResp)
	if err != nil {
		return err
	}

	return nil
}
