package main

import (
	"fmt"
	"github.com/piteur/modulable-zork/src/loader"
	"github.com/piteur/modulable-zork/src/game"
)

func main() {
	fmt.Println("Welcome to 'modulable Zork'")

	loader.LoadStories()
	choice := loader.ChooseHistory()

	// launching the loaded story
	game.Play(choice)
}
