package go_word

import (
	"encoding/json"
	"io"
)

// return Story struct given Json file
func CreateStory(r io.Reader) (Story, error) {

	//better for reading from stream
	d := json.NewDecoder(r)

	//create var story of type Story
	var story Story
	//decode json value of struct Story and store it into var story

	if err := d.Decode(&story); err != nil {
		return nil, err
	}

	return story, nil

}

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json: "title"`
	Paragraphs []string `json: "paragraphs"`
	Options    []Option `json: "options"`
}

type Option struct {
	Text    string `"json: text`
	Chapter string `json: chapter`
}
