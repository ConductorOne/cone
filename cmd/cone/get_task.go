package main

// import (
// 	"fmt"

// 	"github.com/conductorone/cone/internal/c1api"
// 	"github.com/conductorone/cone/pkg/client"
// 	"github.com/conductorone/cone/pkg/output"
// 	"github.com/spf13/cobra"
// )

// func getTaskCmd() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "get-task",
// 		Short: "",
// 		RunE:  getTaskRun,
// 	}

// 	return cmd
// }

// func getTaskRun(cmd *cobra.Command, args []string) error {
// 	ctx := cmd.Context()

// 	v, err := getSubViperForProfile(cmd)
// 	if err != nil {
// 		return err
// 	}

// 	clientId, clientSecret, err := getCredentials(v)
// 	if err != nil {
// 		return err
// 	}

// 	if len(args) != 1 {
// 		return fmt.Errorf("expected 1 argument, got %d", len(args))
// 	}

// 	taskID := args[0]

// 	c, err := client.New(ctx, clientId, clientSecret, client.WithDebug(v.GetBool("debug")))
// 	if err != nil {
// 		return err
// 	}

// 	taskResp, err := c.GetTask(ctx, taskID)
// 	if err != nil {
// 		return err
// 	}

// 	resp := C1A(*userResp)
// 	outputManager := output.NewManager(ctx, v)
// 	err = outputManager.Output(ctx, &resp)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// type C1ApiTaskV1TaskServiceGetResponse c1api.C1ApiTaskV1Task

// func (r *C1ApiTaskV1TaskServiceGetResponse) Header() []string {
// 	return []string{
// 		"Id",
// 		"Email",
// 		"Status",
// 		"Job Title",
// 		"Department",
// 		"Employment Status",
// 		"Employment Type",
// 		"Created At",
// 	}
// }

// func (r *C1ApiUserV1UserServiceGetResponse) Rows() [][]string {
// 	return [][]string{
// 		{
// 			client.StringFromPtr(r.Id),
// 			client.StringFromPtr(r.Email),
// 			client.StringFromPtr(r.Status),
// 			client.StringFromPtr(r.JobTitle),
// 			client.StringFromPtr(r.Department),
// 			client.StringFromPtr(r.EmploymentStatus),
// 			client.StringFromPtr(r.EmploymentType),
// 			output.FormatTime(r.CreatedAt),
// 		},
// 	}
// }
