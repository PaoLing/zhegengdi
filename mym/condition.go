package mym

import (
	"fmt"
)

type Condition struct {
	WhereSQL  string
	HasHolder bool
}

func (cond *Condition) Where(query string, params ...interface{}) *Condition {
	if len(params) == 0 {
		cond.WhereSQL = query
		cond.HasHolder = false
	}

	return cond

}
