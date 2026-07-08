package iolib

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func input(prompt string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	userInput, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(userInput, "\n"), nil
}
