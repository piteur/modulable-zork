package story

import (
	"fmt"
	"github.com/piteur/modulable-zork/src/util"
	"strings"
)

// each story position
type StoryPosition struct {
	Name        string
	Description string
	Actions     map[string]StoryAction
	Final       int
}

// display & interact with a story position
func (position StoryPosition) Run() string {
	fmt.Println(position.Description)
	fmt.Println("===========================")

	for {
		action := waitForCorrectInput(position)

		positionId := action.Run()

		// only load a new position when the action tell us to do it
		if positionId != "" {
			return positionId
		}
	}
}

// run the final position
func (position StoryPosition) RunFinalPosition() {
	util.ClearConsole()

	fmt.Println(position.Description)
}

// is a position loaded ? Or empty
func (position StoryPosition) IsLoaded() bool {
	return position.Name != "" && position.Description != ""
}

// is a position final ? aka: last position
func (position StoryPosition) IsFinal() bool {
	return position.Final == 1
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
