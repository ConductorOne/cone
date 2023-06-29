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
		Short: "",
		RunE:  searchTasksRun,
	}

	addAccessReviewIdsFlag(cmd)
	addAppEntitlementIdsFlag(cmd)
	addAppResourceIdsFlag(cmd)
	addAppResourceTypeIdsFlag(cmd)
	addAppUserSubjectIdsFlag(cmd)
	addUserSubjectIdsFlag(cmd)
	addAppApplicationIdsFlag(cmd)
	addAssigneesIds(cmd)
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

	var state []shared.TaskSearchRequestTaskStates
	switch strings.ToLower(v.GetString(stateFlag)) {
	case "open", "task_state_open":
		state = []shared.TaskSearchRequestTaskStates{shared.TaskSearchRequestTaskStatesTaskStateOpen}
	case "closed", "task_state_closed":
		state = []shared.TaskSearchRequestTaskStates{shared.TaskSearchRequestTaskStatesTaskStateClosed}
	case "":
	}

	var taskTypes []shared.TaskType
	var taskType *shared.TaskType
	switch strings.ToLower(v.GetString(taskTypeFlag)) {
	case "grant":
		taskType = &shared.TaskType{Grant: &shared.TaskTypeGrant{}}
	case "revoke":
		taskType = &shared.TaskType{Revoke: &shared.TaskTypeRevoke{}}
	case "certify":
		taskType = &shared.TaskType{Certify: &shared.TaskTypeCertify{}}
	}

	if taskType != nil {
		taskTypes = []shared.TaskType{*taskType}
	}

	taskResp, err := c.SearchTasks(ctx, shared.TaskSearchRequest{
		AccessReviewIds:    v.GetStringSlice(accessReviewIdsFlag),
		AppEntitlementIds:  v.GetStringSlice(appEntitlementIdsFlag),
		AppResourceIds:     v.GetStringSlice(appResourceIdsFlag),
		AppResourceTypeIds: v.GetStringSlice(appResourceTypeIdsFlag),
		AppUserSubjectIds:  v.GetStringSlice(appUserSubjectIdsFlag),
		ApplicationIds:     v.GetStringSlice(appIdsFlag),
		AssigneesInIds:     v.GetStringSlice(assigneeIdsFlag),
		IncludeDeleted:     includeDeleted,
		Query:              query,
		TaskStates:         state,
		TaskTypes:          taskTypes,
		SubjectIds:         v.GetStringSlice(userSubjectIdsFlag),
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
