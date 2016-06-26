package game

import(
	"github.com/piteur/modulable-zork/src/story"
	"fmt"
	"github.com/piteur/modulable-zork/src/util"
)

// launch a story to play it
func Play(story story.Story) {
	// clear console to start on a clean screen
	util.ClearConsole()

	fmt.Println("You'll now play this story: " + story.Name)
	fmt.Println("\n")

	position := story.Positions[story.DefaultPosition]

	position.Play()
}
