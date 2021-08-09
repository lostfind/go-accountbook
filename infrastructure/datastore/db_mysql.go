package datastore

import (
	"database/sql"
	"fmt"
	"go-accountbook/conf"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLDB() (*sql.DB, error) {
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		conf.App.Database.User,
		conf.App.Database.Password,
		conf.App.Database.Host,
		conf.App.Database.Port,
		conf.App.Database.Dbname,
	)

	conn, err := sql.Open("mysql", connectionString+"?parseTime=true&loc=Asia%2FTokyo&timeout=5s")

	if nil != err {
		return conn, fmt.Errorf("sql open error: %w", err)
	}

	if err := conn.Ping(); err != nil {
		return conn, fmt.Errorf("sql ping error: %w", err)
	}

	return conn, nil
}
