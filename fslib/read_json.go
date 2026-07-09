package fslib

import (
	"encoding/json"
	"os"
)

func readJsonObject(path string) (map[string]any, error) {
	decodedJson := map[string]any{}
	content, err := os.ReadFile(path)
	if err != nil {
		return decodedJson, err
	}

	err = json.Unmarshal(content, &decodedJson)
	return decodedJson, err
}

func readJsonArray(path string) ([]any, error) {
	decodedJson := []any{}
	content, err := os.ReadFile(path)
	if err != nil {
		return decodedJson, err
	}

	err = json.Unmarshal(content, &decodedJson)
	return decodedJson, err
}
