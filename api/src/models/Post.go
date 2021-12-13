package models

import (
	"errors"
	"strings"
	"time"
)

// Post struct represents a publication
type Post struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorID   uint64    `json:"authorID,omitempty"`
	AuthorNick string    `json:"authorNick,omitempty"`
	Likes      uint64    `json:"likes"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
}

func (post *Post) Prepare() error {
	if error := post.validate(); error != nil {
		return error
	}

	post.format()
	return nil
}

func (post *Post) validate() error {

	if post.Title == "" {
		return errors.New("Title can't be empty")
	}

	if post.Content == "" {
		return errors.New("Title can't be empty")
	}

	return nil
}

func (post *Post) format() {
	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)
}
