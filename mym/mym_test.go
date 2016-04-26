package mym

import (
	"fmt"
	"testing"
	"time"
)

func TestInsert(t *testing.T) {
	TOpen()
	defer TClose(t)
	now := time.Now()
	createTime := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	User := Zgd_Users_Table{
		User_mobile: "19211111111",
		User_name:   "zhe_user_3387",
		Password:    "pass_0001",
		Nickname:    "狗蛋",
		Create_time: createTime,
	}

	mym := NewORM()
	// Insert testing.
	id, err := mym.Insert(&User)
	if err != nil {
		t.Error(fmt.Sprintf("Insert Error: %s id: %d", err, id))
	}

	// Update testing
	User.User_name = "new_name_1"
	User.Nickname = "豆芽"

	_, err = mym.Update(&User)
	if err != nil {
		t.Error("Update error: ", err)
	}
}

func TestIsZeroedValue(t *testing.T) {
	values := []interface{}{0, false, "", 0.0, 1}
	for _, v := range values {
		r := IsZeroedValue(v)
		fmt.Println(r)
	}
}
