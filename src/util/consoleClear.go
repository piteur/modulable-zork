package util

import (
	"os"
	"os/exec"
	"runtime"
	"fmt"
)

var clearFunc map[string]func()

func initMap() {
	if len(clearFunc) != 0 {
		return
	}

	clearFunc = make(map[string]func())

	// linux
	clearFunc["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	// mac
	clearFunc["darwin"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	// win
	clearFunc["windows"] = func() {
		cmd := exec.Command("cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func ClearConsole() {
	initMap()

	clear, exist := clearFunc[runtime.GOOS]

	if exist {
		clear()
	} else {
		fmt.Println("\n\n\n\n")
	}
}
