package story

// general story content
type story struct {
	StoryId         int
	Name            string
	Description     string
	StoryPositions  map[int]storyPosition
	DefaultPosition int
}

// each story position
type storyPosition struct {
	InternalName string
	InternalCode int
	Name         string
	Description  []byte
	Actions      map[string]storyAction
}

// each action detail
type storyAction struct {
	Text   string
	MoveTo int

	SetGlobalCondition string
	SetCondition       string

	TestGlobalCondition string
	TestCondition       string

	valid   *storyAction
	invalid *storyAction
}
