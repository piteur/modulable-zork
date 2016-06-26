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
