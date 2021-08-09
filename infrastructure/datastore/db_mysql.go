package datastore

import (
	"fmt"

	"go-accountbook/conf"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"golang.org/x/xerrors"
)

// NewMySQLDB returns a gorm DB connection.
func NewMySQLDB() (*gorm.DB, error) {
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		conf.App.Database.User,
		conf.App.Database.Password,
		conf.App.Database.Host,
		conf.App.Database.Port,
		conf.App.Database.Dbname,
	)

	conn, err := gorm.Open("mysql", connectionString+"?parseTime=true&loc=Asia%2FTokyo&timeout=5s")

	if nil != err {
		return conn, xerrors.Errorf("gorm open error: %w", err)
	}

	if err := conn.DB().Ping(); err != nil {
		return conn, xerrors.Errorf("gorm ping error: %w", err)
	}

	conn.LogMode(true)

	conn.Set("gorm:table_options", "ENGINE=InnoDB")

	return conn, nil
}
