package game

import(
	"github.com/piteur/modulable-zork/src/story"
	"fmt"
	"github.com/piteur/modulable-zork/src/util"
)

// launch a story to play it
func Run(loadedStory story.Story) {
	// clear console to start on a clean screen
	util.ClearConsole()

	fmt.Println("You'll now play this story: " + loadedStory.Name)
	fmt.Println("\n")

	position := loadedStory.LoadDefaultPosition()

	for {
		if (!position.IsLoaded()) {
			fmt.Println("No more position to load")
			// no more position to load, stopping the game
			break
		}

		positionId := position.Run()

		position = loadedStory.LoadPosition(positionId)
	}
}
