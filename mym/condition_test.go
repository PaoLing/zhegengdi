package mym

import (
	"testing"
)

/*
fake usage:

db := NewOrm()
var Users []User
db.Where("name = ?", "bob").Exec(&Users)
db.Where("name = ?", "bob").Exec(&Users)
db.Where("name = ? AND age = ?", "bob", 20).Exec(&Users)
db.Where("name = ? AND age = ?", "bob", 20).Filter("id", "profile").Exec(&Users)
db.Complex().Where("age").LessThan(38).And("job").Equal("cook").Binary().LimitEnd(100).Exec(&User)
*/

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
	mym.Where("name = ?", "zhg")
}
