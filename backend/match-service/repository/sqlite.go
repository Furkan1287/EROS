// sqlite.go - SQLite bağlantı ve işlemleri (Match Service)
package repository

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

func NewSQLiteDB(dbPath string) (*sql.DB, error) {
    return sql.Open("sqlite3", dbPath)
}

// InitMatchDatabase - Match service veritabanı tablolarını oluştur
func InitMatchDatabase(db *sql.DB) error {
    // Matches tablosu
    _, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS matches (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            user1_id INTEGER NOT NULL,
            user2_id INTEGER NOT NULL,
            match_type TEXT NOT NULL,
            status TEXT DEFAULT 'active',
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            expires_at DATETIME,
            FOREIGN KEY (user1_id) REFERENCES users (id),
            FOREIGN KEY (user2_id) REFERENCES users (id)
        )
    `)
    if err != nil {
        return err
    }

    // Swipes tablosu
    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS swipes (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            user_id INTEGER NOT NULL,
            target_id INTEGER NOT NULL,
            direction TEXT NOT NULL,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (user_id) REFERENCES users (id),
            FOREIGN KEY (target_id) REFERENCES users (id)
        )
    `)
    if err != nil {
        return err
    }

    // Blind messages tablosu
    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS blind_messages (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            match_id INTEGER NOT NULL,
            user_id INTEGER NOT NULL,
            message TEXT NOT NULL,
            is_ai BOOLEAN DEFAULT FALSE,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (match_id) REFERENCES matches (id),
            FOREIGN KEY (user_id) REFERENCES users (id)
        )
    `)
    if err != nil {
        return err
    }

    // Date tasks tablosu
    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS date_tasks (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            match_id INTEGER NOT NULL,
            title TEXT NOT NULL,
            description TEXT NOT NULL,
            location TEXT NOT NULL,
            duration TEXT NOT NULL,
            difficulty TEXT NOT NULL,
            status TEXT DEFAULT 'pending',
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (match_id) REFERENCES matches (id)
        )
    `)
    if err != nil {
        return err
    }

    return nil
} 