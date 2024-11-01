package models

import (
	"database/sql"
	_ "embed"
	"fmt"
	"log"
)

const dbFileName = "./lbrd.db"

//go:embed schema.sql
var sqlInitStatement string

func InitializeDB() error {
	db, err := ConnectToSQLiteDB()
	if err != nil {
		return err
	}
	_, err = db.Exec(sqlInitStatement)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlInitStatement)
	}
	return err
}

func ConnectToSQLiteDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbFileName)
	if err != nil {
		return nil, err
	}

	// Ping the database to ensure connectivity
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to SQLite Database!")
	return db, nil
}
