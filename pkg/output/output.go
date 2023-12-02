package output

import (
	"context"

	"github.com/pterm/pterm"
	"github.com/spf13/viper"
)

type Manager interface {
	Output(ctx context.Context, out interface{}, opts ...outputOption) error
}

func NewManager(ctx context.Context, v *viper.Viper) Manager {
	var area *pterm.AreaPrinter
	if v.GetBool("wait") {
		area, _ = pterm.DefaultArea.Start()
	}

	switch v.GetString("output") {
	case "table":
		return &tableManager{area: area, isWide: false}
	case "json":
		return &jsonManager{}
	case "json-pretty":
		return &jsonManager{pretty: true}
	case "wide":
		return &tableManager{area: area, isWide: true}
	default:
		return &tableManager{area: area, isWide: false}
	}
}

type outputConfig struct {
	isTransposed bool
}

type outputOption func(*outputConfig)

func WithTransposeTable() outputOption {
	return func(o *outputConfig) {
		o.isTransposed = true
	}
}

func IsWide(v *viper.Viper) bool {
	return v.GetString("output") == "wide"
}

type OutputList[T TablePrint] struct {
	Items []T
}

func (r *OutputList[T]) Header() []string {
	if len(r.Items) == 0 {
		return []string{}
	}
	return r.Items[0].Header()
}

func (r *OutputList[T]) Rows() [][]string {
	var rows [][]string
	for _, item := range r.Items {
		// Assumes that there is only one row per item
		rows = append(rows, item.Rows()[0])
	}
	return rows
}

type WideOutputList[T WideTablePrint] struct {
	Items []T
}

func (r WideOutputList[T]) WideHeader() []string {
	if len(r.Items) == 0 {
		return []string{}
	}
	return r.Items[0].WideHeader()
}

func (r WideOutputList[T]) WideRows() [][]string {
	var rows [][]string
	for _, item := range r.Items {
		// Assumes that there is only one row per item
		rows = append(rows, item.WideRows()[0])
	}
	return rows
}
