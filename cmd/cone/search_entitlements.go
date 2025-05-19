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
		Short: "Search for entitlements in ConductorOne",
		Long: `Search for entitlements in ConductorOne using various filters.
This command allows you to:
- Search by entitlement name or alias
- Filter by app name
- Show only granted or not granted entitlements
- Include deleted entitlements

The search results will show:
- Whether you have access to each entitlement
- The entitlement's alias and display name
- The app it belongs to
- The resource type and resource name`,
		RunE: searchEntitlementsRun,
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

	query := v.GetString(queryFlag)
	alias := v.GetString(entitlementAliasFlag)
	if len(args) == 1 {
		alias = args[0]
	}

	grantedStatus := shared.GrantedStatusAll
	if v.GetBool(grantedFlag) {
		grantedStatus = shared.GrantedStatusGranted
	} else if v.GetBool(notGrantedFlag) {
		grantedStatus = shared.GrantedStatusNotGranted
	}

	// Phase 1: Try exact alias match first if an alias is provided
	var entitlements []*client.EntitlementWithBindings
	if alias != "" {
		exactMatchEntitlements, err := c.SearchEntitlements(ctx, &client.SearchEntitlementsFilter{
			EntitlementAlias:         alias,
			GrantedStatus:            grantedStatus,
			AppDisplayName:           v.GetString(appDisplayNameFlag),
			IncludeDeleted:           v.GetBool(includeDeletedFlag),
			AppEntitlementExpandMask: shared.AppEntitlementExpandMask{Paths: []string{"app_id", "app_resource_type_id", "app_resource_id"}},
		})
		if err != nil {
			return err
		}
		entitlements = exactMatchEntitlements
	}

	// Phase 2: If no exact matches found and we have a query, try query search
	if len(entitlements) == 0 && query != "" {
		queryEntitlements, err := c.SearchEntitlements(ctx, &client.SearchEntitlementsFilter{
			Query:                    query,
			GrantedStatus:            grantedStatus,
			AppDisplayName:           v.GetString(appDisplayNameFlag),
			IncludeDeleted:           v.GetBool(includeDeletedFlag),
			AppEntitlementExpandMask: shared.AppEntitlementExpandMask{Paths: []string{"app_id", "app_resource_type_id", "app_resource_id"}},
		})
		if err != nil {
			return err
		}
		entitlements = queryEntitlements
	}

	// If still no results and we have both alias and query, try combined search
	if len(entitlements) == 0 && alias != "" && query != "" {
		combinedEntitlements, err := c.SearchEntitlements(ctx, &client.SearchEntitlementsFilter{
			Query:                    query,
			EntitlementAlias:         alias,
			GrantedStatus:            grantedStatus,
			AppDisplayName:           v.GetString(appDisplayNameFlag),
			IncludeDeleted:           v.GetBool(includeDeletedFlag),
			AppEntitlementExpandMask: shared.AppEntitlementExpandMask{Paths: []string{"app_id", "app_resource_type_id", "app_resource_id"}},
		})
		if err != nil {
			return err
		}
		entitlements = combinedEntitlements
	}

	// If no alias or query provided, show all entitlements
	if len(entitlements) == 0 && alias == "" && query == "" {
		allEntitlements, err := c.SearchEntitlements(ctx, &client.SearchEntitlementsFilter{
			GrantedStatus:            grantedStatus,
			AppDisplayName:           v.GetString(appDisplayNameFlag),
			IncludeDeleted:           v.GetBool(includeDeletedFlag),
			AppEntitlementExpandMask: shared.AppEntitlementExpandMask{Paths: []string{"app_id", "app_resource_type_id", "app_resource_id"}},
		})
		if err != nil {
			return err
		}
		entitlements = allEntitlements
	}

	outputManager := output.NewManager(ctx, v)
	resp := &ExpandedEntitlementsResponse{
		Entitlements: entitlements,
	}
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

func (r *ExpandedEntitlementsResponse) Rows() [][]string {
	rows := [][]string{}
	for _, e := range r.Entitlements {
		granted := output.Checkmark
		if len(e.Bindings) == 0 {
			granted = output.Unchecked
		}
		app := client.GetExpanded[shared.App](e, client.ExpandedApp)
		appResourceType := client.GetExpanded[shared.AppResourceType](e, client.ExpandedAppResourceType)
		appResource := client.GetExpanded[shared.AppResource](e, client.ExpandedAppResource)
		rows = append(rows, []string{
			granted,
			client.StringFromPtr(e.Entitlement.Alias),
			client.StringFromPtr(e.Entitlement.DisplayName),
			client.StringFromPtr(app.GetDisplayName()),
			client.StringFromPtr(appResourceType.GetDisplayName()),
			client.StringFromPtr(appResource.GetDisplayName()),
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
