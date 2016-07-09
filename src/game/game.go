package game

import (
	"fmt"
	"github.com/piteur/modulable-zork/src/story"
	"github.com/piteur/modulable-zork/src/util"
)

// launch a story to play it
func Run(loadedStory story.Story) int {
	// clear console to start on a clean screen
	util.ClearConsole()

	fmt.Println("You'll now play this story: " + loadedStory.Name)
	fmt.Println("\n")

	position := loadedStory.LoadDefaultPosition()

	for {
		// the player has arrived to the last position, stopping the game
		if position.IsFinal() {
			position.RunFinalPosition()

			return 0
		}

		// an error occurred on the next position load, stopping
		if !position.IsLoaded() {
			fmt.Println("No more position to load, the story seems missconfigured")

			return 1
		}

		// waiting for the next position to be played
		nextPosition := position.Run()

		position = loadedStory.LoadPosition(nextPosition)
	}
}
