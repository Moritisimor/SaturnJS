package oslib

import "os"

func getUser() string {
	return os.Getenv("USER")
}
