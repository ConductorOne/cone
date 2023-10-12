package main

import (
	"github.com/spf13/cobra"

	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
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
	addIncludeDeletedFlag(cmd)
	cmd.MarkFlagsMutuallyExclusive(grantedFlag, notGrantedFlag)
	return cmd
}

func searchEntitlementsRun(cmd *cobra.Command, args []string) error {
	ctx, c, v, err := cmdContext(cmd)
	if err != nil {
		return err
	}
	if err := validateArgLenth(0, args, cmd); err != nil {
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
		Query:                    query,
		EntitlementAlias:         alias,
		GrantedStatus:            grantedStatus,
		AppDisplayName:           v.GetString(appDisplayNameFlag),
		IncludeDeleted:           v.GetBool(includeDeletedFlag),
		AppEntitlementExpandMask: shared.AppEntitlementExpandMask{Paths: []string{"*"}},
	})
	if err != nil {
		return err
	}
	resp := &ExpandedEntitlementsResponse{
		Entitlements: entitlements,
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
}

const DisplayNameHeader = "Display Name"
const AppHeader = "App"

func (r *ExpandedEntitlementsResponse) Header() []string {
	return []string{"", "Alias", DisplayNameHeader, AppHeader, "Resource Type", "Resource"}
}

func (r *ExpandedEntitlementsResponse) WideHeader() []string {
	return append(r.Header(), "Description", "App ID", "Entitlement ID")
}

func (r *ExpandedEntitlementsResponse) GetExpandedDisplayName(pathname string, e *client.EntitlementWithBindings) string {
	app := e.Expanded[pathname]
	if app.AdditionalProperties["displayName"] == nil {
		return ""
	}
	return app.AdditionalProperties["displayName"].(string)
}

func (r *ExpandedEntitlementsResponse) Rows() [][]string {
	rows := [][]string{}
	for _, e := range r.Entitlements {
		granted := output.Checkmark
		if len(e.Bindings) == 0 {
			granted = output.Unchecked
		}
		rows = append(rows, []string{
			granted,
			client.StringFromPtr(e.Entitlement.Alias),
			client.StringFromPtr(e.Entitlement.DisplayName),
			r.GetExpandedDisplayName("App", e),
			r.GetExpandedDisplayName("AppResourceType", e),
			r.GetExpandedDisplayName("AppResource", e),
		})
	}
	return rows
}

func (r *ExpandedEntitlementsResponse) OrderedSortColumns() []string {
	return []string{
		DisplayNameHeader,
		AppHeader,
	}
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
