package models

import (
	"encoding/json"
	"io"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// ConvertJSONToPosts converts the JSON data into an array of Post objects.
func ConvertJSONToPosts(data io.Reader, posts *[]Post) error {
	decoder := json.NewDecoder(data)
	err := decoder.Decode(posts)
	if err != nil {
		return err
	}
	return nil
}

func ConvertJSONToPost(data io.Reader, posts *Post) error {
	decoder := json.NewDecoder(data)
	err := decoder.Decode(posts)
	if err != nil {
		return err
	}
	return nil
}
