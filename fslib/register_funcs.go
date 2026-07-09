package fslib

import (
	"os"

	"github.com/dop251/goja"
)

func RegisterFuncs(js *goja.Runtime) error {
	return js.Set("fs", map[string]any{
		"readFile": readFile,
		"readFileLines": readFileLines,
		"readFileBytes": os.ReadFile,
		"readDir": readDir,
	})
}
