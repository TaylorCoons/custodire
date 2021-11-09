package user

import (
	"database/sql"
	"errors"
	"fmt"
)

// TODO: Move these to a models folder
type UserModelInput struct {
	Username string `json:"username"`
	Hash     string `json:"hash"`
}

type UserExistsError struct {
	User string
}

func (uee *UserExistsError) Error() string {
	return fmt.Sprintf("user: %s already exists. ", uee.User)
}

func CreateUser(db *sql.DB, user UserModelInput) error {
	// TODO: Check user exists
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
	fmt.Println("This called 2")
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

func DeleteUser(db *sql.DB, user UserModelInput) error {
	// TODO: Delete user
	return errors.New("not implimented")
}
