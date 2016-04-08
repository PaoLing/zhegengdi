package zhegengdi

import (
	"testing"
)

const RAWURL1 string = "https://github.com/go-sql-driver/mysql/blob/master/dsn_test.go"

func TestCheckErr(t *testing.T) {
	_, err := ParseURL(RAWURL1)

	if err != nil {
		t.Error(err.Error())
	}
}
