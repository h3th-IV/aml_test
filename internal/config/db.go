package config

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type Database struct {
	DB *sql.DB
}

func ConnectDatabase() (*Database, error) {
	db, err := sql.Open("sqlite", "./aml_test.db")
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	if err = CreateSchema(db); err != nil {
		return nil, err
	}
	return &Database{DB: db}, nil
}

func CreateSchema(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		name TEXT NOT NULL, 
		email TEXT NOT NULL, 
		gender TEXT NOT NULL, 
		dob TEXT NOT NULL,
		address TEXT NOT NULL
	)`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) Close() error {
	return d.DB.Close()
}
