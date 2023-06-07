package main

import (
	"fmt"

	"github.com/conductorone/cone/internal/c1api"
	"github.com/conductorone/cone/pkg/client"
	"github.com/conductorone/cone/pkg/output"
	"github.com/spf13/cobra"
)

func getUserCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-user",
		Short: "Get a user by id",
		RunE:  getUserRun,
	}

	return cmd
}

func getUserRun(cmd *cobra.Command, args []string) error {
	ctx, c, v, err := cmdContext(cmd)
	if err != nil {
		return err
	}

	if len(args) != 1 {
		return fmt.Errorf("expected 1 argument, got %d", len(args))
	}
	userID := args[0]

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

type C1ApiUserV1UserServiceGetResponse c1api.C1ApiUserV1User

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
			client.StringFromPtr(r.Id),
			client.StringFromPtr(r.Email),
			client.StringFromPtr(r.Status),
			client.StringFromPtr(r.JobTitle),
			client.StringFromPtr(r.Department),
			client.StringFromPtr(r.EmploymentStatus),
			client.StringFromPtr(r.EmploymentType),
			output.FormatTime(r.CreatedAt),
		},
	}
}
