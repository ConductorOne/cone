package output

import (
	"context"

	"github.com/pterm/pterm"
)

type Validator[T any] interface {
	IsValid(txt string) (T, bool)
	Prompt(isFirstRun bool)
}

func GetValidInput[T any](ctx context.Context, initialValue string, validator Validator[T]) (T, error) {
	userInput := pterm.DefaultInteractiveTextInput.WithMultiLine(false)
	isFirstRun := true
	txt := initialValue
	var t T
	for {
		select {
		case <-ctx.Done():
			return t, ctx.Err()
		default:
		}

		if !isFirstRun {
			var err error
			txt, err = userInput.Show()
			if err != nil {
				return t, err
			}
		}

		valid, ok := validator.IsValid(txt)
		if ok {
			return valid, nil
		}

		validator.Prompt(isFirstRun)
		isFirstRun = false
	}
}
