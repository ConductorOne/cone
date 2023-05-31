package output

import (
	"context"
	"fmt"
	"time"

	"github.com/pterm/pterm"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type tableManager struct{}

func (c *tableManager) Output(ctx context.Context, out interface{}) error {
	m, ok := out.(TablePrint)
	if !ok {
		return fmt.Errorf("unexpected output model")
	}

	tableData := pterm.TableData{m.Header()}
	tableData = append(tableData, m.Rows()...)
	err := pterm.DefaultTable.WithHasHeader().WithData(tableData).Render()
	if err != nil {
		return err
	}
	return nil
}

func FormatTime(ts *time.Time) string {
	if ts == nil {
		return ""
	}

	return ts.Format(time.RFC3339)
}

func FormatTimestamp(ts *timestamppb.Timestamp) string {
	if ts == nil {
		return ""
	}

	return ts.AsTime().Format(time.RFC3339)
}

func FromPtr(s *string) string {
	if s == nil {
		return ""
	}

	return *s
}

type TablePrint interface {
	Header() []string
	Rows() [][]string
}
