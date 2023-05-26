package output

import (
	"encoding/json"
	"fmt"
	"os"
)

func PrintOutput(data interface{}, pretty bool) error {
	if pretty {
		prettyJSON, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			return err
		}
		fmt.Fprint(os.Stdout, string(prettyJSON))
		return nil
	}

	plainJSON, err := json.Marshal(data)
	if err != nil {
		return err
	}
	fmt.Fprint(os.Stdout, string(plainJSON))

	return nil
}
