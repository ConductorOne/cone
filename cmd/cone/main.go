package main

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var version = "dev"

func main() {
	os.Exit(runCli())
}

func runCli() int {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := initConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return 1
	}

	cliCmd := &cobra.Command{
		Use:     "cone",
		Short:   "cone is... a cone", // TODO: Change this
		Version: version,
	}

	cliCmd.PersistentFlags().StringP("profile", "p", "default", "The conig profile to use.")
	cliCmd.PersistentFlags().BoolP("non-interactive", "i", false, "Disable prompts.")
	viper.GetString("profile")
	cliCmd.AddCommand(getCmd())
	cliCmd.AddCommand(dropCmd())
	cliCmd.AddCommand(whoAmICmd())

	err = cliCmd.ExecuteContext(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return 1
	}

	return 0
}
