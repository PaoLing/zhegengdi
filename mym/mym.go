package mym

import (
	"database/sql"
	"errors"
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
//
// mym := NewORM()
//
func NewORM() *MyM {
	mym := new(MyM)
	mym.db = opened
	mym.cond = &Condition{}

	return mym
}

type MyM struct {
	db    *sql.DB
	model interface{} // store the model
	Table string      // table name
	cond  *Condition
}

func (mym *MyM) RegisteModel(model interface{}) {

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
func (my *MyM) Insert(model interface{}) (int64, error) {
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
				SQLInsertTail = SQLInsertTail + "'" + value + "'" + ","
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
	rest, err := insertStmt.Exec()
	if err != nil {
		panic(fmt.Sprintf("Insert row error:%s", err.Error()))
	}

	id, err := rest.LastInsertId()
	if err != nil {
		return id, err
	}

	return storeIDToModel(id, model)
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

func (my *MyM) Update(model interface{}) (int64, error) {
	v, t, _ := CheckDestValid(model)
	setSQL, id := GenSetSQL(v, t)
	tableName, _ := GetTableName(model)

	SQLUpdate := fmt.Sprintf("UPDATE %s SET %s WHERE id=%d", tableName, setSQL, id)
	fmt.Println(SQLUpdate)
	stmtUpd, err := opened.Prepare(SQLUpdate)
	if err != nil {
		return id, err
	}
	_, err = stmtUpd.Exec()
	if err != nil {
		return id, err
	}
	return id, nil
}

// storeIDToModel store Database Id to model.
func storeIDToModel(id int64, model interface{}) (int64, error) {
	v, _, _ := CheckDestValid(model)
	f := v.FieldByName("Id")
	if v.CanSet() {
		f.SetInt(id)
		return id, nil
	}
	return id, errors.New("Struct's Id can't set")
}

func GetSingleRowId(v reflect.Value) (int64, error) {
	fieldId := v.FieldByName("Id")
	if fieldId.IsValid() {
		return fieldId.Int(), nil
	}
	return 0, errors.New(fmt.Sprintf("%v field Id not exist.", v))
}

// GenSetSQL generate part of UPDATE SQL.
func GenSetSQL(v reflect.Value, t reflect.Type) (string, int64) {
	var SQL string
	var fields = make(map[string]string)

	id, err := GetSingleRowId(v)
	if err != nil {
		log.Fatal(err)
	}

	numField := t.NumField()
	for i := 0; i < numField; i++ {
		f := t.Field(i)
		vf := v.Field(i)
		if vf.CanInterface() {
			value := GetKindValue(vf)
			fname := strings.ToLower(f.Name)
			if fname == "id" {
				continue
			}
			value = "'" + value + "'"
			fields[fname] = value
		}
	}

	for k, v := range fields {
		SQL = SQL + k + "=" + v + ","
	}
	return SQL[:len(SQL)-1], id
}

// Delete delete a single row by id column.
func (mym *MyM) Delete(model interface{}) bool {
	v, _, _ := CheckDestValid(model)
	id, err := GetSingleRowId(v)
	if err != nil {
		log.Fatal(err)
	}
	tableName, _ := GetTableName(model)
	SQL := fmt.Sprintf("DELETE FROM %s WHERE id=?", tableName)
	stmtDelete, err := opened.Prepare(SQL)
	if err != nil {
		return false
	}
	_, err = stmtDelete.Exec(id)
	if err != nil {
		return false
	}
	return true
}

func (mym *MyM) Where(query string, params ...interface{}) {
	mym.cond.Where(query, params...)
}

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
