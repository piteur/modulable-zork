package story

var staticConditions map[string]StoryCondition
/*
func (storyCondition storyCondition) read() {
	if storyCondition.Test == "" {
		staticConditions[storyCondition.Key] = storyCondition
	}
}

func (storyCondition storyCondition) verify() bool {
	// if no condition exists, return false
	if _, exist := staticConditions[storyCondition.Key]; !exist {
		return false
	}

	key := staticConditions[storyCondition.Key]

	switch storyCondition.Test {
		case "=" && "==":
			return key.Value == storyCondition.Value
		case ">":
			return key > storyCondition.Value
		case "<":
			return key < storyCondition.Value
		case "!=" && "<>":
			return key != storyCondition.Value
	}

	return false
}
*/
