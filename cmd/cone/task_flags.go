package main

import "github.com/spf13/cobra"

const (
	accessReviewIdsFlag    = "access-review-ids"
	appEntitlementIdsFlag  = "app-entitlement-ids"
	appResourceIdsFlag     = "app-resource-ids"
	appResourceTypeIdsFlag = "app-resource-type-ids"
	appUserSubjectIdsFlag  = "app-user-subject-ids"
	userSubjectIdsFlag     = "user-subject-ids"
	appIdsFlag             = "app-ids"
	assigneeIdsFlag        = "assignee-ids"
	stateFlag              = "state"
	includeDeletedFlag     = "include-deleted"
	commentFlag            = "comment"
)

func addCommentFlag(cmd *cobra.Command) {
	cmd.Flags().String(commentFlag, "", "Comment to add to the task when performing an action")
}
func addAccessReviewIdsFlag(cmd *cobra.Command) {
	cmd.Flags().StringSlice(accessReviewIdsFlag, nil, "Filter tasks by access review ids (access review campaign this task belongs to)")
}
func addAppEntitlementIdsFlag(cmd *cobra.Command) {
	cmd.Flags().StringSlice(appEntitlementIdsFlag, nil, "Filter tasks by app entitlement ids (target app entitlement of the ticket)")
}
func addAppResourceIdsFlag(cmd *cobra.Command) {
	cmd.Flags().StringSlice(appResourceIdsFlag, nil, "Filter tasks by app resource ids (target resource of the ticket)")
}
func addAppResourceTypeIdsFlag(cmd *cobra.Command) {
	cmd.Flags().StringSlice(appResourceTypeIdsFlag, nil, "Filter tasks by app resource type ids (target resource type of the task)")
}
func addAppUserSubjectIdsFlag(cmd *cobra.Command) {
	cmd.Flags().StringSlice(appUserSubjectIdsFlag, nil, "Filter tasks by app user subject ids (target of the task)")
}
func addUserSubjectIdsFlag(cmd *cobra.Command) {
	cmd.Flags().StringSlice(userSubjectIdsFlag, nil, "Filter tasks by user subject ids (c1 user target of the task)")
}
func addAppApplicationIdsFlag(cmd *cobra.Command) {
	cmd.Flags().StringSlice(appIdsFlag, nil, "Filter tasks by app application ids (target application of the task)")
}
func addAssigneesIds(cmd *cobra.Command) {
	cmd.Flags().StringSlice(assigneeIdsFlag, nil, "Filter tasks by who is currently assigned to them")
}

func addQueryTaskFlag(cmd *cobra.Command) {
	cmd.Flags().String(queryFlag, "", "Query string to filter tasks")
}
func addTaskStatesFlag(cmd *cobra.Command) {
	cmd.Flags().String(stateFlag, "", "Filter tasks by their state (open, closed)")
}
func addIncludeDeletedFlag(cmd *cobra.Command) {
	cmd.Flags().Bool(includeDeletedFlag, false, "Include deleted tasks in the results")
}
