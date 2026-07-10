package oslib

import (
	"os"

	"github.com/dop251/goja"
)

func RegisterFuncs(js *goja.Runtime) error {
	return js.Set("os", map[string]any{
		"cwd":     os.Getwd,
		"chdir":   os.Chdir,
		"getEnv":  os.Getenv,
		"setEnv":  os.Setenv,
		"getArgs": GetArgs,
		"getUser": GetUser,
		"getHome": GetHome,
		"exit":    os.Exit,
	})
}
