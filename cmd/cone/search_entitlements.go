package main

import (
	"github.com/spf13/cobra"

	"github.com/conductorone/cone/pkg/client"
	"github.com/conductorone/cone/pkg/output"
)

func searchEntitlementsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "search",
		Short: "",
		RunE:  searchEntitlementsRun,
	}
	addEntitlementAliasFlag(cmd)
	addQueryFlag(cmd)
	addGrantedFlag(cmd)
	addNotGrantedFlag(cmd)
	addAppDisplayNameFlag(cmd)
	cmd.MarkFlagsMutuallyExclusive(grantedFlag, notGrantedFlag)
	return cmd
}

func searchEntitlementsRun(cmd *cobra.Command, args []string) error {
	ctx, c, v, err := cmdContext(cmd)
	if err != nil {
		return err
	}

	query := v.GetString(queryFlag)
	alias := v.GetString(entitlementAliasFlag)
	grantedStatus := client.GrantedStatusAll
	if v.GetBool(grantedFlag) {
		grantedStatus = client.GrantedStatusGranted
	} else if v.GetBool(notGrantedFlag) {
		grantedStatus = client.GrantedStatusNotGranted
	}

	// TODO(morgabra) 2-phase search: Accept a positional arg:
	// 1. Test if it's a direct alias
	// 2. Use it as a query
	entitlements, err := c.SearchEntitlements(ctx, &client.SearchEntitlementsFilter{
		Query:            query,
		EntitlementAlias: alias,
		GrantedStatus:    grantedStatus,
		AppDisplayName:   v.GetString(appDisplayNameFlag),
	})
	if err != nil {
		return err
	}

	expander, err := c.ExpandEntitlements(ctx, entitlements)
	if err != nil {
		return err
	}

	resp := &ExpandedEntitlementsResponse{
		Entitlements: entitlements,
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
	Entitlements []*client.EntitlementWithBindings `json:"entitlements"`
	expander     *client.Expander
}

const DisplayNameHeader = "Display Name"

func (r *ExpandedEntitlementsResponse) Header() []string {
	return []string{"", "Alias", DisplayNameHeader, "App", "Resource Type", "Resource"}
}

func (r *ExpandedEntitlementsResponse) WideHeader() []string {
	return append(r.Header(), "Description", "App ID", "Entitlement ID")
}

func (r *ExpandedEntitlementsResponse) Rows() [][]string {
	rows := [][]string{}
	for _, e := range r.Entitlements {
		app, _ := r.expander.GetApp(client.StringFromPtr(e.Entitlement.AppID))
		resourceType, _ := r.expander.GetResourceType(
			client.StringFromPtr(e.Entitlement.AppID),
			client.StringFromPtr(e.Entitlement.AppResourceTypeID),
		)
		resource, _ := r.expander.GetResource(
			client.StringFromPtr(e.Entitlement.AppID),
			client.StringFromPtr(e.Entitlement.AppResourceTypeID),
			client.StringFromPtr(e.Entitlement.AppResourceID),
		)

		granted := output.Checkmark
		if len(e.Bindings) == 0 {
			granted = output.Unchecked
		}

		rows = append(rows, []string{
			granted,
			client.StringFromPtr(e.Entitlement.Alias),
			client.StringFromPtr(e.Entitlement.DisplayName),
			client.StringFromPtr(app.DisplayName),
			client.StringFromPtr(resourceType.DisplayName),
			client.StringFromPtr(resource.DisplayName),
		})
	}
	return rows
}
func (r *ExpandedEntitlementsResponse) SortByColumnName() string {
	return DisplayNameHeader
}

func (r *ExpandedEntitlementsResponse) WideRows() [][]string {
	rows := r.Rows()
	for i, e := range r.Entitlements {
		rows[i] = append(rows[i],
			client.StringFromPtr(e.Entitlement.Description),
			client.StringFromPtr(e.Entitlement.AppID),
			client.StringFromPtr(e.Entitlement.ID),
		)
	}
	return rows
}
