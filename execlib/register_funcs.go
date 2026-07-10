package execlib

import "github.com/dop251/goja"

func RegisterFuncs(js *goja.Runtime) error {
	return js.Set("exec", map[string]any{
		"execute": Execute,
	})
}
