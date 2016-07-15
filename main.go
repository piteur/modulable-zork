package main

import (
	"fmt"
	"github.com/piteur/modular-zork/src/game"
	"github.com/piteur/modular-zork/src/loader"
	"os"
	"github.com/piteur/modular-zork/src/util"
)

func main() {
	fmt.Println("Welcome to 'modular Zork'")

	loader.LoadStories()
	story := loader.ChooseHistory()

	// launch the loaded story
	exitCode := game.Run(story)

	// wait before killing the program
	// useful for users who double-click the binary
	fmt.Println("")
	_, _ = util.ReadString()

	os.Exit(exitCode)
}
