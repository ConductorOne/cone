package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func hasCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "has <app-id> <app-entitlement-id>",
		Short: "Check if the current user ",
		RunE:  hasRun,
	}

	return cmd
}

func hasRun(cmd *cobra.Command, args []string) error {
	ctx, c, v, err := cmdContext(cmd)
	if err != nil {
		return err
	}

	if len(args) != 2 {
		return fmt.Errorf("expected 2 argument, got %d", len(args))
	}

	userIntro, err := c.AuthIntrospect(ctx)
	if err != nil {
		return err
	}

	appID := args[0]
	entitlementID := args[1]

	introspectResp, err := c.GetGrantsForIdentity(ctx, appID, entitlementID, *userIntro.UserID)
	if err != nil {
		return err
	}

	for x := range introspectResp {
		fmt.Println(introspectResp[x].AppEntitlementID)
	}
	print(v)

	/*
		resp := User(*userResp)
		outputManager := output.NewManager(ctx, v)
		err = outputManager.Output(ctx, &resp)
		if err != nil {
			return err
		}
	*/
	return nil

}
