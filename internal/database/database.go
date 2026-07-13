package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB(dbPath string) error {
	dir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create database directory: %w", err)
	}

	var err error
	DB, err = sql.Open("sqlite", dbPath+"?_journal_mode=WAL")
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	if err = DB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	return initTables()
}

func initTables() error {
	migrations := []string{
		`CREATE TABLE IF NOT EXISTS projects (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			description TEXT,
			password_hash TEXT NOT NULL,
			salt TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			storage_path TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS issues (
			id TEXT PRIMARY KEY,
			project_id TEXT NOT NULL,
			title TEXT NOT NULL,
			description TEXT,
			status TEXT DEFAULT 'pending',
			contact_person TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE
		)`,
		`CREATE TABLE IF NOT EXISTS documents (
			id TEXT PRIMARY KEY,
			issue_id TEXT,
			project_id TEXT NOT NULL,
			title TEXT NOT NULL,
			file_path TEXT NOT NULL,
			encrypted_name TEXT NOT NULL,
			file_type TEXT,
			file_size INTEGER,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (issue_id) REFERENCES issues(id) ON DELETE SET NULL,
			FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE
		)`,
		`CREATE TABLE IF NOT EXISTS tags (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL UNIQUE,
			color TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS issue_documents (
			issue_id TEXT NOT NULL,
			document_id TEXT NOT NULL,
			PRIMARY KEY (issue_id, document_id),
			FOREIGN KEY (issue_id) REFERENCES issues(id) ON DELETE CASCADE,
			FOREIGN KEY (document_id) REFERENCES documents(id) ON DELETE CASCADE
		)`,
		`CREATE TABLE IF NOT EXISTS document_tags (
			document_id TEXT NOT NULL,
			tag_id TEXT NOT NULL,
			PRIMARY KEY (document_id, tag_id),
			FOREIGN KEY (document_id) REFERENCES documents(id) ON DELETE CASCADE,
			FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE
		)`,
		`CREATE TABLE IF NOT EXISTS settings (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			key TEXT NOT NULL UNIQUE,
			value TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS activity_logs (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			action TEXT NOT NULL,
			entity_type TEXT,
			entity_id TEXT,
			details TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)`,
	}

	for _, migration := range migrations {
		if _, err := DB.Exec(migration); err != nil {
			return fmt.Errorf("failed to execute migration: %w", err)
		}
	}

	return nil
}

func Close() {
	if DB != nil {
		DB.Close()
	}
}
