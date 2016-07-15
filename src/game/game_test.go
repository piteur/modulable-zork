package game

import(
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/piteur/modular-zork/src/story"
)

func TestFailingStory(t *testing.T) {
	assert := assert.New(t)

	positions := make(map[string]story.StoryPosition)
	positions["0"] = story.StoryPosition{}

	invalidStory := story.Story{
		StoryId: 42,
		Name: "don't care",
		Description: "don't care",
		Positions: positions,
		DefaultPosition: "0",
	}

	assert.Equal(Run(invalidStory), 1)
}

func TestValidStory(t *testing.T) {
	assert := assert.New(t)

	positions := make(map[string]story.StoryPosition)
	positions["0"] = story.StoryPosition{
		Final: 1,
	}

	invalidStory := story.Story{
		StoryId: 42,
		Name: "don't care",
		Description: "don't care",
		Positions: positions,
		DefaultPosition: "0",
	}

	assert.Equal(Run(invalidStory), 0)
}
