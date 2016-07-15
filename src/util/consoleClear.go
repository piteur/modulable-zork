package util

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func ClearConsole() {
	command := ""

	switch runtime.GOOS {
		case "darwin", "freebsd", "linux", "netbsd", "openbsd":
			command = "clear"
		case "windows":
			command = "cls"
	}

	if command != "" {
		cmd := exec.Command(command)
		cmd.Stdout = os.Stdout
		cmd.Run()

		return
	}

	// what the fuck is the current OS ?
	// Better have a fallback in case :)
	fmt.Println("\n\n\n\n")
}
