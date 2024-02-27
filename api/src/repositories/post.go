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

// Searches for a single post on database
func (repository Posts) SearchById(postId uint64) (models.Post, error) {
	rows, err := repository.db.Query(`
		SELECT p.*, u.nick
		FROM posts p INNER JOIN users u
		ON u.id = p.author_id WHERE p.id = ?`,
		postId,
	)
	if err != nil {
		return models.Post{}, err
	}
	defer rows.Close()

	var post models.Post

	for rows.Next() {
		if err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return models.Post{}, err
		}
	}

	return post, nil
}
