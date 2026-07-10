package netlib

import (
	"github.com/Moritisimor/SaturnJS/netlib/httplib"
	"github.com/dop251/goja"
)

func RegisterFuncs(js *goja.Runtime) error {
	return js.Set("net", map[string]any{
		"http": map[string]any{
			"get": httplib.Get,
			"post": httplib.Post,
			"request": httplib.Request,
		},

		"tcp": map[string]any{}, // Only placeholders for now
		"udp": map[string]any{},
	})
}
