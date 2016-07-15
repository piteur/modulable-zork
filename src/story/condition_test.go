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

func TestIsInitialized(t *testing.T) {
	assert := assert.New(t)

	storyCondition := StoryCondition{Key: "test"}

	assert.Equal(storyCondition.isInitialized(), false)
	storyCondition.Init()
	assert.Equal(storyCondition.isInitialized(), true)
}

func TestIsTestable(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(StoryCondition{Test: "="}.IsTestable(), true)
	assert.Equal(StoryCondition{}.IsTestable(), false)
}

func TestVerify(t *testing.T) {
	assert := assert.New(t)

	StoryCondition{Key: "test", Value: "42"}.Init()

	// valid test
	assert.Equal(
		StoryCondition{Key: "test", Value: "42", Test: "="}.Verify(),
		true)
	assert.Equal(
		StoryCondition{Key: "test", Value: "42", Test: "=="}.Verify(),
		true)
	assert.Equal(
		StoryCondition{Key: "test", Value: "42", Test: ">="}.Verify(),
		true)
	assert.Equal(
		StoryCondition{Key: "test", Value: "42", Test: "<="}.Verify(),
		true)

	// invalid test
	assert.Equal(
		StoryCondition{Key: "test", Value: "42", Test: ">"}.Verify(),
		false)
	assert.Equal(
		StoryCondition{Key: "test", Value: "42", Test: "<"}.Verify(),
		false)
	assert.Equal(
		StoryCondition{Key: "test", Value: "42", Test: "<>"}.Verify(),
		false)
	assert.Equal(
		StoryCondition{Key: "test", Value: "42", Test: "!="}.Verify(),
		false)
	assert.Equal(
		StoryCondition{Key: "test", Value: "42", Test: "!=="}.Verify(),
		false)

	assert.Equal(
		StoryCondition{Key: "test", Value: "42", Test: "brainFuckComparaison"}.Verify(),
		false)
}
