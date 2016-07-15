package story

import(
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadPosition(t *testing.T) {
	assert := assert.New(t)

	positions := make(map[string]StoryPosition)
	positions["test"] = StoryPosition{Name: "test"}
	positions["test2"] = StoryPosition{Name: "test2"}

 	story := Story{
		DefaultPosition: "test",
		Positions: positions,
	}

	assert.Equal(story.LoadDefaultPosition(), StoryPosition{Name: "test"})
	assert.Equal(story.LoadPosition("test2"), StoryPosition{Name: "test2"})
	assert.Equal(story.LoadPosition("test3"), StoryPosition{})
}
