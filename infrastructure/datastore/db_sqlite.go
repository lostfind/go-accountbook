package datastore

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"golang.org/x/xerrors"
)

// NewSqliteDB returns a gorm DB connection.
func NewSqliteDB() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "../../tmp/accountbook.db")

	if err != nil {
		return db, xerrors.Errorf("gorm open error: %w", err)
	}

	if err := db.DB().Ping(); err != nil {
		return db, xerrors.Errorf("gorm ping error: %w", err)
	}

	db.LogMode(true)

	return db, err
}
