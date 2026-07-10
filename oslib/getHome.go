package oslib

import "os"

func GetHome() (string, error) {
	return os.UserHomeDir()
}
