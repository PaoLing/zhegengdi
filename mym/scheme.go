package mym

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"reflect"
	"strings"
)

var (
	opened *sql.DB
)

// DB stroied the connection informations.
type DB struct {
	Driver   string
	DSN      string
	MaxConns int32
}

func (db *DB) Open() {
	mdb, err := sql.Open(db.Driver, db.DSN)
	if err != nil {
		panic(err)
	}
	opened = mdb
}

func Open(driverName, dataSourceName string) {
	db := &DB{
		Driver:   driverName,
		DSN:      dataSourceName,
		MaxConns: 1024,
	}
	db.Open()
}

// Close closes the the database, releasing any open resources.
func Close() error {
	return opened.Close()
}

// Q receive a user-defined Database table struct, check out it's
func Q(model interface{}) (q *Query, err error) {
	v, _, _ := CheckDestValid(model)

	tableName, _ := GetTableName(model)

	NumField := v.NumField()
	fields := make([]interface{}, 0, NumField)

	for i := 0; i < NumField; i++ {
		field := v.Field(i)
		if field.CanInterface() {
			fields = append(fields, field.Addr().Interface())
		}
	}

	q = &Query{
		Arch:      v.Addr().Interface(),
		Results:   fields,
		TableName: tableName,
	}
	return q, nil
}

// Query store the table model and the result lists that can be used by sql.Scan.
type Query struct {
	Arch      interface{}
	Results   []interface{}
	TableName string
}

// QueryAll query all rows.
func (q *Query) QueryAll() (allRows []interface{}, err error) {
	SQLQueryAll := fmt.Sprintf("SELECT * FROM %s", q.TableName)

	rows, err := opened.Query(SQLQueryAll)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	allRows = make([]interface{}, 0)
	for rows.Next() {
		if err := rows.Scan(q.Results...); err != nil {
			log.Fatal("Scan error: ", err)
		}
		row := CopyRow(q.Arch)
		allRows = append(allRows, row)
	}
	return allRows, nil
}

// QueryByID query one row by given ID.
func (q *Query) QueryByID(id int32) (interface{}, error) {
	SQLQueryID := fmt.Sprintf("SELECT * FROM %s WHERE id=%d", q.TableName, id)
	err := opened.QueryRow(SQLQueryID).Scan(q.Results...)
	if err != nil {
		return nil, err
	}
	return CopyRow(q.Arch), nil
}

func CopyRow(arch interface{}) interface{} {
	t := reflect.ValueOf(arch).Elem()
	return t.Interface()
}

// CheckDestValid check out the model valid.
func CheckDestValid(model interface{}) (reflect.Value, reflect.Type, error) {
	v := reflect.ValueOf(model)

	// if model not a pointer, painc.
	if v.Kind() != reflect.Ptr {
		panic(fmt.Sprintf("%v not a pointer", v.Type()))
	}
	// if model's value is nil, painc.
	if v.IsNil() {
		panic(fmt.Sprintf("(%v %v) Must have no-nil value", v, v.Type()))
	}

	t := reflect.TypeOf(model).Elem()

	return v.Elem(), t, nil
}

// GetTableName get the table name.
func GetTableName(arch interface{}) (tableName string, err error) {
	t := reflect.TypeOf(arch)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	tableName = t.Name()
	return strings.ToLower(tableName), nil
}
