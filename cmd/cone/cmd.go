package main

import (
	"context"
	"fmt"
	"os"

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

	// Priority 1: CONDUCTORONE_ACCESS_TOKEN -- pre-exchanged bearer token
	if accessToken := os.Getenv(client.EnvAccessToken); accessToken != "" {
		clientID := v.GetString("client-id")
		if clientID == "" {
			clientID = os.Getenv(client.EnvClientID)
		}
		c, err := client.NewWithAccessToken(ctx, accessToken, clientID, v, getCmdName(cmd))
		if err != nil {
			return nil, nil, nil, err
		}
		return ctx, c, v, nil
	}

	// Priority 2: CONDUCTORONE_OIDC_TOKEN -- RFC 8693 token exchange
	if oidcToken := os.Getenv(client.EnvOIDCToken); oidcToken != "" {
		clientID := v.GetString("client-id")
		if clientID == "" {
			clientID = os.Getenv(client.EnvClientID)
		}
		if clientID == "" {
			return nil, nil, nil, fmt.Errorf("%s requires --client-id, CONE_CLIENT_ID, or %s", client.EnvOIDCToken, client.EnvClientID)
		}
		c, err := client.NewWithOIDCToken(ctx, oidcToken, clientID, v, getCmdName(cmd))
		if err != nil {
			return nil, nil, nil, err
		}
		return ctx, c, v, nil
	}

	// Priority 3: existing client-id + client-secret flow
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
