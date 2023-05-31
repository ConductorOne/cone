package main

import (
	"context"
	"fmt"

	"github.com/conductorone/cone/internal/c1api"
	"github.com/conductorone/cone/pkg/client"
	"github.com/conductorone/cone/pkg/output"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	createdAt := r.UserView.GetUser().CreatedAt
	var ts *timestamppb.Timestamp
	if createdAt != nil {
		ts = r.UserView.GetUser().CreatedAt.(*timestamppb.Timestamp)
	}
	return [][]string{
		{
			fmt.Sprintf("%+v", r.UserView.GetUser().Id),
			fmt.Sprintf("%+v", r.UserView.GetUser().Email),
			fmt.Sprintf("%+v", r.UserView.GetUser().Status),
			fmt.Sprintf("%+v", r.UserView.GetUser().JobTitle),
			fmt.Sprintf("%+v", r.UserView.GetUser().Department),
			fmt.Sprintf("%+v", r.UserView.GetUser().EmploymentStatus),
			fmt.Sprintf("%+v", r.UserView.GetUser().EmploymentType),
			fmt.Sprintf("%+v", output.FormatTimestamp(ts)),
		},
	}
}
