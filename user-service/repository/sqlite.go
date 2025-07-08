package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// NewSQLiteDB - SQLite veritabanı bağlantısı oluşturur
func NewSQLiteDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// InitDatabase - Temel tabloları oluşturur
func InitDatabase(db *sql.DB) error {
	userTable := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		email TEXT UNIQUE,
		password TEXT,
		bio TEXT,
		seriousness INTEGER,
		height INTEGER,
		weight INTEGER,
		smokes BOOLEAN,
		drinks BOOLEAN,
		job TEXT,
		education TEXT
	);`
	_, err := db.Exec(userTable)
	if err != nil {
		return err
	}

	photoTable := `CREATE TABLE IF NOT EXISTS photos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		url TEXT,
		position INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	);`
	_, err = db.Exec(photoTable)
	if err != nil {
		return err
	}

	return nil
}
