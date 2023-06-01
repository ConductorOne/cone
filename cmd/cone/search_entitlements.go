package main

import (
	"context"

	"github.com/conductorone/cone/internal/c1api"
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
	entitlements []*c1api.C1ApiAppV1AppEntitlement
	expander     *client.Expander
}

func (r *ExpandedEntitlementsResponse) Header() []string {
	return []string{"Id", "Display Name", "App", "Resource Type", "Resource", "Slug", "Alias", "Description"}
}
func (r *ExpandedEntitlementsResponse) Rows() [][]string {
	rows := [][]string{}
	for _, entitlement := range r.entitlements {
		app, _ := r.expander.GetApp(client.StringFromPtr(entitlement.AppId))
		resourceType, _ := r.expander.GetResourceType(
			client.StringFromPtr(entitlement.AppId),
			client.StringFromPtr(entitlement.AppResourceTypeId),
		)
		resource, _ := r.expander.GetResource(
			client.StringFromPtr(entitlement.AppId),
			client.StringFromPtr(entitlement.AppResourceTypeId),
			client.StringFromPtr(entitlement.AppResourceId),
		)

		rows = append(rows, []string{
			client.StringFromPtr(entitlement.Id),
			client.StringFromPtr(entitlement.DisplayName),
			client.StringFromPtr(app.DisplayName),
			client.StringFromPtr(resourceType.DisplayName),
			client.StringFromPtr(resource.DisplayName),
			client.StringFromPtr(entitlement.Slug),
			client.StringFromPtr(entitlement.Alias),
			client.StringFromPtr(entitlement.Description),
		})
	}
	return rows
}
