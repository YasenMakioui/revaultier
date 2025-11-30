package database

import (
	"database/sql"
	"log"
	"revaultier/configuration"

	_ "github.com/mattn/go-sqlite3"
)

func NewDatabase(cfg *configuration.Config) *sql.DB {
	db, err := sql.Open("sqlite3", cfg.Database.Database)
	if err != nil {
		log.Fatal(err)
	}
	// InitDB

	sqlStmt := `
    CREATE TABLE IF NOT EXISTS user (
		id TEXT PRIMARY KEY,
        username     TEXT   NOT NULL,
        password TEXT    NOT NULL
    );
    `

	_, err = db.Exec(sqlStmt)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Table 'user' created successfully")

	return db
}
