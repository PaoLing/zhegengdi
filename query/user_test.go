package query

import (
	"database/sql"
	"fmt"
	"testing"
)

var (
	driverName     = "mysql"
	dataSourceName = "root:7756789w@/zhegengdi"
	queryAll       = "SELECT * FROM zgd_users_table WHERE user_name=?"
	username       = "zhe_user_3387"
)

type User struct {
	user_id     int64
	user_name   string
	user_mobile string
	password    string
	email       sql.NullString
	nickname    string
	level       byte
	locked      bool
	create_time string
	comment     sql.NullString
}

func TestQueryUser(t *testing.T) {
	db := QueryUser(driverName, dataSourceName)

	defer func() {
		DBClose <- true
	}()

	stmtResults, err := db.Prepare(queryAll)
	defer stmtResults.Close()

	var r *User = new(User)

	fmt.Println(r)

	arr := []interface{}{
		&r.user_id, &r.user_name, &r.user_mobile, &r.password, &r.email,
		&r.nickname, &r.level, &r.locked, &r.create_time, &r.comment,
	}

	err = stmtResults.QueryRow(username).Scan(arr...)

	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println(r)
}
