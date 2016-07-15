package util

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var input = bufio.NewReader(os.Stdin)

// read string from stdin
func ReadString() (string, error) {
	choice, err := input.ReadString('\n')

	if err != nil {
		return "", err
	}

	cleanedChoice := strings.TrimSpace(strings.TrimRight(choice, "\n"))

	return cleanedChoice, err
}

// read int from stdin
func ReadInt() (int, error) {
	choice, err := input.ReadString('\n')

	if err != nil {
		return -0, err
	}

	cleanedChoice := strings.TrimSpace(strings.TrimRight(choice, "\n"))

	return strconv.Atoi(cleanedChoice)
}
