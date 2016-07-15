package story

import(
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasCondition(t *testing.T) {
	assert := assert.New(t)

	storyAction := StoryAction{
		Condition: StoryCondition{Key: "test"},
	}

	assert.Equal(storyAction.hasCondition(), true)
	assert.Equal(StoryAction{}.hasCondition(), false)
}

func TestHasMoveTo(t *testing.T) {
	assert := assert.New(t)

	storyAction := StoryAction{
		MoveTo: "something",
	}

	assert.Equal(StoryAction{}.hasMoveTo(), false)
	assert.Equal(storyAction.hasMoveTo(), true)
}
