package fslib

import (
	"io"
	"os"
)

func Copy(oldPath, newPath string) error {
	oldFile, err := os.Open(oldPath)
	if err != nil {
		return err
	}

	defer oldFile.Close()
	newFile, err := os.Create(newPath)
	if err != nil {
		return err
	}

	defer newFile.Close()
	_, err = io.Copy(newFile, oldFile)
	return err
}
