package main

import (
	"github.com/Moritisimor/SaturnJS/fs"
	"github.com/Moritisimor/SaturnJS/io"
	"github.com/dop251/goja"
)

func main() {
	js := goja.New()
	io.RegisterFuncs(js)
	fs.RegisterFuncs(js)

	js.RunString("io.println('Hello from JavaScript!')")
	js.RunString("io.println(fs.readFile(io.input('Enter file you want to read: ')))")
}
