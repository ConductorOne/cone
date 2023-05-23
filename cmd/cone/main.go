package main

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "dev"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := initConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	cliCmd := &cobra.Command{
		Use:     "cone",
		Short:   "cone is... a cone", // TODO: Change this
		Version: version,
	}

	cliCmd.PersistentFlags().StringP("profile", "p", "default", "The conig profile to use.")
	cliCmd.PersistentFlags().BoolP("non-interactive", "i", false, "Disable prompts.")

	cliCmd.AddCommand(getCmd())
	cliCmd.AddCommand(dropCmd())

	err = cliCmd.ExecuteContext(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
