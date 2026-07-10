package fslib

import "os"

func CreateDir(path string) error {
	return os.Mkdir(path, 0755)
}
