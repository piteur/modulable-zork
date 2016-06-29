package story

// general story content
type Story struct {
	StoryId         int
	Name            string
	Description     string
	Positions       map[string]StoryPosition
	DefaultPosition string
}

// load position from story
func (story Story) LoadPosition(positionId string) (position StoryPosition) {
	position, exist := story.Positions[positionId]

	if !exist {
		return StoryPosition{}
	}

	return
}

// load default position
func (story Story) LoadDefaultPosition() StoryPosition {
	return story.LoadPosition(story.DefaultPosition)
}
