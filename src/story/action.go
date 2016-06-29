package story

import (
	"fmt"
)

// each action detail
type StoryAction struct {
	Text           string
	MoveToPosition string
	Condition      StoryCondition

	Valid          *StoryAction
	Invalid        *StoryAction
}

// Run the action: set/test condition, move to another position or whatever
func (storyAction StoryAction) Run() string {
	fmt.Println(storyAction.Text)

	// do we have a condition ?
	if storyAction.hasCondition() {
		storyAction.Condition.Init()

		if storyAction.Condition.IsTestable() {
			if storyAction.Condition.Verify() {
				fmt.Println()
				return storyAction.Valid.Run()
			} else {
				return storyAction.Invalid.Run()
			}
		}
	}

	// do we need to move to another position ?
	if storyAction.hasMoveToPosition() {
		return storyAction.MoveToPosition
	}

	return ""
}

// does the action have a condition
func (storyAction StoryAction) hasCondition() bool {
	if (storyAction.Condition != StoryCondition{}) {
		return true
	}

	return false
}

// does the action have a condition
func (storyAction StoryAction) hasMoveToPosition() bool {
	if (storyAction.MoveToPosition != "") {
		return true
	}

	return false
}
