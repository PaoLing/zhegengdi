package mym

import (
	"database/sql"
	"fmt"
	"testing"
)

type Zgd_Users_Table struct {
	User_id     int64  `require:"true" action:"auto"`
	User_name   string `require:"true" default:"zhe_user_0001"`
	User_mobile string `require:"true"`
	Password    string `require:"true"`
	holder1     string
	Email       sql.NullString
	Nickname    string `require:"true" default:"little zhe 001"`
	Level       byte   `require:"true" default:"3"`
	Locked      bool   `require:"true" default:"false"`
	Create_time string `require:"true" default:"zhe_user_0001"`
	Comment     sql.NullString
	holder2     sql.NullBool
}

func TestOpen(t *testing.T) {
	db := DB{
		Driver: "mysql",
		DSN:    "root:7756789w@/zhegengdi",
	}

	db.Open()

	defer func() {
		err := Close()
		if err != nil {
			t.Error(err.Error())
		}
	}()

	UserModel := Zgd_Users_Table{User_name: "zhe_user_3387"}

	q, err := Q(&UserModel)
	if err != nil {
		fmt.Println(q)
	}

	var r interface{}
	r, err = q.QueryRows()
	if err != nil {
		t.Error(err.Error())
	}

	if rows, ok := r.([]interface{}); ok {
		for _, r := range rows {
			fmt.Println(r)
		}
	}
}

func TestGetTableName(t *testing.T) {
	var user *Zgd_Users_Table

	_, err := GetTableName(user)
	if err != nil {
		t.Error(err.Error())
	}
}
