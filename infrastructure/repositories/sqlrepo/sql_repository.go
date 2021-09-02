package sqlrepo

import (
	"database/sql"
)

type sqlRepo struct {
	db *sql.DB
}

func NewSqlRepository(db *sql.DB) *sqlRepo {
	return &sqlRepo{db}
}
