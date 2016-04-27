package mym

import (
	"testing"
)

func TestWhere(t *testing.T) {

	/*
		User := Zgd_Users_Table{
			User_mobile: "19211111111",
			User_name:   "zhe_user_3387",
			Password:    "pass_0001",
			Nickname:    "狗蛋",
		}
	*/
	mym := NewORM()
	mym.Where("name = ?")
}
