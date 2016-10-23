package main

import (
	"fmt"
	"github.com/piteur/modular-zork/src/game"
	"github.com/piteur/modular-zork/src/loader"
	"github.com/piteur/modular-zork/src/util"
	"flag"
	"github.com/piteur/modular-zork/src/story"
)

func main() {
	var storyToLoad string

	// flag configuration
	flag.StringVar(&storyToLoad, "story", "", "path to the story json file")
	flag.Parse()

	// load the story
	story := getStoryToLoad(storyToLoad)

	// and launch it
	exitCode := game.Run(story)

	util.DisplayMessage("", exitCode)
}

// get the story to load based on user input & configuration
func getStoryToLoad(storyToload string) (story story.Story) {
	// no specific story to load: we list the folder & prompt the user
	if storyToload == "" {
		stories, err := loader.LoadStories("./stories")

		if err != nil {
			util.StopOnError(err)
		}

		fmt.Println("Welcome to 'modular Zork'")

		return loader.ChooseHistory(stories)
	}

	// we try to load the specified story
	story, err := loader.LoadStory(storyToload)

	if err != nil {
		util.StopOnError(err)
	}

	return
}
