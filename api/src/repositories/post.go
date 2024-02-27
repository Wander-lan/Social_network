package repositories

import (
	"api/src/models"
	"database/sql"
)

// Represents a post repository
type Posts struct {
	db *sql.DB
}

// Create a post repository
func NewPostRepository(db *sql.DB) *Posts {
	return &Posts{db}
}

// Inserts a post on database
func (repository Posts) Create(post models.Post) (uint64, error) {
	statement, err := repository.db.Prepare(
		"INSERT INTO posts (title, content, author_id) values(?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(post.Title, post.Content, post.AuthorID)
	if err != nil {
		return 0, err
	}

	lastIdInserted, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastIdInserted), nil
}
