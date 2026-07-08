package oslib

import (
	"os"

	"github.com/dop251/goja"
)

func RegisterFuncs(js *goja.Runtime) error {
	if err := js.Set("os", map[string]any{
		"getEnv": os.Getenv,
		"setEnv": os.Setenv,
	}); err != nil {
		return err
	}

	return nil
}
