# Contributing to Modular Zork

:tada: First, thanks for taking the time to contribute ! :tada:

## I want to modify the code

Feel free to fork the project & create a new pull request on it !
Don't forget to add the correct label to your PR, to have a better review.

Also, if you have modified the story schema don't forget to add a diff of your changes.

### How does the code works ?

Know a bit about the code logic can help !

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

## I found a bug !

When you think you have found a bug, please report it on this page: https://github.com/piteur/modular-zork/issues

Create a new issue to describe the bug you encountered, don't forget to add:
 - the played story
 - the input order you have used
 - the OS (windows, mac, linux ?)
 - the architecture (32bits, 64 bits, arm ?)
 - the version used (release number)

And don't hesitate to add any other useful information.

You can add the full story on a block like this one to be readable:
```
~~~json
{
    "Name": "Modular Zork: documentation",
    "Description": "This is the documentation & story example of the project",
    "defaultPosition": "0",
    "positions": {
        "0": {
            // ...
        }
    }
}
~~~
```

If the bug is reproduced it'll be resolved as soon as possible.

Don't forget that this project is maintained by only one developer, on his free time :)


Thanks! :heart:

Piteur
