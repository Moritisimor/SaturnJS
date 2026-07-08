package fslib

import (
	"os"

	"github.com/dop251/goja"
)

func RegisterFuncs(js *goja.Runtime) error {
	if err := js.Set("fs", map[string]any{
		"readFile": readFile,
		"readFileLines": readFileLines,
		"readFileBytes": os.ReadFile,
		"readDir": readDir,
	}); err != nil {
		return err
	}

	return nil
}
