package main

import (
	"fmt"

	"github.com/conductorone/cone/internal/c1api"
	"github.com/conductorone/cone/pkg/client"
	"github.com/conductorone/cone/pkg/output"
	"github.com/spf13/cobra"
)

func tasksCommentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "task-comment",
		Short: "",
		RunE:  tasksCommentRun,
	}

	addTaskCommentFlag(cmd)

	return cmd
}

func tasksCommentRun(cmd *cobra.Command, args []string) error {
	ctx, c, v, err := cmdContext(cmd)
	if err != nil {
		return err
	}

	if len(args) != 2 {
		return fmt.Errorf("expected 2 argument, got %d", len(args))
	}
	taskID := args[0]
	comment := args[1]

	userResp, err := c.CommentOnTask(ctx, taskID, comment)
	if err != nil {
		return err
	}

	resp := C1ApiTaskV1TaskActionsServiceCommentResponse(*userResp)
	outputManager := output.NewManager(ctx, v)
	err = outputManager.Output(ctx, &resp)
	if err != nil {
		return err
	}

	return nil
}

type C1ApiTaskV1TaskActionsServiceCommentResponse c1api.C1ApiTaskV1TaskActionsServiceCommentResponse

func (t *C1ApiTaskV1TaskActionsServiceCommentResponse) Header() []string {
	return []string{
		"Id",
		"Display Name",
		"State",
	}
}

func (t *C1ApiTaskV1TaskActionsServiceCommentResponse) Rows() [][]string {
	return [][]string{
		{
			client.StringFromPtr(t.TaskView.GetTask().NumericId),
			client.StringFromPtr(t.TaskView.GetTask().DisplayName),
			client.StringFromPtr(t.TaskView.GetTask().State),
		},
	}
}
