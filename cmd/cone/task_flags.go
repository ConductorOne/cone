package main

import "github.com/spf13/cobra"

const (
	accessReviewIDsFlag    = "access-review-id"
	appEntitlementIDsFlag  = "app-entitlement-id"
	appResourceIDsFlag     = "app-resource-id"
	appResourceTypeIDsFlag = "app-resource-type-id"
	appUserSubjectIDsFlag  = "app-user-subject-id"
	userSubjectIDsFlag     = "user-subject-id"
	appIDsFlag             = "app-id"
	assigneeIDsFlag        = "assignee-id"
	stateFlag              = "state"
	taskTypeFlag           = "task-type"
	includeDeletedFlag     = "include-deleted"
	commentFlag            = "comment"
)

func addCommentFlag(cmd *cobra.Command) {
	cmd.Flags().String(commentFlag, "", "Comment to add to the task when performing an action")
}
func addAccessReviewIDsFlag(cmd *cobra.Command) {
	cmd.Flags().StringSlice(accessReviewIDsFlag, nil, "Filter tasks by access review id(s) (access review campaign this task belongs to)")
}
func addAppEntitlementIDsFlag(cmd *cobra.Command) {
	cmd.Flags().StringSlice(appEntitlementIDsFlag, nil, "Filter tasks by app entitlement id(s) (target app entitlement of the ticket)")
}
func addAppResourceIDsFlag(cmd *cobra.Command) {
	cmd.Flags().StringSlice(appResourceIDsFlag, nil, "Filter tasks by app resource id(s) (target resource of the ticket)")
}
func addAppResourceTypeIDsFlag(cmd *cobra.Command) {
	cmd.Flags().StringSlice(appResourceTypeIDsFlag, nil, "Filter tasks by app resource type id(s) (target resource type of the task)")
}
func addAppUserSubjectIDsFlag(cmd *cobra.Command) {
	cmd.Flags().StringSlice(appUserSubjectIDsFlag, nil, "Filter tasks by app user subject id(s) (target of the task)")
}
func addUserSubjectIDsFlag(cmd *cobra.Command) {
	cmd.Flags().StringSlice(userSubjectIDsFlag, nil, "Filter tasks by user subject id(s) (c1 user target of the task)")
}
func addAppApplicationIDsFlag(cmd *cobra.Command) {
	cmd.Flags().StringSlice(appIDsFlag, nil, "Filter tasks by app application id(s) (target application of the task)")
}
func addAssigneesIDs(cmd *cobra.Command) {
	cmd.Flags().StringSlice(assigneeIDsFlag, nil, "Filter tasks by who is currently assigned to them")
}

func addQueryTaskFlag(cmd *cobra.Command) {
	cmd.Flags().String(queryFlag, "", "Query string to filter tasks")
}
func addTaskStatesFlag(cmd *cobra.Command) {
	cmd.Flags().String(stateFlag, "", "Filter tasks by their state (open, closed)")
}
func addTaskTypesFlag(cmd *cobra.Command) {
	cmd.Flags().String(taskTypeFlag, "", "Filter tasks by their task type (grant, revoke, certify)")
}

func addIncludeDeletedFlag(cmd *cobra.Command) {
	cmd.Flags().Bool(includeDeletedFlag, false, "Include deleted tasks in the results")
}
