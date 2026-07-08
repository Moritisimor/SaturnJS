package main

import (
	"fmt"
	"os"

	"github.com/Moritisimor/SaturnJS/fslib"
	"github.com/Moritisimor/SaturnJS/iolib"
	"github.com/Moritisimor/SaturnJS/oslib"
	"github.com/dop251/goja"
)

func main() {
	js := goja.New()
	if err := iolib.RegisterFuncs(js); err != nil {
		fmt.Printf("Error while registering 'io'-related functions: %s\n", err.Error())
		os.Exit(1)
	}

	if err := fslib.RegisterFuncs(js); err != nil {
		fmt.Printf("Error while registering 'fs'-related functions: %s\n", err.Error())
		os.Exit(1)
	}

	if err := oslib.RegisterFuncs(js); err != nil {
		fmt.Printf("Error while registering 'os'-related functions: %s\n", err.Error())
		os.Exit(1)
	}

	js.RunString("io.println('Hello from JavaScript!')")
	js.RunString("io.println(fs.readFile(io.input('Enter file you want to read: ')))")
	js.RunString("io.println(`Using shell: ${os.getEnv('SHELL')}`)")
}
