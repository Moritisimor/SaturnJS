package fslib

import "strings"

func ReadFileLines(path string) ([]string, error) {
	content, err := ReadFile(path)
	if err != nil {
		return []string{}, err
	}

	return strings.Split(content, "\n"), nil
}
