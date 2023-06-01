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

type ExpandedEntitlement struct {
	Entitlement     *c1api.C1ApiAppV1AppEntitlement
	AppResource     *c1api.C1ApiAppV1AppResource
	AppResourceType *c1api.C1ApiAppV1AppResourceType
	App             *c1api.C1ApiAppV1App
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
	searchResp, err := c.SearchEntitlements(ctx, &client.SearchEntitlementsFilter{
		Query:            query,
		EntitlementAlias: alias,
	})
	if err != nil {
		return err
	}

	entitlements := make([]ExpandedEntitlement, 0)
	for _, item := range searchResp.List {
		app, err := c.GetApp(ctx, *item.AppEntitlement.AppId)
		if err != nil {
			return err
		}

		resourceType, err := c.GetResourceType(ctx, *item.AppEntitlement.AppId, *item.AppEntitlement.AppResourceTypeId)
		if err != nil {
			return err
		}

		resource, err := c.GetResource(ctx, *item.AppEntitlement.AppId, *item.AppEntitlement.AppResourceTypeId, *item.AppEntitlement.AppResourceId)
		if err != nil {
			return err
		}

		entitlements = append(entitlements, ExpandedEntitlement{
			Entitlement:     item.AppEntitlement,
			App:             app,
			AppResource:     resource.AppResourceView.AppResource,
			AppResourceType: resourceType.AppResourceTypeView.AppResourceType,
		})
	}

	resp := ExpandedEntitlementsResponse(entitlements)
	outputManager := output.NewManager(ctx, v)
	err = outputManager.Output(ctx, &resp)
	if err != nil {
		return err
	}

	return nil
}

type ExpandedEntitlementsResponse []ExpandedEntitlement

func (r ExpandedEntitlementsResponse) Header() []string {
	return []string{"Id", "Display Name", "App", "Resource", "Resource Type", "Slug", "Alias", "Description"}
}
func (r ExpandedEntitlementsResponse) Rows() [][]string {
	rows := [][]string{}
	for _, entitlement := range r {
		rows = append(rows, []string{
			output.FromPtr(entitlement.Entitlement.Id),
			output.FromPtr(entitlement.Entitlement.DisplayName),
			output.FromPtr(entitlement.App.DisplayName),
			output.FromPtr(entitlement.AppResource.DisplayName),
			output.FromPtr(entitlement.AppResourceType.DisplayName),
			output.FromPtr(entitlement.Entitlement.Slug),
			output.FromPtr(entitlement.Entitlement.Alias),
			output.FromPtr(entitlement.Entitlement.Description),
		})
	}
	return rows
}
