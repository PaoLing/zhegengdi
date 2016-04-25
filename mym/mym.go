package mym

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"
)

const (
	TagRequire = "require"
	TagDefault = "default"
	PrimaryKey = "primary_key"
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

func getExtraTag(f reflect.StructField, tag string) bool {
	return f.Tag.Get(tag) == "true"
}

func (my *MyM) CreateTable(model interface{}) error {
	return nil
}

func IsZeroedValue(v interface{}) bool {
	switch f := v.(type) {
	case string:
		return f == ""
	case bool:
		return !f
	case int, int8, int16, int32, int64:
		return f == 0
	case uint, uint8, uint16, uint32, uint64:
		return f == 0
	case float32, float64:
		return f == 0.0
	default:
		return false
	}
}

// Insert a row into Database.
func (my *MyM) Insert(model interface{}) (err error) {
	v, t, _ := CheckDestValid(model)

	tableName, _ := GetTableName(model)
	var SQLInsertHead string = SQLSelectStart(tableName)
	var SQLInsertTail string

	for i, n := 0, t.NumField(); i < n; i++ {
		f := t.Field(i)
		vf := v.FieldByName(f.Name)
		if vf.CanInterface() {
			name := strings.ToLower(f.Name)
			if name != "id" {
				SQLInsertHead = SQLInsertHead + name + ","
			}
		}
	}
	SQLInsertHead2 := []rune(SQLInsertHead)
	SQLInsertHead = string(SQLInsertHead2[:len(SQLInsertHead2)-1])

	for i, n := 0, v.NumField(); i < n; i++ {
		f := v.Field(i)
		if f.CanInterface() {
			value := GetKindValue(f)
			if IsZeroedValue(value) {
				value = " "
			}
			ft := t.Field(i)
			if strings.ToLower(ft.Name) != "id" {
				SQLInsertTail = SQLInsertTail + "\"" + value + "\"" + ","
			}
		}
	}
	SQLInsertTail2 := []rune(SQLInsertTail)
	SQLInsertTail = string(SQLInsertTail2[:len(SQLInsertTail2)-1])

	SQLInsert := SQLInsertHead + ") VALUES (" + SQLInsertTail + ");"
	fmt.Println(SQLInsert)

	insertStmt, err := opened.Prepare(SQLInsert)
	if err != nil {
		panic(fmt.Sprintf("Prepare insert query SQL error:%s", err.Error()))
	}
	_, err = insertStmt.Exec()
	if err != nil {
		panic(fmt.Sprintf("Insert row error:%s", err.Error()))
	} else {
		log.Print("Insert row succed")
		return nil
	}
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
		return "struct..."
	default:
		return ""
	}
}
