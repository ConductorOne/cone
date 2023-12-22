package output

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
)

const defaultJSONError = `{"error": "%s"}`

func HandleErrors(ctx context.Context, v *viper.Viper, err error) error {
	outputType := v.GetString("output")
	if outputType != "json" && outputType != "json-pretty" {
		return err
	}

	// TODO: @anthony - handle errors better, for example, HTTP errors could be better, see client.go
	return fmt.Errorf(defaultJSONError, err.Error())
}
