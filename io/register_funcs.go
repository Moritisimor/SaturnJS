package io

import (
	"fmt"

	"github.com/dop251/goja"
)

func RegisterFuncs(js *goja.Runtime) error {
	if err := js.Set("io", map[string]any{
		"print": fmt.Print,
		"println": fmt.Println,
		"input": input,
	}); err != nil {
		return err
	}

	return nil
}
