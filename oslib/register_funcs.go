package oslib

import (
	"os"

	"github.com/dop251/goja"
)

func RegisterFuncs(js *goja.Runtime) error {
	return js.Set("os", map[string]any{
		"getEnv": os.Getenv,
		"setEnv": os.Setenv,
		"getArgs": getArgs,
		"exit": os.Exit,
	})
}
