package main

import (
	"context"

	"github.com/conductorone/cone/pkg/client"
	"github.com/spf13/cobra"
)

func whoAmICmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "whoami",
		Short: "",
		RunE:  whoAmIRun,
	}

	return cmd
}

func whoAmIRun(cmd *cobra.Command, args []string) error {
	ctx := context.Background()

	v, err := getSubViperForProfile(cmd)
	if err != nil {
		return err
	}

	clientId := v.GetString("client_id")
	clientSecret := v.GetString("client_secret")

	c, err := client.New(ctx, clientId, clientSecret)
	if err != nil {
		return err
	}

	_, err = c.WhoAmI(ctx)
	if err != nil {
		return err
	}

	return nil
}
