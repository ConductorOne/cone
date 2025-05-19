package main

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"

	"github.com/conductorone/cone/pkg/output"
)

func searchTasksCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "search",
		Short: "Search for tasks in ConductorOne",
		Long: `Search for tasks in ConductorOne using various filters.
This command allows you to:
- Search by task type (grant, revoke, certify)
- Filter by task state (open, closed)
- Filter by app, entitlement, resource, or user
- Search for specific access reviews
- Include deleted tasks

The search results will show:
- Task type and state
- Related apps, entitlements, and resources
- Assignees and subjects
- Creation and update timestamps`,
		RunE: searchTasksRun,
	}

	addAccessReviewIDsFlag(cmd)
	addAppEntitlementIDsFlag(cmd)
	addAppResourceIDsFlag(cmd)
	addAppResourceTypeIDsFlag(cmd)
	addAppUserSubjectIDsFlag(cmd)
	addUserSubjectIDsFlag(cmd)
	addAppApplicationIDsFlag(cmd)
	addAssigneesIDs(cmd)
	addQueryTaskFlag(cmd)
	addTaskStatesFlag(cmd)
	addTaskTypesFlag(cmd)
	addIncludeDeletedFlag(cmd)

	return cmd
}

func searchTasksRun(cmd *cobra.Command, args []string) error {
	ctx, c, v, err := cmdContext(cmd)
	if err != nil {
		return err
	}

	if err := validateArgLenth(0, args, cmd); err != nil {
		return err
	}

	var includeDeleted *bool
	if v.Get(includeDeletedFlag) != nil {
		includeDeletedVal := v.GetBool(includeDeletedFlag)
		includeDeleted = &includeDeletedVal
	}

	var query *string
	if v.GetString(queryFlag) != "" {
		queryVal := v.GetString(queryFlag)
		query = &queryVal
	}

	var state []shared.TaskStates
	switch strings.ToLower(v.GetString(stateFlag)) {
	case "open", "task_state_open":
		state = []shared.TaskStates{shared.TaskStatesTaskStateOpen}
	case "closed", "task_state_closed":
		state = []shared.TaskStates{shared.TaskStatesTaskStateClosed}
	case "":
	}

	var taskTypes []shared.TaskTypeInput
	var taskType *shared.TaskTypeInput
	switch strings.ToLower(v.GetString(taskTypeFlag)) {
	case "grant":
		taskType = &shared.TaskTypeInput{TaskTypeGrant: &shared.TaskTypeGrantInput{}}
	case "revoke":
		taskType = &shared.TaskTypeInput{TaskTypeRevoke: &shared.TaskTypeRevokeInput{}}
	case "certify":
		taskType = &shared.TaskTypeInput{TaskTypeCertify: &shared.TaskTypeCertifyInput{}}
	}

	if taskType != nil {
		taskTypes = []shared.TaskTypeInput{*taskType}
	}

	taskResp, err := c.SearchTasks(ctx, shared.TaskSearchRequest{
		AccessReviewIds:    v.GetStringSlice(accessReviewIDsFlag),
		AppEntitlementIds:  v.GetStringSlice(appEntitlementIDsFlag),
		AppResourceIds:     v.GetStringSlice(appResourceIDsFlag),
		AppResourceTypeIds: v.GetStringSlice(appResourceTypeIDsFlag),
		AppUserSubjectIds:  v.GetStringSlice(appUserSubjectIDsFlag),
		ApplicationIds:     v.GetStringSlice(appIDsFlag),
		AssigneesInIds:     v.GetStringSlice(assigneeIDsFlag),
		IncludeDeleted:     includeDeleted,
		Query:              query,
		TaskStates:         state,
		TaskTypes:          taskTypes,
		SubjectIds:         v.GetStringSlice(userSubjectIDsFlag),
	})
	if err != nil {
		return err
	}

	resp := TaskSearchResponse(*taskResp)
	outputManager := output.NewManager(ctx, v)
	err = outputManager.Output(ctx, &resp)
	if err != nil {
		return err
	}

	return nil
}

type TaskSearchResponse shared.TaskSearchResponse

func (r *TaskSearchResponse) Header() []string {
	task := Task{}
	return task.Header()
}

func (r *TaskSearchResponse) Rows() [][]string {
	rows := [][]string{}
	for _, task := range r.List {
		t := Task{
			task:   task.Task,
			client: nil,
		}
		rows = append(rows, t.Rows()[0])
	}

	return rows
}

func (r *TaskSearchResponse) WideHeader() []string {
	task := Task{}
	return task.WideHeader()
}

func (r *TaskSearchResponse) WideRows() [][]string {
	rows := [][]string{}
	for _, task := range r.List {
		t := Task{
			task:   task.Task,
			client: nil,
		}
		rows = append(rows, t.WideRows()[0])
	}

	return rows
}
