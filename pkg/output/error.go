package output

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
)

const defaultJSONError = `{"error": "unable to marshal error to JSON %s"}`

type JSONError struct {
	Error string `json:"error"`
}

func HandleErrors(ctx context.Context, v *viper.Viper, input error) error {
	outputType := v.GetString("output")
	if outputType != "json" && outputType != JSONPretty {
		return input
	}
	// TODO: @anthony - handle errors better, for example, HTTP errors could be better, see client.go
	jsonError, err := MakeJSONFromInterface(ctx, JSONError{Error: input.Error()}, outputType == JSONPretty)
	if err != nil {
		return fmt.Errorf(defaultJSONError, input.Error())
	}

	return fmt.Errorf(string(jsonError))
}
