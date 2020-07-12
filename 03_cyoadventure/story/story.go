package story

import (
	"encoding/json"
	"os"
)

func Read(f *os.File) (Story, error) {
	d := json.NewDecoder(f)
	var story Story

	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

type Story map[string]Chapter

type Chapter struct {
	Title     string   `json:"title"`
	Paragraph []string `json:"story"`
	Options   []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}