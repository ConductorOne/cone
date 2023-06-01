package output

import (
	"context"
	"fmt"
	"time"

	"github.com/pterm/pterm"
)

type tableManager struct {
	waiting bool
}

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

type TablePrint interface {
	Header() []string
	Rows() [][]string
}
