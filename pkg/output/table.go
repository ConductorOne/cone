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

var Checkmark = pterm.Green("✓")

const Unchecked = ""

func (c *tableManager) Output(ctx context.Context, out interface{}) error {
	var header func() []string
	var rows func() [][]string

	tablePrinter, okTable := out.(TablePrint)
	wideTablePrinter, okWideTable := out.(WideTablePrint)
	if !okTable && !okWideTable {
		return errors.New("unexpected output model")
	}
	if c.isWide && okWideTable || !okTable {
		// If we want the wide output, and the model supports it, use it. Or if the model doesn't support the table output, use the wide output.
		header = wideTablePrinter.WideHeader
		rows = wideTablePrinter.WideRows
	} else {
		// Otherwise, use the table output
		header = tablePrinter.Header
		rows = tablePrinter.Rows
	}
	var preTableText string
	if p, ok := out.(PreText); ok {
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
