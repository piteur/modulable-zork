package loader

import (
	"encoding/json"
	"fmt"
	"github.com/piteur/modular-zork/src/story"
	"github.com/piteur/modular-zork/src/util"
	"io/ioutil"
	"errors"
	"bytes"
)

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

			loadedStory, err := load(dir + "/" + file.Name())

			if err != nil {
				return storyMap, err
			}

			loadedStory.StoryId = index
			storyMap[index] = loadedStory
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
	fmt.Println("Choose the story to load:")
	fmt.Println()

	for _, loadedStory := range storyMap {
		fmt.Printf("%v - %s\n", loadedStory.StoryId, loadedStory.Name)
		fmt.Printf("\tDescription:\n\t%s\n", loadedStory.Description)
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
		err = fmt.Errorf("Unable to load content from file '%s'", path)

		return
	}

	err = json.Unmarshal(fileContent, &story)

	// we have a specific error on the json: display the offset of it !
	if syntaxError, ok := err.(*json.SyntaxError); ok {
		jsonErrorContent := make([]byte, 240)

		// place the reader pointer to the error offset
		reader := bytes.NewReader(fileContent)
		reader.Seek(syntaxError.Offset, 0)
		_, _ = reader.Read(jsonErrorContent)

		err = fmt.Errorf(`%s

The error is in the following lines:
----
	%s [...]
----

Check for any suspicous character, or syntax mistake.`,
			err.Error(),
			string(jsonErrorContent),
		)
	}

	if err != nil {
		errorMessage := `/!\/!\ The story '%s' can't be parsed due to an error on the story content /!\/!\
Error message:
	%s

You can check your story formating on this website to catch common mistake: http://jsonlint.com`

		err = fmt.Errorf(errorMessage, path, err.Error())

		return
	}

	return
}
