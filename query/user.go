package query

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DBClose = make(chan bool, 1)
)

// QueryUser for query user table
func QueryUser(driverName, dataSourceName string) (db *sql.DB) {
	db, err := sql.Open(driverName, dataSourceName)

	go func() {
		if <-DBClose {
			fmt.Println("closing database connection...")
			db.Close()
		}
	}()

	if err != nil {
		DBClose <- true
		panic(err.Error())
	}

	err = db.Ping()

	if err != nil {
		DBClose <- true
		panic(err.Error())
	}

	fmt.Println("Database connected")
	return db
}
