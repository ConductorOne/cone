package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/cone/pkg/client"
	"github.com/conductorone/cone/pkg/output"
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

	resp := User(*userResp)
	outputManager := output.NewManager(ctx, v)
	err = outputManager.Output(ctx, &resp, output.WithTransposeTable())
	if err != nil {
		return err
	}

	return nil
}

type User shared.User

func (r *User) Header() []string {
	return []string{
		"Email",
		"Status",
		"Job Title",
		"Department",
		"Employment Status",
		"Employment Type",
		"Created At",
	}
}

func (r *User) WideHeader() []string {
	return append([]string{"Id"}, r.Header()...)
}

func (r *User) rows() []string {
	return []string{
		client.StringFromPtr(r.Email),
		userStatusToString[*r.Status],
		client.StringFromPtr(r.JobTitle),
		client.StringFromPtr(r.Department),
		client.StringFromPtr(r.EmploymentStatus),
		client.StringFromPtr(r.EmploymentType),
		output.FormatTime(r.CreatedAt),
	}
}

func (r *User) WideRows() [][]string {
	return [][]string{append([]string{client.StringFromPtr(r.ID)}, r.rows()...)}
}

func (r *User) Rows() [][]string {
	return [][]string{r.rows()}
}

var userStatusToString = map[shared.UserStatus]string{
	shared.UserStatusEnabled:  "Enabled",
	shared.UserStatusDisabled: "Disabled",
	shared.UserStatusDeleted:  "Deleted",
}
