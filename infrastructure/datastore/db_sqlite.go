package datastore

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func NewSqliteDB() *sql.DB {
	sqliteDatabase, err := sql.Open("sqlite3", os.Getenv("SQLITE_PROD_DB_FILE"))
	if err != nil {
		log.Fatal(err)
	}
	return sqliteDatabase
}

func NewSqliteTestDB() *sql.DB {
	os.Remove(os.Getenv("SQLITE_TEST_DB_FILE"))
	sqliteDatabase, _ := sql.Open("sqlite3", os.Getenv("SQLITE_TEST_DB_FILE"))
	return sqliteDatabase
}

func NewSqliteTestEmptyDB() *sql.DB {
	os.Remove(os.Getenv("SQLITE_EMPTY_DB_FILE"))
	sqliteDatabase, _ := sql.Open("sqlite3", os.Getenv("SQLITE_EMPTY_DB_FILE"))
	return sqliteDatabase
}
