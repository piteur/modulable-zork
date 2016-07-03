package main

import (
	"fmt"
	"github.com/piteur/modulable-zork/src/game"
	"github.com/piteur/modulable-zork/src/loader"
	"os"
)

func main() {
	fmt.Println("Welcome to 'modulable Zork'")

	loader.LoadStories()
	story := loader.ChooseHistory()

	// launch the loaded story & exit
	os.Exit(game.Run(story))
}
