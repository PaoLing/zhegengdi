package mym

import (
	"database/sql"
	"fmt"
	"reflect"
)

// NewORM allocate a value that type is *MyM and return it.
func NewORM() *MyM {
	mym := new(MyM)
	mym.db = opened

	return mym
}

type MyM struct {
	db *sql.DB
}

// Insert a row into Database.
func (my *MyM) Insert(model interface{}) (err error) {
	v, t, _ := CheckDestValid(model)

	tableName, _ := GetTableName(model)
	SQLp1 := SQLSelectStart(tableName)

	for i, n := 0, t.NumField(); i < n; i++ {
		f := t.Field(i)
		vf := v.FieldByName(f.Name)
		if vf.CanInterface() {
			SQLp1 = SQLp1 + f.Name + ", "
		}
	}
	SQLp1 = SQLp1 + ") VALUES ("
	fmt.Println(SQLp1)

	for i, n := 0, v.NumField(); i < n; i++ {
		f := v.Field(i)
		if f.CanInterface() {
			value := GetKindValue(f)
			fmt.Println(value)
		}
	}
	return nil
}

func SQLSelectStart(tableName string) string {
	return fmt.Sprintf("INSERT INTO %s (", tableName)
}

func GetKindValue(f reflect.Value) string {
	switch f.Kind() {
	case reflect.String:
		return fmt.Sprint(f.String())
	case reflect.Bool:
		return fmt.Sprint(f.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprint(f.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprint(f.Uint())
	case reflect.Float32, reflect.Float64:
		return fmt.Sprint(f.Float())
	case reflect.Struct:
		return "struct"
	case reflect.Invalid:
		fmt.Println("invalid...")
		return "invalid"
	default:
		fmt.Println("default...")
		return "defult"
	}
}