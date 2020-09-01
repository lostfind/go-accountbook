package datastore

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lostfind/go-accountbook/conf"

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

	db, err := gorm.Open("mysql", connectionString+"?parseTime=true&loc=Asia%2FTokyo&timeout=5s")

	if nil != err {
		return db, xerrors.Errorf("gorm open error: %w", err)
	}

	if err := db.DB().Ping(); err != nil {
		return db, xerrors.Errorf("gorm ping error: %w", err)
	}

	db.LogMode(true)

	db.Set("gorm:table_options", "ENGINE=InnoDB")

	return db, nil
}
