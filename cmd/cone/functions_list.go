package main

import (
	"fmt"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"

	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/cone/pkg/client"
	"github.com/conductorone/cone/pkg/output"
)

func functionsListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List deployed functions",
		Long: `Lists all functions deployed to ConductorOne.

Requires authentication via 'cone login'.`,
		RunE: functionsListRun,
	}

	return cmd
}

func functionsListRun(cmd *cobra.Command, args []string) error {
	ctx, c1Client, v, err := cmdContext(cmd)
	if err != nil {
		return fmt.Errorf("not logged in - run 'cone login' first: %w", err)
	}

	resp, err := c1Client.SDK().Functions.ListFunctions(ctx)
	if err != nil {
		return fmt.Errorf("failed to list functions: %w", err)
	}

	if resp.FunctionsServiceListFunctionsResponse == nil || len(resp.FunctionsServiceListFunctionsResponse.List) == 0 {
		pterm.Info.Println("No functions found")
		return nil
	}

	functions := resp.FunctionsServiceListFunctionsResponse.List

	out := &FunctionsList{functions: functions}
	outputManager := output.NewManager(ctx, v)
	return outputManager.Output(ctx, out)
}

type FunctionsList struct {
	functions []shared.Function
}

func (f *FunctionsList) Header() []string {
	return []string{"ID", "Name", "Description", "Type"}
}

func (f *FunctionsList) WideHeader() []string {
	return append(f.Header(), "Published Commit", "Created")
}

func (f *FunctionsList) Rows() [][]string {
	rows := make([][]string, 0, len(f.functions))
	for _, fn := range f.functions {
		id := client.StringFromPtr(fn.ID)
		name := client.StringFromPtr(fn.DisplayName)
		desc := client.StringFromPtr(fn.Description)
		fnType := ""
		if fn.FunctionType != nil {
			fnType = string(*fn.FunctionType)
		}
		rows = append(rows, []string{id, name, desc, fnType})
	}
	return rows
}

func (f *FunctionsList) WideRows() [][]string {
	rows := make([][]string, 0, len(f.functions))
	for _, fn := range f.functions {
		id := client.StringFromPtr(fn.ID)
		name := client.StringFromPtr(fn.DisplayName)
		desc := client.StringFromPtr(fn.Description)
		fnType := ""
		if fn.FunctionType != nil {
			fnType = string(*fn.FunctionType)
		}
		publishedCommit := client.StringFromPtr(fn.PublishedCommitID)
		created := ""
		if fn.CreatedAt != nil {
			created = fn.CreatedAt.Format("2006-01-02 15:04")
		}
		rows = append(rows, []string{id, name, desc, fnType, publishedCommit, created})
	}
	return rows
}
