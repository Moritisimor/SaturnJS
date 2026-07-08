package oslib

import "os"

func getArgs() []string {
	return os.Args[1:]
}
