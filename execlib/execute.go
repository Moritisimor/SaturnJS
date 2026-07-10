package execlib

import (
	"os"
	"os/exec"
)

func Execute(program string, args []string) (int, error) {
	cmd := exec.Command(program, args...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	return cmd.ProcessState.ExitCode(), err
}
