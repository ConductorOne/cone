package main

import (
	"fmt"
	"os"

	"github.com/conductorone/cone/pkg/client"
	"github.com/conductorone/cone/pkg/output"
	"github.com/spf13/cobra"
)

type HasAppEntitlement struct {
	Has              string `json:"has"`
	AppId            string `json:"app_id"`
	AppEntitlementId string `json:"entitlement_id"`
	AppName          string `json:"app_name"`
	Entitlement      string `json:"entitlement"`
	UserId           string `json:"user_id"`
}

func hasCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "has <app-id> <app-entitlement-id>",
		Short: "Check if the current user has a specific entitlement for an app",
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
		usageString := "\nUsage:  cone has <app-id> <app-entitlement-id>"
		return fmt.Errorf("expected 2 arguments, got %d"+usageString, len(args))
	}

	appID := args[0]
	entitlementID := args[1]

	userIntro, err := c.AuthIntrospect(ctx)
	if err != nil {
		return err
	}
	grants, err := c.GetGrantsForIdentity(ctx, appID, entitlementID, client.StringFromPtr(userIntro.UserID))
	if err != nil {
		return err
	}
	app, err := c.GetApp(ctx, appID)
	if err != nil {
		return err
	}
	entitlement, err := c.GetEntitlement(ctx, appID, entitlementID)
	if err != nil {
		return err
	}

	hasObj := HasAppEntitlement{
		Has:              output.Unchecked,
		AppEntitlementId: entitlementID,
		AppId:            appID,
		AppName:          client.StringFromPtr(app.DisplayName),
		Entitlement:      client.StringFromPtr(entitlement.DisplayName),
		UserId:           client.StringFromPtr(userIntro.UserID),
	}

	for _, grant := range grants {
		if grant.CreatedAt != nil && grant.DeletedAt == nil {
			hasObj.Has = output.Checkmark
			break
		}
	}

	outputManager := output.NewManager(ctx, v)
	err = outputManager.Output(ctx, &hasObj, output.WithTransposeTable())
	if err != nil {
		return err
	}

	if hasObj.Has == output.Unchecked {
		os.Exit(2)
	}

	return nil
}

func (r *HasAppEntitlement) Header() []string {
	return []string{
		"Entitlement Granted",
		"App Name",
		"Entitlement",
	}
}

func (r *HasAppEntitlement) rows() []string {
	return []string{
		r.Has,
		r.AppName,
		r.Entitlement,
	}
}

func (r *HasAppEntitlement) Rows() [][]string {
	return [][]string{r.rows()}
}
