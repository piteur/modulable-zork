package game

import(
	"github.com/piteur/modulable-zork/src/story"
	"fmt"
	"github.com/piteur/modulable-zork/src/util"
	"strings"
)

// launch a story to play it
func Play(story story.Story) {
	// clear console to start on a clean screen
	util.ClearConsole()

	fmt.Println("You'll now play this story: " + story.Name)
	fmt.Println("\n")

	playPosition(story.Positions[story.DefaultPosition])
}

// display & interact with a story position
func playPosition(position story.StoryPosition) {
	fmt.Println(position.Description)
	fmt.Println("===========================")

	action := waitForCorrectInput(position)

	fmt.Println(action.Text)
}

// wait for the user to input something correct
func waitForCorrectInput(position story.StoryPosition) (action story.StoryAction) {
	for {
		var exist bool

		fmt.Println("\n\tWhat do you want to do ?")
		choice, err := util.ReadString()

		if err != nil {
			fmt.Println("I don't understand what you asked for, sorry.")
			continue
		}

		// lower string, easier to compare
		choice = strings.ToLower(choice)

		action, exist = position.Actions[choice]

		if exist == false {
			fmt.Println("Nope, nothing to do with that. Sorry.")
			continue
		} else {
			break
		}
	}

	return
}
