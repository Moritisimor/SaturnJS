package main

import (
	"fmt"
	"os"

	"github.com/Moritisimor/SaturnJS/colorslib"
	"github.com/Moritisimor/SaturnJS/execlib"
	"github.com/Moritisimor/SaturnJS/fslib"
	"github.com/Moritisimor/SaturnJS/iolib"
	"github.com/Moritisimor/SaturnJS/netlib"
	"github.com/Moritisimor/SaturnJS/oslib"
	"github.com/dop251/goja"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("REPL mode not implemented yet!\n")
		os.Exit(1)
	}

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

	if err := colorslib.RegisterFuncs(js); err != nil {
		fmt.Printf("Error while registering 'color'-related functions: %s\n", err.Error())
		os.Exit(1)
	}

	if err := netlib.RegisterFuncs(js); err != nil {
		fmt.Printf("Error while registering 'net'-related functions: %s\n", err.Error())
		os.Exit(1)
	}

	if err := execlib.RegisterFuncs(js); err != nil {
		fmt.Printf("Error while registering 'exec'-related functions. %s\n", err.Error())
		os.Exit(1)
	}

	scriptPath := os.Args[1]
	content, err := os.ReadFile(scriptPath)
	if err != nil {
		fmt.Printf("Error while reading '%s': %s\n", scriptPath, err.Error())
		os.Exit(1)
	}

	script := string(content)
	if _, err := js.RunString(script); err != nil {
		fmt.Printf("Error while executing script: %s\n", err.Error())
		os.Exit(1)
	}
}
