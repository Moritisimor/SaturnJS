package main

import (
	"fmt"
	"os"

	"github.com/Moritisimor/SaturnJS/fs"
	"github.com/Moritisimor/SaturnJS/io"
	"github.com/dop251/goja"
)

func main() {
	js := goja.New()
	if err := io.RegisterFuncs(js); err != nil {
		fmt.Printf("Error while registering 'io'-related functions: %s\n", err.Error())
		os.Exit(1)
	}

	if err := fs.RegisterFuncs(js); err != nil {
		fmt.Printf("Error while registering 'fs'-related functions: %s\n", err.Error())
		os.Exit(1)
	}

	js.RunString("io.println('Hello from JavaScript!')")
	js.RunString("io.println(fs.readFile(io.input('Enter file you want to read: ')))")
}
