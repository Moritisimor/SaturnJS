package oslib

import "os"

func getHome() (string, error) {
	return os.UserHomeDir()
}
