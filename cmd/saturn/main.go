package main

import (
	"fmt"
	"os"

	"github.com/Moritisimor/SaturnJS/internal"
	"github.com/dop251/goja"
)

func main() {
	js := goja.New()
	if err := internal.StdlibSetup(js); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		if err := internal.RunRepl(js); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		return
	}

	if err := internal.RunScript(js, os.Args[1]); err != nil {
		fmt.Println(err.Error())
	}
}
