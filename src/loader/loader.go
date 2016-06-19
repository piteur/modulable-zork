package loader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"github.com/piteur/modulable-zork/src/story"
	"github.com/piteur/modulable-zork/src/util"
)

var storyMap = map[int]story.Story{}

// Scan the 'stories' folder and load stories name & index to story.storyMap
func LoadStories() {
	files, err := ioutil.ReadDir("./stories")

	if err != nil {
		fmt.Println("No 'stories' folder found. Aborting")
		return
	}

	var fileContent []byte

	for index, file := range files {
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
	var choice int

	fmt.Println("Choose the story to load:\n")

	for _, story := range storyMap {
		fmt.Printf("%v - %s\n", story.StoryId, story.Name)
		fmt.Printf("\tDescription:\n\t%s\n", story.Description)
		fmt.Println("")
	}

	for {
		fmt.Println("")
		fmt.Println("What story do you want to play ? (enter valid number)")

		choice, err := util.ReadInt()

		if err != nil {
			fmt.Println("invalid input ! Please enter a valid number")
			continue
		}

		if _, exist := storyMap[choice]; exist {
			break
		}
	}

	return storyMap[choice]
}

// load a story from a json string
func load(storyToLoad []byte) (story story.Story) {
	json.Unmarshal(storyToLoad, &story)

	return
}
