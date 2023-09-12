package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func main() {
	filename := flag.String("file", "gopher.json", "the json file with the goword story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	d := json.NewDecoder(f)
	var story gowordstory.Story
	err = d.Decode(&story)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story)
}
