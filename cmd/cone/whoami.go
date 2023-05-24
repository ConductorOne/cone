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

	addWaitFlag(cmd)
	addAppIdFlag(cmd)
	addEntitlementIdFlag(cmd)
	addEntitlementAliasFlag(cmd)

	return cmd
}

func whoAmIRun(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	clientId := "panicky-chimera-95203@c1dev.mstanb.dev.ductone.com/pcc"
	clientSecret := "secret-token:conductorone.com:v1:eyJrdHkiOiJPS1AiLCJjcnYiOiJFZDI1NTE5IiwieCI6IlBaZG40eHBXSkxsWGZwcThlT0wydXFvRExFNGxMWHctcWlIbXJGLWs0c3MiLCJkIjoiZ3Jlc0xKLTRpUXY5TG5VTXVGaG1JdVBKOVNMYWRSRlZuYmR2a0xMV3BVayJ9"

	c, err := client.NewC1Client(ctx, clientId, clientSecret)
	if err != nil {
		return err
	}

	_, err = c.WhoAmI(ctx)
	if err != nil {
		return err
	}

	return nil
}
