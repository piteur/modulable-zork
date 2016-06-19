package util

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var input = bufio.NewReader(os.Stdin)

// read string from stdin
func ReadString() (choice string, err error) {
	// check on windows
	choice, err = input.ReadString('\n')

	return strings.TrimRight(choice, "\n"), err
}

// read int from stdin
func ReadInt() (choice int, err error) {
	// check on windows
	stringChoice, err := input.ReadString('\n')

	if err != nil {
		return -0, err
	}

	return strconv.Atoi(strings.TrimRight(stringChoice, "\n"))
}
