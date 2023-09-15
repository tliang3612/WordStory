# WordStory

A choose your own adventure game written in go.
Takes a story structured in a JSON file, parses it, decode it into a struct,
then writes it through http.ResponseWriter.

Story can also be viewed through the command line without sending http commands.
