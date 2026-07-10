package fslib

import "os"

func CreateFile(path string) error {
	_, err := os.Create(path)
	return err
}
