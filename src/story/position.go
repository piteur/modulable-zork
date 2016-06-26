package story

import (
	"fmt"
	"github.com/piteur/modulable-zork/src/util"
	"strings"
)

// each story position
type StoryPosition struct {
	Name         string
	Description  string
	Actions      map[string]StoryAction
}

// display & interact with a story position
func (position StoryPosition) Play() {
	fmt.Println(position.Description)
	fmt.Println("===========================")

	for {
		action := waitForCorrectInput(position)

		fmt.Println(action.Text)

		// do we have a condition ?
		if action.HasCondition() {
			action.Condition.Init()

			if action.Condition.IsTestable() {
				if action.Condition.Verify() {
					fmt.Println(action.Valid.Text)
				} else {
					fmt.Println(action.Invalid.Text)
				}
			}

		}
	}
}

// wait for the user to input something correct, and return the action to be played next
func waitForCorrectInput(position StoryPosition) (action StoryAction) {
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
