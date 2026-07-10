package main

import (
	"fmt"
	"os"

	"github.com/Moritisimor/SaturnJS/pkg"
	"github.com/dop251/goja"
)

func main() {
	js := goja.New()
	if err := pkg.StdlibSetup(js); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		if err := pkg.RunRepl(js); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		return
	}

	if err := pkg.RunScript(js, os.Args[1]); err != nil {
		fmt.Println(err.Error())
	}
}
