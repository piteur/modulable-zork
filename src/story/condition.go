package story

// condition used on a story action
type StoryCondition struct {
	Key   string
	Value string
	Test  string
}

// static storage of all the conditions
var staticConditions map[string]StoryCondition

// read & store a condition if necessary
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
func (condition StoryCondition) Verify() bool {
	if !condition.isInitialized() {
		return false
	}

	if !condition.IsTestable() {
		return false
	}

	storedValue := staticConditions[condition.Key].Value

	switch condition.Test {
	case "=", "==":
		return storedValue == condition.Value
	case ">":
		return storedValue > condition.Value
	case ">=":
		return storedValue >= condition.Value
	case "<":
		return storedValue < condition.Value
	case "<=":
		return storedValue <= condition.Value
	case "!=", "!==", "<>":
		return storedValue != condition.Value
	}

	return false
}

// is the condition has been initialized ?
func (storyCondition StoryCondition) isInitialized() bool {
	if _, exist := staticConditions[storyCondition.Key]; !exist {
		return false
	}

	return true
}

// is the condition as test or a declaration ?
func (storyCondition StoryCondition) IsTestable() bool {
	if storyCondition.Test == "" {
		return false
	}

	return true
}
