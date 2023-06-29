package output

import (
	"context"
	"errors"
	"sort"
	"strconv"
	"time"

	"github.com/pterm/pterm"
)

type tableManager struct {
	area   *pterm.AreaPrinter
	isWide bool
}

func (c *tableManager) isInt(v string) bool {
	_, err := strconv.Atoi(v)
	return err == nil
}

func (c *tableManager) isIntColumn(tableData [][]string, col int) bool {
	for _, row := range tableData {
		if row[col] != "" || !c.isInt(row[col]) {
			return false
		}
	}
	return true
}

func (c *tableManager) sortData(header []string, tableData [][]string, out interface{}) {
	sortCol := -1
	sorter, sorterOk := out.(TableSort)
	if sorterOk {
		sortCol = sorter.SortByColumn()
	}
	if sortCol == -1 {
		for i, col := range tableData[0] {
			if col != "" {
				sortCol = i
				break
			}
		}
	}
	if c.isIntColumn(tableData, sortCol) {
		sort.SliceStable(tableData, func(i, j int) bool {
			ii, _ := strconv.Atoi(tableData[i][sortCol])
			jj, _ := strconv.Atoi(tableData[j][sortCol])
			return ii < jj
		})
	} else {
		sort.SliceStable(tableData, func(i, j int) bool {
			return tableData[i][sortCol] < tableData[j][sortCol]
		})
	}
}

func (c *tableManager) getTableData(out interface{}) (pterm.TableData, error) {
	var getHeader func() []string
	var getRows func() [][]string

	tablePrinter, okTable := out.(TablePrint)
	wideTablePrinter, okWideTable := out.(WideTablePrint)
	if !okTable && !okWideTable {
		return nil, errors.New("unexpected output model")
	}
	if c.isWide && okWideTable || !okTable {
		// If we want the wide output, and the model supports it, use it. Or if the model doesn't support the table output, use the wide output.
		getHeader = wideTablePrinter.WideHeader
		getRows = wideTablePrinter.WideRows
	} else {
		// Otherwise, use the table output
		getHeader = tablePrinter.Header
		getRows = tablePrinter.Rows
	}
	header := getHeader()
	rows := getRows()
	c.sortData(header, rows, out)

	tableData := pterm.TableData{header}
	tableData = append(tableData, rows...)
	return tableData, nil
}

func (c *tableManager) Output(ctx context.Context, out interface{}) error {
	var preTableText string
	if p, ok := out.(PreText); ok {
		preTableText = p.Pretext()
	}

	tableData, err := c.getTableData(out)
	if err != nil {
		return err
	}

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

type TableSort interface {
	SortByColumn() int
}
