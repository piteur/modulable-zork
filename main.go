package main

import (
	"github.com/piteur/modulable-zork/src/story"
	"fmt"
)

func main() {
	fmt.Println("Welcome to 'modulable Zork'")

	story.LoadStories()
	choice := story.ChooseHistory()

	fmt.Println(choice)
}
