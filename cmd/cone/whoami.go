package main

import (
	"github.com/conductorone/cone/pkg/client"
	"github.com/conductorone/cone/pkg/output"
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
	ctx := cmd.Context()

	v, err := getSubViperForProfile(cmd)
	if err != nil {
		return err
	}

	clientId, clientSecret, err := getCredentials(v)
	if err != nil {
		return err
	}

	c, err := client.New(ctx, clientId, clientSecret, client.WithDebug(v.GetBool("debug")))
	if err != nil {
		return err
	}

	introspectResp, err := c.AuthIntrospect(ctx)
	if err != nil {
		return err
	}

	userResp, err := c.GetUser(ctx, introspectResp.GetUserId())
	if err != nil {
		return err
	}

	resp := C1ApiUserV1UserServiceGetResponse(*userResp)
	outputManager := output.NewManager(ctx, v)
	err = outputManager.Output(ctx, &resp)
	if err != nil {
		return err
	}

	return nil
}
