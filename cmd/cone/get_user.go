package main

import (
	"context"
	"fmt"

	"github.com/conductorone/cone/internal/c1api"
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

	clientId, clientSecret, err := getCredentials(v)
	if err != nil {
		return err
	}

	if len(args) != 1 {
		return fmt.Errorf("expected 1 argument, got %d", len(args))
	}

	userID := args[0]

	c, err := client.New(ctx, clientId, clientSecret, client.WithDebug(v.GetBool("debug")))
	if err != nil {
		return err
	}

	userResp, err := c.GetUser(ctx, userID)
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

type C1ApiUserV1UserServiceGetResponse c1api.C1ApiUserV1UserServiceGetResponse

func (r *C1ApiUserV1UserServiceGetResponse) Header() []string {
	return []string{
		"Id",
		"Email",
		"Status",
		"Job Title",
		"Department",
		"Employment Status",
		"Employment Type",
		"Created At",
	}
}

func (r *C1ApiUserV1UserServiceGetResponse) Rows() [][]string {
	return [][]string{
		{
			output.FromPtr(r.UserView.GetUser().Id),
			output.FromPtr(r.UserView.GetUser().Email),
			output.FromPtr(r.UserView.GetUser().Status),
			output.FromPtr(r.UserView.GetUser().JobTitle),
			output.FromPtr(r.UserView.GetUser().Department),
			output.FromPtr(r.UserView.GetUser().EmploymentStatus),
			output.FromPtr(r.UserView.GetUser().EmploymentType),
			output.FormatTime(r.UserView.GetUser().CreatedAt),
		},
	}
}
