package main

import (
	"fmt"
	"github.com/dop251/goja"
)

func main() {
	js := goja.New()
	js.Set("io", map[string]any{
		"println": fmt.Println,
		"print": fmt.Print,
	})

	js.RunString("io.println('Hello from JavaScript!')")
}
