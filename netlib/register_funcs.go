package netlib

import "github.com/dop251/goja"

func RegisterFuncs(js *goja.Runtime) error {
	if err := js.Set("net", map[string]any{
		"httpGet": httpGet,
	}); err != nil {
		return err
	}

	return nil
}
