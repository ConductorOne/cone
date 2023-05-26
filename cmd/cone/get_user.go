package main

import (
	"context"
	"fmt"

	"github.com/conductorone/cone/pkg/client"
	"github.com/conductorone/cone/pkg/output"
	"github.com/spf13/cobra"
)

func getUserCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-user",
		Short: "",
		RunE:  getUserRun,
	}

	return cmd
}

func getUserRun(cmd *cobra.Command, args []string) error {
	ctx := context.Background()

	v, err := getSubViperForProfile(cmd)
	if err != nil {
		return err
	}

	clientId := v.GetString("client_id")
	clientSecret := v.GetString("client_secret")

	if len(args) != 1 {
		return fmt.Errorf("expected 1 argument, got %d", len(args))
	}

	userID := args[0]

	c, err := client.New(ctx, clientId, clientSecret)
	if err != nil {
		return err
	}

	userResp, err := c.GetUser(ctx, userID)
	if err != nil {
		return err
	}

	pretty := v.GetBool("pretty-output")
	output.PrintOutput(userResp, pretty)

	return nil
}
