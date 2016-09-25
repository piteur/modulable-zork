package loader

import (
	"encoding/json"
	"fmt"
	"github.com/piteur/modular-zork/src/story"
	"github.com/piteur/modular-zork/src/util"
	"io/ioutil"
	"errors"
)

var storyMap = map[int]story.Story{}

// Scan the 'stories' folder and load stories name & index to story.storyMap
func LoadStories(dir string) (storyMap map[int]story.Story, err error) {
	files, err := ioutil.ReadDir(dir)
	storyMap = make(map[int]story.Story)

	if err != nil {
		err = errors.New("No" + dir + " folder found. Aborting")

		return
	}

	for index, file := range files {
		if !file.IsDir() {
			if file.Name() == ".gitignore" {
				continue;
			}

			story, err := load(dir + "/" + file.Name())

			if err != nil {
				return storyMap, err
			}

			story.StoryId = index
			storyMap[index] = story
		}
	}

	return
}

// load a specific story
func LoadStory(file string) (story story.Story, err error) {
	story, err = load(file)

	return
}

// Prompt user to choose an history
// Will prompt until a valid int has been typed. Exit on invalid type.
func ChooseHistory(storyMap map[int]story.Story) story.Story {
	fmt.Println("Choose the story to load:\n")

	for _, story := range storyMap {
		fmt.Printf("%v - %s\n", story.StoryId, story.Name)
		fmt.Printf("\tDescription:\n\t%s\n", story.Description)
		fmt.Println("")
	}

	return storyMap[getValidStoryInput(storyMap)]
}

// Get a valid story input from the user
func getValidStoryInput(storyMap map[int]story.Story) int {
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

// load a story from a path
func load(path string) (story story.Story, err error) {
	fileContent, err := ioutil.ReadFile(path)

	if err != nil {
		err = errors.New("Unable to load content from file '" + path + "'")

		return
	}

	err = json.Unmarshal(fileContent, &story)

	if err != nil {
		errorMessage := `/!\/!\ The story '%s' can't be parsed due to an error on the story content /!\/!\
Error message:
	%s

You can check your story formating (json file) on this website to catch common mistake: jsonlint.com
If the issue persist, please open a bug case here: github.com/piteur/modular-zork/issues/new`

		err = errors.New(fmt.Sprintf(errorMessage, path, err.Error()))

		return
	}

	return
}
