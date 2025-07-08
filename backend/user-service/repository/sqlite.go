// sqlite.go - SQLite bağlantı ve işlemleri
package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func NewSQLiteDB(dbPath string) (*sql.DB, error) {
	return sql.Open("sqlite3", dbPath)
}

// InitDatabase - Veritabanı tablolarını oluştur
func InitDatabase(db *sql.DB) error {
	// Users tablosu
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            email TEXT UNIQUE NOT NULL,
            password TEXT NOT NULL,
            bio TEXT,
            age INTEGER,
            age_range TEXT,
            distance INTEGER DEFAULT 50,
            seriousness INTEGER DEFAULT 5,
            height INTEGER,
            weight INTEGER,
            smokes BOOLEAN DEFAULT FALSE,
            drinks BOOLEAN DEFAULT FALSE,
            job TEXT,
            job_category TEXT,
            education TEXT,
            hobbies TEXT, -- JSON string olarak saklanacak
            hobby_categories TEXT, -- JSON string olarak saklanacak
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
        )
    `)
	if err != nil {
		return err
	}

	// Photos tablosu
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS photos (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            user_id INTEGER NOT NULL,
            url TEXT NOT NULL,
            is_primary BOOLEAN DEFAULT FALSE,
            order_index INTEGER DEFAULT 0,
            ai_score REAL DEFAULT 0.0,
            is_verified BOOLEAN DEFAULT FALSE,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (user_id) REFERENCES users (id)
        )
    `)
	if err != nil {
		return err
	}

	// User preferences tablosu
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS user_preferences (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            user_id INTEGER UNIQUE NOT NULL,
            min_age INTEGER DEFAULT 18,
            max_age INTEGER DEFAULT 100,
            min_height INTEGER DEFAULT 150,
            max_height INTEGER DEFAULT 200,
            accepts_smokers BOOLEAN DEFAULT TRUE,
            accepts_drinkers BOOLEAN DEFAULT TRUE,
            min_seriousness INTEGER DEFAULT 1,
            max_seriousness INTEGER DEFAULT 10,
            FOREIGN KEY (user_id) REFERENCES users (id)
        )
    `)
	if err != nil {
		return err
	}

	return nil
}
