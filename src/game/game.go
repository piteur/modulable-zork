package game

import(
	"github.com/piteur/modulable-zork/src/story"
	"fmt"
	"github.com/piteur/modulable-zork/src/util"
	"os"
)

// launch a story to play it
func Run(loadedStory story.Story) {
	// clear console to start on a clean screen
	util.ClearConsole()

	fmt.Println("You'll now play this story: " + loadedStory.Name)
	fmt.Println("\n")

	position := loadedStory.LoadDefaultPosition()

	for {
		// an error occured on the next position load, stopping
		if (!position.IsLoaded()) {
			fmt.Println("No more position to load")

			os.Exit(2)
		}

		// the player has arrived to the last position, stopping the game
		if (position.IsFinal()) {
			util.ClearConsole()

			fmt.Println(position.Description)
			fmt.Println()

			os.Exit(0)
		}

		// waiting for the next position to be played
		nextPosition := position.Run()

		position = loadedStory.LoadPosition(nextPosition)
	}
}
