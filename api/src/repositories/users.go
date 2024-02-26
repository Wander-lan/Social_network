package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

// Users represents a users repository
type Users struct {
	db *sql.DB
}

// Create a user repository
func NewUserRepository(db *sql.DB) *Users {
	return &Users{db}
}

// Inserts a user on the database
func (repository Users) Create(user models.User) (uint64, error) {
	statement, err := repository.db.Prepare(
		"INSERT INTO users (name, nick, email, password) values(?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastIdInserted, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastIdInserted), nil
}

// Finds all users with same name or nick
func (repository Users) Search(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) // %nameOrNick%

	rows, err := repository.db.Query(
		"SELECT id, name, nick, email, createdAt FROM users where name LIKE ? OR nick LIKE ?",
		nameOrNick, nameOrNick,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// Finds a user by ID
func (repository Users) SearchById(id uint64) (models.User, error) {
	rows, err := repository.db.Query(
		"SELECT id, name, nick, email, createdAt FROM users where id = ?",
		id,
	)
	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	var user models.User

	if rows.Next() {
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

// Updates a user data
func (repository Users) Update(id uint64, user models.User) error {
	statement, err := repository.db.Prepare(
		"UPDATE users SET name = ?, nick = ?, email = ?, WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Nick, user.Email, user.ID); err != nil {
		return err
	}

	return nil
}

// Deletes a user
func (repository Users) Delete(id uint64) error {
	statement, err := repository.db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(id); err != nil {
		return err
	}

	return nil
}

// Searches a user by email
func (repository Users) SearchByEmail(email string) (models.User, error) {
	rows, err := repository.db.Query("SELECT id, password FROM users where email = ?", email)
	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	var user models.User

	if rows.Next() {
		if err = rows.Scan(
			&user.ID,
			&user.Password,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

// Makes a user follow another
func (repository Users) Follow(userId, followerId uint64) error {
	statement, err := repository.db.Prepare(
		"INSERT IGNORE INTO followers (user_id, follower_id) values (?, ?)",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userId, followerId); err != nil {
		return err
	}

	return nil
}

// Makes a user unfollow another
func (repository Users) Unfollow(userId, followerId uint64) error {
	statement, err := repository.db.Prepare(
		"DELETE FROM followers WHERE user_id = ? AND follower_id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userId, followerId); err != nil {
		return err
	}

	return nil
}

// Finds all followers of a user
func (repository Users) SearchFollowers(userId uint64) ([]models.User, error) {
	rows, err := repository.db.Query(`
		SELECT u.id, u.name, u.nick, u.email, u.createdAt
		FROM users u INNER JOIN followers s
		ON u.id = s.follower_id WHERE s.user_id = ?
		`, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User

		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// Find who is being followed by a user
func (repository Users) SearchFollowing(userId uint64) ([]models.User, error) {
	rows, err := repository.db.Query(`
		SELECT u.id, u.name, u.nick, u.email, u.createdAt
		FROM users u INNER JOIN followers s
		ON u.id = s.user_id WHERE s.follower_id = ?
		`, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// finds the password of a user by ID
func (repository Users) SearchPassword(userId uint64) (string, error) {
	rows, err := repository.db.Query("SELECT password FROM users WHERE id = ?", userId)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var user models.User

	if rows.Next() {
		if err = rows.Scan(&user.Password); err != nil {
			return "", err
		}
	}

	return user.Password, nil
}

func (repository Users) UpdatePassword(userId uint64, password string) error {
	statement, err := repository.db.Prepare("UPDATE users SET password = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(password, userId); err != nil {
		return err
	}

	return nil
}
