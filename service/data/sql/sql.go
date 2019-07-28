package sql

import (
	"fmt"

	"github.com/tingxin/bingo/common/value"
	"github.com/tingxin/bingo/service/data/model"
	gosqler "github.com/tingxin/go-sqler"
	"github.com/tingxin/go-utility/log"
)

var (
	retryTimes = 3
	operators  map[int8]string
)

func init() {
	initOperators()
}

// BuildSelect user to build select query
func BuildSelect(condition *model.Condition, offset, count int, metas map[string]value.ValueType) (string, error) {
	query := gosqler.NewSelect()
	format := "%s.%s"
	// tables := make(map[string]bool)
	mainTable := ""
	for _, field := range condition.Fields {
		key := fmt.Sprintf(format, field.Table, field.Key)
		query.Select(key)
		mainTable = field.Table
	}

	query.From(mainTable)
	query.Offset(offset)
	query.Limit(count)

	for _, filter := range condition.Filters {
		vtype := metas[filter.Key]
		key := fmt.Sprintf(format, filter.Table, filter.Key)
		op := getOperatorStr(filter.Operator)
		if isArrayOperator(filter.Operator) {
			attachFilter(query, filter, vtype)
		} else {
			query.Where(key, op, filter.Value)
		}

	}

	sql := query.String()
	log.INFO.Printf("Build sql %s", sql)
	return sql, nil
}

func getOperatorStr(index int8) string {
	if v, ok := operators[index]; ok {
		return v
	}
	return "="
}

func isArrayOperator(op int8) bool {
	return op == 7 || op == 70
}

func attachFilter(query gosqler.Select, filter *model.Filter, vtype value.ValueType) {
	array, ok := filter.Value.([]interface{})
	if ok {
		format := "%s.%s"
		count := len(array)

		key := fmt.Sprintf(format, filter.Table, filter.Key)
		op := getOperatorStr(filter.Operator)

		if vtype == value.Int {
			t := make([]int, count, count)
			for i := 0; i < count; i++ {
				t[i] = array[i].(int)
			}
			query.Where(key, op, t)
		} else if vtype == value.Int32 {
			t := make([]int32, count, count)
			for i := 0; i < count; i++ {
				t[i] = array[i].(int32)
			}
			query.Where(key, op, t)
		} else if vtype == value.Int64 {
			t := make([]int64, count, count)
			for i := 0; i < count; i++ {
				t[i] = array[i].(int64)
			}
			query.Where(key, op, t)
		} else if vtype == value.Float32 {
			t := make([]float32, count, count)
			for i := 0; i < count; i++ {
				t[i] = array[i].(float32)
			}
			query.Where(key, op, t)
		} else if vtype == value.Float64 {
			t := make([]float64, count, count)
			for i := 0; i < count; i++ {
				t[i] = array[i].(float64)
			}
			query.Where(key, op, t)
		} else {
			t := make([]string, count, count)
			for i := 0; i < count; i++ {
				t[i] = array[i].(string)
			}
			query.Where(key, op, t)
		}
	}
}

func initOperators() {
	operators = make(map[int8]string)
	operators[1] = "="
	operators[2] = ">"
	operators[3] = "<"
	operators[4] = ">="
	operators[5] = "<="
	operators[6] = "LIKE"
	operators[7] = "IN"
	operators[8] = "IS NOT"
	operators[9] = "BETWEEN"
	operators[10] = "!= "
	operators[60] = "NOT LIKE"
	operators[70] = "NOT IN"
	operators[80] = "IS"
	operators[90] = "NOT BETWEEN"
}
