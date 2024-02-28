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

// Searches for all posts a from a user and from the users followed by him
func (repository Posts) Search(userId uint64) ([]models.Post, error) {
	rows, err := repository.db.Query(`
		SELECT DISTINCT p.*, u.nick FROM posts p
		INNER JOIN users u on u.id = p.author_id
		INNER JOIN followers s on p.author_id = s.user_id
		WHERE u.id = ? or s.follower_id = ?`,
		userId, userId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post

		if err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil

}

// Update the info about a post on database
func (repository Posts) Update(postId uint64, post models.Post) error {
	statement, err := repository.db.Prepare(
		"UPDATE posts SET title = ?, content = ? WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(post.Title, post.Content, post.ID); err != nil {
		return err
	}

	return nil
}

// Deletes a post from database
func (repository Posts) Delete(postId uint64) error {
	statement, err := repository.db.Prepare("DELETE FROM posts WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postId); err != nil {
		return err
	}

	return nil
}

func (repository Posts) SearchByUser(userId uint64) ([]models.Post, error) {
	rows, err := repository.db.Query(`
		SELECT p.*, u.nick FROM posts p
		JOIN users u on u.id = p.author_id
		WHERE p.author_id = ?`,
		userId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post

		if err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}
