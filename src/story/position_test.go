package story

import(
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsLoaded(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(StoryPosition{}.IsLoaded(), false)
	assert.Equal(StoryPosition{Name: "test"}.IsLoaded(), false)
	assert.Equal(StoryPosition{Description: "test"}.IsLoaded(), false)
	assert.Equal(StoryPosition{Name: "test", Description: "test"}.IsLoaded(), true)
}

func TestIsFinal(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(StoryPosition{}.IsFinal(), false)
	assert.Equal(StoryPosition{Final: 1}.IsFinal(), true)
	assert.Equal(StoryPosition{Final: 0}.IsFinal(), false)
	assert.Equal(StoryPosition{Final: 2}.IsFinal(), false)
}

func TestGetAction(t *testing.T) {
	assert := assert.New(t)

	actions := make(map[string]StoryAction)
	actions["Test"] = StoryAction{Text: "test1"}
	actions["test"] = StoryAction{Text: "test2"}
	actions["test 2"] = StoryAction{Text: "test3"}

	storyPosition := StoryPosition{
		Actions: actions,
	}

	// exact match
	assert.Equal(storyPosition.getAction("Test"), StoryAction{Text: "test1"})
	assert.Equal(storyPosition.getAction("test"), StoryAction{Text: "test2"})
	assert.Equal(storyPosition.getAction("test 2"), StoryAction{Text: "test3"})

	// no match
	assert.Equal(storyPosition.getAction("test2"), StoryAction{})
	assert.Equal(storyPosition.getAction("nothing"), StoryAction{})

	// testing space
	assert.Equal(storyPosition.getAction("Test  "), StoryAction{Text: "test1"})
	assert.Equal(storyPosition.getAction("test 2  "), StoryAction{Text: "test3"})
	assert.Equal(storyPosition.getAction("  test 2"), StoryAction{Text: "test3"})
	assert.Equal(storyPosition.getAction("  test 2  "), StoryAction{Text: "test3"})

	// testing case
	assert.Equal(storyPosition.getAction("Test 2"), StoryAction{Text: "test3"})
	assert.Equal(storyPosition.getAction("TEST 2"), StoryAction{Text: "test3"})
	assert.Equal(storyPosition.getAction("teSt"), StoryAction{Text: "test2"})
	assert.Equal(storyPosition.getAction("TEST"), StoryAction{Text: "test2"})
}
