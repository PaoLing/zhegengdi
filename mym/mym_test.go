package mym

import (
	"testing"
)

func TestInsert(t *testing.T) {
	TOpen()
	defer TClose(t)

	User := Zgd_Users_Table{
		User_mobile: "19223459866",
		User_name:   "zhe_user_3387",
		Password:    "pass_0001",
		Nickname:    "段-青",
	}

	mym := NewORM()
	err := mym.Insert(&User)
	if err != nil {
		t.Error(err)
	}

}
