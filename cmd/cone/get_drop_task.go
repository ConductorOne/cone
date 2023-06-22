package main

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xhit/go-str2duration/v2"

	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"

	"github.com/conductorone/cone/pkg/client"
	"github.com/conductorone/cone/pkg/output"
)

const grantDurationErrorMessage = "grant duration must be less than or equal to max provision time"
const durationInputTip = "We accept a sequence of decimal numbers, each with optional fraction and a unit suffix," +
	"such as \"12h\", \"1w2d\" or \"2h45m\". Valid units are (m)inutes, (h)ours, (d)ays, (w)eeks."

func getCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Create an access request for an entitlement by alias",
		RunE:  runGet,
	}
	addGrantDurationFlag(cmd)
	cmd = taskCmd(cmd)
	cmd.MarkFlagRequired(justificationFlag)
	return cmd
}

func dropCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "drop",
		Short: "Create a revoke access ticket for an entitlement by alias",
		RunE:  runDrop,
	}

	return taskCmd(cmd)
}

func taskCmd(cmd *cobra.Command) *cobra.Command {
	addWaitFlag(cmd)
	addAppIdFlag(cmd)
	addEntitlementIdFlag(cmd)
	addJustificationFlag(cmd)
	addQueryFlag(cmd)
	addEntitlementAliasFlag(cmd)
	addForceTaskCreateFlag(cmd)
	return cmd
}

func strToDur(duration string) (*time.Duration, error) {
	if duration == "" {
		return nil, nil
	}

	formattedDuration, err := str2duration.ParseDuration(duration)
	if err != nil {
		return nil, fmt.Errorf("invalid duration input %s", err.Error())
	}

	if formattedDuration <= time.Duration(0) {
		return nil, errors.New("duration must be greater than 0")
	}

	return &formattedDuration, nil
}

func handleDurationNonInteractive(maxProvisionTime *time.Duration, duration string) (*time.Duration, error) {
	formattedDuration, err := strToDur(duration)
	if err != nil {
		pterm.Info.Println(durationInputTip)
		return nil, err
	}

	if err := validateGrantTaskArguments(maxProvisionTime, formattedDuration); err != nil {
		pterm.Info.Println(durationInputTip)
		return nil, err
	}

	return formattedDuration, nil
}

func validateGrantTaskArguments(maxProvisionTime *time.Duration, duration *time.Duration) error {
	// If maxProvisionTime is set, ensure the duration is not nil (which means infinite)
	if maxProvisionTime != nil && duration == nil {
		return fmt.Errorf("%s: %s", grantDurationErrorMessage, str2duration.String(*maxProvisionTime))
	}

	// If maxProvisionTime is set, ensure the duration is not greater than maxProvisionTime
	if maxProvisionTime != nil && *duration > *maxProvisionTime {
		return fmt.Errorf("%s: %s", grantDurationErrorMessage, str2duration.String(*maxProvisionTime))
	}

	return nil
}

func getValidDuration(ctx context.Context, v *viper.Viper, maxProvisionTime *time.Duration, duration string) (*time.Duration, error) {
	// If both are empty that means they are both infinite
	if maxProvisionTime == nil && duration == "" {
		return nil, nil
	}

	if v.GetBool(nonInteractiveFlag) {
		return handleDurationNonInteractive(maxProvisionTime, duration)
	}
	input := pterm.DefaultInteractiveTextInput.WithMultiLine(false)
	firstRun := true
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		if !firstRun {
			var err error
			duration, err = input.Show()
			if err != nil {
				return nil, err
			}
		}
		firstRun = false

		formattedDuration, err := strToDur(duration)
		if err != nil {
			pterm.Error.Println(err.Error())
			pterm.Info.Println(durationInputTip)
			continue
		}

		err = validateGrantTaskArguments(maxProvisionTime, formattedDuration)
		if err != nil {
			pterm.Error.Println(err.Error())
			pterm.Info.Println(durationInputTip)
			continue
		}
		if *formattedDuration > 28*24*time.Hour || *formattedDuration < 5*time.Minute {
			result, _ := pterm.DefaultInteractiveConfirm.Show("The time you entered is outside of the range of 5 minutes - 21 days. Are you sure?")
			if !result {
				pterm.Info.Println(durationInputTip)
				continue
			}
		}

		return formattedDuration, nil
	}
}

func runGet(cmd *cobra.Command, args []string) error {
	_, _, v, err := cmdContext(cmd)
	if err != nil {
		return err
	}

	return runTask(cmd, args, func(c client.C1Client, ctx context.Context, appId string, entitlementId string, userId string, justification string) (*shared.Task, error) {
		duration := v.GetString(durationFlag)

		entitlement, err := c.GetEntitlement(ctx, appId, entitlementId)
		if err != nil {
			return nil, err
		}

		// entitlement.DurationGrant is assumed to be nil or a non-zero parsable string
		durationStr := client.StringFromPtr(entitlement.DurationGrant)
		var maxProvision *time.Duration
		maxProvisionTime, err := time.ParseDuration(durationStr)
		if err == nil {
			maxProvision = &maxProvisionTime
		}

		validDuration, err := getValidDuration(ctx, v, maxProvision, duration)
		if err != nil {
			return nil, err
		}

		apiDuration := ""
		if validDuration != nil {
			// API expects seconds formated like "1s"
			seconds := int(validDuration.Seconds())
			apiDuration = fmt.Sprintf("%ds", seconds)
		}

		accessRequest, err := c.CreateGrantTask(ctx, appId, entitlementId, userId, justification, apiDuration)
		if err != nil {
			errorBody := err.Error()
			if strings.Contains(errorBody, grantDurationErrorMessage) {
				startIndex := strings.Index(errorBody, grantDurationErrorMessage)
				endIndex := strings.LastIndex(errorBody, ")") + 1
				return nil, errors.New(errorBody[startIndex:endIndex])
			}
			return nil, err
		}
		return accessRequest.TaskView.Task, nil
	})
}

func runDrop(cmd *cobra.Command, args []string) error {
	return runTask(cmd, args, func(c client.C1Client, ctx context.Context, appId string, entitlementId string, userId string, justification string) (*shared.Task, error) {
		accessRequest, err := c.CreateRevokeTask(ctx, appId, entitlementId, userId, justification)
		if err != nil {
			return nil, err
		}
		return accessRequest.TaskView.Task, nil
	})
}

func runTask(
	cmd *cobra.Command,
	args []string,
	run func(c client.C1Client, ctx context.Context, appId string, entitlementId string, userId string, justification string) (*shared.Task, error),
) error {
	ctx, c, v, err := cmdContext(cmd)
	if err != nil {
		return err
	}

	justification := v.GetString(justificationFlag)

	entitlementId, appId, err := getEntitlementDetails(ctx, c, v, args)
	if err != nil {
		return err
	}

	resp, err := c.AuthIntrospect(ctx)
	if err != nil {
		return err
	}

	userID := client.StringFromPtr(resp.UserID)

	forceCreate := v.GetBool(forceFlag)
	if !forceCreate {
		grants, err := c.GetGrantsForIdentity(ctx, appId, entitlementId, userID)
		if err != nil {
			return err
		}
		grantCount := 0
		for _, grant := range grants {
			// We only want to check if user has a grant
			if client.StringFromPtr(grant.AppEntitlementID) != "" {
				grantCount++
			}
		}

		// If this is get, and they have grants, just exit
		if cmd.Name() == getCmd().Name() && grantCount > 0 {
			pterm.Println("You already have access to this entitlement. Use --force to override this check.")
			return nil
		}

		if cmd.Name() == dropCmd().Name() && grantCount == 0 {
			pterm.Println("You do not have existing grants to drop for this entitlement. Use --force to override this check.")
			return nil
		}
	}

	task, err := run(c, ctx, appId, entitlementId, client.StringFromPtr(resp.UserID), justification)
	if err != nil {
		return err
	}

	outputManager := output.NewManager(ctx, v)
	taskResp := Task{task: task, client: c}
	err = outputManager.Output(ctx, &taskResp)
	if err != nil {
		return err
	}

	if wait, _ := cmd.Flags().GetBool("wait"); wait {
		err = handleWaitBehavior(ctx, c, task, outputManager)
		if err != nil {
			return err
		}
	}

	return nil
}

func getEntitlementDetails(ctx context.Context, c client.C1Client, v *viper.Viper, args []string) (string, string, error) {
	entitlementId := v.GetString(entitlementIdFlag)
	appId := v.GetString(appIdFlag)
	query := v.GetString(queryFlag)

	alias := v.GetString(entitlementAliasFlag)
	if len(args) == 1 {
		alias = args[0]
	}

	if alias == "" && query == "" && (appId == "" || entitlementId == "") {
		return "", "", fmt.Errorf("must provide either an alias, query string, or an entitlement id and app id")
	}

	if (alias != "" || query != "") && (appId != "" || entitlementId != "") {
		return "", "", fmt.Errorf("cannot provide an alias or query and an entitlement id and app id")
	}

	// If we have an appId and appEntitlementId, just return those
	if appId != "" && entitlementId != "" {
		return entitlementId, appId, nil
	}

	// If we have an alias or query, search
	entitlements, err := c.SearchEntitlements(ctx, &client.SearchEntitlementsFilter{EntitlementAlias: alias, Query: query})
	if err != nil {
		return "", "", err
	}

	if len(entitlements) == 0 {
		return "", "", noEntitlementFoundError(alias, query)
	}

	if len(entitlements) == 1 {
		entitlementId = client.StringFromPtr(entitlements[0].Entitlement.ID)
		appId = client.StringFromPtr(entitlements[0].Entitlement.AppID)
	}
	if len(entitlements) > 1 {
		isNonInteractive := v.GetBool("non-interactive")
		if isNonInteractive {
			return "", "", multipleEntitlmentsFoundError(alias, query)
		}
		optionToEntitlementMap := make(map[string]*client.AppEntitlement)
		entitlementOptions := make([]string, len(entitlements))
		for _, e := range entitlements {
			entitlementOptionName := fmt.Sprintf("%s:%s:%s",
				client.StringFromPtr(e.Entitlement.DisplayName),
				client.StringFromPtr(e.Entitlement.AppID),
				client.StringFromPtr(e.Entitlement.ID),
			)
			entitlementOptions = append(entitlementOptions, entitlementOptionName)
			optionToEntitlementMap[entitlementOptionName] = &e.Entitlement
		}
		selectedOption, _ := pterm.DefaultInteractiveSelect.WithOptions(entitlementOptions).WithDefaultText("Please select an entitlement").Show()
		entitlementId = client.StringFromPtr(optionToEntitlementMap[selectedOption].ID)
		appId = client.StringFromPtr(optionToEntitlementMap[selectedOption].AppID)
	}
	return entitlementId, appId, nil
}

func handleWaitBehavior(ctx context.Context, c client.C1Client, task *shared.Task, outputManager output.Manager) error {
	spinner, _ := pterm.DefaultSpinner.Start("Waiting for ticket to close.")
	var taskItem *shared.Task
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(1 * time.Second):
		}
		task, err := c.GetTask(ctx, client.StringFromPtr(task.ID))
		if err != nil {
			return err
		}

		taskItem = task.TaskView.Task
		taskResp := Task{task: taskItem, client: c}
		err = outputManager.Output(ctx, &taskResp)
		if err != nil {
			return err
		}

		if *taskItem.State == shared.TaskStateTaskStateClosed {
			break
		}
	}
	if taskItem.Type.Grant != nil {
		taskOutcome := taskItem.Type.Grant.Outcome
		if *taskOutcome == shared.TaskTypeGrantOutcomeGrantOutcomeGranted {
			spinner.Success("Entitlement granted successfully.")
		} else {
			spinner.Fail(fmt.Sprintf("Failed to grant entitlement %s", string(*taskOutcome)))
			return fmt.Errorf("failed to grant entitlement %s", string(*taskOutcome))
		}
	}
	if taskItem.Type.Revoke != nil {
		taskOutcome := taskItem.Type.Revoke.Outcome
		if *taskOutcome == shared.TaskTypeRevokeOutcomeRevokeOutcomeRevoked {
			spinner.Success("Entitlement revoked succesfully.")
		} else {
			spinner.Fail(fmt.Sprintf("Failed to revoke entitlement %s", string(*taskOutcome)))
			return fmt.Errorf("failed to revoke entitlement %s", string(*taskOutcome))
		}
	}
	return nil
}

var processStateToString = map[shared.TaskProcessing]string{
	"TASK_PROCESSING_TYPE_UNSPECIFIED": "Unknown Processing",
	"TASK_PROCESSING_TYPE_PROCESSING":  "Processing",
	"TASK_PROCESSING_TYPE_WAITING":     "Waiting for Action",
	"TASK_PROCESSING_TYPE_DONE":        "Done",
}

var taskStateToString = map[shared.TaskState]string{
	"TASK_STATE_OPEN":   "Open",
	"TASK_STATE_CLOSED": "Closed",
}

func noEntitlementFoundError(alias string, query string) error {
	if alias != "" && query != "" {
		return fmt.Errorf("no entitlement found with alias %s and query %s", alias, query)
	}
	if alias != "" {
		return fmt.Errorf("no entitlement found with alias %s", alias)
	}
	if query != "" {
		return fmt.Errorf("no entitlement found with query %s", query)
	}
	return fmt.Errorf("no entitlement found")
}

func multipleEntitlmentsFoundError(alias string, query string) error {
	if alias != "" && query != "" {
		return fmt.Errorf("multiple entitlements found with alias %s and query %s, please specify an entitlement id and app id", alias, query)
	}
	if alias != "" {
		return fmt.Errorf("multiple entitlements found with alias %s, please specify an entitlement id and app id", alias)
	}
	if query != "" {
		return fmt.Errorf("multiple entitlements found with query %s, please specify an entitlement id and app id", query)
	}
	return fmt.Errorf("multiple entitlements found, please specify an entitlement id and app id")
}
