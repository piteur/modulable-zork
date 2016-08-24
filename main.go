package main

import (
	"fmt"
	"github.com/piteur/modular-zork/src/game"
	"github.com/piteur/modular-zork/src/loader"
	"os"
	"github.com/piteur/modular-zork/src/util"
	"flag"
	"github.com/piteur/modular-zork/src/story"
	"errors"
)

func main() {
	var defaultStory, storyToload string
	defaultStory = "no story selected"

	// flag configuration
	flag.StringVar(&storyToload, "story", defaultStory, "path to the story json file")
	flag.Parse()

	// load the story
	story := getStoryToLoad(storyToload, defaultStory)

	// and launch it
	exitCode := game.Run(story)

	// wait before killing the program
	// useful for users who double-click the binary
	fmt.Println("")
	util.ReadString()

	os.Exit(exitCode)
}

// get the story to load based on user input & configuration
func getStoryToLoad(storyToload string, defaultStory string) (story story.Story) {
	// no specific story to load: we list the folder & prompt the user
	if storyToload == defaultStory {
		fmt.Println("Welcome to 'modular Zork'")

		stories, err := loader.LoadStories("./stories")

		if err != nil {
			stopOnError(err)
		}

		return loader.ChooseHistory(stories)
	}

	// we try to load the specified story
	story, err := loader.LoadStory(storyToload)
	err = errors.New("testing stuff")

	if err != nil {
		stopOnError(err)
	}

	return
}

// display a nice formatted message & exit the program
func stopOnError(err error) {
	errorMessage := `
An error occured:
		%s

If you believe this error shouldn't have happen
Please feel free to create a new issue on this page: https://github.com/piteur/modular-zork/issues

Sorry for the inconvenience.
`

	fmt.Printf(errorMessage, err.Error())

	os.Exit(2)
}
