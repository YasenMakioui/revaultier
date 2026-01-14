package auth

import (
	"context"
	"database/sql"
)

type AuthRepository struct {
	dbconn *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{dbconn: db}
}

func (r *AuthRepository) UserExists(ctx context.Context, username string) (bool, error) {
	sqlStmt := "SELECT 1 FROM user WHERE username=?"

	var exists int

	err := r.dbconn.QueryRowContext(ctx, sqlStmt, username).Scan(&exists)

	if err == sql.ErrNoRows {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *AuthRepository) GetUser(ctx context.Context, username string) (User, error) {

	sqlStmt := "SELECT * FROM user WHERE username=?"

	var user User

	err := r.dbconn.QueryRowContext(ctx, sqlStmt, username).Scan(&user.Id, &user.Username, &user.Password)

	if err != nil {
		return User{}, err
	}

	return user, err
}

func (r *AuthRepository) GetUUID(ctx context.Context, username string) (string, error) {

	sqlStmt := "SELECT id from user WHERE username=?"

	var uuid string

	err := r.dbconn.QueryRowContext(ctx, sqlStmt, username).Scan(&uuid)

	if err != nil {
		return "", err
	}

	return uuid, err
}

func (r *AuthRepository) InserUser(ctx context.Context, userId string, username string, passwordHash string) error {
	sqlStmt := "INSERT INTO user VALUES(?,?,?)"

	// generate uuid
	_, err := r.dbconn.ExecContext(ctx, sqlStmt, userId, username, passwordHash)

	if err != nil {
		return err
	}

	return nil
}
