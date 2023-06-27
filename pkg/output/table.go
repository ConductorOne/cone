package output

import (
	"context"
	"errors"
	"time"

	"github.com/pterm/pterm"
)

type tableManager struct {
	area   *pterm.AreaPrinter
	isWide bool
}

func (c *tableManager) Output(ctx context.Context, out interface{}) error {
	var header func() []string
	var rows func() [][]string

	m, okTable := out.(TablePrint)
	widePrinter, okWide := out.(WideTablePrint)
	if !okTable && !okWide {
		return errors.New("unexpected output model")
	}
	// If we want the wide output, and the model supports it, use it. Or if the model doesn't support the table output, use the wide output.
	if c.isWide && okWide || !okTable {
		header = widePrinter.WideHeader
		rows = widePrinter.WideRows
		// Otherwise, use the table output.
	} else {
		header = m.Header
		rows = m.Rows
	}
	var preTableText string
	if p, ok := m.(PreText); ok {
		preTableText = p.Pretext()
	}

	tableData := pterm.TableData{header()}
	tableData = append(tableData, rows()...)
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

type WideTablePrint interface {
	WideHeader() []string
	WideRows() [][]string
}

type PreText interface {
	Pretext() string
}
