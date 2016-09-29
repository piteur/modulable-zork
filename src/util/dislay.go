package util

import (
	"fmt"
	"os"
)

func DisplayMessage(errMessage string, exitCode ...int) {
	fmt.Println(errMessage)
	inputWait()

	if (len(exitCode) > 0) {
		os.Exit(exitCode[0])
	}
}

// display a nice formatted message & exit the program
func StopOnError(err error) {
	errorMessage := `
An error occured:
		%s

----
If you believe this error shouldn't have happen
Please feel free to create a new issue on this page: https://github.com/piteur/modular-zork/issues

Sorry for the inconvenience.
`

	fmt.Printf(errorMessage, err.Error())

	inputWait()
	os.Exit(2)
}

func inputWait() {
	// wait before killing the program
	// useful for users who double-click the binary:
	// basically, all non-dev :)
	fmt.Println("")
	fmt.Println("The game will stop at the next [enter] keyboard input")
	ReadString()
}
