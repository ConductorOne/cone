package resource

import (
	"encoding/json"
	"os"
)

type terraformState struct {
	Outputs map[string]interface{} `json:"outputs"`
}

func mapInputs(inputs []templateData) map[string]templateData {
	inputMap := make(map[string]templateData)
	for _, input := range inputs {
		inputMap[input.GetName()] = input
	}
	return inputMap
}

func createResource(input templateData) error {
	switch input.GetType() {
	case TerraformAppType:

	}
	return nil
}

func readTfState(terraformDir string, inputs []templateData) error {
	inputMap := mapInputs(inputs)
	fileBytes, err := os.ReadFile(terraformDir + "/terraform.tfstate")
	if err != nil {
		return err
	}
	var data terraformState
	err = json.Unmarshal(fileBytes, &data)
	if err != nil {
		return err
	}

	for key, value := range data.Outputs {
		if input, ok := inputMap[key]; ok {
			input.GetResourceType()
		}

	}
	return nil
}
