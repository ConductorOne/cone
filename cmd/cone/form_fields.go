package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/pterm/pterm"
	"github.com/spf13/viper"

	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"

	"github.com/conductorone/cone/pkg/client"
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
	formDataMap := parseFormDataFlag(formDataFlagValue)

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
		return nil, fmt.Errorf("unsupported field type for field: %s", displayName)
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

// collectStringSliceField collects a string slice field value.
func collectStringSliceField(ctx context.Context, field *shared.StringSliceField, displayName, description string) ([]string, error) {
	if description != "" {
		pterm.Info.Println(description)
	}

	prompt := fmt.Sprintf("Enter values for '%s' (comma-separated)", displayName)
	if len(field.DefaultValues) > 0 {
		prompt = fmt.Sprintf("Enter values for '%s' (comma-separated, default: %s)", displayName, strings.Join(field.DefaultValues, ", "))
	}

	userInput := pterm.DefaultInteractiveTextInput.WithMultiLine(false)
	input, err := userInput.Show(prompt)
	if err != nil {
		return nil, err
	}

	if input == "" {
		if len(field.DefaultValues) > 0 {
			return field.DefaultValues, nil
		}
		return []string{}, nil
	}

	// Split by comma and trim whitespace
	values := strings.Split(input, ",")
	result := make([]string, 0, len(values))
	for _, v := range values {
		trimmed := strings.TrimSpace(v)
		if trimmed != "" {
			result = append(result, trimmed)
		}
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

// parseFormDataFlag parses the --form-data flag value.
// Expected format: "field1=value1,field2=value2" or JSON object.
func parseFormDataFlag(formDataFlag string) map[string]any {
	if formDataFlag == "" {
		return nil
	}

	result := make(map[string]any)

	// Try parsing as comma-separated key=value pairs
	pairs := strings.Split(formDataFlag, ",")
	for _, pair := range pairs {
		parts := strings.SplitN(strings.TrimSpace(pair), "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			result[key] = value
		}
	}

	return result
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
		// Allow empty if there's a default value
		if v.field.DefaultValue != nil {
			return *v.field.DefaultValue, true
		}
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

// getFormFromTask retrieves the form definition from a task.
// This is used when the form is only available after task creation.
func getFormFromTask(task *shared.Task) *shared.FormInput {
	if task == nil {
		return nil
	}
	return task.Form
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

		// Check if field is required (this is a simplified check)
		// In practice, you'd check field rules for required status
		_, hasValue := requestData[fieldName]
		if !hasValue {
			// Check if there's a default value
			if getFieldDefaultValue(field) == nil {
				displayName := client.StringFromPtr(field.DisplayName)
				if displayName == "" {
					displayName = fieldName
				}
				return fmt.Errorf("required field '%s' is missing", displayName)
			}
		}
	}

	return nil
}

