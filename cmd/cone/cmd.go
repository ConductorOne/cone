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

	c, err := client.New(ctx, clientId, clientSecret, v)
	if err != nil {
		return nil, nil, nil, err
	}

	return ctx, c, v, nil
}

func validateArgLenth(expectedCount int, args []string, cmd *cobra.Command) error {
	if len(args) == expectedCount {
		return nil
	}

	return fmt.Errorf("expected %d arguments, got %d\n%s", expectedCount, len(args), cmd.UsageString())
}
