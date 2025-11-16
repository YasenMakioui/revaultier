package user

import "database/sql"

type UserRepository struct {
	dbconn *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{dbconn: db}
}
