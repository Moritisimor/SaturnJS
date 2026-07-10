package oslib

import "os"

func GetUser() string {
	return os.Getenv("USER")
}
