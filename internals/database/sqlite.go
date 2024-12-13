package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// InitDB initializes the SQLite database connection
func InitDB(dataSourceName string) (*sql.DB, error) {
	var err error
	db, err = sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		log.Printf("Failed to ping database: %v", err)
		return nil, err
	}

	// Enable foreign keys
	if _, err := db.Exec("PRAGMA foreign_keys = ON"); err != nil {
		log.Printf("Failed to enable foreign keys: %v", err)
		return nil, err
	}

	// Create users table if it doesn't exist
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
		email TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL,
		first_name TEXT,
		last_name TEXT,
		phone TEXT,
		address TEXT,
		image TEXT,
		role TEXT,
		status TEXT,
		verified INTEGER DEFAULT 0,
		votes INTEGER DEFAULT 0,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	if _, err := db.Exec(createTableSQL); err != nil {
		log.Printf("Failed to create users table: %v", err)
		return nil, err
	}

	log.Println("Database initialized successfully")
	return db, nil
}

// GetDB returns the database instance
func GetDB() *sql.DB {
	return db
}
