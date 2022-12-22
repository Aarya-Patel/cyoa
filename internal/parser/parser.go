package parser

import (
	"encoding/json"
	"os"

	"github.com/Aarya-Patel/cyoa/model"
)

func ParseStory(filename string) (model.Story, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return model.Story{}, err
	}

	var story model.Story
	err = json.Unmarshal(data, &story.Arcs)
	if err != nil {
		return model.Story{}, err
	}

	return story, nil
}
