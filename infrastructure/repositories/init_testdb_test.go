package repositories

import (
	"database/sql"
	"go-accountbook/infrastructure/datastore"
	"log"
)

func InitTestDB() *sql.DB {
	db := datastore.NewSqliteTestDB()
	createTestDB(db)
	insertTestDB(db)

	return db
}

func InitEmptyTestDB() *sql.DB {
	db := datastore.NewSqliteTestEmptyDB()
	createTestDB(db)

	return db
}

func createTestDB(db *sql.DB) {
	createHistoryTableSQL := `CREATE TABLE histories (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"account_id" integer,
		"category_id" integer,
		"amount" num,
		"memo" varchar,
		"date" text
	  );`

	statement, err := db.Prepare(createHistoryTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer statement.Close()

	statement.Exec()
	log.Println("test table created")
}

func insertTestDB(db *sql.DB) {
	insertHistoryTableSQL := `INSERT INTO histories (id, account_id, category_id, amount, memo, date) VALUES
	(1, 111, 1111, 50000, '테스트항목1', '2021/08/09'),
	(2, 222, 2222, 51230, '테스트항목2', '2021/08/08'),
	(3, 333, 3333, 1214, '테스트항목3', '2021/08/08');`

	statement, err := db.Prepare(insertHistoryTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer statement.Close()

	statement.Exec()
	log.Println("test table inserted")
}
