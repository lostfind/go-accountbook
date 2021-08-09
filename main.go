package main

import (
	"fmt"

	"github.com/lostfind/go-accountbook/conf"
	"github.com/lostfind/go-accountbook/infrastructure/datastore"
)

func main() {
	conf.Read()

	// db, err := datastore.NewMySQLDB()
	db, err := datastore.NewSqliteDB()

	if err != nil {
		fmt.Println(err)
	}
}
