package fs

import "strings"

func readFileLines(path string) ([]string, error) {
	content, err := readFile(path)
	if err != nil {
		return []string{}, err
	}

	return strings.Split(content, "\n"), nil
}
