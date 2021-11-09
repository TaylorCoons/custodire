package dbstartup

import "database/sql"

func InitializeTables(db *sql.DB) error {
	var err error
	err = InitializeUsers(db)
	if err != nil {
		return err
	}
	err = InitializeDevices(db)
	if err != nil {
		return err
	}
	err = InitializeAccounts(db)
	if err != nil {
		return err
	}
	return nil
}

// TODO: Set max lengths in validation file
func InitializeUsers(db *sql.DB) error {
	_, err := db.Query(
		`
		CREATE TABLE IF NOT EXISTS users (
			user_id INT AUTO_INCREMENT PRIMARY KEY,
			username VARCHAR(255) NOT NULL,
			password_hash VARCHAR(512) NOT NULL 
		);
		`,
	)
	return err
}

func InitializeDevices(db *sql.DB) error {
	_, err := db.Query(
		`
		CREATE TABLE IF NOT EXISTS devices (
			device_id INT AUTO_INCREMENT PRIMARY KEY,
			user_id INT NOT NULL,
			pub_key TEXT NOT NULL,
			FOREIGN KEY (user_id)
				REFERENCES users (user_id)
				ON UPDATE CASCADE ON DELETE CASCADE
		);
		`,
	)
	return err
}

func InitializeAccounts(db *sql.DB) error {
	_, err := db.Query(
		`
		CREATE TABLE IF NOT EXISTS accounts (
			account_id INT AUTO_INCREMENT PRIMARY KEY,
			user_id INT NOT NULL,
			name VARCHAR(255) NOT NULL,
			email VARCHAR(512),
			created DATE,
			updated DATE,
			question1 VARCHAR(512),
			question2 VARCHAR(512),
			question3 VARCHAR(512),
			answer1 VARCHAR(512),
			answer2 VARCHAR(512),
			answer3 VARCHAR(512),
			description TEXT,
			password_enc TEXT,
			pin_enc TEXT,
			FOREIGN KEY (user_id)
				REFERENCES users (user_id)
				ON UPDATE CASCADE ON DELETE CASCADE
		);
		`,
	)
	return err
}
