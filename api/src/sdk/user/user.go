package user

import (
	"database/sql"
	"fmt"
)

// TODO: Move these to a models folder
type UserModelInput struct {
	Username string `json:"username"`
	Hash     string `json:"hash"`
}

// TODO: Move these errors to their own file
type UserExistsError struct {
	User string
}

func (uee *UserExistsError) Error() string {
	return fmt.Sprintf("user: %s already exists", uee.User)
}

type UserDoesNotExistError struct {
	User string
}

func (udne *UserDoesNotExistError) Error() string {
	return fmt.Sprintf("user: %s does not exist", udne.User)
}

func CreateUser(db *sql.DB, user UserModelInput) error {
	results, err := db.Query(
		`
		SELECT user_id FROM users WHERE username=?
		`,
		user.Username,
	)
	if err != nil {
		return err
	}
	defer results.Close()
	if results.Next() {
		return &UserExistsError{user.Username}
	}
	_, err = db.Query(
		`
			INSERT INTO users (username, password_hash) VALUES (?, ?);
		`,
		user.Username,
		user.Hash,
	)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(db *sql.DB, username string) error {
	results, err := db.Query(
		`
		SELECT user_id FROM users WHERE username=?
		`,
		username,
	)
	if err != nil {
		return err
	}
	defer results.Close()
	if !results.Next() {
		return &UserDoesNotExistError{username}
	}
	_, err = db.Query(
		`
			DELETE FROM users WHERE username=?
		`,
		username,
	)
	if err != nil {
		return err
	}
	return nil
}
