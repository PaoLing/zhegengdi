package query

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// QueryUser for query user table
func QueryUser() {
	db, err := sql.Open("mysql", "root:7756789w@/zhegengdi")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Database connected")
}
