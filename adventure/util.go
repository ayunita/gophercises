package util

import (
	"encoding/json"
	"io"
	"os"
)

func ParseJSON(file *os.File) (map[string]Chapter, error) {
	jsonBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var chapters map[string]Chapter
	err = json.Unmarshal(jsonBytes, &chapters)
	if err != nil {
		return nil, err
	}

	return chapters, nil
}

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
