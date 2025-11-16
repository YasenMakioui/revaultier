package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func NewDatabase(path string) *sql.DB {
	db, err := sql.Open("sqlite3", path)
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
