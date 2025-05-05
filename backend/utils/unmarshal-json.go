package utils

import (
	"encoding/json"
	"os"
)

func UnmarshalJSON[T any](filepath string, result *[]T) error {
	// Read JSON file
	jsonContent, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	// Unmarshall JSON into projects
	err = json.Unmarshal(jsonContent, &result)
	if err != nil {
		return err
	}

	return nil
}
