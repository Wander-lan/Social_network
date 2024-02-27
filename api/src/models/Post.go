package models

import (
	"errors"
	"strings"
	"time"
)

// Represents a post made by a user
type Post struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorID   uint64    `json:"authorId,omitempty"`
	AuthorNick uint64    `json:"authorNick,omitempty"`
	Likes      uint64    `json:"likes"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
}

// Validates and formats a new post
func (post *Post) Prepare() error {
	if err := post.validate(); err != nil {
		return err
	}

	post.format()

	return nil
}

func (post *Post) validate() error {
	if post.Title == "" {
		return errors.New("Title is mandatory and cannot be empty")
	}

	if post.Content == "" {
		return errors.New("Content is mandatory and cannot be empty")
	}

	return nil
}

func (post *Post) format() {
	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)
}
