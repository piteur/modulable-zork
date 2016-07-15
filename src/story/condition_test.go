package story

import(
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInit(t *testing.T) {
	assert := assert.New(t)

	// map not initialized
	assert.Equal(len(staticConditions), 0)

	// initialization of the map
	storyCondition := StoryCondition{
		Key: "test1",
	}
	storyCondition.Init()

	assert.Equal(len(staticConditions), 1)

	// adding another condition
	AnotherStoryCondition := StoryCondition{
		Key: "test2",
	}
	AnotherStoryCondition.Init()

	assert.Equal(len(staticConditions), 2)

	// doing it another time do nothing: key already used
	AnotherStoryCondition.Init()

	assert.Equal(len(staticConditions), 2)
}
