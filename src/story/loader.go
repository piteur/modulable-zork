package story

import (
	"fmt"
	"io/ioutil"
	"os"
)

type story struct {
	Name  string
	Index int
}

type storyAction struct {
	Text string
	moveTo int
}

type loadedStory struct {
	InternalName string
	InternalCode int
	Name         string
	Description	string
	Actions map[string]storyAction
}

var storyMap = map[int]story{}

/*
	Scan the 'stories' folder and load stories name & index to story.storyMap
*/
func LoadStories() {
	dirs, err := ioutil.ReadDir("./stories")

	if err != nil {
		fmt.Println("No 'stories' folder found. Aborting")
		return
	}

	for index, dir := range dirs {
		if dir.IsDir() {
			storyMap[index] = story{
				Index: index,
				Name:  dir.Name(),
			}
		}
	}

	return
}

/*
	Prompt user to choose an history
	Will prompt until a valid int has been typed. Exit on invalid type.
*/
func ChooseHistory() story {
	var choice int

	fmt.Println("Choose the story to load:")

	for _, story := range storyMap {
		fmt.Printf("\t%v - %s", story.Index, story.Name)
		fmt.Println("")
	}

	for {
		fmt.Println("")
		fmt.Println("What story do you want to play ? (enter valid number)")

		_, err := fmt.Scanf("%v", &choice)

		if err != nil {
			fmt.Println("invalid input ! Please enter a number")
			os.Exit(0)
		}

		if _, exist := storyMap[choice]; exist {
			break
		}
	}

	return storyMap[choice]
}

func LoadStory(story) (loadedStory loadedStory) {

}
