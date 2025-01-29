package database

import (
	"database/sql"
	"fmt"
	"github.com/pressly/goose/v3"
	"log"
	_ "modernc.org/sqlite"
)

func New(dbName string) (*sql.DB, error) {
	db, err := getConnection(dbName)
	if err != nil {
		return nil, err
	}

	if err := createMigrations(db); err != nil {
		return nil, err
	}

	return db, nil
}

func createMigrations(db *sql.DB) error {
	if err := goose.SetDialect("sqlite3"); err != nil { // TODO: Add this a configuration
		return err
	}

	err := goose.Up(db, "./database/migrations")
	if err != nil {
		return err
	}

	return nil
}

func getConnection(dbName string) (*sql.DB, error) {
	var (
		err error
		db  *sql.DB
	)

	// Create sqlite3 database
	db, err = sql.Open("sqlite", dbName)
	if err != nil {
		return nil, fmt.Errorf("💀 failed to connect to the database: %s", err)
	}

	log.Printf("🚀 Connected Successfully to the Database: %s\n", dbName)

	return db, nil
}
