package mym

import (
	"fmt"
	"strings"
)

type Condition struct {
	whereSQL  string
	hasHolder bool
	pms       []interface{}
}

func (cond *Condition) Where(query string, params ...interface{}) *Condition {
	if len(params) == 0 {
		cond.whereSQL = query
		cond.hasHolder = false
	} else {
		cond.pms = params
		cond.hasHolder = true
	}

	return cond
}
