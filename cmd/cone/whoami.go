package main

import (
	"github.com/conductorone/cone/internal/c1api"
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

	whoamiResp, err := c.WhoAmI(ctx)
	if err != nil {
		return err
	}

	resp := C1ApiAuthV1IntrospectResponse(*whoamiResp)
	outputManager := output.NewManager(ctx, v)
	err = outputManager.Output(ctx, &resp)
	if err != nil {
		return err
	}

	return nil
}

type C1ApiAuthV1IntrospectResponse c1api.C1ApiAuthV1IntrospectResponse

func (r *C1ApiAuthV1IntrospectResponse) Header() []string {
	return []string{"PrincipleId", "UserId", "AccessTokenId"}
}
func (r *C1ApiAuthV1IntrospectResponse) Rows() [][]string {
	return [][]string{{
		client.StringFromPtr(r.PrincipleId),
		client.StringFromPtr(r.UserId),
		client.StringFromPtr(r.AccessTokenId),
	}}
}
