package output

import (
	"context"
	"fmt"
	"time"

	"github.com/pterm/pterm"
)

type tableManager struct {
	area *pterm.AreaPrinter
}

func (c *tableManager) Output(ctx context.Context, out interface{}) error {
	m, ok := out.(TablePrint)
	if !ok {
		return fmt.Errorf("unexpected output model")
	}

	var preTableText string
	if p, ok := m.(PreText); ok {
		preTableText = p.Pretext()
	}

	tableData := pterm.TableData{m.Header()}
	tableData = append(tableData, m.Rows()...)
	table := pterm.DefaultTable.WithHasHeader().WithData(tableData)
	if c.area != nil {
		data, err := table.Srender()
		if err != nil {
			return err
		}
		if preTableText != "" {
			data = pterm.Sprintf("%s\n%s", preTableText, data)
		}

		c.area.Update(data)
	} else {
		if preTableText != "" {
			pterm.Println(preTableText)
		}
		err := table.Render()
		if err != nil {
			return err
		}
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

type PreText interface {
	Pretext() string
}
