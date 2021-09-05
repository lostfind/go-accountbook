package datastore

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func NewSqliteDB() *sql.DB {
	sqliteDatabase, err := sql.Open("sqlite3", "/Users/dwkim/development/accountbook/go-accountbook/db/accountbook.db")
	if err != nil {
		log.Fatal(err)
	}
	return sqliteDatabase
}

func NewSqliteTestDB() *sql.DB {
	os.Remove("/Users/dwkim/development/accountbook/go-accountbook/db/accountbook_test.db")
	sqliteDatabase, _ := sql.Open("sqlite3", "/Users/dwkim/development/accountbook/go-accountbook/db/accountbook_test.db")
	return sqliteDatabase
}

func NewSqliteTestEmptyDB() *sql.DB {
	os.Remove("/Users/dwkim/development/accountbook/go-accountbook/db/accountbook_empty_test.db")
	sqliteDatabase, _ := sql.Open("sqlite3", "/Users/dwkim/development/accountbook/go-accountbook/db/accountbook_empty_test.db")
	return sqliteDatabase
}
