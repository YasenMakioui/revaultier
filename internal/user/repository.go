package user

import (
	"database/sql"
	"errors"
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

func (r *UserRepository) InsertUser(username string, password string) error {
	sqlStmt := "INSERT INTO users(name, password) VALUES(?,?)"

	_, err := r.dbconn.Exec(sqlStmt, username, password)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetUser(username string) (*User, error) {
	sqlStmt := "SELECT name,password FROM users WHERE name=?"

	row := r.dbconn.QueryRow(sqlStmt, username)

	user := &User{}

	if err := row.Scan(&user.username, &user.hashedPassword); err != nil {
		return user, errors.New("user not found")
	}

	return user, nil
}
