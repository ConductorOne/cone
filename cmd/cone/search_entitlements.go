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
	searchResp, err := c.SearchEntitlements(ctx, &client.SearchEntitlementsFilter{
		Query:            query,
		EntitlementAlias: alias,
	})
	if err != nil {
		return err
	}

	resp := C1ApiRequestcatalogV2SearchEntitlementsResponse(*searchResp)
	outputManager := output.NewManager(ctx, v)
	err = outputManager.Output(ctx, &resp)
	if err != nil {
		return err
	}

	return nil
}

type C1ApiRequestcatalogV2SearchEntitlementsResponse c1api.C1ApiRequestcatalogV2SearchEntitlementsResponse

func (r *C1ApiRequestcatalogV2SearchEntitlementsResponse) Header() []string {
	return []string{"Id", "Display Name", "Slug", "Alias", "Description"}
}
func (r *C1ApiRequestcatalogV2SearchEntitlementsResponse) Rows() [][]string {
	rows := [][]string{}
	for _, entitlement := range r.List {
		rows = append(rows, []string{
			output.FromPtr(entitlement.Id),
			output.FromPtr(entitlement.DisplayName),
			output.FromPtr(entitlement.Slug),
			output.FromPtr(entitlement.Alias),
			output.FromPtr(entitlement.Description),
		})
	}
	return rows
}
