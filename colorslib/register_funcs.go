package colorslib

import (
	"github.com/Moritisimor/EpsilonFetch/pkg/color"
	"github.com/dop251/goja"
)

func RegisterFuncs(js *goja.Runtime) error {
	return js.Set("color", map[string]any{
		"redify": color.SprintRed,
		"greenify": color.SprintGreen,
		"blueify": color.SprintBlue,
		"yellowify": color.SprintYellow,
		"blackify": color.SprintBlack,
		"cyanify": color.SprintCyan,
		"magentaify": color.SprintMagenta,
	})
}
