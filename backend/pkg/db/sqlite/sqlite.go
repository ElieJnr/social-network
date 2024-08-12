package sqlite

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	db *sql.DB
}

func (db *DB) Prepare(query string) (*sql.Stmt, error) {
	return db.db.Prepare(query)
}

func NewDatabase() (*DB, error) {
	

	db, err := sql.Open("sqlite3", "./pkg/db/sqlite/repository/donnees.db")
	if err != nil {
		fmt.Println("ok")
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		db.Close() // Ensure the database connection is closed if Ping fails
		return nil, err
	}
	err = upDatabase(db)
	if err != nil {
		fmt.Println("Migration error:", err)
	}
	return &DB{db: db}, nil
}

func (d *DB) Close() {
	err := d.db.Close()
	if err != nil {
		fmt.Println("Error closing the database:", err)
	}
}

func (d *DB) GetDB() *sql.DB {
	return d.db
}

func upDatabase(db *sql.DB) error {
	migrationsDir := "pkg/db/migrations/sqlite/"
	err := filepath.Walk(migrationsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), "table.up.sql") {
			// Exécute le fichier de migration
			_, err := db.Exec(FileToString(path))
			if err != nil {
				fmt.Printf("Error executing migration %s: %v\n", path, err)
				return err
			}
		}
		
		return nil
		
	})
	return err
}

func FileToString(filename string) string {
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Could not read file:", err)
		return ""
	}
	return string(file)
}
