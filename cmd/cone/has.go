package main

import (
	"fmt"

	"github.com/conductorone/cone/pkg/client"
	"github.com/conductorone/cone/pkg/output"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// TODO: support an entitlment id + app id pair
func hasCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "has",
		Short: "",
		RunE:  hasRun,
	}
	addEntitlementAliasFlag(cmd)
	addQueryFlag(cmd)
	return cmd
}

func hasRun(cmd *cobra.Command, args []string) error {
	ctx, c, v, err := cmdContext(cmd)
	if err != nil {
		return err
	}

	query := v.GetString(queryFlag)
	alias := v.GetString(entitlementAliasFlag)

	entitlements, err := c.SearchEntitlements(ctx, &client.SearchEntitlementsFilter{
		Query:            query,
		EntitlementAlias: alias,
	})
	if err != nil {
		return err
	}
	selectedEntitlement := &client.EntitlementWithBindings{}
	if len(entitlements) == 0 {
		return fmt.Errorf("no entitlements found")
	}
	if len(entitlements) == 1 {
		selectedEntitlement = entitlements[0]
	}
	if len(entitlements) > 1 {
		isNonInteractive := v.GetBool("non-interactive")
		if isNonInteractive {
			return multipleEntitlmentsFoundError(alias, query)
		}
		optionToEntitlementMap := make(map[string]*client.EntitlementWithBindings)
		entitlementOptions := make([]string, len(entitlements))
		for _, e := range entitlements {
			entitlementOptionName := fmt.Sprintf("%s:%s:%s",
				client.StringFromPtr(e.Entitlement.DisplayName),
				client.StringFromPtr(e.Entitlement.AppId),
				client.StringFromPtr(e.Entitlement.Id),
			)
			entitlementOptions = append(entitlementOptions, entitlementOptionName)
			optionToEntitlementMap[entitlementOptionName] = e
		}
		selectedOption, _ := pterm.DefaultInteractiveSelect.WithOptions(entitlementOptions).WithDefaultText("Please select an entitlement").Show()
		selectedEntitlement = optionToEntitlementMap[selectedOption]
	}

	if len(selectedEntitlement.Bindings) == 0 {
		return fmt.Errorf("you do not have access to this entitlement")
	}

	expander, err := c.ExpandEntitlements(ctx, []*client.EntitlementWithBindings{selectedEntitlement})
	if err != nil {
		return err
	}

	resp := &HasResponse{
		entitlement: selectedEntitlement,
		expander:    expander,
	}
	outputManager := output.NewManager(ctx, v)
	err = outputManager.Output(ctx, resp)
	if err != nil {
		return err
	}

	return nil
}

type HasResponse struct {
	entitlement *client.EntitlementWithBindings
	expander    *client.Expander
}

func (r *HasResponse) Header() []string {
	return []string{"Granted", "Id", "Display Name", "App", "Resource Type", "Resource", "Alias", "Description"}
}
func (r *HasResponse) Rows() [][]string {
	rows := [][]string{}

	app, _ := r.expander.GetApp(client.StringFromPtr(r.entitlement.Entitlement.AppId))
	resourceType, _ := r.expander.GetResourceType(
		client.StringFromPtr(r.entitlement.Entitlement.AppId),
		client.StringFromPtr(r.entitlement.Entitlement.AppResourceTypeId),
	)
	resource, _ := r.expander.GetResource(
		client.StringFromPtr(r.entitlement.Entitlement.AppId),
		client.StringFromPtr(r.entitlement.Entitlement.AppResourceTypeId),
		client.StringFromPtr(r.entitlement.Entitlement.AppResourceId),
	)

	granted := "✅"
	if len(r.entitlement.Bindings) == 0 {
		granted = "❌"
	}

	rows = append(rows, []string{
		granted,
		client.StringFromPtr(r.entitlement.Entitlement.Id),
		client.StringFromPtr(r.entitlement.Entitlement.DisplayName),
		client.StringFromPtr(app.DisplayName),
		client.StringFromPtr(resourceType.DisplayName),
		client.StringFromPtr(resource.DisplayName),
		client.StringFromPtr(r.entitlement.Entitlement.Alias),
		client.StringFromPtr(r.entitlement.Entitlement.Description),
	})
	return rows
}
