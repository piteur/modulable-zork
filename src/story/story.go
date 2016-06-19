package story

// general story content
type Story struct {
	StoryId         int
	Name            string
	Description     string
	StoryPositions  map[int]StoryPosition
	DefaultPosition int
}

// each story position
type StoryPosition struct {
	InternalName string
	InternalCode int
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
