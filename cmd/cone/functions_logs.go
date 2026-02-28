package main

import (
	"fmt"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"

	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/cone/pkg/client"
	"github.com/conductorone/cone/pkg/output"
)

func functionsLogsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logs <function-id>",
		Short: "View invocation history for a deployed function",
		Long: `Fetches and displays invocation history from ConductorOne.

Requires authentication via 'cone login'.`,
		Args: cobra.ExactArgs(1),
		RunE: functionsLogsRun,
	}

	cmd.Flags().BoolP("follow", "f", false, "Follow log output (not yet implemented)")

	return cmd
}

func functionsLogsRun(cmd *cobra.Command, args []string) error {
	functionID := args[0]

	ctx, c1Client, v, err := cmdContext(cmd)
	if err != nil {
		return fmt.Errorf("not logged in - run 'cone login' first: %w", err)
	}

	follow, _ := cmd.Flags().GetBool("follow")
	if follow {
		return fmt.Errorf("--follow is not yet implemented")
	}

	resp, err := c1Client.SDK().FunctionsInvocation.List(ctx, operations.C1APIFunctionsV1FunctionsInvocationServiceListRequest{
		FunctionID: functionID,
	})
	if err != nil {
		return fmt.Errorf("failed to list invocations: %w", err)
	}

	if resp.FunctionsInvocationServiceListResponse == nil || len(resp.FunctionsInvocationServiceListResponse.List) == 0 {
		pterm.Info.Printf("No invocations found for function %s\n", functionID)
		return nil
	}

	invocations := resp.FunctionsInvocationServiceListResponse.List

	out := &InvocationList{invocations: invocations}
	outputManager := output.NewManager(ctx, v)
	return outputManager.Output(ctx, out)
}

type InvocationList struct {
	invocations []shared.FunctionInvocation
}

func (i *InvocationList) Header() []string {
	return []string{"ID", "Status", "Created", "Commit"}
}

func (i *InvocationList) WideHeader() []string {
	return append(i.Header(), "Error")
}

func (i *InvocationList) Rows() [][]string {
	rows := make([][]string, 0, len(i.invocations))
	for _, inv := range i.invocations {
		id := client.StringFromPtr(inv.ID)
		status := ""
		if inv.Status != nil {
			status = string(*inv.Status)
		}
		created := ""
		if inv.CreatedAt != nil {
			created = inv.CreatedAt.Format("2006-01-02 15:04:05")
		}
		commit := client.StringFromPtr(inv.CommitID)
		rows = append(rows, []string{id, status, created, commit})
	}
	return rows
}

func (i *InvocationList) WideRows() [][]string {
	rows := make([][]string, 0, len(i.invocations))
	for _, inv := range i.invocations {
		id := client.StringFromPtr(inv.ID)
		status := ""
		if inv.Status != nil {
			status = string(*inv.Status)
		}
		created := ""
		if inv.CreatedAt != nil {
			created = inv.CreatedAt.Format("2006-01-02 15:04:05")
		}
		commit := client.StringFromPtr(inv.CommitID)
		errMsg := client.StringFromPtr(inv.Error)
		rows = append(rows, []string{id, status, created, commit, errMsg})
	}
	return rows
}
