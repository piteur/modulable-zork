package story

// condition used on a story action
type StoryCondition struct {
	Key   string
	Value string
	Test  string
}

// static storage of all the conditions
var staticConditions map[string]StoryCondition

// read & store if necessary a condition
func (storyCondition StoryCondition) Init() {
	// init the map if not already done
	if len(staticConditions) == 0 {
		staticConditions = make(map[string]StoryCondition)
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
