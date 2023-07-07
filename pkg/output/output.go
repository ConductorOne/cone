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
