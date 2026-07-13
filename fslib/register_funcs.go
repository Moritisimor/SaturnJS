package fslib

import (
	"os"

	"github.com/dop251/goja"
)

func RegisterFuncs(js *goja.Runtime) error {
	return js.Set("fs", map[string]any{
		"readFile":      ReadFile,
		"readFileLines": ReadFileLines,
		"readFileBytes": os.ReadFile,
		"readDir":       ReadDir,

		"readJsonObject": ReadJsonObject,
		"readJsonArray":  ReadJsonArray,

		"createFile": CreateFile,
		"createDir":  CreateDir,

		"deleteFile": os.Remove,
		"deleteDir":  os.RemoveAll,

		"writeFile": Write,
		"writeJson": WriteJson,

		"move": os.Rename,
		"copy": Copy,

		"pathSep": string(os.PathSeparator),
	})
}
