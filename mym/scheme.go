package mym

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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

type Query struct {
	T       interface{}
	Results []interface{}
}

func Q(model interface{}) (q *Query, err error) {
	CheckDestValid(model)

	q = &Query{
		T: model,
	}
	return q, nil
}

// CheckDestValid check out the model valid, if not a pointer or
// value is nil, panic.
func CheckDestValid(m interface{}) {
	// if a destination type not a point, painc.
	t := reflect.TypeOf(model)
	if t.Kind() != reflect.Ptr {
		panic(fmt.Sprintf("mym.Q: Type <%v> not a pointer", t.Name()))
	}

	tv := reflect.ValueOf(model)
	if tv.IsNil() {
		panic(fmt.Sprintf("Must be a no-nil value"))
	}
}

func (q *Query) QueryAll() {

}

func (q *Query) QueryByID(id int) {

}

// GetTableName get the table name.
func GetTableName(arch interface{}) string {
	t := reflect.TypeOf(arch)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	table := t.Name()
	return strings.ToLower(table)
}
