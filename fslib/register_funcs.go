package fslib

import (
	"os"

	"github.com/dop251/goja"
)

func RegisterFuncs(js *goja.Runtime) error {
	return js.Set("fs", map[string]any{
		"readFile":       ReadFile,
		"readFileLines":  ReadFileLines,
		"readFileBytes":  os.ReadFile,
		"readDir":        ReadDir,
		"readJsonObject": ReadJsonObject,
		"readJsonArray":  ReadJsonArray,
	})
}
