package randlib

import "github.com/dop251/goja"

func RegisterFuncs(js *goja.Runtime) error {
	return js.Set("rand", map[string]any{
		"randInt":   RandInt,
		"randFloat": RandFloat,
		"choice": Choice,
	})
}
