package fslib

import "encoding/json"

func WriteJson(path string, jsonObject any) error {
	content, err := json.MarshalIndent(jsonObject, "", "    ")
	if err != nil {
		return err
	}

	return Write(path, string(content))
}
