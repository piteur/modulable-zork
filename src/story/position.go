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
func (position StoryPosition) Run() string {
	fmt.Println(position.Description)
	fmt.Println("===========================")

	for {
		action := waitForCorrectInput(position)

		positionId := action.Run()

		if positionId != "" {
			return positionId
		}
	}
}

// is a position loaded ? Or empty
func (position StoryPosition) IsLoaded() bool {
	return position.Name != "" &&
		position.Description != ""
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

