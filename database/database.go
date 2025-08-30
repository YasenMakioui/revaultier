package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// type Database struct {
// 	db *sql.DB
// }

func NewDatabase(path string) *sql.DB {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
	}
	// InitDB

	sqlStmt := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        name TEXT,
		password varchar(128)
    );
    `
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Table 'users' created successfully")

	return db
}
