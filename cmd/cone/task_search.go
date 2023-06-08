package main

import (
	"errors"
	"strings"
	"time"

	"github.com/conductorone/cone/internal/c1api"
	"github.com/conductorone/cone/pkg/output"
	"github.com/spf13/cobra"
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
	addCreatedBeforeFlag(cmd)
	addCreatedAfterFlag(cmd)
	addQueryTaskFlag(cmd)
	addTaskStatesFlag(cmd)
	addIncludeDeletedFlag(cmd)

	return cmd
}

var validTimeFormats = []string{
	time.RFC3339,
	time.Layout,
	time.RFC1123,
	time.RFC1123Z,
	time.Stamp,
	time.DateTime,
	time.DateOnly,
	time.TimeOnly,
	time.Kitchen,
}

var noValidTimeFound = errors.New("could not parse time, valid formats are go's RFC3339, Layout, RFC1123, RFC1123Z, Stamp, DateTime, DateOnly, TimeOnly, and Kitchen")

func searchTasksRun(cmd *cobra.Command, args []string) error {
	ctx, c, v, err := cmdContext(cmd)
	if err != nil {
		return err
	}

	var createdAfter *time.Time
	if v.GetString(createdAfterFlag) != "" {
		for _, validTimeFormat := range validTimeFormats {
			createdAfterParsed, err := time.Parse(validTimeFormat, v.GetString(createdAfterFlag))
			if err != nil {
				continue
			}
			createdAfter = &createdAfterParsed
		}
		if createdAfter == nil {
			return noValidTimeFound
		}
	}

	var createdBefore *time.Time
	if v.GetString(createdBeforeFlag) != "" {
		for _, validTimeFormat := range validTimeFormats {
			createdBeforeParsed, err := time.Parse(validTimeFormat, v.GetString(createdBeforeFlag))
			if err != nil {
				continue
			}
			createdBefore = &createdBeforeParsed
		}
		if createdAfter == nil {
			return noValidTimeFound
		}
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
		CreatedAfter:       createdAfter,
		CreatedBefore:      createdBefore,
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
