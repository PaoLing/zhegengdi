package mym

import (
	"testing"
)

type User struct {
	user_id     int64  `require:"true" action:"auto"`
	user_name   string `require:"true" default:"zhe_user_0001"`
	user_mobile string `require:"true"`
	password    string `require:"true"`
	email       sql.NullString
	nickname    string `require:"true" default:"little zhe 001"`
	level       byte   `require:"true" default:"3"`
	locked      bool   `require:"true" default:"false"`
	create_time string `require:"true" default:"zhe_user_0001"`
	comment     sql.NullString
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

	UserModel := &User{user_name: "zhe_user_3387"}
	q := Q(UserModel)

}
