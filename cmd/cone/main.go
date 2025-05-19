package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/conductorone/cone/pkg/client"
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
	cliCmd := rootCmd()
	cliCmd.Version = version
	cliCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		cmd.SetContext(ctx)
		return nil
	}

	err := initConfig(cliCmd)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return 1
	}

	err = cliCmd.ExecuteContext(ctx)
	if err != nil {
		_, _, v, _ := cmdContext(cliCmd)
		fmt.Fprintln(os.Stderr, client.HandleErrors(ctx, v, err))
		return 1
	}

	return 0
}
