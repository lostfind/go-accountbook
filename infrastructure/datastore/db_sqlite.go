package datastore

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func NewSqliteDB() *sql.DB {
	sqliteDatabase, _ := sql.Open("sqlite3", "../../db/accountbook.db")
	return sqliteDatabase
}

func NewSqliteTestDB() *sql.DB {
	os.Remove("../../db/accountbook_test.db")
	sqliteDatabase, _ := sql.Open("sqlite3", "../../db/accountbook_test.db")
	return sqliteDatabase
}

func NewSqliteTestEmptyDB() *sql.DB {
	os.Remove("../../db/accountbook_empty_test.db")
	sqliteDatabase, _ := sql.Open("sqlite3", "../../db/accountbook_empty_test.db")
	return sqliteDatabase
}
