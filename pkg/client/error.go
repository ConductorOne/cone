package client

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/conductorone/cone/pkg/output"
	"github.com/spf13/viper"
)

var (
	ErrHTTPError = fmt.Errorf("HTTP error")
)

const defaultJSONError = `{"error": "unable to marshal error to JSON %s"}`

type JSONError struct {
	Error string `json:"error"`
}

type HTTPError struct {
	StatusCode int    `json:"status_code"`
	Body       string `json:"body"`
	Err        error  `json:"-"`
}

func NewHTTPError(resp *http.Response) *HTTPError {
	// This is added temporarily to ensure we return an error if we get a non-success status code.
	// Eventually (ideally), we'll be generating this error handling as part of the SDK
	if resp.StatusCode >= http.StatusBadRequest {
		var httpErr HTTPError
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			httpErr = HTTPError{
				StatusCode: resp.StatusCode,
				Body:       fmt.Errorf("unable to read response body: %w", err).Error(),
				Err:        ErrHTTPError,
			}
		} else {
			httpErr = HTTPError{
				StatusCode: resp.StatusCode,
				Body:       string(body),
				Err:        ErrHTTPError,
			}
		}
		return &httpErr
	}

	return nil
}

func (e *HTTPError) Unwrap() error {
	return e.Err
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("%d: %s", e.StatusCode, e.Body)
}

func HandleErrors(ctx context.Context, v *viper.Viper, input error) error {
	outputType := v.GetString("output")
	if outputType != "json" && outputType != output.JSONPretty {
		return input
	}
	var jsonError []byte

	if errors.Is(input, ErrHTTPError) {
		httpErr := input.(*HTTPError)
		jsonError, err := output.MakeJSONFromInterface(ctx, httpErr, outputType == output.JSONPretty)
		if err != nil {
			return fmt.Errorf(defaultJSONError, httpErr.Error())
		}
		return fmt.Errorf(string(jsonError))
	}
	jsonError, err := output.MakeJSONFromInterface(ctx, JSONError{Error: input.Error()}, outputType == output.JSONPretty)
	if err != nil {
		return fmt.Errorf(defaultJSONError, input.Error())
	}

	return fmt.Errorf(string(jsonError))
}
