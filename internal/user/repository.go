package user

import (
	"database/sql"
)

type UserRepository struct {
	dbconn *sql.DB
}

func NewRepository(db *sql.DB) *UserRepository {
	return &UserRepository{dbconn: db}
}

func (r *UserRepository) UserExists(username string) bool {
	// Returns true if the user exists
	sqlStmt := "SELECT name FROM users WHERE name=?"

	row := r.dbconn.QueryRow(sqlStmt, username)

	var val string

	if err := row.Scan(&val); err != nil {
		return false
	}

	return true
}

func (r *UserRepository) InsertUser(email string, password string) error {
	sqlStmt := "INSERT INTO users(name, password) VALUES(?,?)"

	_, err := r.dbconn.Exec(sqlStmt, email, password)

	if err != nil {
		return err
	}

	return nil
}
