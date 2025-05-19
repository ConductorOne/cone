package main

import (
	"context"
	"fmt"

	"github.com/conductorone/cone/pkg/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func cmdContext(cmd *cobra.Command) (context.Context, client.C1Client, *viper.Viper, error) {
	ctx := cmd.Context()

	v, err := getSubViperForProfile(cmd)
	if err != nil {
		return nil, nil, nil, err
	}

	clientId, clientSecret, err := getCredentials(v)
	if err != nil {
		return nil, nil, nil, err
	}

	c, err := client.New(ctx, clientId, clientSecret, v, getCmdName(cmd))
	if err != nil {
		return nil, nil, nil, err
	}

	return ctx, c, v, nil
}

func getCmdName(cmd *cobra.Command) string {
	if cmd.HasParent() {
		return getCmdName(cmd.Parent()) + ":" + cmd.Name()
	}
	return cmd.Name()
}

func validateArgLenth(expectedCount int, args []string, cmd *cobra.Command) error {
	if len(args) == expectedCount {
		return nil
	}

	return fmt.Errorf("expected %d arguments, got %d\n%s", expectedCount, len(args), cmd.UsageString())
}

func rootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "cone",
		Short:         "Cone interacts with the ConductorOne API to manage access to entitlements.",
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	cmd.PersistentFlags().StringP("profile", "p", "default", "The config profile to use.")
	cmd.PersistentFlags().BoolP("non-interactive", "i", false, "Disable prompts.")
	cmd.PersistentFlags().String("client-id", "", "Client ID")
	cmd.PersistentFlags().String("client-secret", "", "Client secret")
	cmd.PersistentFlags().String("api-endpoint", "", "Override the API endpoint")
	cmd.PersistentFlags().StringP("output", "o", "table", "Output format. Valid values: table, json, json-pretty, wide.")
	cmd.PersistentFlags().Bool("debug", false, "Enable debug logging")

	cmd.AddCommand(getCmd())
	cmd.AddCommand(dropCmd())
	cmd.AddCommand(approveTasksCmd())
	cmd.AddCommand(denyTasksCmd())
	cmd.AddCommand(getTasksCmd())
	cmd.AddCommand(tasksCommentCmd())
	cmd.AddCommand(escalateTasksCmd())
	cmd.AddCommand(configAwsCmd())
	cmd.AddCommand(whoAmICmd())
	cmd.AddCommand(getUserCmd())
	cmd.AddCommand(searchEntitlementsCmd())
	cmd.AddCommand(tasksCmd())
	cmd.AddCommand(loginCmd())
	cmd.AddCommand(hasCmd())
	cmd.AddCommand(tokenCmd())
	cmd.AddCommand(terraformCmd())
	cmd.AddCommand(decryptCredentialCmd())
	cmd.AddCommand(awsCredentialsCmd())
	cmd.AddCommand(generateAliasCmd())

	return cmd
}
