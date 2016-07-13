package loader

import (
	"encoding/json"
	"fmt"
	"github.com/piteur/modulable-zork/src/story"
	"github.com/piteur/modulable-zork/src/util"
	"io/ioutil"
	"os"
)

var storyMap = map[int]story.Story{}

// Scan the 'stories' folder and load stories name & index to story.storyMap
func LoadStories() {
	files, err := ioutil.ReadDir("./stories")

	if err != nil {
		fmt.Println("No 'stories' folder found. Aborting")
		return
	}

	for index, file := range files {
		var fileContent []byte

		if !file.IsDir() {
			fileContent, err = ioutil.ReadFile("./stories/" + file.Name())

			if err != nil {
				fmt.Println("Unable to load file content")
				os.Exit(2)
			}

			story := load(fileContent)
			story.StoryId = index

			storyMap[index] = story
		}
	}

	return
}

// Prompt user to choose an history
// Will prompt until a valid int has been typed. Exit on invalid type.
func ChooseHistory() story.Story {
	fmt.Println("Choose the story to load:\n")

	for _, story := range storyMap {
		fmt.Printf("%v - %s\n", story.StoryId, story.Name)
		fmt.Printf("\tDescription:\n\t%s\n", story.Description)
		fmt.Println("")
	}

	return storyMap[getValidStoryInput()]
}

// Get a valid story input from the user
func getValidStoryInput() (choice int) {
	for {
		fmt.Println("")
		fmt.Println("What story do you want to play ?")

		choice, err := util.ReadInt()

		if err != nil {
			fmt.Println("invalid input ! Please enter a valid number")
			continue
		}

		if _, exist := storyMap[choice]; exist {
			return choice
		}
	}
}

// load a story from a json string
func load(storyToLoad []byte) (story story.Story) {
	json.Unmarshal(storyToLoad, &story)

	return
}
