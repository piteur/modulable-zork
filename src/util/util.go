package util

import (
	"os"
	"os/exec"
	"runtime"
	"fmt"
)

var clearFunc map[string]func()

func initMap() {
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

	value, exist := clearFunc[runtime.GOOS]

	if exist {
		value()
	} else {
		fmt.Println("\n\n\n\n")
	}
}
