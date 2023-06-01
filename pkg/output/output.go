package output

import (
	"context"

	"github.com/spf13/viper"
)

type Manager interface {
	Output(ctx context.Context, out interface{}) error
}

func NewManager(ctx context.Context, v *viper.Viper) Manager {
	switch v.GetString("output") {
	case "table":
		return &tableManager{}
	case "json":
		return &jsonManager{}
	case "json-pretty":
		return &jsonManager{pretty: true}
	default:
		return &tableManager{}
	}
}
