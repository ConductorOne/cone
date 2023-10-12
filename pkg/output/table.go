package output

import (
	"context"
	"errors"
	"sort"
	"strconv"
	"time"

	"golang.org/x/exp/slices"

	"github.com/pterm/pterm"
)

type tableManager struct {
	area   *pterm.AreaPrinter
	isWide bool
}

var Checkmark = pterm.Green("✓")

const Unchecked = ""

func (c *tableManager) isInt(v string) bool {
	_, err := strconv.Atoi(v)
	return err == nil
}

func (c *tableManager) isIntColumn(tableData [][]string, col int) bool {
	for _, row := range tableData {
		if row[col] != "" && !c.isInt(row[col]) {
			return false
		}
	}
	return true
}
func (c *tableManager) sortData(header []string, tableData [][]string, out interface{}) {
	if len(tableData) == 0 {
		return
	}

	sortCols := []int{}
	sorter, sorterOk := out.(TableSort)
	if sorterOk {
		for _, col := range sorter.OrderedSortColumns() {
			sortCol := slices.Index(header, col)
			if sortCol != -1 {
				sortCols = append(sortCols, sortCol)
			}
		}
	}

	// If we didn't find any columns to sort by, just sort by the first non-empty column
	if len(sortCols) == 0 {
		for i, col := range tableData[0] {
			if col != "" {
				sortCols = append(sortCols, i)
				break
			}
		}
	}

	sort.SliceStable(tableData, func(i, j int) bool {
		for _, sortCol := range sortCols {
			if c.isIntColumn(tableData, sortCol) {
				ii, _ := strconv.Atoi(tableData[i][sortCol])
				jj, _ := strconv.Atoi(tableData[j][sortCol])
				if ii != jj {
					return ii < jj
				}
			} else {
				if tableData[i][sortCol] != tableData[j][sortCol] {
					return tableData[i][sortCol] < tableData[j][sortCol]
				}
			}
		}
		return false
	})

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

// Transpose Table is for single object outputs, instead of a single row table.
func transposeTable(table [][]string) [][]string {
	rows := len(table)
	cols := len(table[0])

	transposed := make([][]string, cols)
	for i := 0; i < cols; i++ {
		transposed[i] = make([]string, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transposed[j][i] = table[i][j]
		}
	}
	for i := 0; i < cols; i++ {
		transposed[i][0] = "\x1b[0;36m" + transposed[i][0] + "\x1b[0m"
	}

	return transposed
}

func (c *tableManager) Output(ctx context.Context, out interface{}, opts ...outputOption) error {
	var preTableText string
	if p, ok := out.(PreText); ok {
		preTableText = p.Pretext()
	}

	outputConfig := &outputConfig{}
	for _, opt := range opts {
		opt(outputConfig)
	}

	tableData, err := c.getTableData(out)
	if err != nil {
		return err
	}

	table := &pterm.DefaultTable
	if outputConfig.isTransposed {
		tableData = transposeTable(tableData)
	} else {
		table = table.WithHasHeader()
	}
	table = table.WithData(tableData)

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
	OrderedSortColumns() []string
}
