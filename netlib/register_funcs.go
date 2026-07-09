package netlib

import "github.com/dop251/goja"

func RegisterFuncs(js *goja.Runtime) error {
	return js.Set("net", map[string]any{
		"httpGet": httpGet,
	})
}
