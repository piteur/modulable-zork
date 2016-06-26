package story

// condition used on a story action
type StoryCondition struct {
	Key   string
	Value string
	Test  string
}

// static storage of all the conditions
var staticConditions map[string]StoryCondition
var staticInit bool

// read & store if necessary a condition
func (storyCondition StoryCondition) Init() {
	if !staticInit {
		staticConditions = make(map[string]StoryCondition)
		staticInit = true
	}

	if !storyCondition.IsTestable() {
		staticConditions[storyCondition.Key] = storyCondition
	}
}

// test a condition
func (storyCondition StoryCondition) Verify() bool {
	if !storyCondition.isInitialized() {
		return false
	}

	if !storyCondition.IsTestable() {
		return false
	}

	storedConditionValue := staticConditions[storyCondition.Key].Value

	switch storyCondition.Test {
		case "=", "==":
			return storedConditionValue == storyCondition.Value
		case ">":
			return storedConditionValue > storyCondition.Value
		case "<":
			return storedConditionValue < storyCondition.Value
		case "!=", "!==", "<>":
			return storedConditionValue != storyCondition.Value
	}

	return false
}

// is the condition as test or a declaration
func (storyCondition StoryCondition) IsTestable() bool {
	if storyCondition.Test == "" {
		return false
	}

	return true
}


// is the condition has been initialized ?
func (storyCondition StoryCondition) isInitialized() bool {
	if _, exist := staticConditions[storyCondition.Key]; !exist {
		return false
	}

	return true
}
