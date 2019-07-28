package meta

import (
	"fmt"
	"strings"
	"time"

	"github.com/tingxin/bingo/common/value"
	"github.com/tingxin/bingo/setting"
	"github.com/tingxin/go-utility/db/mysql"
)

var (
	metaCache        map[string]map[string]value.ValueType
	request          chan *requestEntity
	concurrent       = 2
	mysqlTypeMapping map[string]value.ValueType
)

const (
	sqlMetaFormat = `
	SHOW FULL COLUMNS FROM %s
	`
)

func init() {
	initDBTypeMapping()
}

type requestEntity struct {
	Key  string
	Resp chan map[string]value.ValueType
}

// Run used to start get db meta service
func Run() {
	metaCache = make(map[string]map[string]value.ValueType)
	request = make(chan *requestEntity, concurrent*4)
	for i := 0; i < concurrent; i++ {
		go cacheMetaInBackgournd()
	}
}

// GetMeta used to get table fileds meta data
func GetMeta(db, table string) map[string]value.ValueType {
	key := fmt.Sprintf("%s.%s", db, table)
	req := &requestEntity{
		Key:  key,
		Resp: make(chan map[string]value.ValueType),
	}
	request <- req
	v := <-req.Resp
	return v
}

func cacheMetaInBackgournd() {
	for {
		req := <-request
		if v, ok := metaCache[req.Key]; ok {
			req.Resp <- v
		} else {
			m, err := GetTableFieldsMeta(req.Key)
			if err != nil {
				req.Resp <- nil
			} else {
				metaCache[req.Key] = m
				req.Resp <- m
			}
		}
		time.Sleep(time.Millisecond * 10)
	}
}

// GetTableFieldsMeta used to get table fileds meta data
func GetTableFieldsMeta(table string) (map[string]value.ValueType, error) {
	sql := fmt.Sprintf(sqlMetaFormat, table)
	conn, err := mysql.GetConn(setting.DataDBConnStr)
	if err != nil {
		return nil, err
	}

	gen := mysql.FetchRawGenerator(conn, sql)

	meta := make(map[string]value.ValueType)

	for rawRow := range gen {
		if rawRow.Err != nil {
			return nil, rawRow.Err
		}
		row := rawRow.Data
		name := string(row[0])
		rawfieldType := string(row[1])
		fieldType := getFieldType(rawfieldType)
		meta[name] = fieldType
	}
	return meta, nil
}

func getFieldType(typeStr string) value.ValueType {
	pos := strings.Index(typeStr, "(")
	if pos > 0 {
		typeStr = typeStr[0:pos]
	}
	if v, ok := mysqlTypeMapping[typeStr]; ok {
		return v
	}
	return value.String
}

func initDBTypeMapping() {
	mysqlTypeMapping = map[string]value.ValueType{
		// 整数类型
		"bit":       value.Int8,
		"tinyint":   value.Int8,
		"smallint":  value.Int16,
		"mediumint": value.Int32,
		"int":       value.Int,
		"bigint":    value.Int64,
		"bool":      value.Boolen,
		// 浮点数类型
		"float":   value.Float32,
		"double":  value.Float64,
		"decimal": value.Float64,
		// 字符串类型
		"char":       value.String,
		"varchar":    value.String,
		"tinytext":   value.String,
		"text":       value.String,
		"mediumtext": value.String,
		"longtext":   value.String,
		"tinyblob":   value.String,
		"blob":       value.String,
		"mediumblob": value.String,
		"longblob":   value.String,
		// 日期类型
		"datetime":  value.DateTime,
		"date":      value.Date,
		"timestamp": value.TimeStamp,
		"time":      value.DateTime,
		"year":      value.Date,
		// 其他类型
		"enum":               value.Enum,
		"binary":             value.String,
		"varbinary":          value.String,
		"set":                value.String,
		"geometry":           value.String,
		"point":              value.String,
		"multipoint":         value.String,
		"linestring":         value.String,
		"multilinestring":    value.String,
		"polygon":            value.String,
		"geometrycollection": value.String,
	}
}
