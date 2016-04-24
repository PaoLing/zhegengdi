package mym

import (
	"testing"
	"time"
)

func TestInsert(t *testing.T) {
	TOpen()
	defer TClose(t)

	User := Zgd_Users_Table{
		User_mobile: "19211111111",
		User_name:   "zhe_user_3387",
		Password:    "pass_0001",
		Nickname:    "狗蛋",
		Create_time: time.Now().String(),
	}

	mym := NewORM()
	err := mym.Insert(&User)
	if err != nil {
		t.Error(err)
	}
}
