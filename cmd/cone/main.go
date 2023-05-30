package main

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "dev"

func main() {
	os.Exit(runCli())
}

func runCli() int {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cliCmd := &cobra.Command{
		Use:     "cone",
		Short:   "cone is... a cone", // TODO: Change this
		Version: version,
	}

	cliCmd.PersistentFlags().StringP("profile", "p", "default", "The conig profile to use.")
	cliCmd.PersistentFlags().BoolP("non-interactive", "i", false, "Disable prompts.")
	cliCmd.PersistentFlags().String("client-id", "", "Client ID")
	cliCmd.PersistentFlags().String("client-secret", "", "Client secret")
	cliCmd.PersistentFlags().String("config-path", "", "path to config file")
	cliCmd.PersistentFlags().Bool("pretty-output", false, "Whether to pretty print output")

	err := initConfig(cliCmd)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return 1
	}

	cliCmd.AddCommand(getCmd())
	cliCmd.AddCommand(dropCmd())
	cliCmd.AddCommand(whoAmICmd())
	cliCmd.AddCommand(getUserCmd())

	err = cliCmd.ExecuteContext(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return 1
	}

	return 0
}
