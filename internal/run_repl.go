package internal

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Moritisimor/EpsilonFetch/pkg/color"
	"github.com/chzyer/readline"
	"github.com/dop251/goja"
)

func RunRepl(js *goja.Runtime) error {
	editor, err := readline.New(fmt.Sprintf(
		"🪐 %s %s %s ",
		color.SprintGreen("SaturnJS"),
		color.SprintBlue("REPL"),
		color.SprintMagenta(">>"),
	))

	if err != nil {
		return fmt.Errorf("Error while creating readline instance: %s", err.Error())
	}

	for {
		input, err := editor.Readline()
		if err != nil {
			if errors.Is(err, readline.ErrInterrupt) {
				continue
			}

			return fmt.Errorf("Error while reading line: %s\n", err.Error())
		}

		if strings.TrimSpace(input) == "exit" {
			break
		}

		if strings.TrimSpace(input) == "clear" {
			fmt.Print("\033[H\033[2J\033[3J")
			continue
		}

		val, err := js.RunString(input)
		if err != nil {
			color.PrintRedln("Error while running JS: " + err.Error())
			continue
		}

		fmt.Print(color.SprintCyan("Evaluates to: "))
		fmt.Println(val)
	}

	color.PrintBlueln("Goodbye!")
	return nil
}
