package main

import (
	"fmt"

	"github.com/conductorone/cone/pkg/client"

	"github.com/spf13/cobra"
)

func tokenCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "token",
		Short: "",
		RunE:  tokenRun,
	}

	return cmd
}

func tokenRun(cmd *cobra.Command, args []string) error {
	ctx, _, v, err := cmdContext(cmd)
	if err != nil {
		return err
	}

	clientId, clientSecret, err := getCredentials(v)
	if err != nil {
		return err
	}

	if len(args) != 0 {
		usageErrorString := cmd.UsageString()
		return fmt.Errorf("expected 0 arguments, got %d\n"+usageErrorString, len(args))
	}

	tokenSrc, _, _, err := client.NewC1TokenSource(ctx, clientId, clientSecret)
	if err != nil {
		return err
	}

	token, err := tokenSrc.Token()
	if err != nil {
		return err
	}

	print(token)

	return nil
}
