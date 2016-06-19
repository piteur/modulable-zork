package story

// general story content
type Story struct {
	StoryId         int
	Name            string
	Description     string
	Positions       map[string]StoryPosition
	DefaultPosition string
}

// each story position
type StoryPosition struct {
	Name         string
	Description  string
	Actions      map[string]StoryAction
}

// each action detail
type StoryAction struct {
	Text           string
	MoveToPosition int
	Condition      StoryCondition

	Valid          *StoryAction
	Invalid        *StoryAction
}

// condition (set or test) used on a story action
type StoryCondition struct {
	Key   string
	Value string
	Test  string
}
