package users

import (
	"database/sql"
	"errors"
)

// TODO: Move these to a models folder
type UserModelInput struct {
	Username string `json:"username"`
	Hash     string `json:"hash"`
}

type UserExistsError struct {
	UserExists string
}

func CreateUser(db *sql.DB, user UserModelInput) error {
	// TODO: Check user exists

	// TODO: Create user
	return errors.New("not implimented")
}

func DeleteUser(db *sql.DB, user UserModelInput) error {
	// TODO: Delete user
	return errors.New("not implimented")
}
