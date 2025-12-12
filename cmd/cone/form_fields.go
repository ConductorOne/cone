package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/pterm/pterm"
	"github.com/spf13/viper"

	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"

	"github.com/conductorone/cone/pkg/client"
	"github.com/conductorone/cone/pkg/logging"
	"github.com/conductorone/cone/pkg/output"
)

// collectFormFields collects form field values from the user based on the form definition.
// Returns a map of field names to their values, or nil if no form fields are present.
func collectFormFields(ctx context.Context, v *viper.Viper, form *shared.FormInput) (map[string]any, error) {
	if form == nil || len(form.Fields) == 0 {
		return nil, nil
	}

	requestData := make(map[string]any)
	isNonInteractive := v.GetBool(nonInteractiveFlag)

	// Collect form data from command-line flags if provided
	formDataFlagValue := v.GetString(formDataFlag)
	formDataMap, err := parseFormDataFlag(formDataFlagValue)
	if err != nil {
		return nil, err
	}

	for _, field := range form.Fields {
		fieldName := client.StringFromPtr(field.Name)
		if fieldName == "" {
			continue
		}

		displayName := client.StringFromPtr(field.DisplayName)
		if displayName == "" {
			displayName = fieldName
		}

		description := client.StringFromPtr(field.Description)

		// Check if value was provided via flag
		if val, ok := formDataMap[fieldName]; ok {
			requestData[fieldName] = val
			continue
		}

		// Skip if non-interactive and no flag value provided
		if isNonInteractive {
			// Use default value if available
			if defaultValue := getFieldDefaultValue(field); defaultValue != nil {
				requestData[fieldName] = defaultValue
			}
			continue
		}

		// Collect value interactively
		value, err := collectFieldValue(ctx, field, displayName, description)
		if err != nil {
			return nil, fmt.Errorf("error collecting field %s: %w", fieldName, err)
		}

		if value != nil {
			requestData[fieldName] = value
		}
	}

	if len(requestData) == 0 {
		return nil, nil
	}

	return requestData, nil
}

// collectFieldValue collects a single field value from the user based on field type.
func collectFieldValue(ctx context.Context, field shared.Field, displayName, description string) (any, error) {
	// Check for default value first
	if defaultValue := getFieldDefaultValue(field); defaultValue != nil {
		// Show default value and ask for confirmation
		pterm.Info.Printf("Field '%s' has default value: %v\n", displayName, defaultValue)
		if description != "" {
			pterm.Println(description)
		}
		useDefault, err := pterm.DefaultInteractiveConfirm.Show("Use default value?")
		if err != nil {
			return nil, err
		}
		if useDefault {
			return defaultValue, nil
		}
	}

	// Collect based on field type
	switch {
	case field.StringField != nil:
		return collectStringField(ctx, field.StringField, displayName, description)
	case field.BoolField != nil:
		return collectBoolField(ctx, field.BoolField, displayName, description)
	case field.Int64Field != nil:
		return collectInt64Field(ctx, field.Int64Field, displayName, description)
	case field.StringSliceField != nil:
		return collectStringSliceField(ctx, field.StringSliceField, displayName, description)
	default:
		// Unknown field type - warn and skip to avoid breaking on new field types
		logging.Warnf("Skipping field '%s': unsupported field type. You may need to update cone to handle this field.", displayName)
		return nil, nil
	}
}

// collectStringField collects a string field value with validation.
func collectStringField(ctx context.Context, field *shared.StringField, displayName, description string) (string, error) {
	validator := StringFieldValidator{
		field:       field,
		displayName: displayName,
		description: description,
	}

	defaultValue := ""
	if field.DefaultValue != nil {
		defaultValue = *field.DefaultValue
	}

	value, err := output.GetValidInput(ctx, defaultValue, validator)
	if err != nil {
		return "", err
	}

	return value, nil
}

// collectBoolField collects a boolean field value.
func collectBoolField(ctx context.Context, field *shared.BoolField, displayName, description string) (bool, error) {
	select {
	case <-ctx.Done():
		return false, ctx.Err()
	default:
	}

	if description != "" {
		pterm.Info.Println(description)
	}

	prompt := fmt.Sprintf("Enter value for '%s' (true/false)", displayName)
	if field.DefaultValue != nil {
		prompt = fmt.Sprintf("Enter value for '%s' (true/false, default: %v)", displayName, *field.DefaultValue)
	}

	result, err := pterm.DefaultInteractiveConfirm.Show(prompt)
	if err != nil {
		return false, err
	}

	return result, nil
}

// collectInt64Field collects an int64 field value with validation.
func collectInt64Field(ctx context.Context, field *shared.Int64Field, displayName, description string) (int64, error) {
	validator := Int64FieldValidator{
		field:       field,
		displayName: displayName,
		description: description,
	}

	defaultValue := ""
	if field.DefaultValue != nil {
		defaultValue = strconv.FormatInt(*field.DefaultValue, 10)
	}

	value, err := output.GetValidInput(ctx, defaultValue, validator)
	if err != nil {
		return 0, err
	}

	return value, nil
}

// collectStringSliceField collects a string slice field value using a multi-entry loop.
// User enters one value per line, empty line finishes input.
func collectStringSliceField(ctx context.Context, field *shared.StringSliceField, displayName, description string) ([]string, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	if description != "" {
		pterm.Info.Println(description)
	}

	if len(field.DefaultValues) > 0 {
		pterm.Info.Printf("Default values for '%s': %s\n", displayName, strings.Join(field.DefaultValues, ", "))
		pterm.Info.Println("Press enter with no input to use defaults, or enter new values below.")
	}

	pterm.Info.Printf("Enter values for '%s' (one per line, empty line to finish):\n", displayName)

	var result []string
	userInput := pterm.DefaultInteractiveTextInput.WithMultiLine(false)

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		input, err := userInput.Show(fmt.Sprintf("  [%d]", len(result)+1))
		if err != nil {
			return nil, err
		}

		trimmed := strings.TrimSpace(input)
		if trimmed == "" {
			// Empty line ends input
			break
		}

		result = append(result, trimmed)
	}

	// If no values entered and defaults exist, use defaults
	if len(result) == 0 && len(field.DefaultValues) > 0 {
		return field.DefaultValues, nil
	}

	return result, nil
}

// getFieldDefaultValue extracts the default value from a field based on its type.
func getFieldDefaultValue(field shared.Field) any {
	switch {
	case field.StringField != nil && field.StringField.DefaultValue != nil:
		return *field.StringField.DefaultValue
	case field.BoolField != nil && field.BoolField.DefaultValue != nil:
		return *field.BoolField.DefaultValue
	case field.Int64Field != nil && field.Int64Field.DefaultValue != nil:
		return *field.Int64Field.DefaultValue
	case field.StringSliceField != nil && len(field.StringSliceField.DefaultValues) > 0:
		return field.StringSliceField.DefaultValues
	default:
		return nil
	}
}

// parseFormDataFlag parses the --form-data flag value as JSON.
// Expected format: '{"field1":"value1","field2":"value2"}'.
func parseFormDataFlag(formDataFlag string) (map[string]any, error) {
	if formDataFlag == "" {
		return nil, nil
	}

	result := make(map[string]any)
	if err := json.Unmarshal([]byte(formDataFlag), &result); err != nil {
		return nil, fmt.Errorf("invalid JSON in --form-data flag: %w", err)
	}

	return result, nil
}

// StringFieldValidator validates string field input.
type StringFieldValidator struct {
	field       *shared.StringField
	displayName string
	description string
}

func (v StringFieldValidator) IsValid(txt string) (string, bool) {
	if txt == "" {
		// Check if field is required
		if v.field.StringRules != nil {
			// StringRules might have required field, but we'll be lenient here
			// and allow empty if no explicit requirement
			return txt, true
		}
		return txt, true
	}

	// Apply validation rules if present
	if v.field.StringRules != nil {
		rules := v.field.StringRules
		if rules.MinLen != nil {
			minLen, err := strconv.Atoi(*rules.MinLen)
			if err == nil && len(txt) < minLen {
				return txt, false
			}
		}
		if rules.MaxLen != nil {
			maxLen, err := strconv.Atoi(*rules.MaxLen)
			if err == nil && len(txt) > maxLen {
				return txt, false
			}
		}
		// Additional validations (email, URI, etc.) could be added here
	}

	return txt, true
}

func (v StringFieldValidator) Prompt(isFirstRun bool) {
	if isFirstRun {
		if v.description != "" {
			pterm.Info.Println(v.description)
		}
		prompt := fmt.Sprintf("Enter value for '%s'", v.displayName)
		if v.field.Placeholder != nil {
			prompt = fmt.Sprintf("%s (e.g., %s)", prompt, *v.field.Placeholder)
		}
		if v.field.DefaultValue != nil {
			prompt = fmt.Sprintf("%s (default: %s)", prompt, *v.field.DefaultValue)
		}
		output.InputNeeded.Println(prompt)
	} else {
		output.InputNeeded.Println("Invalid input. Please try again.")
	}
}

// Int64FieldValidator validates int64 field input.
type Int64FieldValidator struct {
	field       *shared.Int64Field
	displayName string
	description string
}

func (v Int64FieldValidator) IsValid(txt string) (int64, bool) {
	if txt == "" {
		// Empty is invalid - default values are handled by passing them as initial value
		return 0, false
	}

	value, err := strconv.ParseInt(txt, 10, 64)
	if err != nil {
		return 0, false
	}

	// Apply validation rules if present
	if v.field.Int64Rules != nil {
		rules := v.field.Int64Rules
		if rules.Const != nil && value != *rules.Const {
			return 0, false
		}
		if rules.Lt != nil && value >= *rules.Lt {
			return 0, false
		}
		if rules.Lte != nil && value > *rules.Lte {
			return 0, false
		}
		if rules.Gt != nil && value <= *rules.Gt {
			return 0, false
		}
		if rules.Gte != nil && value < *rules.Gte {
			return 0, false
		}
	}

	return value, true
}

func (v Int64FieldValidator) Prompt(isFirstRun bool) {
	if isFirstRun {
		if v.description != "" {
			pterm.Info.Println(v.description)
		}
		prompt := fmt.Sprintf("Enter integer value for '%s'", v.displayName)
		if v.field.Placeholder != nil {
			prompt = fmt.Sprintf("%s (e.g., %s)", prompt, *v.field.Placeholder)
		}
		if v.field.DefaultValue != nil {
			prompt = fmt.Sprintf("%s (default: %d)", prompt, *v.field.DefaultValue)
		}
		output.InputNeeded.Println(prompt)
	} else {
		output.InputNeeded.Println("Invalid integer input. Please try again.")
	}
}

// isFieldRequired checks if a field is required based on its validation rules.
func isFieldRequired(field shared.Field) bool {
	switch {
	case field.StringField != nil:
		rules := field.StringField.StringRules
		if rules == nil {
			return false
		}
		// Field is required if IgnoreEmpty is not explicitly true and MinLen >= 1
		ignoreEmpty := rules.IgnoreEmpty != nil && *rules.IgnoreEmpty
		if ignoreEmpty {
			return false
		}
		if rules.MinLen != nil {
			minLen, err := strconv.Atoi(*rules.MinLen)
			if err == nil && minLen >= 1 {
				return true
			}
		}
		return false
	case field.Int64Field != nil:
		// Int64 fields don't have an IgnoreEmpty concept in the same way
		// Consider required if there are validation rules but no default
		return field.Int64Field.Int64Rules != nil && field.Int64Field.DefaultValue == nil
	case field.BoolField != nil:
		// Bool fields typically have a default (false), so rarely required
		return false
	case field.StringSliceField != nil:
		// String slice fields - check if there are rules requiring items
		return false
	default:
		return false
	}
}

// validateFormData validates that all required form fields are present.
func validateFormData(form *shared.FormInput, requestData map[string]any) error {
	if form == nil {
		return nil
	}

	for _, field := range form.Fields {
		fieldName := client.StringFromPtr(field.Name)
		if fieldName == "" {
			continue
		}

		displayName := client.StringFromPtr(field.DisplayName)
		if displayName == "" {
			displayName = fieldName
		}

		// Check if field is required
		if !isFieldRequired(field) {
			continue
		}

		// Check if value was provided
		val, hasValue := requestData[fieldName]
		if !hasValue {
			// Check if there's a default value
			if getFieldDefaultValue(field) == nil {
				return fmt.Errorf("required field '%s' is missing", displayName)
			}
			continue
		}

		// Check if value is empty
		switch v := val.(type) {
		case string:
			if v == "" {
				return fmt.Errorf("required field '%s' cannot be empty", displayName)
			}
		case []string:
			if len(v) == 0 {
				return fmt.Errorf("required field '%s' cannot be empty", displayName)
			}
		}
	}

	return nil
}
