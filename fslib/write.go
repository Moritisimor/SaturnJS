package fslib

import "os"

func Write(path, content string) error {
	return os.WriteFile(path, []byte(content), 0755)
}
