package main

import (
	"github.com/Moritisimor/SaturnJS/io"
	"github.com/dop251/goja"
)

func main() {
	js := goja.New()
	io.RegisterFuncs(js)

	js.RunString("io.println('Hello from JavaScript!')")
	js.RunString("io.println(io.input('Enter something: '))")
}
