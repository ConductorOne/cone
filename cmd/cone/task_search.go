package main

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/conductorone/cone/internal/c1api"
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

	var state []string
	switch strings.ToLower(v.GetString(stateFlag)) {
	case "open", "task_state_open":
		tstate := "TASK_STATE_OPEN"
		state = []string{tstate}
	case "closed", "task_state_closed":
		tstate := "TASK_STATE_CLOSED"
		state = []string{tstate}
	case "":
	}

	taskResp, err := c.SearchTasks(ctx, c1api.C1ApiTaskV1TaskSearchRequest{
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
		SubjectIds:         v.GetStringSlice(userSubjectIdsFlag),
	})
	if err != nil {
		return err
	}

	resp := C1ApiTaskV1TaskSearchResponse(*taskResp)
	outputManager := output.NewManager(ctx, v)
	err = outputManager.Output(ctx, &resp)
	if err != nil {
		return err
	}

	return nil
}

type C1ApiTaskV1TaskSearchResponse c1api.C1ApiTaskV1TaskSearchResponse

func (r *C1ApiTaskV1TaskSearchResponse) Header() []string {
	task := C1ApiTaskV1Task{}
	return task.Header()
}

func (r *C1ApiTaskV1TaskSearchResponse) Rows() [][]string {
	rows := [][]string{}
	for _, task := range r.List {
		t := C1ApiTaskV1Task{
			task:   task.Task,
			client: nil,
		}
		rows = append(rows, t.Rows()[0])
	}

	return rows
}
