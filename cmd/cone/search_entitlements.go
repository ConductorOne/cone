package main

import (
	"context"

	"github.com/conductorone/cone/pkg/client"
	"github.com/conductorone/cone/pkg/output"
	"github.com/spf13/cobra"
)

func searchEntitlementsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "search",
		Short: "",
		RunE:  searchEntitlementsRun,
	}
	addEntitlementAliasFlag(cmd)
	addQueryFlag(cmd)
	return cmd
}

func searchEntitlementsRun(cmd *cobra.Command, args []string) error {
	ctx := context.Background()

	v, err := getSubViperForProfile(cmd)
	if err != nil {
		return err
	}

	clientId, clientSecret, err := getCredentials(v)
	if err != nil {
		return err
	}

	query := v.GetString(queryFlag)
	alias := v.GetString(entitlementAliasFlag)

	c, err := client.New(ctx, clientId, clientSecret)
	if err != nil {
		return err
	}

	// TODO(morgabra) 2-phase search: Accept a positional arg:
	// 1. Test if it's a direct alias
	// 2. Use it as a query
	entitlements, err := c.SearchEntitlements(ctx, &client.SearchEntitlementsFilter{
		Query:            query,
		EntitlementAlias: alias,
	})
	if err != nil {
		return err
	}

	expander, err := c.ExpandEntitlements(ctx, entitlements)
	if err != nil {
		return err
	}

	resp := &ExpandedEntitlementsResponse{
		entitlements: entitlements,
		expander:     expander,
	}
	outputManager := output.NewManager(ctx, v)
	err = outputManager.Output(ctx, resp)
	if err != nil {
		return err
	}

	return nil
}

type ExpandedEntitlementsResponse struct {
	entitlements []*client.EntitlementWithBindings
	expander     *client.Expander
}

func (r *ExpandedEntitlementsResponse) Header() []string {
	return []string{"Granted", "Id", "Display Name", "App", "Resource Type", "Resource", "Alias", "Description"}
}
func (r *ExpandedEntitlementsResponse) Rows() [][]string {
	rows := [][]string{}
	for _, e := range r.entitlements {
		app, _ := r.expander.GetApp(client.StringFromPtr(e.Entitlement.AppId))
		resourceType, _ := r.expander.GetResourceType(
			client.StringFromPtr(e.Entitlement.AppId),
			client.StringFromPtr(e.Entitlement.AppResourceTypeId),
		)
		resource, _ := r.expander.GetResource(
			client.StringFromPtr(e.Entitlement.AppId),
			client.StringFromPtr(e.Entitlement.AppResourceTypeId),
			client.StringFromPtr(e.Entitlement.AppResourceId),
		)

		granted := "✅"
		if len(e.Bindings) == 0 {
			granted = "❌"
		}

		rows = append(rows, []string{
			granted,
			client.StringFromPtr(e.Entitlement.Id),
			client.StringFromPtr(e.Entitlement.DisplayName),
			client.StringFromPtr(app.DisplayName),
			client.StringFromPtr(resourceType.DisplayName),
			client.StringFromPtr(resource.DisplayName),
			client.StringFromPtr(e.Entitlement.Slug),
			client.StringFromPtr(e.Entitlement.Alias),
			client.StringFromPtr(e.Entitlement.Description),
		})
	}
	return rows
}
