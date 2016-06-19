# Modulable Zork
----------------

This project aim to furnish a fully modulable Zork-like game.

For those who don't know what this game is, please take a look here: https://en.wikipedia.org/wiki/Zork

## How does it works ?

The compiled binary will look for a `./stories` folder with json files inside.
Each json file should contain a properly formatted story.

The story format can be seen right there: [doc/storyDocumentation.json](./doc/storyDocumentation.json)

## How to build the binary file

All the code/binary interactions are done throught the `make` command.

```console
$ make
build-linux                    Compile Go binary for Linux
build-windows                  Compile Go binary for Windows
build                          Compile Go binary for the current OS
help                           Display usage
```
The compiled binary will be stored on the `bin/` folder.

You can run it directly and try the demo story, or put it where you want to load your stories.
