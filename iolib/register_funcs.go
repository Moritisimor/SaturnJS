package iolib

import (
	"fmt"

	"github.com/dop251/goja"
)

func RegisterFuncs(js *goja.Runtime) error {
	return js.Set("io", map[string]any{
		"print": fmt.Print,
		"println": fmt.Println,
		"input": input,
	})
}
