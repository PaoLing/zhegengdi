package mym

import (
	///"database/sql"
	"fmt"
	"testing"
)

type Zgd_Users_Table struct {
	Id          int64  `require:"true" auto_increment:"true" primary_key:"true"`
	User_name   string `require:"true"`
	User_mobile string `require:"true"`
	Password    string `require:"true"`
	holder1     string
	Email       string `require:"true" default:""`
	Nickname    string `require:"true" default:"little zhe 001"`
	Level       byte   `require:"true" default:"3"`
	Locked      uint8  `require:"true" default:"false"`
	Create_time string `require:"true" default:"zhe_user_0001"`
	Comment     string `require:"true" defaulr:""`
	holder2     bool
}

func TOpen() {
	db := DB{Driver: "mysql", DSN: "root:7756789w@/zhegengdi"}
	db.Open()
}

func TClose(t *testing.T) {
	err := Close()
	if err != nil {
		t.Error("Close the opened db Error:", err.Error())
	}
}

func TestOpen(t *testing.T) {
	TOpen()
	defer TClose(t)
}

func TestQueryAll(t *testing.T) {
	TOpen()
	defer TClose(t)

	UserModel := Zgd_Users_Table{User_name: "zhe_user_3387"}

	q, err := Q(&UserModel)
	if err != nil {
		fmt.Println(q)
	}

	var rows []interface{}
	rows, err = q.QueryAll()
	if err != nil {
		t.Error(err.Error())
	}

	// database query result rows
	for _, r := range rows {
		fmt.Println(r)
	}
}

func TestGetTableName(t *testing.T) {
	var user *Zgd_Users_Table

	_, err := GetTableName(user)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestQueryById(t *testing.T) {
	TOpen()
	defer TClose(t)

	UserModel := Zgd_Users_Table{}
	q, err := Q(&UserModel)
	if err != nil {
		fmt.Println(q)
	}

	row, err2 := q.QueryByID(2)
	if err2 != nil {
		t.Error("query by id error:", err2)
	}
	fmt.Println("found: ", row)
}
