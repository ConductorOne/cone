package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/conductorone/cone/pkg/client"
	"github.com/conductorone/cone/pkg/logging"
)

var version = "dev"

func main() {
	// Create a channel to receive the signals
	signalCh := make(chan os.Signal, 1)

	// Notify the channel for specified signals
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	ctx := context.WithValue(context.Background(), client.VersionKey, version)
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		// Block until a signal is received
		<-signalCh
		cancel()
	}()

	result := runCli(ctx)
	cancel()
	os.Exit(result)
}

func runCli(ctx context.Context) int {
	cliCmd := &cobra.Command{
		Use:     "cone",
		Short:   "Cone interacts with the ConductorOne API to manage access to entitlements.",
		Version: version,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			cmd.SetContext(ctx)
			return nil
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	cliCmd.PersistentFlags().StringP("profile", "p", "default", "The config profile to use.")
	cliCmd.PersistentFlags().BoolP("non-interactive", "i", false, "Disable prompts.")
	cliCmd.PersistentFlags().String("client-id", "", "Client ID")
	cliCmd.PersistentFlags().String("client-secret", "", "Client secret")
	cliCmd.PersistentFlags().String("api-endpoint", "", "Override the API endpoint")
	cliCmd.PersistentFlags().StringP("output", "o", "table", "Output format. Valid values: table, json, json-pretty, wide.")
	cliCmd.PersistentFlags().Bool("debug", false, "Enable HTTP debug logging")
	cliCmd.PersistentFlags().String("log-level", "", "Set log level (debug, info, warn, error)")

	err := initConfig(cliCmd)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return 1
	}

	// Initialize logging based on --log-level flag
	if logLevel := viper.GetString("log-level"); logLevel != "" {
		logging.Init(logging.Level(logLevel))
	}

	cliCmd.AddCommand(getCmd())
	cliCmd.AddCommand(dropCmd())
	cliCmd.AddCommand(whoAmICmd())
	cliCmd.AddCommand(getUserCmd())
	cliCmd.AddCommand(searchEntitlementsCmd())
	cliCmd.AddCommand(tasksCmd())
	cliCmd.AddCommand(loginCmd())
	cliCmd.AddCommand(hasCmd())
	cliCmd.AddCommand(tokenCmd())
	cliCmd.AddCommand(decryptCredentialCmd())
	cliCmd.AddCommand(functionsCmd())

	err = cliCmd.ExecuteContext(ctx)
	if err != nil {
		_, _, v, _ := cmdContext(cliCmd)
		fmt.Fprintln(os.Stderr, client.HandleErrors(ctx, v, err))
		return 1
	}

	return 0
}
