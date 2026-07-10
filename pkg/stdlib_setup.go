package pkg

import (
	"fmt"

	"github.com/Moritisimor/SaturnJS/colorslib"
	"github.com/Moritisimor/SaturnJS/execlib"
	"github.com/Moritisimor/SaturnJS/fslib"
	"github.com/Moritisimor/SaturnJS/iolib"
	"github.com/Moritisimor/SaturnJS/netlib"
	"github.com/Moritisimor/SaturnJS/oslib"
	"github.com/dop251/goja"
)

func StdlibSetup(js *goja.Runtime) error {
	if err := iolib.RegisterFuncs(js); err != nil {
		return fmt.Errorf("Error while registering 'io'-related functions: %s\n", err.Error())
	}

	if err := fslib.RegisterFuncs(js); err != nil {
		return fmt.Errorf("Error while registering 'fs'-related functions: %s\n", err.Error())
	}

	if err := oslib.RegisterFuncs(js); err != nil {
		return fmt.Errorf("Error while registering 'os'-related functions: %s\n", err.Error())
	}

	if err := colorslib.RegisterFuncs(js); err != nil {
		return fmt.Errorf("Error while registering 'color'-related functions: %s\n", err.Error())
	}

	if err := netlib.RegisterFuncs(js); err != nil {
		return fmt.Errorf("Error while registering 'net'-related functions: %s\n", err.Error())
	}

	if err := execlib.RegisterFuncs(js); err != nil {
		return fmt.Errorf("Error while registering 'exec'-related functions. %s\n", err.Error())

	}

	return nil
}
