package main

import (
	"fmt"

	"go-accountbook/conf"
	"go-accountbook/infrastructure/datastore"
)

func main() {
	conf.Read()

	// db, err := datastore.NewMySQLDB()
	db, err := datastore.NewSqliteDB()

	if err != nil {
		fmt.Println(err)
	}
}
