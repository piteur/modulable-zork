# Modular Zork [![Build Status](https://travis-ci.org/piteur/modular-zork.svg?branch=master)](https://travis-ci.org/piteur/modular-zork)

This project aim to furnish a fully modular Zork-like game.

For those who don't know what this game is, please take a look here: https://en.wikipedia.org/wiki/Zork

## How does it works ?

The compiled binary will look for a `./stories` folder with json files inside.
Each json file should contain a properly formatted story.

The story format can be seen right there: [stories/storyDocumentation.json](./stories/storyDocumentation.json)

## How to build the binary file

All the code/binary interactions are done through the `make` command.

```console
$ make
build-all                      install gox & build all the possible binaries
build                          compile Go binary for the current OS
help                           display usage
install                        install dependencies with glide
run                            compile & run the binary
test                           run go tests
```

The compiled binary will be stored on the `bin/` folder.


You can run it directly and try the demo story, or put it where you want to load your stories.

## A look inside the code logic

If you want to do any PR on this project, know a bit about the code logic can help !

We have the following struct used on the code:
 - Story
 - StoryPosition
 - StoryAction
 - StoryCondition

A *Story* represent the whole json file you've created: description, name and all the positions.

A *StoryPosition* represent the available possible action like "put my finger in my nose", and
also a description.

A *StoryAction* represent the action that are executed when you select a specific text from
the current position. It can set or test condition, display text or move the story to a new
position.

A *StoryCondition* can represent both a test or a set for a condition, like setting the value
of `myValue` to `42`. The condition will contain a `Valid` and an `Invalid` action, and they will be
executed depending of the condition test.
