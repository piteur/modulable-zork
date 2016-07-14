package main

import (
	"fmt"
	"github.com/piteur/modular-zork/src/game"
	"github.com/piteur/modular-zork/src/loader"
	"os"
)

func main() {
	fmt.Println("Welcome to 'modular Zork'")

	loader.LoadStories()
	story := loader.ChooseHistory()

	// launch the loaded story & exit
	os.Exit(game.Run(story))
}
