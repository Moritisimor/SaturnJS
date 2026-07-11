package sqlitelib

import "github.com/dop251/goja"

func RegisterFuncs(js *goja.Runtime) error {
	return js.Set("sqlite", map[string]any{
		"connect": Connect,
	})
}
