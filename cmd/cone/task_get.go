package main

import (
	"fmt"

	"github.com/conductorone/cone/pkg/output"
	"github.com/spf13/cobra"
)

func getTasksCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "",
		RunE:  getTaskRun,
	}

	addTaskIdFlag(cmd)
	return cmd
}

func getTaskRun(cmd *cobra.Command, args []string) error {
	ctx, c, v, err := cmdContext(cmd)
	if err != nil {
		return err
	}

	taskId := v.GetString(taskIdFlag)
	if len(args) == 0 && taskId == "" {
		return fmt.Errorf("expected a task id as an argument or --task-id param")
	}
	if len(args) > 0 {
		if taskId != "" {
			return fmt.Errorf("task id should be passed as an argument or --task-id param, not both")
		}
		taskId = args[0]
	}

	taskResp, err := c.GetTask(ctx, taskId)
	if err != nil {
		return err
	}

	resp := C1ApiTaskV1TaskServiceGetResponse(*taskResp)
	outputManager := output.NewManager(ctx, v)
	err = outputManager.Output(ctx, &resp)
	if err != nil {
		return err
	}

	return nil
}
