package fslib

import (
	"encoding/json"
	"os"
)

func ReadJsonObject(path string) (map[string]any, error) {
	decodedJson := map[string]any{}
	content, err := os.ReadFile(path)
	if err != nil {
		return decodedJson, err
	}

	err = json.Unmarshal(content, &decodedJson)
	return decodedJson, err
}

func ReadJsonArray(path string) ([]any, error) {
	decodedJson := []any{}
	content, err := os.ReadFile(path)
	if err != nil {
		return decodedJson, err
	}

	err = json.Unmarshal(content, &decodedJson)
	return decodedJson, err
}
