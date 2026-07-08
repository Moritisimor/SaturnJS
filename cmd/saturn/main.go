package main

import (
	"fmt"
	"github.com/dop251/goja"
)

func main() {
	js := goja.New()
	js.Set("println", fmt.Println)
	js.RunString("println('Hello from JavaScript!')")
}
