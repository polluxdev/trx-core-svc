package utils

import (
	"fmt"
	"strings"

	"github.com/polluxdev/trx-core-svc/application/global"
)

type ConditionalBuilder struct {
	Column   string
	Value    interface{}
	Logical  string
	Operator string
}

func ConstructConditionalClause(builder []*ConditionalBuilder) (string, []interface{}) {
	var (
		clauses []string
		args    []interface{}
	)

	for _, v := range builder {
		switch v.Logical {
		case "LIKE":
			clauses = append(clauses, fmt.Sprintf("%s %s %s ?", v.Operator, v.Column, v.Logical))
			args = append(args, fmt.Sprintf(global.LIKE_CONDITION, v.Value))
		case "IN":
			clauses = append(clauses, fmt.Sprintf("%s %s %s (?)", v.Operator, v.Column, v.Logical))
			args = append(args, v.Value)
		case "BETWEEN":
			clauses = append(clauses, fmt.Sprintf("%s %s %s ? AND ?", v.Operator, v.Column, v.Logical))
			args = append(args, v.Value)
		case "IS NULL", "IS NOT NULL":
			clauses = append(clauses, fmt.Sprintf("%s %s %s", v.Operator, v.Column, v.Logical))
		default:
			clauses = append(clauses, fmt.Sprintf("%s %s %s ?", v.Operator, v.Column, v.Logical))
			args = append(args, v.Value)
		}
	}

	return "1 = 1 " + strings.Join(clauses, " "), args
}

func SetDefaultClause() string {
	return "1 = 1"
}
