package fs

import "github.com/dop251/goja"

func RegisterFuncs(js *goja.Runtime) error {
	if err := js.Set("fs", map[string]any{
		"readFile": readFile,
	}); err != nil {
		return err
	}

	return nil
}
