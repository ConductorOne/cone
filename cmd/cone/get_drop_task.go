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

const durationErrorMessage = "grant duration must be less than or equal to max provision time"
const durationInputTip = "We accept a sequence of decimal numbers, each with optional fraction and a unit suffix," +
	"such as \"12h\", \"1w2d\" or \"2h45m\". Valid units are (m)inutes, (h)ours, (d)ays, (w)eeks."
const justificationWarningMessage = "please provide a justification when requesting access to an entitlement"
const justificationInputTip = "You can add a justification using -j or --justification"
const appUserMultipleUsersWarningMessage = "this app has multiple users. Please select any one. "

func getCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get <alias> [flags]\n  cone get --query <query> [flags]\n  cone get --app-id <app-id> --entitlement-id <entitlement-id> [flags]",
		Short: "Create an access request for an entitlement by alias",
		RunE:  runGet,
	}
	addGrantDurationFlag(cmd)
	addEmergencyAccessFlag(cmd)
	return taskCmd(cmd)
}

func dropCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "drop <alias> [flags]\n  cone drop --query <query> [flags]\n  cone drop --app-id <app-id> --entitlement-id <entitlement-id> [flags]",
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
	addEntitlementDetailsFlag(cmd)
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
		return fmt.Errorf("%s: %s", durationErrorMessage, str2duration.String(*maxProvisionTime))
	}

	// If maxProvisionTime is set, ensure the duration is not greater than maxProvisionTime
	if maxProvisionTime != nil && *duration > *maxProvisionTime {
		return fmt.Errorf("%s: %s", durationErrorMessage, str2duration.String(*maxProvisionTime))
	}

	return nil
}

type JustificationValidator struct{}

func (j JustificationValidator) IsValid(txt string) (string, bool) {
	return txt, strings.TrimSpace(txt) != ""
}

func (j JustificationValidator) Prompt(isFirstRun bool) {
	if isFirstRun {
		pterm.Info.Println(justificationInputTip)
	}
	output.InputNeeded.Println(justificationWarningMessage)
}

func getValidJustification(ctx context.Context, v *viper.Viper, justification string) (string, error) {
	if strings.TrimSpace(justification) != "" {
		return justification, nil
	}

	if v.GetBool(nonInteractiveFlag) {
		pterm.Info.Println(justificationInputTip)
		return "", errors.New(justificationWarningMessage)
	}
	justificationInput, err := output.GetValidInput[string](ctx, justification, JustificationValidator{})
	if err != nil {
		return "", err
	}
	return justificationInput, nil
}

type DurationValidator struct {
	maxProvisionTime *time.Duration
}

func (d DurationValidator) IsValid(txt string) (time.Duration, bool) {
	var t time.Duration
	formattedDuration, err := strToDur(txt)
	if err != nil {
		return t, false
	}

	err = validateGrantTaskArguments(d.maxProvisionTime, formattedDuration)
	if err != nil {
		return t, false
	}

	const daysInAMonth = 28
	const upperBound = 24 * time.Hour * daysInAMonth
	const lowerBound = 5 * time.Minute
	if *formattedDuration < lowerBound || *formattedDuration > upperBound {
		warningMessage := fmt.Sprintf("The time you entered is outside of the range of %d minutes - %d days. Are you sure?", int(lowerBound.Minutes()), daysInAMonth)
		result, _ := pterm.DefaultInteractiveConfirm.Show(warningMessage)
		if !result {
			return t, false
		}
	}
	return *formattedDuration, true
}

func (d DurationValidator) Prompt(isFirstRun bool) {
	if isFirstRun {
		pterm.Info.Println(durationInputTip)
	}
	pterm.Error.Println(durationErrorMessage)
}

func getValidDuration(ctx context.Context, v *viper.Viper, maxProvisionTime *time.Duration, duration string) (*time.Duration, error) {
	// If both are empty that means they are both infinite
	if maxProvisionTime == nil && duration == "" {
		return nil, nil
	}

	if v.GetBool(nonInteractiveFlag) {
		return handleDurationNonInteractive(maxProvisionTime, duration)
	}

	durationInput, err := output.GetValidInput[time.Duration](ctx, duration, DurationValidator{maxProvisionTime})
	if err != nil {
		return nil, err
	}
	return &durationInput, nil
}

func runGet(cmd *cobra.Command, args []string) error {
	_, _, v, err := cmdContext(cmd)
	if err != nil {
		return err
	}

	return runTask(cmd, args, func(c client.C1Client, ctx context.Context, appId string, entitlementId string, userId string, justification string) (*shared.Task, error) {
		duration := v.GetString(durationFlag)
		emergencyAccess := v.GetBool(emergencyAccessFlag)

		entitlement, err := c.GetEntitlement(ctx, appId, entitlementId)
		if err != nil {
			return nil, err
		}

		justification, err = getValidJustification(ctx, v, justification)
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

		appUserId, err := getAppUserId(ctx, c, v, appId, userId)
		if err != nil {
			return nil, err
		}

		apiDuration := ""
		if validDuration != nil {
			// API expects seconds formated like "1s"
			seconds := int(validDuration.Seconds())
			apiDuration = fmt.Sprintf("%ds", seconds)
		}

		accessRequest, err := c.CreateGrantTask(ctx, appId, entitlementId, userId, appUserId, justification, apiDuration, emergencyAccess)
		if err != nil {
			errorBody := err.Error()
			if strings.Contains(errorBody, durationErrorMessage) {
				startIndex := strings.Index(errorBody, durationErrorMessage)
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

// createAWSSSOProfileIfNeeded checks if the entitlement is an AWS permission set and creates the profile if needed.
func createAWSSSOProfileIfNeeded(ctx context.Context, c client.C1Client, task *shared.Task, outputManager output.Manager) error {
	if task.TaskType.TaskTypeGrant == nil {
		return nil
	}

	grantTask := task.TaskType.TaskTypeGrant
	if grantTask.AppID == nil || grantTask.AppEntitlementID == nil {
		// Skip if required fields are missing - not an error for this function
		return nil
	}

	appID := *grantTask.AppID
	entitlementID := *grantTask.AppEntitlementID

	// Get the entitlement details
	entitlement, err := c.GetEntitlement(ctx, appID, entitlementID)
	if err != nil {
		return fmt.Errorf("failed to get entitlement details: %w", err)
	}

	// Check for nil pointers before dereferencing
	if entitlement.AppResourceTypeID == nil {
		// Not an error condition, just skip AWS SSO profile creation
		return nil
	}

	// Get the resource type
	resourceType, err := c.GetResourceType(ctx, appID, *entitlement.AppResourceTypeID)
	if err != nil {
		return fmt.Errorf("failed to get resource type: %w", err)
	}

	// Check if this is an AWS permission set
	if client.IsAWSPermissionSet(entitlement, resourceType) {
		if entitlement.AppResourceID == nil {
			return fmt.Errorf("entitlement AppResourceID is nil, cannot create AWS SSO profile")
		}

		// Get the resource details
		resource, err := c.GetResource(ctx, appID, *entitlement.AppResourceTypeID, *entitlement.AppResourceID)
		if err != nil {
			return fmt.Errorf("failed to get resource details: %w", err)
		}

		if err := client.CreateAWSSSOProfile(entitlement, resource); err != nil {
			return fmt.Errorf("failed to create AWS SSO profile: %w", err)
		}
	}
	return nil
}

// handleWaitBehavior manages the waiting state for task completion.
func handleWaitBehavior(ctx context.Context, c client.C1Client, task *shared.Task, outputManager output.Manager) error {
	// Validate input parameters
	if task == nil || task.ID == nil {
		return fmt.Errorf("task or task ID is nil")
	}

	spinner, _ := pterm.DefaultSpinner.Start("Waiting for task to complete...")
	defer func() { _ = spinner.Stop() }()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(2 * time.Second):
			// Continue with polling
		}

		updatedTask, err := c.GetTask(ctx, *task.ID)
		if err != nil {
			return err
		}

		// Check for nil pointers before dereferencing
		if updatedTask.TaskView == nil || updatedTask.TaskView.Task == nil {
			return fmt.Errorf("received incomplete task response")
		}

		if updatedTask.TaskView.Task.State == nil {
			return fmt.Errorf("task state is nil")
		}

		if *updatedTask.TaskView.Task.State == shared.TaskStateTaskStateClosed {
			if updatedTask.TaskView.Task.TaskType.TaskTypeGrant != nil {
				taskOutcome := updatedTask.TaskView.Task.TaskType.TaskTypeGrant.Outcome
				if taskOutcome == nil {
					return fmt.Errorf("task closed but no outcome provided")
				}

				switch *taskOutcome {
				case shared.TaskTypeGrantOutcomeGrantOutcomeGranted:
					spinner.Success("Entitlement granted successfully.")
				case shared.TaskTypeGrantOutcomeGrantOutcomeDenied:
					spinner.Fail("Entitlement request was denied.")
					return fmt.Errorf("entitlement request was denied")
				default:
					spinner.Fail(fmt.Sprintf("Task completed with unexpected outcome: %s", *taskOutcome))
					return fmt.Errorf("task completed with unexpected outcome: %s", *taskOutcome)
				}
			} else if updatedTask.TaskView.Task.TaskType.TaskTypeRevoke != nil {
				taskOutcome := updatedTask.TaskView.Task.TaskType.TaskTypeRevoke.Outcome
				if taskOutcome == nil {
					return fmt.Errorf("task closed but no outcome provided")
				}

				switch *taskOutcome {
				case shared.TaskTypeRevokeOutcomeRevokeOutcomeRevoked:
					spinner.Success("Entitlement revoked successfully.")
				case shared.TaskTypeRevokeOutcomeRevokeOutcomeDenied:
					spinner.Fail("Entitlement revoke request was denied.")
					return fmt.Errorf("entitlement revoke request was denied")
				default:
					spinner.Fail(fmt.Sprintf("Task completed with unexpected outcome: %s", *taskOutcome))
					return fmt.Errorf("task completed with unexpected outcome: %s", *taskOutcome)
				}
			}
			break
		}
	}

	return nil
}

// runTask executes the task and handles the response.
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

	entitlementId, appId, err := getEntitlementDetails(ctx, c, v, args, cmd)
	if err != nil {
		return err
	}

	resp, err := c.AuthIntrospect(ctx)
	if err != nil {
		return err
	}

	task, err := run(c, ctx, appId, entitlementId, client.StringFromPtr(resp.UserID), justification)
	if err != nil {
		return err
	}

	outputManager := output.NewManager(ctx, v)

	// Show detailed app and entitlement information if requested
	if v.GetBool(extraDetailsFlag) {
		err = showDetailedEntitlementInfo(ctx, c, appId, entitlementId, v)
		if err != nil {
			pterm.Warning.Printf("Failed to show detailed entitlement information: %v\n", err)
		}
	}

	taskResp := Task{task: task, client: c}
	err = outputManager.Output(ctx, &taskResp, output.WithTransposeTable())
	if err != nil {
		return err
	}

	// Create AWS SSO profile immediately after task creation
	if err := createAWSSSOProfileIfNeeded(ctx, c, task, outputManager); err != nil {
		pterm.Warning.Printf("Failed to create AWS SSO profile: %v\n", err)
	}

	if wait, _ := cmd.Flags().GetBool("wait"); wait {
		err = handleWaitBehavior(ctx, c, task, outputManager)
		if err != nil {
			return err
		}
	}

	return nil
}

func getAppUserProfileAttribute(appUserProfile map[string]any, profileAttribute string) string {
	if len(appUserProfile) == 0 || profileAttribute == "" {
		return ""
	}

	attrValue, ok := appUserProfile[profileAttribute]
	if !ok {
		return ""
	}

	return fmt.Sprintf("%v", attrValue)
}

func getAppUserId(ctx context.Context, c client.C1Client, v *viper.Viper, appId, userId string) (string, error) {
	appUsers, err := c.ListAppUsersForUser(ctx, appId, userId)
	if err != nil {
		return "", err
	}

	switch len(appUsers) {
	case 0:
		return "", nil
	case 1:
		return client.StringFromPtr(appUsers[0].ID), nil
	default:
		if v.GetBool(nonInteractiveFlag) {
			return "", errors.New(appUserMultipleUsersWarningMessage)
		}

		output.InputNeeded.Println(appUserMultipleUsersWarningMessage)

		optionToAppUsersMap := make(map[string]*shared.AppUser, len(appUsers))
		appUsersOptions := make([]string, 0, len(appUsers))

		for _, au := range appUsers {
			appUser := au
			appUserOptionName := fmt.Sprintf("%s:%s:%s",
				client.StringFromPtr(au.DisplayName),
				client.StringFromPtr(au.AppID),
				client.StringFromPtr(au.ID),
			)
			// If exists, append the aws user type to help differentiate between appUsers with the same name.
			awsUserType := getAppUserProfileAttribute(au.GetProfile(), "aws_user_type")
			if awsUserType != "" {
				appUserOptionName = fmt.Sprintf("%s (%s)", appUserOptionName, awsUserType)
			}
			appUsersOptions = append(appUsersOptions, appUserOptionName)
			optionToAppUsersMap[appUserOptionName] = &appUser
		}

		selectedOption, err := pterm.DefaultInteractiveSelect.
			WithMaxHeight(len(appUsersOptions)).
			WithOptions(appUsersOptions).
			WithDefaultText("Please select a user").
			Show()
		if err != nil {
			return "", err
		}

		return client.StringFromPtr(optionToAppUsersMap[selectedOption].ID), nil
	}
}

func getEntitlementDetails(ctx context.Context, c client.C1Client, v *viper.Viper, args []string, cmd *cobra.Command) (string, string, error) {
	entitlementId := v.GetString(entitlementIdFlag)
	appId := v.GetString(appIdFlag)
	query := v.GetString(queryFlag)

	alias := v.GetString(entitlementAliasFlag)
	if len(args) == 1 {
		alias = args[0]
	}

	if alias == "" && query == "" && (appId == "" || entitlementId == "") {
		return "", "", fmt.Errorf("must provide either an alias, query string, or an entitlement id and app id\n%s", cmd.UsageString())
	}

	if (alias != "" || query != "") && (appId != "" || entitlementId != "") {
		return "", "", fmt.Errorf("cannot provide an alias or query and an entitlement id and app id\n%s", cmd.UsageString())
	}

	// If we have an appId and appEntitlementId, just return those
	if appId != "" && entitlementId != "" {
		return entitlementId, appId, nil
	}

	// If we have an alias or query, search
	entitlements, err := c.SearchEntitlements(ctx, &client.SearchEntitlementsFilter{
		EntitlementAlias: alias,
		Query:            query,
		GrantedStatus:    shared.GrantedStatusAll,
	})
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
			optionToEntitlementMap[entitlementOptionName] = &(e.Entitlement)
		}
		selectedOption, _ := pterm.DefaultInteractiveSelect.WithOptions(entitlementOptions).WithDefaultText("Please select an entitlement").Show()
		entitlementId = client.StringFromPtr(optionToEntitlementMap[selectedOption].ID)
		appId = client.StringFromPtr(optionToEntitlementMap[selectedOption].AppID)
	}
	return entitlementId, appId, nil
}

var processStateToString = map[shared.Processing]string{
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

// showDetailedEntitlementInfo displays detailed information about the app and entitlement.
func showDetailedEntitlementInfo(ctx context.Context, c client.C1Client, appID string, entitlementID string, v *viper.Viper) error {
	// Skip detailed output for JSON formats to avoid mixing plain text with structured output
	outputFormat := v.GetString("output")
	if outputFormat == "json" || outputFormat == "json-pretty" {
		return nil
	}

	pterm.DefaultSection.Println("Entitlement Details")

	// Get app details
	app, err := c.GetApp(ctx, appID)
	if err != nil {
		return fmt.Errorf("failed to get app details: %w", err)
	}

	// Get entitlement details
	entitlement, err := c.GetEntitlement(ctx, appID, entitlementID)
	if err != nil {
		return fmt.Errorf("failed to get entitlement details: %w", err)
	}

	// Display app information
	pterm.DefaultBasicText.Printf("App: %s\n", client.StringFromPtr(app.DisplayName))
	if app.Description != nil && *app.Description != "" {
		pterm.DefaultBasicText.Printf("App Description: %s\n", client.StringFromPtr(app.Description))
	}

	// Display entitlement information
	pterm.DefaultBasicText.Printf("Entitlement: %s\n", client.StringFromPtr(entitlement.DisplayName))
	if entitlement.Description != nil && *entitlement.Description != "" {
		pterm.DefaultBasicText.Printf("Entitlement Description: %s\n", client.StringFromPtr(entitlement.Description))
	}

	// Show resource type and resource information if available
	if entitlement.AppResourceTypeID != nil {
		resourceType, err := c.GetResourceType(ctx, appID, *entitlement.AppResourceTypeID)
		if err == nil {
			pterm.DefaultBasicText.Printf("Resource Type: %s\n", client.StringFromPtr(resourceType.DisplayName))

			if entitlement.AppResourceID != nil {
				resource, err := c.GetResource(ctx, appID, *entitlement.AppResourceTypeID, *entitlement.AppResourceID)
				if err == nil {
					pterm.DefaultBasicText.Printf("Resource: %s\n", client.StringFromPtr(resource.DisplayName))
					if resource.Description != nil && *resource.Description != "" {
						pterm.DefaultBasicText.Printf("Resource Description: %s\n", client.StringFromPtr(resource.Description))
					}
				}
			}
		}
	}

	// Show duration information if available
	if entitlement.DurationGrant != nil && *entitlement.DurationGrant != "" {
		pterm.DefaultBasicText.Printf("Max Grant Duration: %s\n", *entitlement.DurationGrant)
	}

	pterm.Println() // Add spacing before task output
	return nil
}
