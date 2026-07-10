package pkg

import (
	"fmt"
	"os"

	"github.com/dop251/goja"
)

func RunScript(js *goja.Runtime, scriptPath string) error {
	content, err := os.ReadFile(scriptPath)
	if err != nil {
		return fmt.Errorf("Error while reading '%s': %s\n", scriptPath, err.Error())
	}

	script := string(content)
	if _, err = js.RunString(script); err != nil {
		return fmt.Errorf("Error while executing script: %s\n", err.Error())
	}

	return nil
}
