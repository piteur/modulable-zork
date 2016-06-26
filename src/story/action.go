package story

// each action detail
type StoryAction struct {
	Text           string
	MoveToPosition int
	Condition      StoryCondition

	Valid          *StoryAction
	Invalid        *StoryAction
}

// does the action have a condition
func (storyAction StoryAction) HasCondition() bool {
	if (storyAction.Condition != StoryCondition{}) {
		return true
	}

	return false
}
