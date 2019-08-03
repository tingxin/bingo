package data

import (
	"database/sql"
	"fmt"

	"github.com/kataras/iris"
	"github.com/tingxin/bingo/common/meta"
	"github.com/tingxin/bingo/common/value"
	"github.com/tingxin/bingo/service/data/model"
	command "github.com/tingxin/bingo/service/data/sql"
	"github.com/tingxin/bingo/setting"
	"github.com/tingxin/go-utility/db/mysql"
	"github.com/tingxin/go-utility/log"
	"github.com/tingxin/go-utility/tools/conv"

	m "github.com/tingxin/bingo/model"
)

func (p *service) query(ctx iris.Context) {
	log.INFO.Printf("begin query %s", "data")
	offset := ctx.Params().GetIntDefault("offset", 0)
	pageCount := ctx.Params().GetIntDefault("count", 20)

	condition := model.Condition{}
	err := ctx.ReadJSON(&condition)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}
	metas := meta.GetMeta(condition.DB, condition.Table)
	command, err := command.BuildSelect(&condition, offset, pageCount, metas)

	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}

	log.DEBUG.Printf("build sql %s", command)

	var errs error
	var code = iris.StatusOK
	for {
		conn, err := mysql.GetConn(setting.DataDBConnStr)
		if err != nil {
			errs = fmt.Errorf("failed to connect the db server due to %v", err)
			code = iris.StatusInternalServerError
			break
		}

		rows, err := mysql.FetchRawWithConn(conn, command)
		if err != nil {
			errs = fmt.Errorf("failed to excuete sql %s  due to %v", command, err)
			code = iris.StatusInternalServerError
			break
		}
		count := len(rows)
		result := make([][]interface{}, count, count)
		columnsCount := len(condition.Fields)
		for i, row := range rows {
			rowt := make([]interface{}, columnsCount, columnsCount)
			for t, item := range row {
				rowt[t] = getFiledValue(t, item, &condition, metas)
			}
			result[i] = rowt
		}
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(result)
		break
	}
	if errs != nil {
		ctx.StatusCode(code)
		errStr := fmt.Sprintf("query error : %v", errs)
		ctx.JSON(m.NewResponse(errStr, nil))
	}
}

func getFiledValue(rowIndex int, rawData sql.RawBytes, condition *model.Condition, metas map[string]value.ValueType) interface{} {
	field := condition.Fields[rowIndex]
	valueType := metas[field.Key]
	if valueType == value.Int {
		return conv.ToIntFromBytes(rawData)
	} else if valueType == value.Int32 {
		return conv.ToIntFromBytes(rawData)
	} else if valueType == value.Int64 {
		return conv.ToInt64FromBytes(rawData)
	} else if valueType == value.Float32 {
		return conv.ToFloat32FromBytes(rawData)
	} else if valueType == value.Float64 {
		return conv.ToFloat64FromBytes(rawData)
	} else if valueType == value.Date {
		return conv.ToStrFromBytes(rawData)
	} else {
		return string(rawData)
	}
}
