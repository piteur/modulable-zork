# Modular Zork [![Build Status](https://travis-ci.org/piteur/modular-zork.svg?branch=master)](https://travis-ci.org/piteur/modular-zork)

This project aim to furnish a fully modular Zork-like game.

For those who don't know what this game is, please take a look here: https://en.wikipedia.org/wiki/Zork

## How do I play it ?

First, you'll need to download the game.
You can list all the releases here: https://github.com/piteur/modular-zork/releases

Click on the link to download the correct version for your system.

Then, simply double-click on the executable file to launch the game.
The stories you want to play should be located on the `./stories` folder.

You can launch a specific story by using the `story` flag like this:

```console
$ ./bin/modular-zork --story ./stories/storyDocumentation.json
```

### How to create my own story ?

You can create your own story by creating a `.json` file, and by respecting the following formatting: [stories/storyDocumentation.json](./stories/storyDocumentation.json)

You can validate the json formatting on the [jsonlint.com](http://jsonlint.com) website.

## Technical details

### How to build the binary file

All the code/binary interactions are done through the `make` command.
    
```console
$ make
build-all                      install gox & build all the possible binaries
build                          compile Go binary for the current OS
generate-download-page         generate the download page to list available binaries
help                           display usage
install                        install dependencies with glide
run                            compile & run the binary
test                           run go tests
```

The compiled binary will be stored on the `bin/` folder.
