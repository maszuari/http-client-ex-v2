package model

import (
	"encoding/json"
	"time"
)

type Article struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	URL         string    `json:"url"`
	Imageurl    string    `json:"imageUrl"`
	Newssite    string    `json:"newsSite"`
	Summary     string    `json:"summary"`
	Publishedat time.Time `json:"publishedAt"`
	Updatedat   time.Time `json:"updatedAt"`
	Featured    bool      `json:"featured"`
	Launches    []struct {
		ID       string `json:"id"`
		Provider string `json:"provider"`
	} `json:"launches"`
	Events []struct {
		ID       string `json:"id"`
		Provider string `json:"provider"`
	} `json:"events"`
}

func UnmarshalJSON(data []byte) ([]Article, error) {

	var art []Article
	err := json.Unmarshal(data, &art)
	if err != nil {
		return art, err
	} else {
		return art, nil
	}
}
